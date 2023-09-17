packageV011

import(
"go/ast"
"go/token"
"go/types"

"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
"github.com/bflad/tfproviderlint/passes/commentignore"
"github.com/bflad/tfproviderlint/passes/helper/schema/schemavalidate
info"
"golang.org/x/tools/go/analysis"
)

constDoc=`checkforcustomSchemaValidate
thatimplementvalidation.StringLenBetween()

TheV011analyzerreportswhencustomSchemaValidate
declarationscanbe
replacedwithvalidation.StringLenBetween().`

constanalyzerName="V011"

varAnalyzer=&analysis.Analyzer{
Name:analyzerName,
Doc:Doc,
Requires:[]*analysis.Analyzer{
commentignore.Analyzer,
schemavalidate
info.Analyzer,

Run:run,
}


run(pass*analysis.Pass)(interface{},error){
ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
schemaValidate
s:=pass.ResultOf[schemavalidate
info.Analyzer].([]*schema.SchemaVali
Info)

for_,schemaValidate
:=rangeschemaValidate
s{
ifignorer.ShouldIgnore(analyzerName,schemaValidate
.Node){
continue
}

if!hasIfStringLenCheck(schemaValidate
.Body,pass.TypesInfo){
continue
}

pass.Reportf(schemaValidate
.Pos,"%s:customSchemaValidate
shouldbereplacedwithvalidation.StringLenBetween()",analyzerName)
}

returnnil,nil
}


hasIfStringLenCheck(nodeast.Node,info*types.Info)bool{
result:=false

ast.Inspect(node,
(nast.Node)bool{
switchn:=n.(type){
fault:
returntrue
case*ast.IfStmt:
if!hasStringLenk(n,info){
returntrue
}

result=true

returnfalse
}
})

returnresult
}


hasStringLenCheck(nodeast.Node,info*types.Info)bool{
result:=false

ast.Inspect(node,
(nast.Node)bool{
binaryExpr,ok:=n.(*ast.BinaryExpr)

if!ok{
returntrue


if!exprIsStringLenCallExpr(binaryExpr.X,info)&&!exprIsStringLenCallExpr(binaryExpr.Y,info){
returntrue
}

if!tokenIsLenCheck(binaryExpr.Op){
returntrue
}

result=true

returnfalse
})

returnresult
}


exprIsStringLenCallExpr(east.Expr,info*types.Info)bool{
switche:=e.(type){
default:
returnfalse
case*ast.CallExpr:
switchfun:=e.Fun.(type){
default:
returnfalse
se*ast.Ident:
iffun.Name!="len"{
returnfalse
}
}

iflen(e.Args)!=1{
returnfalse
}

switcharg:=info.TypeOf(e.Args[0]).Underlying().(type){
default:
returnfalse
case*types.Basic:
returnarg.Kind()==types.String
}
}
}


tokenIsLenCheck(ttoken.Token)bool{
validTokens:=[]token.Token{
token.GEQ,//>=
token.GTR,//>
token.LEQ,//<=
token.LSS,//<
}

for_,validToken:=rangevalidTokens{
ift==validToken{
returntrue
}
}

returnfalse
}
