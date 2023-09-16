// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_ec2_network_insights_analysis", name="Network Insights Analysis")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourceNetworkInsightsAnalysisCreate,
		ReadWithoutTimeout:ourceNetworkInsightsAnalysisRead,
		UpdateWithoutTimeout: resourceNetworkInsightsAnalysisUpdate,
		DeleteWithoutTimeout: resourceNetworkInsightsAnalysisDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"alternate_path_hints": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"component_arn": {
							Type:eString,
							Computed: true,
						},
						"component_id": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"arn": {
				Type:eString,
				Computed: true,
			},
			"explanations": networkInsightsAnalysisExplanationsSchema,
			"filter_in_arns": {
				Type:eSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type:schema.TypeString,
					Validate
func: verify.ValidARN,
func,
			"forward_path_components": networkInsightsAnalysisPathComponentsSchema,
			"network_insights_path_id": {
				Type:eString,
				Required: true,
				ForceNew: true,
			},
			"path_found": {
				Type:eBool,
				Computed: true,
			},
			"return_path_components": networkInsightsAnalysisPathComponentsSchema,
			"start_date": {
				Type:eString,
				Computed: true,
			},
			"status": {
				Type:eString,
				Computed: true,
			},
			"status_message": {
				Type:eString,
				Computed: true,
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"wait_for_completion": {
				Type:eBool,
				Optional: true,
				Default:true,
			},
			"warning_message": {
				Type:eString,
				Computed: true,
			},
		},

		CustomizeDiff: verify.SetTagsDiff,
	}
}

var networkInsightsAnalysisPathComponentsSchema = &schema.Schema{
	Type:eList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"acl_rule": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:eString,
							Computed: true,
						},
						"egress": {
							Type:eBool,
							Computed: true,
						},
						"port_range": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:eInt,
										Computed: true,
									},
									"to": {
										Type:eInt,
										Computed: true,
									},
								},
							},
						},
						"protocol": {
							Type:eString,
							Computed: true,
						},
						"rule_action": {
							Type:eString,
							Computed: true,
						},
						"rule_number": {
							Type:eInt,
							Computed: true,
						},
					},
				},
			},
			"additional_details": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_detail_type": {
							Type:eString,
							Computed: true,
						},
						"component": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"arn": {
										Type:eString,
										Computed: true,
									},
									"id": {
										Type:eString,
										Computed: true,
									},
									"name": {
										Type:eString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"attached_to": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"component": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"destination_vpc": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"inbound_header": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_addresses": {
							Type:eList,
							Computed: true,
							Elem:hema{Type: schema.TypeString},
						},
						"destination_port_ranges": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:eInt,
										Computed: true,
									},
									"to": {
										Type:eInt,
										Computed: true,
									},
								},
							},
						},
						"protocol": {
							Type:eString,
							Computed: true,
						},
						"source_addresses": {
							Type:eList,
							Computed: true,
							Elem:hema{Type: schema.TypeString},
						},
						"source_port_ranges": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:eInt,
										Computed: true,
									},
									"to": {
										Type:eInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"outbound_header": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_addresses": {
							Type:eList,
							Computed: true,
							Elem:hema{Type: schema.TypeString},
						},
						"destination_port_ranges": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:eInt,
										Computed: true,
									},
									"to": {
										Type:eInt,
										Computed: true,
									},
								},
							},
						},
						"protocol": {
							Type:eString,
							Computed: true,
						},
						"source_addresses": {
							Type:eList,
							Computed: true,
							Elem:hema{Type: schema.TypeString},
						},
						"source_port_ranges": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:eInt,
										Computed: true,
									},
									"to": {
										Type:eInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"route_table_route": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_cidr": {
							Type:eString,
							Computed: true,
						},
						"destination_prefix_list_id": {
							Type:eString,
							Computed: true,
						},
						"egress_only_internet_gateway_id": {
							Type:eString,
							Computed: true,
						},
						"gateway_id": {
							Type:eString,
							Computed: true,
						},
						"instance_id": {
							Type:eString,
							Computed: true,
						},
						"nat_gateway_id": {
							Type:eString,
							Computed: true,
						},
						"network_interface_id": {
							Type:eString,
							Computed: true,
						},
						"origin": {
							Type:eString,
							Computed: true,
						},
						"transit_gateway_id": {
							Type:eString,
							Computed: true,
						},
						"vpc_peering_connection_id": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"security_group_rule": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:eString,
							Computed: true,
						},
						"direction": {
							Type:eString,
							Computed: true,
						},
						"port_range": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:eInt,
										Computed: true,
									},
									"to": {
										Type:eInt,
										Computed: true,
									},
								},
							},
						},
						"prefix_list_id": {
							Type:eString,
							Computed: true,
						},
						"protocol": {
							Type:eString,
							Computed: true,
						},
						"security_group_id": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"sequence_number": {
				Type:eInt,
				Computed: true,
			},
			"source_vpc": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"subnet": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"transit_gateway": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"transit_gateway_route_table_route": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attachment_id": {
							Type:eString,
							Computed: true,
						},
						"destination_cidr": {
							Type:eString,
							Computed: true,
						},
						"prefix_list_id": {
							Type:eString,
							Computed: true,
						},
						"resource_id": {
							Type:eString,
							Computed: true,
						},
						"resource_type": {
							Type:eString,
							Computed: true,
						},
						"route_origin": {
							Type:eString,
							Computed: true,
						},
						"state": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"vpc": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
		},
	},
}

