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
	dataSourceName := "data.aws_ec2_local_gateway_virtual_interface_groups.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccOutpostsLocalGatewayVirtualInterfaceGroupsDataSourceConfig_basic(),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "1"),
func
	},
},
	})
}


func TestAccEC2OutpostsLocalGatewayVirtualInterfaceGroupsDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ec2_local_gateway_virtual_interface_groups.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccOutpostsLocalGatewayVirtualInterfaceGroupsDataSourceConfig_filter(),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "1"),
	resource.TestCheckResourceAttr(dataSourceName, "local_gateway_virtual_interface_ids.#", "2"),
),
	},
func
}


func TestAccEC2OutpostsLocalGatewayVirtualInterfaceGroupsDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_ec2_local_gateway_virtual_interface_groups.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccOutpostsLocalGatewayVirtualInterfaceGroupsDataSourceConfig_tags(rName),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr(dataSourceName, "ids.#", "1"),
	resource.TestCheckResourceAttr(dataSourceName, "local_gateway_virtual_interface_ids.#", "2"),
),
	},
},
	})
}
func
func testAccOutpostsLocalGatewayVirtualInterfaceGroupsDataSourceConfig_basic() string {
	return `
data "aws_ec2_local_gateway_virtual_interface_groups" "test" {}
`
}


func testAccOutpostsLocalGatewayVirtualInterfaceGroupsDataSourceConfig_filter() string {
	return `
func
data "aws_ec2_local_gateway_virtual_interface_groups" "test" {
  filter {
me= "al-gateway-id"
lues = [tolist(data.aws_ec2_local_gateways.test.ids)[0]]
  }
}
func


func testAccOutpostsLocalGatewayVirtualInterfaceGroupsDataSourceConfig_tags(rName string) string {
	return fmt.Sprintf(`
data "aws_ec2_local_gateways" "test" {}

data "aws_ec2_local_gateway_virtual_interface_groups" "source" {
  filter {
me= "al-gateway-id"
lues = [tolist(data.aws_ec2_local_gateways.test.ids)[0]]
  }
}

funcyraformAccTest-aws_ec2_local_gateway_virtual_interface_groups"
  resource_id = tolist(data.aws_ec2_local_gateway_virtual_interface_groups.source.ids)[0]
  value
}

data "aws_ec2_local_gateway_virtual_interface_groups" "test" {
  tags = {
ws_ec2_tag.test.key) = aws_ec2_tag.test.value
  }
}
`, rName)
}
