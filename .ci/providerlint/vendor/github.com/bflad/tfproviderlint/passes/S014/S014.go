//PackageS014definesanAnalyzerthatchecksfor
//SchemathatwithinElem,Computed,Optional,andRequired
//arenotconfigured
packageS014

import(
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemamapcompositelit"
)

constDoc=`checkforSchemathatElemdoesnotcontainextraneousfields

TheS014analyzerreportscasesofschemaswhichwithinElem,that
Computed,Optional,andRequiredarenotconfigured,whichwillfail
providerschemavalidation.`

constanalyzerName="S014"

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

			elemKvExpr:=schemaInfo.Fields[schema.SchemaFieldElem]

			ifelemKvExpr==nil{
				continue
			}

			//searchwithinElem
			switchelemValue:=elemKvExpr.Value.(type){
			default:
				continue
			case*ast.UnaryExpr:
				ifelemValue.Op!=token.AND||!schema.IsTypeSchema(pass.TypesInfo.TypeOf(elemValue.X)){
					continue
				}

				switchtElemSchema:=elemValue.X.(type){
				default:
					continue
				case*ast.CompositeLit:
					elemSchema:=schema.NewSchemaInfo(tElemSchema,pass.TypesInfo)

					for_,field:=range[]string{schema.SchemaFieldComputed,schema.SchemaFieldOptional,schema.SchemaFieldRequired}{
						ifkvExpr:=elemSchema.Fields[field];kvExpr!=nil{
							pass.Reportf(kvExpr.Pos(),"%s:schemawithinElemshouldnotconfigureComputed,Optional,orRequired",analyzerName)
							break
						}
					}
				}
			}
		}
	}

	returnnil,nil
}
