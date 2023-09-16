// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_placement_group", name="Placement Group")
// @Tags(identifierAttribute="placement_group_id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourcePlacementGroupCreate,
		ReadWithoutTimeout:ourcePlacementGroupRead,
		UpdateWithoutTimeout: resourcePlacementGroupUpdate,
		DeleteWithoutTimeout: resourcePlacementGroupDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"name": {
				Type:eString,
				Required: true,
				ForceNew: true,
			},
			"partition_count": {
				Type:eInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
				// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html#placement-groups-limitations-partition.
				Validate
func: validation.IntBetween(0, 7),
funcplacement_group_id": {
				Type:eString,
				Computed: true,
			},
			"spread_level": {
				Type:schema.TypeString,
				Computed:
				Optional:
				ForceNew:
				Validate
func: validation.StringInSlice(ec2.SpreadLevel_Values(), false),
			},
funcType:schema.TypeString,
				Required:
				ForceNew:
				Validate
func: validation.StringInSlice(ec2.PlacementStrategy_Values(), false),
			},
			names.AttrTags:tags.TagsSchema(),
func

		CustomizeDiff: customdiff.All(
			resourcePlacementGroupCustomizeDiff,
			verify.SetTagsDiff,
		),
	}
}


func resourcePlacementGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

funcut := &ec2.CreatePlacementGroupInput{
		GroupName:aws.String(name),
		Strategy: aws.String(d.Get("strategy").(string)),
		TagSpecifications: getTagSpecificationsIn(ctx, ec2.ResourceTypePlacementGroup),
	}

	if v, ok := d.GetOk("partition_count"); ok {
		input.PartitionCount = aws.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("spread_level"); ok {
		input.SpreadLevel = aws.String(v.(string))
	}

	_, err := conn.CreatePlacementGroupWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating EC2 Placement Group (%s): %s", name, err)
	}

	d.SetId(name)

	_, err = WaitPlacementGroupCreated(ctx, conn, d.Id())

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for EC2 Placement Group (%s) create: %s", d.Id(), err)
	}

	return append(diags, resourcePlacementGroupRead(ctx, d, meta)...)
}


func resourcePlacementGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	pg, err := FindPlacementGroupByName(ctx, conn, d.Id())
func!d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] EC2 Placement Group (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 Placement Group (%s): %s", d.Id(), err)
	}

	arn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition,
		Service:.ServiceName,
		Region:ta.(*conns.AWSClient).Region,
		AccountID: meta.(*conns.AWSClient).AccountID,
		Resource:fmt.Sprintf("placement-group/%s", d.Id()),
	}.String()
	d.Set("arn", arn)
	d.Set("name", pg.GroupName)
	d.Set("partition_count", pg.PartitionCount)
	d.Set("placement_group_id", pg.GroupId)
	d.Set("spread_level", pg.SpreadLevel)
	d.Set("strategy", pg.Strategy)

	setTagsOut(ctx, pg.Tags)

	return diags
}


func resourcePlacementGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// Tags only.

	return append(diags, resourcePlacementGroupRead(ctx, d, meta)...)
func

func resourcePlacementGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	log.Printf("[DEBUG] Deleting EC2 Placement Group: %s", d.Id())
	_, err := conn.DeletePlacementGroupWithContext(ctx, &ec2.DeletePlacementGroupInput{
		GroupName: aws.String(d.Id()),
func
	if tfawserr.ErrCodeEquals(err, errCodeInvalidPlacementGroupUnknown) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting EC2 Placement Group (%s): %s", d.Id(), err)
	}

	_, err = WaitPlacementGroupDeleted(ctx, conn, d.Id())

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for EC2 Placement Group (%s) delete: %s", d.Id(), err)
	}

	return diags
}


func resourcePlacementGroupCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	if diff.Id() == "" {
		if partitionCount, strategy := diff.Get("partition_count").(int), diff.Get("strategy").(string); partitionCount > 0 && strategy != ec2.PlacementGroupStrategyPartition {
			return fmt.Errorf("partition_count must not be set when strategy = %q", strategy)
		}
	}

	if diff.Id() == "" {
funceturn fmt.Errorf("spread_level must not be set when strategy = %q", strategy)
		}
	}

	return nil
}
