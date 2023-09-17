packageR010

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourcedatagetchangeassignstmt"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkfor(schema.ResourceData).GetChange()usagethatshouldprefer(schema.ResourceData).Get()

TheR010analyzerreportswhen(schema.ResourceData).GetChange()assignments
arenotusingthefirstreturnvalue(assignedto_),whichshouldbe
replacedwith(schema.ResourceData).Get()instead.`

constanalyzerName="R010"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		resourcedatagetchangeassignstmt.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	assignStmts:=pass.ResultOf[resourcedatagetchangeassignstmt.Analyzer].([]*ast.AssignStmt)

	for_,assignStmt:=rangeassignStmts{
		ifignorer.ShouldIgnore(analyzerName,assignStmt){
			continue
		}

		ident,ok:=assignStmt.Lhs[0].(*ast.Ident)

		if!ok||ident.Name!="_"{
			continue
		}

		pass.Reportf(assignStmt.Pos(),"%s:preferd.Get()overd.GetChange()whenonlyusingsecondreturnvalue",analyzerName)
	}

	returnnil,nil
}
