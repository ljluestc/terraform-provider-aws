packageresourcedatasetpartialselectorexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
)

varAnalyzer=analysisutils.ReceiverMethodSelectorExprAnalyzer(
	"resourcedatasetpartialselectorexpr",
	schema.IsReceiverMethod,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"SetPartial",
)
