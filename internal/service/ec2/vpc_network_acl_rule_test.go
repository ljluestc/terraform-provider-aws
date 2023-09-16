// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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
	resource1Name := "aws_network_acl_rule.test1"
	resource2Name := "aws_network_acl_rule.test2"
	resource3Name := "aws_network_acl_rule.test3"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkACLRuleConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNetworkACLRuleExists(ctx, resource1Name),
functAccCheckNetworkACLRuleExists(ctx, resource3Name),

	resource.TestCheckResourceAttr(resource1Name, "cidr_block", "0.0.0.0/0"),
	resource.TestCheckResourceAttr(resource1Name, "egress", "false"),
	resource.TestCheckResourceAttr(resource1Name, "from_port", "22"),
	resource.TestCheckResourceAttr(resource1Name, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resource1Name, "protocol", "6"),
	resource.TestCheckResourceAttr(resource1Name, "rule_action", "allow"),
	resource.TestCheckResourceAttr(resource1Name, "rule_number", "200"),
	resource.TestCheckResourceAttr(resource1Name, "to_port", "22"),

	resource.TestCheckResourceAttr(resource2Name, "cidr_block", "0.0.0.0/0"),
	resource.TestCheckResourceAttr(resource2Name, "egress", "false"),
	resource.TestCheckResourceAttr(resource2Name, "icmp_code", "-1"),
	resource.TestCheckResourceAttr(resource2Name, "icmp_type", "0"),
	resource.TestCheckResourceAttr(resource2Name, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resource2Name, "protocol", "1"),
	resource.TestCheckResourceAttr(resource2Name, "rule_action", "allow"),
	resource.TestCheckResourceAttr(resource2Name, "rule_number", "300"),

	resource.TestCheckResourceAttr(resource3Name, "cidr_block", "0.0.0.0/0"),
	resource.TestCheckResourceAttr(resource3Name, "egress", "false"),
	resource.TestCheckResourceAttr(resource3Name, "icmp_code", "-1"),
	resource.TestCheckResourceAttr(resource3Name, "icmp_type", "-1"),
	resource.TestCheckResourceAttr(resource3Name, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resource3Name, "protocol", "1"),
	resource.TestCheckResourceAttr(resource3Name, "rule_action", "allow"),
	resource.TestCheckResourceAttr(resource3Name, "rule_number", "400"),
),
	},
	{
ResourceName:Name,
ImportState:
ImportStateId
func: testAccNetworkACLRuleImportStateId
func(resource1Name, "tcp"),
ImportStateVerify: true,
func
funcrtState:
ImportStateId
func: testAccNetworkACLRuleImportStateId
func(resource2Name, "icmp"),
ImportStateVerify: true,
	},
	{
funcrtState:
func: testAccNetworkACLRuleImportStateId
func(resource3Name, "icmp"),
ImportStateVerify: true,
	},
},
	})
}
func
func := acctest.Context(t)
	resourceName := "aws_network_acl_rule.test1"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkACLRuleConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkACLRuleExists(ctx, resourceName),
func
ExpectNonEmptyPlan: true,
	},
},
	})
}


func := acctest.Context(t)
	resourceName := "aws_network_acl_rule.test1"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCNetworkACLRuleConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkACLRuleExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceNetworkACL(), "aws_network_acl.test"),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccVPCNetworkACLRule_Disappears_ingressEgressSameNumber(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_network_acl_rule.test1"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkACLRuleConfig_ingressEgressSameNumberMissing(rName),
Check: resource.ComposeTestCheck
functAccCheckNetworkACLRuleExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceNetworkACLRule(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
func

func TestAccVPCNetworkACLRule_ipv6(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_network_acl_rule.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkACLRuleConfig_ipv6(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNetworkACLRuleExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "cidr_block", ""),
funcource.TestCheckResourceAttr(resourceName, "from_port", "22"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", "::/0"),
	resource.TestCheckResourceAttr(resourceName, "protocol", "6"),
	resource.TestCheckResourceAttr(resourceName, "rule_action", "allow"),
	resource.TestCheckResourceAttr(resourceName, "rule_number", "150"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "22"),
),
func
ResourceName:ame,
ImportState:
ImportStateId
func: testAccNetworkACLRuleImportStateId
func(resourceName, "tcp"),
ImportStateVerify: true,
	},
func
}


func TestAccVPCNetworkACLRule_ipv6ICMP(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_network_acl_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCNetworkACLRuleConfig_ipv6ICMP(rName),
func(
	testAccCheckNetworkACLRuleExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "egress", "false"),
	resource.TestCheckResourceAttr(resourceName, "icmp_code", "-1"),
	resource.TestCheckResourceAttr(resourceName, "icmp_type", "-1"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", "::/0"),
	resource.TestCheckResourceAttr(resourceName, "protocol", "58"),
funcource.TestCheckResourceAttr(resourceName, "rule_number", "150"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func(resourceName, "58"),
ImportStateVerify: true,
	},
},
	})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/6710
