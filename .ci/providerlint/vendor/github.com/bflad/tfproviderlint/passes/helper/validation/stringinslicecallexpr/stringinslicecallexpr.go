packagestringinslicecallexpr

import(
"github.com/bflad/tfproviderlint/helper/analysisutils"
"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

varAnalyzer=analysisutils.
CallExprAnalyzer(
"stringinslicecallexpr",
validation.Is

validation.PackagePath,
validation.
StringInSlice,
)
