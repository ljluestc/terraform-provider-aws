//PackageS017definesanAnalyzerthatchecksfor
//SchemaincludingMaxItemsorMinItemswithoutTypeList,
//TypeMap,orTypeSet
packageS017

import(
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
)

constDoc=`checkforSchemaincludingMaxItemsorMinItemswithoutproperType

TheS017analyzerreportscasesofschemaincludingMaxItemsorMinItemswithout
TypeList,TypeMap,orTypeSet,whichwillfailschemavalidation.`

constanalyzerName="S017"

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

		if!schemaInfo.DeclaresField(schema.SchemaFieldMaxItems)&&!schemaInfo.DeclaresField(schema.SchemaFieldMinItems){
			continue
		}

		ifschemaInfo.IsOneOfTypes(schema.SchemaValueTypeList,schema.SchemaValueTypeMap,schema.SchemaValueTypeSet){
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemaMaxItemsorMinItemsshouldonlybeincludedforTypeList,TypeMap,orTypeSet",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemaMaxItemsorMinItemsshouldonlybeincludedforTypeList,TypeMap,orTypeSet",analyzerName)
		}
	}

	returnnil,nil
}
