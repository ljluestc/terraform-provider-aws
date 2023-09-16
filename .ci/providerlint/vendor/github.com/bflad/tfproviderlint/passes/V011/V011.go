package V011

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemavalidate
info"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for custom SchemaValidate
 that implement validation.StringLenBetween()

The V011 analyzer reports when custom SchemaValidate
 declarations can be
replaced with validation.StringLenBetween().`

const analyzerName = "V011"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		schemavalidate
info.Analyzer,

	Run: run,
}


 run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaValidate
s := pass.ResultOf[schemavalidate
info.Analyzer].([]*schema.SchemaVali
Info)

	for _, schemaValidate
 := range schemaValidate
s {
		if ignorer.ShouldIgnore(analyzerName, schemaValidate
.Node) {
			continue
		}

		if !hasIfStringLenCheck(schemaValidate
.Body, pass.TypesInfo) {
			continue
		}

		pass.Reportf(schemaValidate
.Pos, "%s: custom SchemaValidate
 should be replaced with validation.StringLenBetween()", analyzerName)
	}

	return nil, nil
}


 hasIfStringLenCheck(node ast.Node, info *types.Info) bool {
	result := false

	ast.Inspect(node, 
(n ast.Node) bool {
		switch n := n.(type) {
fault:
			return true
		case *ast.IfStmt:
			if !hasStringLenk(n, info) {
				return true
			}

			result = true

			return false
		}
	})

	return result
}


 hasStringLenCheck(node ast.Node, info *types.Info) bool {
	result := false

	ast.Inspect(node, 
(n ast.Node) bool {
		binaryExpr, ok := n.(*ast.BinaryExpr)

		if !ok {
			return true


		if !exprIsStringLenCallExpr(binaryExpr.X, info) && !exprIsStringLenCallExpr(binaryExpr.Y, info) {
			return true
		}

		if !tokenIsLenCheck(binaryExpr.Op) {
			return true
		}

		result = true

		return false
	})

	return result
}


 exprIsStringLenCallExpr(e ast.Expr, info *types.Info) bool {
	switch e := e.(type) {
	default:
		return false
	case *ast.CallExpr:
		switch fun := e.Fun.(type) {
		default:
			return false
se *ast.Ident:
			if fun.Name != "len" {
				return false
			}
		}

		if len(e.Args) != 1 {
			return false
		}

		switch arg := info.TypeOf(e.Args[0]).Underlying().(type) {
		default:
			return false
		case *types.Basic:
			return arg.Kind() == types.String
		}
	}
}


 tokenIsLenCheck(t token.Token) bool {
	validTokens := []token.Token{
		token.GEQ, // >=
		token.GTR, // >
		token.LEQ, // <=
		token.LSS, // <
	}

	for _, validToken := range validTokens {
		if t == validToken {
			return true
		}
	}

	return false
}
