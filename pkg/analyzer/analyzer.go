package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "Simple",
	Doc:      "Tests that no function name contains silly",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// The ResultOf field provides the results computed by the analyzers required by this one,
	// as expressed in its Analyzer.Requires field.

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspector.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(n ast.Node) {
		f, ok := n.(*ast.FuncDecl)
		if !ok {
			return
		}
		if strings.Contains(strings.ToLower(f.Name.Name), "silly") {
			pass.Reportf(n.Pos(), "function '%s', should not contain silly", f.Name.Name)
		}
	})

	return nil, nil
}
