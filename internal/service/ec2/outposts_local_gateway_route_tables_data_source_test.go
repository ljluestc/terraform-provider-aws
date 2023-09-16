// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	dataSourceName := "data.aws_ec2_local_gateway_route_tables.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccOutpostsLocalGatewayRouteTablesDataSourceConfig_basic(),
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "ids.#", 0),
func
},
	})
}


func TestAccEC2OutpostsLocalGatewayRouteTablesDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ec2_local_gateway_route_tables.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccOutpostsLocalGatewayRouteTablesDataSourceConfig_filter(),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "1"),
),
	},
},
func


func testAccOutpostsLocalGatewayRouteTablesDataSourceConfig_basic() string {
	return `
data "aws_ec2_local_gateway_route_tables" "test" {}
`
}

func testAccOutpostsLocalGatewayRouteTablesDataSourceConfig_filter() string {
	return `
data "aws_ec2_local_gateway_route_tables" "all" {}

data "aws_ec2_local_gateway_route_tables" "test" {
filter {
me= "al-gateway-route-table-id"
func
}
`
}
