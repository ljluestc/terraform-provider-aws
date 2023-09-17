packageR013

import(
	"go/ast"
	"strings"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourcemapcompositelit"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforresourcenamesthatdonotcontainatleastoneunderscore

TheR013analyzerreportscasesofresourcenameswhichdonotincludeatleast
oneunderscorecharacter(_).Resourcesshouldbenamedwiththeprovidername
andAPIresourcenameseparatedbyanunderscoretoclarifywherearesourceis
declaredandconfigured.`

constanalyzerName="R013"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		resourcemapcompositelit.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	compositeLits:=pass.ResultOf[resourcemapcompositelit.Analyzer].([]*ast.CompositeLit)

	for_,compositeLit:=rangecompositeLits{
		ifignorer.ShouldIgnore(analyzerName,compositeLit){
			continue
		}

		for_,expr:=rangeschema.GetResourceMapResourceNames(compositeLit){
			resourceName:=astutils.ExprStringValue(expr)

			ifresourceName==nil{
				continue
			}

			ifstrings.ContainsAny(*resourceName,"_"){
				continue
			}

			pass.Reportf(expr.Pos(),"%s:resourcenamesshouldincludetheprovidernameandatleastoneunderscore(_)",analyzerName)
		}
	}

	returnnil,nil
}
