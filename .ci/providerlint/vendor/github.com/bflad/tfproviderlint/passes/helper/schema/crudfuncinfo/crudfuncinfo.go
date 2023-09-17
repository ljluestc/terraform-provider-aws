packagecrud
info

import(
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

varAnalyzeranalysis.Analyzer{
	Name:"crud
info",
	Doc:"findgithub.com/hashicorp/terraform-plugin-sdk/helper/schemaCreate
,CreateContext
,Read
,ReadContext
,Update
,UpdateContext
lete
,andDeleteContext
declarationsforlaterpasses",
	Require]*analysis.Analyzer{
		inspecalyzer,
	},
	Run:run,
	ResultType:reflect.TypeOf([]*schema.CRUD
Info{}),
}


run(pass*analysis.Pass)(interface{},error){
	inspect:=pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter:=[]ast.Node{
		(*ast.
Decl)(nil),
		(*ast.
Lit)(nil),
	}
	varresult[]*schema.CRUD
Info

	inspect.Preorder(nodeFilter,
(nast.Node){
		ifschema.Is
TypeCRUD
(n,pass.TypesInfo){
			result=append(result,schema.NewCRUD
Info(n,pass.TypesInfo))
		}
	})

	returnresult,nil
}