func TestAccVPCNetworkACLRule_ipv6VPCAssignGeneratedIPv6CIDRBlockUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Vpc
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vpcResourceName := "aws_vpc.test"
	resourceName := "aws_network_acl_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkACLRuleConfig_ipv6NotAssignGeneratedIPv6CIDRBlockUpdate(rName),
func(
funcource.TestCheckResourceAttr(vpcResourceName, "assign_generated_ipv6_cidr_block", "false"),
	resource.TestCheckResourceAttr(vpcResourceName, "ipv6_cidr_block", ""),
),
	},
	{
Config: testAccVPCNetworkACLRuleConfig_ipv6AssignGeneratedIPv6CIDRBlockUpdate(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckVPCExists(ctx, vpcResourceName, &v),
funcource.TestCheckResourceAttr(vpcResourceName, "assign_generated_ipv6_cidr_block", "true"),
	resource.TestMatchResourceAttr(vpcResourceName, "ipv6_cidr_block", regexache.MustCompile(`/56$`)),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccNetworkACLRuleImportStateId
funcrtStateVerify: true,
	},
},
	})
}


func TestAccVPCNetworkACLRule_allProtocol(t *testing.T) {
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_network_acl_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
func
Config: testAccVPCNetworkACLRuleConfig_allProtocol(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNetworkACLRuleExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "cidr_block", "0.0.0.0/0"),
	resource.TestCheckResourceAttr(resourceName, "egress", "false"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "22"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "protocol", "-1"),
	resource.TestCheckResourceAttr(resourceName, "rule_action", "allow"),
funcource.TestCheckResourceAttr(resourceName, "to_port", "22"),
func
	{
Config:tAccVPCNetworkACLRuleConfig_allProtocolNoRealUpdate(rName),
PlanOnly: true,
	},
},
	})
}
func
func TestAccVPCNetworkACLRule_tcpProtocol(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_network_acl_rule.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkACLRuleConfig_tcpProtocol(rName),
Check: resource.ComposeAggregateTestCheck
functAccCheckNetworkACLRuleExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "cidr_block", "0.0.0.0/0"),
	resource.TestCheckResourceAttr(resourceName, "egress", "true"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "22"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "protocol", "6"),
	resource.TestCheckResourceAttr(resourceName, "rule_action", "deny"),
	resource.TestCheckResourceAttr(resourceName, "rule_number", "150"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "22"),
),
	},
	{
Config:tAccVPCNetworkACLRuleConfig_tcpProtocolNoRealUpdate(rName),
PlanOnly: true,
	},
},
	})
}


func testAccCheckNetworkACLRuleDestroy(ctx context.Context) resource.TestCheck
funcurn 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_network_acl_rule" {
continue
func
	egress, err := strconv.ParseBool(rs.Primary.Attributes["egress"])

	if err != nil {
return err
	}

	naclID := rs.Primary.Attributes["network_acl_id"]
funceNumber, err := strconv.Atoi(rs.Primary.Attributes["rule_number"])

	if err != nil {
return err
	}

	_, err = tfec2.FindNetworkACLEntryByThreePartKey(ctx, conn, naclID, egress, ruleNumber)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Network ACL Rule %s still exists", rs.Primary.ID)
}

return nil
	}
func
func testAccCheckNetworkACLRuleExists(ctx context.Context, n string) resource.TestCheck
func {
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Network ACL Rule ID is set: %s", n)
}

egress, err := strconv.ParseBool(rs.Primary.Attributes["egress"])

if err != nil {
	return err
}

naclID := rs.Primary.Attributes["network_acl_id"]

ruleNumber, err := strconv.Atoi(rs.Primary.Attributes["rule_number"])

if err != nil {
	return err
}

_, err = tfec2.FindNetworkACLEntryByThreePartKey(ctx, conn, naclID, egress, ruleNumber)

return err
	}
}


func testAccVPCNetworkACLRuleConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.3.0.0/16"

