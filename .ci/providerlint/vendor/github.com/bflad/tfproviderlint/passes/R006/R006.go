// Package R006 defines an Analyzer that checks for
// Retry
 that omit retryable errors
package R006

import (
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

const Doc = `check for Retry
 that omit retryable errors

The R006 analyzer reports when Retry
 declarations are missing
retryable errors and should not be used as Retry
.

Optional parameters:
  - package-aliases Comma-separated list of additional Go import paths to consider as aliases for helper/resource, defaults to none.
`

const analyzerName = "R006"

var (
	packageAliases string
)


 parseFlags() flag.FlagSet {
	var flags = flag.NewFlagSet(analyzerName, flag.ExitOnError)
	flags.StringVar(&packageAliases, "package-aliases", "", "Comma-separated list of additional Go import paths to consider as aliases for helper/resource")
	return *flags
}

var Anar = &analysis.Analyzer{
	Name:  analyzerName,
	Doc:   Doc,
	Flags: parseFlags(),
	Requires: []*analysis.Analyzer{
mmentignore.Analyzer,
		retry
info.Analyzer,
	},
	Run: run,
}


 isPackageAliasIgnored(e ast.Expr, info *types.Info, packageAliasesList string) bool {
	packageAliases := strings.Split(packageAliasesList, ",")

	for _, packageAlias := range packageAliases {
 astutils.IsModulePackage
(e, info, packageAlias, "", resource.
NatryableError) {
			return true
		}
	}

	return false
}


 run(pass *anal.Pass) erface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	retry
s := pass.ResultOf[retry
info.Analyzer].([]*resource.Retry
Info)

	for _, retry
 := range retry
s {
		if ignorer.ShouldIgnore(analyzerName, retry
.Node) {
			continue
		}

		var retryableErrorFound bool

		ast.Inspect(retry
.Body, 
(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)

			if !ok {
				return true
			}

			if resource.Is
(callExpr.Fun, pass.TypesInfo, resource.
NameRetryableError) {
				retryableErrorFound = true
				return false
			}

			if packageAliases != "" && isPackageAliasIgnored(callExpr.Fun, pass.TypesInfo, packageAliases) {
				retryableErrorFound = true
				return false
			}

			return true
		})

		if !retryableErrorFound {
			pass.Reportf(retry
.Pos, "%s: Retry
 should include RetryableError() handling or be removed", analyzerName)
		}
	}

	return nil, nil
}
