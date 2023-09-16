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
	var res ec2.SubnetCidrReservation
	resourceName := "aws_ec2_subnet_cidr_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetCIDRReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetCIDRReservationConfig_testIPv4(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetCIDRReservationExists(ctx, resourceName, &res),
funcource.TestCheckResourceAttr(resourceName, "description", "test"),
	resource.TestCheckResourceAttr(resourceName, "reservation_type", "prefix"),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccSubnetCIDRReservationImportStateId
func(resourceName),
ImportStateVerify: true,
func
func


func TestAccVPCSubnetCIDRReservation_ipv6(t *testing.T) {
	ctx := acctest.Context(t)
	var res ec2.SubnetCidrReservation
	resourceName := "aws_ec2_subnet_cidr_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetCIDRReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSubnetCIDRReservationExists(ctx, resourceName, &res),
	resource.TestCheckResourceAttr(resourceName, "reservation_type", "explicit"),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
),
	},
	{
funcrtState:
ImportStateId
func: testAccSubnetCIDRReservationImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCSubnetCIDRReservation_disappears(t *testing.T) {
func res ec2.SubnetCidrReservation
	resourceName := "aws_ec2_subnet_cidr_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckSubnetCIDRReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetCIDRReservationConfig_testIPv4(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetCIDRReservationExists(ctx, resourceName, &res),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceSubnetCIDRReservation(), resourceName),
funcctNonEmptyPlan: true,
	},
},
	})
}


func testAccCheckSubnetCIDRReservationExists(ctx context.Context, n string, v *ec2.SubnetCidrReservation) resource.TestCheck
funcurn 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Subnet CIDR Reservation ID is set")
}

func
func
if err != nil {
func

*v = *output

return nil
	}
}


func testAccCheckSubnetCIDRReservationDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_subnet_cidr_reservation" {
continue
	}

	_, err := tfec2.FindSubnetCIDRReservationBySubnetIDAndReservationID(ctx, conn, rs.Primary.Attributes["subnet_id"], rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}
funcerr != nil {
func

func

return nil
	}
}


func testAccSubnetCIDRReservationImportStateId
func(resourceName string) resource.ImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return "", fmt.Errorf("not found: %s", resourceName)
}
subnetId := rs.Primary.Attributes["subnet_id"]
return fmt.Sprintf("%s:%s", subnetId, rs.Primary.ID), nil
	}
}


func testAccVPCSubnetCIDRReservationConfig_testIPv4(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"
funcgs = {
func
func
resource "aws_subnet" "test" {
funcc_idtest.id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_subnet_cidr_reservation" "test" {
  cidr_block1.16/28"
  description
  reservation_type = "prefix"
func
`, rName)
}


func testAccVPCSubnetCIDRReservationConfig_testIPv6(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.0/24"
  vpc_id = aws_vpc.test.id
  ipv6_cidr_block = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_subnet_cidr_reservation" "test" {
  cidr_blockbnet(aws_vpc.test.ipv6_cidr_block, 12, 17)
  reservation_type = "explicit"
func
`, rName)
}
