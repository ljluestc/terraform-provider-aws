package stringdoesnotmatchcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.
CallExprAnalyzer(
	"stringdoesnotmatchcallexpr",
	validation.Is

	validation.PackagePath,
	validation.
StringDoesNotMatch,
)
