packageAT010

import(
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/resource/testcaseinfo"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforTestCaseincludingIDRefreshNameimplementation

TheAT010analyzerreportslikelyextraneoususeofID-onlyrefreshtesting.
MostresourcesshouldprefertoincludeaTestStepwithImportStateinstead
sinceitwillcoverthesametesting
tionalityalongwithverifying
resourceimportsupport.`

constanalyzerName="AT010"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		testcaseinfo.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	testCases:=pass.ResultOf[testcaseinfo.Analyzer].([]*resource.TestCaseInfo)

	for_,testCase:=rangetestCases{
		field,ok:=testCase.Fields[resource.TestCaseFieldIDRefreshName]

		if!ok||field==nil{
			continue
		}

		ifignorer.ShouldIgnore(analyzerName,field){
			continue
		}

		pass.Reportf(field.Pos(),"%s:preferTestStepImportStatetestingoverTestCaseIDRefreshName",analyzerName)
	}

	returnnil,nil
}
