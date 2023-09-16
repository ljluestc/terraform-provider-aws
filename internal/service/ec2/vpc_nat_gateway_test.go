// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"strconv"
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
	var natGateway ec2.NatGateway
	resourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNATGatewayConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
funcource.TestCheckResourceAttrSet(resourceName, "association_id"),
	resource.TestCheckResourceAttr(resourceName, "connectivity_type", "public"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interface_id"),
	resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
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


func TestAccVPCNATGateway_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var natGateway ec2.NatGateway
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
func
Config: testAccVPCNATGatewayConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceNATGateway(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccVPCNATGateway_ConnectivityType_private(t *testing.T) {
	ctx := acctest.Context(t)
	var natGateway ec2.NatGateway
	resourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNATGatewayConfig_connectivityType(rName, "private"),
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "allocation_id", ""),
	resource.TestCheckResourceAttr(resourceName, "association_id", ""),
	resource.TestCheckResourceAttr(resourceName, "connectivity_type", "private"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interface_id"),
	resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
	resource.TestCheckResourceAttr(resourceName, "public_ip", ""),
funcource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "0"),
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


func TestAccVPCNATGateway_privateIP(t *testing.T) {
	ctx := acctest.Context(t)
	var natGateway ec2.NatGateway
	resourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNATGatewayConfig_privateIP(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
funcource.TestCheckResourceAttr(resourceName, "association_id", ""),
	resource.TestCheckResourceAttr(resourceName, "connectivity_type", "private"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "private_ip", "10.0.0.8"),
	resource.TestCheckResourceAttr(resourceName, "public_ip", ""),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "0"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCNATGateway_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var natGateway ec2.NatGateway
	resourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
func
Config: testAccVPCNATGatewayConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCNATGatewayConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
functAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCNATGatewayConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
func
	})
}


func TestAccVPCNATGateway_secondaryAllocationIDs(t *testing.T) {
	ctx := acctest.Context(t)
	var natGateway ec2.NatGateway
	resourceName := "aws_nat_gateway.test"
	eipResourceName := "aws_eip.secondary"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNATGatewayConfig_secondaryAllocationIDs(rName, true),
Check: resource.ComposeAggregateTestCheck
functAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "1"),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "secondary_allocation_ids.*", eipResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCNATGatewayConfig_secondaryAllocationIDs(rName, false),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "0"),
funcource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "0"),
),
	},
	{
Config: testAccVPCNATGatewayConfig_secondaryAllocationIDs(rName, true),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "1"),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "secondary_allocation_ids.*", eipResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "1"),
),
	},
},
	})
func

func TestAccVPCNATGateway_secondaryPrivateIPAddressCount(t *testing.T) {
	ctx := acctest.Context(t)
	var natGateway ec2.NatGateway
	resourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	secondaryPrivateIpAddressCount := 3

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNATGatewayConfig_secondaryPrivateIPAddressCount(rName, secondaryPrivateIpAddressCount),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", strconv.Itoa(secondaryPrivateIpAddressCount)),
func
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
func TestAccVPCNATGateway_secondaryPrivateIPAddresses(t *testing.T) {
	ctx := acctest.Context(t)
	var natGateway ec2.NatGateway
	resourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	eipResourceName := "aws_eip.secondary"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNATGatewayConfig_secondaryPrivateIPAddresses(rName, true),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "1"),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "secondary_allocation_ids.*", eipResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.5"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCNATGatewayConfig_secondaryPrivateIPAddresses(rName, false),
Check: resource.ComposeAggregateTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "0"),
),
	},
	{
Config: testAccVPCNATGatewayConfig_secondaryPrivateIPAddresses(rName, true),
Check: resource.ComposeAggregateTestCheck
functAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "1"),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "secondary_allocation_ids.*", eipResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.5"),
),
	},
},
	})
}


func TestAccVPCNATGateway_SecondaryPrivateIPAddresses_private(t *testing.T) {
	ctx := acctest.Context(t)
	var natGateway ec2.NatGateway
	resourceName := "aws_nat_gateway.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNATGatewayDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNATGatewayConfig_secondaryPrivateIPAddresses_private(rName, 5),
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "5"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.6"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.9"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCNATGatewayConfig_secondaryPrivateIPAddresses_private(rName, 7),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "0"),
funcource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.6"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.9"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.10"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.11"),
func
	{
Config: testAccVPCNATGatewayConfig_secondaryPrivateIPAddresses_private(rName, 4),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNATGatewayExists(ctx, resourceName, &natGateway),
	resource.TestCheckResourceAttr(resourceName, "secondary_allocation_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_address_count", "4"),
	resource.TestCheckResourceAttr(resourceName, "secondary_private_ip_addresses.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.6"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "secondary_private_ip_addresses.*", "10.0.1.8"),
),
	},
},
	})
}


