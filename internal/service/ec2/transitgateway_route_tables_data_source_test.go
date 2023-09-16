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
	dataSourceName := "data.aws_ec2_transit_gateway_route_tables.test"

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTablesDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "ids.#", 0),
func
},
	})
}


func testAccTransitGatewayRouteTablesDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccTransitGatewayRouteTablesDataSourceConfig_filter(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "2"),
),
	},
},
func


func testAccTransitGatewayRouteTablesDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_ec2_transit_gateway_route_tables.test"

	resource.Test(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTablesDataSourceConfig_tags(rName),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr(dataSourceName, "ids.#", "1"),
),
	},
},
	})
}

func testAccTransitGatewayRouteTablesDataSource_empty(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_ec2_transit_gateway_route_tables.test"

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funcs: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTablesDataSourceConfig_empty(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "0"),
),
func
	})
}


func testAccTransitGatewayRouteTablesDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
funcgs = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id

  tags = {
func
}

data "aws_ec2_transit_gateway_route_tables" "test" {
  depends_on = [aws_ec2_transit_gateway_route_table.test]
}
`, rName)
}


func testAccTransitGatewayRouteTablesDataSourceConfig_filter(rName string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id

  tags = {
me = %[1]q
func

data "aws_ec2_transit_gateway_route_tables" "test" {
  filter {
me= "nsit-gateway-id"
lues = [aws_ec2_transit_gateway.test.id]
  }

  depends_on = [aws_ec2_transit_gateway_route_table.test]
}
`, rName)
}


func testAccTransitGatewayRouteTablesDataSourceConfig_tags(rName string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id

  tags = {
me = %[1]q
  }
func
data "aws_ec2_transit_gateway_route_tables" "test" {
  tags = {
me = %[1]q
  }

  depends_on = [aws_ec2_transit_gateway_route_table.test]
}
`, rName)
}


func testAccTransitGatewayRouteTablesDataSourceConfig_empty(rName string) string {
	return fmt.Sprintf(`
data "aws_ec2_transit_gateway_route_tables" "test" {
  tags = {
me = %[1]q
  }
}
`, rName)
}
func