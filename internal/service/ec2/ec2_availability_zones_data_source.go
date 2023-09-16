// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)

// @SDKDataSource("aws_availability_zones")

funcurn &schema.Resource{
		ReadWithoutTimeout: dataSourceAvailabilityZonesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"all_availability_zones": {
				Type:eBool,
				Optional: true,
			},
			"exclude_names": {
				Type:eSet,
				Optional: true,
				Elem:hema{Type: schema.TypeString},
			},
			"exclude_zone_ids": {
				Type:eSet,
				Optional: true,
				Elem:hema{Type: schema.TypeString},
			},
			"filter": CustomFiltersSchema(),
			"group_names": {
				Type:eSet,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"names": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"state": {
				Type:schema.TypeString,
				Optional:
				Validate
func: validation.StringInSlice(ec2.AvailabilityZoneState_Values(), false),
funczone_ids": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
		},
	}
}


func dataSourceAvailabilityZonesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	log.Printf("[DEBUG] Reading Availability Zones.")

	request := &ec2.DescribeAvailabilityZonesInput{}

	if v, ok := d.GetOk("all_availability_zones"); ok {
		request.AllAvailabilityZones = aws.Bool(v.(bool))
	}

	if v, ok := d.GetOk("state"); ok {
		request.Filters = []*ec2.Filter{
			{
				Name:.String("state"),
				Values: []*string{aws.String(v.(string))},
			},
		}
	}

	if filters, filtersOk := d.GetOk("filter"); filtersOk {
		request.Filters = append(request.Filters, BuildCustomFilterList(
			filters.(*schema.Set),
		)...)
	}

	if len(request.Filters) == 0 {
		// Don't send an empty filters list; the EC2 API won't accept it.
		request.Filters = nil
	}

	log.Printf("[DEBUG] Reading Availability Zones: %s", request)
	resp, err := conn.DescribeAvailabilityZonesWithContext(ctx, request)
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "fetching Availability Zones: %s", err)
	}

	sort.Slice(resp.AvailabilityZones, 
func(i, j int) bool {
		return aws.StringValue(resp.AvailabilityZones[i].ZoneName) < aws.StringValue(resp.AvailabilityZones[j].ZoneName)
	})
funcludeNames := d.Get("exclude_names").(*schema.Set)
	excludeZoneIDs := d.Get("exclude_zone_ids").(*schema.Set)

	groupNames := schema.NewSet(schema.HashString, nil)
	names := []string{}
	zoneIds := []string{}
	for _, v := range resp.AvailabilityZones {
		groupName := aws.StringValue(v.GroupName)
		name := aws.StringValue(v.ZoneName)
		zoneID := aws.StringValue(v.ZoneId)

		if excludeNames.Contains(name) {
			continue
		}

		if excludeZoneIDs.Contains(zoneID) {
			continue
		}

		if !groupNames.Contains(groupName) {
			groupNames.Add(groupName)
		}

		names = append(names, name)
		zoneIds = append(zoneIds, zoneID)
	}

	d.SetId(meta.(*conns.AWSClient).Region)

	if err := d.Set("group_names", groupNames); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting group_names: %s", err)
	}
	if err := d.Set("names", names); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting Availability Zone names: %s", err)
	}
	if err := d.Set("zone_ids", zoneIds); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting Availability Zone IDs: %s", err)
	}

	return diags
}
