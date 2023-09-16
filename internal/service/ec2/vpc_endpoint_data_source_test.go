// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	resourceName := "aws_vpc_endpoint.test"
	datasourceName := "data.aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointDataSourceConfig_gatewayBasic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
funcource.TestCheckResourceAttrPair(datasourceName, "dns_entry.#", resourceName, "dns_entry.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_options.#", resourceName, "dns_options.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
	resource.TestCheckResourceAttrPair(datasourceName, "ip_address_type", resourceName, "ip_address_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "network_interface_ids.#", resourceName, "network_interface_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "policy", resourceName, "policy"),
	resource.TestCheckResourceAttrPair(datasourceName, "prefix_list_id", resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_enabled", resourceName, "private_dns_enabled"),
	resource.TestCheckResourceAttrPair(datasourceName, "requester_managed", resourceName, "requester_managed"),
	resource.TestCheckResourceAttrPair(datasourceName, "route_table_ids.#", resourceName, "route_table_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "security_group_ids.#", resourceName, "security_group_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "service_name", resourceName, "service_name"),
	resource.TestCheckResourceAttrPair(datasourceName, "state", resourceName, "state"),
	resource.TestCheckResourceAttrPair(datasourceName, "subnet_ids.#", resourceName, "subnet_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_endpoint_type", resourceName, "vpc_endpoint_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_id", resourceName, "vpc_id"),
),
	},
},
	})
}


func TestAccVPCEndpointDataSource_byID(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_vpc_endpoint.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccVPCEndpointDataSourceConfig_byID(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
	resource.TestCheckResourceAttrPair(datasourceName, "cidr_blocks.#", resourceName, "cidr_blocks.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_entry.#", resourceName, "dns_entry.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_options.#", resourceName, "dns_options.#"),
funcource.TestCheckResourceAttrPair(datasourceName, "ip_address_type", resourceName, "ip_address_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "network_interface_ids.#", resourceName, "network_interface_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "policy", resourceName, "policy"),
	resource.TestCheckResourceAttrPair(datasourceName, "prefix_list_id", resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_enabled", resourceName, "private_dns_enabled"),
	resource.TestCheckResourceAttrPair(datasourceName, "requester_managed", resourceName, "requester_managed"),
	resource.TestCheckResourceAttrPair(datasourceName, "route_table_ids.#", resourceName, "route_table_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "security_group_ids.#", resourceName, "security_group_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "service_name", resourceName, "service_name"),
	resource.TestCheckResourceAttrPair(datasourceName, "state", resourceName, "state"),
	resource.TestCheckResourceAttrPair(datasourceName, "subnet_ids.#", resourceName, "subnet_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_endpoint_type", resourceName, "vpc_endpoint_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_id", resourceName, "vpc_id"),
),
	},
},
	})
}


func TestAccVPCEndpointDataSource_byFilter(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_vpc_endpoint.test"
	datasourceName := "data.aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointDataSourceConfig_byFilter(rName),
Check: resource.ComposeAggregateTestCheck
funcource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
	resource.TestCheckResourceAttrPair(datasourceName, "cidr_blocks.#", resourceName, "cidr_blocks.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_entry.#", resourceName, "dns_entry.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_options.#", resourceName, "dns_options.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
	resource.TestCheckResourceAttrPair(datasourceName, "ip_address_type", resourceName, "ip_address_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "network_interface_ids.#", resourceName, "network_interface_ids.#"),
funcource.TestCheckResourceAttrPair(datasourceName, "policy", resourceName, "policy"),
	resource.TestCheckResourceAttrPair(datasourceName, "prefix_list_id", resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_enabled", resourceName, "private_dns_enabled"),
	resource.TestCheckResourceAttrPair(datasourceName, "requester_managed", resourceName, "requester_managed"),
	resource.TestCheckResourceAttrPair(datasourceName, "route_table_ids.#", resourceName, "route_table_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "security_group_ids.#", resourceName, "security_group_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "service_name", resourceName, "service_name"),
	resource.TestCheckResourceAttrPair(datasourceName, "state", resourceName, "state"),
	resource.TestCheckResourceAttrPair(datasourceName, "subnet_ids.#", resourceName, "subnet_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_endpoint_type", resourceName, "vpc_endpoint_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_id", resourceName, "vpc_id"),
),
	},
},
	})
}


func TestAccVPCEndpointDataSource_byTags(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_vpc_endpoint.test"
	datasourceName := "data.aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointDataSourceConfig_byTags(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
	resource.TestCheckResourceAttrPair(datasourceName, "cidr_blocks.#", resourceName, "cidr_blocks.#"),
funcource.TestCheckResourceAttrPair(datasourceName, "dns_options.#", resourceName, "dns_options.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
	resource.TestCheckResourceAttrPair(datasourceName, "ip_address_type", resourceName, "ip_address_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "network_interface_ids.#", resourceName, "network_interface_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "policy", resourceName, "policy"),
	resource.TestCheckResourceAttrPair(datasourceName, "prefix_list_id", resourceName, "prefix_list_id"),
funcource.TestCheckResourceAttrPair(datasourceName, "requester_managed", resourceName, "requester_managed"),
	resource.TestCheckResourceAttrPair(datasourceName, "route_table_ids.#", resourceName, "route_table_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "security_group_ids.#", resourceName, "security_group_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "service_name", resourceName, "service_name"),
	resource.TestCheckResourceAttrPair(datasourceName, "state", resourceName, "state"),
	resource.TestCheckResourceAttrPair(datasourceName, "subnet_ids.#", resourceName, "subnet_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_endpoint_type", resourceName, "vpc_endpoint_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_id", resourceName, "vpc_id"),
),
	},
},
	})
}


