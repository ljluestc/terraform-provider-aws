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
	dataSourceName := "data.aws_ec2_coip_pools.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccOutpostsCoIPPoolsDataSourceConfig_basic(),
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "pool_ids.#", 0),
func
},
	})
}


func TestAccEC2OutpostsCoIPPoolsDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ec2_coip_pools.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccOutpostsCoIPPoolsDataSourceConfig_filter(),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "pool_ids.#", "1"),
),
	},
},
func


func testAccOutpostsCoIPPoolsDataSourceConfig_basic() string {
	return `
data "aws_ec2_coip_pools" "test" {}
`
}

func testAccOutpostsCoIPPoolsDataSourceConfig_filter() string {
	return `
data "aws_ec2_coip_pools" "all" {}

data "aws_ec2_coip_pools" "test" {
  filter {
me= "p-pool.pool-id"
func
}
`
}
