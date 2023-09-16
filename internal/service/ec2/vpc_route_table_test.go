// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
functest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
Config: testAccVPCRouteTableConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceRouteTable(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccVPCRouteTable_Disappears_subnetAssociation(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_subnetAssociation(rName),
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceRouteTable(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
func

func TestAccVPCRouteTable_ipv4ToInternetGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	igwResourceName := "aws_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr1 := "10.2.0.0/16"
	destinationCidr2 := "10.3.0.0/16"
	destinationCidr3 := "10.4.0.0/16"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_ipv4InternetGateway(rName, destinationCidr1, destinationCidr2),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
functest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "gateway_id", igwResourceName, "id"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "gateway_id", igwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
func
	{
Config: testAccVPCRouteTableConfig_ipv4InternetGateway(rName, destinationCidr2, destinationCidr3),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "gateway_id", igwResourceName, "id"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr3, "gateway_id", igwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_ipv4ToInstance(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	instanceResourceName := "aws_instance.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_ipv4Instance(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
functAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "network_interface_id", instanceResourceName, "primary_network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccVPCRouteTable_ipv6ToEgressOnlyInternetGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	eoigwResourceName := "aws_egress_only_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr, "egress_only_gateway_id", eoigwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
funcig:tAccVPCRouteTableConfig_ipv6EgressOnlyInternetGateway(rName, "::0/0"),
PlanOnly: true,
	},
},
	})
}


func TestAccVPCRouteTable_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
funcig: testAccVPCRouteTableConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}


func TestAccVPCRouteTable_requireRouteDestination(t *testing.T) {
functesting.Short() {
t.Skip("skipping long-running test in short mode")
	}

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config:CRouteTableConfig_noDestination(rName),
ExpectError: regexache.MustCompile("creating route: one of `cidr_block"),
	},
},
	})
}


func := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
funcig:CRouteTableConfig_noTarget(rName),
ExpectError: regexache.MustCompile(`creating route: one of .*\begress_only_gateway_id\b`),
	},
},
	})
}


