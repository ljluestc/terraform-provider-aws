// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package connectimport (
"context""github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
)// @SDKDataSource("aws_connect_lambda_function_association")
func DataSourceLambdafunctionAssociation() *schema.Resource {
return &schema.Resource{
ReadWithoutTimeout: dataSourceLambdafunctionAssociationRead,
Schema: map[string]*schema.Schema{
"function_arn": {
Type:schema.TypeString,
Required:true,
Validatefunc: verify.ValidARN,
},
"instance_id": {
Type:schema.TypeString,
Required: true,
},
},
}
}func dataSourceLambdafunctionAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
conn := meta.(*conns.AWSClient).ConnectConn(ctx)
functionArn := d.Get("function_arn")
instanceID := d.Get("instance_id")lfaArn, err := FindLambdafunctionAssociationByARNWithContext(ctx, conn, instanceID.(string), functionArn.(string))
if err != nil {
return diag.Errorf("finding Connect Lambda function Association by ARN (%s): %s", functionArn, err)
}if lfaArn == "" {
return diag.Errorf("finding Connect Lambda function Association by ARN (%s): not found", functionArn)
}d.SetId(meta.(*conns.AWSClient).Region)
d.Set("function_arn", functionArn)
d.Set("instance_id", instanceID)return nil
}
