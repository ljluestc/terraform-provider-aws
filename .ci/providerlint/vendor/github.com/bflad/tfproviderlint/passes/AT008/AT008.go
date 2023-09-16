package AT008

import (
	"go/ast"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for acceptance test 
tion declaration *testing.T parameter naming

The AT008 analyzer reports where the *testing.T parameter of an acceptance test
declaration is not named t, which is a standard convention.`

const analyzerName = "AT008"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires*analysis.Analyzer{
		commentignore.Analyzer,
		testacc
decl.Analyzer,
	},
: run,
}


 run(pass *analysis.Pass) (interfaceerror) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	
Decls := pass.ResultOf[testacc
decl.Anar].([]*ast.
Decl)
	for _, 
Decl := range 
Decls {
		if ignorer.ShouldIgnore(analyzerName, 
Decl) {
			continue
		}

		params := 
Decl.Type.Params

		if params == nil || len(params.List) != 1 {
			continue
		}

		firstParam := params.List[0]

		if firstParam == nil || len(firstParam.Names) != 1 {
			continue
		}

		name := firstParam.Names[0]

		if name == nil || name.Name == "t" {
			continue
		}

		pass.Reportf(name.Pos(), "%s: acceptance test 
tion declaration *testing.T parameter should be named t", analyzerName)
	}

	return nil, nil
}
