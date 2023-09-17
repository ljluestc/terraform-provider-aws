packagetestmatchresourceattrcallexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
)

varAnalyzer=analysisutils.
CallExprAnalyzer(
	"testmatchresourceattrcallexpr",
	resource.Is

	resource.PackagePath,
	resource.
TestMatchResourceAttr,
)
