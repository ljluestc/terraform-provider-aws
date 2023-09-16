// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
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


func := acctest.Context(t)
	var transitGatewayRouteTable1 ec2.TransitGatewayRouteTable
	resourceName := "aws_ec2_transit_gateway_route_table.test"
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTableConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayRouteTableExists(ctx, resourceName, &transitGatewayRouteTable1),
funcource.TestCheckResourceAttr(resourceName, "default_association_route_table", "false"),
	resource.TestCheckResourceAttr(resourceName, "default_propagation_route_table", "false"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", transitGatewayResourceName, "id"),
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


func testAccTransitGatewayRouteTable_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGatewayRouteTable1 ec2.TransitGatewayRouteTable
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayRouteTableDestroy(ctx),
func
Config: testAccTransitGatewayRouteTableConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayRouteTableExists(ctx, resourceName, &transitGatewayRouteTable1),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGatewayRouteTable(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func testAccTransitGatewayRouteTable_disappears_TransitGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGateway1 ec2.TransitGateway
	var transitGatewayRouteTable1 ec2.TransitGatewayRouteTable
	resourceName := "aws_ec2_transit_gateway_route_table.test"
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
func
	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTableConfig_basic(rName),
func(
	testAccCheckTransitGatewayExists(ctx, transitGatewayResourceName, &transitGateway1),
	testAccCheckTransitGatewayRouteTableExists(ctx, resourceName, &transitGatewayRouteTable1),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGateway(), transitGatewayResourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func testAccTransitGatewayRouteTable_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGatewayRouteTable1, transitGatewayRouteTable2, transitGatewayRouteTable3 ec2.TransitGatewayRouteTable
	resourceName := "aws_ec2_transit_gateway_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayRouteTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTableConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayRouteTableExists(ctx, resourceName, &transitGatewayRouteTable1),
funcource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
funcig: testAccTransitGatewayRouteTableConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayRouteTableExists(ctx, resourceName, &transitGatewayRouteTable2),
	testAccCheckTransitGatewayRouteTableNotRecreated(&transitGatewayRouteTable1, &transitGatewayRouteTable2),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccTransitGatewayRouteTableConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
functAccCheckTransitGatewayRouteTableNotRecreated(&transitGatewayRouteTable2, &transitGatewayRouteTable3),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}


func testAccCheckTransitGatewayRouteTableExists(ctx context.Context, n string, v *ec2.TransitGatewayRouteTable) resource.TestCheck
funcurn 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Transit Gateway Route Table ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
funcut, err := tfec2.FindTransitGatewayRouteTableByID(ctx, conn, rs.Primary.ID)
funcrr != nil {
	return err
func
*v = *output

return nil
	}
}


func testAccCheckTransitGatewayRouteTableDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_transit_gateway_route_table" {
continue
	}

	_, err := tfec2.FindTransitGatewayRouteTableByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

funcrn err
func
	return fmt.Errorf("EC2 Transit Gateway Route Table %s still exists", rs.Primary.ID)
func
return nil
	}
}


func testAccCheckTransitGatewayRouteTableNotRecreated(i, j *ec2.TransitGatewayRouteTable) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if aws.StringValue(i.TransitGatewayRouteTableId) != aws.StringValue(j.TransitGatewayRouteTableId) {
	return errors.New("EC2 Transit Gateway Route Table was recreated")
}

return nil
	}
}


func testAccTransitGatewayRouteTableConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
tags = {
me = %[1]q
}
}
funcurce "aws_ec2_transit_gateway_route_table" "test" {
func
`, rName)
func

func testAccTransitGatewayRouteTableConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
tags = {
me = %[1]q
}
}

funcansit_gateway_id = aws_ec2_transit_gateway.test.id

tags = {
2]q = %[3]q
}
}
`, rName, tagKey1, tagValue1)
}


func testAccTransitGatewayRouteTableConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
tags = {
me = %[1]q
func

resource "aws_ec2_transit_gateway_route_table" "test" {
transit_gateway_id = aws_ec2_transit_gateway.test.id

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}
func