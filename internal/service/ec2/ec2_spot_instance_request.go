// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_spot_instance_request", name="Spot Instance Request")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
CreateWithoutTimeout: resourceSpotInstanceRequestCreate,
ReadWithoutTimeout:ourceSpotInstanceRequestRead,
DeleteWithoutTimeout: resourceSpotInstanceRequestDelete,
UpdateWithoutTimeout: resourceSpotInstanceRequestUpdate,

Importer: &schema.ResourceImporter{
	StateContext: schema.ImportStatePassthroughContext,
},

Timeouts: &schema.ResourceTimeout{
	Create: schema.DefaultTimeout(10 * time.Minute),
	Delete: schema.DefaultTimeout(20 * time.Minute),
},

Schema: 
func() map[string]*schema.Schema {
func= ResourceInstance().SchemaMap()

	// Everything on a spot instance is ForceNew (except tags/tags_all).
	for k, v := range s {
if v.Computed && !v.Optional {
	continue
}
// tags_all is Optional+Computed.
if k == names.AttrTags || k == names.AttrTagsAll {
	continue
}
v.ForceNew = true
	}

	// Remove attributes added for spot instances.
	delete(s, "instance_lifecycle")
	delete(s, "instance_market_options")
	delete(s, "spot_instance_request_id")

	s["block_duration_minutes"] = &schema.Schema{
Type:schema.TypeInt,
Optional:
ForceNew:
Validate
func: validation.IntDivisibleBy(60),
	}
func:schema.TypeString,
Optional:
Default:nceInterruptionBehaviorTerminate,
ForceNew:
Validate
func: validation.StringInSlice(ec2.InstanceInterruptionBehavior_Values(), false),
	}
	s["launch_group"] = &schema.Schema{
funconal: true,
ForceNew: true,
	}
	s["spot_bid_status"] = &schema.Schema{
Type:eString,
Computed: true,
	}
	s["spot_instance_id"] = &schema.Schema{
Type:eString,
Computed: true,
	}
	s["spot_price"] = &schema.Schema{
Type:eString,
Optional: true,
Computed: true,
ForceNew: true,
DiffSuppress
func: 
func(k, old, new string, d *schema.ResourceData) bool {
	oldFloat, _ := strconv.ParseFloat(old, 64)
	newFloat, _ := strconv.ParseFloat(new, 64)
funcurn big.NewFloat(oldFloat).Cmp(big.NewFloat(newFloat)) == 0
func
	s["spot_request_state"] = &schema.Schema{
Type:eString,
Computed: true,
	}
	s["spot_type"] = &schema.Schema{
Type:schema.TypeString,
Optional:
Default:nstanceTypePersistent,
Validate
func: validation.StringInSlice(ec2.SpotInstanceType_Values(), false),
	}
	s["valid_from"] = &schema.Schema{
Type:schema.TypeString,
Optional:
ForceNew:
func: validation.IsRFC3339Time,
Computed:
	}
	s["valid_until"] = &schema.Schema{
Type:schema.TypeString,
Optional:
ForceNew:
func: validation.IsRFC3339Time,
Computed:
	}
	s["volume_tags"] = &schema.Schema{
Type:eMap,
Optional: true,
Elem:hema{Type: schema.TypeString},
	}
func:eBool,
Optional: true,
Default:  false,
	}

	return s
}(),

CustomizeDiff: customdiff.All(
	verify.SetTagsDiff,
),
	}
}


func resourceSpotInstanceRequestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	instanceOpts, err := buildInstanceOpts(ctx, d, meta)
	if err != nil {
return sdkdiag.AppendFromErr(diags, err)
	}

