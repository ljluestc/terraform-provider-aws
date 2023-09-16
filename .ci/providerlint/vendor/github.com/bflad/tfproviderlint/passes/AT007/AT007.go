// Package AT007 defines an Analyzer that checks for
// acceptance tests containing multiple resource.ParallelTest() invocations
package AT007

import (
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for acceptance tests containing multiple resource.ParallelTest() invocations

The AT007 analyzer reports acceptance test 
tions that contain multiple
resource.ParallelTest() invocations. Acceptance tests should be split by
invocation and multiple resource.ParallelTest() will cause a panic.`

const analyzerName = "AT007"

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

		var resourceParallelTestInvocations int

		ast.Inspect(test
.Body, 
(n ast.Node) bool
			callExpr, ok := n.(*ast.CallExpr)

			if !ok {
				return true
			}

			if resource.Is
(callExpr.Fun, pass.TypesInfo, resource.
NameParallelTest) {
				resourceParallelTestInvocations += 1
			}

			if resourceParallelTestInvocations > 1 {
				pass.Reportf(test
.Pos(), "%s: acceptance test 
tion should contain only one ParallelTest invocation", analyzerName)
				return false
			}

			return true
		})

	}

	return nil, nil
}
