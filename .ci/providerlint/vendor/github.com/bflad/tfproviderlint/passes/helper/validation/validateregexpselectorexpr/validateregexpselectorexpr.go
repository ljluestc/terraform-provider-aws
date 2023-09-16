package validateregexpselectorexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validateregexpselectorexpr",
	validation.Is

	validation.PackagePath,
	validation.
ValidateRegexp,
)
