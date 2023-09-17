packagevalidaterfc3339timestringselectorexpr

import(
"github.com/bflad/tfproviderlint/helper/analysisutils"
"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

varAnalyzer=analysisutils.SelectorExprAnalyzer(
"validaterfc3339timestringselectorexpr",
validation.Is

validation.PackagePath,
validation.
ValidateRFC3339TimeString,
)
