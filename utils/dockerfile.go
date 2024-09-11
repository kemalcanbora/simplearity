package utils

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"os"
)

type config struct {
	Image struct {
		Base        string   `yaml:"base"`
		Packages    []string `yaml:"packages"`
		Environment []string `yaml:"environment"`
	} `yaml:"image"`
	Install []string `yaml:"install"`
	Data    []struct {
		Mount string `yaml:"mount"`
	} `yaml:"data"`
	Code []struct {
		Path string `yaml:"path"`
		Dest string `yaml:"dest"`
	} `yaml:"code"`
	Run struct {
		Command string   `yaml:"command"`
		Args    []string `yaml:"args"`
	} `yaml:"run"`
}

func YamlToDockerfile(yamlFile string) (string, error) {
	var c config

	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return "", err
	}

	dockerfile := fmt.Sprintf("FROM %s\n\n", c.Image.Base)

	if len(c.Image.Packages) > 0 {
		dockerfile += "RUN pip install --no-cache-dir \\\n"
		for _, pkg := range c.Image.Packages {
			dockerfile += fmt.Sprintf("    %s \\\n", pkg)
		}
		dockerfile = dockerfile[:len(dockerfile)-3] + "\n\n"
	}

	if len(c.Image.Environment) > 0 {
		for _, env := range c.Image.Environment {
			dockerfile += fmt.Sprintf("ENV %s\n", env)
		}
		dockerfile += "\n"
	}

	if len(c.Install) > 0 {
		for _, env := range c.Install {
			dockerfile += fmt.Sprintf("RUN %s\n", env)
		}
		dockerfile += "\n"
	}

	if len(c.Data) > 0 {
		for _, data := range c.Data {
			dockerfile += fmt.Sprintf("VOLUME %s\n", data.Mount)
		}
		dockerfile += "\n"
	}

	if len(c.Code) > 0 {
		for _, code := range c.Code {
			dockerfile += fmt.Sprintf("COPY %s %s\n", code.Path, code.Dest)
		}
		dockerfile += "\n"
	}

	dockerfile += "WORKDIR /app\n\n"

	if c.Run.Command != "" {
		cmd := c.Run.Command
		if len(c.Run.Args) > 0 {
			cmd += " " + fmt.Sprint(c.Run.Args)
		}
		dockerfile += fmt.Sprintf("CMD %s\n", cmd)
	}

	return dockerfile, nil
}

func BuildDockerImage(imageName string, dockerfileContent string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	buildContext, err := createBuildContext(dockerfileContent)
	if err != nil {
		return fmt.Errorf("failed to create build context: %w", err)
	}

	buildOptions := types.ImageBuildOptions{
		Tags:       []string{imageName},
		Dockerfile: "Dockerfile",
		Remove:     true,
	}

	buildResponse, err := cli.ImageBuild(context.Background(), buildContext, buildOptions)
	if err != nil {
		return fmt.Errorf("failed to build image: %w", err)
	}
	defer buildResponse.Body.Close()

	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		return fmt.Errorf("failed to read build response: %w", err)
	}

	return nil
}

func createBuildContext(dockerfileContent string) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerfileHeader := &tar.Header{
		Name: "Dockerfile",
		Size: int64(len(dockerfileContent)),
		Mode: 0600,
	}

	if err := tw.WriteHeader(dockerfileHeader); err != nil {
		return nil, fmt.Errorf("failed to write Dockerfile header: %w", err)
	}

	if _, err := tw.Write([]byte(dockerfileContent)); err != nil {
		return nil, fmt.Errorf("failed to write Dockerfile content: %w", err)
	}

	return buf, nil
}

func RemoveOldDockerImages(imageName string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == imageName {
			err = cli.ContainerRemove(ctx, container.ID, containertypes.RemoveOptions{true,
				true, true})
			if err != nil {
				panic(err)
			}
		}
	}
	return err
}
