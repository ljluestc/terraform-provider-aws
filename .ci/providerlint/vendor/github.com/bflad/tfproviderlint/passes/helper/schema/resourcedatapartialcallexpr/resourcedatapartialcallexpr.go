packageresourcedatapartialcallexpr

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
)

varAnalyzer=analysisutils.ReceiverMethodCallExprAnalyzer(
	"resourcedatapartialcallexpr",
	schema.IsReceiverMethod,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"Partial",
)
