package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rahmatrdn/ai-guidelines-generator/internal/generator"
	"github.com/rahmatrdn/ai-guidelines-generator/internal/parser"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync [source_file]",
	Short: "Generate AI rules/guidelines from source to other formats",
	Long: `Reads the source file (e.g., CLAUDE.md) and updates/generates other AI guideline files
(GEMINI.md, .cursor/rules/general.mdc) to keep them in sync.`,
	Args: cobra.ExactArgs(1),
	RunE: runSync,
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initialize empty AI guideline files",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	targets := []string{
		"CLAUDE.md",
		"GEMINI.md",
		"AGENTS.md",
		".cursor/rules/general.mdc",
	}

	defaultContent := "How smart the AI is depends on us!\nWrite standard guidelines here according to your project’s needs ....\n"

	for _, target := range targets {
		filename := filepath.Join(dir, target)

		// Check if file exists
		if _, err := os.Stat(filename); err == nil {
			fmt.Printf("File already exists, skipping: %s\n", filename)
			continue
		}

		// Ensure directory exists
		targetDir := filepath.Dir(filename)
		if targetDir != "." {
			if err := os.MkdirAll(targetDir, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", targetDir, err)
				continue
			}
		}

		if err := os.WriteFile(filename, []byte(defaultContent), 0644); err != nil {
			fmt.Printf("Error writing to %s: %v\n", filename, err)
		} else {
			fmt.Printf("Initialized: %s\n", filename)
		}
	}

	return nil
}

func runSync(cmd *cobra.Command, args []string) error {
	sourceFile := args[0]
	fmt.Printf("Syncing from source: %s\n", sourceFile)

	// Verify source exists
	contentBytes, err := os.ReadFile(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}
	content := string(contentBytes)

	// 1. Parse
	// TODO: Detect parser based on extension or flags. For now, default to Markdown.
	p := parser.NewMarkdownParser()
	guidelines, err := p.Parse(content)
	if err != nil {
		return fmt.Errorf("failed to parse guidelines: %w", err)
	}

	fmt.Printf("Parsed %d rules and %d commands\n", len(guidelines.Rules), len(guidelines.Commands))

	// 2. Define Targets
	// TODO: Make this configurable or auto-detect
	targets := map[string]generator.Generator{
		"CLAUDE.md":                 generator.NewMarkdownGenerator(),
		"GEMINI.md":                 generator.NewMarkdownGenerator(),
		"AGENTS.md":                 generator.NewMarkdownGenerator(),
		".cursor/rules/general.mdc": generator.NewMarkdownGenerator(),
	}

	// 3. Generate and Write
	for filename, gen := range targets {
		// Skip writing back to the source file to avoid potential dataloss/loops if logic isn't perfect
		// or if we want to support editing the source.
		// For now, let's allow overwriting everything including source if it matches target name.
		// But maybe safer to skip exact match strings.
		if filename == sourceFile {
			continue
		}

		genContent, err := gen.Generate(guidelines)
		if err != nil {
			fmt.Printf("Error generating for %s: %v\n", filename, err)
			continue
		}

		// Ensure directory exists
		dir := filepath.Dir(filename)
		if dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dir, err)
				continue
			}
		}

		if err := os.WriteFile(filename, []byte(genContent), 0644); err != nil {
			fmt.Printf("Error writing to %s: %v\n", filename, err)
		} else {
			fmt.Printf("Updated %s\n", filename)
		}
	}

	return nil
}
