package generator

import (
	"fmt"

	"github.com/hsoul/skconf/internal/ast"
)

type CodeGenerator interface {
	Generate(node ast.Node) string
}

type GeneratorConstructor func() (CodeGenerator, error)

var generators = make(map[string]GeneratorConstructor)

func Register(lang string, constructor GeneratorConstructor) {
	generators[lang] = constructor
}

func New(lang string) (CodeGenerator, error) {
	if constructor, ok := generators[lang]; ok {
		return constructor()
	}
	return nil, fmt.Errorf("unsupported language: %s", lang)
}
