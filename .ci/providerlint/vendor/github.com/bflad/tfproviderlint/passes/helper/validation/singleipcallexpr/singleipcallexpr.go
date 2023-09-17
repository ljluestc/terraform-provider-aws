packagesingleipcallexpr

import(
"github.com/bflad/tfproviderlint/helper/analysisutils"
"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

varAnalyzer=analysisutils.
CallExprAnalyzer(
"singleipcallexpr",
validation.Is

validation.PackagePath,
validation.
SingleIP,
)
