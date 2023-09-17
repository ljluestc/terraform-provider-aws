//PackageAT002definesanAnalyzerthatchecksfor
//acceptancetestnamesincludingthewordimport
packageAT002

import(
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
)

constDoc=`checkforacceptancetest
tionnamesincludingthewordimport

TheAT002analyzerreportswherethewordimportorImportisused
inanacceptancetest
tionname,whichgenerallymeansthereisanextraneous
acceptancetest.ImportStatetestingshouldbeincludedasaTestStepwitheach
applicableacceptancetest,ratherthanaseparatetestthatonlyverifiesimport
ofasingletestconfiguration.`

constanalyzerName="AT002"

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
pass.ResultOstacc
decl.Analyzer].([]*ast.
Decl)
	for_,testAcc
:=rangetestAcc
s{
		ifignorer.ShouldIgnore(analyzerName,testAcc
){
			continue
		}

		
Name:=testAcc
.Name.Name

		ifstrings.Contains(
Name,"_import")||strings.Contains(
Name,"_Import"){
			pass.Reportf(testAcc
.Name.NamePos,"%s:acceptancetest
tionnameshouldnotincludeimport",analyzerName)
		}
	}

	returnnil,nil
}
