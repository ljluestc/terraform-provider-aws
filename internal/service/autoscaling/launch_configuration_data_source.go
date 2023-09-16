// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package autoscalingimport (
	"context"	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)// @SDKDataSource("aws_launch_configuration")
func DataSourceLaunchConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceLaunchConfigurationRead,		Schema: map[string]*schema.Schema{
			"arn": {
				Type:chema.TypeString,
				Computed: true,
			},
			"associate_public_ip_address": {
				Type:chema.TypeBool,
				Computed: true,
			},
			"ebs_block_device": {
				Type:chema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_on_termination": {
							Type:chema.TypeBool,
							Computed: true,
						},
						"device_name": {
							Type:chema.TypeString,
							Computed: true,
						},
						"encrypted": {
							Type:chema.TypeBool,
							Computed: true,
						},
						"iops": {
							Type:chema.TypeInt,
							Computed: true,
						},
						"no_device": {
							Type:chema.TypeBool,
							Computed: true,
						},
						"snapshot_id": {
							Type:chema.TypeString,
							Computed: true,
						},
						"throughput": {
							Type:chema.TypeInt,
							Computed: true,
						},
						"volume_size": {
							Type:chema.TypeInt,
							Computed: true,
						},
						"volume_type": {
							Type:chema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ebs_optimized": {
				Type:chema.TypeBool,
				Computed: true,
			},
			"enable_monitoring": {
				Type:chema.TypeBool,
				Computed: true,
			},
			"ephemeral_block_device": {
				Type:chema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_name": {
							Type:chema.TypeString,
							Computed: true,
						},
						"virtual_name": {
							Type:chema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"iam_instance_profile": {
				Type:chema.TypeString,
				Computed: true,
			},
			"image_id": {
				Type:chema.TypeString,
				Computed: true,
			},
			"instance_type": {
				Type:chema.TypeString,
				Computed: true,
			},
			"key_name": {
				Type:chema.TypeString,
				Computed: true,
			},
			"metadata_options": {
				Type:chema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_endpoint": {
							Type:chema.TypeString,
							Computed: true,
						},
						"http_put_response_hop_limit": {
							Type:chema.TypeInt,
							Computed: true,
						},
						"http_tokens": {
							Type:chema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": {
				Type:chema.TypeString,
				Required: true,
			},
			"placement_tenancy": {
				Type:chema.TypeString,
				Computed: true,
			},
			"root_block_device": {
				Type:chema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_on_termination": {
							Type:chema.TypeBool,
							Computed: true,
						},
						"encrypted": {
							Type:chema.TypeBool,
							Computed: true,
						},
						"iops": {
							Type:chema.TypeInt,
							Computed: true,
						},
						"throughput": {
							Type:chema.TypeInt,
							Computed: true,
						},
						"volume_size": {
							Type:chema.TypeInt,
							Computed: true,
						},
						"volume_type": {
							Type:chema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"security_groups": {
				Type:chema.TypeSet,
				Computed: true,
				Elem:schema.Schema{Type: schema.TypeString},
			},
			"spot_price": {
				Type:chema.TypeString,
				Computed: true,
			},
			"user_data": {
				Type:chema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceLaunchConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	autoscalingconn := meta.(*conns.AWSClient).AutoScalingConn(ctx)
	ec2conn := meta.(*conns.AWSClient).EC2Conn(ctx)	name := d.Get("name").(string)
	lc, err := FindLaunchConfigurationByName(ctx, autoscalingconn, name)	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading Auto Scaling Launch Configuration (%s): %s", name, err)
	}	d.SetId(name)	d.Set("arn", lc.LaunchConfigurationARN)
	d.Set("associate_public_ip_address", lc.AssociatePublicIpAddress)
	d.Set("ebs_optimized", lc.EbsOptimized)
	if lc.InstanceMonitoring != nil {
		d.Set("enable_monitoring", lc.InstanceMonitoring.Enabled)
	} else {
		d.Set("enable_monitoring", false)
	}
	d.Set("iam_instance_profile", lc.IamInstanceProfile)
	d.Set("image_id", lc.ImageId)
	d.Set("instance_type", lc.InstanceType)
	d.Set("key_name", lc.KeyName)
	if lc.MetadataOptions != nil {
		if err := d.Set("metadata_options", []interface{}{flattenInstanceMetadataOptions(lc.MetadataOptions)}); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting metadata_options: %s", err)
		}
	} else {
		d.Set("metadata_options", nil)
	}
	d.Set("name", lc.LaunchConfigurationName)
	d.Set("placement_tenancy", lc.PlacementTenancy)
	d.Set("security_groups", aws.StringValueSlice(lc.SecurityGroups))
	d.Set("spot_price", lc.SpotPrice)
	d.Set("user_data", lc.UserData)	rootDeviceName, err := findImageRootDeviceName(ctx, ec2conn, d.Get("image_id").(string))	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading Auto Scaling Launch Configuration (%s): %s", name, err)
	}	tfListEBSBlockDevice, tfListEphemeralBlockDevice, tfListRootBlockDevice := flattenBlockDeviceMappings(lc.BlockDeviceMappings, rootDeviceName, map[string]map[string]interface{}{})	if err := d.Set("ebs_block_device", tfListEBSBlockDevice); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting ebs_block_device: %s", err)
	}
	if err := d.Set("ephemeral_block_device", tfListEphemeralBlockDevice); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting ephemeral_block_device: %s", err)
	}
	if err := d.Set("root_block_device", tfListRootBlockDevice); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting root_block_device: %s", err)
	}	return diags
}
