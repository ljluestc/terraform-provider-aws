packageR014

import(
	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/crud
info"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforCreate
,CreateContext
,Delete
,DeleteContext
,Read
,ReadContext
,Update
,andUpdateContext
parameternaming

TheR014analyzerreportswhenCreate
,CreateContext
,Delete
,
Deleteext
,Read
,ReadContext
,Update
,andUpdateContext

declarationsdonotusedasthenameforthe*schema.ResourceDataparameter
ormesthenamefortheintee{}parameter.Thisparameterngisthe
standardconventionforresources.`

constanalyzerName="R014"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		crud
info.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	crud
s:=pass.ResultOf[crud
info.Analyzer].([]*schema.CRUD
Info)

	for_,crud
:=rangecrud
s{
		ifignorer.ShouldIgnore(analyzerName,crud
.Node){
			continue
		}

		params:=crud
.Type.Params
		paramCount:=len(params.List)

		switchparamCount{
		case2:
			ifname:=astutils.FieldListName(params,0,0);name!=nil&&*name!="_"&&*name!="d"{
				pass.Reportf(params.List[0].Pos(),"%s:*schema.ResourceDataparameterofCreate
,Read
,Update
,orDelete
shouldbenamedd",analyzerName)
			}

			ifname:=astutils.FieldListName(params,1,0);name!=nil&&*name!="_"&&*name!="meta"{
				pass.Reportf(params.List[1].Pos(),"%s:interface{}parameterofCreate
,Read
,Update
,orDelete
shouldbenamedmeta",analyzerName)
			}
		case3:
			ifname:=astutils.FieldListName(params,1,0);name!=nil&&*name!="_"&&*name!="d"{
				pass.Reportf(params.List[1].Pos(),"%s:*schema.ResourceDataparameterofCreateContext
,ReadContext
,UpdateContext
,orDeleteContext
shouldbenamedd",analyzerName)
			}

			ifname:=astutils.FieldListName(params,2,0);name!=nil&&*name!="_"&&*name!="meta"{
				pass.Reportf(params.List[2].Pos(),"%s:interface{}parameterofCreateContext
,ReadContext
,UpdateContext
,orDeleteContext
shouldbenamedmeta",analyzerName)
			}
		}
	}

	returnnil,nil
}
