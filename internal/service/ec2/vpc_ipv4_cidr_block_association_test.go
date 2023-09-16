// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

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
	var associationSecondary, associationTertiary ec2.VpcCidrBlockAssociation
	resource1Name := "aws_vpc_ipv4_cidr_block_association.secondary_cidr"
	resource2Name := "aws_vpc_ipv4_cidr_block_association.tertiary_cidr"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCIPv4CIDRBlockAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCIPv4CIDRBlockAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCIPv4CIDRBlockAssociationExists(ctx, resource1Name, &associationSecondary),
functAccCheckVPCIPv4CIDRBlockAssociationExists(ctx, resource2Name, &associationTertiary),
	testAccCheckAdditionalVPCIPv4CIDRBlock(&associationTertiary, "170.2.0.0/16"),
),
	},
	{
ResourceName:Name,
ImportState:
ImportStateVerify: true,
	},
	{
ResourceName:Name,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCIPv4CIDRBlockAssociation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var associationSecondary, associationTertiary ec2.VpcCidrBlockAssociation
funcource2Name := "aws_vpc_ipv4_cidr_block_association.tertiary_cidr"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCIPv4CIDRBlockAssociationDestroy(ctx),
func
Config: testAccVPCIPv4CIDRBlockAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCIPv4CIDRBlockAssociationExists(ctx, resource1Name, &associationSecondary),
	testAccCheckVPCIPv4CIDRBlockAssociationExists(ctx, resource2Name, &associationTertiary),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPCIPv4CIDRBlockAssociation(), resource1Name),
),
func
},
	})
}


func TestAccVPCIPv4CIDRBlockAssociation_ipamBasic(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCIPv4CIDRBlockAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCIPv4CIDRBlockAssociationConfig_ipam(rName, 28),
func(
	testAccCheckVPCIPv4CIDRBlockAssociationExists(ctx, "aws_vpc_ipv4_cidr_block_association.secondary_cidr", &associationSecondary),
	testAccCheckVPCAssociationCIDRPrefix(&associationSecondary, "28"),
),
	},
},
	})
}
func
func TestAccVPCIPv4CIDRBlockAssociation_ipamBasicExplicitCIDR(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var associationSecondary ec2.VpcCidrBlockAssociation
	cidr := "172.2.0.32/28"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCIPv4CIDRBlockAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCIPv4CIDRBlockAssociationConfig_ipamExplicit(rName, cidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCIPv4CIDRBlockAssociationExists(ctx, "aws_vpc_ipv4_cidr_block_association.secondary_cidr", &associationSecondary),
func
},
	})
}


func testAccCheckAdditionalVPCIPv4CIDRBlock(association *ec2.VpcCidrBlockAssociation, expected string) resource.TestCheck
func {
func(s *terraform.State) error {
CIDRBlock := association.CidrBlock
if *CIDRBlock != expected {
	return fmt.Errorf("Bad CIDR: %s", *association.CidrBlock)
}

return nil
	}
}
func
func {
	return 
functrings.Split(aws.StringValue(association.CidrBlock), "/")[1] != expected {
	return fmt.Errorf("Bad cidr prefix: %s", aws.StringValue(association.CidrBlock))
}

return nil
	}
}


func testAccCheckVPCIPv4CIDRBlockAssociationDestroy(ctx context.Context) resource.TestCheck
func {
func(s *terraform.State) error {
func
for _, rs := range s.RootModule().Resources {
funcinue
	}

	_, _, err := tfec2.FindVPCCIDRBlockAssociationByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
func
funcurn fmt.Errorf("EC2 VPC IPv4 CIDR Block Association %s still exists", rs.Primary.ID)
}
funcrn nil
	}
}


func testAccCheckVPCIPv4CIDRBlockAssociationExists(ctx context.Context, n string, v *ec2.VpcCidrBlockAssociation) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 VPC IPv4 CIDR Block Association is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, _, err := tfec2.FindVPCCIDRBlockAssociationByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

func
func
}
func
func testAccVPCIPv4CIDRBlockAssociationConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_ipv4_cidr_block_association" "secondary_cidr" {
  vpc_idtest.id
  cidr_block = "172.2.0.0/16"
}

resource "aws_vpc_ipv4_cidr_block_association" "tertiary_cidr" {
  vpc_idtest.id
  cidr_block = "170.2.0.0/16"
}
`, rName)
}


func testAccVPCIPv4CIDRBlockAssociationConfig_ipam(rName string, netmaskLength int) string {
funcurce "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_ipv4_cidr_block_association" "secondary_cidr" {
  ipv4_ipam_pool_idws_vpc_ipam_pool.test.id
  ipv4_netmask_length = %[2]d
  vpc_idtest.id

  depends_on = [aws_vpc_ipam_pool_cidr.test]
}
`, rName, netmaskLength))
}


func testAccVPCIPv4CIDRBlockAssociationConfig_ipamExplicit(rName, cidr string) string {
	return acctest.ConfigCompose(testAccVPCConfig_baseIPAMIPv4(rName), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"
funcgs = {
me = %[1]q
  }
}

resource "aws_vpc_ipv4_cidr_block_association" "secondary_cidr" {
  ipv4_ipam_pool_id = aws_vpc_ipam_pool.test.id
  cidr_block
  vpc_idws_vpc.test.id

  depends_on = [aws_vpc_ipam_pool_cidr.test]
}
`, rName, cidr))
}
func