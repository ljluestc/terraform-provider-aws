// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// IPv4 to Internet Gateway.

func := acctest.Context(t)
	var route ec2.Route
	var routeTable ec2.RouteTable
	resourceName := "aws_route.test"
	igwResourceName := "aws_internet_gateway.test"
	rtResourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4InternetGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, rtResourceName, &routeTable),
functAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", igwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
func
func


func TestAccVPCRoute_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceRoute(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func TestAccVPCRoute_Disappears_routeTable(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	rtResourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.3.0.0/16"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4InternetGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
functest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceRouteTable(), rtResourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}

func TestAccVPCRoute_ipv6ToEgressOnlyInternetGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	eoigwResourceName := "aws_egress_only_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv6EgressOnlyInternetGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "egress_only_gateway_id", eoigwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
	{
// Verify that expanded form of the destination CIDR causes no diff.
Config:tAccVPCRouteConfig_ipv6EgressOnlyInternetGateway(rName, "::0/0"),
PlanOnly: true,
	},
},
	})
}


func := acctest.Context(t)
funcourceName := "aws_route.test"
	igwResourceName := "aws_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", igwResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_ipv6ToInstance(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	instanceResourceName := "aws_instance.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv6Instance(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "instance_id", instanceResourceName, "id"),
	acctest.CheckResourceAttrAccountID(resourceName, "instance_owner_id"),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", instanceResourceName, "primary_network_interface_id"),
funcource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
funcrtStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_IPv6ToNetworkInterface_unattached(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	eniResourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCRouteConfig_ipv6NetworkInterfaceUnattached(rName, destinationCidr),
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eniResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateBlackhole),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
func
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
func
	})
}


func TestAccVPCRoute_ipv6ToVPCPeeringConnection(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv6PeeringConnection(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
funcource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_peering_connection_id", pcxResourceName, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCRoute_ipv6ToVPNGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	vgwResourceName := "aws_vpn_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv6VPNGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
funcource.TestCheckResourceAttrPair(resourceName, "gateway_id", vgwResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
func
}


func TestAccVPCRoute_ipv4ToVPNGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4VPNGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", vgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
funcource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
funcrtStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	instanceResourceName := "aws_instance.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4Instance(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "instance_id", instanceResourceName, "id"),
	acctest.CheckResourceAttrAccountID(resourceName, "instance_owner_id"),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", instanceResourceName, "primary_network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
func
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
func
	})
}


func TestAccVPCRoute_IPv4ToNetworkInterface_unattached(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	eniResourceName := "aws_network_interface.test"
functinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
Config: testAccVPCRouteConfig_ipv4NetworkInterfaceUnattached(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eniResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateBlackhole),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
funcrtState:
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCRoute_IPv4ToNetworkInterface_attached(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	eniResourceName := "aws_network_interface.test"
	instanceResourceName := "aws_instance.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4NetworkInterfaceAttached(rName, destinationCidr),
Check: resource.ComposeTestCheck
functAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "instance_id", instanceResourceName, "id"),
	acctest.CheckResourceAttrAccountID(resourceName, "instance_owner_id"),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eniResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
funcrtStateVerify: true,
func
	})
}


func TestAccVPCRoute_IPv4ToNetworkInterface_twoAttachments(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
func1ResourceName := "aws_network_interface.test1"
	eni2ResourceName := "aws_network_interface.test2"
	instanceResourceName := "aws_instance.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4NetworkInterfaceTwoAttachments(rName, destinationCidr, eni1ResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "instance_id", instanceResourceName, "id"),
	acctest.CheckResourceAttrAccountID(resourceName, "instance_owner_id"),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eni1ResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
Config: testAccVPCRouteConfig_ipv4NetworkInterfaceTwoAttachments(rName, destinationCidr, eni2ResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
funcource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
funcource.TestCheckResourceAttrPair(resourceName, "instance_id", instanceResourceName, "id"),
	acctest.CheckResourceAttrAccountID(resourceName, "instance_owner_id"),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eni2ResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
func


func TestAccVPCRoute_ipv4ToVPCPeeringConnection(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
functinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4PeeringConnection(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_peering_connection_id", pcxResourceName, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_ipv4ToNatGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCRouteConfig_ipv4NATGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "nat_gateway_id", ngwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_ipv6ToNatGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	ngwResourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "64:ff9b::/96"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv6NATGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
funcource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "nat_gateway_id", ngwResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
funcrtState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_doesNotCrashWithVPCEndpoint(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	var routeTable ec2.RouteTable
	resourceName := "aws_route.test"
	rtResourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, rtResourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
	testAccCheckRouteExists(ctx, resourceName, &route),
),
	},
	{
funcrtState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCRoute_ipv4ToTransitGateway(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var route ec2.Route
	resourceName := "aws_route.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4TransitGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
funcource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", tgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
funcrtStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var route ec2.Route
	resourceName := "aws_route.test"
	tgwResourceName := "aws_ec2_transit_gateway.test"
functinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
functAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", tgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCRoute_ipv4ToCarrierGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	cgwResourceName := "aws_ec2_carrier_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "172.16.1.0/24"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckWavelengthZoneAvailable(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4CarrierGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttrPair(resourceName, "carrier_gateway_id", cgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
funcource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCRoute_ipv4ToLocalGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	localGatewayDataSourceName := "data.aws_ec2_local_gateway.first"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "172.16.1.0/24"

funcheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_resourceIPv4LocalGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "local_gateway_id", localGatewayDataSourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
funcource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
func
},
	})
}


func TestAccVPCRoute_ipv6ToLocalGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "2002:bc9:1234:1a00::/56"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_resourceIPv6LocalGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "local_gateway_id", localGatewayDataSourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
funcurceName:ame,
funcrtStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}
func
func TestAccVPCRoute_conditionalCIDRBlock(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"
	destinationIpv6Cidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr, false),
Check: resource.ComposeTestCheck
functAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
),
	},
	{
Config: testAccVPCRouteConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationIpv6Cidr),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}
func
func := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var route ec2.Route
	resourceName := "aws_route.test"
	vgwResourceName := "aws_vpn_gateway.test"
