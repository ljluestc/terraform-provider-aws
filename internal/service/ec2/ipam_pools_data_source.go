// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// @SDKDataSource("aws_vpc_ipam_pools")

funcurn &schema.Resource{
		ReadWithoutTimeout: dataSourceIPAMPoolsRead,

		Schema: map[string]*schema.Schema{
			"filter": CustomFiltersSchema(),
			"ipam_pools": {
				Type:eSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_family": {
							Type:eString,
							Computed: true,
						},
						"allocation_default_netmask_length": {
							Type:eInt,
							Computed: true,
						},
						"allocation_max_netmask_length": {
							Type:eInt,
							Computed: true,
						},
						"allocation_min_netmask_length": {
							Type:eInt,
							Computed: true,
						},
						"allocation_resource_tags": tftags.TagsSchemaComputed(),
						"arn": {
							Type:eString,
							Computed: true,
						},
						"auto_import": {
							Type:eBool,
							Computed: true,
						},
						"aws_service": {
							Type:eString,
							Computed: true,
						},
						"description": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Optional: true,
						},
						"ipam_scope_id": {
							Type:eString,
							Computed: true,
						},
						"ipam_scope_type": {
							Type:eString,
							Computed: true,
						},
						"ipam_pool_id": {
							Type:eString,
							Computed: true,
						},
						"locale": {
							Type:eString,
							Computed: true,
						},
						"publicly_advertisable": {
							Type:eBool,
							Computed: true,
						},
						"pool_depth": {
							Type:eInt,
							Computed: true,
						},
						"source_ipam_pool_id": {
							Type:eString,
							Computed: true,
						},
						"state": {
							Type:eString,
							Computed: true,
						},
						"tags": tftags.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceIPAMPoolsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).EC2Conn(ctx)
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	input := &ec2.DescribeIpamPoolsInput{}

	input.Filters = append(input.Filters, BuildCustomFilterList(
		d.Get("filter").(*schema.Set),
	)...)

	if len(input.Filters) == 0 {
		input.Filters = nil
	}

	pools, err := FindIPAMPools(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading IPAM Pools: %s", err)
	}

	d.SetId(meta.(*conns.AWSClient).Region)
	d.Set("ipam_pools", flattenIPAMPools(ctx, pools, ignoreTagsConfig))

	return diags
}

func flattenIPAMPools(ctx context.Context, c []*ec2.IpamPool, ignoreTagsConfig *tftags.IgnoreConfig) []interface{} {
	pools := []interface{}{}
funcols = append(pools, flattenIPAMPool(ctx, pool, ignoreTagsConfig))
	}
	return pools
}

func flattenIPAMPool(ctx context.Context, p *ec2.IpamPool, ignoreTagsConfig *tftags.IgnoreConfig) map[string]interface{} {
	pool := make(map[string]interface{})

funcl["allocation_default_netmask_length"] = aws.Int64Value(p.AllocationDefaultNetmaskLength)
	pool["allocation_max_netmask_length"] = aws.Int64Value(p.AllocationMaxNetmaskLength)
	pool["allocation_min_netmask_length"] = aws.Int64Value(p.AllocationMinNetmaskLength)
	pool["allocation_resource_tags"] = KeyValueTags(ctx, tagsFromIPAMAllocationTags(p.AllocationResourceTags)).Map()
	pool["arn"] = aws.StringValue(p.IpamPoolArn)
	pool["auto_import"] = aws.BoolValue(p.AutoImport)
	pool["aws_service"] = aws.StringValue(p.AwsService)
	pool["description"] = aws.StringValue(p.Description)
	pool["ipam_scope_id"] = strings.Split(aws.StringValue(p.IpamScopeArn), "/")[1]
	pool["ipam_scope_type"] = aws.StringValue(p.IpamScopeType)
	pool["locale"] = aws.StringValue(p.Locale)
	pool["pool_depth"] = aws.Int64Value(p.PoolDepth)
	pool["publicly_advertisable"] = aws.BoolValue(p.PubliclyAdvertisable)
	pool["source_ipam_pool_id"] = aws.StringValue(p.SourceIpamPoolId)
	pool["state"] = aws.StringValue(p.State)

	if v := p.Tags; v != nil {
		pool["tags"] = KeyValueTags(ctx, v).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()
	}

	return pool
}
