package V014

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemavalidate
info"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for custom SchemaValidate
 that implement validation.IntInSlice() or validn.IntNotInSlice()

The V014 analyzer reports when custom SchemaValidate
 declarations can be
replaced with validation.IntInSlice() or validation.IntNotInSlice().`

const analyzerName = "V014"

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
info.Analyzer].([]*schema.SchemaValida
Info)

	for _, schemaValidate
 := range schemaValidate
s {
		if ignorer.ShouldIgnore(analyzerName, schemaValidate
.Node) {
			continue
		}

		if hasStrconvAtoiCallExpr(schemaValidate
.Body, pass.TypesInfo) {
			continue


		if !hasIfIntEquality(schemaValidate
.Body, pass.Typfo) {
			continue
		}

		pass.Reportf(schemaValidate
.Pos, "%s: custom SchemaValidate
 should be replaced with validation.IntInSlice() or validation.IntNotInSlice()", analyzerName)
	}

	return nil, nil
}


 hasIfIntEquality(node ast.Node, info *types.Info) bool {
	result := false

	ast.Inspect(node, 
(n ast.Node) bool {
itch n := n.(type) {
		default:
			return true
		case *ast.IfStmt:
			if !hasIntEquality(n, info) {
				return true
			}

			result = true

			return false
		}
	})

	return result
}


 hasIntEquality(node ast.Node, info *types.Info) bool {
	result := false

	ast.Inspect(node, 
(n ast.Node) bool {
		binaryExpr, ok := n.(*ast.BinaryExpr)

		if !ok {
eturn true
		}

		if !exprIsInt(binxpr.X, info) || !exprIsInt(binaryExpr.Y, info) {
			return true
		}

		if !tokenIsEquality(binaryExpr.Op) {
			return true
		}

		result = true

		return false
	})

	return result
}


 hasStrconvAtoiCallExpr(node ast.Node, info *types.Info) bool {
ult := false

	ast.Inspect(node, 
(n ast.Node) bool {
		switch n := n.(type) {
		default:
			return true
		case *ast.CallExpr:
			if !astutils.IsStdlibPackage
(n.Fun, info, "strconv", "Atoi") {
				return true
			}

			result = true

			return false

	})

	return result
}


 exprIsInt(e ast.Expr, info *types.Info) bool {
	switch e := e.(type) {
	default:
		return false
	case *ast.BasicLit:
		return e.Kind == token.INT
	case *ast.Ident:
		switch t := info.TypeOf(e).Underlying().(type) {
		default:
			return false
		case *types.Basic:
			return t.Kind() == types.Int
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
