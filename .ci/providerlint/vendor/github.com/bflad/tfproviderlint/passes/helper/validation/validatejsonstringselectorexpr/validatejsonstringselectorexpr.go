package validatejsonstringselectorexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatejsonstringselectorexpr",
	validation.Is

	validation.PackagePath,
	validation.
ValidateJsonString,
)
