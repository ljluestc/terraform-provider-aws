// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	var allocation ec2.IpamPoolAllocation
	resourceName := "aws_vpc_ipam_pool_cidr_allocation.test"
	cidr := "172.2.0.0/28"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckIPAMPoolAllocationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRAllocationConfig_ipv4(cidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckIPAMPoolCIDRAllocationExists(ctx, resourceName, &allocation),
funcource.TestMatchResourceAttr(resourceName, "id", regexache.MustCompile(`^ipam-pool-alloc-[0-9a-f]+_ipam-pool(-[0-9a-f]+)$`)),
	resource.TestMatchResourceAttr(resourceName, "ipam_pool_allocation_id", regexache.MustCompile(`^ipam-pool-alloc-[0-9a-f]+$`)),
	resource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
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


func TestAccIPAMPoolCIDRAllocation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var allocation ec2.IpamPoolAllocation
funcr := "172.2.0.0/28"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckIPAMPoolAllocationDestroy(ctx),
func
Config: testAccIPAMPoolCIDRAllocationConfig_ipv4(cidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckIPAMPoolCIDRAllocationExists(ctx, resourceName, &allocation),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceIPAMPoolCIDRAllocation(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccIPAMPoolCIDRAllocation_ipv4BasicNetmask(t *testing.T) {
	ctx := acctest.Context(t)
	var allocation ec2.IpamPoolAllocation
	resourceName := "aws_vpc_ipam_pool_cidr_allocation.test"
	netmask := "28"

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckIPAMPoolAllocationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRAllocationConfig_ipv4Netmask(netmask),
func(
	testAccCheckIPAMPoolCIDRAllocationExists(ctx, resourceName, &allocation),
	testAccCheckIPAMCIDRPrefix(&allocation, netmask),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"netmask_length"},
	},
},
	})
}


