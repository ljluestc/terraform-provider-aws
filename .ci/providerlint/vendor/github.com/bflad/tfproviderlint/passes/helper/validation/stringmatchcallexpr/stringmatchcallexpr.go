packagestringmatchcallexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

varAnalyzer=analysisutils.
CallExprAnalyzer(
	"stringmatchcallexpr",
	validation.Is

	validation.PackagePath,
	validation.
StringMatch,
)
