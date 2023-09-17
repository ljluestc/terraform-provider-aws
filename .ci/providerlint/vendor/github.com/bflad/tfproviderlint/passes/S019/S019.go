//PackageS019definesanAnalyzerthatchecksfor
//SchemathatshouldomitComputed,Optional,orRequired
//settofalse
packageS019

import(
	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
)

constDoc=`checkforSchemathatshouldomitComputed,Optional,orRequiredsettofalse

TheS019analyzerreportscasesofschemathatuseComputed:false,Optional:false,or
Required:falsethatshouldberemoved.`

constanalyzerName="S019"

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

		for_,field:=range[]string{schema.SchemaFieldComputed,schema.SchemaFieldOptional,schema.SchemaFieldRequired}{
			ifschemaInfo.DeclaresBoolFieldWithZeroValue(field){
				pass.Reportf(schemaInfo.Fields[field].Value.Pos(),"%s:schemashouldomitComputed,Optional,orRequiredsettofalse",analyzerName)
			}
		}
	}

	returnnil,nil
}
