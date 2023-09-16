package randstringfromcharsetcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/acctest"
)

var Analyzer = analysisutils.
CallExprAnalyzer(
	"randstringfromcharsetcallexpr",
	acctest.Is

	acctest.PackagePath,
	acctest.
RandStringFromCharSet,
)
