package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "Simple",
	Doc:  "Tests that no function name contains silly",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			f, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}
			if strings.Contains(strings.ToLower(f.Name.Name), "silly") {
				pass.Reportf(n.Pos(), "function '%s', should not contain silly", f.Name.Name)
			}

			return true
		})
	}

	return nil, nil
}
