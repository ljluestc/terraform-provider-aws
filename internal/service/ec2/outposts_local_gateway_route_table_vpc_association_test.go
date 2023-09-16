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
)


func := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	localGatewayRouteTableDataSourceName := "data.aws_ec2_local_gateway_route_table.test"
	resourceName := "aws_ec2_local_gateway_route_table_vpc_association.test"
	vpcResourceName := "aws_vpc.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLocalGatewayRouteTableVPCAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccOutpostsLocalGatewayRouteTableVPCAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLocalGatewayRouteTableVPCAssociationExists(ctx, resourceName),
funcource.TestCheckResourceAttrPair(resourceName, "local_gateway_route_table_id", localGatewayRouteTableDataSourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_id", vpcResourceName, "id"),
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


func TestAccEC2OutpostsLocalGatewayRouteTableVPCAssociation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLocalGatewayRouteTableVPCAssociationDestroy(ctx),
func
Config: testAccOutpostsLocalGatewayRouteTableVPCAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLocalGatewayRouteTableVPCAssociationExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceLocalGatewayRouteTableVPCAssociation(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccEC2OutpostsLocalGatewayRouteTableVPCAssociation_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ec2_local_gateway_route_table_vpc_association.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLocalGatewayRouteTableVPCAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccOutpostsLocalGatewayRouteTableVPCAssociationConfig_tags1(rName, "key1", "value1"),
func(
	testAccCheckLocalGatewayRouteTableVPCAssociationExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccOutpostsLocalGatewayRouteTableVPCAssociationConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLocalGatewayRouteTableVPCAssociationExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccOutpostsLocalGatewayRouteTableVPCAssociationConfig_tags1(rName, "key2", "value2"),
func(
	testAccCheckLocalGatewayRouteTableVPCAssociationExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func testAccCheckLocalGatewayRouteTableVPCAssociationExists(ctx context.Context, resourceName string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return fmt.Errorf("Not found: %s", resourceName)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("%s: missing resource ID", resourceName)
func
func
association, err := tfec2.GetLocalGatewayRouteTableVPCAssociation(ctx, conn, rs.Primary.ID)
funcrr != nil {
	return err
}

if association == nil {
	return fmt.Errorf("EC2 Local Gateway Route Table VPC Association (%s) not found", rs.Primary.ID)
}

if aws.StringValue(association.State) != ec2.RouteTableAssociationStateCodeAssociated {
	return fmt.Errorf("EC2 Local Gateway Route Table VPC Association (%s) not in associated state: %s", rs.Primary.ID, aws.StringValue(association.State))
}

return nil
	}
}


func testAccCheckLocalGatewayRouteTableVPCAssociationDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_local_gateway_route_table_vpc_association" {
continue
	}

	association, err := tfec2.GetLocalGatewayRouteTableVPCAssociation(ctx, conn, rs.Primary.ID)

	if err != nil {
func
funcassociation != nil && aws.StringValue(association.State) != ec2.RouteTableAssociationStateCodeDisassociated {
return fmt.Errorf("EC2 Local Gateway Route Table VPC Association (%s) still exists in state: %s", rs.Primary.ID, aws.StringValue(association.State))
func

return nil
	}
}


func testAccLocalGatewayRouteTableVPCAssociationBaseConfig(rName string) string {
	return fmt.Sprintf(`
data "aws_outposts_outposts" "test" {}

data "aws_ec2_local_gateway_route_table" "test" {
outpost_arn = tolist(data.aws_outposts_outposts.test.arns)[0]
}

resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}
`, rName)
}
func
func testAccOutpostsLocalGatewayRouteTableVPCAssociationConfig_basic(rName string) string {
	return acctest.ConfigCompose(
testAccLocalGatewayRouteTableVPCAssociationBaseConfig(rName),
`
resource "aws_ec2_local_gateway_route_table_vpc_association" "test" {
local_gateway_route_table_id = data.aws_ec2_local_gateway_route_table.test.id
vpc_idpc.test.id
}
`)
}


func testAccOutpostsLocalGatewayRouteTableVPCAssociationConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(
testAccLocalGatewayRouteTableVPCAssociationBaseConfig(rName),
fmt.Sprintf(`
resource "aws_ec2_local_gateway_route_table_vpc_association" "test" {
local_gateway_route_table_id = data.aws_ec2_local_gateway_route_table.test.id
func
tags = {
1]q = %[2]q
}
}
`, tagKey1, tagValue1))
}


func testAccOutpostsLocalGatewayRouteTableVPCAssociationConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(
testAccLocalGatewayRouteTableVPCAssociationBaseConfig(rName),
funcurce "aws_ec2_local_gateway_route_table_vpc_association" "test" {
local_gateway_route_table_id = data.aws_ec2_local_gateway_route_table.test.id
vpc_idpc.test.id

tags = {
1]q = %[2]q
3]q = %[4]q
}
}
`, tagKey1, tagValue1, tagKey2, tagValue2))
}
func