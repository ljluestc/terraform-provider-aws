//PackageS005definesanAnalyzerthatchecksfor
//SchemawithComputedenabledandDefaultconfigured
packageS005

import(
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
)

constDoc=`checkforSchemawithComputedenabledandDefaultconfigured

TheS005analyzerreportscasesofschemaswhichenablesComputed
andconfiguresDefault,whichwillfailproviderschemavalidation.`

constanalyzerName="S005"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		schemainfo.Analyzer,
		commentignore.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos:=pass.ResultOf[schemainfo.Analyzer].([]*schema.SchemaInfo)
	for_,schemaInfo:=rangeschemaInfos{
		ifignorer.ShouldIgnore(analyzerName,schemaInfo.AstCompositeLit){
			continue
		}

		if!schemaInfo.Schema.Computed||schemaInfo.Schema.Default==nil{
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemashouldnotenableComputedandconfigureDefault",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemashouldnotenableComputedandconfigureDefault",analyzerName)
		}
	}

	returnnil,nil
}
