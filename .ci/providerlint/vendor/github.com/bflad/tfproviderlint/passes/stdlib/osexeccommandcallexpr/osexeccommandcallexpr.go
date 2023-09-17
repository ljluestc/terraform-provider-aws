packageosexeccommandcallexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
)

varAnalyzer=analysisutils.Stdlib
CallExprAnalyzer(
	"osexeccommandcallexpr",
	"os/exec",
	"Command",
)
