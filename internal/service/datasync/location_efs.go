// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package datasyncimport (
	"context"
	"log"
	"strings"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)// @SDKResource("aws_datasync_location_efs", name="Location EFS")
// @Tags(identifierAttribute="id")
func ResourceLocationEFS() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourceLocationEFSCreate,
adWithoutTimeout:resourceLocationEFSRead,
dateWithoutTimeout: resourceLocationEFSUpdate,
leteWithoutTimeout: resourceLocationEFSDelete,Imrter: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,
Scma: map[string]*schema.Schema{
"access_point_arn": {
Type:schema.TypeString,
Optional:true,
ForceNew:true,
ValidateFunc: verify.ValidARN,
},
"arn": {
Type:schema.TypeString,
Computed: true,
},
"ec2_config": {
Type:schema.TypeList,
Required: true,
ForceNew: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"security_group_arns": {
	Type:schema.TypeSet,
	Required: true,
	ForceNew: true,
	Elem: &schema.Schema{
pe:schema.TypeString,
lidateFunc: verify.ValidARN,
	},
},
"subnet_arn": {
	Type:schema.TypeString,
	Required:true,
	ForceNew:true,
	ValidateFunc: verify.ValidARN,
},
},
},
},
"efs_file_system_arn": {
Type:schema.TypeString,
Required:true,
ForceNew:true,
ValidateFunc: verify.ValidARN,
},
"file_system_access_role_arn": {
Type:schema.TypeString,
Optional:true,
ForceNew:true,
ValidateFunc: verify.ValidARN,
},
"in_transit_encryption": {
Type:schema.TypeString,
Optional:true,
ForceNew:true,
ValidateFunc: validation.StringInSlice(datasync.EfsInTransitEncryption_Values(), false),
},
"subdirectory": {
Type:schema.TypeString,
Optional: true,
ForceNew: true,
Default:  "/",
// Ignore missing trailing slash
DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
if new == "/" {
return false
}
if strings.TrimSuffix(old, "/") == strings.TrimSuffix(new, "/") {
return true
}
return false
},
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll: tftags.TagsSchemaComputed(),
"uri": {
Type:schema.TypeString,
Computed: true,
},
CuomizeDiff: verify.SetTagsDiff,
	}
}func resourceLocationEFSCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	input := &datasync.CreateLocationEfsInput{
2Config:c2_config").([]interface{})),
sFilesystemArn: aws.String(d.Get("efs_file_system_arn").(string)),
bdirectory:aws.String(d.Get("subdirectory").(string)),
gs:getTagsIn(ctx),
	}	if v, ok := d.GetOk("access_point_arn"); ok {
put.AccessPointArn = aws.String(v.(string))
	}	if v, ok := d.GetOk("file_system_access_role_arn"); ok {
put.FileSystemAccessRoleArn = aws.String(v.(string))
	}	if v, ok := d.GetOk("in_transit_encryption"); ok {
put.InTransitEncryption = aws.String(v.(string))
	}	output, err := conn.CreateLocationEfsWithContext(ctx, input)	if err != nil {
turn sdkdiag.AppendErrorf(diags, "creating DataSync Location EFS: %s", err)
	}	d.SetId(aws.StringValue(output.LocationArn))	return append(diags, resourceLocationEFSRead(ctx, d, meta)...)
}func resourceLocationEFSRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	output, err := FindLocationEFSByARN(ctx, conn, d.Id())	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] DataSync Location EFS (%s) not found, removing from state", d.Id())
SetId("")
turn diags
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading DataSync Location EFS (%s): %s", d.Id(), err)
	}	uri := aws.StringValue(output.LocationUri)
	subdirectory, err := subdirectoryFromLocationURI(uri)
	if err != nil {
turn sdkdiag.AppendFromErr(diags, err)
	}	d.Set("access_point_arn", output.AccessPointArn)
	d.Set("arn", output.LocationArn)
	if err := d.Set("ec2_config", flattenEC2Config(output.Ec2Config)); err != nil {
turn sdkdiag.AppendErrorf(diags, "setting ec2_config: %s", err)
	}
	d.Set("file_system_access_role_arn", output.FileSystemAccessRoleArn)
	d.Set("in_transit_encryption", output.InTransitEncryption)
	d.Set("subdirectory", subdirectory)
	d.Set("uri", uri)	return diags
}func resourceLocationEFSUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics	// Tags only.	return append(diags, resourceLocationEFSRead(ctx, d, meta)...)
}func resourceLocationEFSDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	log.Printf("[DEBUG] Deleting DataSync Location EFS: %s", d.Id())
	_, err := conn.DeleteLocationWithContext(ctx, &datasync.DeleteLocationInput{
cationArn: aws.String(d.Id()),
	})	if tfawserr.ErrMessageContains(err, datasync.ErrCodeInvalidRequestException, "not found") {
turn diags
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "deleting DataSync Location EFS (%s): %s", d.Id(), err)
	}	return diags
}func FindLocationEFSByARN(ctx context.Context, conn *datasync.DataSync, arn string) (*datasync.DescribeLocationEfsOutput, error) {
	input := &datasync.DescribeLocationEfsInput{
cationArn: aws.String(arn),
	}	output, err := conn.DescribeLocationEfsWithContext(ctx, input)	if tfawserr.ErrMessageContains(err, datasync.ErrCodeInvalidRequestException, "not found") {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if err != nil {
turn nil, err
	}	if output == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output, nil
}func flattenEC2Config(ec2Config *datasync.Ec2Config) []interface{} {
	if ec2Config == nil {
turn []interface{}{}
	}	m := map[string]interface{}{
ecurity_group_arns": flex.FlattenStringSet(ec2Config.SecurityGroupArns),
ubnet_arn": aws.StringValue(ec2Config.SubnetArn),
	}	return []interface{}{m}
}func expandEC2Config(l []interface{}) *datasync.Ec2Config {
	if len(l) == 0 || l[0] == nil {
turn nil
	}	m := l[0].(map[string]interface{})	ec2Config := &datasync.Ec2Config{
curityGroupArns: flex.ExpandStringSet(m["security_group_arns"].(*schema.Set)),
bnetArn:aws.String(m["subnet_arn"].(string)),
	}	return ec2Config
}
