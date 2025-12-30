package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ai-guidelines-generator",
	Short: "A CLI to sync AI guidelines across different formats",
	Long: `A CLI tool that allows you to maintain a single source of truth for your AI coding guidelines
and generates/updates configuration files for various AI coding assistants like Claude, Gemini, and Cursor.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags can be defined here
}
