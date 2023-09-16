// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"os"
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
	var conf ec2.Address
	resourceName := "aws_eip.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_basic,
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
funcource.TestCheckResourceAttrSet(resourceName, "public_ip"),
	testAccCheckEIPPublicDNS(resourceName),
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


func TestAccEC2EIP_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
func
Config: testAccEIPConfig_basic,
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceEIP(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccEC2EIP_migrateVPCToDomain(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
	resourceName := "aws_eip.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:test.ErrorCheck(t, ec2.EndpointsID),
CheckDestroy: testAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
ExternalProviders: map[string]resource.ExternalProvider{
	"aws": {
funcionConstraint: "4.67.0",
	},
},
Config: testAccEIPConfig_vpc,
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "domain", "vpc"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
	testAccCheckEIPPublicDNS(resourceName),
),
	},
	{
funcig:stAccEIPConfig_basic,
PlanOnly:true,
	},
},
	})
}


func TestAccEC2EIP_noVPC(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
	resourceName := "aws_eip.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_noVPC,
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
funcource.TestCheckResourceAttrSet(resourceName, "public_ip"),
	testAccCheckEIPPublicDNS(resourceName),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}


func TestAccEC2EIP_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
	resourceName := "aws_eip.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccEIPConfig_tags1("key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccEIPConfig_tags2("key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
functAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccEIPConfig_tags1("key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
func
	})
}


func TestAccEC2EIP_instance(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
	instanceResourceName := "aws_instance.test"
	resourceName := "aws_eip.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_instance(rName),
Check: resource.ComposeTestCheck
functAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrPair(resourceName, "instance", instanceResourceName, "id"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}

// Regression test for https://github.com/hashicorp/terraform/issues/3429 (now
// https://github.com/hashicorp/terraform-provider-aws/issues/42)

func TestAccEC2EIP_Instance_reassociate(t *testing.T) {
func conf ec2.Address
	instanceResourceName := "aws_instance.test"
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_instanceReassociate(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrPair(resourceName, "instance", instanceResourceName, "id"),
func
	{
Config: testAccEIPConfig_instanceReassociate(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrPair(resourceName, "instance", instanceResourceName, "id"),
),
Taint: []string{resourceName},
func
	})
}

// This test is an expansion of TestAccEC2EIP_Instance_associatedUserPrivateIP, by testing the
// associated Private EIPs of two instances

func TestAccEC2EIP_Instance_associatedUserPrivateIP(t *testing.T) {
func conf ec2.Address
	instance1ResourceName := "aws_instance.test.1"
	instance2ResourceName := "aws_instance.test.0"
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_instanceAssociated(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrSet(resourceName, "association_id"),
	resource.TestCheckResourceAttrPair(resourceName, "instance", instance1ResourceName, "id"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"associate_with_private_ip"},
	},
	{
Config: testAccEIPConfig_instanceAssociatedSwitch(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttrSet(resourceName, "association_id"),
	resource.TestCheckResourceAttrPair(resourceName, "instance", instance2ResourceName, "id"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
	},
},
	})
}
func
func TestAccEC2EIP_Instance_notAssociated(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
	instanceResourceName := "aws_instance.test"
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "association_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance", ""),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
Config: testAccEIPConfig_instance(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrSet(resourceName, "association_id"),
	resource.TestCheckResourceAttrPair(resourceName, "instance", instanceResourceName, "id"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
func
	})
}


func TestAccEC2EIP_networkInterface(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_networkInterface(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	testAccCheckEIPPrivateDNS(resourceName),
funcource.TestCheckResourceAttrSet(resourceName, "association_id"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
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
func TestAccEC2EIP_NetworkInterface_twoEIPsOneInterface(t *testing.T) {
	ctx := acctest.Context(t)
	var one, two ec2.Address
	resource1Name := "aws_eip.test.0"
	resource2Name := "aws_eip.test.1"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_multiNetworkInterface(rName),
func(
	testAccCheckEIPExists(ctx, resource1Name, &one),
	resource.TestCheckResourceAttrSet(resource1Name, "association_id"),
	resource.TestCheckResourceAttrSet(resource1Name, "public_ip"),

	testAccCheckEIPExists(ctx, resource2Name, &two),
	resource.TestCheckResourceAttrSet(resource2Name, "association_id"),
	resource.TestCheckResourceAttrSet(resource2Name, "public_ip"),
),
	},
},
	})
}


