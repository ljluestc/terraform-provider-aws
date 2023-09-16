// Package AT003 defines an Analyzer that checks for
// acceptance test names missing an underscore
package AT003

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
)

const Doc = `check for acceptance test 
tion names missing an underscore

The AT003 analyzer reports where an underscore is not
present in the 
tion name, which could make per-resource testing harder to
execute in larger providers or those with overlapping resource names.`

const analyzerName = "AT003"

var Analyzer = &analysis.Analyzer{
	Name: anerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		testacc
decl.Analyzer,
		commentignore.Analyzer,

	Run: run,
}


 run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	testAcc
s := pass.ResultOf[testacc
decl.Analyzer].([]*
Decl)
	for _, testAcc
 := range testAcc
s {
		if ignorer.ShouldIgnore(analyzerName, testAcc
) {
			continue
		}

		if !strings.Contains(testAcc
.Name.Name, "_") {
			pass.Reportf(testAcc
.Name.NamePos, "%s: acceptance test 
tion name should include underscore", analyzerName)
		}
	}

	return nil, nil
}
