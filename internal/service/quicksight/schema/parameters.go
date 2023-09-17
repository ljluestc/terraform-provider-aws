// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package schemaimport (
"time""github.com/YakDriver/regexache"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/quicksight"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
"github.com/hashicorp/terraform-provider-aws/internal/flex"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
)func ParametersSchema() *schema.Schema {
return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Parameters.html
Type:schema.TypeList,
xItems: 1,
tional: true,
mputed: true,
em: &schema.Resource{
Schema: map[string]*schema.Schema{
"date_time_parameters": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DateTimeParameter.html
Type:schema.TypeList,
MinItems: 1,
MaxItems: 100,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
ame": stringSchema(true, validation.StringMatch(regexache.MustCompile(`.*\S.*`), "")),
alues": {
Type:schema.TypeList,
MinItems: 1,
Required: true,
Elem: &schema.Schema{
Type:schema.TypeString,
Validate
func: verify.ValidUTCTimestamp,
},},
},
},
"decimal_parameters": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalParameter.html
Type:schema.TypeList,
MinItems: 1,
MaxItems: 100,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
ame": stringSchema(true, validation.StringMatch(regexache.MustCompile(`.*\S.*`), "")),
alues": {
Type:schema.TypeList,
MinItems: 1,
Required: true,
Elem: &schema.Schema{
Type: schema.TypeFloat,
},},
},
},
"integer_parameters": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_IntegerParameter.html
Type:schema.TypeList,
MinItems: 1,
MaxItems: 100,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
ame": stringSchema(true, validation.StringMatch(regexache.MustCompile(`.*\S.*`), "")),
alues": {
Type:schema.TypeList,
MinItems: 1,
Required: true,
Elem: &schema.Schema{
Type: schema.TypeInt,
},},
},
},
"string_parameters": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StringParameter.html
Type:schema.TypeList,
MinItems: 1,
MaxItems: 100,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
ame": stringSchema(true, validation.StringMatch(regexache.MustCompile(`.*\S.*`), "")),
alues": {
Type:schema.TypeList,
MinItems: 1,
Required: true,
Elem: &schema.Schema{
Type: schema.TypeString,
},},
},
},
},}
}func ExpandParameters(tfList []interface{}) *quicksight.Parameters {
if len(tfList) == 0 || tfList[0] == nil {
turn nil
}tfMap, ok := tfList[0].(map[string]interface{})
if !ok {
turn nil
}parameters := &quicksight.Parameters{}if v, ok := tfMap["date_time_parameters"].([]interface{}); ok && len(v) > 0 {
rameters.DateTimeParameters = expandDateTimeParameters(v)
}
if v, ok := tfMap["decimal_parameters"].([]interface{}); ok && len(v) > 0 {
rameters.DecimalParameters = expandDecimalParameters(v)
}
if v, ok := tfMap["integer_parameters"].([]interface{}); ok && len(v) > 0 {
rameters.IntegerParameters = expandIntegerParameters(v)
}
if v, ok := tfMap["string_parameters"].([]interface{}); ok && len(v) > 0 {
rameters.StringParameters = expandStringParameters(v)
}return parameters
}func expandDateTimeParameters(tfList []interface{}) []*quicksight.DateTimeParameter {
if len(tfList) == 0 {
turn nil
}var parameters []*quicksight.DateTimeParameter
for _, tfMapRaw := range tfList {
Map, ok := tfMapRaw.(map[string]interface{})
 !ok {
continue
ameter := expandDateTimeParameter(tfMap)
 parameter == nil {
continue
ameters = append(parameters, parameter)
}return parameters
}func expandDateTimeParameter(tfMap map[string]interface{}) *quicksight.DateTimeParameter {
if tfMap == nil {
turn nil
}parameter := &quicksight.DateTimeParameter{}if v, ok := tfMap["name"].(string); ok && v != "" {
rameter.Name = aws.String(v)
}
if v, ok := tfMap["values"].([]interface{}); ok && len(v) > 0 {
rameter.Values = flex.ExpandStringTimeList(v, time.RFC3339)
}return parameter
}func expandDecimalParameters(tfList []interface{}) []*quicksight.DecimalParameter {
if len(tfList) == 0 {
turn nil
}var parameters []*quicksight.DecimalParameter
for _, tfMapRaw := range tfList {
Map, ok := tfMapRaw.(map[string]interface{})
 !ok {
continue
ameter := expandDecimalParameter(tfMap)
 parameter == nil {
continue
ameters = append(parameters, parameter)
}return parameters
}func expandDecimalParameter(tfMap map[string]interface{}) *quicksight.DecimalParameter {
if tfMap == nil {
turn nil
}parameter := &quicksight.DecimalParameter{}if v, ok := tfMap["name"].(string); ok && v != "" {
rameter.Name = aws.String(v)
}
if v, ok := tfMap["values"].([]interface{}); ok && len(v) > 0 {
rameter.Values = flex.ExpandFloat64List(v)
}return parameter
}func expandIntegerParameters(tfList []interface{}) []*quicksight.IntegerParameter {
if len(tfList) == 0 {
turn nil
}var parameters []*quicksight.IntegerParameter
for _, tfMapRaw := range tfList {
Map, ok := tfMapRaw.(map[string]interface{})
 !ok {
continue
ameter := expandIntegerParameter(tfMap)
 parameter == nil {
continue
ameters = append(parameters, parameter)
}return parameters
}func expandIntegerParameter(tfMap map[string]interface{}) *quicksight.IntegerParameter {
if tfMap == nil {
turn nil
}parameter := &quicksight.IntegerParameter{}if v, ok := tfMap["name"].(string); ok && v != "" {
rameter.Name = aws.String(v)
}
if v, ok := tfMap["values"].([]interface{}); ok && len(v) > 0 {
rameter.Values = flex.ExpandInt64List(v)
}return parameter
}func expandStringParameters(tfList []interface{}) []*quicksight.StringParameter {
if len(tfList) == 0 {
turn nil
}var parameters []*quicksight.StringParameter
for _, tfMapRaw := range tfList {
Map, ok := tfMapRaw.(map[string]interface{})
 !ok {
continue
ameter := expandStringParameter(tfMap)
 parameter == nil {
continue
ameters = append(parameters, parameter)
}return parameters
}func expandStringParameter(tfMap map[string]interface{}) *quicksight.StringParameter {
if tfMap == nil {
turn nil
}parameter := &quicksight.StringParameter{}if v, ok := tfMap["name"].(string); ok && v != "" {
rameter.Name = aws.String(v)
}
if v, ok := tfMap["values"].([]interface{}); ok && len(v) > 0 {
rameter.Values = flex.ExpandStringList(v)
}return parameter
}
