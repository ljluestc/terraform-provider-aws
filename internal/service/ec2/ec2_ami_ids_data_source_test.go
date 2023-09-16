// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	datasourceName := "data.aws_ami_ids.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccAMIIDsDataSourceConfig_basic,
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue(datasourceName, "ids.#", 0),
func
},
	})
}


func TestAccEC2AMIIDsDataSource_sorted(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_ami_ids.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccAMIIDsDataSourceConfig_sorted(false),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(datasourceName, "ids.#", "2"),
	resource.TestCheckResourceAttrPair(datasourceName, "ids.0", "data.aws_ami.test2", "id"),
	resource.TestCheckResourceAttrPair(datasourceName, "ids.1", "data.aws_ami.test1", "id"),
),
func
Config: testAccAMIIDsDataSourceConfig_sorted(true),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(datasourceName, "ids.#", "2"),
	resource.TestCheckResourceAttrPair(datasourceName, "ids.0", "data.aws_ami.test1", "id"),
	resource.TestCheckResourceAttrPair(datasourceName, "ids.1", "data.aws_ami.test2", "id"),
),
	},
func
}


func TestAccEC2AMIIDsDataSource_includeDeprecated(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_ami_ids.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccAMIIDsDataSourceConfig_includeDeprecated(true),
Check: resource.ComposeTestCheck
func(
func
	},
},
	})
}

const testAccAMIIDsDataSourceConfig_basic = `
funcners = ["099720109477"]

filter {
me= "e"
lues = ["ubuntu/images/ubuntu-*-*-amd64-server-*"]
}
}
`


func testAccAMIIDsDataSourceConfig_sorted(sortAscending bool) string {
	return fmt.Sprintf(`
data "aws_ami" "test1" {
owners = ["amazon"]

filter {
me= "e"
lues = ["amzn-ami-hvm-2018.03.0.20221018.0-x86_64-gp2"]
}
}
func "aws_ami" "test2" {
owners = ["amazon"]

filter {
me= "e"
lues = ["amzn-ami-hvm-2018.03.0.20221209.1-x86_64-gp2"]
}
}

data "aws_ami_ids" "test" {
owners = ["amazon"]

filter {
me= "e"
lues = [data.aws_ami.test1.name, data.aws_ami.test2.name]
}

sort_ascending = %[1]t
}
`, sortAscending)
}


func testAccAMIIDsDataSourceConfig_includeDeprecated(includeDeprecated bool) string {
	return fmt.Sprintf(`
data "aws_ami_ids" "test" {
owners["099720109477"]
include_deprecated = %[1]t

filter {
me= "e"
lues = ["ubuntu/images/ubuntu-*-*-amd64-server-*"]
}
}
func
