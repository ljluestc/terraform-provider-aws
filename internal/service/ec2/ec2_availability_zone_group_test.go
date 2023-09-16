// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	resourceName := "aws_ec2_availability_zone_group.test"

	// Filter to one Availability Zone Group per Region as Local Zones become available
	// e.g. ensure there are not two us-west-2-XXX when adding to this list
	// (Not including in config to avoid lintignoring entire config.)
	localZone := "us-west-2-lax-1" // lintignore:AWSAT003

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckRegion(t, endpoints.UsWest2RegionID) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:l,
Steps: []resource.TestStep{
	{
Config: testAccAvailabilityZoneGroupConfig_optInStatus(localZone, ec2.AvailabilityZoneOptInStatusOptedIn),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "opt_in_status", ec2.AvailabilityZoneOptInStatusOptedIn),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	// InvalidOptInStatus: Opting out of Local Zones is not currently supported. Contact AWS Support for additional assistance.
	/*
{
	Config: testAccAvailabilityZoneGroupConfig_optInStatus(ec2.AvailabilityZoneOptInStatusNotOptedIn),
	Check: resource.ComposeTestCheck
func(
resource.TestCheckResourceAttr(resourceName, "opt_in_status", ec2.AvailabilityZoneOptInStatusNotOptedIn),
	),
func
	Config: testAccAvailabilityZoneGroupConfig_optInStatus(ec2.AvailabilityZoneOptInStatusOptedIn),
	Check: resource.ComposeTestCheck
func(
resource.TestCheckResourceAttr(resourceName, "opt_in_status", ec2.AvailabilityZoneOptInStatusOptedIn),
	),
},
func
	})
}


func testAccAvailabilityZoneGroupConfig_optInStatus(name, optInStatus string) string {
	return fmt.Sprintf(`
data "aws_availability_zones" "test" {
  all_availability_zones = true

func "group-name"
lues = [


  }
}

resource "aws_ec2_availability_zone_group" "test" {
  # The above group-name filter should ensure one Availability Zone Group per Region
  group_nametolist(data.aws_availability_zones.test.group_names)[0]
  opt_in_status = %[2]q
}
`, name, optInStatus)
}
