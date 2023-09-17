// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package connectimport (
"context"
"log""github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/connect"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
)// @SDKResource("aws_connect_lambda_function_association")
func ResourceLambdafunctionAssociation() *schema.Resource {
return &schema.Resource{
CreateWithoutTimeout: resourceLambdafunctionAssociationCreate,
ReadWithoutTimeout:resourceLambdafunctionAssociationRead,
DeleteWithoutTimeout: resourceLambdafunctionAssociationDelete,
Importer: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,
},
Schema: map[string]*schema.Schema{
"function_arn": {
Type:schema.TypeString,
Required:true,
ForceNew:true,
Validatefunc: verify.ValidARN,
},
"instance_id": {
Type:schema.TypeString,
Required: true,
ForceNew: true,
},
},
}
}func resourceLambdafunctionAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
conn := meta.(*conns.AWSClient).ConnectConn(ctx)instanceId := d.Get("instance_id").(string)
functionArn := d.Get("function_arn").(string)input := &connect.AssociateLambdafunctionInput{
InstanceId:aws.String(instanceId),
functionArn: aws.String(functionArn),
}_, err := conn.AssociateLambdafunctionWithContext(ctx, input)
if err != nil {
return diag.Errorf("creating Connect Lambda function Association (%s,%s): %s", instanceId, functionArn, err)
}d.SetId(LambdafunctionAssociationCreateResourceID(instanceId, functionArn))return resourceLambdafunctionAssociationRead(ctx, d, meta)
}func resourceLambdafunctionAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
conn := meta.(*conns.AWSClient).ConnectConn(ctx)instanceID, functionArn, err := LambdafunctionAssociationParseResourceID(d.Id())if err != nil {
return diag.FromErr(err)
}lfaArn, err := FindLambdafunctionAssociationByARNWithContext(ctx, conn, instanceID, functionArn)if !d.IsNewResource() && tfresource.NotFound(err) {
log.Printf("[WARN] Connect Lambda function Association (%s) not found, removing from state", d.Id())
d.SetId("")
return nil
}if err != nil {
return diag.Errorf("finding Connect Lambda function Association by function ARN (%s): %s", functionArn, err)
}d.Set("function_arn", lfaArn)
d.Set("instance_id", instanceID)return nil
}func resourceLambdafunctionAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
conn := meta.(*conns.AWSClient).ConnectConn(ctx)instanceID, functionArn, err := LambdafunctionAssociationParseResourceID(d.Id())
if err != nil {
return diag.FromErr(err)
}input := &connect.DisassociateLambdafunctionInput{
InstanceId:aws.String(instanceID),
functionArn: aws.String(functionArn),
}_, err = conn.DisassociateLambdafunctionWithContext(ctx, input)if tfawserr.ErrCodeEquals(err, connect.ErrCodeResourceNotFoundException) {
return nil
}if err != nil {
return diag.Errorf("deleting Connect Lambda function Association (%s): %s", d.Id(), err)
}return nil
}
