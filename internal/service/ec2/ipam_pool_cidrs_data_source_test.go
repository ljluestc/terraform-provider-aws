// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	dataSourceName := "data.aws_vpc_ipam_pool_cidrs.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolCIDRsDataSourceConfig_basicOneCIDRs,
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ipam_pool_cidrs.#", "1"),
func
	{
Config: testAccIPAMPoolCIDRsDataSourceConfig_basicTwoCIDRs,
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ipam_pool_cidrs.#", "2"),
),
func
Config: testAccIPAMPoolCIDRsDataSourceConfig_basicTwoCIDRsFiltered,
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ipam_pool_cidrs.#", "1"),
),
	},
func
}

var testAccIPAMPoolCIDRsDataSourceConfig_basicOneCIDRs = acctest.ConfigCompose(testAccIPAMPoolConfig_basic, `
resource "aws_vpc_ipam_pool_cidr" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr.2.0.0/16"
}

data "aws_vpc_ipam_pool_cidrs" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id

depends_on = [
s_vpc_ipam_pool_cidr.test
]
}
`)

var testAccIPAMPoolCIDRsDataSourceConfig_basicTwoCIDRs = acctest.ConfigCompose(testAccIPAMPoolConfig_basic, `
resource "aws_vpc_ipam_pool_cidr" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr.2.0.0/16"
}

resource "aws_vpc_ipam_pool_cidr" "testtwo" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr2.0.0/16"
}

data "aws_vpc_ipam_pool_cidrs" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id

depends_on = [
s_vpc_ipam_pool_cidr.test,
s_vpc_ipam_pool_cidr.testtwo,
]
}
`)

var testAccIPAMPoolCIDRsDataSourceConfig_basicTwoCIDRsFiltered = acctest.ConfigCompose(testAccIPAMPoolConfig_basic, `
resource "aws_vpc_ipam_pool_cidr" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr.2.0.0/16"
}

resource "aws_vpc_ipam_pool_cidr" "testtwo" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr2.0.0/16"
}

data "aws_vpc_ipam_pool_cidrs" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id

filter {
me= "r"
lues = ["10.*"]
}

depends_on = [
s_vpc_ipam_pool_cidr.test,
s_vpc_ipam_pool_cidr.testtwo,
]
}
`)
