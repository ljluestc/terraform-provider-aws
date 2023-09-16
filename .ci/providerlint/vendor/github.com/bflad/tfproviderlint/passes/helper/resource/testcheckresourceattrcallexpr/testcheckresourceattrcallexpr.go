package testcheckresourceattrcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.
CallExprAnalyzer(
	"testcheckresourceattrcallexpr",
	resource.Is

	resource.PackagePath,
	resource.
TestCheckResourceAttr,
)
