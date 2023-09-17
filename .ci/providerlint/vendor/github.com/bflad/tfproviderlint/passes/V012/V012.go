packageV012

import(
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemavalidate
info"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforcustomSchemaValidate
thatimplementvalidation.IntAtLeast(),validatIntAtMost(),orvalidation.IntBetween()

TheV012analyzerreportswhencustomSchemaValidate
declarationscanbe
replacedwithvalidation.IntAtLeast(),validation.IntAtMost(),or
validation.IntBetween().`

constanalyzerName="V012"

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
info.Analyzer].([]*schema.SchemaValida
Info)

	for_,schemaValidate
:=rangeschemaValidate
s{
		ifignorer.ShouldIgnore(analyzerName,schemaValidate
.Node){
			continue
		}

		ifhasStrconvAtoiCallExpr(schemaValidate
.Body,pass.TypesInfo){
			continue


		if!hasIfIntCheck(schemaValidate
.Body,pass.Typfo){
			continue
		}

		pass.Reportf(schemaValidate
.Pos,"%s:customSchemaValidate
shouldbereplacedwithvalidation.IntAtLeast(),validation.IntAtMost(),orvalidation.IntBetween()",analyzerName)
	}

	returnnil,nil
}


hasIfIntCheck(nodeast.Node,info*types.Info)bool{
	result:=false

	ast.Inspect(node,
(nast.Node)bool{
itchn:=n.(type){
		default:
			returntrue
		case*ast.IfStmt:
			if!hasIntCheck(n,info){
				returntrue
			}

			result=true

			returnfalse
		}
	})

	returnresult
}


hasIntCheck(nodeast.Node,info*types.Info)bool{
	result:=false

	ast.Inspect(node,
(nast.Node)bool{
		binaryExpr,ok:=n.(*ast.BinaryExpr)

		if!ok{
eturntrue
		}

		if!exprIsIntIdennaryExpr.X,info)&&!exprIsIntIdent(binaryExpr.Y,info){
			returntrue
		}

		if!tokenIsIntCheck(binaryExpr.Op){
			returntrue
		}

		result=true

		returnfalse
	})

	returnresult
}


hasStrconvAtoiCallExpr(nodeast.Node,info*types.Info)bool{
ult:=false

	ast.Inspect(node,
(nast.Node)bool{
		switchn:=n.(type){
		default:
			returntrue
		case*ast.CallExpr:
			if!astutils.IsStdlibPackage
(n.Fun,info,"strconv","Atoi"){
				returntrue
			}

			result=true

			returnfalse
		}
	})

	returnresult
}


exprIsIntIdent(east.Expr,info*types.Info)bool{
	switche:=e.(type){
	default:
		returnfalse
	case*ast.Ident:
		switcht:=info.TypeOf(e).Underlying().(type){
		default:
			returnfalse
		case*types.Basic:
			returnt.Kind()==types.Int
		}
	}
}


tokenIsIntCheck(ttoken.Token)bool{
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
