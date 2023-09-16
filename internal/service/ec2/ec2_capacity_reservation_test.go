// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
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
	var cr ec2.CapacityReservation
	availabilityZonesDataSourceName := "data.aws_availability_zones.available"
	resourceName := "aws_ec2_capacity_reservation.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccCapacityReservationConfig_basic,
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
funcource.TestCheckResourceAttrPair(resourceName, "availability_zone", availabilityZonesDataSourceName, "names.0"),
	resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "false"),
	resource.TestCheckResourceAttr(resourceName, "end_date", ""),
	resource.TestCheckResourceAttr(resourceName, "end_date_type", "unlimited"),
	resource.TestCheckResourceAttr(resourceName, "ephemeral_storage", "false"),
	resource.TestCheckResourceAttr(resourceName, "instance_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_match_criteria", "open"),
	resource.TestCheckResourceAttr(resourceName, "instance_platform", "Linux/UNIX"),
	resource.TestCheckResourceAttr(resourceName, "instance_type", "t2.micro"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "placement_group_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tenancy", "default"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccEC2CapacityReservation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var cr ec2.CapacityReservation
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
func
Config: testAccCapacityReservationConfig_basic,
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceCapacityReservation(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccEC2CapacityReservation_ebsOptimized(t *testing.T) {
	ctx := acctest.Context(t)
	var cr ec2.CapacityReservation
	resourceName := "aws_ec2_capacity_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccCapacityReservationConfig_ebsOptimized(rName, true),
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccEC2CapacityReservation_endDate(t *testing.T) {
	ctx := acctest.Context(t)
	var cr ec2.CapacityReservation
	endDate1 := time.Now().UTC().Add(1 * time.Hour).Format(time.RFC3339)
	endDate2 := time.Now().UTC().Add(2 * time.Hour).Format(time.RFC3339)
	resourceName := "aws_ec2_capacity_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccCapacityReservationConfig_endDate(rName, endDate1),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
funcource.TestCheckResourceAttr(resourceName, "end_date_type", "limited"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
funcig: testAccCapacityReservationConfig_endDate(rName, endDate2),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "end_date", endDate2),
	resource.TestCheckResourceAttr(resourceName, "end_date_type", "limited"),
),
	},
},
	})
}


func TestAccEC2CapacityReservation_endDateType(t *testing.T) {
func cr ec2.CapacityReservation
	endDate := time.Now().UTC().Add(12 * time.Hour).Format(time.RFC3339)
	resourceName := "aws_ec2_capacity_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
func
Config: testAccCapacityReservationConfig_endDateType(rName, "unlimited"),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "end_date_type", "unlimited"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
	{
Config: testAccCapacityReservationConfig_endDate(rName, endDate),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
funcource.TestCheckResourceAttr(resourceName, "end_date_type", "limited"),
),
	},
	{
Config: testAccCapacityReservationConfig_endDateType(rName, "unlimited"),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "end_date_type", "unlimited"),
),
	},
},
	})
func

func TestAccEC2CapacityReservation_ephemeralStorage(t *testing.T) {
	ctx := acctest.Context(t)
	var cr ec2.CapacityReservation
	resourceName := "aws_ec2_capacity_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccCapacityReservationConfig_ephemeralStorage(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2CapacityReservation_instanceCount(t *testing.T) {
	ctx := acctest.Context(t)
	var cr ec2.CapacityReservation
	resourceName := "aws_ec2_capacity_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccCapacityReservationConfig_instanceCount(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "instance_count", "1"),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccCapacityReservationConfig_instanceCount(rName, 2),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "instance_count", "2"),
),
	},
},
	})
}


func := acctest.Context(t)
	var cr ec2.CapacityReservation
	resourceName := "aws_ec2_capacity_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "instance_match_criteria", "targeted"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}


