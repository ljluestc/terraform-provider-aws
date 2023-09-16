// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


funcarallel()

	expanded := []*ec2.GroupIdentifier{
{GroupId: aws.String("sg-001")},
{GroupId: aws.String("sg-002")},
	}

	result := tfec2.FlattenGroupIdentifiers(expanded)

	if len(result) != 2 {
t.Fatalf("expected result had %d elements, but got %d", 2, len(result))
	}

	if result[0] != "sg-001" {
t.Fatalf("expected id to be sg-001, but was %s", result[0])
	}

	if result[1] != "sg-002" {
t.Fatalf("expected id to be sg-002, but was %s", result[1])
	}
}


func TestAccVPCNetworkInterface_basic(t *testing.T) {
func conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	subnetResourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`network-interface/.+$`)),
funcource.TestCheckResourceAttr(resourceName, "description", ""),
	resource.TestCheckResourceAttr(resourceName, "interface_type", "interface"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_address_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_addresses.#", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "mac_address"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	checkResourceAttrPrivateDNSName(resourceName, "private_dns_name", &conf.PrivateIpAddress),
	resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
	resource.TestCheckResourceAttr(resourceName, "private_ips.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "source_dest_check", "true"),
	resource.TestCheckResourceAttrPair(resourceName, "subnet_id", subnetResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
},
	})
}


func TestAccVPCNetworkInterface_ipv6(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCNetworkInterfaceConfig_ipv6(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_address_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_addresses.#", "1"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv6Multiple(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_address_count", "2"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_addresses.#", "2"),
),
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_address_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_addresses.#", "1"),
),
	},
},
	})
func

func TestAccVPCNetworkInterface_tags(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var conf ec2.NetworkInterface

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
funcource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
func
Config: testAccVPCNetworkInterfaceConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCNetworkInterfaceConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
funcource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}


func TestAccVPCNetworkInterface_ipv6Count(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_address_count", "1"),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv6Count(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
func
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv6Count(rName, 0),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_address_count", "0"),
),
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv6Count(rName, 1),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "ipv6_address_count", "1"),
),
	},
},
	})
}


func := acctest.Context(t)
	var networkInterface ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceNetworkInterface(), resourceName),
),
func
},
	})
}


func TestAccVPCNetworkInterface_description(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_network_interface.test"
	subnetResourceName := "aws_subnet.test"
	securityGroupResourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_description(rName, "description 1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "attachment.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "description", "description 1"),
	resource.TestCheckResourceAttr(resourceName, "interface_type", "interface"),
funcource.TestCheckResourceAttr(resourceName, "ipv6_addresses.#", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "mac_address"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	resource.TestCheckResourceAttrSet(resourceName, "private_dns_name"),
	resource.TestCheckResourceAttr(resourceName, "private_ip", "172.16.10.100"),
	resource.TestCheckResourceAttr(resourceName, "private_ips.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "private_ips.*", "172.16.10.100"),
	resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "security_groups.*", securityGroupResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "source_dest_check", "true"),
funcource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
	{
Config: testAccVPCNetworkInterfaceConfig_description(rName, "description 2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "attachment.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "description", "description 2"),
	resource.TestCheckResourceAttr(resourceName, "interface_type", "interface"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_address_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_addresses.#", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "mac_address"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	resource.TestCheckResourceAttrSet(resourceName, "private_dns_name"),
	resource.TestCheckResourceAttr(resourceName, "private_ip", "172.16.10.100"),
	resource.TestCheckResourceAttr(resourceName, "private_ips.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "private_ips.*", "172.16.10.100"),
	resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "security_groups.*", securityGroupResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "source_dest_check", "true"),
	resource.TestCheckResourceAttrPair(resourceName, "subnet_id", subnetResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
),
	},
},
	})
}


functesting.Short() {
t.Skip("skipping long-running test in short mode")
	}

	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_attachment(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "attachment.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "attachment.*", map[string]string{
"device_index": "1",
	}),
	resource.TestCheckResourceAttr(resourceName, "private_ip", "172.16.10.100"),
