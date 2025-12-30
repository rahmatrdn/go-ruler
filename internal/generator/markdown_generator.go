package generator

import (
	"fmt"
	"strings"

	"github.com/rahmatrdn/go-ruler/internal/model"
)

type MarkdownGenerator struct {
	HeaderStyle string // e.g. "#" or "##"
}

func NewMarkdownGenerator() *MarkdownGenerator {
	return &MarkdownGenerator{
		HeaderStyle: "#",
	}
}

func (g *MarkdownGenerator) Generate(guidelines *model.Guidelines) (string, error) {
	// If Raw content is available, use it to ensure exact sync
	if guidelines.Raw != "" {
		return guidelines.Raw, nil
	}

	var sb strings.Builder

	// Rules Section
	if len(guidelines.Rules) > 0 {
		sb.WriteString(fmt.Sprintf("%s Rules\n\n", g.HeaderStyle))
		for _, rule := range guidelines.Rules {
			sb.WriteString(fmt.Sprintf("- %s\n", rule))
		}
		sb.WriteString("\n")
	}

	// Commands Section
	if len(guidelines.Commands) > 0 {
		sb.WriteString(fmt.Sprintf("%s Commands\n\n", g.HeaderStyle))
		for _, cmd := range guidelines.Commands {
			desc := ""
			if cmd.Description != "" {
				desc = fmt.Sprintf(": %s", cmd.Description)
			}
			sb.WriteString(fmt.Sprintf("- %s%s\n", cmd.Name, desc))
		}
		sb.WriteString("\n")
	}

	return sb.String(), nil
}