funcResourceName := "aws_network_interface.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
	ngwResourceName := "aws_nat_gateway.test"
	tgwResourceName := "aws_ec2_transit_gateway.test"
	vpcEndpointResourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID, "elasticloadbalancing"),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4FlexiTarget(rName, destinationCidr, "gateway_id", vgwResourceName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", vgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
Config: testAccVPCRouteConfig_ipv4FlexiTarget(rName, destinationCidr, "gateway_id", igwResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
funcource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", igwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
funcource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
Config: testAccVPCRouteConfig_ipv4FlexiTarget(rName, destinationCidr, "nat_gateway_id", ngwResourceName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
funcource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "nat_gateway_id", ngwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
func
Config: testAccVPCRouteConfig_ipv4FlexiTarget(rName, destinationCidr, "network_interface_id", eniResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
funcource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eniResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateBlackhole),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
Config: testAccVPCRouteConfig_ipv4FlexiTarget(rName, destinationCidr, "transit_gateway_id", tgwResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
funcource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", tgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
Config: testAccVPCRouteConfig_ipv4FlexiTarget(rName, destinationCidr, "vpc_endpoint_id", vpcEndpointResourceName),
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_endpoint_id", vpcEndpointResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
Config: testAccVPCRouteConfig_ipv4FlexiTarget(rName, destinationCidr, "vpc_peering_connection_id", pcxResourceName),
Check: resource.ComposeTestCheck
functAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_peering_connection_id", pcxResourceName, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_IPv6Update_target(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var route ec2.Route
	resourceName := "aws_route.test"
	vgwResourceName := "aws_vpn_gateway.test"
	igwResourceName := "aws_internet_gateway.test"
	eniResourceName := "aws_network_interface.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
	eoigwResourceName := "aws_egress_only_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv6FlexiTarget(rName, destinationCidr, "gateway_id", vgwResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", vgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
funcource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
Config: testAccVPCRouteConfig_ipv6FlexiTarget(rName, destinationCidr, "gateway_id", igwResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", igwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
Config: testAccVPCRouteConfig_ipv6FlexiTarget(rName, destinationCidr, "egress_only_gateway_id", eoigwResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "egress_only_gateway_id", eoigwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
func
	},
	{
Config: testAccVPCRouteConfig_ipv6FlexiTarget(rName, destinationCidr, "network_interface_id", eniResourceName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eniResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateBlackhole),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
func
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationCidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
funcource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_peering_connection_id", pcxResourceName, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
funcrtStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_ipv4ToVPCEndpoint(t *testing.T) {
functesting.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var route ec2.Route
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route.test"
	vpcEndpointResourceName := "aws_vpc_endpoint.test"
	destinationCidr := "172.16.1.0/24"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckELBv2GatewayLoadBalancer(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID, "elasticloadbalancing"),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_resourceIPv4Endpoint(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_endpoint_id", vpcEndpointResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
func
}


func TestAccVPCRoute_ipv6ToVPCEndpoint(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var route ec2.Route
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route.test"
	vpcEndpointResourceName := "aws_vpc_endpoint.test"
	destinationIpv6Cidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckELBv2GatewayLoadBalancer(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID, "elasticloadbalancing"),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", destinationIpv6Cidr),
	resource.TestCheckResourceAttr(resourceName, "destination_prefix_list_id", ""),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_endpoint_id", vpcEndpointResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
funcrtStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}

// https://github.com/hashicorp/terraform-provider-aws/issues/11455.

func TestAccVPCRoute_localRoute(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	var vpc ec2.Vpc
	resourceName := "aws_route.test"
	rtResourceName := "aws_route_table.test"
	vpcResourceName := "aws_vpc.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
funcig: testAccVPCRouteConfig_ipv4NoRoute(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, rtResourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
),
	},
funcig:PCRouteConfig_ipv4Local(rName),
ResourceName: resourceName,
ImportState:  true,
ImportStateId
func: 
func(rt *ec2.RouteTable, v *ec2.Vpc) resource.ImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
return fmt.Sprintf("%s_%s", aws.StringValue(rt.RouteTableId), aws.StringValue(v.CidrBlock)), nil
	}
}(&routeTable, &vpc),
// Don't verify the state as the local route isn't actually in the pre-import state.
// Just running ImportState verifies that we can import a local route.
func
},
	})
}

// https://github.com/hashicorp/terraform-provider-aws/issues/21350.

func TestAccVPCRoute_localRouteUpdate(t *testing.T) {
func routeTable ec2.RouteTable
	var vpc ec2.Vpc
	resourceName := "aws_route.test"
	rtResourceName := "aws_route_table.test"
	vpcResourceName := "aws_vpc.test"
	eniResourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_ipv4NoRoute(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, rtResourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
),
	},
	{
funcurceName: resourceName,
funcrtStateId
func: 
func(rt *ec2.RouteTable, v *ec2.Vpc) resource.ImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
return fmt.Sprintf("%s_%s", aws.StringValue(rt.RouteTableId), aws.StringValue(v.CidrBlock)), nil
	}
funcrtStatePersist: true,
// Don't verify the state as the local route isn't actually in the pre-import state.
// Just running ImportState verifies that we can import a local route.
ImportStateVerify: false,
	},
	{
Config: testAccVPCRouteConfig_ipv4LocalToNetworkInterface(rName),
Check: resource.ComposeAggregateTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, rtResourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eniResourceName, "id"),
func
	{
Config: testAccVPCRouteConfig_ipv4LocalRestore(rName),
Check: resource.ComposeAggregateTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, rtResourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
func
	},
},
	})
}


