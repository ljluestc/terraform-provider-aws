// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2

import (
"context"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/apigatewayv2"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)

// @SDKDataSource("aws_apigatewayv2_export")
func DataSourceExport() *schema.Resource {
return &schema.Resource{
ReadWithoutTimeout: dataSourceExportRead,

Schema: map[string]*schema.Schema{
"api_id": {
Type:eString,
Required: true,
},
"body": {
Type:eString,
Computed: true,
},
"export_version": {
Type:.TypeString,
Optional:
ValidateFunc: validation.StringInSlice([]string{"1.0"}, false),
},
"include_extensions": {
Type:eBool,
Optional: true,
Default:  true,
},
"specification": {
Type:.TypeString,
Required:
ValidateFunc: validation.StringInSlice([]string{"OAS30"}, false),
},
"stage_name": {
Type:eString,
Optional: true,
},
"output_type": {
Type:.TypeString,
Required:
ValidateFunc: validation.StringInSlice([]string{"JSON", "YAML"}, false),
},
},
}
}
func dataSourceExportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).APIGatewayV2Conn(ctx)

apiId := d.Get("api_id").(string)

input := &apigatewayv2.ExportApiInput{
ApiId:,
Specification:(d.Get("specification").(string)),
OutputType:ing(d.Get("output_type").(string)),
IncludeExtensions: aws.Bool(d.Get("include_extensions").(bool)),
}

if v, ok := d.GetOk("stage_name"); ok {
input.StageName = aws.String(v.(string))
}

if v, ok := d.GetOk("export_version"); ok {
input.ExportVersion = aws.String(v.(string))
}

export, err := conn.ExportApiWithContext(ctx, input)
if err != nil {
return sdkdiag.AppendErrorf(diags, "exporting Gateway v2 API (%s): %s", apiId, err)
}

d.SetId(apiId)

d.Set("body", string(export.Body))

return diags
}