func TestAccEC2EIP_association(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
funcResourceName := "aws_network_interface.test"
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccEIPConfig_associationNone(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "association_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance", ""),
funcource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
	},
	{
Config: testAccEIPConfig_associationENI(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrSet(resourceName, "association_id"),
	resource.TestCheckResourceAttr(resourceName, "instance", ""),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface", eniResourceName, "id"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrSet(resourceName, "association_id"),
	resource.TestCheckResourceAttrPair(resourceName, "instance", instanceResourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "network_interface", instanceResourceName, "primary_network_interface_id"),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
	},
},
func


func TestAccEC2EIP_PublicIPv4Pool_default(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_publicIPv4PoolDefault(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttrSet(resourceName, "public_ip"),
	resource.TestCheckResourceAttr(resourceName, "public_ipv4_pool", "amazon"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccEC2EIP_PublicIPv4Pool_custom(t *testing.T) {
	ctx := acctest.Context(t)
	key := "AWS_EC2_EIP_PUBLIC_IPV4_POOL"
	poolName := os.Getenv(key)
	if poolName == "" {
t.Skipf("Environment variable %s is not set", key)
	}

	var conf ec2.Address
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
	resource.TestCheckResourceAttr(resourceName, "public_ipv4_pool", poolName),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccEC2EIP_customerOwnedIPv4Pool(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_customerOwnedIPv4Pool(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestMatchResourceAttr(resourceName, "customer_owned_ipv4_pool", regexache.MustCompile(`^ipv4pool-coip-.+$`)),
	resource.TestMatchResourceAttr(resourceName, "customer_owned_ip", regexache.MustCompile(`\d+\.\d+\.\d+\.\d+`)),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func := acctest.Context(t)
	var conf ec2.Address
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_networkBorderGroup(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "network_border_group", acctest.Region()),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
	resource.TestCheckResourceAttr(resourceName, "public_ipv4_pool", "amazon"),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccEC2EIP_carrierIP(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckWavelengthZoneAvailable(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_carrierIP(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrSet(resourceName, "carrier_ip"),
funcource.TestCheckResourceAttr(resourceName, "public_ip", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
func
}


func TestAccEC2EIP_BYOIPAddress_default(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.Address
	resourceName := "aws_eip.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_byoipAddressCustomDefault(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
),
	},
},
func


func TestAccEC2EIP_BYOIPAddress_custom(t *testing.T) {
	ctx := acctest.Context(t)
	key := "AWS_EC2_EIP_BYOIP_ADDRESS"
	address := os.Getenv(key)
	if address == "" {
func

	var conf ec2.Address
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_byoipAddressCustom(rName, address),
Check: resource.ComposeTestCheck
func(
	testAccCheckEIPExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "public_ip", address),
),
	},
},
	})
}


func := acctest.Context(t)
	key := "AWS_EC2_EIP_BYOIP_ADDRESS"
	address := os.Getenv(key)
	if address == "" {
t.Skipf("Environment variable %s is not set", key)
	}

	key = "AWS_EC2_EIP_PUBLIC_IPV4_POOL"
funcpoolName == "" {
t.Skipf("Environment variable %s is not set", key)
	}

	var conf ec2.Address
	resourceName := "aws_eip.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEIPDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEIPConfig_byoipAddressCustomPublicIPv4Pool(rName, address, poolName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "public_ip", address),
	resource.TestCheckResourceAttr(resourceName, "public_ipv4_pool", poolName),
),
	},
},
	})
}


func testAccCheckEIPExists(ctx context.Context, n string, v *ec2.Address) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
funcurn fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 EIP ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
funcut, err := tfec2.FindEIPByAllocationID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output

return nil
	}
func

func testAccCheckEIPDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_eip" {
continue
	}

	_, err := tfec2.FindEIPByAllocationID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
func

	return fmt.Errorf("EC2 EIP %s still exists", rs.Primary.ID)
}

return nil
	}
}
func
func testAccCheckEIPPrivateDNS(resourceName string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return fmt.Errorf("Not found: %s", resourceName)
}

