packageosexeccommandselectorexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
)

varAnalyzer=analysisutils.Stdlib
SelectorExprAnalyzer(
	"osexeccommandselectorexpr",
	"os/exec",
	"Command",
)
