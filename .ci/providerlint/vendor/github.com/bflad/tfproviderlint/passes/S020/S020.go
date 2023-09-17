//PackageS020definesanAnalyzerthatchecksfor
//SchemawithonlyComputedenabledandValidate
configured
packageS020

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfocomputedonly"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforSchemawithonlyComputedenabledandForceNewenabled

TheS020analyzerreportscasesofschemaswhichonlyenablesComputed
andenablesForceNew,whichisnotvalid.`

constanalyzerName="S020"

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

		if!schemaInfo.Schema.Computed||schemaInfo.Schema.Optional||schemaInfo.Schema.Required{
			continue
		}

		if!schemaInfo.Schema.ForceNew{
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemashouldnotonlyenableComputedandenableForceNew",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemashouldnotonlyenableComputedandenableForceNew",analyzerName)
		}
	}

	returnnil,nil
}
