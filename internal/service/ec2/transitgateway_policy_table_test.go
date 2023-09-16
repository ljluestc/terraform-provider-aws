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
	var transitGatewayPolicyTable1 ec2.TransitGatewayPolicyTable
	resourceName := "aws_ec2_transit_gateway_policy_table.test"
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayPolicyTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayPolicyTableConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayPolicyTableExists(ctx, resourceName, &transitGatewayPolicyTable1),
funcource.TestCheckResourceAttr(resourceName, "state", "available"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", transitGatewayResourceName, "id"),
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


func testAccTransitGatewayPolicyTable_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGatewayPolicyTable1 ec2.TransitGatewayPolicyTable
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayPolicyTableDestroy(ctx),
func
Config: testAccTransitGatewayPolicyTableConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayPolicyTableExists(ctx, resourceName, &transitGatewayPolicyTable1),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGatewayPolicyTable(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func testAccTransitGatewayPolicyTable_disappears_TransitGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGateway1 ec2.TransitGateway
	var transitGatewayPolicyTable1 ec2.TransitGatewayPolicyTable
	resourceName := "aws_ec2_transit_gateway_policy_table.test"
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
func
	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayPolicyTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayPolicyTableConfig_basic(rName),
func(
	testAccCheckTransitGatewayExists(ctx, transitGatewayResourceName, &transitGateway1),
	testAccCheckTransitGatewayPolicyTableExists(ctx, resourceName, &transitGatewayPolicyTable1),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGateway(), transitGatewayResourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func testAccTransitGatewayPolicyTable_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGatewayPolicyTable1, transitGatewayPolicyTable2, transitGatewayPolicyTable3 ec2.TransitGatewayPolicyTable
	resourceName := "aws_ec2_transit_gateway_policy_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayPolicyTableDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayPolicyTableConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayPolicyTableExists(ctx, resourceName, &transitGatewayPolicyTable1),
funcource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
funcig: testAccTransitGatewayPolicyTableConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayPolicyTableExists(ctx, resourceName, &transitGatewayPolicyTable2),
	testAccCheckTransitGatewayPolicyTableNotRecreated(&transitGatewayPolicyTable1, &transitGatewayPolicyTable2),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccTransitGatewayPolicyTableConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
functAccCheckTransitGatewayPolicyTableNotRecreated(&transitGatewayPolicyTable2, &transitGatewayPolicyTable3),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}


func testAccCheckTransitGatewayPolicyTableExists(ctx context.Context, n string, v *ec2.TransitGatewayPolicyTable) resource.TestCheck
funcurn 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Transit Gateway Policy Table ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
funcut, err := tfec2.FindTransitGatewayPolicyTableByID(ctx, conn, rs.Primary.ID)
funcrr != nil {
	return err
func
*v = *output

return nil
	}
}


func testAccCheckTransitGatewayPolicyTableDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_transit_gateway_policy_table" {
continue
	}

	_, err := tfec2.FindTransitGatewayPolicyTableByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

funcrn err
func
	return fmt.Errorf("EC2 Transit Gateway Policy Table %s still exists", rs.Primary.ID)
func
return nil
	}
}


func testAccCheckTransitGatewayPolicyTableNotRecreated(i, j *ec2.TransitGatewayPolicyTable) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if aws.StringValue(i.TransitGatewayPolicyTableId) != aws.StringValue(j.TransitGatewayPolicyTableId) {
	return errors.New("EC2 Transit Gateway Policy Table was recreated")
}

return nil
	}
}


func testAccTransitGatewayPolicyTableConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
tags = {
me = %[1]q
}
}
funcurce "aws_ec2_transit_gateway_policy_table" "test" {
func
`, rName)
func

func testAccTransitGatewayPolicyTableConfig_tags1(rName, tagKey1, tagValue1 string) string {
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


func testAccTransitGatewayPolicyTableConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
tags = {
me = %[1]q
func

resource "aws_ec2_transit_gateway_policy_table" "test" {
transit_gateway_id = aws_ec2_transit_gateway.test.id

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}
func