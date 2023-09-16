package stringmatchcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.
CallExprAnalyzer(
	"stringmatchcallexpr",
	validation.Is

	validation.PackagePath,
	validation.
StringMatch,
)
