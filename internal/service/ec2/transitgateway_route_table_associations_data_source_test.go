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
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_ec2_transit_gateway_route_table_associations.test"

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTableAssociationsDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "ids.#", 0),
func
},
	})
}


func testAccTransitGatewayRouteTableAssociationsDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccTransitGatewayRouteTableAssociationsDataSourceConfig_filter(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "1"),
),
	},
},
func


func testAccTransitGatewayRouteTableAssociationsDataSourceConfig_base(rName string) string {
	return acctest.ConfigCompose(testAccTransitGatewayRouteTableAssociationConfig_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids= aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id
funcansit_gateway_default_route_table_association = false

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table_association" "test" {
  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.test.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id
}
`, rName))
}


func testAccTransitGatewayRouteTableAssociationsDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccTransitGatewayRouteTableAssociationsDataSourceConfig_base(rName), `
data "aws_ec2_transit_gateway_route_table_associations" "test" {
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id

  depends_on = [aws_ec2_transit_gateway_route_table_association.test]
}
func


func testAccTransitGatewayRouteTableAssociationsDataSourceConfig_filter(rName string) string {
	return acctest.ConfigCompose(testAccTransitGatewayRouteTableAssociationsDataSourceConfig_base(rName), `
data "aws_ec2_transit_gateway_route_table_associations" "test" {
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id

  filter {
me= "nsit-gateway-attachment-id"
lues = [aws_ec2_transit_gateway_vpc_attachment.test.id]
func
  depends_on = [aws_ec2_transit_gateway_route_table_association.test]
}
`)
}
