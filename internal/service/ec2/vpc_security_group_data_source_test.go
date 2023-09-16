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
Config: testAccVPCSecurityGroupDataSourceConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccSecurityGroupCheckDataSource("data.aws_security_group.by_id"),
functAccSecurityGroupCheckDataSource("data.aws_security_group.by_filter"),
	testAccSecurityGroupCheckDataSource("data.aws_security_group.by_name"),
),
	},
},
	})
}


func testAccSecurityGroupCheckDataSource(dataSourceName string) resource.TestCheck
func {
	resourceName := "aws_security_group.test"
funcurn resource.ComposeAggregateTestCheck
funcurce.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
func
}


func testAccVPCSecurityGroupDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "172.16.0.0/16"

  tags = {
func
}

resource "aws_security_group" "test" {
  vpc_id = aws_vpc.test.id
  name[1]q

  tags = {
me = %[1]q
  }
}

data "aws_security_group" "by_id" {
  id = aws_security_group.test.id
}

data "aws_security_group" "by_name" {
  name = aws_security_group.test.name
}

data "aws_security_group" "by_tag" {
  tags = {
me = aws_security_group.test.tags["Name"]
  }
}

data "aws_security_group" "by_filter" {
  filter {
me= "up-name"
lues = [aws_security_group.test.name]
  }
}
`, rName)
}
