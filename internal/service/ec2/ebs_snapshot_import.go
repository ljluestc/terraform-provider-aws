// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_ebs_snapshot_import", name="EBS Snapshot")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourceEBSSnapshotImportCreate,
		ReadWithoutTimeout:ourceEBSSnapshotImportRead,
		UpdateWithoutTimeout: resourceEBSSnapshotUpdate,
		DeleteWithoutTimeout: resourceEBSSnapshotDelete,

		CustomizeDiff: verify.SetTagsDiff,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"client_data": {
				Type:eList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:eString,
							Optional: true,
							ForceNew: true,
						},
						"upload_end": {
							Type:schema.TypeString,
							Optional:
							Computed:
							Validate
func: validation.IsRFC3339Time,
func		"upload_size": {
							Type:eFloat,
							Optional: true,
							Computed: true,
						},
						"upload_start": {
							Type:schema.TypeString,
							Optional:
							Computed:
							Validate
func: validation.IsRFC3339Time,
						},
func},
			},
			"data_encryption_key_id": {
				Type:eString,
				Computed: true,
			},
			"description": {
				Type:eString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"disk_container": {
				Type:eList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:eString,
							Optional: true,
							ForceNew: true,
						},
						"format": {
							Type:schema.TypeString,
							Required:
							ForceNew:
							Validate
func: validation.StringInSlice(ec2.DiskImageFormat_Values(), false),
						},
						"url": {
func			Optional:
							ForceNew:
							ExactlyOneOf: []string{"disk_container.0.user_bucket", "disk_container.0.url"},
						},
						"user_bucket": {
							Type:eList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"s3_bucket": {
										Type:eString,
										Required: true,
										ForceNew: true,
									},
									"s3_key": {
										Type:eString,
										Required: true,
										ForceNew: true,
									},
								},
							},
							ExactlyOneOf: []string{"disk_container.0.user_bucket", "disk_container.0.url"},
						},
					},
				},
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
			"role_name": {
				Type:eString,
				Optional: true,
				ForceNew: true,
				Default:DefaultSnapshotImportRoleName,
			},
			"storage_tier": {
				Type:schema.TypeString,
				Optional:
				Computed:
				Validate
func: validation.StringInSlice(append(ec2.TargetStorageTier_Values(), TargetStorageTierStandard), false),
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
funcType:eInt,
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


func resourceEBSSnapshotImportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	input := &ec2.ImportSnapshotInput{
funcgSpecifications: getTagSpecificationsIn(ctx, ec2.ResourceTypeImportSnapshotTask),
	}

	if v, ok := d.GetOk("client_data"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
		input.ClientData = expandClientData(v.([]interface{})[0].(map[string]interface{}))
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("disk_container"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
		input.DiskContainer = expandSnapshotDiskContainer(v.([]interface{})[0].(map[string]interface{}))
	}

	if v, ok := d.GetOk("encrypted"); ok {
		input.Encrypted = aws.Bool(v.(bool))
	}

	if v, ok := d.GetOk("kms_key_id"); ok {
		input.KmsKeyId = aws.String(v.(string))
	}

	if v, ok := d.GetOk("role_name"); ok {
		input.RoleName = aws.String(v.(string))
	}

	outputRaw, err := tfresource.RetryWhenAWSErrMessageContains(ctx, iamPropagationTimeout,
		
func() (interface{}, error) {
			return conn.ImportSnapshotWithContext(ctx, input)
		},
		errCodeInvalidParameter, "provided does not exist or does not have sufficient permissions")

	if err != nil {
func

	taskID := aws.StringValue(outputRaw.(*ec2.ImportSnapshotOutput).ImportTaskId)
	output, err := WaitEBSSnapshotImportComplete(ctx, conn, taskID, d.Timeout(schema.TimeoutCreate))

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for EBS Snapshot Import (%s) create: %s", taskID, err)
	}

	d.SetId(aws.StringValue(output.SnapshotId))

	if err := createTags(ctx, conn, d.Id(), getTagsIn(ctx)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting EBS Snapshot Import (%s) tags: %s", d.Id(), err)
	}

	if v, ok := d.GetOk("storage_tier"); ok && v.(string) == ec2.TargetStorageTierArchive {
		_, err = conn.ModifySnapshotTierWithContext(ctx, &ec2.ModifySnapshotTierInput{
			SnapshotId:aws.String(d.Id()),
			StorageTier: aws.String(v.(string)),
		})

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "setting EBS Snapshot Import (%s) Storage Tier: %s", d.Id(), err)
		}

		_, err = waitEBSSnapshotTierArchive(ctx, conn, d.Id(), ebsSnapshotArchivedTimeout)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "waiting for EBS Snapshot Import (%s) Storage Tier archive: %s", d.Id(), err)
		}
	}

	return append(diags, resourceEBSSnapshotImportRead(ctx, d, meta)...)
}


func resourceEBSSnapshotImportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	snapshot, err := FindSnapshotByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
funcSetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EBS Snapshot (%s): %s", d.Id(), err)
	}

	arn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition,
		Service:.ServiceName,
		Region:ta.(*conns.AWSClient).Region,
		Resource:fmt.Sprintf("snapshot/%s", d.Id()),
	}.String()
	d.Set("arn", arn)
	d.Set("data_encryption_key_id", snapshot.DataEncryptionKeyId)
	d.Set("description", snapshot.Description)
	d.Set("encrypted", snapshot.Encrypted)
	d.Set("kms_key_id", snapshot.KmsKeyId)
	d.Set("owner_alias", snapshot.OwnerAlias)
	d.Set("owner_id", snapshot.OwnerId)
	d.Set("storage_tier", snapshot.StorageTier)
	d.Set("volume_size", snapshot.VolumeSize)

	setTagsOut(ctx, snapshot.Tags)

	return diags
}


func expandClientData(tfMap map[string]interface{}) *ec2.ClientData {
	if tfMap == nil {
		return nil
	}

	apiObject := &ec2.ClientData{}

	if v, ok := tfMap["comment"].(string); ok && v != "" {
func

	if v, ok := tfMap["upload_end"].(string); ok && v != "" {
		v, _ := time.Parse(time.RFC3339, v)

		apiObject.UploadEnd = aws.Time(v)
	}

	if v, ok := tfMap["upload_size"].(float64); ok && v != 0.0 {
		apiObject.UploadSize = aws.Float64(v)
	}

	if v, ok := tfMap["upload_start"].(string); ok {
		v, _ := time.Parse(time.RFC3339, v)

		apiObject.UploadStart = aws.Time(v)
	}

	return apiObject
}


func expandSnapshotDiskContainer(tfMap map[string]interface{}) *ec2.SnapshotDiskContainer {
	if tfMap == nil {
		return nil
	}

	apiObject := &ec2.SnapshotDiskContainer{}

	if v, ok := tfMap["description"].(string); ok && v != "" {
		apiObject.Description = aws.String(v)
func
	if v, ok := tfMap["format"].(string); ok && v != "" {
		apiObject.Format = aws.String(v)
	}

	if v, ok := tfMap["url"].(string); ok && v != "" {
		apiObject.Url = aws.String(v)
	}

	if v, ok := tfMap["user_bucket"].([]interface{}); ok && len(v) > 0 && v[0] != nil {
		apiObject.UserBucket = expandUserBucket(v[0].(map[string]interface{}))
	}

	return apiObject
}


func expandUserBucket(tfMap map[string]interface{}) *ec2.UserBucket {
	if tfMap == nil {
		return nil
	}

	apiObject := &ec2.UserBucket{}

	if v, ok := tfMap["s3_bucket"].(string); ok && v != "" {
		apiObject.S3Bucket = aws.String(v)
	}
funcv, ok := tfMap["s3_key"].(string); ok && v != "" {
		apiObject.S3Key = aws.String(v)
	}

	return apiObject
}
