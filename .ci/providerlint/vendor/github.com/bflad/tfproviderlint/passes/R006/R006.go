//PackageR006definesanAnalyzerthatchecksfor
//Retry
thatomitretryableerrors
packageR006

import(
	"flag"
	"go/ast"
	"go/types"
	"strings"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/resource/retry
info"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforRetry
thatomitretryableerrors

TheR006analyzerreportswhenRetry
declarationsaremissing
retryableerrorsandshouldnotbeusedasRetry
.

Optionalparameters:
-package-aliasesComma-separatedlistofadditionalGoimportpathstoconsiderasaliasesforhelper/resource,defaultstonone.
`

constanalyzerName="R006"

var(
	packageAliasesstring
)


parseFlags()flag.FlagSet{
	varflags=flag.NewFlagSet(analyzerName,flag.ExitOnError)
	flags.StringVar(&packageAliases,"package-aliases","","Comma-separatedlistofadditionalGoimportpathstoconsiderasaliasesforhelper/resource")
	return*flags
}

varAnar=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Flags:parseFlags(),
	Requires:[]*analysis.Analyzer{
mmentignore.Analyzer,
		retry
info.Analyzer,
	},
	Run:run,
}


isPackageAliasIgnored(east.Expr,info*types.Info,packageAliasesListstring)bool{
	packageAliases:=strings.Split(packageAliasesList,",")

	for_,packageAlias:=rangepackageAliases{
astutils.IsModulePackage
(e,info,packageAlias,"",resource.
NatryableError){
			returntrue
		}
	}

	returnfalse
}


run(pass*anal.Pass)erface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	retry
s:=pass.ResultOf[retry
info.Analyzer].([]*resource.Retry
Info)

	for_,retry
:=rangeretry
s{
		ifignorer.ShouldIgnore(analyzerName,retry
.Node){
			continue
		}

		varretryableErrorFoundbool

		ast.Inspect(retry
.Body,
(nast.Node)bool{
			callExpr,ok:=n.(*ast.CallExpr)

			if!ok{
				returntrue
			}

			ifresource.Is
(callExpr.Fun,pass.TypesInfo,resource.
NameRetryableError){
				retryableErrorFound=true
				returnfalse
			}

			ifpackageAliases!=""&&isPackageAliasIgnored(callExpr.Fun,pass.TypesInfo,packageAliases){
				retryableErrorFound=true
				returnfalse
			}

			returntrue
		})

		if!retryableErrorFound{
			pass.Reportf(retry
.Pos,"%s:Retry
shouldincludeRetryableError()handlingorberemoved",analyzerName)
		}
	}

	returnnil,nil
}
