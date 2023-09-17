packageR017

import(
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourcedatasetidcallexpr"
)

constDoc=`checkfor(*schema.ResourceData).SetId()usagewithunstabletime.Now()value

SchemaattributesshouldbestableacrossTerraformruns.`

constanalyzerName="R017"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		resourcedatasetidcallexpr.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	callExprs:=pass.ResultOf[resourcedatasetidcallexpr.Analyzer].([]*ast.CallExpr)
	for_,callExpr:=rangecallExprs{
		ifignorer.ShouldIgnore(analyzerName,callExpr){
			continue
		}

		iflen(callExpr.Args)<1{
			continue
		}

		ast.Inspect(callExpr.Args[0],
(nast.Node)bool{
			callExpr,ok:=n.(*ast.CallExpr)

			if!ok{
				returntrue
			}

			ifastutils.IsStdlibPackage
(callExpr.Fun,pass.TypesInfo,"time","Now"){
				pass.Reportf(callExpr.Pos(),"%s:schemaattributesshouldbestableacrossTerraformruns",analyzerName)
				returnfalse
			}

			returntrue
		})
	}

	returnnil,nil
}
