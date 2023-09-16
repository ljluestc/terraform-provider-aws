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
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKDataSource("aws_subnet")

funcurn &schema.Resource{
		ReadWithoutTimeout: dataSourceSubnetRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"assign_ipv6_address_on_creation": {
				Type:eBool,
				Computed: true,
			},
			"availability_zone": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			"availability_zone_id": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			"available_ip_address_count": {
				Type:eInt,
				Computed: true,
			},
			"cidr_block": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			"customer_owned_ipv4_pool": {
				Type:eString,
				Computed: true,
			},
			"default_for_az": {
				Type:eBool,
				Optional: true,
				Computed: true,
			},
			"enable_dns64": {
				Type:eBool,
				Computed: true,
			},
			"enable_lni_at_device_index": {
				Type:eInt,
				Computed: true,
			},
			"enable_resource_name_dns_aaaa_record_on_launch": {
				Type:eBool,
				Computed: true,
			},
			"enable_resource_name_dns_a_record_on_launch": {
				Type:eBool,
				Computed: true,
			},
			"filter": CustomFiltersSchema(),
			"id": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			"ipv6_cidr_block": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			"ipv6_cidr_block_association_id": {
				Type:eString,
				Computed: true,
			},
			"ipv6_native": {
				Type:eBool,
				Computed: true,
			},
			"map_customer_owned_ip_on_launch": {
				Type:eBool,
				Computed: true,
			},
			"map_public_ip_on_launch": {
				Type:eBool,
				Computed: true,
			},
			"outpost_arn": {
				Type:eString,
				Computed: true,
			},
			"owner_id": {
				Type:eString,
				Computed: true,
			},
			"private_dns_hostname_type_on_launch": {
				Type:eString,
				Computed: true,
			},
			"state": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			"tags": tftags.TagsSchemaComputed(),
			"vpc_id": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceSubnetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).EC2Conn(ctx)
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	input := &ec2.DescribeSubnetsInput{}

	if id, ok := d.GetOk("id"); ok {
		input.SubnetIds = []*string{aws.String(id.(string))}
	}

	// We specify default_for_az as boolean, but EC2 filters want
	// it to be serialized as a string. Note that setting it to
	// "false" here does not actually filter by it *not* being
	// the default, because Terraform can't distinguish between
	// "false" and "not set".
	defaultForAzStr := ""
	if d.Get("default_for_az").(bool) {
		defaultForAzStr = "true"
	}

	filters := map[string]string{
		"availabilityZone":et("availability_zone").(string),
		"availabilityZoneId": d.Get("availability_zone_id").(string),
		"defaultForAz":orAzStr,
		"state":te").(string),
		"vpc-id":Get("vpc_id").(string),
	}

	if v, ok := d.GetOk("cidr_block"); ok {
		filters["cidrBlock"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_cidr_block"); ok {
		filters["ipv6-cidr-block-association.ipv6-cidr-block"] = v.(string)
	}

	input.Filters = BuildAttributeFilterList(filters)

	if tags, tagsOk := d.GetOk("tags"); tagsOk {
		input.Filters = append(input.Filters, BuildTagFilterList(
			Tags(tftags.New(ctx, tags.(map[string]interface{}))),
		)...)
	}

	input.Filters = append(input.Filters, BuildCustomFilterList(
		d.Get("filter").(*schema.Set),
	)...)
	if len(input.Filters) == 0 {
		// Don't send an empty filters list; the EC2 API won't accept it.
		input.Filters = nil
	}

	subnet, err := FindSubnet(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("EC2 Subnet", err))
	}

	d.SetId(aws.StringValue(subnet.SubnetId))

	d.Set("arn", subnet.SubnetArn)
	d.Set("assign_ipv6_address_on_creation", subnet.AssignIpv6AddressOnCreation)
	d.Set("availability_zone_id", subnet.AvailabilityZoneId)
	d.Set("availability_zone", subnet.AvailabilityZone)
	d.Set("available_ip_address_count", subnet.AvailableIpAddressCount)
	d.Set("cidr_block", subnet.CidrBlock)
	d.Set("customer_owned_ipv4_pool", subnet.CustomerOwnedIpv4Pool)
	d.Set("default_for_az", subnet.DefaultForAz)
	d.Set("enable_dns64", subnet.EnableDns64)
	d.Set("enable_lni_at_device_index", subnet.EnableLniAtDeviceIndex)
	d.Set("ipv6_native", subnet.Ipv6Native)

	// Make sure those values are set, if an IPv6 block exists it'll be set in the loop.
	d.Set("ipv6_cidr_block_association_id", nil)
	d.Set("ipv6_cidr_block", nil)

	for _, v := range subnet.Ipv6CidrBlockAssociationSet {
		if v.Ipv6CidrBlockState != nil && aws.StringValue(v.Ipv6CidrBlockState.State) == ec2.VpcCidrBlockStateCodeAssociated { //we can only ever have 1 IPv6 block associated at once
			d.Set("ipv6_cidr_block_association_id", v.AssociationId)
			d.Set("ipv6_cidr_block", v.Ipv6CidrBlock)
		}
	}

	d.Set("map_customer_owned_ip_on_launch", subnet.MapCustomerOwnedIpOnLaunch)
	d.Set("map_public_ip_on_launch", subnet.MapPublicIpOnLaunch)
	d.Set("outpost_arn", subnet.OutpostArn)
	d.Set("owner_id", subnet.OwnerId)
	d.Set("state", subnet.State)

	if subnet.PrivateDnsNameOptionsOnLaunch != nil {
		d.Set("enable_resource_name_dns_aaaa_record_on_launch", subnet.PrivateDnsNameOptionsOnLaunch.EnableResourceNameDnsAAAARecord)
		d.Set("enable_resource_name_dns_a_record_on_launch", subnet.PrivateDnsNameOptionsOnLaunch.EnableResourceNameDnsARecord)
		d.Set("private_dns_hostname_type_on_launch", subnet.PrivateDnsNameOptionsOnLaunch.HostnameType)
	} else {
		d.Set("enable_resource_name_dns_aaaa_record_on_launch", nil)
		d.Set("enable_resource_name_dns_a_record_on_launch", nil)
		d.Set("private_dns_hostname_type_on_launch", nil)
	}

	if err := d.Set("tags", KeyValueTags(ctx, subnet.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
	}

	d.Set("vpc_id", subnet.VpcId)

	return diags
}
