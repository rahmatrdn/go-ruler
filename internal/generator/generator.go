package generator

import "github.com/rahmatrdn/ai-guidelines-generator/internal/model"

type Generator interface {
	Generate(g *model.Guidelines) (string, error)
}
