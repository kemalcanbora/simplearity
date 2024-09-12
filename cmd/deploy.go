package cmd

import (
	"github.com/spf13/cobra"
	"simplearity/utils/deploy"
)

func DeployCmd(hpcUsername, dockerHubUsername, imageName, jobName, mem, partition string, cpus int) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploy to HPC",
		Run: func(cmd *cobra.Command, args []string) {
			deploy.Deploy(hpcUsername, dockerHubUsername, imageName, jobName, mem, partition, cpus)
		},
	}
}
