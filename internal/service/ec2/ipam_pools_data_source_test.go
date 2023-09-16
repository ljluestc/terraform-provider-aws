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
	dataSourceName := "data.aws_vpc_ipam_pools.test"
	dataSourceNameTwo := "data.aws_vpc_ipam_pools.testtwo"
	resourceName := "aws_vpc_ipam_pool.testthree"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccIPAMPoolsDataSourceConfig_basic,
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "ipam_pools.#", 0),
func
	{
Config: testAccIPAMPoolsDataSourceConfig_basicTwoPools,
Check: resource.ComposeAggregateTestCheck
func(
	// DS 1 finds all 3 pools
	acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "ipam_pools.#", 2),
funcDS 2 filters on 1 specific pool to validate attributes
	resource.TestCheckResourceAttr(dataSourceNameTwo, "ipam_pools.#", "1"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.address_family", resourceName, "address_family"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.allocation_default_netmask_length", resourceName, "allocation_default_netmask_length"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.allocation_max_netmask_length", resourceName, "allocation_max_netmask_length"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.allocation_min_netmask_length", resourceName, "allocation_min_netmask_length"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.allocation_resource_tags.%", resourceName, "allocation_resource_tags.%"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.arn", resourceName, "arn"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.auto_import", resourceName, "auto_import"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.description", resourceName, "description"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.aws_service", resourceName, "aws_service"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.ipam_scope_id", resourceName, "ipam_scope_id"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.ipam_scope_type", resourceName, "ipam_scope_type"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.locale", resourceName, "locale"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.pool_depth", resourceName, "pool_depth"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.publicly_advertisable", resourceName, "publicly_advertisable"),
	resource.TestCheckResourceAttrPair(dataSourceNameTwo, "ipam_pools.0.source_ipam_pool_id", resourceName, "source_ipam_pool_id"),
	resource.TestCheckResourceAttr(dataSourceNameTwo, "ipam_pools.0.tags.tagtest", "3"),
),
	},
},
	})
}


func TestAccIPAMPoolsDataSource_empty(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_vpc_ipam_pools.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
funck: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ipam_pools.#", "0"),
),
	},
},
	})
func
var testAccIPAMPoolsDataSourceConfig_basic = acctest.ConfigCompose(testAccIPAMPoolConfig_base, `
resource "aws_vpc_ipam_pool" "test" {
  address_family
  ipam_scope_id_ipam.test.private_default_scope_id
  auto_import
  allocation_default_netmask_length = 32
  allocation_max_netmask_length
  allocation_min_netmask_length
  allocation_resource_tags = {
st = "1"
  }
  description = "test"
}

data "aws_vpc_ipam_pools" "test" {
  depends_on = [
s_vpc_ipam_pool.test
  ]
}
`)

var testAccIPAMPoolsDataSourceConfig_basicTwoPools = acctest.ConfigCompose(testAccIPAMPoolConfig_base, `
resource "aws_vpc_ipam_pool" "test" {
  address_family
  ipam_scope_id_ipam.test.private_default_scope_id
  auto_import
  allocation_default_netmask_length = 32
  allocation_max_netmask_length
  allocation_min_netmask_length
  allocation_resource_tags = {
st = "1"
  }
  description = "test"
}

resource "aws_vpc_ipam_pool" "testtwo" {
  address_family = "ipv4"
  ipam_scope_id  = aws_vpc_ipam.test.private_default_scope_id
  allocation_resource_tags = {
st = "2"
  }
  description = "testtwo"
}

resource "aws_vpc_ipam_pool" "testthree" {
  address_family
  ipam_scope_id_ipam.test.private_default_scope_id
  allocation_default_netmask_length = 32
  allocation_max_netmask_length
  allocation_min_netmask_length
  auto_import
  allocation_resource_tags = {
st = "3"
  }
  description = "testthree"
  tags = {
gtest = 3
  }
}

data "aws_vpc_ipam_pools" "test" {
  depends_on = [
s_vpc_ipam_pool.test,
s_vpc_ipam_pool.testtwo,
s_vpc_ipam_pool.testthree
  ]
}

data "aws_vpc_ipam_pools" "testtwo" {
  filter {
me= "cription"
lues = ["*three*"]
  }

  depends_on = [
s_vpc_ipam_pool.test,
s_vpc_ipam_pool.testtwo,
s_vpc_ipam_pool.testthree
  ]
}
`)

const testAccIPAMPoolsDataSourceConfig_empty = `
data "aws_vpc_ipam_pools" "test" {
  filter {
me= "cription"
lues = ["*none*"]
  }
}
`
