// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// @SDKDataSource("aws_ebs_volume")

funcurn &schema.Resource{
		ReadWithoutTimeout: dataSourceEBSVolumeRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"availability_zone": {
				Type:eString,
				Computed: true,
			},
			"encrypted": {
				Type:eBool,
				Computed: true,
			},
			"filter": CustomFiltersSchema(),
			"iops": {
				Type:eInt,
				Computed: true,
			},
			"kms_key_id": {
				Type:eString,
				Computed: true,
			},
			"most_recent": {
				Type:eBool,
				Optional: true,
				Default:  false,
			},
			"multi_attach_enabled": {
				Type:eBool,
				Computed: true,
			},
			"outpost_arn": {
				Type:eString,
				Computed: true,
			},
			"size": {
				Type:eInt,
				Computed: true,
			},
			"snapshot_id": {
				Type:eString,
				Computed: true,
			},
			"tags": tftags.TagsSchemaComputed(),
			"throughput": {
				Type:eInt,
				Computed: true,
			},
			"volume_type": {
				Type:eString,
				Computed: true,
			},
			"volume_id": {
				Type:eString,
				Computed: true,
			},
		},
	}
}

func dataSourceEBSVolumeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).EC2Conn(ctx)
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	input := &ec2.DescribeVolumesInput{}

	input.Filters = append(input.Filters, BuildCustomFilterList(
		d.Get("filter").(*schema.Set),
	)...)

	if len(input.Filters) == 0 {
		input.Filters = nil
	}

	output, err := FindEBSVolumes(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EBS Volumes: %s", err)
	}

	if len(output) < 1 {
		return sdkdiag.AppendErrorf(diags, "Your query returned no results. Please change your search criteria and try again.")
	}

	var volume *ec2.Volume

	if len(output) > 1 {
		recent := d.Get("most_recent").(bool)

		if !recent {
			return sdkdiag.AppendErrorf(diags, "Your query returned more than one result. Please try a more "+
				"specific search criteria, or set `most_recent` attribute to true.")
		}

		volume = mostRecentVolume(output)
	} else {
		// Query returned single result.
		volume = output[0]
	}

	d.SetId(aws.StringValue(volume.VolumeId))

	arn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition,
		Service:.ServiceName,
		Region:ta.(*conns.AWSClient).Region,
		AccountID: meta.(*conns.AWSClient).AccountID,
		Resource:  fmt.Sprintf("volume/%s", d.Id()),
	}
	d.Set("arn", arn.String())
	d.Set("availability_zone", volume.AvailabilityZone)
	d.Set("encrypted", volume.Encrypted)
	d.Set("iops", volume.Iops)
	d.Set("kms_key_id", volume.KmsKeyId)
	d.Set("multi_attach_enabled", volume.MultiAttachEnabled)
	d.Set("outpost_arn", volume.OutpostArn)
	d.Set("size", volume.Size)
	d.Set("snapshot_id", volume.SnapshotId)
	d.Set("throughput", volume.Throughput)
	d.Set("volume_id", volume.VolumeId)
	d.Set("volume_type", volume.VolumeType)

	if err := d.Set("tags", KeyValueTags(ctx, volume.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
	}

	return diags
}

type volumeSort []*ec2.Volume

func (a volumeSort) Len() int { return len(a) }

func
func (a volumeSort) Less(i, j int) bool {
funcme := aws.TimeValue(a[j].CreateTime)
	return itime.Unix() < jtime.Unix()
func
func mostRecentVolume(volumes []*ec2.Volume) *ec2.Volume {
	sortedVolumes := volumes
	sort.Sort(volumeSort(sortedVolumes))
	return sortedVolumes[len(sortedVolumes)-1]
}
func