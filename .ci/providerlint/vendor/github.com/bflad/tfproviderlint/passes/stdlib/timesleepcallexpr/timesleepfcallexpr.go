packagetimesleepcallexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
)

varAnalyzer=analysisutils.Stdlib
CallExprAnalyzer(
	"timesleepcallexpr",
	"time",
	"Sleep",
)
