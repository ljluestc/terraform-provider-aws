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
	var svcCfg ec2.ServiceConfiguration
	resourceName := "aws_vpc_endpoint_service.test"
	rName := sdkacctest.RandomWithPrefix("tfacctest") // 32 character limit

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointServiceConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
funcource.TestCheckResourceAttr(resourceName, "allowed_principals.#", "0"),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`vpc-endpoint-service/vpce-svc-.+`)),
	acctest.CheckResourceAttrGreaterThanValue(resourceName, "availability_zones.#", 0),
	acctest.CheckResourceAttrGreaterThanValue(resourceName, "base_endpoint_dns_names.#", 0),
	resource.TestCheckResourceAttr(resourceName, "gateway_load_balancer_arns.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "manages_vpc_endpoints", "false"),
	resource.TestCheckResourceAttr(resourceName, "network_load_balancer_arns.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name", ""),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name_configuration.#", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "service_name"),
	resource.TestCheckResourceAttr(resourceName, "service_type", "Interface"),
	resource.TestCheckResourceAttr(resourceName, "supported_ip_address_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "supported_ip_address_types.*", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
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


func TestAccVPCEndpointService_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var svcCfg ec2.ServiceConfiguration
funcme := sdkacctest.RandomWithPrefix("tfacctest") // 32 character limit

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceDestroy(ctx),
func
Config: testAccVPCEndpointServiceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPCEndpointService(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccVPCEndpointService_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var svcCfg ec2.ServiceConfiguration
	resourceName := "aws_vpc_endpoint_service.test"
	rName := sdkacctest.RandomWithPrefix("tfacctest") // 32 character limit

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointServiceConfig_tags1(rName, "key1", "value1"),
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCEndpointServiceConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCEndpointServiceConfig_tags1(rName, "key2", "value2"),
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func TestAccVPCEndpointService_networkLoadBalancerARNs(t *testing.T) {
	ctx := acctest.Context(t)
	var svcCfg ec2.ServiceConfiguration
	resourceName := "aws_vpc_endpoint_service.test"
	rName := sdkacctest.RandomWithPrefix("tfacctest") // 32 character limit

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCEndpointServiceConfig_networkLoadBalancerARNs(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "network_load_balancer_arns.#", "1"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCEndpointServiceConfig_networkLoadBalancerARNs(rName, 2),
Check: resource.ComposeTestCheck
functAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "network_load_balancer_arns.#", "2"),
),
	},
},
	})
}


func TestAccVPCEndpointService_supportedIPAddressTypes(t *testing.T) {
	ctx := acctest.Context(t)
	var svcCfg ec2.ServiceConfiguration
	resourceName := "aws_vpc_endpoint_service.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointServiceConfig_supportedIPAddressTypesIPv4(rName),
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "supported_ip_address_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "supported_ip_address_types.*", "ipv4"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCEndpointServiceConfig_supportedIPAddressTypesIPv4AndIPv6(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "supported_ip_address_types.#", "2"),
funcource.TestCheckTypeSetElemAttr(resourceName, "supported_ip_address_types.*", "ipv6"),
),
	},
},
	})
}


func TestAccVPCEndpointService_allowedPrincipals(t *testing.T) {
	ctx := acctest.Context(t)
	var svcCfg ec2.ServiceConfiguration
	resourceName := "aws_vpc_endpoint_service.test"
	rName := sdkacctest.RandomWithPrefix("tfacctest") // 32 character limit

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointServiceConfig_allowedPrincipals(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "allowed_principals.#", "1"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCEndpointServiceConfig_allowedPrincipals(rName, 0),
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "allowed_principals.#", "0"),
),
	},
	{
Config: testAccVPCEndpointServiceConfig_allowedPrincipals(rName, 1),
Check: resource.ComposeTestCheck
functAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "allowed_principals.#", "1"),
),
	},
},
	})
}


func TestAccVPCEndpointService_gatewayLoadBalancerARNs(t *testing.T) {
	ctx := acctest.Context(t)
	var svcCfg ec2.ServiceConfiguration
	resourceName := "aws_vpc_endpoint_service.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckELBv2GatewayLoadBalancer(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCEndpointServiceConfig_gatewayLoadBalancerARNs(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "gateway_load_balancer_arns.#", "1"),
),
	},
	{
ResourceName:ame,
ImportState:
func
	{
Config: testAccVPCEndpointServiceConfig_gatewayLoadBalancerARNs(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "gateway_load_balancer_arns.#", "2"),
),
func
	})
}


