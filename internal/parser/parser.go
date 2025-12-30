package parser

import "github.com/rahmatrdn/ai-ruler/internal/model"

type Parser interface {
	Parse(content string) (*model.Guidelines, error)
}