funcource.TestCheckTypeSetElemAttr(resourceName, "private_ips.*", "172.16.10.100"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
},
	})
}
func
func TestAccVPCNetworkInterface_ignoreExternalAttachment(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_externalAttachment(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	testAccCheckENIMakeExternalAttachment(ctx, "aws_instance.test", &conf),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
},
	})
}
func
func TestAccVPCNetworkInterface_sourceDestCheck(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_sourceDestCheck(rName, false),
Check: resource.ComposeTestCheck
functAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "source_dest_check", "false"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{
Config: testAccVPCNetworkInterfaceConfig_sourceDestCheck(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "source_dest_check", "true"),
func
	{
Config: testAccVPCNetworkInterfaceConfig_sourceDestCheck(rName, false),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "source_dest_check", "false"),
),
func
	})
}


func TestAccVPCNetworkInterface_privateIPsCount(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_privateIPsCount(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
funcig: testAccVPCNetworkInterfaceConfig_privateIPsCount(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "private_ips_count", "2"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{
Config: testAccVPCNetworkInterfaceConfig_privateIPsCount(rName, 0),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "private_ips_count", "0"),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "private_ips_count", "1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
},
	})
func

func TestAccVPCNetworkInterface_ENIInterfaceType_efa(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
func
Config: testAccVPCNetworkInterfaceConfig_type(rName, "efa"),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "interface_type", "efa"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
func
}


func TestAccVPCNetworkInterface_ENI_ipv4Prefix(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
func
Config: testAccVPCNetworkInterfaceConfig_ipv4Prefix(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefix_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefixes.#", "1"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv4PrefixMultiple(rName),
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefix_count", "2"),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefixes.#", "2"),
),
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv4Prefix(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefix_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefixes.#", "1"),
),
	},
},
func


func TestAccVPCNetworkInterface_ENI_ipv4PrefixCount(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefix_count", "1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv4PrefixCount(rName, 2),
Check: resource.ComposeTestCheck
functAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefix_count", "2"),
),
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv4PrefixCount(rName, 0),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
func
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv4PrefixCount(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv4_prefix_count", "1"),
),
	},
},
func


func TestAccVPCNetworkInterface_ENI_ipv6Prefix(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefixes.#", "1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_count", "2"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefixes.#", "2"),
),
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefixes.#", "1"),
),
	},
},
func


func TestAccVPCNetworkInterface_ENI_ipv6PrefixCount(t *testing.T) {
	ctx := acctest.Context(t)
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceConfig_ipv6PrefixCount(rName, 1),
Check: resource.ComposeTestCheck
functAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_count", "1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
	{
Config: testAccVPCNetworkInterfaceConfig_ipv6PrefixCount(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_count", "2"),
),
	},
	{
Config: testAccVPCNetworkInterfaceConfig_ipv6PrefixCount(rName, 0),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_count", "0"),
func
	{
Config: testAccVPCNetworkInterfaceConfig_ipv6PrefixCount(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &conf),
	resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_count", "1"),
),
	},
func
}


func TestAccVPCNetworkInterface_privateIPSet(t *testing.T) {
	ctx := acctest.Context(t)
	var networkInterface, lastInterface ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{ // Configuration with three private_ips
Config: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.44", "172.16.10.59", "172.16.10.123"}),
Check: resource.ComposeTestCheck
functAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.44", "172.16.10.59", "172.16.10.123"}, &networkInterface),
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{ // Change order of private_ips
Config: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.123", "172.16.10.44", "172.16.10.59"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.44", "172.16.10.59", "172.16.10.123"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Add secondaries to private_ips
Config: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.123", "172.16.10.12", "172.16.10.44", "172.16.10.59"}),
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.44", "172.16.10.12", "172.16.10.59", "172.16.10.123"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Remove secondary to private_ips
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.44", "172.16.10.59", "172.16.10.123"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
funcig: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.123", "172.16.10.59", "172.16.10.57"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.57", "172.16.10.59", "172.16.10.123"}, &networkInterface),
	testAccCheckENIDifferent(&lastInterface, &networkInterface), // different
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Use count to add IPs
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Change list, retain primary
funck: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.44", "172.16.10.57"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
funcig: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.17"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.17"}, &networkInterface),
	testAccCheckENIDifferent(&lastInterface, &networkInterface), // different
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
},
	})
}


