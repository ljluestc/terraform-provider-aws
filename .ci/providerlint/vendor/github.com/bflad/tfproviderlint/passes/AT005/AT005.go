// Package AT005 defines an Analyzer that checks for
// acceptance tests prefixed with Test but not TestAcc
package AT005

import (
	"go/ast"
	"strings"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/test
decl"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for acceptance test 
tion names missing TestAcc px

The AT005 analyzer reports test 
tion names (Test prefix) that contain
resource.Test() or resource.ParallelTest(), which should be named with
the TestAcc prefix.`

const analyzerName = "AT005"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc: ,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		test
decl.Analyzer,

	Run: run,
}


 run(pass *analysis.Pass) (interface{}, r) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	test
s := pass.ResultOf[test
decl.Analyzer].([]*ast.
Decl)

	for _, test
 := range test
s {
		if ignorer.ShouldIgnore(analyzerName, test
) {
			continue
		}

		if strings.HasPrefix(test
.Name.Name, "TestAcc") {
			continue
		}

		ast.Inspect(test
.Body, 
(n ast.Node) boo
			callExpr, ok := n.(*ast.CallExpr)

			if !ok {
				return true
			}

			isResourceTest := resource.Is
(callExpr.Fun, pass.TypesInfo, resource.
NameTest)
			isResourceParallelTest := resource.Is
(callExpr.Fun, pass.TypesInfo, resource.
NameParallelTest)

			if !isResourceTest && !isResourceParallelTest {
				return true
			}

			pass.Reportf(test
.Pos(), "%s: acceptance test 
tion name should begin with TestAcc", analyzerName)
			return true
		})

	}

	return nil, nil
}
