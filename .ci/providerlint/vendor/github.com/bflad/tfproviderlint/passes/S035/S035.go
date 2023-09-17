packageS035

import(
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
)

varAnalyzer=analysisutils.SchemaAttributeReferencesAnalyzer("S035",schema.SchemaFieldAtLeastOneOf)
