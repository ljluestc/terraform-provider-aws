// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package schemaimport (
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/quicksight"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)func emptyVisualSchema() *schema.Schema {
return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_EmptyVisual.html
Type:schema.TypeList,
tional: true,
nItems: 1,
xItems: 1,
em: &schema.Resource{
Schema: map[string]*schema.Schema{
"data_set_identifier": stringSchema(true, validation.StringLenBetween(1, 2048)),
"visual_id":idSchema(),
"actions": visualCustomActionsSchema(customActionsMaxItems), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_VisualCustomAction.html
},}
}func expandEmptyVisual(tfList []interface{}) *quicksight.EmptyVisual {
if len(tfList) == 0 || tfList[0] == nil {
turn nil
}tfMap, ok := tfList[0].(map[string]interface{})
if !ok {
turn nil
}visual := &quicksight.EmptyVisual{}if v, ok := tfMap["data_set_identifier"].(string); ok && v != "" {
sual.DataSetIdentifier = aws.String(v)
}
if v, ok := tfMap["visual_id"].(string); ok && v != "" {
sual.VisualId = aws.String(v)
}
if v, ok := tfMap["actions"].([]interface{}); ok && len(v) > 0 {
sual.Actions = expandVisualCustomActions(v)
}return visual
}func flattenEmptyVisual(apiObject *quicksight.EmptyVisual) []interface{} {
if apiObject == nil {
turn nil
}tfMap := map[string]interface{}{
ata_set_identifier": aws.StringValue(apiObject.DataSetIdentifier),
isual_id":aws.StringValue(apiObject.VisualId),
}
if apiObject.Actions != nil {
Map["actions"] = flattenVisualCustomAction(apiObject.Actions)
}return []interface{}{tfMap}
}