func TestAccEC2CapacityReservation_instanceType(t *testing.T) {
	ctx := acctest.Context(t)
	var cr ec2.CapacityReservation
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
func
Config: testAccCapacityReservationConfig_instanceType(rName, "t2.micro"),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "instance_type", "t2.micro"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccCapacityReservationConfig_instanceType(rName, "t2.small"),
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "instance_type", "t2.small"),
),
	},
},
	})
}
func
func TestAccEC2CapacityReservation_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var cr ec2.CapacityReservation
	resourceName := "aws_ec2_capacity_reservation.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccCapacityReservationConfig_tags1("key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccCapacityReservationConfig_tags2("key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
funcource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccCapacityReservationConfig_tags1("key2", "value2"),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func TestAccEC2CapacityReservation_tenancy(t *testing.T) {
	ctx := acctest.Context(t)
	var cr ec2.CapacityReservation
	resourceName := "aws_ec2_capacity_reservation.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckCapacityReservation(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckCapacityReservationDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckCapacityReservationExists(ctx, resourceName, &cr),
	resource.TestCheckResourceAttr(resourceName, "tenancy", "dedicated"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}


func testAccCheckCapacityReservationExists(ctx context.Context, n string, v *ec2.CapacityReservation) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
func

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Capacity Reservation ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

func
if err != nil {
	return err
}

*v = *output

return nil
func


func testAccCheckCapacityReservationDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_capacity_reservation" {
continue
	}

	_, err := tfec2.FindCapacityReservationByID(ctx, conn, rs.Primary.ID)
functfresource.NotFound(err) {
func

funcrn err
	}

	return fmt.Errorf("EC2 Capacity Reservation %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccPreCheckCapacityReservation(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	input := &ec2.DescribeCapacityReservationsInput{
MaxResults: aws.Int64(1),
	}

	_, err := conn.DescribeCapacityReservationsWithContext(ctx, input)

	if acctest.PreCheckSkipError(err) {
t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
func
func
var testAccCapacityReservationConfig_basic = acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), `
funcailability_zone = data.aws_availability_zones.available.names[0]
  instance_count1
  instance_platform = "Linux/UNIX"
  instance_typeo"
}
`)


func testAccCapacityReservationConfig_ebsOptimized(rName string, ebsOptimized bool) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  ebs_optimized
  instance_count1
  instance_platform = "Linux/UNIX"
  instance_typee"

  tags = {
me = %[1]q
  }
}
`, rName, ebsOptimized))
}


func testAccCapacityReservationConfig_endDate(rName, endDate string) string {
funcurce "aws_ec2_capacity_reservation" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  end_date = %[2]q
  end_date_type"
  instance_count1
  instance_platform = "Linux/UNIX"
  instance_typeo"

  tags = {
me = %[1]q
  }
}
`, rName, endDate))
}


func testAccCapacityReservationConfig_endDateType(rName, endDateType string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  end_date_type
  instance_count1
  instance_platform = "Linux/UNIX"
  instance_typeo"

  tags = {
me = %[1]q
  }
funcName, endDateType))
}


func testAccCapacityReservationConfig_ephemeralStorage(rName string, ephemeralStorage bool) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  ephemeral_storage = %[2]t
  instance_count1
  instance_platform = "Linux/UNIX"
  instance_typeum"

  tags = {
me = %[1]q
  }
}
func


func testAccCapacityReservationConfig_instanceCount(rName string, instanceCount int) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  instance_count%[2]d
  instance_platform = "Linux/UNIX"
  instance_typeo"

  tags = {
me = %[1]q
  }
}
`, rName, instanceCount))
}

func testAccCapacityReservationConfig_instanceMatchCriteria(rName, instanceMatchCriteria string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
  availability_zonews_availability_zones.available.names[0]
  instance_count = 1
  instance_platform/UNIX"
  instance_match_criteria = %[2]q
  instance_type  = "t2.micro"

  tags = {
me = %[1]q
  }
}
`, rName, instanceMatchCriteria))
}


funcurn acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  instance_count1
  instance_platform = "Linux/UNIX"
  instance_type

  tags = {
me = %[1]q
  }
}
`, rName, instanceType))
}


func testAccCapacityReservationConfig_tags1(tag1Key, tag1Value string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
funcailability_zone = data.aws_availability_zones.available.names[0]
  instance_count1
  instance_platform = "Linux/UNIX"
  instance_typeo"

  tags = {
1]q = %[2]q
  }
}
`, tag1Key, tag1Value))
}


func testAccCapacityReservationConfig_tags2(tag1Key, tag1Value, tag2Key, tag2Value string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
funcstance_count1
  instance_platform = "Linux/UNIX"
  instance_typeo"

  tags = {
1]q = %[2]q
3]q = %[4]q
  }
}
`, tag1Key, tag1Value, tag2Key, tag2Value))
}


func testAccCapacityReservationConfig_tenancy(rName, tenancy string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
  availability_zone = data.aws_availability_zones.available.names[1]
funcstance_platform = "Linux/UNIX"
  instance_typerge"
  tenancy  = %[2]q

  tags = {
me = %[1]q
  }
}
`, rName, tenancy))
}
funcfuncfunc