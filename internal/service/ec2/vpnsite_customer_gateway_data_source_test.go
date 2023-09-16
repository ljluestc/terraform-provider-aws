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
	dataSourceName := "data.aws_customer_gateway.test"
	resourceName := "aws_customer_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	asn := sdkacctest.RandIntRange(64512, 65534)
	hostOctet := sdkacctest.RandIntRange(1, 254)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCustomerGatewayDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNCustomerGatewayDataSourceConfig_filter(rName, asn, hostOctet),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(resourceName, "arn", dataSourceName, "arn"),
funcource.TestCheckResourceAttrPair(resourceName, "certificate_arn", dataSourceName, "certificate_arn"),
	resource.TestCheckResourceAttrPair(resourceName, "device_name", dataSourceName, "device_name"),
	resource.TestCheckResourceAttrPair(resourceName, "ip_address", dataSourceName, "ip_address"),
	resource.TestCheckResourceAttrPair(resourceName, "tags.%", dataSourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "type"),
),
	},
},
	})
}


func TestAccSiteVPNCustomerGatewayDataSource_id(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_customer_gateway.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	asn := sdkacctest.RandIntRange(64512, 65534)
	hostOctet := sdkacctest.RandIntRange(1, 254)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCustomerGatewayDestroy(ctx),
func
Config: testAccSiteVPNCustomerGatewayDataSourceConfig_id(rName, asn, hostOctet),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(resourceName, "arn", dataSourceName, "arn"),
	resource.TestCheckResourceAttrPair(resourceName, "bgp_asn", dataSourceName, "bgp_asn"),
	resource.TestCheckResourceAttrPair(resourceName, "certificate_arn", dataSourceName, "certificate_arn"),
	resource.TestCheckResourceAttrPair(resourceName, "device_name", dataSourceName, "device_name"),
funcource.TestCheckResourceAttrPair(resourceName, "tags.%", dataSourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "type"),
),
	},
},
	})
}


func testAccSiteVPNCustomerGatewayDataSourceConfig_filter(rName string, asn, hostOctet int) string {
	return fmt.Sprintf(`
resource "aws_customer_gateway" "test" {
bgp_asn%[2]d
ip_address = "50.0.0.%[3]d"
type.1"
funcgs = {
me = %[1]q
}
}

data "aws_customer_gateway" "test" {
filter {
me= ":Name"
lues = [aws_customer_gateway.test.tags.Name]
}
}
`, rName, asn, hostOctet)
}


func testAccSiteVPNCustomerGatewayDataSourceConfig_id(rName string, asn, hostOctet int) string {
	return fmt.Sprintf(`
resource "aws_customer_gateway" "test" {
bgp_asn
ip_address= "50.0.0.%[3]d"
device_name = "test"
typec.1"
funcgs = {
me = %[1]q
}
}

data "aws_customer_gateway" "test" {
id = aws_customer_gateway.test.id
}
`, rName, asn, hostOctet)
}