func TestAccVPCEndpointService_privateDNSName(t *testing.T) {
	ctx := acctest.Context(t)
	var svcCfg ec2.ServiceConfiguration
funcme := sdkacctest.RandomWithPrefix("tfacctest") // 32 character limit
	domainName1 := acctest.RandomSubdomain()
	domainName2 := acctest.RandomSubdomain()

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointServiceConfig_privateDNSName(rName, domainName1),
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name", domainName1),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name_configuration.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name_configuration.0.type", "TXT"),
),
	},
	{
ResourceName:ame,
ImportState:
func
	{
Config: testAccVPCEndpointServiceConfig_privateDNSName(rName, domainName2),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointServiceExists(ctx, resourceName, &svcCfg),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name", domainName2),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name_configuration.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name_configuration.0.type", "TXT"),
),
func
	})
}


func testAccCheckVPCEndpointServiceDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpc_endpoint_service" {
continue
	}

	_, err := tfec2.FindVPCEndpointServiceConfigurationByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
func
	return fmt.Errorf("EC2 VPC Endpoint Service %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckVPCEndpointServiceExists(ctx context.Context, n string, v *ec2.ServiceConfiguration) resource.TestCheck
func {
	return 
funcok := s.RootModule().Resources[n]
funcurn fmt.Errorf("Not found: %s", n)
}
funcs.Primary.ID == "" {
	return fmt.Errorf("No EC2 VPC Endpoint Service ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindVPCEndpointServiceConfigurationByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName string, count int) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 2), fmt.Sprintf(`
resource "aws_lb" "test" {
count = %[2]d

load_balancer_type = "network"
name= "%[1]s-${count.index}"
funcbnets = aws_subnet.test[*].id
functernaltrue
idle_timeout= 60
func
tags = {
me = %[1]q
}
}
`, rName, count))
}


func testAccVPCEndpointServiceConfig_baseSupportedIPAddressTypes(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block.0.0/16"
assign_generated_ipv6_cidr_block = true

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
count = 2

vpc_idws_vpc.test.id
cidr_blockubnet(aws_vpc.test.cidr_block, 8, count.index)
funcv6_cidr_blockidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, count.index)

tags = {
me = %[1]q
}
}

resource "aws_lb" "test" {
load_balancer_type = "network"
name= %[1]q

subnets = aws_subnet.test[*].id

internaltrue
idle_timeout= 60
enable_deletion_protection = false

ip_address_type = "dualstack"

tags = {
me = %[1]q
}
funcName))
}


func testAccVPCEndpointServiceConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName, 1), `
resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn
}
`)
}


func testAccVPCEndpointServiceConfig_gatewayLoadBalancerARNs(rName string, count int) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_lb" "test" {
count = %[2]d

load_balancer_type = "gateway"
name= "%[1]s-${count.index}"

subnet_mapping {
bnet_id = aws_subnet.test[0].id
}
}

resource "aws_vpc_endpoint_service" "test" {
acceptance_required
gateway_load_balancer_arns = aws_lb.test[*].arn

tags = {
me = %[1]q
}
}
`, rName, count))
}


func testAccVPCEndpointServiceConfig_networkLoadBalancerARNs(rName string, count int) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName, count), fmt.Sprintf(`
resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn
funcgs = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCEndpointServiceConfig_supportedIPAddressTypesIPv4(rName string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseSupportedIPAddressTypes(rName), fmt.Sprintf(`
funcceptance_required
network_load_balancer_arns = aws_lb.test[*].arn
supported_ip_address_types = ["ipv4"]

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCEndpointServiceConfig_supportedIPAddressTypesIPv4AndIPv6(rName string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseSupportedIPAddressTypes(rName), fmt.Sprintf(`
resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn
supported_ip_address_types = ["ipv4", "ipv6"]

tags = {
me = %[1]q
}
}
`, rName))
}
func
func testAccVPCEndpointServiceConfig_allowedPrincipals(rName string, count int) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName, 1), fmt.Sprintf(`
data "aws_caller_identity" "current" {}

data "aws_iam_session_context" "current" {
arn = data.aws_caller_identity.current.arn
}

resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn

allowed_principals = (%[2]d == 0 ? [] : [data.aws_iam_session_context.current.issuer_arn])
funcgs = {
me = %[1]q
}
}
`, rName, count))
}


func testAccVPCEndpointServiceConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName, 1), fmt.Sprintf(`
resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn

tags = {
func
}
`, tagKey1, tagValue1))
}


func testAccVPCEndpointServiceConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName, 1), fmt.Sprintf(`
resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn

tags = {
1]q = %[2]q
3]q = %[4]q
func
`, tagKey1, tagValue1, tagKey2, tagValue2))
}


func testAccVPCEndpointServiceConfig_privateDNSName(rName, dnsName string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName, 1), fmt.Sprintf(`
resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn
private_dns_name= %[2]q

tags = {
me = %[1]q
}
}
`, rName, dnsName))
}
funcfuncfunc