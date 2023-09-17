packageS025

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfocomputedonly"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforSchemawithonlyComputedenabledandAtLeastOneOfconfigured

TheS025analyzerreportscasesofschemaswhichonlyenablesComputed
andconfiguresAtLeastOneOf,whichisnotvalid.`

constanalyzerName="S025"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		schemainfocomputedonly.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos:=pass.ResultOf[schemainfocomputedonly.Analyzer].([]*schema.SchemaInfo)
	for_,schemaInfo:=rangeschemaInfos{
		ifignorer.ShouldIgnore(analyzerName,schemaInfo.AstCompositeLit){
			continue
		}

		ifschemaInfo.Schema.AtLeastOneOf==nil{
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemashouldnotonlyenableComputedandconfigureAtLeastOneOf",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemashouldnotonlyenableComputedandconfigureAtLeastOneOf",analyzerName)
		}
	}

	returnnil,nil
}
