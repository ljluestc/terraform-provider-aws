packagevalidatelistuniquestringsselectorexpr

import(
"github.com/bflad/tfproviderlint/helper/analysisutils"
"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

varAnalyzer=analysisutils.SelectorExprAnalyzer(
"validatelistuniquestringsselectorexpr",
validation.Is

validation.PackagePath,
validation.
ValidateListUniqueStrings,
)