var networkInsightsAnalysisExplanationsSchema = &schema.Schema{
	Type:eList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"acl": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"acl_rule": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:eString,
							Computed: true,
						},
						"egress": {
							Type:eBool,
							Computed: true,
						},
						"port_range": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:eInt,
										Computed: true,
									},
									"to": {
										Type:eInt,
										Computed: true,
									},
								},
							},
						},
						"protocol": {
							Type:eString,
							Computed: true,
						},
						"rule_action": {
							Type:eString,
							Computed: true,
						},
						"rule_number": {
							Type:eInt,
							Computed: true,
						},
					},
				},
			},
			"address": {
				Type:eString,
				Computed: true,
			},
			"addresses": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"attached_to": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"availability_zones": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"cidrs": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"classic_load_balancer_listener": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_port": {
							Type:eInt,
							Computed: true,
						},
						"load_balancer_port": {
							Type:eInt,
							Computed: true,
						},
					},
				},
			},
			"component": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"customer_gateway": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"destination": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"destination_vpc": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"direction": {
				Type:eString,
				Computed: true,
			},
			"elastic_load_balancer_listener": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"explanation_code": {
				Type:eString,
				Computed: true,
			},
			"ingress_route_table": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"internet_gateway": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"load_balancer_arn": {
				Type:eString,
				Computed: true,
			},
			"load_balancer_listener_port": {
				Type:eInt,
				Computed: true,
			},
			"load_balancer_target_group": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"load_balancer_target_groups": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"load_balancer_target_port": {
				Type:eInt,
				Computed: true,
			},
			"missing_component": {
				Type:eString,
				Computed: true,
			},
			"nat_gateway": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"network_interface": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"packet_field": {
				Type:eString,
				Computed: true,
			},
			"port": {
				Type:eInt,
				Computed: true,
			},
			"port_ranges": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from": {
							Type:eInt,
							Computed: true,
						},
						"to": {
							Type:eInt,
							Computed: true,
						},
					},
				},
			},
			"prefix_list": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"protocols": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"route_table": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"route_table_route": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_cidr": {
							Type:eString,
							Computed: true,
						},
						"destination_prefix_list_id": {
							Type:eString,
							Computed: true,
						},
						"egress_only_internet_gateway_id": {
							Type:eString,
							Computed: true,
						},
						"gateway_id": {
							Type:eString,
							Computed: true,
						},
						"instance_id": {
							Type:eString,
							Computed: true,
						},
						"nat_gateway_id": {
							Type:eString,
							Computed: true,
						},
						"network_interface_id": {
							Type:eString,
							Computed: true,
						},
						"origin": {
							Type:eString,
							Computed: true,
						},
						"transit_gateway_id": {
							Type:eString,
							Computed: true,
						},
						"vpc_peering_connection_id": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"security_group": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"security_group_rule": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:eString,
							Computed: true,
						},
						"direction": {
							Type:eString,
							Computed: true,
						},
						"port_range": {
							Type:eList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:eInt,
										Computed: true,
									},
									"to": {
										Type:eInt,
										Computed: true,
									},
								},
							},
						},
						"prefix_list_id": {
							Type:eString,
							Computed: true,
						},
						"protocol": {
							Type:eString,
							Computed: true,
						},
						"security_group_id": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"security_groups": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"source_vpc": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:eString,
				Computed: true,
			},
			"subnet": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"subnet_route_table": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"transit_gateway": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"transit_gateway_attachment": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"transit_gateway_route_table": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"transit_gateway_route_table_route": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attachment_id": {
							Type:eString,
							Computed: true,
						},
						"destination_cidr": {
							Type:eString,
							Computed: true,
						},
						"prefix_list_id": {
							Type:eString,
							Computed: true,
						},
						"resource_id": {
							Type:eString,
							Computed: true,
						},
						"resource_type": {
							Type:eString,
							Computed: true,
						},
						"route_origin": {
							Type:eString,
							Computed: true,
						},
						"state": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"vpc": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"vpc_endpoint": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"vpc_peering_connection": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"vpn_connection": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"vpn_gateway": {
				Type:eList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:eString,
							Computed: true,
						},
						"id": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
		},
	},
}


func resourceNetworkInsightsAnalysisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)
funcut := &ec2.StartNetworkInsightsAnalysisInput{
		NetworkInsightsPathId: aws.String(d.Get("network_insights_path_id").(string)),
		TagSpecifications:ificationsIn(ctx, ec2.ResourceTypeNetworkInsightsAnalysis),
	}

	if v, ok := d.GetOk("filter_in_arns"); ok && v.(*schema.Set).Len() > 0 {
		input.FilterInArns = flex.ExpandStringSet(v.(*schema.Set))
	}

	log.Printf("[DEBUG] Creating EC2 Network Insights Analysis: %s", input)
	output, err := conn.StartNetworkInsightsAnalysisWithContext(ctx, input)

	if err != nil {
		return diag.Errorf("creating EC2 Network Insights Analysis: %s", err)
	}

	d.SetId(aws.StringValue(output.NetworkInsightsAnalysis.NetworkInsightsAnalysisId))

	if d.Get("wait_for_completion").(bool) {
		if _, err := WaitNetworkInsightsAnalysisCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
			return diag.Errorf("waiting for EC2 Network Insights Analysis (%s) create: %s", d.Id(), err)
		}
	}

	return resourceNetworkInsightsAnalysisRead(ctx, d, meta)
}


func resourceNetworkInsightsAnalysisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

func
	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] EC2 Network Insights Analysis (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.Errorf("reading EC2 Network Insights Analysis (%s): %s", d.Id(), err)
	}

	if err := d.Set("alternate_path_hints", flattenAlternatePathHints(output.AlternatePathHints)); err != nil {
		return diag.Errorf("setting alternate_path_hints: %s", err)
	}
	d.Set("arn", output.NetworkInsightsAnalysisArn)
	if err := d.Set("explanations", flattenExplanations(output.Explanations)); err != nil {
		return diag.Errorf("setting explanations: %s", err)
	}
	d.Set("filter_in_arns", aws.StringValueSlice(output.FilterInArns))
	if err := d.Set("forward_path_components", flattenPathComponents(output.ForwardPathComponents)); err != nil {
		return diag.Errorf("setting forward_path_components: %s", err)
	}
	d.Set("network_insights_path_id", output.NetworkInsightsPathId)
	d.Set("path_found", output.NetworkPathFound)
	if err := d.Set("return_path_components", flattenPathComponents(output.ReturnPathComponents)); err != nil {
		return diag.Errorf("setting return_path_components: %s", err)
	}
	d.Set("start_date", output.StartDate.Format(time.RFC3339))
	d.Set("status", output.Status)
	d.Set("status_message", output.StatusMessage)
	d.Set("warning_message", output.WarningMessage)

	setTagsOut(ctx, output.Tags)

	return nil
}


func resourceNetworkInsightsAnalysisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Tags only.
	return resourceNetworkInsightsAnalysisRead(ctx, d, meta)
}
func
func resourceNetworkInsightsAnalysisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	log.Printf("[DEBUG] Deleting EC2 Network Insights Analysis: %s", d.Id())
	_, err := conn.DeleteNetworkInsightsAnalysisWithContext(ctx, &ec2.DeleteNetworkInsightsAnalysisInput{
func

	if tfawserr.ErrCodeEquals(err, errCodeInvalidNetworkInsightsAnalysisIdNotFound) {
		return nil
	}

	if err != nil {
		return diag.Errorf("deleting EC2 Network Insights Analysis (%s): %s", d.Id(), err)
	}

	return nil
}


func flattenAdditionalDetail(apiObject *ec2.AdditionalDetail) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
funcv := apiObject.AdditionalDetailType; v != nil {
		tfMap["additional_detail_type"] = aws.StringValue(v)
	}

	if v := apiObject.Component; v != nil {
		tfMap["component"] = []interface{}{flattenAnalysisComponent(v)}
	}

	return tfMap
}


func flattenAdditionalDetails(apiObjects []*ec2.AdditionalDetail) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}

	var tfList []interface{}

func apiObject == nil {
			continue
		}

		tfList = append(tfList, flattenAdditionalDetail(apiObject))
	}

	return tfList
}


func flattenAlternatePathHint(apiObject *ec2.AlternatePathHint) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.ComponentArn; v != nil {
func

	if v := apiObject.ComponentId; v != nil {
		tfMap["component_id"] = aws.StringValue(v)
	}

	return tfMap
}


func flattenAlternatePathHints(apiObjects []*ec2.AlternatePathHint) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}

	var tfList []interface{}

	for _, apiObject := range apiObjects {
		if apiObject == nil {
func

		tfList = append(tfList, flattenAlternatePathHint(apiObject))
	}

	return tfList
}


func flattenAnalysisAclRule(apiObject *ec2.AnalysisAclRule) map[string]interface{} { // nosemgrep:ci.caps2-in-
func-name
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.Cidr; v != nil {
		tfMap["cidr"] = aws.StringValue(v)
func
funcMap["egress"] = aws.BoolValue(v)
	}

	if v := apiObject.PortRange; v != nil {
		tfMap["port_range"] = []interface{}{flattenPortRange(v)}
	}

	if v := apiObject.Protocol; v != nil {
		tfMap["protocol"] = aws.StringValue(v)
	}

	if v := apiObject.RuleAction; v != nil {
		tfMap["rule_action"] = aws.StringValue(v)
	}

	if v := apiObject.RuleNumber; v != nil {
		tfMap["rule_number"] = aws.Int64Value(v)
	}

	return tfMap
}


func flattenAnalysisLoadBalancerListener(apiObject *ec2.AnalysisLoadBalancerListener) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.InstancePort; v != nil {
		tfMap["instance_port"] = aws.Int64Value(v)
	}

	if v := apiObject.LoadBalancerPort; v != nil {
func

	return tfMap
}


func flattenAnalysisComponent(apiObject *ec2.AnalysisComponent) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.Arn; v != nil {
		tfMap["arn"] = aws.StringValue(v)
	}

	if v := apiObject.Id; v != nil {
		tfMap["id"] = aws.StringValue(v)
func
	if v := apiObject.Name; v != nil {
		tfMap["name"] = aws.StringValue(v)
	}

	return tfMap
}


func flattenAnalysisComponents(apiObjects []*ec2.AnalysisComponent) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}

	var tfList []interface{}

	for _, apiObject := range apiObjects {
		if apiObject == nil {
			continue
		}

		tfList = append(tfList, flattenAnalysisComponent(apiObject))
	}
funcurn tfList
}


func flattenAnalysisLoadBalancerTarget(apiObject *ec2.AnalysisLoadBalancerTarget) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.Address; v != nil {
		tfMap["address"] = aws.StringValue(v)
	}

	if v := apiObject.AvailabilityZone; v != nil {
		tfMap["availability_zone"] = aws.StringValue(v)
	}

funcMap["instance"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.Port; v != nil {
		tfMap["port"] = aws.Int64Value(v)
	}

	return tfMap
}


