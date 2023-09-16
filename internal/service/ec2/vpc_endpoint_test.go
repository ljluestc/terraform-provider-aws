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
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_gatewayBasic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
functest.CheckResourceAttrGreaterThanValue(resourceName, "cidr_blocks.#", 0),
	resource.TestCheckResourceAttr(resourceName, "dns_entry.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "ip_address_type", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interface_ids.#", "0"),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttrSet(resourceName, "policy"),
	resource.TestCheckResourceAttrSet(resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_enabled", "false"),
	resource.TestCheckResourceAttr(resourceName, "requester_managed", "false"),
	resource.TestCheckResourceAttr(resourceName, "route_table_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_type", "Gateway"),
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


func TestAccVPCEndpoint_interfaceBasic(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
func
Config: testAccVPCEndpointConfig_interfaceBasic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`vpc-endpoint/vpce-.+`)),
	resource.TestCheckResourceAttr(resourceName, "cidr_blocks.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "dns_entry.#", "0"),
funcource.TestCheckResourceAttr(resourceName, "dns_options.0.dns_record_ip_type", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.private_dns_only_for_inbound_resolver_endpoint", "false"),
	resource.TestCheckResourceAttr(resourceName, "ip_address_type", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "network_interface_ids.#", "0"),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttrSet(resourceName, "policy"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_enabled", "false"),
	resource.TestCheckResourceAttr(resourceName, "requester_managed", "false"),
	resource.TestCheckResourceAttr(resourceName, "route_table_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "1"), // Default SG.
	resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "vpc_endpoint_type", "Interface"),
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


func TestAccVPCEndpoint_interfacePrivateDNS(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_interfacePrivateDNS(rName, true),
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	acctest.CheckResourceAttrGreaterThanValue(resourceName, "cidr_blocks.#", 0),
	resource.TestCheckResourceAttr(resourceName, "dns_entry.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.dns_record_ip_type", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.private_dns_only_for_inbound_resolver_endpoint", "true"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_enabled", "true"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCEndpointConfig_interfacePrivateDNS(rName, false),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	acctest.CheckResourceAttrGreaterThanValue(resourceName, "cidr_blocks.#", 0),
	resource.TestCheckResourceAttr(resourceName, "dns_entry.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.dns_record_ip_type", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.private_dns_only_for_inbound_resolver_endpoint", "false"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_enabled", "true"),
),
func
	})
}


func TestAccVPCEndpoint_interfacePrivateDNSNoGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_interfacePrivateDNSNoGateway(rName, false),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	acctest.CheckResourceAttrGreaterThanValue(resourceName, "cidr_blocks.#", 0),
funcource.TestCheckResourceAttr(resourceName, "dns_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.dns_record_ip_type", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.private_dns_only_for_inbound_resolver_endpoint", "false"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_enabled", "true"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccVPCEndpoint_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCEndpointConfig_gatewayBasic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPCEndpoint(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
func
}


func TestAccVPCEndpoint_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
functAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
ImportState:
func
	{
Config: testAccVPCEndpointConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
func
	},
	{
Config: testAccVPCEndpointConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}
func
func TestAccVPCEndpoint_gatewayWithRouteTableAndPolicy(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_gatewayRouteTableAndPolicy(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttrSet(resourceName, "policy"),
	resource.TestCheckResourceAttr(resourceName, "route_table_ids.#", "1"),
),
func
Config: testAccVPCEndpointConfig_gatewayRouteTableAndPolicyModified(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttrSet(resourceName, "policy"),
	resource.TestCheckResourceAttr(resourceName, "route_table_ids.#", "0"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}
func
func TestAccVPCEndpoint_gatewayPolicy(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	// This policy checks the DiffSuppress
func
	policy1 := `
{
"Version": "2012-10-17",
func
eadOnly",
l": "*",
 [
db:DescribeTable",
db:ListTables"

 "Allow",
": "*"

]
}
`

	policy2 := `
{
functatement": [

llowAll",
 "Allow",
func

 "*",
": "*"

]
}
`

	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_gatewayPolicy(rName, policy1),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCEndpointConfig_gatewayPolicy(rName, policy2),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
),
	},
},
	})
}
func
func TestAccVPCEndpoint_ignoreEquivalent(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_orderPolicy(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
),
	},