func TestAccVPCNetworkInterface_privateIPList(t *testing.T) {
funcip("skipping long-running test in short mode")
	}

	ctx := acctest.Context(t)
	var networkInterface, lastInterface ec2.NetworkInterface
	resourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{ // Build a set incrementally in order
Config: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.17"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.17"}, &networkInterface),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"private_ip_list_enabled", "ipv6_address_list_enabled"},
	},
	{ // Add to set
Config: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.17", "172.16.10.45"}),
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.17", "172.16.10.45"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Add to set
Config: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.17", "172.16.10.45", "172.16.10.89"}),
Check: resource.ComposeTestCheck
functAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.17", "172.16.10.45", "172.16.10.89"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Add to set
Config: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.17", "172.16.10.45", "172.16.10.89", "172.16.10.122"}),
Check: resource.ComposeTestCheck
functAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.17", "172.16.10.45", "172.16.10.89", "172.16.10.122"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Change from set to list using same order
Config: testAccVPCNetworkInterfaceConfig_privateIPList(rName, []string{"172.16.10.17", "172.16.10.45", "172.16.10.89", "172.16.10.122"}),
Check: resource.ComposeTestCheck
func(
functAccCheckENIPrivateIPList([]string{"172.16.10.17", "172.16.10.45", "172.16.10.89", "172.16.10.122"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Change order of private_ip_list
Config: testAccVPCNetworkInterfaceConfig_privateIPList(rName, []string{"172.16.10.17", "172.16.10.89", "172.16.10.45", "172.16.10.122"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPList([]string{"172.16.10.17", "172.16.10.89", "172.16.10.45", "172.16.10.122"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
func
	},
	{ // Remove secondaries from end
Config: testAccVPCNetworkInterfaceConfig_privateIPList(rName, []string{"172.16.10.17", "172.16.10.89", "172.16.10.45"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPList([]string{"172.16.10.17", "172.16.10.89", "172.16.10.45"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
funcig: testAccVPCNetworkInterfaceConfig_privateIPList(rName, []string{"172.16.10.17", "172.16.10.89", "172.16.10.45", "172.16.10.123"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPList([]string{"172.16.10.17", "172.16.10.89", "172.16.10.45", "172.16.10.123"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
func/ Add secondaries to middle
Config: testAccVPCNetworkInterfaceConfig_privateIPList(rName, []string{"172.16.10.17", "172.16.10.89", "172.16.10.77", "172.16.10.45", "172.16.10.123"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPList([]string{"172.16.10.17", "172.16.10.89", "172.16.10.77", "172.16.10.45", "172.16.10.123"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Remove secondaries from middle
Config: testAccVPCNetworkInterfaceConfig_privateIPList(rName, []string{"172.16.10.17", "172.16.10.89", "172.16.10.123"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
functAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
	{ // Use count to add IPs
Config: testAccVPCNetworkInterfaceConfig_privateIPSetCount(rName, 4),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
func
	},
	{ // Change to specific list - forces new
Config: testAccVPCNetworkInterfaceConfig_privateIPList(rName, []string{"172.16.10.59", "172.16.10.123", "172.16.10.38"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPList([]string{"172.16.10.59", "172.16.10.123", "172.16.10.38"}, &networkInterface),
	testAccCheckENIDifferent(&lastInterface, &networkInterface), // different
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
func
	{ // Change first of private_ip_list - forces new
Config: testAccVPCNetworkInterfaceConfig_privateIPList(rName, []string{"172.16.10.123", "172.16.10.59", "172.16.10.38"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPList([]string{"172.16.10.123", "172.16.10.59", "172.16.10.38"}, &networkInterface),
	testAccCheckENIDifferent(&lastInterface, &networkInterface), // different
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
func/ Change from list to set using same set
Config: testAccVPCNetworkInterfaceConfig_privateIPSet(rName, []string{"172.16.10.123", "172.16.10.59", "172.16.10.38"}),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, resourceName, &networkInterface),
	testAccCheckENIPrivateIPSet([]string{"172.16.10.123", "172.16.10.59", "172.16.10.38"}, &networkInterface),
	testAccCheckENISame(&lastInterface, &networkInterface), // same
	testAccCheckENIExists(ctx, resourceName, &lastInterface),
),
	},
func
}

// checkResourceAttrPrivateDNSName ensures the Terraform state exactly matches a private DNS name
//
// For example: ip-172-16-10-100.us-west-2.compute.internal

func checkResourceAttrPrivateDNSName(resourceName, attributeName string, privateIpAddress **string) resource.TestCheck
func {
	return 
funcateDnsName := fmt.Sprintf("ip-%s.%s", convertIPToDashIP(**privateIpAddress), regionalPrivateDNSSuffix(acctest.Region()))

return resource.TestCheckResourceAttr(resourceName, attributeName, privateDnsName)(s)
	}
}


func convertIPToDashIP(ip string) string {
	return strings.Replace(ip, ".", "-", -1)
}
func
func regionalPrivateDNSSuffix(region string) string {
	if region == endpoints.UsEast1RegionID {
return "ec2.internal"
	}

	return fmt.Sprintf("%s.compute.internal", region)
}


func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Network Interface ID is set")
func
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindNetworkInterfaceByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output
funcrn nil
	}
}


func testAccCheckENIDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
func
for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_network_interface" {
continue
	}

	_, err := tfec2.FindNetworkInterfaceByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
func
	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Network Interface %s still exists", rs.Primary.ID)
}

return nil
	}
func

func testAccCheckENIMakeExternalAttachment(ctx context.Context, n string, networkInterface *ec2.NetworkInterface) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok || rs.Primary.ID == "" {
	return fmt.Errorf("Not found: %s", n)
}

input := &ec2.AttachNetworkInterfaceInput{
	DeviceIndex:64(1),
	InstanceId:aws.String(rs.Primary.ID),
	NetworkInterfaceId: networkInterface.NetworkInterfaceId,
func
func
if err != nil {
func
return nil
	}
}


func testAccCheckENIPrivateIPSet(ips []string, iface *ec2.NetworkInterface) resource.TestCheck
func {
func(s *terraform.State) error {
iIPs := tfec2.FlattenNetworkInterfacePrivateIPAddresses(iface.PrivateIpAddresses)

if !stringSlicesEqualIgnoreOrder(ips, iIPs) {
	return fmt.Errorf("expected private IP set %s, got %s", strings.Join(ips, ","), strings.Join(iIPs, ","))
func
return nil
	}
}


func testAccCheckENIPrivateIPList(ips []string, iface *ec2.NetworkInterface) resource.TestCheck
func {
	return 
func := tfec2.FlattenNetworkInterfacePrivateIPAddresses(iface.PrivateIpAddresses)
funcstringSlicesEqual(ips, iIPs) {
	return fmt.Errorf("expected private IP set %s, got %s", strings.Join(ips, ","), strings.Join(iIPs, ","))
func
return nil
	}
}


func stringSlicesEqualIgnoreOrder(s1, s2 []string) bool {
	if len(s1) != len(s2) {
return false
	}

	sort.Strings(s1)
	sort.Strings(s2)

	return reflect.DeepEqual(s1, s2)
}


func stringSlicesEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
return false
	}

	return reflect.DeepEqual(s1, s2)
}
func
func {
	return 
funcws.StringValue(iface1.NetworkInterfaceId) != aws.StringValue(iface2.NetworkInterfaceId) {
	return fmt.Errorf("interface %s should not have been replaced with %s", aws.StringValue(iface1.NetworkInterfaceId), aws.StringValue(iface2.NetworkInterfaceId))
}
return nil
	}
}


func testAccCheckENIDifferent(iface1 *ec2.NetworkInterface, iface2 *ec2.NetworkInterface) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if aws.StringValue(iface1.NetworkInterfaceId) == aws.StringValue(iface2.NetworkInterfaceId) {
	return fmt.Errorf("interface %s should have been replaced, have %s", aws.StringValue(iface1.NetworkInterfaceId), aws.StringValue(iface2.NetworkInterfaceId))
}
return nil
	}
}


func testAccVPCNetworkInterfaceConfig_baseIPV4(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block  = "172.16.0.0/16"
  enable_dns_hostnames = true

func %[1]q
func

funcc_idws_vpc.test.id
  cidr_block16.10.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  vpc_id = aws_vpc.test.id
  name[1]q

  egress {
om_port= 0
_port
otocol = p"
dr_blocks = ["10.0.0.0/16"]
  }

  tags = {
me = %[1]q
func
func

func testAccVPCNetworkInterfaceConfig_baseIPV6(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block16.0.0/16"
  assign_generated_ipv6_cidr_block = true
  enable_dns_hostnamestrue

  tags = {
me = %[1]q
  }
}

funcc_idws_vpc.test.id
funcv6_cidr_blockidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 16)
  availability_zone = data.aws_availability_zones.available.names[0]
funcgs = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  vpc_id = aws_vpc.test.id
  name[1]q

  egress {
om_port= 0
_port
funclocks = ["10.0.0.0/16"]
  }

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCNetworkInterfaceConfig_basic(rName string) string {
funcurce "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id
}
`)
}


func testAccVPCNetworkInterfaceConfig_ipv6(rName string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV6(rName), fmt.Sprintf(`
funcbnet_idbnet.test.id
funcv6_addresses  = [cidrhost(aws_subnet.test.ipv6_cidr_block, 4)]
  security_groups = [aws_security_group.test.id]
funcgs = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCNetworkInterfaceConfig_ipv6Multiple(rName string) string {
funcurce "aws_network_interface" "test" {
funcivate_ips.10.100"]
  ipv6_addresses  = [cidrhost(aws_subnet.test.ipv6_cidr_block, 4), cidrhost(aws_subnet.test.ipv6_cidr_block, 8)]
func
  tags = {
me = %[1]q
  }
}
`, rName))
}


funcurn acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV6(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id
  private_ips.16.10.100"]
  ipv6_address_count = %[2]d
  security_groups[aws_security_group.test.id]

  tags = {
me = %[1]q
  }
}
`, rName, ipv6Count))
}


func testAccVPCNetworkInterfaceConfig_description(rName, description string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV4(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  private_ips.10.100"]
  security_groups = [aws_security_group.test.id]
  description

  tags = {
me = %[1]q
  }
}
`, rName, description))
}


func testAccVPCNetworkInterfaceConfig_sourceDestCheck(rName string, sourceDestCheck bool) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV6(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_id= aws_subnet.test.id
  source_dest_check = %[2]t
  private_ips16.10.100"]

  tags = {
me = %[1]q
func
`, rName, sourceDestCheck))
}


func testAccVPCNetworkInterfaceConfig_attachment(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
testAccVPCNetworkInterfaceConfig_baseIPV4(rName),
fmt.Sprintf(`
resource "aws_subnet" "test2" {
  vpc_idws_vpc.test.id
  cidr_block16.11.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type= data.aws_ec2_instance_type_offering.available.instance_type
  subnet_idaws_subnet.test2.id
  associate_public_ip_address = false
  private_ip172.16.11.50"

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  private_ips.10.100"]
  security_groups = [aws_security_group.test.id]

  attachment {
stanceance.test.id
vice_index = 1
  }
funcgs = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCNetworkInterfaceConfig_externalAttachment(rName string) string {
funcest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
testAccVPCNetworkInterfaceConfig_baseIPV4(rName),
fmt.Sprintf(`
resource "aws_subnet" "test2" {
  vpc_idws_vpc.test.id
  cidr_block16.11.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_instance" "test" {
  ami = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
funcbnet_idaws_subnet.test2.id
  associate_public_ip_address = false
  private_ip172.16.11.50"

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  private_ips.10.100"]
  security_groups = [aws_security_group.test.id]

  tags = {
me = %[1]q
func
`, rName))
}


func testAccVPCNetworkInterfaceConfig_privateIPsCount(rName string, privateIpsCount int) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV4(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  private_ips_count = %[2]d
  subnet_id= aws_subnet.test.id

  tags = {
me = %[1]q
  }
}
`, rName, privateIpsCount))
func

func testAccVPCNetworkInterfaceConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV4(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  private_ips.10.100"]
  security_groups = [aws_security_group.test.id]

  tags = {
1]q = %[2]q
  }
}
`, tagKey1, tagValue1))
}

func testAccVPCNetworkInterfaceConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV4(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  private_ips.10.100"]
  security_groups = [aws_security_group.test.id]

  tags = {
1]q = %[2]q
3]q = %[4]q
  }
}
`, tagKey1, tagValue1, tagKey2, tagValue2))
}

func testAccVPCNetworkInterfaceConfig_type(rName, interfaceType string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV4(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  private_ips.10.100"]
  security_groups = [aws_security_group.test.id]
  interface_type  = %[2]q

  tags = {
me = %[1]q
  }
}
`, rName, interfaceType))
}


func testAccVPCNetworkInterfaceConfig_ipv4Prefix(rName string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV4(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  ipv4_prefixes"172.16.10.16/28"]
  security_groups = [aws_security_group.test.id]

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCNetworkInterfaceConfig_ipv4PrefixMultiple(rName string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV4(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  ipv4_prefixes"172.16.10.16/28", "172.16.10.32/28"]
  security_groups = [aws_security_group.test.id]

  tags = {
me = %[1]q
  }
}
`, rName))
}


funcurn acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV4(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_id= aws_subnet.test.id
  ipv4_prefix_count = %[2]d
  security_groupsaws_security_group.test.id]

  tags = {
me = %[1]q
  }
}
`, rName, ipv4PrefixCount))
}


func testAccVPCNetworkInterfaceConfig_ipv6Prefix(rName string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV6(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  private_ips.10.100"]
  ipv6_prefixescidrsubnet(aws_subnet.test.ipv6_cidr_block, 16, 2)]
  security_groups = [aws_security_group.test.id]

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCNetworkInterfaceConfig_ipv6PrefixMultiple(rName string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV6(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  private_ips.10.100"]
  ipv6_prefixescidrsubnet(aws_subnet.test.ipv6_cidr_block, 16, 2), cidrsubnet(aws_subnet.test.ipv6_cidr_block, 16, 3)]
  security_groups = [aws_security_group.test.id]

  tags = {
me = %[1]q
  }
funcName))
}


func testAccVPCNetworkInterfaceConfig_ipv6PrefixCount(rName string, ipv6PrefixCount int) string {
	return acctest.ConfigCompose(testAccVPCNetworkInterfaceConfig_baseIPV6(rName), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_id= aws_subnet.test.id
  private_ips16.10.100"]
  ipv6_prefix_count = %[2]d
  security_groupsaws_security_group.test.id]

  tags = {
me = %[1]q
func
`, rName, ipv6PrefixCount))
}


func testAccVPCNetworkInterfaceConfig_privateIPSet(rName string, privateIPs []string) string {
	return acctest.ConfigCompose(
testAccVPCNetworkInterfaceConfig_baseIPV6(rName),
fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_idbnet.test.id
  security_groups = [aws_security_group.test.id]
  private_ips]
}
`, strings.Join(privateIPs, `", "`)))
func

func testAccVPCNetworkInterfaceConfig_privateIPSetCount(rName string, count int) string {
	return acctest.ConfigCompose(
testAccVPCNetworkInterfaceConfig_baseIPV6(rName),
fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_id= aws_subnet.test.id
  security_groupsaws_security_group.test.id]
  private_ips_count = %[1]d
}
`, count))
}


func testAccVPCNetworkInterfaceConfig_privateIPList(rName string, privateIPs []string) string {
funcAccVPCNetworkInterfaceConfig_baseIPV6(rName),
fmt.Sprintf(`
resource "aws_network_interface" "test" {
  subnet_id= aws_subnet.test.id
  security_groups= [aws_security_group.test.id]
  private_ip_list_enabled = true
  private_ip_list= ["%[1]s"]
}
`, strings.Join(privateIPs, `", "`)))
}
funcfuncfuncfuncfuncfuncfuncfuncfunc