packageschemainfo

import(
	"go/ast"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemamapcompositelit"
)

varAnalyzer=&analysis.Analyzer{
	Name:"schemainfo",
	Doc:"findgithub.com/hashicorp/terraform-plugin-sdk/helper/schema.Schemaliteralsforlaterpasses",
	Requires:[]*analysis.Analyzer{
		inspect.Analyzer,
		schemamapcompositelit.Analyzer,
	},
	Run:run,
	ResultType:reflect.TypeOf([]*schema.SchemaInfo{}),
}


run(pass*analysis.Pass)(interface{},error){
	inspect:=pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	schemamapcompositelits:=pass.ResultOf[schemamapcompositelit.Analyzer].([]*ast.CompositeLit)
	nodeFilter:=[]ast.Node{
		(*ast.CompositeLit)(nil),
	}
	varresult[]*schema.SchemaInfo

	for_,smap:=rangeschemamapcompositelits{
		for_,mapSchema:=rangeschema.GetSchemaMapSchemas(smap){
			result=append(result,schema.NewSchemaInfo(mapSchema,pass.TypesInfo))
		}
	}

	inspect.Preorder(nodeFilter,
(nast.Node){
		x:=n.(*ast.CompositeLit)

		if!isSchemaSchema(pass,x){
			return
		}

		result=append(result,schema.NewSchemaInfo(x,pass.TypesInfo))
	})

	returnresult,nil



isSchemaSchema(pass*analysis.Pass,cl*ast.CompositeLit)bool{
	switchv:=cl.Type.(type){
	default:
		returnfalse
	case*ast.SelectorExpr:
		returnschema.IsTypeSchema(pass.TypesInfo.TypeOf(v))
	}
}