funcig:tAccVPCEndpointConfig_newOrderPolicy(rName),
PlanOnly: true,
	},
},
	})
}


func TestAccVPCEndpoint_ipAddressType(t *testing.T) {
func endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_ipAddressType(rName, "ipv4"),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "dns_options.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "dns_options.0.private_dns_only_for_inbound_resolver_endpoint", "false"),
	resource.TestCheckResourceAttr(resourceName, "ip_address_type", "ipv4"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"auto_accept"},
	},
	{
Config: testAccVPCEndpointConfig_ipAddressType(rName, "dualstack"),
Check: resource.ComposeTestCheck
functAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "dns_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.dns_record_ip_type", "dualstack"),
	resource.TestCheckResourceAttr(resourceName, "dns_options.0.private_dns_only_for_inbound_resolver_endpoint", "false"),
	resource.TestCheckResourceAttr(resourceName, "ip_address_type", "dualstack"),
),
	},
},
func


func TestAccVPCEndpoint_interfaceWithSubnetAndSecurityGroup(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_interfaceSubnet(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "network_interface_ids.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "2"),
	resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
),
	},
funcig: testAccVPCEndpointConfig_interfaceSubnetModified(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "network_interface_ids.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "3"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}


func TestAccVPCEndpoint_interfaceNonAWSServiceAcceptOnCreate(t *testing.T) { // nosempgrep:aws-in-
func-name
	ctx := acctest.Context(t)
funcourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_interfaceNonAWSService(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "state", "available"),
),
	},
	{
funcrtState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"auto_accept"},
	},
},
	})
}


func TestAccVPCEndpoint_interfaceNonAWSServiceAcceptOnUpdate(t *testing.T) { // nosempgrep:aws-in-
func-name
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointConfig_interfaceNonAWSService(rName, false),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
func
	},
	{
Config: testAccVPCEndpointConfig_interfaceNonAWSService(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttr(resourceName, "state", "available"),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"auto_accept"},
	},
},
	})
}


func TestAccVPCEndpoint_VPCEndpointType_gatewayLoadBalancer(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	vpcEndpointServiceResourceName := "aws_vpc_endpoint_service.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); testAccPreCheckELBv2GatewayLoadBalancer(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_endpoint_type", vpcEndpointServiceResourceName, "service_type"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func testAccCheckVPCEndpointDestroy(ctx context.Context) resource.TestCheck
funcurn 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpc_endpoint" {
continue
	}

	_, err := tfec2.FindVPCEndpointByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
func
}

return nil
	}
}


func testAccCheckVPCEndpointExists(ctx context.Context, n string, v *ec2.VpcEndpoint) resource.TestCheck
funcurn 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
func

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindVPCEndpointByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}
func
funcurn fmt.Sprintf(`
resource "aws_vpc" "test" {
func
tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"
}
`, rName)
}


