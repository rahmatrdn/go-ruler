package parser

import (
	"bufio"
	"strings"

	"github.com/rahmatrdn/go-ruler/internal/model"
)

type MarkdownParser struct{}

func NewMarkdownParser() *MarkdownParser {
	return &MarkdownParser{}
}

func (p *MarkdownParser) Parse(content string) (*model.Guidelines, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	guidelines := &model.Guidelines{
		Raw: content,
	}

	var currentSection string

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "#") {
			lowerLine := strings.ToLower(trimmedLine)
			if strings.Contains(lowerLine, "command") {
				currentSection = "commands"
				continue
			} else {
				// Reset section or treat as just a header
				currentSection = ""
			}
		}

		if currentSection == "commands" {
			if strings.HasPrefix(trimmedLine, "-") || strings.HasPrefix(trimmedLine, "*") {
				cmdText := strings.TrimSpace(strings.TrimLeft(trimmedLine, "-* "))
				parts := strings.SplitN(cmdText, ":", 2)
				cmd := model.Command{
					Name: strings.TrimSpace(parts[0]),
				}
				if len(parts) > 1 {
					cmd.Description = strings.TrimSpace(parts[1])
				}
				guidelines.Commands = append(guidelines.Commands, cmd)
			}
		} else {
			// Treat as Rule if it looks like a list item
			if strings.HasPrefix(trimmedLine, "-") || strings.HasPrefix(trimmedLine, "*") {
				rule := strings.TrimSpace(strings.TrimLeft(trimmedLine, "-* "))
				if rule != "" {
					guidelines.Rules = append(guidelines.Rules, rule)
				}
			}
		}
	}

	return guidelines, nil
}
