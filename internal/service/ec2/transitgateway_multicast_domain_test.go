// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
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
	var v ec2.TransitGatewayMulticastDomain
	resourceName := "aws_ec2_transit_gateway_multicast_domain.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayMulticastDomainDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayMulticastDomainConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayMulticastDomainExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttr(resourceName, "auto_accept_shared_associations", "disable"),
	resource.TestCheckResourceAttr(resourceName, "igmpv2_support", "disable"),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "static_sources_support", "disable"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "transit_gateway_id"),
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


func testAccTransitGatewayMulticastDomain_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TransitGatewayMulticastDomain
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayMulticastDomainDestroy(ctx),
func
Config: testAccTransitGatewayMulticastDomainConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayMulticastDomainExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGatewayMulticastDomain(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func testAccTransitGatewayMulticastDomain_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TransitGatewayMulticastDomain
	resourceName := "aws_ec2_transit_gateway_multicast_domain.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTransitGatewayMulticastDomainDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayMulticastDomainConfig_tags1(rName, "key1", "value1"),
func(
	testAccCheckTransitGatewayMulticastDomainExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccTransitGatewayMulticastDomainConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccTransitGatewayMulticastDomainConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}


func := acctest.Context(t)
	var v ec2.TransitGatewayMulticastDomain
	resourceName := "aws_ec2_transit_gateway_multicast_domain.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccTransitGatewayMulticastDomainConfig_igmpv2Support(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckTransitGatewayMulticastDomainExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "auto_accept_shared_associations", "enable"),
	resource.TestCheckResourceAttr(resourceName, "igmpv2_support", "enable"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func testAccCheckTransitGatewayMulticastDomainExists(ctx context.Context, n string, v *ec2.TransitGatewayMulticastDomain) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Transit Gateway Multicast Domain ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
funcut, err := tfec2.FindTransitGatewayMulticastDomainByID(ctx, conn, rs.Primary.ID)
funcrr != nil {
	return err
func
*v = *output

return nil
	}
}


func testAccCheckTransitGatewayMulticastDomainDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_transit_gateway_multicast_domain" {
continue
	}

	_, err := tfec2.FindTransitGatewayMulticastDomainByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

funcrn err
func
	return fmt.Errorf("EC2 Transit Gateway Multicast Domain %s still exists", rs.Primary.ID)
func
return nil
	}
}


func testAccTransitGatewayMulticastDomainConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  multicast_support = "enable"

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_multicast_domain" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id
}
`, rName)
}


func testAccTransitGatewayMulticastDomainConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_multicast_domain" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id

  tags = {
2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}


funcurn fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  multicast_support = "enable"

  tags = {
me = %[1]q
  }
}
resource "aws_ec2_transit_gateway_multicast_domain" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id

  tags = {
2]q = %[3]q
4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}


func testAccTransitGatewayMulticastDomainConfig_igmpv2Support(rName string) string {
funcurce "aws_ec2_transit_gateway" "test" {
  multicast_support = "enable"

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway_multicast_domain" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id

  auto_accept_shared_associations = "enable"
  igmpv2_supportenable"

  tags = {
me = %[1]q
  }
}
`, rName)
}
func