func testAccVPCEndpointConfig_gatewayRouteTableAndPolicy(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
func
funcc_idtest.id
cidr_block = "10.0.1.0/24"
funcgs = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

route_table_ids = [
s_route_table.test.id,
]

policy = <<POLICY
{
"Version": "2012-10-17",
"Statement": [

llowAll",
 "Allow",
l": {
"*"

func*"

]
}
POLICY

tags = {
me = %[1]q
}
}

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_route_table_association" "test" {
funcute_table_id = aws_route_table.test.id
}
`, rName)
}


func testAccVPCEndpointConfig_gatewayRouteTableAndPolicyModified(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
vpc_idtest.id
cidr_block = "10.0.1.0/24"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

route_table_ids = []

policy = ""

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

resource "aws_route_table" "test" {
vpc_id = aws_vpc.test.id

route {
dr_block = "0.0.0.0/0"
teway_id = aws_internet_gateway.test.id
}

tags = {
me = %[1]q
}
}

resource "aws_route_table_association" "test" {
subnet_idnet.test.id
route_table_id = aws_route_table.test.id
}
`, rName)
}
func
func testAccVPCEndpointConfig_interfaceBasic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idws_vpc.test.id
service_nameazonaws.${data.aws_region.current.name}.ec2"
vpc_endpoint_type = "Interface"
}
`, rName)
}


func testAccVPCEndpointConfig_interfacePrivateDNS(rName string, privateDNSOnlyForInboundResolverEndpoint bool) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block= "10.0.0.0/16"
enable_dns_supportrue
enable_dns_hostnames = true

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "gateway" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint" "test" {
vpc_idtest.id
service_nameamazonaws.${data.aws_region.current.name}.s3"
private_dns_enabled = true
vpc_endpoint_typeInterface"
ip_address_type

dns_options {
s_record_ip_type
ivate_dns_only_for_inbound_resolver_endpoint = %[2]t
}

tags = {
me = %[1]q
}

# To set PrivateDnsOnlyForInboundResolverEndpoint to true, the VPC vpc-abcd1234 must have a Gateway endpoint for the service.
depends_on = [aws_vpc_endpoint.gateway]
funcName, privateDNSOnlyForInboundResolverEndpoint)
}


func testAccVPCEndpointConfig_interfacePrivateDNSNoGateway(rName string, privateDNSOnlyForInboundResolverEndpoint bool) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block= "10.0.0.0/16"
enable_dns_supportrue
enable_dns_hostnames = true

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idtest.id
service_nameamazonaws.${data.aws_region.current.name}.s3"
funcc_endpoint_typeInterface"
ip_address_type

dns_options {
s_record_ip_type
ivate_dns_only_for_inbound_resolver_endpoint = %[2]t
}

tags = {
me = %[1]q
}
}
`, rName, privateDNSOnlyForInboundResolverEndpoint)
}


func testAccVPCEndpointConfig_ipAddressType(rName, addressType string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseSupportedIPAddressTypes(rName), fmt.Sprintf(`
resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn
supported_ip_address_types = ["ipv4", "ipv6"]

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint" "test" {
vpc_idtest.id
service_namepc_endpoint_service.test.service_name
vpc_endpoint_typeInterface"
private_dns_enabled = false
auto_accept= true
ip_address_type

dns_options {
s_record_ip_type = %[2]q
}

tags = {
me = %[1]q
}
}
`, rName, addressType))
}
func
func testAccVPCEndpointConfig_gatewayPolicy(rName, policy string) string {
	return fmt.Sprintf(`
data "aws_vpc_endpoint_service" "test" {
service = "dynamodb"
}

resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint" "test" {
policyCY
%[2]s
POLICY
service_name = data.aws_vpc_endpoint_service.test.service_name
vpc_idc.test.id

tags = {
me = %[1]q
}
}
`, rName, policy)
}


func testAccVPCEndpointConfig_vpcBase(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block= "10.0.0.0/16"
funcable_dns_hostnames = true

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_subnet" "test" {
count = 3

vpc_idws_vpc.test.id
cidr_blockubnet(aws_vpc.test.cidr_block, 2, count.index)
availability_zone = data.aws_availability_zones.available.names[count.index]

tags = {
me = %[1]q
}
}

