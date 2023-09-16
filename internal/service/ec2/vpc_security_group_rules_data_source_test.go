// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupRulesDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_vpc_security_group_rules.test", "ids.#", "1"),
func
},
	})
}


func TestAccVPCSecurityGroupRulesDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccVPCSecurityGroupRulesDataSourceConfig_tags(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_vpc_security_group_rules.test", "ids.#", "2"),
),
	},
},
func


func testAccVPCSecurityGroupRulesDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
security_group_id = aws_security_group.test.id

cidr_ipv410.0.0.0/8"
func_protocol = "tcp"
to_port
}

data "aws_vpc_security_group_rules" "test" {
filter {
me= "urity-group-rule-id"
lues = [aws_vpc_security_group_ingress_rule.test.id]
}
}
`)
}


func testAccVPCSecurityGroupRulesDataSourceConfig_tags(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), fmt.Sprintf(`
resource "aws_vpc_security_group_ingress_rule" "test" {
security_group_id = aws_security_group.test.id

cidr_ipv410.0.0.0/8"
from_port0
func_port

tags = {
me = %[1]q
}
}

resource "aws_vpc_security_group_egress_rule" "test" {
security_group_id = aws_security_group.test.id

cidr_ipv410.0.0.0/8"
from_port0
ip_protocol = "tcp"
to_port

tags = {
me = %[1]q
}
}

data "aws_vpc_security_group_rules" "test" {
tags = {
me = %[1]q
}

depends_on = [aws_vpc_security_group_ingress_rule.test, aws_vpc_security_group_egress_rule.test]
}
`, rName))
}
