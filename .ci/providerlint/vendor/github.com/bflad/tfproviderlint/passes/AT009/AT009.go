packageAT009

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/acctest"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/acctest/randstringfromcharsetcallexpr"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforacctest.RandStringFromCharSet()callsthatcanbeacctest.RandString()

TheAT009analyzerreportswherethesecondparameterofa
RandStringFromCharSetcallisacctest.CharSetAlpha,whichisequivalentto
callingRandString.`

constanalyzerName="AT009"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		randstringfromcharsetcallexpr.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	callExprs:=pass.ResultOf[randstringfromcharsetcallexpr.Analyzer].([]*ast.CallExpr)
	for_,callExpr:=rangecallExprs{
		ifignorer.ShouldIgnore(analyzerName,callExpr){
			continue
		}

		iflen(callExpr.Args)<2{
			continue
		}

		if!acctest.IsConst(callExpr.Args[1],pass.TypesInfo,acctest.ConstNameCharSetAlphaNum){
			continue
		}

		pass.Reportf(callExpr.Pos(),"%s:shoulduseRandStringcallinstead",analyzerName)
	}

	returnnil,nil
}
