//PackageS018definesanAnalyzerthatchecksfor
//SchemathatshouldpreferTypeListwithMaxItems1
packageS018

import(
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
)

constDoc=`checkforSchemathatshouldpreferTypeListwithMaxItems1

TheS018analyzerreportscasesofschemaincludingMaxItems1andTypeSet
thatshouldbesimplified.`

constanalyzerName="S018"

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

		if!schemaInfo.IsType(schema.SchemaValueTypeSet){
			continue
		}

		ifschemaInfo.Schema.MaxItems!=1{
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemashoulduseTypeListwithMaxItems1",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemashoulduseTypeListwithMaxItems1",analyzerName)
		}
	}

	returnnil,nil
}