func flattenAnalysisPacketHeader(apiObject *ec2.AnalysisPacketHeader) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.DestinationAddresses; v != nil {
		tfMap["destination_addresses"] = aws.StringValueSlice(v)
	}

	if v := apiObject.DestinationPortRanges; v != nil {
		tfMap["destination_port_ranges"] = flattenPortRanges(v)
	}

	if v := apiObject.Protocol; v != nil {
func

	if v := apiObject.SourceAddresses; v != nil {
		tfMap["source_addresses"] = aws.StringValueSlice(v)
	}

	if v := apiObject.SourcePortRanges; v != nil {
		tfMap["source_port_ranges"] = flattenPortRanges(v)
	}

	return tfMap
}


func flattenAnalysisRouteTableRoute(apiObject *ec2.AnalysisRouteTableRoute) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.DestinationCidr; v != nil {
		tfMap["destination_cidr"] = aws.StringValue(v)
	}

	if v := apiObject.DestinationPrefixListId; v != nil {
		tfMap["destination_prefix_list_id"] = aws.StringValue(v)
	}

	if v := apiObject.EgressOnlyInternetGatewayId; v != nil {
		tfMap["egress_only_internet_gateway_id"] = aws.StringValue(v)
func
	if v := apiObject.GatewayId; v != nil {
		tfMap["gateway_id"] = aws.StringValue(v)
	}

	if v := apiObject.InstanceId; v != nil {
		tfMap["instance_id"] = aws.StringValue(v)
	}

	if v := apiObject.NatGatewayId; v != nil {
		tfMap["nat_gateway_id"] = aws.StringValue(v)
	}

	if v := apiObject.NetworkInterfaceId; v != nil {
		tfMap["network_interface_id"] = aws.StringValue(v)
	}

	if v := apiObject.Origin; v != nil {
		tfMap["origin"] = aws.StringValue(v)
	}

	if v := apiObject.TransitGatewayId; v != nil {
		tfMap["transit_gateway_id"] = aws.StringValue(v)
	}

	if v := apiObject.VpcPeeringConnectionId; v != nil {
		tfMap["vpc_peering_connection_id"] = aws.StringValue(v)
	}

	return tfMap
}


func flattenAnalysisSecurityGroupRule(apiObject *ec2.AnalysisSecurityGroupRule) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.Cidr; v != nil {
		tfMap["cidr"] = aws.StringValue(v)
	}

	if v := apiObject.PortRange; v != nil {
		tfMap["port_range"] = []interface{}{flattenPortRange(v)}
	}

	if v := apiObject.PrefixListId; v != nil {
		tfMap["prefix_list_id"] = aws.StringValue(v)
	}
funcv := apiObject.Protocol; v != nil {
		tfMap["protocol"] = aws.StringValue(v)
	}

	if v := apiObject.SecurityGroupId; v != nil {
		tfMap["security_group_id"] = aws.StringValue(v)
	}

	return tfMap
}


