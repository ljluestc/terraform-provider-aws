// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	instanceRouteResourceName := "aws_route.instance"
	pcxRouteResourceName := "aws_route.vpc_peering_connection"
	rtResourceName := "aws_route_table.test"
	instanceResourceName := "aws_instance.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
	datasource1Name := "data.aws_route.by_destination_cidr_block"
	datasource2Name := "data.aws_route.by_instance_id"
	datasource3Name := "data.aws_route.by_peering_connection_id"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	// By destination CIDR.
funcource.TestCheckResourceAttrPair(datasource1Name, "route_table_id", rtResourceName, "id"),

	// By instance ID.
	resource.TestCheckResourceAttrPair(datasource2Name, "destination_cidr_block", instanceRouteResourceName, "destination_cidr_block"),
	resource.TestCheckResourceAttrPair(datasource2Name, "instance_id", instanceResourceName, "id"),
	resource.TestCheckResourceAttrPair(datasource2Name, "route_table_id", rtResourceName, "id"),

	// By VPC peering connection ID.
	resource.TestCheckResourceAttrPair(datasource3Name, "destination_cidr_block", pcxRouteResourceName, "destination_cidr_block"),
	resource.TestCheckResourceAttrPair(datasource3Name, "route_table_id", rtResourceName, "id"),
	resource.TestCheckResourceAttrPair(datasource3Name, "vpc_peering_connection_id", pcxResourceName, "id"),
),
	},
},
	})
}


func TestAccVPCRouteDataSource_transitGatewayID(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
func

	dataSourceName := "data.aws_route.test"
	resourceName := "aws_route.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
Config: testAccVPCRouteDataSourceConfig_ipv4TransitGateway(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(resourceName, "destination_cidr_block", dataSourceName, "destination_cidr_block"),
	resource.TestCheckResourceAttrPair(resourceName, "route_table_id", dataSourceName, "route_table_id"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", dataSourceName, "transit_gateway_id"),
),
func
	})
}


func TestAccVPCRouteDataSource_ipv6DestinationCIDR(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_route.test"
	resourceName := "aws_route.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteDataSourceConfig_ipv6EgressOnlyInternetGateway(rName),
func(
	resource.TestCheckResourceAttrPair(resourceName, "destination_ipv6_cidr_block", dataSourceName, "destination_ipv6_cidr_block"),
	resource.TestCheckResourceAttrPair(resourceName, "route_table_id", dataSourceName, "route_table_id"),
),
	},
},
	})
}
func
func TestAccVPCRouteDataSource_localGatewayID(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_route.by_local_gateway_id"
	resourceName := "aws_route.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteDataSourceConfig_ipv4LocalGateway(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(resourceName, "destination_cidr_block", dataSourceName, "destination_cidr_block"),
funcource.TestCheckResourceAttrPair(resourceName, "local_gateway_id", dataSourceName, "local_gateway_id"),
),
	},
},
	})
}


func := acctest.Context(t)
	dataSourceName := "data.aws_route.test"
	resourceName := "aws_route.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckWavelengthZoneAvailable(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
func
Config: testAccVPCRouteDataSourceConfig_ipv4CarrierGateway(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(resourceName, "destination_cidr_block", dataSourceName, "destination_cidr_block"),
	resource.TestCheckResourceAttrPair(resourceName, "route_table_id", dataSourceName, "route_table_id"),
	resource.TestCheckResourceAttrPair(resourceName, "carrier_gateway_id", dataSourceName, "carrier_gateway_id"),
),
func
	})
}


func TestAccVPCRouteDataSource_destinationPrefixListID(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_route.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteDataSourceConfig_prefixListNATGateway(rName),
func(
	resource.TestCheckResourceAttrPair(resourceName, "destination_prefix_list_id", dataSourceName, "destination_prefix_list_id"),
	resource.TestCheckResourceAttrPair(resourceName, "nat_gateway_id", dataSourceName, "nat_gateway_id"),
	resource.TestCheckResourceAttrPair(resourceName, "route_table_id", dataSourceName, "route_table_id"),
),
	},
},
	})
func

func TestAccVPCRouteDataSource_gatewayVPCEndpoint(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	var vpce ec2.VpcEndpoint
	rtResourceName := "aws_route_table.test"
	vpceResourceName := "aws_vpc_endpoint.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteDataSourceConfig_gatewayEndpointNo(rName),
Check: resource.ComposeTestCheck
functAccCheckRouteTableExists(ctx, rtResourceName, &routeTable),
	testAccCheckVPCEndpointExists(ctx, vpceResourceName, &vpce),
	testAccCheckRouteTableWaitForVPCEndpointRoute(ctx, &routeTable, &vpce),
),
	},
	{
Config:CRouteDataSourceConfig_gatewayEndpoint(rName),
ExpectError: regexache.MustCompile(`No routes matching supplied arguments found in Route Table`),
	},
},
func


func testAccVPCRouteDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptInDefaultExclude(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
funcurce "aws_vpc" "test" {
cidr_block = "172.16.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_vpc" "target" {
cidr_block = "172.17.0.0/16"

tags = {
me = %[1]q
}
}
funcurce "aws_vpc_peering_connection" "test" {
peer_vpc_id = aws_vpc.target.id
vpc_id.test.id
auto_accept = true

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
cidr_block16.0.0/24"
vpc_idws_vpc.test.id
availability_zone = data.aws_availability_zones.available.names[0]

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

resource "aws_route" "vpc_peering_connection" {
route_table_idws_route_table.test.id
vpc_peering_connection_id = aws_vpc_peering_connection.test.id
destination_cidr_block"10.0.2.0/24"
}

resource "aws_route_table_association" "a" {
subnet_idnet.test.id
route_table_id = aws_route_table.test.id
}

resource "aws_instance" "test" {
ami= data.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type
subnet_idet.test.id

tags = {
me = %[1]q
}
}

resource "aws_route" "instance" {
route_table_id= aws_route_table.test.id
destination_cidr_block = "10.0.1.0/24"
network_interface_idws_instance.test.primary_network_interface_id
}

data "aws_route" "by_peering_connection_id" {
route_table_idws_route_table.test.id
vpc_peering_connection_id = aws_route.vpc_peering_connection.vpc_peering_connection_id
}

data "aws_route" "by_destination_cidr_block" {
route_table_id= aws_route_table.test.id
destination_cidr_block = aws_route.instance.destination_cidr_block
}

data "aws_route" "by_instance_id" {
route_table_id = aws_route_table.test.id
instance_idaws_route.instance.instance_id
}
`, rName))
}


func testAccVPCRouteDataSourceConfig_ipv4TransitGateway(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptInDefaultExclude(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
cidr_block.0.0/24"
vpc_idws_vpc.test.id

tags = {
me = %[1]q
}
}

funcgs = {
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

resource "aws_route" "test" {
destination_cidr_block = "0.0.0.0/0"
route_table_id= aws_vpc.test.default_route_table_id
transit_gateway_idtransit_gateway_vpc_attachment.test.transit_gateway_id
}

data "aws_route" "test" {
route_table_ide.test.route_table_id
transit_gateway_id = aws_route.test.transit_gateway_id
}
`, rName))
}


func testAccVPCRouteDataSourceConfig_ipv6EgressOnlyInternetGateway(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

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
func

resource "aws_route" "test" {
route_table_ide_table.test.id
destination_ipv6_cidr_block = "::/0"
egress_only_gateway_idess_only_internet_gateway.test.id
}

data "aws_route" "test" {
route_table_ide.test.route_table_id
destination_ipv6_cidr_block = aws_route.test.destination_ipv6_cidr_block
}
`, rName)
}


func testAccVPCRouteDataSourceConfig_ipv4LocalGateway(rName string) string {
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
route_table_id= aws_route_table.test.id
destination_cidr_block = "172.16.1.0/24"
local_gateway_idws_ec2_local_gateway.first.id
depends_on[aws_ec2_local_gateway_route_table_vpc_association.example]
}

data "aws_route" "by_local_gateway_id" {
route_table_idws_route_table.test.id
local_gateway_id = data.aws_ec2_local_gateway.first.id
depends_onoute.test]
}
`, rName)
}


func testAccVPCRouteDataSourceConfig_ipv4CarrierGateway(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

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
funcurce "aws_route" "test" {
destination_cidr_block = "0.0.0.0/0"
route_table_id= aws_route_table.test.id
carrier_gateway_idcarrier_gateway.test.id
}

data "aws_route" "test" {
route_table_ide.test.route_table_id
carrier_gateway_id = aws_route.test.carrier_gateway_id
}
`, rName)
}


func testAccVPCRouteDataSourceConfig_prefixListNATGateway(rName string) string {
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
name= %[1]q
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

data "aws_route" "test" {
route_table_idaws_route.test.route_table_id
destination_prefix_list_id = aws_route.test.destination_prefix_list_id
nat_gateway_idaws_route.test.nat_gateway_id
}
`, rName)
}


func testAccVPCRouteDataSourceConfig_gatewayEndpointNo(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_id = aws_vpc.test.id
service_name"com.amazonaws.${data.aws_region.current.name}.s3"
route_table_ids = [aws_route_table.test.id]
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}
`, rName)
}
func
func testAccVPCRouteDataSourceConfig_gatewayEndpoint(rName string) string {
	return acctest.ConfigCompose(testAccVPCRouteDataSourceConfig_gatewayEndpointNo(rName), `
data "aws_prefix_list" "test" {
name = aws_vpc_endpoint.test.service_name
}

data "aws_route" "test" {
route_table_idaws_route_table.test.id
destination_prefix_list_id = data.aws_prefix_list.test.id
}
`)
}
func