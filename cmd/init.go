package cmd

import (
	"github.com/spf13/cobra"
	"simplearity/utils"
)

func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize configuration files",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CreateFileFromGitHub("simplearity.yaml", "https://gist.githubusercontent.com/kemalcanbora/eaa28ea6858961d0f0a56ee955e25368/raw/99658dad11ded0d6c459b648e6c0097ee2c4701b/simplearity.yaml.template")
			utils.CreateFileFromGitHub("simplearity.env", "https://gist.githubusercontent.com/kemalcanbora/eaa28ea6858961d0f0a56ee955e25368/raw/99658dad11ded0d6c459b648e6c0097ee2c4701b/simplearity.env.template")
		},
	}
}