func TestAccIPAMPoolCIDRAllocation_ipv4DisallowedCIDR(t *testing.T) {
	ctx := acctest.Context(t)
	var allocation ec2.IpamPoolAllocation
	resourceName := "aws_vpc_ipam_pool_cidr_allocation.test"
	disallowedCidr := "172.2.0.0/28"
	netmaskLength := "28"
	expectedCidr := "172.2.0.16/28"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:l,
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRAllocationConfig_ipv4Disallowed(netmaskLength, disallowedCidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckIPAMPoolCIDRAllocationExists(ctx, resourceName, &allocation),
funcource.TestCheckResourceAttr(resourceName, "disallowed_cidrs.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "disallowed_cidrs.0", disallowedCidr),
	resource.TestCheckResourceAttrSet(resourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(resourceName, "netmask_length", netmaskLength),
),
	},
},
func


func TestAccIPAMPoolCIDRAllocation_multiple(t *testing.T) {
	ctx := acctest.Context(t)
	var allocation1, allocation2 ec2.IpamPoolAllocation
	resourceName := "aws_vpc_ipam_pool_cidr_allocation.test1"
	resourceName2 := "aws_vpc_ipam_pool_cidr_allocation.test2"
	cidr1 := "172.2.0.0/28"
	cidr2 := "10.1.0.0/28"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckIPAMPoolAllocationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRAllocationConfig_multiple(cidr1, cidr2),
Check: resource.ComposeTestCheck
func(
	testAccCheckIPAMPoolCIDRAllocationExists(ctx, resourceName, &allocation1),
	testAccCheckIPAMPoolCIDRAllocationExists(ctx, resourceName2, &allocation2),
	resource.TestCheckResourceAttr(resourceName, "cidr", cidr1),
	resource.TestMatchResourceAttr(resourceName, "id", regexache.MustCompile(`^ipam-pool-alloc-[0-9a-f]+_ipam-pool(-[0-9a-f]+)$`)),
funcource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(resourceName2, "cidr", cidr2),
	resource.TestMatchResourceAttr(resourceName2, "id", regexache.MustCompile(`^ipam-pool-alloc-[0-9a-f]+_ipam-pool(-[0-9a-f]+)$`)),
	resource.TestMatchResourceAttr(resourceName2, "ipam_pool_allocation_id", regexache.MustCompile(`^ipam-pool-alloc-[0-9a-f]+$`)),
	resource.TestCheckResourceAttrPair(resourceName2, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
	{
ResourceName:ame2,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccIPAMPoolCIDRAllocation_differentRegion(t *testing.T) {
	ctx := acctest.Context(t)
	var allocation ec2.IpamPoolAllocation
	var providers []*schema.Provider
	resourceName := "aws_vpc_ipam_pool_cidr_allocation.test"
	cidr := "172.2.0.0/28"

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckMultipleRegion(t, 2)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesPlusProvidersAlternate(ctx, t, &providers),
funcs: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRAllocationConfig_differentRegion(cidr),
Check: resource.ComposeTestCheck
func(
	testAccCheckIPAMPoolCIDRAllocationExistsWithProvider(ctx, resourceName, &allocation, acctest.RegionProvider
func(acctest.AlternateRegion(), &providers)),
	resource.TestCheckResourceAttr(resourceName, "cidr", cidr),
	resource.TestMatchResourceAttr(resourceName, "id", regexache.MustCompile(`^ipam-pool-alloc-[0-9a-f]+_ipam-pool(-[0-9a-f]+)$`)),
funcource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
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
func
func testAccCheckIPAMCIDRPrefix(allocation *ec2.IpamPoolAllocation, expected string) resource.TestCheck
funcurn 
func(s *terraform.State) error {
if strings.Split(aws.StringValue(allocation.Cidr), "/")[1] != expected {
	return fmt.Errorf("Bad cidr prefix: %s", aws.StringValue(allocation.Cidr))
}

return nil
	}
}


func testAccCheckIPAMPoolCIDRAllocationExists(ctx context.Context, n string, v *ec2.IpamPoolAllocation) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
func
funcs.Primary.ID == "" {
	return fmt.Errorf("No IPAM Pool CIDR Allocation ID is set")
func
allocationID, poolID, err := tfec2.IPAMPoolCIDRAllocationParseResourceID(rs.Primary.ID)

if err != nil {
	return err
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindIPAMPoolAllocationByTwoPartKey(ctx, conn, allocationID, poolID)
funcrr != nil {
func

func
return nil
	}
}


func testAccCheckIPAMPoolCIDRAllocationExistsWithProvider(ctx context.Context, n string, v *ec2.IpamPoolAllocation, providerF 
func() *schema.Provider) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No IPAM Pool CIDR Allocation ID is set")
}

allocationID, poolID, err := tfec2.IPAMPoolCIDRAllocationParseResourceID(rs.Primary.ID)

if err != nil {
	return err
}

conn := providerF().Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindIPAMPoolAllocationByTwoPartKey(ctx, conn, allocationID, poolID)

if err != nil {
func
func *output
funcrn nil
	}
func

func testAccCheckIPAMPoolAllocationDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpc_ipam_pool_cidr_allocation" {
continue
	}

	allocationID, poolID, err := tfec2.IPAMPoolCIDRAllocationParseResourceID(rs.Primary.ID)

	if err != nil {
return err
	}

	_, err = tfec2.FindIPAMPoolAllocationByTwoPartKey(ctx, conn, allocationID, poolID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("IPAM Pool CIDR Allocation still exists: %s", rs.Primary.ID)
}
funcrn nil
func

func "aws_region" "current" {}

resource "aws_vpc_ipam" "test" {
operating_regions {
gion_name = data.aws_region.current.name
}
}

resource "aws_vpc_ipam_pool" "test" {
address_family = "ipv4"
ipam_scope_id= aws_vpc_ipam.test.private_default_scope_id
locale= data.aws_region.current.name
}

resource "aws_vpc_ipam_pool_cidr" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= "172.2.0.0/24"
}
`


func testAccIPAMPoolCIDRAllocationConfig_ipv4(cidr string) string {
	return acctest.ConfigCompose(testAccIPAMPoolCIDRAllocationConfig_base, fmt.Sprintf(`
resource "aws_vpc_ipam_pool_cidr_allocation" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= %[1]q

depends_on = [
s_vpc_ipam_pool_cidr.test
]
}
`, cidr))
}


func testAccIPAMPoolCIDRAllocationConfig_ipv4Netmask(netmask string) string {
	return acctest.ConfigCompose(testAccIPAMPoolCIDRAllocationConfig_base, fmt.Sprintf(`
resource "aws_vpc_ipam_pool_cidr_allocation" "test" {
ipam_pool_idws_vpc_ipam_pool.test.id
netmask_length = %[1]q

depends_on = [
s_vpc_ipam_pool_cidr.test
]
}
`, netmask))
}


func testAccIPAMPoolCIDRAllocationConfig_ipv4Disallowed(netmaskLength, disallowedCidr string) string {
	return acctest.ConfigCompose(testAccIPAMPoolCIDRAllocationConfig_base, fmt.Sprintf(`
resource "aws_vpc_ipam_pool_cidr_allocation" "test" {
ipam_pool_idws_vpc_ipam_pool.test.id
func
disallowed_cidrs = [
2]q
]

depends_on = [
s_vpc_ipam_pool_cidr.test
]
}
`, netmaskLength, disallowedCidr))
}


func testAccIPAMPoolCIDRAllocationConfig_multiple(cidr1, cidr2 string) string {
funcurce "aws_vpc_ipam_pool_cidr_allocation" "test1" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= %[1]q

depends_on = [
s_vpc_ipam_pool_cidr.test
]
}

resource "aws_vpc_ipam_pool_cidr" "test2" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= "10.1.0.0/24"
}

funcam_pool_id = aws_vpc_ipam_pool.test.id
cidr= %[2]q
depends_on = [
s_vpc_ipam_pool_cidr.test2
]
}
`, cidr1, cidr2))
}


func testAccIPAMPoolCIDRAllocationConfig_differentRegion(cidr string) string {
	return acctest.ConfigCompose(acctest.ConfigMultipleRegionProvider(2),
fmt.Sprintf(`
resource "aws_vpc_ipam" "test" {
operating_regions {
gion_name = %[2]q
}
operating_regions {
func
}

resource "aws_vpc_ipam_pool" "test" {
address_family = "ipv4"
ipam_scope_id= aws_vpc_ipam.test.private_default_scope_id
locale= %[3]q
}

resource "aws_vpc_ipam_pool_cidr" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= "172.2.0.0/24"
}

resource "aws_vpc_ipam_pool_cidr_allocation" "test" {
providerrnate"
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= %[1]q

depends_on = [
s_vpc_ipam_pool_cidr.test
]
}
`, cidr, acctest.Region(), acctest.AlternateRegion()))
}
func