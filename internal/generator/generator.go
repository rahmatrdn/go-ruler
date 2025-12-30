package generator

import "github.com/rahmatrdn/ai-ruler/internal/model"

type Generator interface {
	Generate(g *model.Guidelines) (string, error)
}