func TestAccVPCRouteTable_Route_mode(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	igwResourceName := "aws_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
functinationCidr2 := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_ipv4InternetGateway(rName, destinationCidr1, destinationCidr2),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "gateway_id", igwResourceName, "id"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "gateway_id", igwResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCRouteTableConfig_modeNoBlocks(rName),
Check: resource.ComposeTestCheck
functAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "gateway_id", igwResourceName, "id"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "gateway_id", igwResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCRouteTableConfig_modeZeroed(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_ipv4ToTransitGateway(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	tgwResourceName := "aws_ec2_transit_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_ipv4TransitGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "transit_gateway_id", tgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_ipv4ToVPCEndpoint(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
func
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	vpceResourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "0.0.0.0/0"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckELBv2GatewayLoadBalancer(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID, "elasticloadbalancing"),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_ipv4EndpointID(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "vpc_endpoint_id", vpceResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_ipv4ToCarrierGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	cgwResourceName := "aws_ec2_carrier_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "0.0.0.0/0"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckWavelengthZoneAvailable(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "carrier_gateway_id", cgwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCRouteTable_ipv4ToLocalGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	lgwDataSourceName := "data.aws_ec2_local_gateway.first"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "0.0.0.0/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_ipv4LocalGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
functAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "local_gateway_id", lgwDataSourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_ipv4ToVPCPeeringConnection(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_route_table.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_ipv4PeeringConnection(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
funcource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "vpc_peering_connection_id", pcxResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_vgwRoutePropagation(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	vgwResourceName1 := "aws_vpn_gateway.test1"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "1"),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "propagating_vgws.*", vgwResourceName1, "id"),
funcource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
Config: testAccVPCRouteTableConfig_vgwPropagation(rName, vgwResourceName2),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "1"),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "propagating_vgws.*", vgwResourceName2, "id"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccVPCRouteTable_conditionalCIDRBlock(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"
	destinationIpv6Cidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr, false),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "gateway_id", igwResourceName, "id"),
),
	},
	{
Config: testAccVPCRouteTableConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_ipv4ToNatGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	ngwResourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_ipv4NATGateway(rName, destinationCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
funcource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "nat_gateway_id", ngwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	eniResourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr, "network_interface_id", eniResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccVPCRouteTable_IPv4ToNetworkInterfaces_unattached(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	eni1ResourceName := "aws_network_interface.test1"
	eni2ResourceName := "aws_network_interface.test2"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr1 := "10.2.0.0/16"
	destinationCidr2 := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_basic(rName),
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
),
	},
funcig: testAccVPCRouteTableConfig_ipv4TwoNetworkInterfacesUnattached(rName, destinationCidr1, destinationCidr2),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
functAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "network_interface_id", eni1ResourceName, "id"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "network_interface_id", eni2ResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCRouteTableConfig_ipv4TwoNetworkInterfacesUnattached(rName, destinationCidr2, destinationCidr1),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "network_interface_id", eni1ResourceName, "id"),
funcource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
Config: testAccVPCRouteTableConfig_modeZeroed(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
funcource.TestCheckResourceAttr(resourceName, "route.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
},
	})
}
func
func TestAccVPCRouteTable_vpcMultipleCIDRs(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
Config: testAccVPCRouteTableConfig_multipleCIDRs(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccVPCRouteTable_gatewayVPCEndpoint(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	var vpce ec2.VpcEndpoint
	resourceName := "aws_route_table.test"
	vpceResourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
Config: testAccVPCRouteTableConfig_gatewayEndpoint(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckVPCEndpointExists(ctx, vpceResourceName, &vpce),
	testAccCheckRouteTableWaitForVPCEndpointRoute(ctx, &routeTable, &vpce),
	// Refresh the route table once the VPC endpoint route is present.
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccVPCRouteTable_multipleRoutes(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	eoigwResourceName := "aws_egress_only_internet_gateway.test"
	igwResourceName := "aws_internet_gateway.test"
funcResourceName := "aws_vpc_peering_connection.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr1 := "10.2.0.0/16"
	destinationCidr2 := "10.3.0.0/16"
	destinationCidr3 := "10.4.0.0/16"
	destinationCidr4 := "2001:db8::/122"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_multiples(rName,
	"cidr_block", destinationCidr1, "gateway_id", fmt.Sprintf(`%s.%s`, igwResourceName, "id"),
	"cidr_block", destinationCidr2, "network_interface_id", fmt.Sprintf(`%s.%s`, instanceResourceName, "primary_network_interface_id"),
	"ipv6_cidr_block", destinationCidr4, "egress_only_gateway_id", fmt.Sprintf(`%s.%s`, eoigwResourceName, "id")),
Check: resource.ComposeTestCheck
func(
functAccCheckRouteTableNumberOfRoutes(&routeTable, 5),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "3"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "gateway_id", igwResourceName, "id"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "network_interface_id", instanceResourceName, "primary_network_interface_id"),
	testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr4, "egress_only_gateway_id", eoigwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
func
	{
Config: testAccVPCRouteTableConfig_multiples(rName,
	"cidr_block", destinationCidr1, "vpc_peering_connection_id", fmt.Sprintf(`%s.%s`, pcxResourceName, "id"),
	"cidr_block", destinationCidr3, "network_interface_id", fmt.Sprintf(`%s.%s`, instanceResourceName, "primary_network_interface_id"),
	"ipv6_cidr_block", destinationCidr4, "egress_only_gateway_id", fmt.Sprintf(`%s.%s`, eoigwResourceName, "id")),
Check: resource.ComposeTestCheck
func(
functAccCheckRouteTableNumberOfRoutes(&routeTable, 5),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "3"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "vpc_peering_connection_id", pcxResourceName, "id"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr3, "network_interface_id", instanceResourceName, "primary_network_interface_id"),
	testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr4, "egress_only_gateway_id", eoigwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
Config: testAccVPCRouteTableConfig_multiples(rName,
	"ipv6_cidr_block", destinationCidr4, "vpc_peering_connection_id", fmt.Sprintf(`%s.%s`, pcxResourceName, "id"),
	"cidr_block", destinationCidr3, "gateway_id", fmt.Sprintf(`%s.%s`, igwResourceName, "id"),
	"cidr_block", destinationCidr2, "network_interface_id", fmt.Sprintf(`%s.%s`, instanceResourceName, "primary_network_interface_id")),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 5),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "3"),
functAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr3, "gateway_id", igwResourceName, "id"),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "network_interface_id", instanceResourceName, "primary_network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	igwResourceName := "aws_internet_gateway.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableConfig_prefixListInternetGateway(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`route-table/.+$`)),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
	testAccCheckRouteTablePrefixListRoute(resourceName, plResourceName, "gateway_id", igwResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccVPCRouteTable_localRoute(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	var vpc ec2.Vpc
	resourceName := "aws_route_table.test"
	vpcResourceName := "aws_vpc.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
Config: testAccVPCRouteTableConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
),
	},
	{
Config:tAccVPCRouteTableConfig_ipv4Local(rName),
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTable_localRouteAdoptUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	var vpc ec2.Vpc
funcResourceName := "aws_vpc.test"
	eniResourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vpcCIDR := "10.1.0.0/16"
	localGatewayCIDR := "10.1.0.0/16"
	localGatewayCIDRBad := "10.2.0.0/16"
	subnetCIDR := "10.1.1.0/24"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config:CRouteTableConfig_ipv4NetworkInterfaceToLocal(rName, vpcCIDR, localGatewayCIDRBad, subnetCIDR),
ExpectError: regexache.MustCompile("must exist to be adopted"),
	},
funcig: testAccVPCRouteTableConfig_ipv4NetworkInterfaceToLocal(rName, vpcCIDR, localGatewayCIDR, subnetCIDR),
Check: resource.ComposeAggregateTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "route.*", map[string]string{
"gateway_id": "local",
"cidr_block": localGatewayCIDR,
	}),
),
	},
	{
Config: testAccVPCRouteTableConfig_ipv4LocalNetworkInterface(rName, vpcCIDR, subnetCIDR),
Check: resource.ComposeAggregateTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", vpcCIDR, "network_interface_id", eniResourceName, "id"),
),
	},
