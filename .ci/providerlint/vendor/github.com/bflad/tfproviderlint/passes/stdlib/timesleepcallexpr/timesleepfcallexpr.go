package timesleepcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.Stdlib
CallExprAnalyzer(
	"timesleepcallexpr",
	"time",
	"Sleep",
)
