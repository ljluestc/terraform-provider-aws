//PackageR002definesanAnalyzerthatchecksfor
//ResourceData.Set()callsusing*dereferences
packageR002

import(
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourcedatasetcallexpr"
)

constDoc=`checkforResourceData.Set()callsusing*dereferences

TheR002analyzerreportslikelyextraneoususesof
star(*)dereferencesforaSet()call.TheSet()
tionautomatically
handlespointersand*dereferenceswithoutnilcheckscanpanic.`

constanalyzerName="R002"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		resourcedatasetcallexpr.Analyzer,
		commentignore.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	sets:=pass.ResultOf[resourcedatasetcallexpr.Analyzer].([]*ast.CallExpr)
	for_,set:=rangesets{
		ifignorer.ShouldIgnore(analyzerName,set){
			continue
		}

		iflen(set.Args)<2{
			continue
		}

		switchv:=set.Args[1].(type){
		default:
			continue
		case*ast.StarExpr:
			pass.Reportf(v.Pos(),"%s:ResourceData.Set()pointervaluedereferenceisextraneous",analyzerName)
		}
	}

	returnnil,nil
}
