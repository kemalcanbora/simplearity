package helper

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ConvertYAML(yamlFile, outputFile string) {
	if yamlFile == "" {
		yamlFile = "simplearity.yaml"
	}
	if outputFile == "" {
		outputFile = "Dockerfile"
	}

	dockerfileContent, err := YamlToDockerfile(yamlFile)
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
}