funcig: testAccVPCRouteTableConfig_ipv4NetworkInterfaceToLocal(rName, vpcCIDR, localGatewayCIDR, subnetCIDR),
Check: resource.ComposeAggregateTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "route.*", map[string]string{
"gateway_id": "local",
"cidr_block": localGatewayCIDR,
	}),
func
},
	})
}


func TestAccVPCRouteTable_localRouteImportUpdate(t *testing.T) {
	ctx := acctest.Context(t)
func vpc ec2.Vpc
	resourceName := "aws_route_table.test"
	rteResourceName := "aws_route.test"
	vpcResourceName := "aws_vpc.test"
	eniResourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vpcCIDR := "10.1.0.0/16"
	localGatewayCIDR := "10.1.0.0/16"
	subnetCIDR := "10.1.1.0/24"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
funccould figure anyway) to use aws_route_table to import a local
	// route and then persist it to the next step since the route is
	// inline rather than a separate resource. Instead, it uses
	// aws_route config rather than aws_route_table w/ inline routes
	// for steps 1-3 and then does slight of hand, switching
	// to aws_route_table to finish the test.
	{
Config: testAccVPCRouteConfig_ipv4NoRoute(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
),
	},
funcig:PCRouteConfig_ipv4Local(rName),
ResourceName: rteResourceName,
ImportState:true,
ImportStateId
func: 
func(rt *ec2.RouteTable, v *ec2.Vpc) resource.ImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
return fmt.Sprintf("%s_%s", aws.StringValue(rt.RouteTableId), aws.StringValue(v.CidrBlock)), nil
	}
}(&routeTable, &vpc),
funcon't verify the state as the local route isn't actually in the pre-import state.
// Just running ImportState verifies that we can import a local route.
ImportStateVerify: false,
	},
	{
Config: testAccVPCRouteConfig_ipv4LocalToNetworkInterface(rName),
Check: resource.ComposeAggregateTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	resource.TestCheckResourceAttr(rteResourceName, "gateway_id", ""),
	resource.TestCheckResourceAttrPair(rteResourceName, "network_interface_id", eniResourceName, "id"),
func
	{
Config: testAccVPCRouteTableConfig_ipv4LocalNetworkInterface(rName, vpcCIDR, subnetCIDR),
Check: resource.ComposeAggregateTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", vpcCIDR, "network_interface_id", eniResourceName, "id"),
),
func
Config: testAccVPCRouteTableConfig_ipv4NetworkInterfaceToLocal(rName, vpcCIDR, localGatewayCIDR, subnetCIDR),
Check: resource.ComposeAggregateTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "route.*", map[string]string{
"gateway_id": "local",
"cidr_block": localGatewayCIDR,
	}),
),
	},
	{
Config: testAccVPCRouteTableConfig_ipv4LocalNetworkInterface(rName, vpcCIDR, subnetCIDR),
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &vpc),
	testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
	testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
	testAccCheckRouteTableRoute(resourceName, "cidr_block", vpcCIDR, "network_interface_id", eniResourceName, "id"),
),
	},
},
	})
}


