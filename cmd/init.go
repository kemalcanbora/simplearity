package cmd

import (
	"github.com/spf13/cobra"
	"simplearity/utils/helper"
)

func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize configuration files",
		Run: func(cmd *cobra.Command, args []string) {
			helper.FetchFileFromGitHub("simplearity.yaml", "https://gist.githubusercontent.com/kemalcanbora/eaa28ea6858961d0f0a56ee955e25368/raw/beddec4f740a96ce6f4d594c7ae0d3d6436a5ab0/simplearity.yaml.template")
			helper.FetchFileFromGitHub("simplearity.env", "https://gist.githubusercontent.com/kemalcanbora/eaa28ea6858961d0f0a56ee955e25368/raw/7ecb21095c3ba7f6b15c65880d94bc340bed5640/simplearity.env.template")
		},
	}
}