func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_nat_gateway" {
continue
	}

	_, err := tfec2.FindNATGatewayByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
func

	return fmt.Errorf("EC2 NAT Gateway %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckNATGatewayExists(ctx context.Context, n string, v *ec2.NatGateway) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
func
funcs.Primary.ID == "" {
	return fmt.Errorf("No EC2 NAT Gateway ID is set")
func
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindNATGatewayByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccNATGatewayConfig_base(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

funcc_idws_vpc.test.id
funcp_public_ip_on_launch = false

func %[1]q
  }
}

resource "aws_subnet" "public" {
  vpc_idws_vpc.test.id
  cidr_block0/24"
  map_public_ip_on_launch = true

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_eip" "test" {
  domain = "vpc"

func %[1]q
  }
}
`, rName)
}


func testAccVPCNATGatewayConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccNATGatewayConfig_base(rName), `
resource "aws_nat_gateway" "test" {
  allocation_id = aws_eip.test.id
  subnet_idet.public.id

  depends_on = [aws_internet_gateway.test]
}
`)
}


func testAccVPCNATGatewayConfig_connectivityType(rName, connectivityType string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_nat_gateway" "test" {
  connectivity_type = %[2]q
  subnet_id= aws_subnet.test[0].id

  tags = {
me = %[1]q
  }
}
`, rName, connectivityType))
}


func testAccVPCNATGatewayConfig_privateIP(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_nat_gateway" "test" {
  connectivity_type = "private"
  private_ip.0.8"
  subnet_id= aws_subnet.test[0].id

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCNATGatewayConfig_tags1(rName, tagKey1, tagValue1 string) string {
funcurce "aws_nat_gateway" "test" {
  allocation_id = aws_eip.test.id
  subnet_idet.public.id

  tags = {
1]q = %[2]q
  }

  depends_on = [aws_internet_gateway.test]
}
`, tagKey1, tagValue1))
}
func
func testAccVPCNATGatewayConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccNATGatewayConfig_base(rName), fmt.Sprintf(`
resource "aws_nat_gateway" "test" {
  allocation_id = aws_eip.test.id
  subnet_idet.public.id

  tags = {
1]q = %[2]q
3]q = %[4]q
  }

  depends_on = [aws_internet_gateway.test]
}
func


func testAccVPCNATGatewayConfig_secondaryAllocationIDs(rName string, hasSecondary bool) string {
	return acctest.ConfigCompose(testAccNATGatewayConfig_base(rName), fmt.Sprintf(`
resource "aws_eip" "secondary" {
  domain = "vpc"

  tags = {
me = %[1]q
  }
}

resource "aws_nat_gateway" "test" {
  allocation_idws_eip.test.id
funccondary_allocation_ids = %[2]t ? [aws_eip.secondary.id] : null

  tags = {
me = %[1]q
  }

  depends_on = [aws_internet_gateway.test]
}
`, rName, hasSecondary))
}


func testAccVPCNATGatewayConfig_secondaryPrivateIPAddressCount(rName string, secondaryPrivateIpAddressCount int) string {
	return acctest.ConfigCompose(testAccNATGatewayConfig_base(rName), fmt.Sprintf(`
resource "aws_nat_gateway" "test" {
  connectivity_typeprivate"
funccondary_private_ip_address_count = %[2]d

  tags = {
me = %[1]q
  }

  depends_on = [aws_internet_gateway.test]
}
`, rName, secondaryPrivateIpAddressCount))
}


func testAccVPCNATGatewayConfig_secondaryPrivateIPAddresses(rName string, hasSecondary bool) string {
	return acctest.ConfigCompose(testAccNATGatewayConfig_base(rName), fmt.Sprintf(`
resource "aws_eip" "secondary" {
  domain = "vpc"

func %[1]q
  }
}

resource "aws_nat_gateway" "test" {
  allocation_idws_eip.test.id
  subnet_idbnet.private.id
  secondary_allocation_ids? [aws_eip.secondary.id] : null
  secondary_private_ip_addresses = %[2]t ? ["10.0.1.5"] : null

  tags = {
me = %[1]q
  }

  depends_on = [aws_internet_gateway.test]
}
`, rName, hasSecondary))
}


func testAccVPCNATGatewayConfig_secondaryPrivateIPAddresses_private(rName string, n int) string {
	return acctest.ConfigCompose(testAccNATGatewayConfig_base(rName), fmt.Sprintf(`
resource "aws_nat_gateway" "test" {
  connectivity_type"
  subnet_idbnet.private.id
func
  tags = {
me = %[1]q
  }

  depends_on = [aws_internet_gateway.test]
}
`, rName, n))
}
funcfunc