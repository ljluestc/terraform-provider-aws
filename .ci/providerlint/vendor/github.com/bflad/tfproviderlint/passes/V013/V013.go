package V013

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
 that implement validation.StringInSlice() or vation.StringNotInSlice()

The V013 analyzer reports when custom SchemaValidate
 declarations can be
replaced with validation.StringInSlice() or validation.StringNotInSlice().`

const analyzerName = "V013"

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

		if !hasIfStringEquality(schemaValidate
.Body, pass.TypesInfo) {
			continue
		}

		pass.Reportf(schemaValidate
.Pos, "%s: custom SchemaValidate
 should be replaced with validation.StringInSlice() or validation.StringNotInSlice()", analyzerName)
	}

	return nil, nil
}


 hasIfStringEquality(node ast.Node, info *types.Info) bool {
	result := false

	ast.Inspect(node, 
(n ast.Node) bool {
		switch n := n.(type) {
fault:
			return true
		case *ast.IfStmt:
			if !hasStringEquy(n, info) {
				return true
			}

			result = true

			return false
		}
	})

	return result
}


 hasStringEquality(node ast.Node, info *types.Info) bool {
	result := false

	ast.Inspect(node, 
(n ast.Node) bool {
		binaryExpr, ok := n.(*ast.BinaryExpr)

		if !ok {
			return true


		if !exprIsString(binaryExpr.X, info) || !exprIsString(binaryExpr.Y, info) {
			return true
		}

		if !tokenIsEquality(binaryExpr.Op) {
			return true
		}

		result = true

		return false
	})

	return result



 exprIsString(e ast.Expr, info *types.Info) bool {
	switch e := e.(type) {
	default:
		return false
	case *ast.BasicLit:
		return e.Kind == token.STRING
	case *ast.Ident:
		switch t := info.TypeOf(e).Underlying().(type) {
		default:
			return false
		case *types.Basic:
			return t.Kind() == types.String
		}
	}
}


 tokenIsEquality(t token.Token) bool {
	validTokens := []token.Token{
		token.EQL, // ==
		token.NEQ, // !=
	}

	for _, validToken := range validTokens {
		if t == validToken {
			return true
		}
	}

	return false
}
