// Package AT002 defines an Analyzer that checks for
// acceptance test names including the word import
package AT002

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
)

const Doc = `check for acceptance test 
tion names including the word import

The AT002 analyzer reports where the word import or Import is used
in an acceptance test 
tion name, which generally means there is an extraneous
acceptance test. ImportState testing should be included as a TestStep with each
applicable acceptance test, rather than a separate test that only verifies import
of a single test configuration.`

const analyzerName = "AT002"

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
 pass.ResultOstacc
decl.Analyzer].([]*ast.
Decl)
	for _, testAcc
 := range testAcc
s {
		if ignorer.ShouldIgnore(analyzerName, testAcc
) {
			continue
		}

		
Name := testAcc
.Name.Name

		if strings.Contains(
Name, "_import") || strings.Contains(
Name, "_Import") {
			pass.Reportf(testAcc
.Name.NamePos, "%s: acceptance test 
tion name should not include import", analyzerName)
		}
	}

	return nil, nil
}
