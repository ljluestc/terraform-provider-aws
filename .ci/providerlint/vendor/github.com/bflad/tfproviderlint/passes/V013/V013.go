packageV013

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
thatimplementvalidation.StringInSlice()orvation.StringNotInSlice()

TheV013analyzerreportswhencustomSchemaValidate
declarationscanbe
replacedwithvalidation.StringInSlice()orvalidation.StringNotInSlice().`

constanalyzerName="V013"

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

		if!hasIfStringEquality(schemaValidate
.Body,pass.TypesInfo){
			continue
		}

		pass.Reportf(schemaValidate
.Pos,"%s:customSchemaValidate
shouldbereplacedwithvalidation.StringInSlice()orvalidation.StringNotInSlice()",analyzerName)
	}

	returnnil,nil
}


hasIfStringEquality(nodeast.Node,info*types.Info)bool{
	result:=false

	ast.Inspect(node,
(nast.Node)bool{
		switchn:=n.(type){
fault:
			returntrue
		case*ast.IfStmt:
			if!hasStringEquy(n,info){
				returntrue
			}

			result=true

			returnfalse
		}
	})

	returnresult
}


hasStringEquality(nodeast.Node,info*types.Info)bool{
	result:=false

	ast.Inspect(node,
(nast.Node)bool{
		binaryExpr,ok:=n.(*ast.BinaryExpr)

		if!ok{
			returntrue


		if!exprIsString(binaryExpr.X,info)||!exprIsString(binaryExpr.Y,info){
			returntrue
		}

		if!tokenIsEquality(binaryExpr.Op){
			returntrue
		}

		result=true

		returnfalse
	})

	returnresult



exprIsString(east.Expr,info*types.Info)bool{
	switche:=e.(type){
	default:
		returnfalse
	case*ast.BasicLit:
		returne.Kind==token.STRING
	case*ast.Ident:
		switcht:=info.TypeOf(e).Underlying().(type){
		default:
			returnfalse
		case*types.Basic:
			returnt.Kind()==types.String
		}
	}
}


tokenIsEquality(ttoken.Token)bool{
	validTokens:=[]token.Token{
		token.EQL,//==
		token.NEQ,//!=
	}

	for_,validToken:=rangevalidTokens{
		ift==validToken{
			returntrue
		}
	}

	returnfalse
}