func flattenExplanation(apiObject *ec2.Explanation) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.Acl; v != nil {
		tfMap["acl"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.AclRule; v != nil {
		tfMap["acl_rule"] = []interface{}{flattenAnalysisAclRule(v)}
	}

	if v := apiObject.Address; v != nil {
		tfMap["address"] = aws.StringValue(v)
	}

funcMap["addresses"] = aws.StringValueSlice(v)
	}

	if v := apiObject.AttachedTo; v != nil {
		tfMap["attached_to"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.AvailabilityZones; v != nil {
		tfMap["availability_zones"] = aws.StringValueSlice(v)
	}

	if v := apiObject.Cidrs; v != nil {
		tfMap["cidrs"] = aws.StringValueSlice(v)
	}

	if v := apiObject.ClassicLoadBalancerListener; v != nil {
		tfMap["classic_load_balancer_listener"] = []interface{}{flattenAnalysisLoadBalancerListener(v)}
	}

	if v := apiObject.Component; v != nil {
		tfMap["component"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.CustomerGateway; v != nil {
		tfMap["customer_gateway"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.Destination; v != nil {
		tfMap["destination"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.DestinationVpc; v != nil {
		tfMap["destination_vpc"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.Direction; v != nil {
		tfMap["direction"] = aws.StringValue(v)
	}

	if v := apiObject.ElasticLoadBalancerListener; v != nil {
		tfMap["elastic_load_balancer_listener"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.ExplanationCode; v != nil {
		tfMap["explanation_code"] = aws.StringValue(v)
	}

	if v := apiObject.IngressRouteTable; v != nil {
		tfMap["ingress_route_table"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.InternetGateway; v != nil {
		tfMap["internet_gateway"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.LoadBalancerArn; v != nil {
		tfMap["load_balancer_arn"] = aws.StringValue(v)
	}

	if v := apiObject.LoadBalancerListenerPort; v != nil {
		tfMap["load_balancer_listener_port"] = aws.Int64Value(v)
	}

	if v := apiObject.LoadBalancerTarget; v != nil {
		tfMap["load_balancer_target"] = []interface{}{flattenAnalysisLoadBalancerTarget(v)}
	}

	if v := apiObject.LoadBalancerTargetGroup; v != nil {
		tfMap["load_balancer_target_group"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.LoadBalancerTargetGroups; v != nil {
		tfMap["load_balancer_target_group"] = flattenAnalysisComponents(v)
	}

	if v := apiObject.LoadBalancerTargetPort; v != nil {
		tfMap["load_balancer_target_port"] = aws.Int64Value(v)
	}

	if v := apiObject.MissingComponent; v != nil {
		tfMap["missing_component"] = aws.StringValue(v)
	}

	if v := apiObject.NatGateway; v != nil {
		tfMap["nat_gateway"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.NetworkInterface; v != nil {
		tfMap["network_interface"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.PacketField; v != nil {
		tfMap["packet_field"] = aws.StringValue(v)
	}

	if v := apiObject.Port; v != nil {
		tfMap["port"] = aws.Int64Value(v)
	}

	if v := apiObject.PortRanges; v != nil {
		tfMap["port_ranges"] = flattenPortRanges(v)
	}

	if v := apiObject.PrefixList; v != nil {
		tfMap["prefix_list"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.Protocols; v != nil {
		tfMap["protocols"] = aws.StringValueSlice(v)
	}

	if v := apiObject.RouteTable; v != nil {
		tfMap["route_table"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.RouteTableRoute; v != nil {
		tfMap["route_table_route"] = []interface{}{flattenAnalysisRouteTableRoute(v)}
	}

	if v := apiObject.SecurityGroup; v != nil {
		tfMap["security_group"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.SecurityGroupRule; v != nil {
		tfMap["security_group_rule"] = []interface{}{flattenAnalysisSecurityGroupRule(v)}
	}

	if v := apiObject.SecurityGroups; v != nil {
		tfMap["security_groups"] = flattenAnalysisComponents(v)
	}

	if v := apiObject.SourceVpc; v != nil {
		tfMap["source_vpc"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.State; v != nil {
		tfMap["state"] = aws.StringValue(v)
	}

	if v := apiObject.Subnet; v != nil {
		tfMap["subnet"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.SubnetRouteTable; v != nil {
		tfMap["subnet_route_table"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.TransitGateway; v != nil {
		tfMap["transit_gateway"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.TransitGatewayAttachment; v != nil {
		tfMap["transit_gateway_attachment"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.TransitGatewayRouteTable; v != nil {
		tfMap["transit_gateway_route_table"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.TransitGatewayRouteTableRoute; v != nil {
		tfMap["transit_gateway_route_table_route"] = []interface{}{flattenTransitGatewayRouteTableRoute(v)}
	}

	if v := apiObject.Vpc; v != nil {
		tfMap["vpc"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.VpcEndpoint; v != nil {
		tfMap["vpc_endpoint"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.VpcPeeringConnection; v != nil {
		tfMap["vpc_peering_connection"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.VpnConnection; v != nil {
		tfMap["vpn_connection"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.VpnGateway; v != nil {
		tfMap["vpn_gateway"] = []interface{}{flattenAnalysisComponent(v)}
	}

	return tfMap
}


func flattenExplanations(apiObjects []*ec2.Explanation) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}

	var tfList []interface{}

	for _, apiObject := range apiObjects {
		if apiObject == nil {
			continue
		}

		tfList = append(tfList, flattenExplanation(apiObject))
	}

	return tfList
}


func flattenPathComponent(apiObject *ec2.PathComponent) map[string]interface{} {
functurn nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.AclRule; v != nil {
		tfMap["acl_rule"] = []interface{}{flattenAnalysisAclRule(v)}
	}

	if v := apiObject.AdditionalDetails; v != nil {
		tfMap["additional_details"] = flattenAdditionalDetails(v)
	}

	if v := apiObject.Component; v != nil {
		tfMap["component"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.DestinationVpc; v != nil {
		tfMap["destination_vpc"] = []interface{}{flattenAnalysisComponent(v)}
func
	if v := apiObject.InboundHeader; v != nil {
		tfMap["inbound_header"] = []interface{}{flattenAnalysisPacketHeader(v)}
	}

	if v := apiObject.OutboundHeader; v != nil {
		tfMap["outbound_header"] = []interface{}{flattenAnalysisPacketHeader(v)}
	}

	if v := apiObject.RouteTableRoute; v != nil {
		tfMap["route_table_route"] = []interface{}{flattenAnalysisRouteTableRoute(v)}
	}

	if v := apiObject.SecurityGroupRule; v != nil {
		tfMap["security_group_rule"] = []interface{}{flattenAnalysisSecurityGroupRule(v)}
	}

	if v := apiObject.SequenceNumber; v != nil {
		tfMap["sequence_number"] = aws.Int64Value(v)
	}

	if v := apiObject.SourceVpc; v != nil {
		tfMap["source_vpc"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.Subnet; v != nil {
		tfMap["subnet"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.TransitGateway; v != nil {
		tfMap["transit_gateway"] = []interface{}{flattenAnalysisComponent(v)}
	}

	if v := apiObject.TransitGatewayRouteTableRoute; v != nil {
		tfMap["transit_gateway_route_table_route"] = []interface{}{flattenTransitGatewayRouteTableRoute(v)}
	}

	if v := apiObject.Vpc; v != nil {
		tfMap["vpc"] = []interface{}{flattenAnalysisComponent(v)}
	}

	return tfMap
}


func flattenPathComponents(apiObjects []*ec2.PathComponent) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}

	var tfList []interface{}

	for _, apiObject := range apiObjects {
		if apiObject == nil {
			continue
		}

		tfList = append(tfList, flattenPathComponent(apiObject))
	}

	return tfList
}


func flattenPortRange(apiObject *ec2.PortRange) map[string]interface{} {
	if apiObject == nil {
		return nil
func
	tfMap := map[string]interface{}{}

	if v := apiObject.From; v != nil {
		tfMap["from"] = aws.Int64Value(v)
	}

	if v := apiObject.To; v != nil {
		tfMap["to"] = aws.Int64Value(v)
	}

	return tfMap
}


func flattenPortRanges(apiObjects []*ec2.PortRange) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}
func tfList []interface{}

	for _, apiObject := range apiObjects {
		if apiObject == nil {
			continue
		}

		tfList = append(tfList, flattenPortRange(apiObject))
	}

	return tfList
}


func flattenTransitGatewayRouteTableRoute(apiObject *ec2.TransitGatewayRouteTableRoute) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

func
	if v := apiObject.AttachmentId; v != nil {
		tfMap["attachment_id"] = aws.StringValue(v)
	}

	if v := apiObject.DestinationCidr; v != nil {
		tfMap["destination_cidr"] = aws.StringValue(v)
	}

	if v := apiObject.PrefixListId; v != nil {
		tfMap["prefix_list_id"] = aws.StringValue(v)
	}

	if v := apiObject.ResourceId; v != nil {
		tfMap["resource_id"] = aws.StringValue(v)
	}

	if v := apiObject.ResourceType; v != nil {
		tfMap["resource_type"] = aws.StringValue(v)
func
	if v := apiObject.RouteOrigin; v != nil {
		tfMap["route_origin"] = aws.StringValue(v)
	}

	if v := apiObject.State; v != nil {
		tfMap["state"] = aws.StringValue(v)
	}

	return tfMap
}
