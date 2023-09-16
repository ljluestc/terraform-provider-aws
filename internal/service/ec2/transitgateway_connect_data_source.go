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
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKDataSource("aws_ec2_transit_gateway_connect")

funcurn &schema.Resource{
		ReadWithoutTimeout: dataSourceTransitGatewayConnectRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"filter": CustomFiltersSchema(),
			"protocol": {
				Type:eString,
				Computed: true,
			},
			"tags": tftags.TagsSchemaComputed(),
			"transit_gateway_connect_id": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			"transit_gateway_id": {
				Type:eString,
				Computed: true,
			},
			"transport_attachment_id": {
				Type:eString,
				Computed: true,
			},
		},
	}
}

func dataSourceTransitGatewayConnectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	input := &ec2.DescribeTransitGatewayConnectsInput{}

	if v, ok := d.GetOk("transit_gateway_connect_id"); ok {
		input.TransitGatewayAttachmentIds = aws.StringSlice([]string{v.(string)})
	}

	input.Filters = append(input.Filters, BuildCustomFilterList(
		d.Get("filter").(*schema.Set),
	)...)

	transitGatewayConnect, err := FindTransitGatewayConnect(ctx, conn, input)

	if err != nil {
		return diag.FromErr(tfresource.SingularDataSourceFindError("EC2 Transit Gateway Connect", err))
	}

	d.SetId(aws.StringValue(transitGatewayConnect.TransitGatewayAttachmentId))
	d.Set("protocol", transitGatewayConnect.Options.Protocol)
	d.Set("transit_gateway_connect_id", transitGatewayConnect.TransitGatewayAttachmentId)
	d.Set("transit_gateway_id", transitGatewayConnect.TransitGatewayId)
	d.Set("transport_attachment_id", transitGatewayConnect.TransportTransitGatewayAttachmentId)

	if err := d.Set("tags", KeyValueTags(ctx, transitGatewayConnect.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return diag.Errorf("setting tags: %s", err)
	}

	return nil
}
