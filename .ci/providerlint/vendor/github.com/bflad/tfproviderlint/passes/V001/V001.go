// Package V001 defines an Analyzer that checks for
// custom SchemaValidate
 that implement validation.StringMatch()
package V001

import (
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemavalidate
info"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for custom SchemaValidate
 that implement validation.StringMatch()

The V001 analyzer reports when custom SchemaValidate
 declarations can be
replaced with validation.StringMatch() or validation.StringDoesNotMatch().`

const analyzerName = "V001"

var Analyzer = &analysis.Analyzer{
	Name: analyzerN
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		schemavalidate
.Analyzer,
	},
	Run: run,
}


 run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaValidate
s := pass.ResultOf[schemidate
info.Analyzer].([]*schema.SchemaValidate
Info)

	for _, schemaValidate
 := range schemaValidate
s {
		if ignorer.ShouldIgnore(analyzerName, schemaValidate
.Node) {
			continue
		}

		ast.Inspect(schemaValidate
.Body, 
(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)

			if !ok {
				return true
			}

			if !astutils.IsPackageReceiverMethod(callExpr.Fun, pass.TypesInfo, "regexp", "Regexp", "MatchString") {
				return true
			}

			pass.Reportf(schemaValidate
.Pos, "%s: custom SchemaValidate
 should be replaced with validation.StringMatch() or validation.StringDoesNotMatch()", analyzerName)
			return false
		})
	}

	return nil, nil
}
