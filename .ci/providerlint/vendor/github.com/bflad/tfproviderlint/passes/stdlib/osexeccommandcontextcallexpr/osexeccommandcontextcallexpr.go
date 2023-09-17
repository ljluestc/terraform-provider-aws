packageosexeccommandcontextcallexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
)

varAnalyzer=analysisutils.Stdlib
CallExprAnalyzer(
	"osexeccommandcontextcallexpr",
	"os/exec",
	"CommandContext",
)
