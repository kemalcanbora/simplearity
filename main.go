package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"simplearity/cmd"
	"simplearity/utils"
)

func main() {
	var rootCmd = &cobra.Command{Use: "simplearity"}

	// Add basic commands
	rootCmd.AddCommand(cmd.InitCmd())
	rootCmd.AddCommand(cmd.CreateCmd())
	rootCmd.AddCommand(cmd.GpuCmd())
	rootCmd.AddCommand(cmd.JobsCmd())

	isInitialized := false
	if _, err := os.Stat("simplearity.env"); err == nil {
		isInitialized = true
		// Load environment variables and configuration
		err := utils.LoadEnv("simplearity.env")
		if err != nil {
			fmt.Printf("Error loading simplearity.env file: %v\n", err)
			fmt.Println("Please run 'simplearity init' to create a proper configuration.")
			os.Exit(1)
		}

		config, err := utils.LoadConfig()
		if err != nil {
			fmt.Printf("Error parsing configuration: %v\n", err)
			fmt.Println("Please run 'simplearity init' to reset the configuration.")
			os.Exit(1)
		}
		// Add Docker-related commands
		rootCmd.AddCommand(cmd.BuildCmd(config.AppImageName))
		rootCmd.AddCommand(cmd.RunCmd(config.AppImageName, config.Singularity))
	}

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if cmd.Use == "init" {
			return nil
		}
		if !isInitialized {
			if len(args) == 0 {
				fmt.Println("SimpleArity is not initialized. Please run 'simplearity init' to set up.")
			} else {
				fmt.Printf("Cannot run '%s' command. SimpleArity is not initialized.\n", args[0])
				fmt.Println("Please run 'simplearity init' first.")
			}
			os.Exit(1)
		}
		return nil
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