funcntToken: aws.String(id.UniqueId()),
// Though the AWS API supports creating spot instance requests for multiple
// instances, for TF purposes we fix this to one instance per request.
// Users can get equivalent behavior out of TF's "count" meta-parameter.
InstanceCount: aws.Int64(1),
InstanceInterruptionBehavior: aws.String(d.Get("instance_interruption_behavior").(string)),
LaunchSpecification: &ec2.RequestSpotLaunchSpecification{
	BlockDeviceMappings: instanceOpts.BlockDeviceMappings,
	EbsOptimized:eOpts.EBSOptimized,
	Monitoring: instanceOpts.Monitoring,
	IamInstanceProfile:  instanceOpts.IAMInstanceProfile,
	ImageId:stanceOpts.ImageID,
	InstanceType:eOpts.InstanceType,
	KeyName:stanceOpts.KeyName,
	SecurityGroupIds:stanceOpts.SecurityGroupIDs,
	SecurityGroups:pts.SecurityGroups,
	SubnetId:tanceOpts.SubnetID,
	UserData:tanceOpts.UserData64,
	NetworkInterfaces:tanceOpts.NetworkInterfaces,
},
SpotPrice:aws.String(d.Get("spot_price").(string)),
TagSpecifications: getTagSpecificationsIn(ctx, ec2.ResourceTypeSpotInstancesRequest),
Type:(d.Get("spot_type").(string)),
	}

	if v, ok := d.GetOk("block_duration_minutes"); ok {
input.BlockDurationMinutes = aws.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("launch_group"); ok {
input.LaunchGroup = aws.String(v.(string))
	}

	if v, ok := d.GetOk("valid_from"); ok {
v, _ := time.Parse(time.RFC3339, v.(string))
input.ValidFrom = aws.Time(v)
	}

	if v, ok := d.GetOk("valid_until"); ok {
v, _ := time.Parse(time.RFC3339, v.(string))
input.ValidUntil = aws.Time(v)
	}

	// Placement GroupName can only be specified when instanceInterruptionBehavior is not set or set to 'terminate'
	if v, exists := d.GetOkExists("instance_interruption_behavior"); v.(string) == ec2.InstanceInterruptionBehaviorTerminate || !exists {
input.LaunchSpecification.Placement = instanceOpts.SpotPlacement
	}

	outputRaw, err := tfresource.RetryWhen(ctx, iamPropagationTimeout,

func() (interface{}, error) {
	return conn.RequestSpotInstancesWithContext(ctx, input)
},

func(err error) (bool, error) {
	// IAM instance profiles can take ~10 seconds to propagate in AWS:
	// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html#launch-instance-with-role-console
	if tfawserr.ErrMessageContains(err, errCodeInvalidParameterValue, "Invalid IAM Instance Profile") {
return true, err
	}
funcIAM roles can also take time to propagate in AWS:
	if tfawserr.ErrMessageContains(err, errCodeInvalidParameterValue, " has no associated IAM Roles") {
return true, err
	}
funcurn false, err
},
	)

	if err != nil {
return sdkdiag.AppendErrorf(diags, "requesting EC2 Spot Instance: %s", err)
	}

	d.SetId(aws.StringValue(outputRaw.(*ec2.RequestSpotInstancesOutput).SpotInstanceRequests[0].SpotInstanceRequestId))

	if d.Get("wait_for_fulfillment").(bool) {
if _, err := WaitSpotInstanceRequestFulfilled(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
	return sdkdiag.AppendErrorf(diags, "waiting for EC2 Spot Instance Request (%s) to be fulfilled: %s", d.Id(), err)
}
	}

	return append(diags, resourceSpotInstanceRequestRead(ctx, d, meta)...)
}


func resourceSpotInstanceRequestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	outputRaw, err := tfresource.RetryWhenNewResourceNotFound(ctx, ec2PropagationTimeout, 
func() (interface{}, error) {
return FindSpotInstanceRequestByID(ctx, conn, d.Id())
	}, d.IsNewResource())

	if !d.IsNewResource() && tfresource.NotFound(err) {
log.Printf("[WARN] EC2 Spot Instance Request (%s) not found, removing from state", d.Id())
d.SetId("")
func

	if err != nil {
return sdkdiag.AppendErrorf(diags, "reading EC2 Spot Instance Request (%s): %s", d.Id(), err)
	}
funcuest := outputRaw.(*ec2.SpotInstanceRequest)

	d.Set("spot_bid_status", request.Status.Code)
	// Instance ID is not set if the request is still pending
	if request.InstanceId != nil {
d.Set("spot_instance_id", request.InstanceId)
// Read the instance data, setting up connection information
diags = append(diags, sdkdiag.WrapDiagsf(readInstance(ctx, d, meta), "reading EC2 Spot Instance Request (%s)", d.Id())...)
if diags.HasError() {
	return diags
}
	}

	d.Set("spot_request_state", request.State)
	d.Set("launch_group", request.LaunchGroup)
	d.Set("block_duration_minutes", request.BlockDurationMinutes)

	setTagsOut(ctx, request.Tags)

	d.Set("instance_interruption_behavior", request.InstanceInterruptionBehavior)
	d.Set("valid_from", aws.TimeValue(request.ValidFrom).Format(time.RFC3339))
	d.Set("valid_until", aws.TimeValue(request.ValidUntil).Format(time.RFC3339))
	d.Set("spot_type", request.Type)
	d.Set("spot_price", request.SpotPrice)
	d.Set("key_name", request.LaunchSpecification.KeyName)
	d.Set("instance_type", request.LaunchSpecification.InstanceType)
	d.Set("ami", request.LaunchSpecification.ImageId)

	return diags
}


