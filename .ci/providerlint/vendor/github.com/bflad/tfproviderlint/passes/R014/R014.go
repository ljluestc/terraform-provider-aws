package R014

import (
	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/crud
info"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for Create
, CreateContext
, Delete
, DeleteContext
, Read
, ReadContext
, Update
, and UpdateContext
 parameter naming

The R014 analyzer reports when Create
, CreateContext
, Delete
,
Deleteext
, Read
, ReadContext
, Update
, and UpdateContext

declarations do not use d as the name for the *schema.ResourceData parameter
or mes the name for the intee{} parameter. This parameter ng is the
standard convention for resources.`

const analyzerName = "R014"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		crud
info.Analyzer,
	},
	Run: run,
}


 run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	crud
s := pass.ResultOf[crud
info.Analyzer].([]*schema.CRUD
Info)

	for _, crud
 := range crud
s {
		if ignorer.ShouldIgnore(analyzerName, crud
.Node) {
			continue
		}

		params := crud
.Type.Params
		paramCount := len(params.List)

		switch paramCount {
		case 2:
			if name := astutils.FieldListName(params, 0, 0); name != nil && *name != "_" && *name != "d" {
				pass.Reportf(params.List[0].Pos(), "%s: *schema.ResourceData parameter of Create
, Read
, Update
, or Delete
 should be named d", analyzerName)
			}

			if name := astutils.FieldListName(params, 1, 0); name != nil && *name != "_" && *name != "meta" {
				pass.Reportf(params.List[1].Pos(), "%s: interface{} parameter of Create
, Read
, Update
, or Delete
 should be named meta", analyzerName)
			}
		case 3:
			if name := astutils.FieldListName(params, 1, 0); name != nil && *name != "_" && *name != "d" {
				pass.Reportf(params.List[1].Pos(), "%s: *schema.ResourceData parameter of CreateContext
, ReadContext
, UpdateContext
, or DeleteContext
 should be named d", analyzerName)
			}

			if name := astutils.FieldListName(params, 2, 0); name != nil && *name != "_" && *name != "meta" {
				pass.Reportf(params.List[2].Pos(), "%s: interface{} parameter of CreateContext
, ReadContext
, UpdateContext
, or DeleteContext
 should be named meta", analyzerName)
			}
		}
	}

	return nil, nil
}