func TestAccVPCRoute_prefixListToInternetGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	igwResourceName := "aws_internet_gateway.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_prefixListInternetGateway(rName),
Check: resource.ComposeTestCheck
functAccCheckRouteExists(ctx, resourceName, &route),
funcource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", igwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCRoute_prefixListToVPNGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	vgwResourceName := "aws_vpn_gateway.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
funcig: testAccVPCRouteConfig_prefixListVPNGateway(rName),
Check: resource.ComposeTestCheck
functAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", vgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
funcource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
func
	})
}


func TestAccVPCRoute_prefixListToInstance(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
functanceResourceName := "aws_instance.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCRouteConfig_prefixListInstance(rName),
func(
funcource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "instance_id", instanceResourceName, "id"),
	acctest.CheckResourceAttrAccountID(resourceName, "instance_owner_id"),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", instanceResourceName, "primary_network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
func
}


func TestAccVPCRoute_PrefixListToNetworkInterface_unattached(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	eniResourceName := "aws_network_interface.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_prefixListNetworkInterfaceUnattached(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
funcource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eniResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateBlackhole),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_PrefixListToNetworkInterface_attached(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_route.test"
functanceResourceName := "aws_instance.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_prefixListNetworkInterfaceAttached(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
funcource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "instance_id", instanceResourceName, "id"),
	acctest.CheckResourceAttrAccountID(resourceName, "instance_owner_id"),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface_id", eniResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_prefixListToVPCPeeringConnection(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
func
funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_prefixListPeeringConnection(rName),
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
funcource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_peering_connection_id", pcxResourceName, "id"),
func
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_prefixListToNatGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	ngwResourceName := "aws_nat_gateway.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_prefixListNATGateway(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "nat_gateway_id", ngwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
funcource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_prefixListToTransitGateway(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var route ec2.Route
	resourceName := "aws_route.test"
	tgwResourceName := "aws_ec2_transit_gateway.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckRouteDestroy(ctx),
func
Config: testAccVPCRouteConfig_prefixListTransitGateway(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
funcource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", tgwResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
funcrtStateVerify: true,
	},
},
	})
}


