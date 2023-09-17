packagetestacc
decl

import(
	"go/ast"
	"reflect"
	"strings"

	"github.com/bflad/tfproviderlint/passes/test
decl"
	"golang.org/x/tools/go/analysis"
)

varAnalyzer=&anis.Analyzer{
	Name:"testacc
de
	Doc:"find*ast.
DeclwithTestAccprefixednamesforlaterpasses",
	Requires:[]*analysis.Analyzer{
		test
decl.Analyzer,

	Run:run,
	ResultType:reflect.TypeOf([]*ast.
Decl{}),
}


run(pass*analysis.Pass)(rface{},error){
	test
s:=pass.ResultOf[test
decl.Analyzer].([]*ast.
Decl)

	varresult[]*ast.
Decl

	for_,test
:=rangetest
s{
		ifstrings.HasPrefix(test
.Name.Name,"TestAcc"){
			result=append(result,test
)
		}
	}

	returnresult,nil
}
