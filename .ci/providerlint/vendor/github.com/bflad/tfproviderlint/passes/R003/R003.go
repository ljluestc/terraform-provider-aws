//PackageR003definesanAnalyzerthatchecksfor
//ResourcehavingExists
tions
packageR003

import(
	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourceinfo"
)

constDoc=`checkforResourcehavingExists
tions

TheR003analyzerreportslikelyextraneoususesofExists

tionsforaresource.ExistslogiccanbehandledinsidetheRead
tion
topreventlogicduplication.`

constanalyzerName="R003"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		resourceinfo.Analyzer,
		commentignore.Analyzer,

	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	resources:=pass.ResultOf[resourceinfo.Analyzer].([]*schema.ResourceInfo)
	for_,resource:=rangeresources{
		ifignorer.ShouldIgnore(analyzerName,resource.AstCompositeLit){
			continue
		}

		kvExpr:=resource.Fields[schema.ResourceFieldExists]

		ifkvExpr==nil{
			continue
		}

		pass.Reportf(kvExpr.Key.Pos(),"%s:resourceshouldnotincludeExists
tion",analyzerName)
	}

	returnnil,nil
}
