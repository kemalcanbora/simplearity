package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CreateFileFromGitHub(destination, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching file from GitHub: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	file, err := os.Create(destination)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", destination, err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", destination, err)
		os.Exit(1)
	}

	fmt.Printf("Created file: %s\n", destination)
}
