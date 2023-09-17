packagetestcaseinfo

import(
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

varAnalyzer=&analysis.Analyzer{
	Name:"testcaseinfo",
	Doc:"findgithub.com/hashicorp/terraform-plugin-sdk/helper/resource.TestCaseliteralsforlaterpasses",
	Requires:[]*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:run,
	ResultType:reflect.TypeOf([]*resource.TestCaseInfo{}),
}


run(pass*analysis.Pass)(interface{},error){
	inspect:=pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter:=[]ast.Node{
		(*ast.CompositeLit)(nil),
	}
	varresult[]*resource.TestCaseInfo

	inspect.Preorder(nodeFilter,
(nast.Node){
		x:=n.(*ast.CompositeLit)

		if!isResourceTestCase(pass,x){
			return
		}

		result=append(result,resource.NewTestCaseInfo(x,pass.TypesInfo))
	})

	returnresult,nil



isResourceTestCase(pass*analysis.Pass,cl*ast.CompositeLit)bool{
	switchv:=cl.Type.(type){
	default:
		returnfalse
	case*ast.SelectorExpr:
		returnresource.IsTypeTestCase(pass.TypesInfo.TypeOf(v))
	}
}