tags = {
me = %[1]q
}
func
funcc_id = aws_vpc.test.id

func %[1]q
}
}

resource "aws_network_acl_rule" "test1" {
network_acl_id = aws_network_acl.test.id
rule_number200
egress= false
protocol
rule_action"allow"
cidr_block/0"
from_port
to_port
}

resource "aws_network_acl_rule" "test2" {
network_acl_id = aws_network_acl.test.id
rule_number300
protocol
rule_action"allow"
cidr_block/0"
icmp_type
icmp_code
}

resource "aws_network_acl_rule" "test3" {
network_acl_id = aws_network_acl.test.id
rule_number400
protocol
rule_action"allow"
cidr_block/0"
icmp_type
func
`, rName)
}


func testAccVPCNetworkACLRuleConfig_allProtocolNoRealUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.3.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_network_acl_rule" "test" {
network_acl_id = aws_network_acl.test.id
rule_number150
egress= false
protocol
rule_action"allow"
cidr_block/0"
from_port
to_port
}
`, rName)
}


func testAccVPCNetworkACLRuleConfig_tcpProtocolNoRealUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.3.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
func

resource "aws_network_acl_rule" "test" {
network_acl_id = aws_network_acl.test.id
rule_number150
egress= true
protocol
rule_action"deny"
cidr_block/0"
from_port
to_port
}
`, rName)
}


func testAccVPCNetworkACLRuleConfig_allProtocol(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.3.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
func
resource "aws_network_acl_rule" "test" {
network_acl_id = aws_network_acl.test.id
rule_number150
egress= false
protocol
rule_action"allow"
cidr_block/0"
from_port
to_port
}
`, rName)
}


func testAccVPCNetworkACLRuleConfig_tcpProtocol(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.3.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}
funcurce "aws_network_acl_rule" "test" {
network_acl_id = aws_network_acl.test.id
rule_number150
egress= true
protocol
rule_action"deny"
cidr_block/0"
from_port
to_port
}
`, rName)
}


func testAccVPCNetworkACLRuleConfig_ipv6(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.3.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

functwork_acl_id= aws_network_acl.test.id
rule_number
egress = false
protocol
rule_action
ipv6_cidr_block = "::/0"
from_port
to_port= 22
}
`, rName)
}


func testAccVPCNetworkACLRuleConfig_ingressEgressSameNumberMissing(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.3.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_network_acl_rule" "test1" {
funcle_number100
egress= false
protocol
rule_action"allow"
cidr_block/0"
from_port
to_port
}

resource "aws_network_acl_rule" "test2" {
network_acl_id = aws_network_acl.test.id
rule_number100
egress= true
protocol
rule_action"allow"
cidr_block/0"
from_port
to_port
}
`, rName)
}


func testAccVPCNetworkACLRuleConfig_ipv6ICMP(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.3.0.0/16"

tags = {
me = %[1]q
}
}
funcurce "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_network_acl_rule" "test" {
icmp_code
icmp_type
ipv6_cidr_block = "::/0"
network_acl_id= aws_network_acl.test.id
protocol
rule_action
rule_number
}
`, rName, rName)
}


func testAccVPCNetworkACLRuleConfig_ipv6AssignGeneratedIPv6CIDRBlockUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
assign_generated_ipv6_cidr_block = true
cidr_block.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_network_acl_rule" "test" {
from_port
ipv6_cidr_block = aws_vpc.test.ipv6_cidr_block
funcotocol
rule_action
rule_number
to_port= 22
}
`, rName)
}


func testAccVPCNetworkACLRuleConfig_ipv6NotAssignGeneratedIPv6CIDRBlockUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
assign_generated_ipv6_cidr_block = false
cidr_block.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_network_acl" "test" {
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}`, rName)
}


func testAccNetworkACLRuleImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return "", fmt.Errorf("Not found: %s", resourceName)
}

naclID := rs.Primary.Attributes["network_acl_id"]
ruleNumber := rs.Primary.Attributes["rule_number"]
protocol := rs.Primary.Attributes["protocol"]
// Ensure the resource's ID will be determined from the original protocol value set in the resource's config
if protocol != resourceProtocol {
	protocol = resourceProtocol
}
egress := rs.Primary.Attributes["egress"]

return strings.Join([]string{naclID, ruleNumber, protocol, egress}, tfec2.NetworkACLRuleImportIDSeparator), nil
	}
}
funcfuncfuncfuncfunc