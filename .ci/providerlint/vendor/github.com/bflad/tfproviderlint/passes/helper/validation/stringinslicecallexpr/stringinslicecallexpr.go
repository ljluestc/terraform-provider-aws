package stringinslicecallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.
CallExprAnalyzer(
	"stringinslicecallexpr",
	validation.Is

	validation.PackagePath,
	validation.
StringInSlice,
)
