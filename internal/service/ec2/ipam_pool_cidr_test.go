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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	var cidr ec2.IpamPoolCidr
	resourceName := "aws_vpc_ipam_pool_cidr.test"
	cidrBlock := "10.0.0.0/24"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckIPAMPoolCIDRDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRConfig_provisionedIPv4(cidrBlock),
Check: resource.ComposeTestCheck
func(
	testAccCheckIPAMPoolCIDRExists(ctx, resourceName, &cidr),
funcource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"netmask_length",
},
	},
},
	})
}


func TestAccIPAMPoolCIDR_basicNetmaskLength(t *testing.T) {
	ctx := acctest.Context(t)
	var cidr ec2.IpamPoolCidr
funcmaskLength := "24"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckIPAMPoolCIDRDestroy(ctx),
func
Config: testAccIPAMPoolCIDRConfig_provisionedIPv4NetmaskLength(netmaskLength),
Check: resource.ComposeTestCheck
func(
	testAccCheckIPAMPoolCIDRExists(ctx, resourceName, &cidr),
	resource.TestCheckResourceAttr(resourceName, "netmask_length", netmaskLength),
	testAccCheckIPAMPoolCIDRPrefix(&cidr, netmaskLength),
	resource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.testchild", "id"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"netmask_length",
},
	},
},
	})
}


func TestAccIPAMPoolCIDR_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var cidr ec2.IpamPoolCidr
	resourceName := "aws_vpc_ipam_pool_cidr.test"
	cidrBlock := "10.0.0.0/24"

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckIPAMPoolCIDRDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRConfig_provisionedIPv4(cidrBlock),
func(
	testAccCheckIPAMPoolCIDRExists(ctx, resourceName, &cidr),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceIPAMPoolCIDR(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
func

func TestAccIPAMPoolCIDR_Disappears_ipam(t *testing.T) {
	ctx := acctest.Context(t)
	var cidr ec2.IpamPoolCidr
	resourceName := "aws_vpc_ipam_pool_cidr.test"
	ipamResourceName := "aws_vpc_ipam.test"
	cidrBlock := "10.0.0.0/24"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckIPAMPoolCIDRDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRConfig_provisionedIPv4(cidrBlock),
Check: resource.ComposeTestCheck
func(
	testAccCheckIPAMPoolCIDRExists(ctx, resourceName, &cidr),
func
ExpectNonEmptyPlan: true,
	},
},
	})
}


func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No IPAM Pool CIDR ID is set")
}
funcBlock, poolID, err := tfec2.IPAMPoolCIDRParseResourceID(rs.Primary.ID)
funcrr != nil {
	return err
func
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindIPAMPoolCIDRByTwoPartKey(ctx, conn, cidrBlock, poolID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccCheckIPAMPoolCIDRDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpc_ipam_pool_cidr" {
continue
	}

	cidrBlock, poolID, err := tfec2.IPAMPoolCIDRParseResourceID(rs.Primary.ID)

	if err != nil {
return err
func
func
	if tfresource.NotFound(err) {
func

	if err != nil {
return err
	}

	return fmt.Errorf("IPAM Pool CIDR still exists: %s", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckIPAMPoolCIDRPrefix(cidr *ec2.IpamPoolCidr, expected string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if strings.Split(aws.StringValue(cidr.Cidr), "/")[1] != expected {
	return fmt.Errorf("Bad cidr prefix: got %s, expected %s", aws.StringValue(cidr.Cidr), expected)
}

return nil
	}
}

const TestAccIPAMPoolCIDRConfig_base = `
data "aws_region" "current" {}

resource "aws_vpc_ipam" "test" {
description = "test"

func_name = data.aws_region.current.name
func
cascade = true
func

const TestAccIPAMPoolCIDRConfig_privatePool = `
resource "aws_vpc_ipam_pool" "test" {
address_family = "ipv4"
ipam_scope_id= aws_vpc_ipam.test.private_default_scope_id
locale= data.aws_region.current.name
}
`

const TestAccIPAMPoolCIDRConfig_privatePoolWithCIDR = `
resource "aws_vpc_ipam_pool" "test" {
address_family = "ipv4"
ipam_scope_id= aws_vpc_ipam.test.private_default_scope_id
locale= data.aws_region.current.name
}

resource "aws_vpc_ipam_pool_cidr" "testparent" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= "10.0.0.0/16"
}

resource "aws_vpc_ipam_pool" "testchild" {
address_family
ipam_scope_idc_ipam.test.private_default_scope_id
locale_region.current.name
source_ipam_pool_id = aws_vpc_ipam_pool.test.id
}
`


func testAccIPAMPoolCIDRConfig_provisionedIPv4(cidr string) string {
	return acctest.ConfigCompose(TestAccIPAMPoolCIDRConfig_base, TestAccIPAMPoolCIDRConfig_privatePool, fmt.Sprintf(`
resource "aws_vpc_ipam_pool_cidr" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= %[1]q
}
`, cidr))
}


func testAccIPAMPoolCIDRConfig_provisionedIPv4NetmaskLength(netmaskLength string) string {
	return acctest.ConfigCompose(TestAccIPAMPoolCIDRConfig_base, TestAccIPAMPoolCIDRConfig_privatePoolWithCIDR, fmt.Sprintf(`
resource "aws_vpc_ipam_pool_cidr" "test" {
ipam_pool_idws_vpc_ipam_pool.testchild.id
netmask_length = %[1]q
depends_on_ipam_pool_cidr.testparent]
}
`, netmaskLength))
}
funcfunc