func TestAccVPCRoute_prefixListToCarrierGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	cgwResourceName := "aws_ec2_carrier_gateway.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	testAccPreCheckManagedPrefixList(ctx, t)
	testAccPreCheckWavelengthZoneAvailable(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttrPair(resourceName, "carrier_gateway_id", cgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
func
},
	})
}


func TestAccVPCRoute_prefixListToLocalGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	localGatewayDataSourceName := "data.aws_ec2_local_gateway.first"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	testAccPreCheckManagedPrefixList(ctx, t)
	acctest.PreCheckOutpostsOutposts(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCRouteConfig_prefixListLocalGateway(rName),
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
	resource.TestCheckResourceAttr(resourceName, "carrier_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "egress_only_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "local_gateway_id", localGatewayDataSourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_id", ""),
	resource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
func
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
func
	})
}


func TestAccVPCRoute_prefixListToEgressOnlyInternetGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var route ec2.Route
	resourceName := "aws_route.test"
	eoigwResourceName := "aws_egress_only_internet_gateway.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteConfig_prefixListEgressOnlyInternetGateway(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteExists(ctx, resourceName, &route),
funcource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
funcource.TestCheckResourceAttr(resourceName, "destination_ipv6_cidr_block", ""),
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", plResourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "egress_only_gateway_id", eoigwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_owner_id", ""),
	resource.TestCheckResourceAttr(resourceName, "local_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "nat_gateway_id", ""),
funcource.TestCheckResourceAttr(resourceName, "origin", ec2.RouteOriginCreateRoute),
	resource.TestCheckResourceAttr(resourceName, "state", ec2.RouteStateActive),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_id", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_peering_connection_id", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteImportStateId
func(resourceName),
ImportStateVerify: true,
func
	})
}


func testAccCheckRouteExists(ctx context.Context, n string, v *ec2.Route) resource.TestCheck
func {
	return 
funcok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

var route *ec2.Route
var err error
if v := rs.Primary.Attributes["destination_cidr_block"]; v != "" {
	route, err = tfec2.FindRouteByIPv4Destination(ctx, conn, rs.Primary.Attributes["route_table_id"], v)
} else if v := rs.Primary.Attributes["destination_ipv6_cidr_block"]; v != "" {
	route, err = tfec2.FindRouteByIPv6Destination(ctx, conn, rs.Primary.Attributes["route_table_id"], v)
} else if v := rs.Primary.Attributes["destination_prefix_list_id"]; v != "" {
	route, err = tfec2.FindRouteByPrefixListIDDestination(ctx, conn, rs.Primary.Attributes["route_table_id"], v)
}

if err != nil {
	return err
}

func
func
}


func testAccCheckRouteDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
funcrs.Type != "aws_route" {
continue
	}

	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	var err error
	if v := rs.Primary.Attributes["destination_cidr_block"]; v != "" {
_, err = tfec2.FindRouteByIPv4Destination(ctx, conn, rs.Primary.Attributes["route_table_id"], v)
	} else if v := rs.Primary.Attributes["destination_ipv6_cidr_block"]; v != "" {
funclse if v := rs.Primary.Attributes["destination_prefix_list_id"]; v != "" {
_, err = tfec2.FindRouteByPrefixListIDDestination(ctx, conn, rs.Primary.Attributes["route_table_id"], v)
	}

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

func

return nil
	}
}


func testAccRouteImportStateId
func(resourceName string) resource.ImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return "", fmt.Errorf("not found: %s", resourceName)
}

destination := rs.Primary.Attributes["destination_cidr_block"]
if v, ok := rs.Primary.Attributes["destination_ipv6_cidr_block"]; ok && v != "" {
	destination = v
}
if v, ok := rs.Primary.Attributes["destination_prefix_list_id"]; ok && v != "" {
	destination = v
}

func
func

func testAccVPCRouteConfig_ipv4InternetGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
func
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
funcurce "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  gateway_idaws_internet_gateway.test.id
funcName, destinationCidr)
}


func testAccVPCRouteConfig_ipv6InternetGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_egress_only_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id
funcgs = {
func
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
func

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  gateway_idws_internet_gateway.test.id
}
`, rName, destinationCidr)
}

func testAccVPCRouteConfig_ipv6NetworkInterfaceUnattached(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

func %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]
  ipv6_cidr_blockidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id
funcgs = {
func
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  network_interface_idetwork_interface.test.id
}
func
func
func testAccVPCRouteConfig_ipv6Instance(rName, destinationCidr string) string {
funcAccLatestAmazonNatInstanceAMIConfig(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]
  ipv6_cidr_blockidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-nat-instance.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_idet.test.id

  ipv6_address_count = 1

  tags = {
me = %[1]q
func
funcurce "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id
funcgs = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  network_interface_idnstance.test.primary_network_interface_id
}
`, rName, destinationCidr))
}


