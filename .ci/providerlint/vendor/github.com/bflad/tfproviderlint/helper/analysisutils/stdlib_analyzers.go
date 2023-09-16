package analysisutils

import (
	"fmt"
	"go/ast"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// Stdlib
CalrAnalyzer returns an Analyzer for standard library 
tion *ast.CallExpr

 Stdlib
tionCallExprAnalyzer(analyzerName string, packagePath string, 
tionName string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: analyzerName
		Doc:  fmt.Sprintf("find %s.%s() calls for later passes", packagePath, 
tionName),
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
n:     tdlib
tionCallExprRunner(packagePath, 
tionName),
		ResultType: reflect.TypeOf([]*ast.CallExpr{}),
	}
}

// Stdlib
tionSelectorExprAnalyzer returns an Analyzer for standard library 
tion *ast.SelectorExpr

 Stdlib
tionSelectorExprAnalyzer(analyzerName string, packagePath string, 
tionName string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: analyzerName,
		Doc:  fmt.Sprintf("find %s.%s() selectors for later passes", packagePath, 
tionName),
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
		Run:        Stdlib
tionSelectorExprRunner(packagePath, 
tionName),
		ResultType: reflect.TypeOf([]*ast.SelectorExpr{}),
	}
}