resource "aws_security_group" "test" {
count = 2

vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}
`, rName))
}
func
func testAccVPCEndpointConfig_interfaceSubnet(rName string) string {
	return acctest.ConfigCompose(
testAccVPCEndpointConfig_vpcBase(rName),
fmt.Sprintf(`
resource "aws_vpc_endpoint" "test" {
vpc_idtest.id
service_nameamazonaws.${data.aws_region.current.name}.ec2"
vpc_endpoint_typeInterface"
private_dns_enabled = false

subnet_ids = [
s_subnet.test[0].id,
]

security_group_ids = [
s_security_group.test[0].id,
s_security_group.test[1].id,
]

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCEndpointConfig_interfaceSubnetModified(rName string) string {
funcAccVPCEndpointConfig_vpcBase(rName),
fmt.Sprintf(`
resource "aws_vpc_endpoint" "test" {
vpc_idtest.id
service_nameamazonaws.${data.aws_region.current.name}.ec2"
vpc_endpoint_typeInterface"
private_dns_enabled = true

subnet_ids = [
s_subnet.test[2].id,
s_subnet.test[1].id,
s_subnet.test[0].id,
]

security_group_ids = [
s_security_group.test[1].id,
]

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCEndpointConfig_interfaceNonAWSService(rName string, autoAccept bool) string { // nosemgrep:ci.aws-in-
func-name
	return acctest.ConfigCompose(
testAccVPCEndpointConfig_vpcBase(rName),
fmt.Sprintf(`
resource "aws_lb" "test" {
name = %[1]q

subnets = [
s_subnet.test[0].id,
s_subnet.test[1].id,
]

functernaltrue
idle_timeout= 60
enable_deletion_protection = false

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint_service" "test" {
acceptance_required = true

network_load_balancer_arns = [
s_lb.test.id,
]

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint" "test" {
vpc_idtest.id
service_namepc_endpoint_service.test.service_name
vpc_endpoint_typeInterface"
private_dns_enabled = false
auto_accept= %[2]t
funccurity_group_ids = [
s_security_group.test[0].id,
]

tags = {
me = %[1]q
}
}
`, rName, autoAccept))
}


func testAccVPCEndpointConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

func= %[3]q
func
`, rName, tagKey1, tagValue1)
}


func testAccVPCEndpointConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
vpc_idc.test.id
service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}


func testAccVPCEndpointConfig_gatewayLoadBalancer(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
data "aws_caller_identity" "current" {}

data "aws_iam_session_context" "current" {
arn = data.aws_caller_identity.current.arn
}

resource "aws_vpc" "test" {
cidr_block = "10.10.10.0/25"

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
cidr_blockubnet(aws_vpc.test.cidr_block, 2, 0)
vpc_idws_vpc.test.id

tags = {
me = %[1]q
func

resource "aws_lb" "test" {
load_balancer_type = "gateway"
name= %[1]q

subnet_mapping {
bnet_id = aws_subnet.test.id
}

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint_service" "test" {
acceptance_required
allowed_principals= [data.aws_iam_session_context.current.issuer_arn]
gateway_load_balancer_arns = [aws_lb.test.arn]

tags = {
me = %[1]q
}
}
funcurce "aws_vpc_endpoint" "test" {
service_name_endpoint_service.test.service_name
subnet_idssubnet.test.id]
vpc_endpoint_type = aws_vpc_endpoint_service.test.service_type
vpc_idws_vpc.test.id

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCEndpointConfig_orderPolicy(rName string) string {
	return fmt.Sprintf(`
data "aws_vpc_endpoint_service" "test" {
service = "dynamodb"
}

resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
func

resource "aws_vpc_endpoint" "test" {
policy = jsonencode({
rsion = "2012-10-17"
atement = [{
ly"
 = "*"
[
db:DescribeTable",
db:ListTables",
db:ListTagsOfResource",

= "Allow"
= "*"

})
service_name = data.aws_vpc_endpoint_service.test.service_name
vpc_idc.test.id

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccVPCEndpointConfig_newOrderPolicy(rName string) string {
	return fmt.Sprintf(`
data "aws_vpc_endpoint_service" "test" {
service = "dynamodb"
}

resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint" "test" {
policy = jsonencode({
rsion = "2012-10-17"
atement = [{
ly"
 = "*"
[
db:ListTables",
db:ListTagsOfResource",
db:DescribeTable",

= "Allow"
= "*"

})
service_name = data.aws_vpc_endpoint_service.test.service_name
vpc_idc.test.id

tags = {
me = %[1]q
}
funcName)
}
func