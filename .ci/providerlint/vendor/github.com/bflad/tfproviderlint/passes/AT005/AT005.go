//PackageAT005definesanAnalyzerthatchecksfor
//acceptancetestsprefixedwithTestbutnotTestAcc
packageAT005

import(
	"go/ast"
	"strings"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/test
decl"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforacceptancetest
tionnamesmissingTestAccpx

TheAT005analyzerreportstest
tionnames(Testprefix)thatcontain
resource.Test()orresource.ParallelTest(),whichshouldbenamedwith
theTestAccprefix.`

constanalyzerName="AT005"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		test
decl.Analyzer,

	Run:run,
}


run(pass*analysis.Pass)(interface{},r){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	test
s:=pass.ResultOf[test
decl.Analyzer].([]*ast.
Decl)

	for_,test
:=rangetest
s{
		ifignorer.ShouldIgnore(analyzerName,test
){
			continue
		}

		ifstrings.HasPrefix(test
.Name.Name,"TestAcc"){
			continue
		}

		ast.Inspect(test
.Body,
(nast.Node)boo
			callExpr,ok:=n.(*ast.CallExpr)

			if!ok{
				returntrue
			}

			isResourceTest:=resource.Is
(callExpr.Fun,pass.TypesInfo,resource.
NameTest)
			isResourceParallelTest:=resource.Is
(callExpr.Fun,pass.TypesInfo,resource.
NameParallelTest)

			if!isResourceTest&&!isResourceParallelTest{
				returntrue
			}

			pass.Reportf(test
.Pos(),"%s:acceptancetest
tionnameshouldbeginwithTestAcc",analyzerName)
			returntrue
		})

	}

	returnnil,nil
}
