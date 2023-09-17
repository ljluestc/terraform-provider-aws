packageS029

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfocomputedonly"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforSchemawithonlyComputedenabledandExactlyOneOfconfigured

TheS029analyzerreportscasesofschemaswhichonlyenablesComputed
andconfiguresExactlyOneOf,whichisnotvalid.`

constanalyzerName="S029"

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

		ifschemaInfo.Schema.ExactlyOneOf==nil{
			continue
		}

		switcht:=schemaInfo.AstCompositeLit.Type.(type){
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemashouldnotonlyenableComputedandconfigureExactlyOneOf",analyzerName)
		case*ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(),"%s:schemashouldnotonlyenableComputedandconfigureExactlyOneOf",analyzerName)
		}
	}

	returnnil,nil
}
