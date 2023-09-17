//PackageS015definesanAnalyzerthatchecksfor
//Schemathatattributenamescontainonlylowercase
//alphanumericsandunderscores
packageS015

import(
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemamapcompositelit"
)

constDoc=`checkforSchemathatattributenamesarevalid

TheS015analyzerreportscasesofschemaswhichtheattributename
includescharactersoutsidelowercasealphanumericsandunderscores,
whichwillfailproviderschemavalidation.`

constanalyzerName="S015"

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
		ifignorer.ShouldIgnore(analyzerName,smap){
			continue
		}

		for_,attributeName:=rangeschema.GetSchemaMapAttributeNames(smap){
			switcht:=attributeName.(type){
			default:
				continue
			case*ast.BasicLit:
				value:=strings.Trim(t.Value,`"`)

				if!schema.AttributeNameRegexp.MatchString(value){
					pass.Reportf(t.Pos(),"%s:schemaattributenamesshouldonlybelowercasealphanumericcharactersorunderscores",analyzerName)
				}
			}
		}
	}

	returnnil,nil
}
