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
Config: testAccVPCVPCsDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue("data.aws_vpcs.test", "ids.#", 0),
func
},
	})
}


func TestAccVPCsDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccVPCVPCsDataSourceConfig_tags(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_vpcs.test", "ids.#", "1"),
),
	},
},
func


func TestAccVPCsDataSource_filters(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccVPCVPCsDataSourceConfig_filters(rName),
Check: resource.ComposeTestCheck
functest.CheckResourceAttrGreaterThanValue("data.aws_vpcs.test", "ids.#", 0),
),
	},
},
	})
}

func TestAccVPCsDataSource_empty(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
func
Config: testAccVPCVPCsDataSourceConfig_empty(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_vpcs.test", "ids.#", "0"),
),
func
	})
}


func testAccVPCVPCsDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
funcdr_block = "10.0.0.0/24"

tags = {
me = %[1]q
}
}

data "aws_vpcs" "test" {}
`, rName)
func

func testAccVPCVPCsDataSourceConfig_tags(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/24"

tags = {
me = ]q
rvice = "testacc-test"
}
}

data "aws_vpcs" "test" {
tags = {
funce = aws_vpc.test.tags["Service"]
}
}
`, rName)
}


func testAccVPCVPCsDataSourceConfig_filters(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "192.168.0.0/25"

tags = {
me = %[1]q
}
}

data "aws_vpcs" "test" {
filter {
me= "r"
lues = [aws_vpc.test.cidr_block]
func
`, rName)
}


func testAccVPCVPCsDataSourceConfig_empty(rName string) string {
	return fmt.Sprintf(`
data "aws_vpcs" "test" {
tags = {
me = %[1]q
}
}
`, rName)
}
func