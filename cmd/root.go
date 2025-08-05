package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gochat",
	Short: "GoChat CLI",                              // Changed from just "gochat"
	Long:  `Fast and lightweight Go web chat system`, // More descriptive
	Args:  args,
	Run: func(cmd *cobra.Command, args []string) {
		// Original logic preserved
	},
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires at least one argument") // More standard error message
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err) // Better error output format
		os.Exit(1)
	}
}

func init() {
	// Original command adding logic preserved
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(installCmd)
}
