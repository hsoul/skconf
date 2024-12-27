package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hsoul/skconf/internal/ast"
	"github.com/hsoul/skconf/internal/generator"
	"github.com/hsoul/skconf/internal/generator/languages/lua"
	"github.com/hsoul/skconf/internal/lexer"
	"github.com/hsoul/skconf/internal/syntax"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: dsl <input_file> <output_dir>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputDir := os.Args[2]

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	l := lexer.New(string(input))
	p := syntax.New(l, inputFile)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Printf("Parser errors:\n")
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(1)
	}

	astTree := ast.PrintTree(program)
	astFile := filepath.Join(outputDir, "ast_tree.dot")
	if err := os.WriteFile(astFile, []byte(astTree), 0644); err != nil {
		log.Printf("Warning: Failed to write AST tree to file: %v", err)
	}

	gen, err := generator.New(lua.Language)
	if err != nil {
		log.Fatalf("Error creating generator: %v", err)
	}
	output := gen.Generate(program)

	fileName := filepath.Base(inputFile)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	outputFile := filepath.Join(outputDir, baseName+".lua")
	if err := os.WriteFile(outputFile, []byte(output), 0644); err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	fmt.Printf("Successfully generated:\n")
	fmt.Printf("- AST: %s\n", astFile)
	fmt.Printf("- Lua: %s\n", outputFile)
}
