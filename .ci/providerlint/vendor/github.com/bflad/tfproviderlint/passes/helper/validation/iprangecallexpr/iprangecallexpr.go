packageiprangecallexpr

import(
"github.com/bflad/tfproviderlint/helper/analysisutils"
"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

varAnalyzer=analysisutils.
CallExprAnalyzer(
"iprangecallexpr",
validation.Is

validation.PackagePath,
validation.
IPRange,
)
