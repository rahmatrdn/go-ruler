package generator

import "github.com/rahmatrdn/go-ruler/internal/model"

type Generator interface {
	Generate(g *model.Guidelines) (string, error)
}
