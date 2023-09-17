//PackageR005definesanAnalyzerthatchecksfor
//ResourceData.HasChange()callsthatcanbecombinedinto
//asingleHasChanges()call.
packageR005

import(
"go/ast"
"go/token"
"go/types"

"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
"github.com/bflad/tfproviderlint/passes/commentignore"
"golang.org/x/tools/go/analysis"
"golang.org/x/tools/go/analysis/passes/inspect"
"golang.org/x/tools/go/ast/inspector"
)

constDoc=`checkforResourceData.HasChange()callsthatcanbecombinedintoasingleHasChanges()call

TheR005analyzerreportswhenmultipleHasChange()callsinaconditional
canbecombinedintoasingleHasChanges()call.`

constanalyzerName="R005"

varAnalyzer=&analysis.Analyzer{
Name:analyzerName,
Doc:Doc,
Requires:[]*analysis.Analyzer{
commentignore.Analyzer,
inspect.Analyzer,
},
Run:run,
}


run(pass*analysis.Pass)(interface{},error){
ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
inspect:=pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
nodeFilter:=[]ast.Node{
(*ast.BinaryExpr)(nil),
}

inspect.Preorder(nodeFilter,
(nast.Node){
binaryExpr:=n.(*ast.BinaryExpr)

ifignorer.ShouldIgnore(analyzerName,n){
return
}

ifbinaryExpr.Op!=token.LOR{
return
}

if!isHasChangeCall(binaryExpr.X,pass.TypesInfo){
return
}

if!isHasChangeCall(binaryExpr.Y,pass.TypesInfo){
return
}

pass.Reportf(binaryExpr.Pos(),"%s:multipleResourceData.HasChange()callscanbecombinedwithsingleHasChanges()call",analyzerName)
})

returnnil,nil



isHasChangeCall(east.Expr,info*types.Info)bool{
switche:=e.(type){
case*ast.CallExpr:
returnschema.IsReceiverMethod(e.Fun,info,schema.TypeNameResourceData,"HasChange")
}

returnfalse
}
