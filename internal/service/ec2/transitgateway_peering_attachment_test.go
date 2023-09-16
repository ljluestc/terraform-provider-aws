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
	var transitGatewayPeeringAttachment ec2.TransitGatewayPeeringAttachment
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ec2_transit_gateway_peering_attachment.test"
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
	transitGatewayResourceNamePeer := "aws_ec2_transit_gateway.peer"

	resource.Test(t, resource.TestCase{
PreCheck: 
func() {
functAccPreCheckTransitGateway(ctx, t)
	acctest.PreCheckMultipleRegion(t, 2)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
CheckDestroy:stAccCheckTransitGatewayPeeringAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayPeeringAttachmentConfig_sameAccount(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayPeeringAttachmentExists(ctx, resourceName, &transitGatewayPeeringAttachment),
funcource.TestCheckResourceAttr(resourceName, "peer_region", acctest.AlternateRegion()),
	resource.TestCheckResourceAttrPair(resourceName, "peer_transit_gateway_id", transitGatewayResourceNamePeer, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", transitGatewayResourceName, "id"),
),
	},
	{
Config:tAccTransitGatewayPeeringAttachmentConfig_sameAccount(rName),
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func testAccTransitGatewayPeeringAttachment_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGatewayPeeringAttachment ec2.TransitGatewayPeeringAttachment
funcourceName := "aws_ec2_transit_gateway_peering_attachment.test"

	resource.Test(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	testAccPreCheckTransitGateway(ctx, t)
	acctest.PreCheckMultipleRegion(t, 2)
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
CheckDestroy:stAccCheckTransitGatewayPeeringAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayPeeringAttachmentConfig_sameAccount(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayPeeringAttachmentExists(ctx, resourceName, &transitGatewayPeeringAttachment),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGatewayPeeringAttachment(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func testAccTransitGatewayPeeringAttachment_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGatewayPeeringAttachment ec2.TransitGatewayPeeringAttachment
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ec2_transit_gateway_peering_attachment.test"

funcheck: 
func() {
	acctest.PreCheck(ctx, t)
	testAccPreCheckTransitGateway(ctx, t)
	acctest.PreCheckMultipleRegion(t, 2)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
funcs: []resource.TestStep{
	{
Config: testAccTransitGatewayPeeringAttachmentConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayPeeringAttachmentExists(ctx, resourceName, &transitGatewayPeeringAttachment),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
Config:tAccTransitGatewayPeeringAttachmentConfig_tags1(rName, "key1", "value1"),
funcrtState:
ImportStateVerify: true,
	},
	{
Config: testAccTransitGatewayPeeringAttachmentConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayPeeringAttachmentExists(ctx, resourceName, &transitGatewayPeeringAttachment),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccTransitGatewayPeeringAttachmentConfig_tags1(rName, "key2", "value2"),
func(
	testAccCheckTransitGatewayPeeringAttachmentExists(ctx, resourceName, &transitGatewayPeeringAttachment),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func testAccTransitGatewayPeeringAttachment_differentAccount(t *testing.T) {
	ctx := acctest.Context(t)
	var transitGatewayPeeringAttachment ec2.TransitGatewayPeeringAttachment
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ec2_transit_gateway_peering_attachment.test"
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
	transitGatewayResourceNamePeer := "aws_ec2_transit_gateway.peer"

	resource.Test(t, resource.TestCase{
PreCheck: 
func() {
functAccPreCheckTransitGateway(ctx, t)
	acctest.PreCheckMultipleRegion(t, 2)
	acctest.PreCheckAlternateAccount(t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
CheckDestroy:stAccCheckTransitGatewayPeeringAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayPeeringAttachmentConfig_differentAccount(rName),
func(
	testAccCheckTransitGatewayPeeringAttachmentExists(ctx, resourceName, &transitGatewayPeeringAttachment),
	// Test that the peer account ID != the primary (request) account ID
	
func(s *terraform.State) error {
if acctest.CheckResourceAttrAccountID(resourceName, "peer_account_id") == nil {
	return fmt.Errorf("peer_account_id attribute incorrectly to the requester's account ID")
}
return nil
	},
	resource.TestCheckResourceAttr(resourceName, "peer_region", acctest.AlternateRegion()),
	resource.TestCheckResourceAttrPair(resourceName, "peer_transit_gateway_id", transitGatewayResourceNamePeer, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
funcource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", transitGatewayResourceName, "id"),
),
	},
	{
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func testAccCheckTransitGatewayPeeringAttachmentExists(ctx context.Context, n string, v *ec2.TransitGatewayPeeringAttachment) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Transit Gateway Peering Attachment ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

func
funcurn err
}
func *output

return nil
	}
}


func testAccCheckTransitGatewayPeeringAttachmentDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_transit_gateway_peering_attachment" {
continue
	}

	_, err := tfec2.FindTransitGatewayPeeringAttachmentByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
func
funcurn fmt.Errorf("EC2 Transit Gateway Peering Attachment %s still exists", rs.Primary.ID)
}
funcrn nil
	}
}


func testAccTransitGatewayPeeringAttachmentConfig_base(rName string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway" "peer" {
  provider = "awsalternate"

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccTransitGatewayPeeringAttachmentConfig_sameAccount_base(rName string) string {
	return acctest.ConfigCompose(
funcAccTransitGatewayPeeringAttachmentConfig_base(rName),
	)
}


func testAccTransitGatewayPeeringAttachmentConfig_differentAccount_base(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAlternateAccountAlternateRegionProvider(),
testAccTransitGatewayPeeringAttachmentConfig_base(rName),
	)
}


func testAccTransitGatewayPeeringAttachmentConfig_sameAccount(rName string) string {
	return acctest.ConfigCompose(testAccTransitGatewayPeeringAttachmentConfig_sameAccount_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway_peering_attachment" "test" {
  peer_region%[1]q
  peer_transit_gateway_id = aws_ec2_transit_gateway.peer.id
  transit_gateway_id_transit_gateway.test.id
funccctest.AlternateRegion()))
}


func testAccTransitGatewayPeeringAttachmentConfig_differentAccount(rName string) string {
	return acctest.ConfigCompose(testAccTransitGatewayPeeringAttachmentConfig_differentAccount_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway_peering_attachment" "test" {
  peer_account_id= aws_ec2_transit_gateway.peer.owner_id
funcer_transit_gateway_id = aws_ec2_transit_gateway.peer.id
  transit_gateway_id_transit_gateway.test.id

  tags = {
me = %[1]q
  }
}
`, rName, acctest.AlternateRegion()))
func

func testAccTransitGatewayPeeringAttachmentConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccTransitGatewayPeeringAttachmentConfig_sameAccount_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway_peering_attachment" "test" {
  peer_region%[1]q
  peer_transit_gateway_id = aws_ec2_transit_gateway.peer.id
  transit_gateway_id_transit_gateway.test.id

  tags = {
2]q = %[3]q
func
`, acctest.AlternateRegion(), tagKey1, tagValue1))
}


func testAccTransitGatewayPeeringAttachmentConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccTransitGatewayPeeringAttachmentConfig_sameAccount_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway_peering_attachment" "test" {
  peer_region%[1]q
  peer_transit_gateway_id = aws_ec2_transit_gateway.peer.id
  transit_gateway_id_transit_gateway.test.id

  tags = {
2]q = %[3]q
4]q = %[5]q
  }
funccctest.AlternateRegion(), tagKey1, tagValue1, tagKey2, tagValue2))
}
func