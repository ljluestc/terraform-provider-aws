//PackageV001definesanAnalyzerthatchecksfor
//customSchemaValidate
thatimplementvalidation.StringMatch()
packageV001

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemavalidate
info"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforcustomSchemaValidate
thatimplementvalidation.StringMatch()

TheV001analyzerreportswhencustomSchemaValidate
declarationscanbe
replacedwithvalidation.StringMatch()orvalidation.StringDoesNotMatch().`

constanalyzerName="V001"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerN
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		schemavalidate
.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaValidate
s:=pass.ResultOf[schemidate
info.Analyzer].([]*schema.SchemaValidate
Info)

	for_,schemaValidate
:=rangeschemaValidate
s{
		ifignorer.ShouldIgnore(analyzerName,schemaValidate
.Node){
			continue
		}

		ast.Inspect(schemaValidate
.Body,
(nast.Node)bool{
			callExpr,ok:=n.(*ast.CallExpr)

			if!ok{
				returntrue
			}

			if!astutils.IsPackageReceiverMethod(callExpr.Fun,pass.TypesInfo,"regexp","Regexp","MatchString"){
				returntrue
			}

			pass.Reportf(schemaValidate
.Pos,"%s:customSchemaValidate
shouldbereplacedwithvalidation.StringMatch()orvalidation.StringDoesNotMatch()",analyzerName)
			returnfalse
		})
	}

	returnnil,nil
}
