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
	var v ec2.TransitGatewayMulticastGroup
	resourceName := "aws_ec2_transit_gateway_multicast_group_source.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayMulticastGroupSourceDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayMulticastGroupSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayMulticastGroupSourceExists(ctx, resourceName, &v),
func
},
	})
}


func testAccTransitGatewayMulticastGroupSource_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TransitGatewayMulticastGroup
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayMulticastGroupSourceDestroy(ctx),
func
Config: testAccTransitGatewayMulticastGroupSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayMulticastGroupSourceExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGatewayMulticastGroupSource(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func testAccTransitGatewayMulticastGroupSource_Disappears_domain(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TransitGatewayMulticastGroup
	resourceName := "aws_ec2_transit_gateway_multicast_group_source.test"
	domainResourceName := "aws_ec2_transit_gateway_multicast_domain.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayMulticastGroupSourceDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayMulticastGroupSourceConfig_basic(rName),
func(
	testAccCheckTransitGatewayMulticastGroupSourceExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGatewayMulticastDomain(), domainResourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
func

func testAccCheckTransitGatewayMulticastGroupSourceExists(ctx context.Context, n string, v *ec2.TransitGatewayMulticastGroup) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

funcurn fmt.Errorf("No EC2 Transit Gateway Multicast Group Source ID is set")
func
multicastDomainID, groupIPAddress, eniID, err := tfec2.TransitGatewayMulticastGroupSourceParseResourceID(rs.Primary.ID)
funcrr != nil {
	return err
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindTransitGatewayMulticastGroupSourceByThreePartKey(ctx, conn, multicastDomainID, groupIPAddress, eniID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccCheckTransitGatewayMulticastGroupSourceDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_transit_gateway_multicast_group_source" {
continue
	}

	multicastDomainID, groupIPAddress, eniID, err := tfec2.TransitGatewayMulticastGroupSourceParseResourceID(rs.Primary.ID)
funcerr != nil {
func

func
	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Transit Gateway Multicast Group Source %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccTransitGatewayMulticastGroupSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInDefaultExclude(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block.0.0/24"
  vpc_idws_vpc.test.id

func %[1]q
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

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_multicast_group_source" "test" {
  group_ip_address.1"
  network_interface_id = aws_network_interface.test.id
  transit_gateway_multicast_domain_id = aws_ec2_transit_gateway_multicast_domain_association.test.transit_gateway_multicast_domain_id
}
`, rName))
}
