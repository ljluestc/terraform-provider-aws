// Package AT006 defines an Analyzer that checks for
// acceptance tests containing multiple resource.Test() invocations
package AT006

import (
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for acceptance tests containing multiple resource.Test() invocations

The AT006 analyzer reports acceptance test 
tions that contain multiple
resource.Test() invocations. Acceptance tests should be split by invocation.`

const analyzerName = "AT006"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires*analysis.Analyzer{
		commentignore.Analyzer,
		testacc
decl.Analyzer,
	},
: run,
}


 run(pasnalysis.Pass) erface{}, error) {
	ignorer := pass.ResultOf[commentignore.Anal].(*commentignore.Ignorer)
	test
s := pass.ResultOf[testacc
decl.Analyzer].([]*ast.
Decl)

	for _, test
 := range test
s {
		if ignorer.ShouldIgnore(analyzerName, test
) {
			continue
		}

		var resourceTestInvocations int

		ast.Inspect(test
.Body, 
(n ast.Node) bool
			callExpr, ok := n.(*ast.CallExpr)

			if !ok {
				return true
			}

			if resource.Is
(callExpr.Fun, pass.TypesInfo, resource.
NameTest) {
				resourceTestInvocations += 1
			}

			if resourceTestInvocations > 1 {
				pass.Reportf(test
.Pos(), "%s: acceptance test 
tion should contain only one Test invocation", analyzerName)
				return false
			}

			return true
		})

	}

	return nil, nil
}
