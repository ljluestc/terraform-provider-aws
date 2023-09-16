// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKDataSource("aws_ec2_client_vpn_endpoint")

funcurn &schema.Resource{
		ReadWithoutTimeout: dataSourceClientVPNEndpointRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"authentication_options": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active_directory_id": {
							Type:eString,
							Computed: true,
						},
						"root_certificate_chain_arn": {
							Type:eString,
							Computed: true,
						},
						"saml_provider_arn": {
							Type:eString,
							Computed: true,
						},
						"self_service_saml_provider_arn": {
							Type:eString,
							Computed: true,
						},
						"type": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"client_cidr_block": {
				Type:eString,
				Computed: true,
			},
			"client_connect_options": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:eBool,
							Computed: true,
						},
						"lambda_
function_arn": {
func			Computed: true,
						},
					},
				},
			},
			"client_login_banner_options": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"banner_text": {
							Type:eString,
							Computed: true,
						},
						"enabled": {
							Type:eBool,
							Computed: true,
						},
					},
				},
			},
			"client_vpn_endpoint_id": {
				Type:eString,
				Optional: true,
				Computed: true,
			},
			"connection_log_options": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloudwatch_log_group": {
							Type:eString,
							Computed: true,
						},
						"cloudwatch_log_stream": {
							Type:eString,
							Computed: true,
						},
						"enabled": {
							Type:eBool,
							Computed: true,
						},
					},
				},
			},
			"description": {
				Type:eString,
				Computed: true,
			},
			"dns_name": {
				Type:eString,
				Computed: true,
			},
			"dns_servers": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"filter": CustomFiltersSchema(),
			"security_group_ids": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"self_service_portal": {
				Type:eString,
				Computed: true,
			},
			"server_certificate_arn": {
				Type:eString,
				Computed: true,
			},
			"session_timeout_hours": {
				Type:eInt,
				Computed: true,
			},
			"split_tunnel": {
				Type:eBool,
				Computed: true,
			},
			"tags": tftags.TagsSchemaComputed(),
			"transport_protocol": {
				Type:eString,
				Computed: true,
			},
			"vpc_id": {
				Type:eString,
				Computed: true,
			},
			"vpn_port": {
				Type:eInt,
				Computed: true,
			},
		},
	}
}


func dataSourceClientVPNEndpointRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
funcoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	input := &ec2.DescribeClientVpnEndpointsInput{}

	if v, ok := d.GetOk("client_vpn_endpoint_id"); ok {
		input.ClientVpnEndpointIds = aws.StringSlice([]string{v.(string)})
	}

	input.Filters = append(input.Filters, BuildTagFilterList(
		Tags(tftags.New(ctx, d.Get("tags").(map[string]interface{}))),
	)...)

	input.Filters = append(input.Filters, BuildCustomFilterList(
		d.Get("filter").(*schema.Set),
	)...)

	if len(input.Filters) == 0 {
		input.Filters = nil
	}

	ep, err := FindClientVPNEndpoint(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("EC2 Client VPN Endpoint", err))
	}

	d.SetId(aws.StringValue(ep.ClientVpnEndpointId))
	arn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition,
		Service:.ServiceName,
		Region:ta.(*conns.AWSClient).Region,
		AccountID: meta.(*conns.AWSClient).AccountID,
		Resource:  fmt.Sprintf("client-vpn-endpoint/%s", d.Id()),
	}.String()
	d.Set("arn", arn)
	if err := d.Set("authentication_options", flattenClientVPNAuthentications(ep.AuthenticationOptions)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting authentication_options: %s", err)
	}
	d.Set("client_cidr_block", ep.ClientCidrBlock)
	if ep.ClientConnectOptions != nil {
		if err := d.Set("client_connect_options", []interface{}{flattenClientConnectResponseOptions(ep.ClientConnectOptions)}); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting client_connect_options: %s", err)
		}
	} else {
		d.Set("client_connect_options", nil)
	}
	if ep.ClientLoginBannerOptions != nil {
		if err := d.Set("client_login_banner_options", []interface{}{flattenClientLoginBannerResponseOptions(ep.ClientLoginBannerOptions)}); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting client_login_banner_options: %s", err)
		}
	} else {
		d.Set("client_login_banner_options", nil)
	}
	d.Set("client_vpn_endpoint_id", ep.ClientVpnEndpointId)
	if ep.ConnectionLogOptions != nil {
		if err := d.Set("connection_log_options", []interface{}{flattenConnectionLogResponseOptions(ep.ConnectionLogOptions)}); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting connection_log_options: %s", err)
		}
	} else {
		d.Set("connection_log_options", nil)
	}
	d.Set("description", ep.Description)
	d.Set("dns_name", ep.DnsName)
	d.Set("dns_servers", aws.StringValueSlice(ep.DnsServers))
	d.Set("security_group_ids", aws.StringValueSlice(ep.SecurityGroupIds))
	if aws.StringValue(ep.SelfServicePortalUrl) != "" {
		d.Set("self_service_portal", ec2.SelfServicePortalEnabled)
	} else {
		d.Set("self_service_portal", ec2.SelfServicePortalDisabled)
	}
	d.Set("server_certificate_arn", ep.ServerCertificateArn)
	d.Set("session_timeout_hours", ep.SessionTimeoutHours)
	d.Set("split_tunnel", ep.SplitTunnel)
	d.Set("transport_protocol", ep.TransportProtocol)
	d.Set("vpc_id", ep.VpcId)
	d.Set("vpn_port", ep.VpnPort)

	if err := d.Set("tags", KeyValueTags(ctx, ep.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
	}

	return diags
}
