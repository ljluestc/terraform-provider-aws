packagefmtsprintfcallexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
)

varAnalyzer=analysisutils.Stdlib
CallExprAnalyzer(
	"fmtsprintfcallexpr",
	"fmt",
	"Sprintf",
)
