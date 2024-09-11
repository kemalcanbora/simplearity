package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func RunCmd(AppImageName string, useSingularity bool) *cobra.Command {
	var err error
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the application using Docker or Singularity",
		Run: func(cmd *cobra.Command, args []string) {
			if useSingularity {
				err = runWithSingularity(AppImageName)
			} else {
				err = runWithDocker(AppImageName)
			}

			if err != nil {
				fmt.Printf("Error running the application: %v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().BoolVarP(&useSingularity, "singularity", "s", false, "Use Singularity instead of Docker")

	return cmd
}

func runWithDocker(AppImageName string) error {
	fmt.Println("Running with Docker...")

	cmd := exec.Command("docker", "run", AppImageName+":latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func runWithSingularity(AppImageName string) error {
	fmt.Println("Running with Singularity...")

	cmd := exec.Command("singularity", "run", "docker://"+AppImageName+":latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
