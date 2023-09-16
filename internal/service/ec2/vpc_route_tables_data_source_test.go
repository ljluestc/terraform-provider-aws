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
CheckDestroy:stAccCheckVPCDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTablesDataSourceConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_route_tables.by_vpc_id", "ids.#", "2"), // Add the default route table.
funcource.TestCheckResourceAttr("data.aws_route_tables.by_filter", "ids.#", "6"), // Add the default route tables.
	resource.TestCheckResourceAttr("data.aws_route_tables.empty", "ids.#", "0"),
),
	},
},
	})
}


func testAccVPCRouteTablesDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test1" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "test2" {
  cidr_block = "172.16.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test1_public" {
  vpc_id = aws_vpc.test1.id

  tags = {
me
er"
mponent = "Frontend"
  }
}

resource "aws_route_table" "test1_private1" {
  vpc_id = aws_vpc.test1.id

  tags = {
me
ere"
mponent = "Database"
  }
}

resource "aws_route_table" "test1_private2" {
  vpc_id = aws_vpc.test1.id

  tags = {
me
ere"
mponent = "AppServer"
  }
}

resource "aws_route_table" "test2_public" {
  vpc_id = aws_vpc.test2.id

  tags = {
me
er"
mponent = "Frontend"
  }
}

data "aws_route_tables" "by_vpc_id" {
  vpc_id = aws_vpc.test2.id

  depends_on = [aws_route_table.test1_public, aws_route_table.test1_private1, aws_route_table.test1_private2, aws_route_table.test2_public]
}

data "aws_route_tables" "by_tags" {
  tags = {
er = "Public"
  }

  depends_on = [aws_route_table.test1_public, aws_route_table.test1_private1, aws_route_table.test1_private2, aws_route_table.test2_public]
}

data "aws_route_tables" "by_filter" {
  filter {
me= "-id"
lues = [aws_vpc.test1.id, aws_vpc.test2.id]
  }

  depends_on = [aws_route_table.test1_public, aws_route_table.test1_private1, aws_route_table.test1_private2, aws_route_table.test2_public]
}

data "aws_route_tables" "empty" {
  vpc_id = aws_vpc.test2.id

  tags = {
er = "Private"
  }

  depends_on = [aws_route_table.test1_public, aws_route_table.test1_private1, aws_route_table.test1_private2, aws_route_table.test2_public]
}
`, rName)
}
