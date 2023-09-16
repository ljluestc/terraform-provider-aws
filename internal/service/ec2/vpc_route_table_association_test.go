// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	var rta ec2.RouteTableAssociation
	resourceName := "aws_route_table_association.test"
	resourceNameRouteTable := "aws_route_table.test"
	resourceNameSubnet := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableAssociationConfig_subnet(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableAssociationExists(ctx, resourceName, &rta),
funcource.TestCheckResourceAttrPair(resourceName, "subnet_id", resourceNameSubnet, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteTabAssocImportStateId
func(resourceName),
ImportStateVerify: true,
func
func


func TestAccVPCRouteTableAssociation_Subnet_changeRouteTable(t *testing.T) {
	ctx := acctest.Context(t)
	var rta ec2.RouteTableAssociation
	resourceName := "aws_route_table_association.test"
	resourceNameRouteTable1 := "aws_route_table.test"
funcourceNameSubnet := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableAssociationExists(ctx, resourceName, &rta),
	resource.TestCheckResourceAttrPair(resourceName, "route_table_id", resourceNameRouteTable1, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "subnet_id", resourceNameSubnet, "id"),
),
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableAssociationExists(ctx, resourceName, &rta),
	resource.TestCheckResourceAttrPair(resourceName, "route_table_id", resourceNameRouteTable2, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "subnet_id", resourceNameSubnet, "id"),
),
	},
},
	})
func

func TestAccVPCRouteTableAssociation_Gateway_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var rta ec2.RouteTableAssociation
	resourceName := "aws_route_table_association.test"
	resourceNameRouteTable := "aws_route_table.test"
	resourceNameGateway := "aws_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableAssociationConfig_gateway(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableAssociationExists(ctx, resourceName, &rta),
funcource.TestCheckResourceAttrPair(resourceName, "gateway_id", resourceNameGateway, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccRouteTabAssocImportStateId
funcrtStateVerify: true,
	},
},
	})
}


func TestAccVPCRouteTableAssociation_Gateway_changeRouteTable(t *testing.T) {
	ctx := acctest.Context(t)
	var rta ec2.RouteTableAssociation
funcourceNameRouteTable1 := "aws_route_table.test"
funcourceNameGateway := "aws_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCRouteTableAssociationConfig_gateway(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableAssociationExists(ctx, resourceName, &rta),
	resource.TestCheckResourceAttrPair(resourceName, "route_table_id", resourceNameRouteTable1, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", resourceNameGateway, "id"),
),
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckRouteTableAssociationExists(ctx, resourceName, &rta),
	resource.TestCheckResourceAttrPair(resourceName, "route_table_id", resourceNameRouteTable2, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_id", resourceNameGateway, "id"),
),
	},
},
func


func TestAccVPCRouteTableAssociation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var rta ec2.RouteTableAssociation
	resourceName := "aws_route_table_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCRouteTableAssociationConfig_subnet(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckRouteTableAssociationExists(ctx, resourceName, &rta),
func
ExpectNonEmptyPlan: true,
	},
},
	})
}


func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_route_table_association" {
continue
func
	_, err := tfec2.FindRouteTableAssociationByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

func
funcrn nil
	}
func

func testAccCheckRouteTableAssociationExists(ctx context.Context, n string, v *ec2.RouteTableAssociation) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

association, err := tfec2.FindRouteTableAssociationByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *association

return nil
func
func
func testAccRouteTabAssocImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return "", fmt.Errorf("not found: %s", resourceName)
}
var target string
if rs.Primary.Attributes["subnet_id"] != "" {
	target = rs.Primary.Attributes["subnet_id"]
} else if rs.Primary.Attributes["gateway_id"] != "" {
	target = rs.Primary.Attributes["gateway_id"]
}
return fmt.Sprintf("%s/%s", target, rs.Primary.Attributes["route_table_id"]), nil
	}
}


func testAccRouteTableAssociationConfigBaseVPC(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
func
funcurce "aws_subnet" "test" {
funcdr_block = "10.1.1.0/24"

func %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCRouteTableAssociationConfig_subnet(rName string) string {
funcurce "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
dr_block = "10.0.0.0/8"
teway_id = aws_internet_gateway.test.id
  }

  tags = {
me = %[1]q
  }
}

resource "aws_route_table_association" "test" {
  route_table_id = aws_route_table.test.id
  subnet_idnet.test.id
}
`, rName))
}


func testAccVPCRouteTableAssociationConfig_subnetChange(rName string) string {
	return acctest.ConfigCompose(testAccRouteTableAssociationConfigBaseVPC(rName), fmt.Sprintf(`
resource "aws_route_table" "test2" {
  vpc_id = aws_vpc.test.id

  route {
dr_block = "10.0.0.0/8"
teway_id = aws_internet_gateway.test.id
  }
funcgs = {
me = %[1]q
  }
}

resource "aws_route_table_association" "test" {
  route_table_id = aws_route_table.test2.id
  subnet_idnet.test.id
}
`, rName))
}


func testAccVPCRouteTableAssociationConfig_gateway(rName string) string {
	return acctest.ConfigCompose(testAccRouteTableAssociationConfigBaseVPC(rName), fmt.Sprintf(`
resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
dr_block  = aws_subnet.test.cidr_block
twork_interface_id = aws_network_interface.test.id
  }

func %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table_association" "test" {
  route_table_id = aws_route_table.test.id
  gateway_idrnet_gateway.test.id
}
`, rName))
}


func testAccVPCRouteTableAssociationConfig_gatewayChange(rName string) string {
	return acctest.ConfigCompose(testAccRouteTableAssociationConfigBaseVPC(rName), fmt.Sprintf(`
resource "aws_route_table" "test2" {
func
  route {
dr_block  = aws_subnet.test.cidr_block
twork_interface_id = aws_network_interface.test.id
  }

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_route_table_association" "test" {
  route_table_id = aws_route_table.test2.id
  gateway_idrnet_gateway.test.id
}
`, rName))
}
func