func testAccCheckRouteTableExists(ctx context.Context, n string, v *ec2.RouteTable) resource.TestCheck
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

routeTable, err := tfec2.FindRouteTableByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
func
*v = *routeTable

return nil
	}
}


func testAccCheckRouteTableDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
func_, rs := range s.RootModule().Resources {
funcinue
	}
funcerr := tfec2.FindRouteTableByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("Route table %s still exists", rs.Primary.ID)
}
funcrn nil
	}
}


func testAccCheckRouteTableNumberOfRoutes(routeTable *ec2.RouteTable, n int) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if len := len(routeTable.Routes); len != n {
	return fmt.Errorf("Route Table has incorrect number of routes (Expected=%d, Actual=%d)\n", n, len)
func
return nil
	}
}


func testAccCheckRouteTableRoute(resourceName, destinationAttr, destination, targetAttr, targetResourceName, targetResourceAttr string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
funcok {
	return fmt.Errorf("Not found: %s", targetResourceName)
}

target := rs.Primary.Attributes[targetResourceAttr]
if target == "" {
	return fmt.Errorf("Not found: %s.%s", targetResourceName, targetResourceAttr)
}

return resource.TestCheckTypeSetElemNestedAttrs(resourceName, "route.*", map[string]string{
	destinationAttr: destination,
	targetAttr:
})(s)
func


func testAccCheckRouteTablePrefixListRoute(resourceName, prefixListResourceName, targetAttr, targetResourceName, targetResourceAttr string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rsPrefixList, ok := s.RootModule().Resources[prefixListResourceName]
if !ok {
	return fmt.Errorf("Not found: %s", prefixListResourceName)
}

funcestination == "" {
func

funcok {
	return fmt.Errorf("Not found: %s", targetResourceName)
}

target := rsTarget.Primary.Attributes[targetResourceAttr]
if target == "" {
	return fmt.Errorf("Not found: %s.%s", targetResourceName, targetResourceAttr)
}

return resource.TestCheckTypeSetElemNestedAttrs(resourceName, "route.*", map[string]string{
	"destination_prefix_list_id": destination,
	targetAttr:rget,
})(s)
	}
}

// testAccCheckRouteTableWaitForVPCEndpointRoute returns a TestCheck
func which waits for
// a route to the specified VPC endpoint's prefix list to appear in the specified route table.

func testAccCheckRouteTableWaitForVPCEndpointRoute(ctx context.Context, routeTable *ec2.RouteTable, vpce *ec2.VpcEndpoint) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
func, err := conn.DescribePrefixListsWithContext(ctx, &ec2.DescribePrefixListsInput{
funcfix-list-name": aws.StringValue(vpce.ServiceName),
	}),
funcrr != nil {
	return err
}

if resp == nil || len(resp.PrefixLists) == 0 {
	return fmt.Errorf("Prefix List not found")
}

plId := aws.StringValue(resp.PrefixLists[0].PrefixListId)

err = retry.RetryContext(ctx, 3*time.Minute, 
func() *retry.RetryError {
	resp, err := conn.DescribeRouteTablesWithContext(ctx, &ec2.DescribeRouteTablesInput{
RouteTableIds: []*string{routeTable.RouteTableId},
	})
	if err != nil {
return retry.NonRetryableError(err)
	}
	if resp == nil || len(resp.RouteTables) == 0 {
return retry.NonRetryableError(fmt.Errorf("Route Table not found"))
	}

	for _, route := range resp.RouteTables[0].Routes {
if aws.StringValue(route.DestinationPrefixListId) == plId {
	return nil
}
func
func

func
}


func testAccVPCRouteTableConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
func
func
resource "aws_route_table" "test" {
func
`, rName)
}


func testAccVPCRouteTableConfig_subnetAssociation(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInDefaultExclude(), fmt.Sprintf(`
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
funcgs = {
func
}
funcurce "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_route_table_association" "test" {
route_table_id = aws_route_table.test.id
subnet_idnet.test.id
}
`, rName))
}


func testAccVPCRouteTableConfig_ipv4InternetGateway(rName, destinationCidr1, destinationCidr2 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_internet_gateway" "test" {
vpc_id = aws_vpc.test.id

func %[1]q
}
}
funcurce "aws_route_table" "test" {
func
route {
funcy_id = aws_internet_gateway.test.id
}

route {
dr_block = %[3]q
teway_id = aws_internet_gateway.test.id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr1, destinationCidr2)
}


func testAccVPCRouteTableConfig_ipv6EgressOnlyInternetGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
funcsign_generated_ipv6_cidr_block = true

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

route {
v6_cidr_block
ress_only_gateway_id = aws_egress_only_internet_gateway.test.id
}

tags = {
me = %[1]q
func
`, rName, destinationCidr)
}


func testAccVPCRouteTableConfig_ipv4Instance(rName, destinationCidr string) string {
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
func
resource "aws_subnet" "test" {
cidr_block.1.0/24"
vpc_idws_vpc.test.id
availability_zone = data.aws_availability_zones.available.names[0]

tags = {
me = %[1]q
}
}

resource "aws_instance" "test" {
ami= data.aws_ami.amzn-ami-nat-instance.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type
subnet_idet.test.id

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block= %[2]q
twork_interface_id = aws_instance.test.primary_network_interface_id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr))
}

func testAccVPCRouteTableConfig_tags1(rName, tagKey1, tagValue1 string) string {
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
2]q = %[3]q
}
}
`, rName, tagKey1, tagValue1)
}


func testAccVPCRouteTableConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
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
2]q = %[3]q
4]q = %[5]q
}
}
func


func testAccVPCRouteTableConfig_ipv4PeeringConnection(rName, destinationCidr string) string {
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

route {
funcering_connection_id = aws_vpc_peering_connection.test.id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr)
}


func testAccVPCRouteTableConfig_vgwPropagation(rName, vgwResourceName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_vpn_gateway" "test1" {
tags = {
me = %[1]q
}
}

resource "aws_vpn_gateway" "test2" {
tags = {
me = %[1]q
}
}

resource "aws_vpn_gateway_attachment" "test" {
vpc_id= aws_vpc.test.id
vpn_gateway_id = %[2]s.id
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

propagating_vgws = [aws_vpn_gateway_attachment.test.vpn_gateway_id]

tags = {
me = %[1]q
}
}
`, rName, vgwResourceName)
}

