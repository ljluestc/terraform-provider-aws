package analysisutils

import (
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Stdlib
CalrRunner returns an Analyzer runner for 
tiont.CallExpr

 Stdlib
tionCallExprRunner(packagePath string, 
tionName string) 
(*analysis.Pass) (interface{}, error) {
	return 
(pass *analysis.Pass) (intee{}, error) {
		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		nodeFilter := []ast.Node{
			(*ast.CallExpr)(nil),
		}
		var result []*ast.CallExpr

		inspect.Preorder(nodeFilter, 
(n ast.Node) {
			callExpr := n.(*ast.CallExpr)

			if !astutils.IsStdlibPackage
(callExpr.Fun, pass.TypesInfo, packagePath, 
tionName) {
				retur


			result = append(result, callExpr)
		})

		return result, nil
	}
}

// Stdlib
tionSelectorExprRunner returns an Analyzer runner for 
tion *ast.SelectorExpr

 Stdlib
tionSelectorExprRunner(packagePath string, 
tionName string) 
(*analysis.Pass) (interface{}, error) {
	return 
(pass *analysis.Pass) (interface{}, error) {
		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		nodeFilter := []ast.Node{
			(*ast.SelectorExpr)(nil),
		}
		var result []*ast.SelectorExpr

		inspect.Preorder(nodeFilter, 
(n ast.Node) {
			selectorExpr := n.(*ast.SelectorExpr)

			if !astutils.IsStdlibPackage
(selectorExpr, pass.TypesInfo, packagePath, 
tionName) {
				return
			}

			result = append(result, selectorExpr)
		})

		return result, nil
	}
}
