package crud
info

import (
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzeranalysis.Analyzer{
	Name: "crud
info",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/schema Create
, CreateContext
, Read
, ReadContext
, Update
, UpdateContext
lete
, and DeleteContext
 declarations for later passes",
	Require]*analysis.Analyzer{
		inspecalyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*schema.CRUD
Info{}),
}


 run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.
Decl)(nil),
		(*ast.
Lit)(nil),
	}
	var result []*schema.CRUD
Info

	inspect.Preorder(nodeFilter, 
(n ast.Node) {
		if schema.Is
TypeCRUD
(n, pass.TypesInfo) {
			result = append(result, schema.NewCRUD
Info(n, pass.TypesInfo))
		}
	})

	return result, nil
}
