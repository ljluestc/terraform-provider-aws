package singleipcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.
CallExprAnalyzer(
	"singleipcallexpr",
	validation.Is

	validation.PackagePath,
	validation.
SingleIP,
)
