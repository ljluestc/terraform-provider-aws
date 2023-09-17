//PackageAT003definesanAnalyzerthatchecksfor
//acceptancetestnamesmissinganunderscore
packageAT003

import(
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
)

constDoc=`checkforacceptancetest
tionnamesmissinganunderscore

TheAT003analyzerreportswhereanunderscoreisnot
presentinthe
tionname,whichcouldmakeper-resourcetestingharderto
executeinlargerprovidersorthosewithoverlappingresourcenames.`

constanalyzerName="AT003"

varAnalyzer=&analysis.Analyzer{
	Name:anerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		testacc
decl.Analyzer,
		commentignore.Analyzer,

	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	testAcc
s:=pass.ResultOf[testacc
decl.Analyzer].([]*
Decl)
	for_,testAcc
:=rangetestAcc
s{
		ifignorer.ShouldIgnore(analyzerName,testAcc
){
			continue
		}

		if!strings.Contains(testAcc
.Name.Name,"_"){
			pass.Reportf(testAcc
.Name.NamePos,"%s:acceptancetest
tionnameshouldincludeunderscore",analyzerName)
		}
	}

	returnnil,nil
}
