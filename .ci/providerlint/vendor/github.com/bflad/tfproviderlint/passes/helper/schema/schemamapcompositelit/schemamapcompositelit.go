packageschemamapcompositelit

import(
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

varAnalyzer=&analysis.Analyzer{
	Name:"schemamapcompositelit",
	Doc:"findmap[string]*schema.Schemaliteralsforlaterpasses",
	Requires:[]*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:run,
	ResultType:reflect.TypeOf([]*ast.CompositeLit{}),
}


run(pass*analysis.Pass)(interface{},error){
	inspect:=pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter:=[]ast.Node{
		(*ast.CompositeLit)(nil),
	}
	varresult[]*ast.CompositeLit

	inspect.Preorder(nodeFilter,
(nast.Node){
		x:=n.(*ast.CompositeLit)

		if!schema.IsMapStringSchema(x,pass.TypesInfo){
			return
		}

		result=append(result,x)
	})

	returnresult,nil
}
