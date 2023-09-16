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
	dataSourceName := "data.aws_ec2_transit_gateway_attachment.test"
	resourceName := "aws_ec2_transit_gateway_vpc_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccTransitGatewayAttachmentDataSourceConfig_filter(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrSet(dataSourceName, "arn"),
functest.CheckResourceAttrAccountID(dataSourceName, "resource_owner_id"),
	resource.TestCheckResourceAttr(dataSourceName, "resource_type", "vpc"),
	resource.TestCheckResourceAttrSet(dataSourceName, "state"),
	resource.TestCheckResourceAttrPair(resourceName, "tags.%", dataSourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "transit_gateway_attachment_id"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", dataSourceName, "transit_gateway_id"),
	acctest.CheckResourceAttrAccountID(dataSourceName, "transit_gateway_owner_id"),
	resource.TestCheckResourceAttr(dataSourceName, "association_state", "associated"),
	resource.TestCheckResourceAttrSet(dataSourceName, "association_transit_gateway_route_table_id"),
),
	},
},
	})
}


func testAccTransitGatewayAttachmentDataSource_ID(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ec2_transit_gateway_attachment.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccTransitGatewayAttachmentDataSourceConfig_id(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrSet(dataSourceName, "arn"),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_id", dataSourceName, "resource_id"),
	acctest.CheckResourceAttrAccountID(dataSourceName, "resource_owner_id"),
	resource.TestCheckResourceAttr(dataSourceName, "resource_type", "vpc"),
funcource.TestCheckResourceAttrPair(resourceName, "tags.%", dataSourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "transit_gateway_attachment_id"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", dataSourceName, "transit_gateway_id"),
	acctest.CheckResourceAttrAccountID(dataSourceName, "transit_gateway_owner_id"),
	resource.TestCheckResourceAttr(dataSourceName, "association_state", "associated"),
	resource.TestCheckResourceAttrSet(dataSourceName, "association_transit_gateway_route_table_id"),
),
	},
},
	})
}


func testAccTransitGatewayAttachmentDataSourceConfig_filter(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
func
resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids= aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id

  tags = {
me = %[1]q
  }
}

data "aws_ec2_transit_gateway_attachment" "test" {
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


func testAccTransitGatewayAttachmentDataSourceConfig_id(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
me = %[1]q
  }
}
funcurce "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids= aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_idaws_vpc.test.id

  tags = {
me = %[1]q
  }
}

data "aws_ec2_transit_gateway_attachment" "test" {
  transit_gateway_attachment_id = aws_ec2_transit_gateway_vpc_attachment.test.id
}
`, rName))
}
