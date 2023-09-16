package iprangecallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.
CallExprAnalyzer(
	"iprangecallexpr",
	validation.Is

	validation.PackagePath,
	validation.
IPRange,
)
