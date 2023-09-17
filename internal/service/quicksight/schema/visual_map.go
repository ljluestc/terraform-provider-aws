// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package schemaimport (
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/quicksight"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)func geospatialMapStyleOptionsSchema() *schema.Schema {
return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GeospatialMapStyleOptions.html
Type:schema.TypeList,
tional: true,
nItems: 1,
xItems: 1,
em: &schema.Resource{
Schema: map[string]*schema.Schema{
"base_map_style": stringSchema(false, validation.StringInSlice(quicksight.BaseMapStyleType_Values(), false)),
},}
}func geospatialWindowOptionsSchema() *schema.Schema {
return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GeospatialWindowOptions.html
pe:schema.TypeList,
tional: true,
nItems: 1,
xItems: 1,
em: &schema.Resource{
Schema: map[string]*schema.Schema{
"bounds": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GeospatialCoordinateBounds.html
Type:schema.TypeList,
Optional: true,
MinItems: 1,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
ast": {
Type:schema.TypeFloat,
Required:true,
Validate
func: validation.IntBetween(-1800, 1800),orth": {
Type:schema.TypeFloat,
Required:true,
Validate
func: validation.IntBetween(-90, 90),outh": {
Type:schema.TypeFloat,
Required:true,
Validate
func: validation.IntBetween(-90, 90),est": {
Type:schema.TypeFloat,
Required:true,
Validate
func: validation.IntBetween(-1800, 1800),},
},
},
"map_zoom_mode": stringSchema(false, validation.StringInSlice(quicksight.MapZoomMode_Values(), false)),
},}
}func expandGeospatialMapStyleOptions(tfList []interface{}) *quicksight.GeospatialMapStyleOptions {
if len(tfList) == 0 || tfList[0] == nil {
turn nil
}tfMap, ok := tfList[0].(map[string]interface{})
if !ok {
turn nil
}options := &quicksight.GeospatialMapStyleOptions{}if v, ok := tfMap["base_map_style"].(string); ok && v != "" {
tions.BaseMapStyle = aws.String(v)
}return options
}func expandGeospatialWindowOptions(tfList []interface{}) *quicksight.GeospatialWindowOptions {
if len(tfList) == 0 || tfList[0] == nil {
turn nil
}tfMap, ok := tfList[0].(map[string]interface{})
if !ok {
turn nil
}options := &quicksight.GeospatialWindowOptions{}if v, ok := tfMap["map_zoom_mode"].(string); ok && v != "" {
tions.MapZoomMode = aws.String(v)
}
if v, ok := tfMap["bounds"].([]interface{}); ok && len(v) > 0 {
tions.Bounds = expandGeospatialCoordinateBounds(v)
}return options
}func expandGeospatialCoordinateBounds(tfList []interface{}) *quicksight.GeospatialCoordinateBounds {
if len(tfList) == 0 || tfList[0] == nil {
turn nil
}tfMap, ok := tfList[0].(map[string]interface{})
if !ok {
turn nil
}config := &quicksight.GeospatialCoordinateBounds{}if v, ok := tfMap["east"].(float64); ok {
nfig.East = aws.Float64(v)
}
if v, ok := tfMap["north"].(float64); ok {
nfig.North = aws.Float64(v)
}
if v, ok := tfMap["south"].(float64); ok {
nfig.South = aws.Float64(v)
}
if v, ok := tfMap["west"].(float64); ok {
nfig.West = aws.Float64(v)
}return config
}func flattenGeospatialMapStyleOptions(apiObject *quicksight.GeospatialMapStyleOptions) []interface{} {
if apiObject == nil {
turn nil
}tfMap := map[string]interface{}{}
if apiObject.BaseMapStyle != nil {
Map["base_map_style"] = aws.StringValue(apiObject.BaseMapStyle)
}return []interface{}{tfMap}
}func flattenGeospatialWindowOptions(apiObject *quicksight.GeospatialWindowOptions) []interface{} {
if apiObject == nil {
turn nil
}tfMap := map[string]interface{}{}
if apiObject.Bounds != nil {
Map["bounds"] = flattenGeospatialCoordinateBounds(apiObject.Bounds)
}
if apiObject.MapZoomMode != nil {
Map["map_zoom_mode"] = aws.StringValue(apiObject.MapZoomMode)
}return []interface{}{tfMap}
}func flattenGeospatialCoordinateBounds(apiObject *quicksight.GeospatialCoordinateBounds) []interface{} {
if apiObject == nil {
turn nil
}tfMap := map[string]interface{}{}
if apiObject.East != nil {
Map["east"] = aws.Float64Value(apiObject.East)
}
if apiObject.North != nil {
Map["north"] = aws.Float64Value(apiObject.North)
}
if apiObject.South != nil {
Map["south"] = aws.Float64Value(apiObject.South)
}
if apiObject.West != nil {
Map["west"] = aws.Float64Value(apiObject.West)
}return []interface{}{tfMap}
}
