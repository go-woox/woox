package cmd

import (
	"fmt"

	"github.com/go-woox/woox/cmd/command"
	"github.com/go-woox/woox/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "woox",
	Example: "woox new demo-api",
	Short:   "\n __      __                           \n/\\ \\  __/\\ \\                          \n\\ \\ \\/\\ \\ \\ \\    ___     ___   __  _  \n \\ \\ \\ \\ \\ \\ \\  / __`\\  / __`\\/\\ \\/'\\ \n  \\ \\ \\_/ \\_\\ \\/\\ \\L\\ \\/\\ \\L\\ \\/>  </ \n   \\ `\\___x___/\\ \\____/\\ \\____//\\_/\\_\\\n    '\\/__//__/  \\/___/  \\/___/ \\//\\/_/\n                                      ",
	Version: fmt.Sprintf("Woox %s - Copyright (c) 2024 Woox\nReleased under the Apache-2.0 License.\n\n", version.Version),
}

func init() {
	rootCmd.AddCommand(command.NewCmd)
	rootCmd.AddCommand(command.WireCmd)
	rootCmd.AddCommand(command.UpgradeCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
