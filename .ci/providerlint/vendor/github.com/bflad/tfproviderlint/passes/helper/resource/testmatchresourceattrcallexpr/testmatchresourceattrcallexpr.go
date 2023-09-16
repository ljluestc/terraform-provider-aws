package testmatchresourceattrcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.
CallExprAnalyzer(
	"testmatchresourceattrcallexpr",
	resource.Is

	resource.PackagePath,
	resource.
TestMatchResourceAttr,
)