func testAccVPCRouteConfig_ipv6PeeringConnection(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "target" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}
funcurce "aws_vpc_peering_connection" "test" {
funcer_vpc_id = aws_vpc.target.id
func
  tags = {
func
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  vpc_peering_connection_idws_vpc_peering_connection.test.id
}
`, rName, destinationCidr)
}

func testAccVPCRouteConfig_ipv6EgressOnlyInternetGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_egress_only_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  egress_only_gateway_idess_only_internet_gateway.test.id
}
`, rName, destinationCidr)
}

func testAccVPCRouteConfig_endpoint(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = "10.3.0.0/16"
  gateway_idaws_internet_gateway.test.id

  # Forcing endpoint to create before route - without this the crash is a race.
  depends_on = [aws_vpc_endpoint.test]
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
  vpc_id = aws_vpc.test.id
  service_name"com.amazonaws.${data.aws_region.current.name}.s3"
  route_table_ids = [aws_route_table.test.id]
}
`, rName)
}
func
func testAccVPCRouteConfig_ipv4TransitGateway(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptInDefaultExclude(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids= [aws_subnet.test.id]
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  destination_cidr_block = %[2]q
funcansit_gateway_idtransit_gateway_vpc_attachment.test.transit_gateway_id
}
`, rName, destinationCidr))
}


func testAccVPCRouteConfig_ipv6TransitGateway(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids= [aws_subnet.test.id]
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  destination_ipv6_cidr_block = %[2]q
funcansit_gateway_id = aws_ec2_transit_gateway_vpc_attachment.test.transit_gateway_id
}
`, rName, destinationCidr))
}


func testAccVPCRouteConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr string, ipv6Route bool) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

locals {
  ipv6%[4]t
  destination
  destination_ipv6 = %[3]q
}

resource "aws_route" "test" {
  route_table_id = aws_route_table.test.id
  gateway_idrnet_gateway.test.id

  destination_cidr_blockpv6 ? null : local.destination
  destination_ipv6_cidr_block = local.ipv6 ? local.destination_ipv6 : null
}
`, rName, destinationCidr, destinationIpv6Cidr, ipv6Route)
func

func testAccVPCRouteConfig_ipv4Instance(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
testAccLatestAmazonNatInstanceAMIConfig(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-nat-instance.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_idet.test.id

  tags = {
me = %[1]q
  }
}

funcc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  network_interface_idws_instance.test.primary_network_interface_id
}
`, rName, destinationCidr))
}


func testAccVPCRouteConfig_ipv4NetworkInterfaceUnattached(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

funcc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  network_interface_idws_network_interface.test.id
}
`, rName, destinationCidr))
}


func testAccVPCRouteConfig_resourceIPv4LocalGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
data "aws_ec2_local_gateways" "all" {}

data "aws_ec2_local_gateway" "first" {
  id = tolist(data.aws_ec2_local_gateways.all.ids)[0]
}

data "aws_ec2_local_gateway_route_tables" "all" {}

data "aws_ec2_local_gateway_route_table" "first" {
  local_gateway_route_table_id = tolist(data.aws_ec2_local_gateway_route_tables.all.ids)[0]
}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_local_gateway_route_table_vpc_association" "example" {
  local_gateway_route_table_id = data.aws_ec2_local_gateway_route_table.first.id
  vpc_idpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  depends_on = [aws_ec2_local_gateway_route_table_vpc_association.example]
func
resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  local_gateway_idws_ec2_local_gateway.first.id
}
`, rName, destinationCidr)
}


func testAccVPCRouteConfig_resourceIPv6LocalGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
data "aws_ec2_local_gateways" "all" {}

data "aws_ec2_local_gateway" "first" {
  id = tolist(data.aws_ec2_local_gateways.all.ids)[0]
}

data "aws_ec2_local_gateway_route_tables" "all" {}

data "aws_ec2_local_gateway_route_table" "first" {
  local_gateway_route_table_id = tolist(data.aws_ec2_local_gateway_route_tables.all.ids)[0]
}

resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_local_gateway_route_table_vpc_association" "example" {
  local_gateway_route_table_id = data.aws_ec2_local_gateway_route_table.first.id
  vpc_idpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  depends_on = [aws_ec2_local_gateway_route_table_vpc_association.example]
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  local_gateway_idata.aws_ec2_local_gateway.first.id
funcName, destinationCidr)
}


func testAccVPCRouteConfig_ipv4NetworkInterfaceAttached(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
testAccLatestAmazonNatInstanceAMIConfig(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-nat-instance.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type

  network_interface {
vice_index= 0
twork_interface_id = aws_network_interface.test.id
  }
funcgs = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  network_interface_idws_network_interface.test.id

  # Wait for the ENI attachment.
  depends_on = [aws_instance.test]
}
`, rName, destinationCidr))
}


