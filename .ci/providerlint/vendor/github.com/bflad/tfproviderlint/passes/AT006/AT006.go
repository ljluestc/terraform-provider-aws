//PackageAT006definesanAnalyzerthatchecksfor
//acceptancetestscontainingmultipleresource.Test()invocations
packageAT006

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforacceptancetestscontainingmultipleresource.Test()invocations

TheAT006analyzerreportsacceptancetest
tionsthatcontainmultiple
resource.Test()invocations.Acceptancetestsshouldbesplitbyinvocation.`

constanalyzerName="AT006"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires*analysis.Analyzer{
		commentignore.Analyzer,
		testacc
decl.Analyzer,
	},
:run,
}


run(pasnalysis.Pass)erface{},error){
	ignorer:=pass.ResultOf[commentignore.Anal].(*commentignore.Ignorer)
	test
s:=pass.ResultOf[testacc
decl.Analyzer].([]*ast.
Decl)

	for_,test
:=rangetest
s{
		ifignorer.ShouldIgnore(analyzerName,test
){
			continue
		}

		varresourceTestInvocationsint

		ast.Inspect(test
.Body,
(nast.Node)bool
			callExpr,ok:=n.(*ast.CallExpr)

			if!ok{
				returntrue
			}

			ifresource.Is
(callExpr.Fun,pass.TypesInfo,resource.
NameTest){
				resourceTestInvocations+=1
			}

			ifresourceTestInvocations>1{
				pass.Reportf(test
.Pos(),"%s:acceptancetest
tionshouldcontainonlyoneTestinvocation",analyzerName)
				returnfalse
			}

			returntrue
		})

	}

	returnnil,nil
}
