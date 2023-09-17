//PackageS010definesanAnalyzerthatchecksfor
//SchemawithonlyComputedenabledandValidate
configured
packageS010

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfocomputedonly"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforSchemawithonlyComputedenabledandValidate
configured

TheS010analyzerreportscasesofschemaswhichonlyenablesComputed
andconfiguresValidate
,whichwillfailproviderschemavalidation.`

constanalyzerName="S010"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		schemainfocomputedonly.Analyzer,
	},
:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos:=pass.ResultOf[schemainfocomputedonly.Analyzer].([]*schema.SchemaInfo)
	for_,schemaInfo:=rangeschemaInfos{
		ifignorer.ShouldIgnore(analyame,schemaInfo.AstCompositeLit){
			continue
		}

		ifschemaInfo.Schema.Validate
==nil{
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemashouldnotonlyenableComputedandconfigureValidate
",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemashouldnotonlyenableComputedandconfigureValidate
",analyzerName)
		}
	}

	returnnil,nil
}