func TestAccVPCEndpointDataSource_gatewayWithRouteTableAndTags(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_vpc_endpoint.test"
	datasourceName := "data.aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccVPCEndpointDataSourceConfig_gatewayRouteTableAndTags(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
	resource.TestCheckResourceAttrPair(datasourceName, "cidr_blocks.#", resourceName, "cidr_blocks.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_entry.#", resourceName, "dns_entry.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_options.#", resourceName, "dns_options.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
funcource.TestCheckResourceAttrPair(datasourceName, "network_interface_ids.#", resourceName, "network_interface_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "policy", resourceName, "policy"),
	resource.TestCheckResourceAttrPair(datasourceName, "prefix_list_id", resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_enabled", resourceName, "private_dns_enabled"),
	resource.TestCheckResourceAttrPair(datasourceName, "requester_managed", resourceName, "requester_managed"),
	resource.TestCheckResourceAttrPair(datasourceName, "route_table_ids.#", resourceName, "route_table_ids.#"),
funcource.TestCheckResourceAttrPair(datasourceName, "service_name", resourceName, "service_name"),
	resource.TestCheckResourceAttrPair(datasourceName, "state", resourceName, "state"),
	resource.TestCheckResourceAttrPair(datasourceName, "subnet_ids.#", resourceName, "subnet_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_endpoint_type", resourceName, "vpc_endpoint_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_id", resourceName, "vpc_id"),
),
	},
},
	})
}


func TestAccVPCEndpointDataSource_interface(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_vpc_endpoint.test"
	datasourceName := "data.aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointDataSourceConfig_interface(rName),
Check: resource.ComposeAggregateTestCheck
funcource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
	resource.TestCheckResourceAttrPair(datasourceName, "cidr_blocks.#", resourceName, "cidr_blocks.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_entry.#", resourceName, "dns_entry.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "dns_options.#", resourceName, "dns_options.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
	resource.TestCheckResourceAttrPair(datasourceName, "ip_address_type", resourceName, "ip_address_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "network_interface_ids.#", resourceName, "network_interface_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "owner_id", resourceName, "owner_id"),
funcource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_enabled", resourceName, "private_dns_enabled"),
	resource.TestCheckResourceAttrPair(datasourceName, "requester_managed", resourceName, "requester_managed"),
	resource.TestCheckResourceAttrPair(datasourceName, "route_table_ids.#", resourceName, "route_table_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "security_group_ids.#", resourceName, "security_group_ids.#"),
	resource.TestCheckResourceAttrPair(datasourceName, "service_name", resourceName, "service_name"),
	resource.TestCheckResourceAttrPair(datasourceName, "state", resourceName, "state"),
funcource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_endpoint_type", resourceName, "vpc_endpoint_type"),
	resource.TestCheckResourceAttrPair(datasourceName, "vpc_id", resourceName, "vpc_id"),
),
	},
},
	})
}


func testAccVPCEndpointDataSourceConfig_gatewayBasic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

tags = {
me = %[1]q
func

data "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = aws_vpc_endpoint.test.service_name
statelable"
}
`, rName)
}


func testAccVPCEndpointDataSourceConfig_byID(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

tags = {
me = %[1]q
}
func
data "aws_vpc_endpoint" "test" {
id = aws_vpc_endpoint.test.id
}
`, rName)
}


func testAccVPCEndpointDataSourceConfig_byFilter(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

tags = {
me = %[1]q
}
}
func "aws_vpc_endpoint" "test" {
filter {
me= "-endpoint-id"
lues = [aws_vpc_endpoint.test.id]
}
}
`, rName)
}


func testAccVPCEndpointDataSourceConfig_byTags(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

tags = {
me = %[1]q
y1 = "Value1"
y2 = "Value2"
y3 = "Value3"
func

data "aws_vpc_endpoint" "test" {
vpc_id = aws_vpc_endpoint.test.vpc_id

tags = {
me = %[1]q
y1 = "Value1"
y2 = "Value2"
y3 = "Value3"
}
}
`, rName)
}


func testAccVPCEndpointDataSourceConfig_gatewayRouteTableAndTags(rName string) string {
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

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
func
route_table_ids = [
s_route_table.test.id,
]

tags = {
me = %[1]q
}
}

data "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = aws_vpc_endpoint.test.service_name
statelable"
}
`, rName)
}


func testAccVPCEndpointDataSourceConfig_interface(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
vpc_idws_vpc.test.id
cidr_blockpc.test.cidr_block
availability_zone = data.aws_availability_zones.available.names[0]

tags = {
me = %[1]q
}
}

resource "aws_security_group" "test" {
vpc_id = aws_vpc.test.id
name[1]q
funcgs = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idtest.id
vpc_endpoint_typeInterface"
service_nameamazonaws.${data.aws_region.current.name}.ec2"
private_dns_enabled = false

subnet_ids = [
s_subnet.test.id,
]

security_group_ids = [
s_security_group.test.id,
]

tags = {
me = %[1]q
}
}

data "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = aws_vpc_endpoint.test.service_name
statelable"
}
`, rName))
}