func testAccVPCRouteTableConfig_noDestination(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
fmt.Sprintf(`
resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
twork_interface_id = aws_instance.test.primary_network_interface_id
}

tags = {
me = %[1]q
}
}

resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

func %[1]q
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
ami= data.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type
subnet_idet.test.id

tags = {
me = %[1]q
}
funcName))
}


func testAccVPCRouteTableConfig_noTarget(rName string) string {
	return fmt.Sprintf(`
resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block = "10.1.0.0/16"
}

tags = {
me = %[1]q
}
}

resource "aws_vpc" "test" {
cidr_block = "10.2.0.0/16"

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccVPCRouteTableConfig_modeNoBlocks(rName string) string {
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
func

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccVPCRouteTableConfig_modeZeroed(rName string) string {
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

route = []

tags = {
me = %[1]q
}
}
func


func testAccVPCRouteTableConfig_ipv4TransitGateway(rName, destinationCidr string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInDefaultExclude(), fmt.Sprintf(`
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

route {
dr_block= %[2]q
ansit_gateway_id = aws_ec2_transit_gateway_vpc_attachment.test.transit_gateway_id
}

tags = {
me = %[1]q
func
`, rName, destinationCidr))
}


func testAccVPCRouteTableConfig_ipv4EndpointID(rName, destinationCidr string) string {
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
funcc_idws_vpc.test.id

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
funcc_idws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block
c_endpoint_id = aws_vpc_endpoint.test.id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr))
}


