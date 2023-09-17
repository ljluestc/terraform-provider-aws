// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package kmsimport (
"context"
"log""github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/kms"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/create"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)// @SDKResource("aws_kms_alias")
func ResourceAlias() *schema.Resource {
return &schema.Resource{
CreateWithoutTimeout: resourceAliasCreate,
adWithoutTimeout:resourceAliasRead,
dateWithoutTimeout: resourceAliasUpdate,
leteWithoutTimeout: resourceAliasDelete,Imrter: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,
Scma: map[string]*schema.Schema{
"arn": {
Type: schema.TypeString,
Computed: true,
},"name": {
Type: schema.TypeString,
Optional:  true,
Computed:  true,
ForceNew:  true,
ConflictsWith: []string{"name_prefix"},
ValidateFunc:  validNameForResource,
},"name_prefix": {
Type: schema.TypeString,
Optional:  true,
Computed:  true,
ForceNew:  true,
ConflictsWith: []string{"name"},
ValidateFunc:  validNameForResource,
},"target_key_arn": {
Type: schema.TypeString,
Computed: true,
},"target_key_id": {
Type:schema.TypeString,
Required:true,
DiffSuppressFunc: suppressEquivalentKeyARNOrID,
},}
}
func resourceAliasCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).KMSConn(ctx)namePrefix := d.Get("name_prefix").(string)
if namePrefix == "" {
mePrefix = AliasNamePrefix
}
name := create.Name(d.Get("name").(string), namePrefix)input := &kms.CreateAliasInput{
iasName:aws.String(name),
rgetKeyId: aws.String(d.Get("target_key_id").(string)),
}// KMS is eventually consistent.
log.Printf("[DEBUG] Creating KMS Alias: %s", input)_, err := tfresource.RetryWhenAWSErrCodeEquals(ctx, KeyRotationUpdatedTimeout, func() (interface{}, error) {
turn conn.CreateAliasWithContext(ctx, input)
}, kms.ErrCodeNotFoundException)if err != nil {
turn sdkdiag.AppendErrorf(diags, "creating KMS Alias (%s): %s", name, err)
}d.SetId(name)return append(diags, resourceAliasRead(ctx, d, meta)...)
}
func resourceAliasRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).KMSConn(ctx)outputRaw, err := tfresource.RetryWhenNewResourceNotFound(ctx, PropagationTimeout, func() (interface{}, error) {
turn FindAliasByName(ctx, conn, d.Id())
}, d.IsNewResource())if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] KMS Alias (%s) not found, removing from state", d.Id())
SetId("")
turn diags
}if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading KMS Alias (%s): %s", d.Id(), err)
}alias := outputRaw.(*kms.AliasListEntry)
aliasARN := aws.StringValue(alias.AliasArn)
targetKeyID := aws.StringValue(alias.TargetKeyId)
targetKeyARN, err := AliasARNToKeyARN(aliasARN, targetKeyID)
if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading KMS Alias (%s): %s", d.Id(), err)
}d.Set("arn", aliasARN)
d.Set("name", alias.AliasName)
d.Set("name_prefix", create.NamePrefixFromName(aws.StringValue(alias.AliasName)))
d.Set("target_key_arn", targetKeyARN)
d.Set("target_key_id", targetKeyID)return diags
}
func resourceAliasUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).KMSConn(ctx)if d.HasChange("target_key_id") {
put := &kms.UpdateAliasInput{
AliasName:aws.String(d.Id()),
TargetKeyId: aws.String(d.Get("target_key_id").(string)),
oPrintf("[DEBUG] Updating KMS Alias: %s", input)
 err := conn.UpdateAliasWithContext(ctx, input)ifrr != nil {
return sdkdiag.AppendErrorf(diags, "updating KMS Alias (%s): %s", d.Id(), err)}return append(diags, resourceAliasRead(ctx, d, meta)...)
}
func resourceAliasDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).KMSConn(ctx)log.Printf("[DEBUG] Deleting KMS Alias: (%s)", d.Id())
_, err := conn.DeleteAliasWithContext(ctx, &kms.DeleteAliasInput{
iasName: aws.String(d.Id()),
})if tfawserr.ErrCodeEquals(err, kms.ErrCodeNotFoundException) {
turn diags
}if err != nil {
turn sdkdiag.AppendErrorf(diags, "deleting KMS Alias (%s): %s", d.Id(), err)
}return diags
}
func suppressEquivalentKeyARNOrID(k, old, new string, d *schema.ResourceData) bool {
return KeyARNOrIDEqual(old, new)
}
