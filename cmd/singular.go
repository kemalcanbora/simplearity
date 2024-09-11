package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func GpuCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gpu",
		Short: "Display GPU information",
		Run: func(cmd *cobra.Command, args []string) {
			output, err := exec.Command("bash", "-c", "sinfo -o \"%N %G %m\" | awk '/gpu/ {print $1, $2, $3}'").Output()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				os.Exit(1)
			}

			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.TrimSpace(line) == "" {
					continue
				}
				parts := strings.Fields(line)
				if len(parts) > 1 {
					color.Set(color.FgGreen)
					fmt.Printf("%s ", parts[0])
					color.Set(color.FgYellow)
					fmt.Printf("%s ", parts[1])
					color.Set(color.FgCyan)
					fmt.Printf("%s\n", parts[2])
					color.Unset()
				}
			}
		},
	}
}

func JobsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "jobs",
		Short: "Display Singularity job information",
		Run: func(cmd *cobra.Command, args []string) {
			output, err := exec.Command("squeue", "--me", "--format=%i %j %T %M %l %R").Output()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				os.Exit(1)
			}

			lines := strings.Split(string(output), "\n")
			for i, line := range lines {
				if strings.TrimSpace(line) == "" {
					continue
				}
				parts := strings.Fields(line)
				if len(parts) >= 6 {
					if i == 0 {
						// Print header
						color.Set(color.FgWhite, color.Bold)
					} else {
						// Print job information
						switch parts[2] {
						case "RUNNING":
							color.Set(color.FgGreen)
						case "PENDING":
							color.Set(color.FgYellow)
						default:
							color.Set(color.FgRed)
						}
					}
					fmt.Printf("%-12s %-20s %-10s %-10s %-12s %s\n", parts[0], parts[1], parts[2], parts[3], parts[4], parts[5])
					color.Unset()
				}
			}
		},
	}
}