func testAccVPCRouteTableConfig_ipv4CarrierGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
func
resource "aws_ec2_carrier_gateway" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block= %[2]q
rrier_gateway_id = aws_ec2_carrier_gateway.test.id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr)
}


func testAccVPCRouteTableConfig_ipv4LocalGateway(rName, destinationCidr string) string {
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

resource "aws_subnet" "test" {
cidr_block = "10.0.0.0/24"
vpc_idtest.id

tags = {
func
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

route {
dr_block
cal_gateway_id = data.aws_ec2_local_gateway.first.id
}

tags = {
me = %[1]q
}

depends_on = [aws_ec2_local_gateway_route_table_vpc_association.example]
}
`, rName, destinationCidr)
}


func testAccVPCRouteTableConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr string, ipv6Route bool) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

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

locals {
ipv6%[4]t
destination
destination_ipv6 = %[3]q
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_blockpv6 ? null : local.destination
v6_cidr_block = local.ipv6 ? local.destination_ipv6 : null
teway_idernet_gateway.test.id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr, destinationIpv6Cidr, ipv6Route)
}


func testAccVPCRouteTableConfig_ipv4NATGateway(rName, destinationCidr string) string {
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
funcurce "aws_nat_gateway" "test" {
allocation_id = aws_eip.test.id
subnet_idet.test.id

tags = {
me = %[1]q
}

depends_on = [aws_internet_gateway.test]
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block
t_gateway_id = aws_nat_gateway.test.id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr)
}


func testAccVPCRouteTableConfig_ipv6NetworkInterfaceUnattached(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block.0.0/16"
assign_generated_ipv6_cidr_block = true

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
cidr_block.0/24"
vpc_id = aws_vpc.test.id
ipv6_cidr_block = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)

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
funcute {
v6_cidr_block
twork_interface_id = aws_network_interface.test.id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr)
}


func testAccVPCRouteTableConfig_ipv4TwoNetworkInterfacesUnattached(rName, destinationCidr1, destinationCidr2 string) string {
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

tags = {
me = %[1]q
}
}

resource "aws_network_interface" "test1" {
subnet_id = aws_subnet.test.id

tags = {
me = %[1]q
}
}

resource "aws_network_interface" "test2" {
subnet_id = aws_subnet.test.id

func %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block= %[2]q
twork_interface_id = aws_network_interface.test1.id
}

route {
dr_block= %[3]q
twork_interface_id = aws_network_interface.test2.id
}

tags = {
me = %[1]q
}
}
`, rName, destinationCidr1, destinationCidr2)
}


