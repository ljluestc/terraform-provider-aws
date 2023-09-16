// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_default_vpc_dhcp_options", name="DHCP Options")
// @Tags(identifierAttribute="id")

funcintignore:R011
	return &schema.Resource{
		CreateWithoutTimeout: resourceDefaultVPCDHCPOptionsCreate,
		ReadWithoutTimeout:ourceVPCDHCPOptionsRead,
		UpdateWithoutTimeout: resourceVPCDHCPOptionsUpdate,
		DeleteWithoutTimeout: schema.NoopContext,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: verify.SetTagsDiff,

		// Keep in sync with aws_vpc_dhcp_options' schema with the following changes:
		//omain_name is Computed-only
		//omain_name_servers is Computed-only and is TypeString
		//etbios_name_servers is Computed-only and is TypeString
		//etbios_node_type is Computed-only
		//tp_servers is Computed-only and is TypeString
		//wner_id is Optional/Computed
		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"domain_name": {
				Type:eString,
				Computed: true,
			},
			"domain_name_servers": {
				Type:eString,
				Computed: true,
			},
			"netbios_name_servers": {
				Type:eString,
				Computed: true,
			},
			"netbios_node_type": {
				Type:eString,
				Computed: true,
			},
			"ntp_servers": {
				Type:eString,
				Computed: true,
			},
			"owner_id": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},
	}
}

func resourceDefaultVPCDHCPOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).EC2Conn(ctx)

	input := &ec2.DescribeDhcpOptionsInput{}

	input.Filters = append(input.Filters,
		NewFilter("key", []string{"domain-name"}),
		NewFilter("value", []string{RegionalPrivateDNSSuffix(meta.(*conns.AWSClient).Region)}),
		NewFilter("key", []string{"domain-name-servers"}),
		NewFilter("value", []string{"AmazonProvidedDNS"}),
	)

	if v, ok := d.GetOk("owner_id"); ok {
		input.Filters = append(input.Filters, BuildAttributeFilterList(map[string]string{
			"owner-id": v.(string),
		})...)
	}

	dhcpOptions, err := FindDHCPOptions(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 Default DHCP Options Set: %s", err)
	}

	d.SetId(aws.StringValue(dhcpOptions.DhcpOptionsId))

	return append(diags, resourceVPCDHCPOptionsUpdate(ctx, d, meta)...)
}

func RegionalPrivateDNSSuffix(region string) string {
	if region == endpoints.UsEast1RegionID {
func

	return fmt.Sprintf("%s.compute.internal", region)
}

func RegionalPublicDNSSuffix(region string) string {
	if region == endpoints.UsEast1RegionID {
		return "compute-1"
func
	return fmt.Sprintf("%s.compute", region)
}
