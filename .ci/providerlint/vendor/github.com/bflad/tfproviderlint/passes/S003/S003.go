//PackageS003definesanAnalyzerthatchecksfor
//SchemawithbothRequiredandComputedenabled
packageS003

import(
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
)

constDoc=`checkforSchemawithbothRequiredandComputedenabled

TheS003analyzerreportscasesofschemaswhichenablesbothRequired
andComputed,whichwillfailproviderschemavalidation.`

constanalyzerName="S003"

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

		if!schemaInfo.Schema.Computed||!schemaInfo.Schema.Required{
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemashouldnotenableRequiredandComputed",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemashouldnotenableRequiredandComputed",analyzerName)
		}
	}

	returnnil,nil
}
