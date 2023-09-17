//PackageS013definesanAnalyzerthatchecksfor
//SchemathatoneofComputed,Optional,orRequired
//isnotconfigured
packageS013

import(
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemamapcompositelit"
)

constDoc=`checkforSchemathataremissingrequiredfields

TheS013analyzerreportscasesofschemaswhichoneofComputed,
Optional,orRequiredisnotconfigured,whichwillfailprovider
schemavalidation.`

constanalyzerName="S013"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		schemamapcompositelit.Analyzer,
		commentignore.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemamapcompositelits:=pass.ResultOf[schemamapcompositelit.Analyzer].([]*ast.CompositeLit)

	for_,smap:=rangeschemamapcompositelits{
		for_,schemaCompositeLit:=rangeschema.GetSchemaMapSchemas(smap){
			schemaInfo:=schema.NewSchemaInfo(schemaCompositeLit,pass.TypesInfo)

			ifignorer.ShouldIgnore(analyzerName,schemaInfo.AstCompositeLit){
				continue
			}

			ifschemaInfo.Schema.Computed||schemaInfo.Schema.Optional||schemaInfo.Schema.Required{
				continue
			}

			switcht:=schemaInfo.AstCompositeLit.Type.(type){
			default:
				pass.Reportf(schemaInfo.AstCompositeLit.Lbrace,"%s:schemashouldconfigureoneofComputed,Optional,orRequired",analyzerName)
			case*ast.SelectorExpr:
				pass.Reportf(t.Sel.Pos(),"%s:schemashouldconfigureoneofComputed,Optional,orRequired",analyzerName)
			}
		}
	}

	returnnil,nil
}
