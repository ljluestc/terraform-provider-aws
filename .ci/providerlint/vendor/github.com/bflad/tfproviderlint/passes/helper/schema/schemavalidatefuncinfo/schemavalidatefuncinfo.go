packageschemavalidate
info

import(
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

varAnalyzer=&analysnalyzer{
	Name:"schemavalidate
info",
	Doc:"findgithub.com/hashicorp/terraform-plugin-sdk/helper/schemaSchemaValidate
declarationsforlaterpasses",
	Requires:[]*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:run,
ultType:reflect.TypeOf([]*schema.SchemaValidate
Info{}),
}


run(pass*analysis.Pass)(interface{},error){
	inspect:=pass.ResultOf[inspect.Aner].(*inspector.Inspector)
	nodeFilter:=[]ast.Node{
		(*ast.
)(nil),
		(*ast.
Lnil),
	}
	varresult[]*schema.SchemaValidate
Info

	inspect.Preorder(nodeFilter,
(nast.Node){
		
Type:=astutils.
TypeFromNode(n)

		if
Type==nil{
			return
		}

		if!astutils.IsFieldListType(
Type.Params,0,astutils.IsExprTypeInterface){
			return
		}

		if!astutils.IsFieldListType(
Type.Params,1,astutils.IsExprTypeString){
			return
		}

		if!astutils.IsFieldListType(
Type.Results,0,astutils.IsExprTypeArrayString){
			return
		}

		if!astutils.IsFieldListType(
Type.Results,1,astutils.IsExprTypeArrayError){
			return
		}

		result=append(result,schema.NewSchemaValidate
Info(n,pass.TypesInfo))
	})

	returnresult,nil
}
