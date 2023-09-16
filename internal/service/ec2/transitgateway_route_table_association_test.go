// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

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
	var v ec2.TransitGatewayRouteTableAssociation
	resourceName := "aws_ec2_transit_gateway_route_table_association.test"
	transitGatewayRouteTableResourceName := "aws_ec2_transit_gateway_route_table.test"
	transitGatewayVpcAttachmentResourceName := "aws_ec2_transit_gateway_vpc_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTableAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayRouteTableAssociationExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttrSet(resourceName, "resource_id"),
	resource.TestCheckResourceAttrSet(resourceName, "resource_type"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_attachment_id", transitGatewayVpcAttachmentResourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_route_table_id", transitGatewayRouteTableResourceName, "id"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"replace_existing_association"},
	},
},
	})
}


func testAccTransitGatewayRouteTableAssociation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TransitGatewayRouteTableAssociation
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayRouteTableAssociationDestroy(ctx),
func
Config: testAccTransitGatewayRouteTableAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayRouteTableAssociationExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGatewayRouteTableAssociation(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func testAccTransitGatewayRouteTableAssociation_replaceExistingAssociation(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TransitGatewayRouteTableAssociation
	resourceName := "aws_ec2_transit_gateway_route_table_association.test"
	transitGatewayRouteTableResourceName := "aws_ec2_transit_gateway_route_table.test"
	transitGatewayVpcAttachmentResourceName := "aws_ec2_transit_gateway_vpc_attachment.test"
func
	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayRouteTableAssociationConfig_replaceExistingAssociation(rName),
func(
	testAccCheckTransitGatewayRouteTableAssociationExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "replace_existing_association", "true"),
	resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
	resource.TestCheckResourceAttrSet(resourceName, "resource_type"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_attachment_id", transitGatewayVpcAttachmentResourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_route_table_id", transitGatewayRouteTableResourceName, "id"),
	resource.TestCheckResourceAttr(transitGatewayVpcAttachmentResourceName, "transit_gateway_default_route_table_association", "true"),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"replace_existing_association"},
	},
},
	})
}


func testAccCheckTransitGatewayRouteTableAssociationExists(ctx context.Context, n string, v *ec2.TransitGatewayRouteTableAssociation) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

funcurn fmt.Errorf("No EC2 Transit Gateway Route Table Association ID is set")
func
transitGatewayRouteTableID, transitGatewayAttachmentID, err := tfec2.TransitGatewayRouteTableAssociationParseResourceID(rs.Primary.ID)
funcrr != nil {
	return err
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindTransitGatewayRouteTableAssociationByTwoPartKey(ctx, conn, transitGatewayRouteTableID, transitGatewayAttachmentID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccCheckTransitGatewayRouteTableAssociationDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_transit_gateway_route_table_association" {
continue
	}

	transitGatewayRouteTableID, transitGatewayAttachmentID, err := tfec2.TransitGatewayRouteTableAssociationParseResourceID(rs.Primary.ID)
funcerr != nil {
func

func
	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Transit Gateway Route Table Association %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccTransitGatewayRouteTableAssociationConfig_base(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
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
funcName))
}


func testAccTransitGatewayRouteTableAssociationConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccTransitGatewayRouteTableAssociationConfig_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_idsubnet.test[*].id
  transit_gateway_default_route_table_association = false
  transit_gateway_id= aws_ec2_transit_gateway.test.id
  vpc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table_association" "test" {
  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.test.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id
funcName))
}


func testAccTransitGatewayRouteTableAssociationConfig_replaceExistingAssociation(rName string) string {
	return acctest.ConfigCompose(testAccTransitGatewayRouteTableAssociationConfig_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids= aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table_association" "test" {
  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.test.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id

  replace_existing_association = true
funcName))
}
