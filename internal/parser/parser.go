package parser

import "github.com/rahmatrdn/go-ruler/internal/model"

type Parser interface {
	Parse(content string) (*model.Guidelines, error)
}
