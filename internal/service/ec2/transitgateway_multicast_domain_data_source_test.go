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
	dataSourceName := "data.aws_ec2_transit_gateway_multicast_domain.test"
	resourceName := "aws_ec2_transit_gateway_multicast_domain.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayMulticastDomainDataSourceConfig_filter(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
funcource.TestCheckResourceAttrPair(dataSourceName, "auto_accept_shared_associations", resourceName, "auto_accept_shared_associations"),
	resource.TestCheckResourceAttrPair(dataSourceName, "igmpv2_support", resourceName, "igmpv2_support"),
	resource.TestCheckResourceAttr(dataSourceName, "members.#", "0"),
	resource.TestCheckResourceAttrPair(dataSourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttr(dataSourceName, "sources.#", "0"),
	resource.TestCheckResourceAttrPair(dataSourceName, "static_sources_support", resourceName, "static_sources_support"),
	resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(dataSourceName, "transit_gateway_id", resourceName, "transit_gateway_id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "transit_gateway_multicast_domain_id", resourceName, "id"),
),
	},
},
	})
}


func testAccTransitGatewayMulticastDomainDataSource_ID(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ec2_transit_gateway_multicast_domain.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccTransitGatewayMulticastDomainDataSourceConfig_iD(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
	resource.TestCheckResourceAttr(dataSourceName, "associations.#", "1"),
	resource.TestCheckResourceAttrPair(dataSourceName, "auto_accept_shared_associations", resourceName, "auto_accept_shared_associations"),
	resource.TestCheckResourceAttrPair(dataSourceName, "igmpv2_support", resourceName, "igmpv2_support"),
funcource.TestCheckResourceAttrPair(dataSourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttr(dataSourceName, "sources.#", "1"),
	resource.TestCheckResourceAttrPair(dataSourceName, "static_sources_support", resourceName, "static_sources_support"),
	resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(dataSourceName, "transit_gateway_id", resourceName, "transit_gateway_id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "transit_gateway_multicast_domain_id", resourceName, "id"),
),
	},
},
	})
}


func testAccTransitGatewayMulticastDomainDataSourceConfig_filter(rName string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
multicast_support = "enable"

tags = {
func
}

resource "aws_ec2_transit_gateway_multicast_domain" "test" {
transit_gateway_id = aws_ec2_transit_gateway.test.id

tags = {
me = %[1]q
}
}

data "aws_ec2_transit_gateway_multicast_domain" "test" {
filter {
me= "nsit-gateway-multicast-domain-id"
lues = [aws_ec2_transit_gateway_multicast_domain.test.id]
}
}
`, rName)
}


func testAccTransitGatewayMulticastDomainDataSourceConfig_iD(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInDefaultExclude(), fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
func

resource "aws_subnet" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
cidr_block.0.0/24"
vpc_idws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_ec2_transit_gateway" "test" {
multicast_support = "enable"

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

resource "aws_ec2_transit_gateway_multicast_domain" "test" {
transit_gateway_id = aws_ec2_transit_gateway.test.id

static_sources_support = "enable"

tags = {
me = %[1]q
}
}

resource "aws_ec2_transit_gateway_multicast_domain_association" "test" {
subnet_idws_subnet.test.id
transit_gateway_attachment_id2_transit_gateway_vpc_attachment.test.id
transit_gateway_multicast_domain_id = aws_ec2_transit_gateway_multicast_domain.test.id
}

resource "aws_network_interface" "test3" {
subnet_id = aws_subnet.test.id

tags = {
me = %[1]q
}
}

resource "aws_ec2_transit_gateway_multicast_group_source" "test" {
group_ip_address.1"
network_interface_id = aws_network_interface.test3.id
transit_gateway_multicast_domain_id = aws_ec2_transit_gateway_multicast_domain_association.test.transit_gateway_multicast_domain_id
}

resource "aws_network_interface" "test1" {
subnet_id = aws_subnet.test.id

tags = {
me = %[1]q
}
}

resource "aws_network_interface" "test2" {
subnet_id = aws_subnet.test.id

tags = {
me = %[1]q
}
}

resource "aws_ec2_transit_gateway_multicast_group_member" "test1" {
group_ip_address.1"
network_interface_id = aws_network_interface.test1.id
transit_gateway_multicast_domain_id = aws_ec2_transit_gateway_multicast_domain_association.test.transit_gateway_multicast_domain_id
}

resource "aws_ec2_transit_gateway_multicast_group_member" "test2" {
group_ip_address.1"
network_interface_id = aws_network_interface.test2.id
transit_gateway_multicast_domain_id = aws_ec2_transit_gateway_multicast_domain_association.test.transit_gateway_multicast_domain_id
}

data "aws_ec2_transit_gateway_multicast_domain" "test" {
transit_gateway_multicast_domain_id = aws_ec2_transit_gateway_multicast_domain.test.id

depends_on = [
s_ec2_transit_gateway_multicast_group_member.test1,
s_ec2_transit_gateway_multicast_group_member.test2,
s_ec2_transit_gateway_multicast_group_source.test,
]
}
`, rName))
}
