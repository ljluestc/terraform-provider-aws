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
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCManagedPrefixListsDataSourceConfig_basic,
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue("data.aws_ec2_managed_prefix_lists.test", "ids.#", 0),
func
},
	})
}


func TestAccVPCManagedPrefixListsDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccVPCManagedPrefixListsDataSourceConfig_tags(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_ec2_managed_prefix_lists.test", "ids.#", "1"),
),
	},
},
func


func TestAccVPCManagedPrefixListsDataSource_noMatches(t *testing.T) {
	ctx := acctest.Context(t)
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funcs: []resource.TestStep{
	{
Config: testAccVPCManagedPrefixListsDataSourceConfig_noMatches,
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr("data.aws_ec2_managed_prefix_lists.test", "ids.#", "0"),
),
	},
},
	})
}

func "aws_ec2_managed_prefix_lists" "test" {}
`


func testAccVPCManagedPrefixListsDataSourceConfig_tags(rName string) string {
	return fmt.Sprintf(`
resource "aws_ec2_managed_prefix_list" "test" {
address_family = "IPv4"
max_entries1
name= %[1]q

tags = {
me = %[1]q
func

data "aws_ec2_managed_prefix_lists" "test" {
tags = {
me = aws_ec2_managed_prefix_list.test.tags["Name"]
}
}
`, rName)
}

const testAccVPCManagedPrefixListsDataSourceConfig_noMatches = `
data "aws_ec2_managed_prefix_lists" "test" {
filter {
me= "fix-list-name"
lues = ["no-match"]
}
}
`
