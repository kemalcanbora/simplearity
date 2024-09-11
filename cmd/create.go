package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"simplearity/utils"
)

func CreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Convert YAML file to Dockerfile",
		Run: func(cmd *cobra.Command, args []string) {
			yamlFile := "simplearity.yaml"
			outputFile := "Dockerfile"

			dockerfileContent, err := utils.YamlToDockerfile(yamlFile)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// Write Dockerfile
			err = ioutil.WriteFile(outputFile, []byte(dockerfileContent), 0644)
			if err != nil {
				fmt.Printf("Error writing Dockerfile: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Dockerfile created successfully: %s\n", outputFile)
		},
	}
}
