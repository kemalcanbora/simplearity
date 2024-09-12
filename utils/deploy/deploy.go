package deploy

import (
	"fmt"
	"os"
	"path/filepath"
	"simplearity/utils"
	"simplearity/utils/helper"
	"simplearity/utils/singularity"
)

func Deploy(hpcUsername, dockerHubUsername, imageName, jobName, mem, partition string, cpus int) {
	dockerImage := fmt.Sprintf("%s/%s:latest", dockerHubUsername, imageName)

	var err error

	client, err := utils.NewSSHClient("hpc.s.upf.edu", "kbora")
	if err != nil {
		fmt.Printf("Error creating SSH client: %v\n", err)
		return
	}
	defer client.Close()
	helper.ConvertYAML("simplearity.yaml", "Dockerfile")

	err = helper.RemoveOldDockerImages("simplearity")
	if err != nil {
		fmt.Println(err)
	}

	err = helper.BuildAndPushDockerImage("simplearity", dockerHubUsername)
	if err != nil {
		fmt.Println(err)
	}

	script := singularity.GenerateScript(jobName, mem, partition, cpus, dockerImage)

	tmpFile, err := os.CreateTemp("", "singularity_script_*.sh")
	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		return
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temporary file when we're done

	// Write the script to the temporary file
	_, err = tmpFile.WriteString(script)
	if err != nil {
		fmt.Printf("Error writing to temporary file: %v\n", err)
		return
	}
	tmpFile.Close()
	// Get the absolute path of the temporary file
	absPath, err := filepath.Abs(tmpFile.Name())
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		return
	}

	// Upload the temporary file to the server
	remoteFilePath := fmt.Sprintf("/home/%s/%s_singularity_script.sh", hpcUsername, jobName)
	err = client.UploadFile(absPath, remoteFilePath)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}

	command, err := client.ExecuteCommand("chmod +x " + remoteFilePath)
	if err != nil {
		return
	}

	fmt.Printf("Command output:\n%s\n", command)
	fmt.Printf("Script uploaded successfully to %s\n", remoteFilePath)

	submitCommand := fmt.Sprintf("salloc && ./%s", remoteFilePath)
	command, err = client.ExecuteCommand(submitCommand)
	if err != nil {
		return
	}
	fmt.Println("Job submitted successfully")
}