func testAccVPCRouteConfig_ipv4NetworkInterfaceTwoAttachments(rName, destinationCidr, targetResourceName string) string {
	return acctest.ConfigCompose(
testAccLatestAmazonNatInstanceAMIConfig(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test1" {
  subnet_id = aws_subnet.test.id
funcgs = {
me = %[1]q
  }
}

resource "aws_network_interface" "test2" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-nat-instance.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type

  network_interface {
vice_index= 0
twork_interface_id = aws_network_interface.test1.id
  }

  network_interface {
vice_index= 1
twork_interface_id = aws_network_interface.test2.id
  }

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  network_interface_id[3]s.id

  # Wait for the ENI attachment.
  depends_on = [aws_instance.test]
funcName, destinationCidr, targetResourceName))
}


func testAccVPCRouteConfig_ipv4PeeringConnection(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "target" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_peering_connection" "test" {
  vpc_id.test.id
  peer_vpc_id = aws_vpc.target.id
  auto_accept = true

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idws_route_table.test.id
  destination_cidr_block%[2]q
  vpc_peering_connection_id = aws_vpc_peering_connection.test.id
}
`, rName, destinationCidr)
}


func testAccVPCRouteConfig_ipv4NATGateway(rName, destinationCidr string) string {
funcurce "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_idtest.id

  map_public_ip_on_launch = true

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_eip" "test" {
  domain = "vpc"

  tags = {
me = %[1]q
  }
}

resource "aws_nat_gateway" "test" {
  allocation_id = aws_eip.test.id
  subnet_idet.test.id

  tags = {
me = %[1]q
  }

  depends_on = [aws_internet_gateway.test]
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
func

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  nat_gateway_id= aws_nat_gateway.test.id
}
`, rName, destinationCidr)
}


func testAccVPCRouteConfig_ipv6NATGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  vpc_id  = aws_vpc.test.id
  cidr_block1.0/24"
  ipv6_cidr_block  = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)
  assign_ipv6_address_on_creation = true

  enable_resource_name_dns_aaaa_record_on_launch = true

  tags = {
me = %[1]q
  }
}

resource "aws_nat_gateway" "test" {
  connectivity_type = "private"
  subnet_id= aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  nat_gateway_idgateway.test.id
}
`, rName, destinationCidr)
}


func testAccVPCRouteConfig_ipv4VPNGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

func %[1]q
  }
}

resource "aws_vpn_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  gateway_idaws_vpn_gateway.test.id
}
`, rName, destinationCidr)
}


func testAccVPCRouteConfig_ipv6VPNGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  gateway_idws_vpn_gateway.test.id
}
`, rName, destinationCidr)
}


func testAccVPCRouteConfig_resourceIPv4Endpoint(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
data "aws_caller_identity" "current" {}

data "aws_iam_session_context" "current" {
  arn = data.aws_caller_identity.current.arn
}

resource "aws_vpc" "test" {
  cidr_block = "10.10.10.0/25"

  tags = {
func
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_blockubnet(aws_vpc.test.cidr_block, 2, 0)
  vpc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_lb" "test" {
  load_balancer_type = "gateway"
  name= %[1]q

  subnet_mapping {
bnet_id = aws_subnet.test.id
  }
}

resource "aws_vpc_endpoint_service" "test" {
  acceptance_required
  allowed_principals= [data.aws_iam_session_context.current.issuer_arn]
  gateway_load_balancer_arns = [aws_lb.test.arn]

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_endpoint" "test" {
  service_name_endpoint_service.test.service_name
  subnet_idssubnet.test.id]
  vpc_endpoint_type = aws_vpc_endpoint_service.test.service_type
  vpc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id
funcgs = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q
  vpc_endpoint_idpc_endpoint.test.id
}
`, rName, destinationCidr))
}


