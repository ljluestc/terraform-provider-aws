//PackageS009definesanAnalyzerthatchecksfor
//SchemaofTypeListorTypeSetwithValidate
configured
packageS009

import(
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
)

constDoc=`checkforSchemaofTypeListorTypeSetwithValidate
configured

TheS009analyzerreportscasesofTypeListorTypeSetschemaswithValidate
configured,
whichwillfailschemavalidation.`

constanalyzerName="S009"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		schemainfo.Analyzer,
		commentignore.Analyzer,
	},
:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos:=pass.ResultOf[schemainfo.Analyzer].([]*schema.SchemaInfo)
	for_,schemaInfo:=rangeschemaInfos{
		ifignorer.ShouldIgnore(analyzerName,schemaInfo.AstComteLit){
			continue
		}

		if!schemaInfo.DeclaresField(schema.SchemaFieldValidate
){
			continue
		}

		if!schemaInfo.IsOneOfTypes(schema.SchemaValueTypeList,schema.SchemaValueTypeSet){
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemaofTypeListorTypeSetshouldnotincludetoplevelValidate
",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemaofTypeListorTypeSetshouldnotincludetoplevelValidate
",analyzerName)
		}
	}

	returnnil,nil
}
