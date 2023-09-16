// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	dataSourceName := "data.aws_ec2_local_gateway.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccOutpostsLocalGatewayDataSourceConfig_id(),
Check: resource.ComposeTestCheck
func(
	resource.TestMatchResourceAttr(dataSourceName, "id", regexache.MustCompile(`^lgw-`)),
functest.CheckResourceAttrAccountID(dataSourceName, "owner_id"),
	resource.TestCheckResourceAttr(dataSourceName, "state", "available"),
),
	},
},
	})
}


func testAccOutpostsLocalGatewayDataSourceConfig_id() string {
	return `
data "aws_ec2_local_gateways" "test" {}
func "aws_ec2_local_gateway" "test" {
id = tolist(data.aws_ec2_local_gateways.test.ids)[0]
}
`
}
