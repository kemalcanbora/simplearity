package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"simplearity/utils"
)

func BuildCmd(name string) *cobra.Command {
	var dockerfileContent string
	var err error

	return &cobra.Command{
		Use:   "build",
		Short: "Build Docker image and remove old images",
		Run: func(cmd *cobra.Command, args []string) {
			imageName := name
			dockerfilePath := "Dockerfile"
			yamlFile := "simplearity.yaml"
			// Check if Dockerfile exists
			if _, err := os.Stat(dockerfilePath); err == nil {
				content, err := ioutil.ReadFile(dockerfilePath)
				if err != nil {
					fmt.Printf("Error reading Dockerfile: %v\n", err)
					os.Exit(1)
				}
				dockerfileContent = string(content)
			} else if os.IsNotExist(err) {
				// Dockerfile doesn't exist, generate from YAML
				dockerfileContent, err = utils.YamlToDockerfile(yamlFile)
				if err != nil {
					fmt.Printf("Error generating Dockerfile from YAML: %v\n", err)
					os.Exit(1)
				}
			} else {
				// Some other error occurred
				fmt.Printf("Error checking Dockerfile: %v\n", err)
				os.Exit(1)
			}

			err = utils.RemoveOldDockerImages(imageName)
			if err != nil {
				fmt.Printf("Error removing old Docker images: %v\n", err)
				os.Exit(1)
			}

			err = utils.BuildDockerImage(imageName, dockerfileContent)
			if err != nil {
				fmt.Printf("Error building Docker image: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Docker image built successfully.")
		},
	}
}
