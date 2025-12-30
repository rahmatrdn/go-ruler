package parser

import "github.com/rahmatrdn/ai-guidelines-generator/internal/model"

type Parser interface {
	Parse(content string) (*model.Guidelines, error)
}
