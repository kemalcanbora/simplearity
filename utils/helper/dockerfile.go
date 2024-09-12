package helper

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os/exec"
	"strings"
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

func BuildAndPushDockerImage(imageName, dockerHubUsername string) error {
	// Build the Docker image
	fmt.Println("Building Docker image...")
	fmt.Println("imageName:", imageName)
	cmd := exec.Command("docker", "build", "-t", imageName, ".")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build image: %w\nOutput: %s", err, out.String())
	}

	fmt.Println("Build output:")
	fmt.Println(out.String())

	fullImageName := fmt.Sprintf("%s/%s:latest", dockerHubUsername, imageName)
	fmt.Printf("Tagging image as %s...\n", fullImageName)
	tagCmd := exec.Command("docker", "tag", imageName, fullImageName)
	tagOut := &bytes.Buffer{}
	tagCmd.Stdout = tagOut
	tagCmd.Stderr = tagOut
	if err := tagCmd.Run(); err != nil {
		return fmt.Errorf("failed to tag image: %w\nOutput: %s", err, tagOut.String())
	}

	// Push to Docker Hub
	fmt.Printf("Pushing image to Docker Hub as %s...\n", fullImageName)
	pushCmd := exec.Command("docker", "push", fullImageName)
	pushOut := &bytes.Buffer{}
	pushCmd.Stdout = pushOut
	pushCmd.Stderr = pushOut

	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("failed to push image to Docker Hub: %w\nOutput: %s", err, pushOut.String())
	}

	fmt.Println("Push output:")
	fmt.Println(pushOut.String())

	fmt.Println("Successfully pushed image to Docker Hub")
	return nil
}

func RemoveOldDockerImages(imageName string) error {
	// List containers using the image
	cmd := exec.Command("docker", "ps", "-a", "--filter", fmt.Sprintf("ancestor=%s", imageName), "--format", "{{.ID}}")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	// Remove containers
	containerIDs := strings.Fields(out.String())
	for _, id := range containerIDs {
		removeCmd := exec.Command("docker", "rm", "-f", id)
		if err := removeCmd.Run(); err != nil {
			return fmt.Errorf("failed to remove container %s: %w", id, err)
		}
	}

	// Remove the image
	removeImageCmd := exec.Command("docker", "rmi", "-f", imageName)
	if err := removeImageCmd.Run(); err != nil {
		// Ignore errors if the image doesn't exist
		if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != 1 {
			return fmt.Errorf("failed to remove image %s: %w", imageName, err)
		}
	}

	return nil
}
