// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)


func := acctest.Context(t)
	dataSourceName := "data.aws_ec2_public_ipv4_pool.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckPublicIPv4Pools(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testPublicIPv4PoolDataSourceConfig_basic,
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrSet(dataSourceName, "total_address_count"),
func
	},
},
	})
}


func testAccPreCheckPublicIPv4Pools(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

func
	if acctest.PreCheckSkipError(err) {
t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
t.Fatalf("unexpected PreCheck error: %s", err)
	}

	// Ensure there is at least one pool.
	if len(output) == 0 {
t.Skip("skipping since no EC2 Public IPv4 Pools found")
	}
}

const testPublicIPv4PoolDataSourceConfig_basic = `
data "aws_ec2_public_ipv4_pools" "test" {}

data "aws_ec2_public_ipv4_pool" "test" {
  pool_id = data.aws_ec2_public_ipv4_pools.test.pool_ids[0]
}
`
