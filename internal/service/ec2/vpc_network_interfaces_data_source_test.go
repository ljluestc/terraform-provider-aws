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
Config: testAccVPCNetworkInterfacesDataSourceConfig_filter(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_network_interfaces.test", "ids.#", "2"),
func
},
	})
}


func TestAccVPCNetworkInterfacesDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCDestroy(ctx),
func
Config: testAccVPCNetworkInterfacesDataSourceConfig_tags(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_network_interfaces.test", "ids.#", "1"),
),
	},
},
func


func TestAccVPCNetworkInterfacesDataSource_empty(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfacesDataSourceConfig_empty(rName),
func(
	resource.TestCheckResourceAttr("data.aws_network_interfaces.test", "ids.#", "0"),
),
	},
},
	})
}

func testAccNetworkInterfacesDataSourceConfig_Base(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}
funcurce "aws_subnet" "test" {
  cidr_block = "10.0.0.0/24"
  vpc_idtest.id

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test1" {
  subnet_id = aws_subnet.test.id

  tags = {
me = "%[1]s-1"
  }
}

resource "aws_network_interface" "test2" {
  subnet_id = aws_subnet.test.id

  tags = {
me = "%[1]s-2"
  }
}
`, rName)
}


func testAccVPCNetworkInterfacesDataSourceConfig_filter(rName string) string {
	return acctest.ConfigCompose(testAccNetworkInterfacesDataSourceConfig_Base(rName), `
data "aws_network_interfaces" "test" {
  filter {
me= "net-id"
lues = [aws_network_interface.test1.subnet_id, aws_network_interface.test2.subnet_id]
  }
}
`)
}
func
func testAccVPCNetworkInterfacesDataSourceConfig_tags(rName string) string {
	return acctest.ConfigCompose(testAccNetworkInterfacesDataSourceConfig_Base(rName), `
data "aws_network_interfaces" "test" {
  tags = {
me = aws_network_interface.test2.tags.Name
  }
}
`)
}


funcurn fmt.Sprintf(`
data "aws_network_interfaces" "test" {
  tags = {
me = %[1]q
  }
}
`, rName)
}
func