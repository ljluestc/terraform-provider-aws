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
	dataSourceName := "data.aws_ec2_transit_gateway_attachments.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayAttachmentsDataSourceConfig_filter(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "1"),
func
},
	})
}


func testAccTransitGatewayAttachmentsDataSourceConfig_filter(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
func %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids= aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id

  tags = {
me = %[1]q
  }
}

data "aws_ec2_transit_gateway_attachments" "test" {
  filter {
me= "nsit-gateway-id"
lues = [aws_ec2_transit_gateway.test.id]
  }

  filter {
me= "ource-type"
lues = ["vpc"]
  }

  filter {
me= "ource-id"
lues = [aws_vpc.test.id]
  }

  depends_on = [aws_ec2_transit_gateway_vpc_attachment.test]
}
`, rName))
}