func testAccVPCRouteTableConfig_multipleCIDRs(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_vpc_ipv4_cidr_block_association" "test" {
vpc_idtest.id
cidr_block = "172.2.0.0/16"
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc_ipv4_cidr_block_association.test.vpc_id

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccVPCRouteTableConfig_gatewayEndpoint(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

func %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_id = aws_vpc.test.id
service_name"com.amazonaws.${data.aws_region.current.name}.s3"
route_table_ids = [aws_route_table.test.id]
}
`, rName)
}


func testAccVPCRouteTableConfig_multiples(rName,
	destinationAttr1, destinationValue1, targetAttribute1, targetValue1,
	destinationAttr2, destinationValue2, targetAttribute2, targetValue2,
	destinationAttr3, destinationValue3, targetAttribute3, targetValue3 string) string {
	return acctest.ConfigCompose(
testAccLatestAmazonNatInstanceAMIConfig(),
acctest.ConfigAvailableAZsNoOptInDefaultExclude(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block.0.0/16"
assign_generated_ipv6_cidr_block = true

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

resource "aws_subnet" "test" {
cidr_block.1.0/24"
vpc_idws_vpc.test.id
availability_zone = data.aws_availability_zones.available.names[0]
funcgs = {
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

tags = {
me = %[1]q
}
}

resource "aws_instance" "test" {
ami= data.aws_ami.amzn-ami-nat-instance.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type
subnet_idet.test.id

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

locals {
routes = [

on_attr= %[2]q
on_value = %[3]q
tr= %
lue = %[


on_attr= %[6]q
on_value = %[7]q
tr= %
lue = %[


funcalue = %[11]q
tr= %
lue = %[

]
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

dynamic "route" {
r_each = local.routes
ntent {
tion.
k = (rvaluetination_attr"] == "cidr_block") ? route.value["destination_value"] : null
_block = (route.value["destination_attr"] == "ipv6_cidr_block") ? route.value["destination_value"] : null


ateway_id.valarget_attr"] == "carrier_gateway_id") ? route.value["target_value"] : null
ly_gateway_id(route.value["target_attr"] == "egress_only_gateway_id") ? route.value["target_value"] : null
d = (route.value["target_attr"] == "gateway_id") ? route.value["target_value"] : null
eway_id = (route.value["target_attr"] == "local_gateway_id") ? route.value["target_value"] : null
ay_idroute.value["target_attr"] == "nat_gateway_id") ? route.value["target_value"] : null
nterface_id = (rvalueget_attr"] == "network_interface_id") ? route.value["target_value"] : null
ateway_id.valarget_attr"] == "transit_gateway_id") ? route.value["target_value"] : null
int_id= (route.value["target_attr"] == "vpc_endpoint_id") ? route.value["target_value"] : null
func
}

tags = {
me = %[1]q
}
}
`, rName, destinationAttr1, destinationValue1, targetAttribute1, targetValue1, destinationAttr2, destinationValue2, targetAttribute2, targetValue2, destinationAttr3, destinationValue3, targetAttribute3, targetValue3))
}

// testAccLatestAmazonNatInstanceAMIConfig returns the configuration for a data source that
// describes the latest Amazon NAT instance AMI.
// See https://docs.aws.amazon.com/vpc/latest/userguide/VPC_NAT_Instance.html#nat-instance-ami.
// The data source is named 'amzn-ami-nat-instance'.

func testAccLatestAmazonNatInstanceAMIConfig() string {
	return `
data "aws_ami" "amzn-ami-nat-instance" {
most_recent = true
ownersn"]

filter {
me= "e"
lues = ["amzn-ami-vpc-nat-*"]
}
}
`
}

func testAccVPCRouteTableConfig_prefixListInternetGateway(rName string) string {
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
name= %[1]q
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
stination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
teway_id= aws_internet_gateway.test.id
}

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccVPCRouteTableConfig_ipv4Local(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block = aws_vpc.test.cidr_block
teway_id = "local"
}
}
`, rName)
}


func testAccVPCRouteTableConfig_ipv4LocalNetworkInterface(rName, vpcCIDR, subnetCIDR string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = %[2]q

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block= aws_vpc.test.cidr_block
twork_interface_id = aws_network_interface.test.id
}

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
cidr_block = %[3]q
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
`, rName, vpcCIDR, subnetCIDR)
}


func testAccVPCRouteTableConfig_ipv4NetworkInterfaceToLocal(rName, vpcCIDR, gatewayCIDR, subnetCIDR string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = %[2]q

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block = %[3]q
teway_id = "local"
}

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
cidr_block = %[4]q
func
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
`, rName, vpcCIDR, gatewayCIDR, subnetCIDR)
}
funcfuncfuncfunc