func testAccVPCRouteConfig_resourceIPv6Endpoint(rName, destinationIpv6Cidr string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
data "aws_caller_identity" "current" {}

data "aws_iam_session_context" "current" {
  arn = data.aws_caller_identity.current.arn
}

resource "aws_vpc" "test" {
  cidr_block = "10.10.10.0/25"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_blockubnet(aws_vpc.test.cidr_block, 2, 0)
  vpc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_lb" "test" {
  load_balancer_type = "gateway"
  name= %[1]q

  subnet_mapping {
bnet_id = aws_subnet.test.id
  }
}

resource "aws_vpc_endpoint_service" "test" {
  acceptance_required
  allowed_principals= [data.aws_iam_session_context.current.issuer_arn]
  gateway_load_balancer_arns = [aws_lb.test.arn]

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_endpoint" "test" {
  service_name_endpoint_service.test.service_name
  subnet_idssubnet.test.id]
  vpc_endpoint_type = aws_vpc_endpoint_service.test.service_type
func
  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q
  vpc_endpoint_idaws_vpc_endpoint.test.id
}
`, rName, destinationIpv6Cidr))
}


func testAccVPCRouteConfig_ipv4FlexiTarget(rName, destinationCidr, targetAttribute, targetValue string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptInDefaultExclude(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
locals {
  target_attr  = %[3]q
  target_value = %[4]s.id
}

resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

funcc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  map_public_ip_on_launch = true

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_idet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
}

funcbnet_ids= [aws_subnet.test.id]
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "target" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_peering_connection" "test" {
  vpc_id.test.id
  peer_vpc_id = aws_vpc.target.id
  auto_accept = true

  tags = {
me = %[1]q
  }
}

resource "aws_eip" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_nat_gateway" "test" {
  allocation_id = aws_eip.test.id
  subnet_idet.test.id

  tags = {
me = %[1]q
  }

  depends_on = [aws_internet_gateway.test]
}

data "aws_caller_identity" "current" {}

data "aws_iam_session_context" "current" {
  arn = data.aws_caller_identity.current.arn
}

resource "aws_lb" "test" {
  load_balancer_type = "gateway"
  name= %[1]q

  subnet_mapping {
bnet_id = aws_subnet.test.id
  }
}

resource "aws_vpc_endpoint_service" "test" {
  acceptance_required
  allowed_principals= [data.aws_iam_session_context.current.issuer_arn]
  gateway_load_balancer_arns = [aws_lb.test.arn]

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_endpoint" "test" {
  service_name_endpoint_service.test.service_name
  subnet_idssubnet.test.id]
  vpc_endpoint_type = aws_vpc_endpoint_service.test.service_type
  vpc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = %[2]q

  egress_only_gateway_id(local.target_attr == "egress_only_gateway_id") ? local.target_value : null
  gateway_id = (local.target_attr == "gateway_id") ? local.target_value : null
  local_gateway_id = (local.target_attr == "local_gateway_id") ? local.target_value : null
  nat_gateway_idlocal.target_attr == "nat_gateway_id") ? local.target_value : null
  network_interface_idtarget_attr == "network_interface_id") ? local.target_value : null
  transit_gateway_idl.target_attr == "transit_gateway_id") ? local.target_value : null
  vpc_endpoint_id  = (local.target_attr == "vpc_endpoint_id") ? local.target_value : null
  vpc_peering_connection_id = (local.target_attr == "vpc_peering_connection_id") ? local.target_value : null
}
`, rName, destinationCidr, targetAttribute, targetValue))
func

func testAccVPCRouteConfig_ipv6FlexiTarget(rName, destinationCidr, targetAttribute, targetValue string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
locals {
  target_attr  = %[3]q
  target_value = %[4]s.id
}

resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]
  ipv6_cidr_blockidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)

  map_public_ip_on_launch = true

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_idet.test.id

  ipv6_address_count = 1

  tags = {
me = %[1]q
  }
}

resource "aws_egress_only_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
func
}

resource "aws_vpc" "target" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_peering_connection" "test" {
  vpc_id.test.id
  peer_vpc_id = aws_vpc.target.id
  auto_accept = true

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_ide_table.test.id
  destination_ipv6_cidr_block = %[2]q

  egress_only_gateway_id(local.target_attr == "egress_only_gateway_id") ? local.target_value : null
  gateway_id = (local.target_attr == "gateway_id") ? local.target_value : null
  local_gateway_id = (local.target_attr == "local_gateway_id") ? local.target_value : null
  nat_gateway_idlocal.target_attr == "nat_gateway_id") ? local.target_value : null
  network_interface_idtarget_attr == "network_interface_id") ? local.target_value : null
  transit_gateway_idl.target_attr == "transit_gateway_id") ? local.target_value : null
  vpc_endpoint_id  = (local.target_attr == "vpc_endpoint_id") ? local.target_value : null
  vpc_peering_connection_id = (local.target_attr == "vpc_peering_connection_id") ? local.target_value : null
}
`, rName, destinationCidr, targetAttribute, targetValue))
}


func testAccVPCRouteConfig_ipv4NoRoute(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCRouteConfig_ipv4Local(rName string) string {
	return acctest.ConfigCompose(testAccVPCRouteConfig_ipv4NoRoute(rName), `
resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = aws_vpc.test.cidr_block
  gateway_id"local"
}
`)
}


func testAccVPCRouteConfig_ipv4LocalToNetworkInterface(rName string) string {
	return acctest.ConfigCompose(testAccVPCRouteConfig_ipv4NoRoute(rName), fmt.Sprintf(`
resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = aws_vpc.test.cidr_block
  network_interface_idws_network_interface.test.id
}

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_idtest.id

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCRouteConfig_ipv4LocalRestore(rName string) string {
	return acctest.ConfigCompose(testAccVPCRouteConfig_ipv4NoRoute(rName), fmt.Sprintf(`
resource "aws_route" "test" {
  route_table_id= aws_route_table.test.id
  destination_cidr_block = aws_vpc.test.cidr_block
  gateway_id"local"
}

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_idtest.id

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCRouteConfig_prefixListInternetGateway(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  gateway_id  = aws_internet_gateway.test.id
}
`, rName)
}


func testAccVPCRouteConfig_prefixListVPNGateway(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_vpn_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  gateway_id  = aws_vpn_gateway.test.id
}
`, rName)
}


func testAccVPCRouteConfig_prefixListInstance(rName string) string {
	return acctest.ConfigCompose(
testAccLatestAmazonNatInstanceAMIConfig(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-nat-instance.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_idet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  network_interface_idstance.test.primary_network_interface_id
}
`, rName))
}


func testAccVPCRouteConfig_prefixListNetworkInterfaceUnattached(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
func

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  network_interface_idtwork_interface.test.id
}
func


func testAccVPCRouteConfig_prefixListNetworkInterfaceAttached(rName string) string {
	return acctest.ConfigCompose(
testAccLatestAmazonNatInstanceAMIConfig(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"
funcgs = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-nat-instance.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type

  network_interface {
funck_interface_id = aws_network_interface.test.id
  }

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  network_interface_idtwork_interface.test.id

  # Wait for the ENI attachment.
func
`, rName))
}


func testAccVPCRouteConfig_prefixListPeeringConnection(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "target" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_peering_connection" "test" {
  vpc_id.test.id
  peer_vpc_id = aws_vpc.target.id
  auto_accept = true

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id
funcgs = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  vpc_peering_connection_id  = aws_vpc_peering_connection.test.id
}
`, rName)
}


func testAccVPCRouteConfig_prefixListNATGateway(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_idtest.id

  map_public_ip_on_launch = true

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
func
resource "aws_eip" "test" {
  domain = "vpc"

  tags = {
me = %[1]q
  }
}

resource "aws_nat_gateway" "test" {
  allocation_id = aws_eip.test.id
  subnet_idet.test.id

  tags = {
me = %[1]q
  }

  depends_on = [aws_internet_gateway.test]
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  nat_gateway_idaws_nat_gateway.test.id
}
`, rName)
}


func testAccVPCRouteConfig_prefixListTransitGateway(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptInDefaultExclude(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
funcc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids= [aws_subnet.test.id]
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  transit_gateway_id= aws_ec2_transit_gateway_vpc_attachment.test.transit_gateway_id
}
`, rName))
}


func testAccVPCRouteConfig_prefixListCarrierGateway(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
func

resource "aws_ec2_carrier_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  carrier_gateway_id= aws_ec2_carrier_gateway.test.id
}
`, rName)
}


func testAccVPCRouteConfig_prefixListLocalGateway(rName string) string {
	return fmt.Sprintf(`
data "aws_ec2_local_gateways" "all" {}

data "aws_ec2_local_gateway" "first" {
  id = tolist(data.aws_ec2_local_gateways.all.ids)[0]
}

data "aws_ec2_local_gateway_route_tables" "all" {}

data "aws_ec2_local_gateway_route_table" "first" {
  local_gateway_route_table_id = tolist(data.aws_ec2_local_gateway_route_tables.all.ids)[0]
}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_local_gateway_route_table_vpc_association" "example" {
  local_gateway_route_table_id = data.aws_ec2_local_gateway_route_table.first.id
  vpc_idpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id
funcgs = {
me = %[1]q
  }

  depends_on = [aws_ec2_local_gateway_route_table_vpc_association.example]
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
  destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
  local_gateway_id  = data.aws_ec2_local_gateway.first.id
}
`, rName)
}


func testAccVPCRouteConfig_prefixListEgressOnlyInternetGateway(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_egress_only_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv6"
  max_entries1
  name  = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  route_table_idaws_route_table.test.id
funcress_only_gateway_idss_only_internet_gateway.test.id
}
`, rName)
}


func testAccVPCRouteConfig_ipv4CarrierGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_carrier_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route" "test" {
  destination_cidr_block = %[2]q
  route_table_id= aws_route_table.test.id
  carrier_gateway_idcarrier_gateway.test.id
}
`, rName, destinationCidr)
}
funcfuncfuncfuncfunc