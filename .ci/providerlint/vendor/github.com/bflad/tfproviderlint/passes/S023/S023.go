//PackageS023definesanAnalyzerthatchecksfor
//SchemathatshouldomitElemwithincompatibleType
packageS023

import(
	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
)

constDoc=`checkforSchemathatshouldomitElemwithincompatibleType

TheS023analyzerreportscasesofschemathatdeclareElemthatshould
beremovedwithincompatibleType.`

constanalyzerName="S023"

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

		ifschemaInfo.IsOneOfTypes(schema.SchemaValueTypeList,schema.SchemaValueTypeMap,schema.SchemaValueTypeSet){
			continue
		}

		if!schemaInfo.DeclaresField(schema.SchemaFieldElem){
			continue
		}

		pass.Reportf(schemaInfo.Fields[schema.SchemaFieldElem].Value.Pos(),"%s:schemashouldnotincludeElemwithincompatibleType",analyzerName)
	}

	returnnil,nil
}
