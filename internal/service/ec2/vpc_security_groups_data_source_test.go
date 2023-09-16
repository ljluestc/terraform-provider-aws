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
	dataSourceName := "data.aws_security_groups.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupsDataSourceConfig_tag(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "arns.#", "3"),
funcource.TestCheckResourceAttr(dataSourceName, "vpc_ids.#", "3"),
),
	},
},
	})
}


func TestAccVPCSecurityGroupsDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccVPCSecurityGroupsDataSourceConfig_filter(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "arns.#", "1"),
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "1"),
	resource.TestCheckResourceAttr(dataSourceName, "vpc_ids.#", "1"),
),
func
	})
}


func TestAccVPCSecurityGroupsDataSource_empty(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_security_groups.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupsDataSourceConfig_empty(rName),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr(dataSourceName, "arns.#", "0"),
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "0"),
	resource.TestCheckResourceAttr(dataSourceName, "vpc_ids.#", "0"),
),
	},
},
	})
func

func testAccVPCSecurityGroupsDataSourceConfig_tag(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "172.16.0.0/16"

tags = {
me = %[1]q
}
}
funcurce "aws_security_group" "test" {
count= 3
vpc_id = aws_vpc.test.id
name%[1]s-${count.index}"

tags = {
me = %[1]q
}
}

data "aws_security_groups" "test" {
tags = {
me = %[1]q
}

depends_on = [aws_security_group.test[0], aws_security_group.test[1], aws_security_group.test[2]]
}
`, rName)
}


func testAccVPCSecurityGroupsDataSourceConfig_filter(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "172.16.0.0/16"

tags = {
me = %[1]q
}
}

funcc_id = aws_vpc.test.id
name[1]q

tags = {
me = %[1]q
}
}

data "aws_security_groups" "test" {
filter {
me= "-id"
lues = [aws_vpc.test.id]
}

filter {
me= "up-name"
lues = [aws_security_group.test.name]
}
}
`, rName)
}


func testAccVPCSecurityGroupsDataSourceConfig_empty(rName string) string {
	return fmt.Sprintf(`
data "aws_security_groups" "test" {
tags = {
me = %[1]q
}
}
`, rName)
}
func