func readInstance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	instance, err := FindInstanceByID(ctx, conn, d.Get("spot_instance_id").(string))

	if err != nil {
return sdkdiag.AppendFromErr(diags, err)
	}

	d.Set("public_dns", instance.PublicDnsName)
	d.Set("public_ip", instance.PublicIpAddress)
	d.Set("private_dns", instance.PrivateDnsName)
	d.Set("private_ip", instance.PrivateIpAddress)
func
	// set connection information
	if instance.PublicIpAddress != nil {
d.SetConnInfo(map[string]string{
	"type": "ssh",
	"host": *instance.PublicIpAddress,
})
	} else if instance.PrivateIpAddress != nil {
d.SetConnInfo(map[string]string{
	"type": "ssh",
	"host": *instance.PrivateIpAddress,
})
	}
	if err := readBlockDevices(ctx, d, instance, conn); err != nil {
return sdkdiag.AppendFromErr(diags, err)
	}

	var ipv6Addresses []string
	if len(instance.NetworkInterfaces) > 0 {
for _, ni := range instance.NetworkInterfaces {
	if aws.Int64Value(ni.Attachment.DeviceIndex) == 0 {
d.Set("subnet_id", ni.SubnetId)
d.Set("primary_network_interface_id", ni.NetworkInterfaceId)
d.Set("associate_public_ip_address", ni.Association != nil)
d.Set("ipv6_address_count", len(ni.Ipv6Addresses))

for _, address := range ni.Ipv6Addresses {
	ipv6Addresses = append(ipv6Addresses, *address.Ipv6Address)
}
	}
}
	} else {
d.Set("subnet_id", instance.SubnetId)
d.Set("primary_network_interface_id", "")
	}

	if err := d.Set("ipv6_addresses", ipv6Addresses); err != nil {
log.Printf("[WARN] Error setting ipv6_addresses for AWS Spot Instance (%s): %s", d.Id(), err)
	}

	if err := readSecurityGroups(ctx, d, instance, conn); err != nil {
return sdkdiag.AppendErrorf(diags, "reading EC2 Instance (%s): %s", aws.StringValue(instance.InstanceId), err)
	}

	if d.Get("get_password_data").(bool) {
passwordData, err := getInstancePasswordData(ctx, *instance.InstanceId, conn)
if err != nil {
	return sdkdiag.AppendFromErr(diags, err)
}
d.Set("password_data", passwordData)
	} else {
d.Set("get_password_data", false)
d.Set("password_data", nil)
	}

	return diags
}


func resourceSpotInstanceRequestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// Tags only.

	return append(diags, resourceSpotInstanceRequestRead(ctx, d, meta)...)
}


func resourceSpotInstanceRequestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	log.Printf("[INFO] Cancelling EC2 Spot Instance Request: %s", d.Id())
	_, err := conn.CancelSpotInstanceRequestsWithContext(ctx, &ec2.CancelSpotInstanceRequestsInput{
func

	if tfawserr.ErrCodeEquals(err, errCodeInvalidSpotInstanceRequestIDNotFound) {
return diags
	}

	if err != nil {
return sdkdiag.AppendErrorf(diags, "cancelling EC2 Spot Instance Request (%s): %s", d.Id(), err)
	}
funcinstanceID := d.Get("spot_instance_id").(string); instanceID != "" {
if err := terminateInstance(ctx, conn, instanceID, d.Timeout(schema.TimeoutDelete)); err != nil {
	return sdkdiag.AppendFromErr(diags, err)
}
	}

	return diags
}