privateDNS := rs.Primary.Attributes["private_dns"]
func-%s.%s",
funcc2.RegionalPrivateDNSSuffix(acctest.Region()),
)
funcrivateDNS != expectedPrivateDNS {
	return fmt.Errorf("expected private_dns value (%s), received: %s", expectedPrivateDNS, privateDNS)
}

return nil
	}
}


func testAccCheckEIPPublicDNS(resourceName string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return fmt.Errorf("Not found: %s", resourceName)
}

publicDNS := rs.Primary.Attributes["public_dns"]
expectedPublicDNS := fmt.Sprintf(
	"ec2-%s.%s.%s",
	tfec2.ConvertIPToDashIP(rs.Primary.Attributes["public_ip"]),
	tfec2.RegionalPublicDNSSuffix(acctest.Region()),
	acctest.PartitionDNSSuffix(),
)
funcublicDNS != expectedPublicDNS {
func

func
}

const testAccEIPConfig_basic = `
resource "aws_eip" "test" {
domain = "vpc"
}
`

const testAccEIPConfig_vpc = `
resource "aws_eip" "test" {
vpc = true
}
`

const testAccEIPConfig_noVPC = `
resource "aws_eip" "test" {
}
`


func testAccEIPConfig_tags1(tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_eip" "test" {
vpc = true

func= %[2]q
func
`, tagKey1, tagValue1)
func

func testAccEIPConfig_tags2(tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

tags = {
1]q = %[2]q
3]q = %[4]q
}
}
`, tagKey1, tagValue1, tagKey2, tagValue2)
}


func testAccEIPConfig_baseInstance(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigVPCWithSubnets(rName, 1),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
funcc_id = aws_vpc.test.id
funcgs = {
me = %[1]q
func

resource "aws_instance" "test" {
ami= data.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type
subnet_idet.test[0].id

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccEIPConfig_instance(rName string) string {
	return acctest.ConfigCompose(testAccEIPConfig_baseInstance(rName), fmt.Sprintf(`
resource "aws_eip" "test" {
instance = aws_instance.test.id
domainvpc"

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccEIPConfig_instanceReassociate(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigVPCWithSubnets(rName, 1),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_eip" "test" {
instance = aws_instance.test.id
domainvpc"

tags = {
func
}

resource "aws_instance" "test" {
ami = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
associate_public_ip_address = true
instance_type= data.aws_ec2_instance_type_offering.available.instance_type
subnet_idaws_subnet.test[0].id

tags = {
me = %[1]q
}

func_before_destroy = true
}
}

resource "aws_internet_gateway" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id
funcute {
dr_block = "0.0.0.0/0"
teway_id = aws_internet_gateway.test.id
}

tags = {
me = %[1]q
}
}

resource "aws_route_table_association" "test" {
subnet_idnet.test[0].id
route_table_id = aws_route_table.test.id
}
`, rName))
}


func testAccEIPConfig_baseInstanceAssociated(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block= "10.0.0.0/16"
enable_dns_hostnames = true
funcgs = {
me = %[1]q
}
}

resource "aws_internet_gateway" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
funcailability_zonews_availability_zones.available.names[0]
cidr_block0/24"
map_public_ip_on_launch = true

depends_on = [aws_internet_gateway.test]

tags = {
me = %[1]q
}
}

resource "aws_instance" "test" {
count = 2

ami= data.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type

private_ip = "10.0.0.1${count.index}"
subnet_id= aws_subnet.test.id

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccEIPConfig_instanceAssociated(rName string) string {
	return acctest.ConfigCompose(testAccEIPConfig_baseInstanceAssociated(rName), fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

instancews_instance.test[1].id
associate_with_private_ip = aws_instance.test[1].private_ip

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccEIPConfig_instanceAssociatedSwitch(rName string) string {
	return acctest.ConfigCompose(testAccEIPConfig_baseInstanceAssociated(rName), fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

instancews_instance.test[0].id
associate_with_private_ip = aws_instance.test[0].private_ip

tags = {
me = %[1]q
}
}
`, rName))
}

func testAccEIPConfig_instanceAssociateNotAssociated(rName string) string {
	return acctest.ConfigCompose(testAccEIPConfig_baseInstance(rName), fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccEIPConfig_networkInterface(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_internet_gateway" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_network_interface" "test" {
subnet_idbnet.test[0].id
private_ips.10"]
security_groups = [aws_vpc.test.default_security_group_id]

tags = {
me = %[1]q
}
}

resource "aws_eip" "test" {
domainvpc"
network_interface = aws_network_interface.test.id

tags = {
me = %[1]q
}

depends_on = [aws_internet_gateway.test]
}
`, rName))
}


func testAccEIPConfig_multiNetworkInterface(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_internet_gateway" "test" {
vpc_id = aws_vpc.test.id

tags = {
func
}

resource "aws_network_interface" "test" {
subnet_idbnet.test[0].id
private_ips.10", "10.0.0.11"]
security_groups = [aws_vpc.test.default_security_group_id]

tags = {
me = %[1]q
}
}

resource "aws_eip" "test" {
count = 2

functwork_interface= aws_network_interface.test.id
associate_with_private_ip = "10.0.0.1${count.index}"

tags = {
me = %[1]q
}

depends_on = [aws_internet_gateway.test]
}
`, rName))
}


func testAccEIPConfig_baseAssociation(rName string) string {
	return acctest.ConfigCompose(testAccEIPConfig_baseInstance(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
funcivate_ips.10"]
security_groups = [aws_vpc.test.default_security_group_id]

tags = {
me = %[1]q
}

depends_on = [aws_internet_gateway.test]
}
`, rName))
}


funcurn acctest.ConfigCompose(testAccEIPConfig_baseAssociation(rName), fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

tags = {
me = %[1]q
}

depends_on = [aws_network_interface.test]
}
`, rName))
}


func testAccEIPConfig_associationENI(rName string) string {
	return acctest.ConfigCompose(testAccEIPConfig_baseAssociation(rName), fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

tags = {
me = %[1]q
}

network_interface = aws_network_interface.test.id
}
`, rName))
}


