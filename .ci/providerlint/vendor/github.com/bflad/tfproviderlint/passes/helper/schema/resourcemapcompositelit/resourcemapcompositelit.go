packageresourcemapcompositelit

import(
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

varAnalyzer=&analysis.Analyzer{
	Name:"resourcemapcompositelit",
	Doc:"findmap[string]*schema.Resourceliteralsforlaterpasses",
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

		if!schema.IsMapStringResource(x,pass.TypesInfo){
			return
		}

		result=append(result,x)
	})

	returnresult,nil
}
