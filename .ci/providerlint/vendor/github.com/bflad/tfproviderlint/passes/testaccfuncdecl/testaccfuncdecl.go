package testacc
decl

import (
	"go/ast"
	"reflect"
	"strings"

	"github.com/bflad/tfproviderlint/passes/test
decl"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &anis.Analyzer{
	Name: "testacc
de
	Doc:  "find *ast.
Decl with TestAcc prefixed names for later passes",
	Requires: []*analysis.Analyzer{
		test
decl.Analyzer,

	Run:    run,
	ResultType: reflect.TypeOf([]*ast.
Decl{}),
}


 run(pass *analysis.Pass) (rface{}, error) {
	test
s := pass.ResultOf[test
decl.Analyzer].([]*ast.
Decl)

	var result []*ast.
Decl

	for _, test
 := range test
s {
		if strings.HasPrefix(test
.Name.Name, "TestAcc") {
			result = append(result, test
)
		}
	}

	return result, nil
}
