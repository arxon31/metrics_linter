package os_exit_analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var OsExitChecker = &analysis.Analyzer{
	Name: "osExitChecker",
	Doc:  "check for os.Exit calls in main package",
	Run:  runOsExitChecker,
}

// runOsExitChecker checks for os.Exit only in main package and main function
func runOsExitChecker(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			funcDecl, ok := node.(*ast.FuncDecl)
			if !ok || funcDecl.Name.Name != "main" {
				return true
			}
			for _, stmt := range funcDecl.Body.List {
				callExpr, ok := stmt.(*ast.ExprStmt)
				if !ok {
					continue
				}
				call, ok := callExpr.X.(*ast.CallExpr)
				if !ok {
					continue
				}
				selectorExpr, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				if ident, ok := selectorExpr.X.(*ast.Ident); !ok || ident.Name != "os" || selectorExpr.Sel.Name != "Exit" {
					continue
				}
				pass.Report(analysis.Diagnostic{
					Pos:     call.Pos(),
					Message: "os.Exit calls are not allowed",
				})
			}
			return true
		})
	}

	return nil, nil
}
