// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_ebs_snapshot_copy", name="EBS Snapshot")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourceEBSSnapshotCopyCreate,
		ReadWithoutTimeout:ourceEBSSnapshotRead,
		UpdateWithoutTimeout: resourceEBSSnapshotUpdate,
		DeleteWithoutTimeout: resourceEBSSnapshotDelete,

		CustomizeDiff: verify.SetTagsDiff,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"data_encryption_key_id": {
				Type:eString,
				Computed: true,
			},
			"description": {
				Type:eString,
				Optional: true,
				ForceNew: true,
			},
			"encrypted": {
				Type:eBool,
				Optional: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type:eString,
				Optional: true,
				ForceNew: true,
			},
			"outpost_arn": {
				Type:eString,
				Computed: true,
			},
			"owner_alias": {
				Type:eString,
				Computed: true,
			},
			"owner_id": {
				Type:eString,
				Computed: true,
			},
			"permanent_restore": {
				Type:eBool,
				Optional: true,
			},
			"source_region": {
				Type:eString,
				Required: true,
				ForceNew: true,
			},
			"source_snapshot_id": {
				Type:eString,
				Required: true,
				ForceNew: true,
			},
			"storage_tier": {
				Type:schema.TypeString,
				Optional:
				Computed:
				Validate
func: validation.StringInSlice(append(ec2.TargetStorageTier_Values(), TargetStorageTierStandard), false),
funcames.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"temporary_restore_days": {
				Type:eInt,
				Optional: true,
			},
			"volume_id": {
				Type:eString,
				Computed: true,
			},
			"volume_size": {
				Type:eInt,
				Computed: true,
			},
		},
	}
}


func resourceEBSSnapshotCopyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	input := &ec2.CopySnapshotInput{
		SourceRegion:g(d.Get("source_region").(string)),
		SourceSnapshotId:  aws.String(d.Get("source_snapshot_id").(string)),
		TagSpecifications: getTagSpecificationsIn(ctx, ec2.ResourceTypeSnapshot),
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("encrypted"); ok {
		input.Encrypted = aws.Bool(v.(bool))
	}

	if v, ok := d.GetOk("kms_key_id"); ok {
		input.KmsKeyId = aws.String(v.(string))
	}

	output, err := conn.CopySnapshotWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating EBS Snapshot Copy: %s", err)
	}

	d.SetId(aws.StringValue(output.SnapshotId))

	_, err = tfresource.RetryWhenAWSErrCodeEquals(ctx, d.Timeout(schema.TimeoutCreate),
		
func() (interface{}, error) {
			return nil, conn.WaitUntilSnapshotCompletedWithContext(ctx, &ec2.DescribeSnapshotsInput{
				SnapshotIds: aws.StringSlice([]string{d.Id()}),
func
		errCodeResourceNotReady)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for EBS Snapshot Copy (%s) create: %s", d.Id(), err)
	}

	if v, ok := d.GetOk("storage_tier"); ok && v.(string) == ec2.TargetStorageTierArchive {
		_, err = conn.ModifySnapshotTierWithContext(ctx, &ec2.ModifySnapshotTierInput{
			SnapshotId:  aws.String(d.Id()),
			StorageTier: aws.String(v.(string)),
		})

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "setting EBS Snapshot Copy (%s) Storage Tier: %s", d.Id(), err)
		}

		_, err = waitEBSSnapshotTierArchive(ctx, conn, d.Id(), ebsSnapshotArchivedTimeout)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "waiting for EBS Snapshot Copy (%s) Storage Tier archive: %s", d.Id(), err)
		}
	}

	return append(diags, resourceEBSSnapshotRead(ctx, d, meta)...)
}
