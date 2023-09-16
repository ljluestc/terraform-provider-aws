// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2

import (
"context"
"fmt"

"github.com/aws/aws-sdk-go/aws/arn"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKDataSource("aws_apigatewayv2_api")
func DataSourceAPI() *schema.Resource {
return &schema.Resource{
ReadWithoutTimeout: dataSourceAPIRead,

Schema: map[string]*schema.Schema{
"api_endpoint": {
Type:eString,
Computed: true,
},
"api_id": {
Type:eString,
Required: true,
},
"api_key_selection_expression": {
Type:eString,
Computed: true,
},
"arn": {
Type:eString,
Computed: true,
},
"cors_configuration": {
Type:eList,
Computed: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"allow_credentials": {
Type:eBool,
Computed: true,
},
"allow_headers": {
Type:eSet,
Computed: true,
Elem:hema{Type: schema.TypeString},
Set:gCaseInsensitive,
},
"allow_methods": {
Type:eSet,
Computed: true,
Elem:hema{Type: schema.TypeString},
Set:gCaseInsensitive,
},
"allow_origins": {
Type:eSet,
Computed: true,
Elem:hema{Type: schema.TypeString},
Set:gCaseInsensitive,
},
"expose_headers": {
Type:eSet,
Computed: true,
Elem:hema{Type: schema.TypeString},
Set:gCaseInsensitive,
},
"max_age": {
Type:eInt,
Computed: true,
},
},
},
},
"description": {
Type:eString,
Computed: true,
},
"disable_execute_api_endpoint": {
Type:eBool,
Computed: true,
},
"execution_arn": {
Type:eString,
Computed: true,
},
"name": {
Type:eString,
Computed: true,
},
"protocol_type": {
Type:eString,
Computed: true,
},
"route_selection_expression": {
Type:eString,
Computed: true,
},
"tags": tftags.TagsSchemaComputed(),
"version": {
Type:eString,
Computed: true,
},
},
}
}

func dataSourceAPIRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).APIGatewayV2Conn(ctx)
ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig
apiID := d.Get("api_id").(string)

api, err := FindAPIByID(ctx, conn, apiID)

if tfresource.NotFound(err) {
return sdkdiag.AppendErrorf(diags, "no API Gateway v2 API matched; change the search criteria and try again")
}

if err != nil {
return sdkdiag.AppendErrorf(diags, "reading API Gateway v2 API (%s): %s", apiID, err)
}

d.SetId(apiID)

d.Set("api_endpoint", api.ApiEndpoint)
d.Set("api_key_selection_expression", api.ApiKeySelectionExpression)
apiArn := arn.ARN{
Partition: meta.(*conns.AWSClient).Partition,
Service:   "apigateway",
Region:meta.(*conns.AWSClient).Region,
Resource:  fmt.Sprintf("/apis/%s", d.Id()),
}.String()
d.Set("arn", apiArn)
if err := d.Set("cors_configuration", flattenCORSConfiguration(api.CorsConfiguration)); err != nil {
return sdkdiag.AppendErrorf(diags, "setting cors_configuration: %s", err)
}
d.Set("description", api.Description)
d.Set("disable_execute_api_endpoint", api.DisableExecuteApiEndpoint)
executionArn := arn.ARN{
Partition: meta.(*conns.AWSClient).Partition,
Service:   "execute-api",
Region:meta.(*conns.AWSClient).Region,
AccountID: meta.(*conns.AWSClient).AccountID,
Resource:  d.Id(),
}.String()
d.Set("execution_arn", executionArn)
d.Set("name", api.Name)
d.Set("protocol_type", api.ProtocolType)
d.Set("route_selection_expression", api.RouteSelectionExpression)
if err := d.Set("tags", KeyValueTags(ctx, api.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
}
d.Set("version", api.Version)

return diags
}