func testAccEIPConfig_associationInstance(rName string) string {
	return acctest.ConfigCompose(testAccEIPConfig_baseAssociation(rName), fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

func %[1]q
}

instance = aws_instance.test.id

depends_on = [aws_network_interface.test]
}
`, rName))
}


func testAccEIPConfig_publicIPv4PoolDefault(rName string) string {
	return fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccEIPConfig_publicIPv4PoolCustom(rName, poolName string) string {
	return fmt.Sprintf(`
resource "aws_eip" "test" {
domain= "vpc"
public_ipv4_pool = %[2]q

tags = {
me = %[1]q
}
}
`, rName, poolName)
}

func testAccEIPConfig_customerOwnedIPv4Pool(rName string) string {
	return fmt.Sprintf(`
data "aws_ec2_coip_pools" "test" {}

resource "aws_eip" "test" {
customer_owned_ipv4_pool = tolist(data.aws_ec2_coip_pools.test.pool_ids)[0]
domain"vpc"

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccEIPConfig_networkBorderGroup(rName string) string {
func "aws_region" current {}

resource "aws_eip" "test" {
domain= "vpc"
network_border_group = data.aws_region.current.name

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccEIPConfig_carrierIP(rName string) string {
funcAccAvailableAZsWavelengthZonesDefaultExcludeConfig(),
fmt.Sprintf(`
data "aws_availability_zone" "available" {
name = data.aws_availability_zones.available.names[0]
}

resource "aws_eip" "test" {
domain= "vpc"
network_border_group = data.aws_availability_zone.available.network_border_group

tags = {
me = %[1]q
}
}
`, rName))
func

func testAccEIPConfig_byoipAddressCustomDefault(rName string) string {
	return fmt.Sprintf(`
resource "aws_eip" "test" {
domain = "vpc"

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccEIPConfig_byoipAddressCustom(rName, address string) string {
	return fmt.Sprintf(`
funcmain= "vpc"
address = %[2]q

tags = {
me = %[1]q
}
}
`, rName, address)
}


func testAccEIPConfig_byoipAddressCustomPublicIPv4Pool(rName, address, poolName string) string {
	return fmt.Sprintf(`
funcmain= "vpc"
address = %[2]q
public_ipv4_pool = %[3]q

tags = {
me = %[1]q
}
}
`, rName, address, poolName)
}
funcfuncfuncfuncfuncfunc