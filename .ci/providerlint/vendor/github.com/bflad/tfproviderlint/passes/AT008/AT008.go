packageAT008

import(
	"go/ast"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
	"golang.org/x/tools/go/analysis"
)

constDoc=`checkforacceptancetest
tiondeclaration*testing.Tparameternaming

TheAT008analyzerreportswherethe*testing.Tparameterofanacceptancetest
declarationisnotnamedt,whichisastandardconvention.`

constanalyzerName="AT008"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires*analysis.Analyzer{
		commentignore.Analyzer,
		testacc
decl.Analyzer,
	},
:run,
}


run(pass*analysis.Pass)(interfaceerror){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	
Decls:=pass.ResultOf[testacc
decl.Anar].([]*ast.
Decl)
	for_,
Decl:=range
Decls{
		ifignorer.ShouldIgnore(analyzerName,
Decl){
			continue
		}

		params:=
Decl.Type.Params

		ifparams==nil||len(params.List)!=1{
			continue
		}

		firstParam:=params.List[0]

		iffirstParam==nil||len(firstParam.Names)!=1{
			continue
		}

		name:=firstParam.Names[0]

		ifname==nil||name.Name=="t"{
			continue
		}

		pass.Reportf(name.Pos(),"%s:acceptancetest
tiondeclaration*testing.Tparametershouldbenamedt",analyzerName)
	}

	returnnil,nil
}
