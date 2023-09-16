// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
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
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
funcource.TestCheckResourceAttr(resourceName, "context", ""),
	resource.TestCheckResourceAttr(resourceName, "excess_capacity_termination_policy", "termination"),
	resource.TestCheckResourceAttr(resourceName, "fleet_state", "active"),
	resource.TestCheckResourceAttr(resourceName, "fulfilled_capacity", "0"),
	resource.TestCheckResourceAttr(resourceName, "fulfilled_on_demand_capacity", "0"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.launch_template_specification.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "launch_template_config.0.launch_template_specification.0.launch_template_id"),
	resource.TestCheckResourceAttrSet(resourceName, "launch_template_config.0.launch_template_specification.0.version"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.allocation_strategy", "lowestPrice"),
	resource.TestCheckResourceAttr(resourceName, "replace_unhealthy_instances", "false"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.allocation_strategy", "lowestPrice"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.instance_interruption_behavior", "terminate"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.instance_pools_to_use_count", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.0.default_target_capacity_type", "spot"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.0.total_target_capacity", "0"),
	resource.TestCheckResourceAttr(resourceName, "terminate_instances", "false"),
	resource.TestCheckResourceAttr(resourceName, "terminate_instances_with_expiration", "false"),
	resource.TestCheckResourceAttr(resourceName, "type", "maintain"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
func
Config: testAccFleetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceFleet(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccEC2Fleet_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_tags1(rName, "key1", "value1"),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccFleetConfig_tags1(rName, "key2", "value2"),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func TestAccEC2Fleet_excessCapacityTerminationPolicy(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccFleetConfig_excessCapacityTerminationPolicy(rName, "no-termination"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "excess_capacity_termination_policy", "no-termination"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_excessCapacityTerminationPolicy(rName, "termination"),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "excess_capacity_termination_policy", "termination"),
),
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateLaunchTemplateSpecification_launchTemplateID(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
funcnchTemplateResourceName2 := "aws_launch_template.test2"
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccFleetConfig_launchTemplateID(rName, launchTemplateResourceName1),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.launch_template_specification.#", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.launch_template_id", launchTemplateResourceName1, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.version", launchTemplateResourceName1, "latest_version"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateID(rName, launchTemplateResourceName2),
Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.launch_template_specification.#", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.launch_template_id", launchTemplateResourceName2, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.version", launchTemplateResourceName2, "latest_version"),
),
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateLaunchTemplateSpecification_launchTemplateName(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	launchTemplateResourceName1 := "aws_launch_template.test1"
funcourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateName(rName, launchTemplateResourceName1),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.launch_template_specification.#", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.launch_template_name", launchTemplateResourceName1, "name"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.version", launchTemplateResourceName1, "latest_version"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
	{
Config: testAccFleetConfig_launchTemplateName(rName, launchTemplateResourceName2),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
funcource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.launch_template_name", launchTemplateResourceName2, "name"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.version", launchTemplateResourceName2, "latest_version"),
),
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateLaunchTemplateSpecification_version(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	var launchTemplate ec2.LaunchTemplate
	launchTemplateResourceName := "aws_launch_template.test"
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateVersion(rName, "t3.micro"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, launchTemplateResourceName, &launchTemplate),
	resource.TestCheckResourceAttr(launchTemplateResourceName, "instance_type", "t3.micro"),
	resource.TestCheckResourceAttr(launchTemplateResourceName, "latest_version", "1"),
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.launch_template_specification.#", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.launch_template_id", launchTemplateResourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.version", launchTemplateResourceName, "latest_version"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
func
Config: testAccFleetConfig_launchTemplateVersion(rName, "t3.small"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, launchTemplateResourceName, &launchTemplate),
	resource.TestCheckResourceAttr(launchTemplateResourceName, "instance_type", "t3.small"),
	resource.TestCheckResourceAttr(launchTemplateResourceName, "latest_version", "2"),
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.launch_template_specification.#", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.launch_template_id", launchTemplateResourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.launch_template_specification.0.version", launchTemplateResourceName, "latest_version"),
),
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_availabilityZone(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	availabilityZonesDataSourceName := "data.aws_availability_zones.available"
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideAvailabilityZone(rName, 0),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.override.0.availability_zone", availabilityZonesDataSourceName, "names.0"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideAvailabilityZone(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
funcource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.override.0.availability_zone", availabilityZonesDataSourceName, "names.1"),
),
	},
},
	})
}

// Pending AWS to provide this attribute back in the `Describe` call.
func TestAccEC2Fleet_LaunchTemplateOverride_imageId(t *testing.T) {
// 	ctx := acctest.Context(t)
// 	var fleet1 ec2.FleetData
// 	awsAmiDataSourceName := "data.aws_ami.amz2"
// 	resourceName := "aws_ec2_fleet.test"
// 	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

// 	resource.ParallelTest(t, resource.TestCase{
// PreCheck:  
func() { acctest.PreCheck(t); testAccPreCheckFleet(ctx, t) },
// ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
// ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
// CheckDestroy:stAccCheckFleetDestroy(ctx),
// Steps: []resource.TestStep{
// 	{
// Config: testAccFleetConfig_launchTemplateOverrideImageId(rName),
func(
// 	testAccCheckFleetExists(ctx, resourceName, &fleet1),
// 	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
// 	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
// 	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.override.0.image_id", awsAmiDataSourceName, "id"),
// ),
// 	},
// },
// 	})
// }


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_memoryMiBAndVCPUCount(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_mib {
00

t {


Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_mib.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_mib.0.min", "500"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.vcpu_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.vcpu_count.0.min", "1"),
),
	},
	{
funcrtState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_mib {
func

t {



Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_mib.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_mib.0.max", "24000"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.vcpu_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.vcpu_count.0.min", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.vcpu_count.0.max", "8"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_acceleratorCount(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_count {


b {
0

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_count.0.min", "1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_count {



func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_count.0.min", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_count.0.max", "4"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_count {


b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_count.0.max", "0"),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_acceleratorManufacturers(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_manufacturers = ["amd"]
b {
00

t {


func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_manufacturers.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_manufacturers.*", "amd"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_manufacturers = ["amazon-web-services", "amd", "nvidia", "xilinx"]
b {
00
func


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
func
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_manufacturers.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_manufacturers.*", "amazon-web-services"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_manufacturers.*", "amd"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_manufacturers.*", "nvidia"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_manufacturers.*", "xilinx"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
func
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_acceleratorNames(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_names = ["a100"]
b {
00

t {


Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.*", "a100"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_names = ["a100", "v100", "k80", "t4", "m60", "radeon-pro-v520", "vu9p"]
b {
00

t {


func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.#", "7"),
funcource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.*", "v100"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.*", "k80"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.*", "t4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.*", "m60"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.*", "radeon-pro-v520"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_names.*", "vu9p"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
func


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_acceleratorTotalMemoryMiB(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_total_memory_mib {
000

b {
00

t {


func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_total_memory_mib.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_total_memory_mib.0.min", "1000"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_total_memory_mib {
4000

b {
00

t {


func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_total_memory_mib.#", "1"),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_total_memory_mib {
000
4000

b {
00

func

Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_total_memory_mib.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_total_memory_mib.0.min", "1000"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_total_memory_mib.0.max", "24000"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_acceleratorTypes(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_types = ["fpga"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_types.*", "fpga"),
),
	},
	{
funcrtState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_types = ["fpga", "gpu", "inference"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_types.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_types.*", "fpga"),
funcource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.accelerator_types.*", "inference"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_allowedInstanceTypes(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`allowed_instance_types = ["m4.large"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.allowed_instance_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.allowed_instance_types.*", "m4.large"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`allowed_instance_types = ["m4.large", "m5.*", "m6*"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.allowed_instance_types.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.allowed_instance_types.*", "m4.large"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.allowed_instance_types.*", "m5.*"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.allowed_instance_types.*", "m6*"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
func

func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_bareMetal(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
funcre_metal = "excluded"
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.bare_metal", "excluded"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`bare_metal = "included"
func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.bare_metal", "included"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
func
00

t {


Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.bare_metal", "required"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
func


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_baselineEBSBandwidthMbps(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`baseline_ebs_bandwidth_mbps {
0

b {
00

t {

funck: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.baseline_ebs_bandwidth_mbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.baseline_ebs_bandwidth_mbps.0.min", "10"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`baseline_ebs_bandwidth_mbps {
0000

b {
00

func

Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.baseline_ebs_bandwidth_mbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.baseline_ebs_bandwidth_mbps.0.max", "20000"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
funcseline_ebs_bandwidth_mbps {
0
0000

b {
00

func

Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.baseline_ebs_bandwidth_mbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.baseline_ebs_bandwidth_mbps.0.min", "10"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.baseline_ebs_bandwidth_mbps.0.max", "20000"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_burstablePerformance(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`burstable_performance = "excluded"
b {
00

t {


func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.burstable_performance", "excluded"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`burstable_performance = "included"
b {
000

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.burstable_performance", "included"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`burstable_performance = "required"
b {
000

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.burstable_performance", "required"),
),
	},
	{
funcrtState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_cpuManufacturers(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`cpu_manufacturers = ["amd"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.cpu_manufacturers.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.cpu_manufacturers.*", "amd"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`cpu_manufacturers = ["amazon-web-services", "amd", "intel"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.cpu_manufacturers.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.cpu_manufacturers.*", "amazon-web-services"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.cpu_manufacturers.*", "amd"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.cpu_manufacturers.*", "intel"),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_excludedInstanceTypes(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`excluded_instance_types = ["t2.nano"]
b {
func
t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.excluded_instance_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.excluded_instance_types.*", "t2.nano"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`excluded_instance_types = ["t2.nano", "t3*", "t4g.*"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.excluded_instance_types.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.excluded_instance_types.*", "t2.nano"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.excluded_instance_types.*", "t3*"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.excluded_instance_types.*", "t4g.*"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_instanceGenerations(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`instance_generations = ["current"]
b {
00

t {

funck: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.instance_generations.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.instance_generations.*", "current"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`instance_generations = ["current", "previous"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.instance_generations.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.instance_generations.*", "current"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.instance_generations.*", "previous"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_localStorage(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage = "excluded"
b {
00
func


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
func
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.local_storage", "excluded"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage = "included"
b {
func
t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.local_storage", "included"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage = "required"
b {
00

func

Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.local_storage", "required"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}

func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_localStorageTypes(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage_types = ["hdd"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.local_storage_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.local_storage_types.*", "hdd"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage_types = ["hdd", "ssd"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.local_storage_types.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.local_storage_types.*", "hdd"),
	resource.TestCheckTypeSetElemAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.local_storage_types.*", "ssd"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_memoryGiBPerVCPU(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_gib_per_vcpu {
.5

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_gib_per_vcpu.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_gib_per_vcpu.0.min", "0.5"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_gib_per_vcpu {
.5

b {
00

t {

funck: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_gib_per_vcpu.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_gib_per_vcpu.0.max", "9.5"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_gib_per_vcpu {
.5
.5

b {
00

func

Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_gib_per_vcpu.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_gib_per_vcpu.0.min", "0.5"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.memory_gib_per_vcpu.0.max", "9.5"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
func

func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_networkBandwidthGbps(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_bandwidth_gbps {
	 = 1.5

b {
00

t {


func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_bandwidth_gbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_bandwidth_gbps.0.min", "1.5"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_bandwidth_gbps {
00

b {
00

t {


Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_bandwidth_gbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_bandwidth_gbps.0.max", "200"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_bandwidth_gbps {
.5
50

b {
00

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_bandwidth_gbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_bandwidth_gbps.0.min", "2.5"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_bandwidth_gbps.0.max", "250"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_networkInterfaceCount(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_interface_count {


b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_interface_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_interface_count.0.min", "1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_interface_count {
0

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_interface_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_interface_count.0.max", "10"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_interface_count {

0

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_interface_count.0.min", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.network_interface_count.0.max", "10"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_onDemandMaxPricePercentageOverLowestPrice(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`on_demand_max_price_percentage_over_lowest_price = 50
func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.on_demand_max_price_percentage_over_lowest_price", "50"),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_requireHibernateSupport(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`require_hibernate_support = false
b {
00

t {

funck: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.require_hibernate_support", "false"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`require_hibernate_support = true
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
func
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.require_hibernate_support", "true"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_spotMaxPricePercentageOverLowestPrice(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
func
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.spot_max_price_percentage_over_lowest_price", "75"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceRequirements_totalLocalStorageGB(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet ec2.FleetData
	resourceName := "aws_ec2_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`total_local_storage_gb {
.5

func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.total_local_storage_gb.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.total_local_storage_gb.0.min", "0.5"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`total_local_storage_gb {
0.5

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.total_local_storage_gb.#", "1"),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`total_local_storage_gb {
.5
0.5

b {
00

t {


func(
	testAccCheckFleetExists(ctx, resourceName, &fleet),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),

	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.total_local_storage_gb.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_requirements.0.total_local_storage_gb.0.max", "20.5"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_instanceType(t *testing.T) {
func fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideInstanceType(rName, "t3.small"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_type", "t3.small"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.instance_type", "t3.medium"),
),
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_maxPrice(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideMaxPrice(rName, "1.01"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.max_price", "1.01"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideMaxPrice(rName, "1.02"),
Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.max_price", "1.02"),
),
	},
},
	})
}

// Pending AWS to provide this attribute back in the `Describe` call.
// 
func TestAccEC2Fleet_LaunchTemplateOverride_placement(t *testing.T) {
// 	ctx := acctest.Context(t)
// 	var fleet1 ec2.FleetData
// 	resourceName := "aws_ec2_fleet.test"
// 	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

// 	resource.ParallelTest(t, resource.TestCase{
// PreCheck:  
func() { acctest.PreCheck(t); testAccPreCheckFleet(ctx, t) },
// ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
// ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
// CheckDestroy:stAccCheckFleetDestroy(ctx),
// Steps: []resource.TestStep{
// 	{
// Config: testAccFleetConfig_launchTemplateOverridePlacement(rName),
// Check: resource.ComposeTestCheck
functestAccCheckFleetExists(ctx, resourceName, &fleet1),
// 	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
// 	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
// 	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.placement", "1"),
// 	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.placement.group_name", rName),
// ),
// 	},
// },
// 	})
// }


func TestAccEC2Fleet_LaunchTemplateOverride_priority(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverridePriority(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.priority", "1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
func
Config: testAccFleetConfig_launchTemplateOverridePriority(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.priority", "2"),
),
	},
},
	})
}


func := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "2"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.priority", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.1.priority", "2"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverridePriorityMultiple(rName, 2, 1),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "2"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.priority", "2"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.1.priority", "1"),
),
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_subnetID(t *testing.T) {
	ctx := acctest.Context(t)
funcnetResourceName1 := "aws_subnet.test.0"
	subnetResourceName2 := "aws_subnet.test.1"
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideSubnetID(rName, 0),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.override.0.subnet_id", subnetResourceName1, "id"),
),
	},
	{
funcrtState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideSubnetID(rName, 1),
Check: resource.ComposeTestCheck
func(
functAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "launch_template_config.0.override.0.subnet_id", subnetResourceName2, "id"),
),
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverride_weightedCapacity(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideWeightedCapacity(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideWeightedCapacity(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "1"),
func
	},
},
	})
}


func TestAccEC2Fleet_LaunchTemplateOverrideWeightedCapacity_multiple(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_launchTemplateOverrideWeightedCapacityMultiple(rName, 1, 1),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "2"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.weighted_capacity", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.1.weighted_capacity", "1"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_launchTemplateOverrideWeightedCapacityMultiple(rName, 1, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.#", "2"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.0.weighted_capacity", "1"),
	resource.TestCheckResourceAttr(resourceName, "launch_template_config.0.override.1.weighted_capacity", "2"),
),
	},
func
}


func TestAccEC2Fleet_OnDemandOptions_allocationStrategy(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_onDemandOptionsAllocationStrategy(rName, "prioritized"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.allocation_strategy", "prioritized"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_onDemandOptionsAllocationStrategy(rName, "lowestPrice"),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.allocation_strategy", "lowestPrice"),
),
	},
},
	})
}

// Pending AWS to provide this attribute back in the `Describe` call.
// 
func TestAccEC2Fleet_OnDemandOptions_CapacityReservationOptions(t *testing.T) {
// 	ctx := acctest.Context(t)
// 	var fleet1 ec2.FleetData
funcrName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

// 	resource.ParallelTest(t, resource.TestCase{
// PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
// ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
// ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
// CheckDestroy:stAccCheckFleetDestroy(ctx),
// Steps: []resource.TestStep{
// 	{
// Config: testAccFleetConfig_onDemandOptionsCapacityReservationOptions(rName, "use-capacity-reservations-first"),
// Check: resource.ComposeTestCheck
func(
funcresource.TestCheckResourceAttr(resourceName, "on_demand_options.#", "1"),
// 	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.capacity_reservation_options.#", "1"),
// 	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.capacity_reservation_options.0.usage_strategy", "use-capacity-reservations-first"),
// ),
// 	},
// },
// 	})
// }
func
func TestAccEC2Fleet_OnDemandOptions_MaxTotalPrice(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_onDemandOptionsMaxTotalPrice(rName, "1.0"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.max_total_price", "1.0"),
),
	},
},
	})
func

func TestAccEC2Fleet_OnDemandOptions_MinTargetCapacity(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccFleetConfig_onDemandOptionsMinTargetCapacity(rName, "1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.min_target_capacity", "1"),
func
},
	})
}


func TestAccEC2Fleet_OnDemandOptions_SingleAvailabilityZone(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_onDemandOptionsSingleAvailabilityZone(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.single_availability_zone", "true"),
func
},
	})
}


func TestAccEC2Fleet_OnDemandOptions_SingleInstanceType(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_onDemandOptionsSingleInstanceType(rName, true),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "on_demand_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "on_demand_options.0.single_instance_type", "true"),
),
	},
},
	})
}

func TestAccEC2Fleet_replaceUnhealthyInstances(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_replaceUnhealthyInstances(rName, true),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "replace_unhealthy_instances", "true"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_replaceUnhealthyInstances(rName, false),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "replace_unhealthy_instances", "false"),
),
	},
},
	})
func

func TestAccEC2Fleet_SpotOptions_allocationStrategy(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_spotOptionsAllocationStrategy(rName, "diversified"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "spot_options.#", "1"),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
funcig: testAccFleetConfig_spotOptionsAllocationStrategy(rName, "lowestPrice"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "spot_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.allocation_strategy", "lowestPrice"),
),
func
	})
}


func TestAccEC2Fleet_SpotOptions_capacityRebalance(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData

	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funclacementStrategy := "launch-before-terminate"
	terminationDelay := "120"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccFleetConfig_spotOptionsCapacityRebalance(rName, allocationStrategy, replacementStrategy, terminationDelay),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "spot_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.allocation_strategy", allocationStrategy),
funcource.TestCheckResourceAttr(resourceName, "spot_options.0.maintenance_strategies.0.capacity_rebalance.0.termination_delay", terminationDelay),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
func

func TestAccEC2Fleet_capacityRebalanceInvalidType(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config:eetConfig_invalidTypeForCapacityRebalance(rName),
ExpectError: regexache.MustCompile(`Capacity Rebalance maintenance strategies can only be specified for fleets of type maintain`),
	},
func
}


func TestAccEC2Fleet_SpotOptions_instanceInterruptionBehavior(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_spotOptionsInstanceInterruptionBehavior(rName, "stop"),
Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "spot_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.instance_interruption_behavior", "stop"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_spotOptionsInstanceInterruptionBehavior(rName, "terminate"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "spot_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.instance_interruption_behavior", "terminate"),
),
func
	})
}


func TestAccEC2Fleet_SpotOptions_instancePoolsToUseCount(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
func
Config: testAccFleetConfig_spotOptionsInstancePoolsToUseCount(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "spot_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.instance_pools_to_use_count", "2"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
func
Config: testAccFleetConfig_spotOptionsInstancePoolsToUseCount(rName, 3),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "spot_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "spot_options.0.instance_pools_to_use_count", "3"),
),
	},
},
func


func TestAccEC2Fleet_TargetCapacitySpecification_defaultTargetCapacityType(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.0.default_target_capacity_type", "on-demand"),
),
	},
	{
Config: testAccFleetConfig_targetCapacitySpecificationDefaultTargetCapacityType(rName, "spot"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.0.default_target_capacity_type", "spot"),
func
},
	})
}


func TestAccEC2Fleet_TargetCapacitySpecificationDefaultTargetCapacityType_onDemand(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_targetCapacitySpecificationDefaultTargetCapacityType(rName, "on-demand"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.#", "1"),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
func
}


func TestAccEC2Fleet_TargetCapacitySpecificationDefaultTargetCapacityType_spot(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_targetCapacitySpecificationDefaultTargetCapacityType(rName, "spot"),
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.0.default_target_capacity_type", "spot"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_TargetCapacitySpecification_totalTargetCapacity(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccFleetConfig_targetCapacitySpecificationTotalTargetCapacity(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.0.total_target_capacity", "1"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_targetCapacitySpecificationTotalTargetCapacity(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetNotRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.0.total_target_capacity", "2"),
),
	},
func
}


func TestAccEC2Fleet_TargetCapacitySpecification_targetCapacityUnitType(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	targetCapacityUnitType := "vcpu"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_targetCapacitySpecificationTargetCapacityUnitType(rName, 1, targetCapacityUnitType),
Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "target_capacity_specification.0.target_capacity_unit_type", targetCapacityUnitType),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"terminate_instances"},
	},
},
	})
}


func TestAccEC2Fleet_terminateInstancesWithExpiration(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1, fleet2 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_terminateInstancesExpiration(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "terminate_instances_with_expiration", "true"),
),
	},
funcurceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	{
Config: testAccFleetConfig_terminateInstancesExpiration(rName, false),
Check: resource.ComposeTestCheck
functAccCheckFleetExists(ctx, resourceName, &fleet2),
	testAccCheckFleetRecreated(&fleet1, &fleet2),
	resource.TestCheckResourceAttr(resourceName, "terminate_instances_with_expiration", "false"),
),
	},
},
	})
}
func
func TestAccEC2Fleet_type(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	excessCapacityTerminationPolicy := "termination"
	fleetType := "maintain"
	terminateInstances := false
funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_type(rName, fleetType, excessCapacityTerminationPolicy, terminateInstances),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "type", fleetType),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"terminate_instances"},
	},
	// This configuration will fulfill immediately, skip until ValidFrom is implemented
func	Config: testAccFleetConfig_type(rName, "request"),
	// 	Check: resource.ComposeTestCheck
func(
	// testAccCheckFleetExists(resourceName, &fleet2),
	// testAccCheckFleetRecreated(&fleet1, &fleet2),
	// resource.TestCheckResourceAttr(resourceName, "type", "request"),
	// 	),
	// },
func
}


func TestAccEC2Fleet_type_instant(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	fleetType := "instant"
	totalTargetCapacity := "2"
	terminateInstances := true
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccFleetConfig_type_instant(rName, fleetType, terminateInstances, totalTargetCapacity),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "type", fleetType),
	resource.TestCheckResourceAttr(resourceName, "fleet_instance_set.#", "1"),
funcource.TestCheckResourceAttrSet(resourceName, "fleet_instance_set.0.instance_ids.0"),
	resource.TestCheckResourceAttrSet(resourceName, "fleet_instance_set.0.instance_type"),
	resource.TestCheckResourceAttrSet(resourceName, "fleet_instance_set.0.lifecycle"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"terminate_instances"},
	},
	// This configuration will fulfill immediately, skip until ValidFrom is implemented
	// {
	// 	Config: testAccFleetConfig_type(rName, "request"),
	// 	Check: resource.ComposeTestCheck
func(
	// testAccCheckFleetExists(resourceName, &fleet2),
	// testAccCheckFleetRecreated(&fleet1, &fleet2),
	// resource.TestCheckResourceAttr(resourceName, "type", "request"),
	// 	),
	// },
},
	})
}

// Test for the bug described in https://github.com/hashicorp/terraform-provider-aws/issues/6777
func TestAccEC2Fleet_templateMultipleNetworkInterfaces(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_multipleNetworkInterfaces(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "type", "maintain"),
	testAccCheckFleetHistory(ctx, resourceName, "The associatePublicIPAddress parameter cannot be specified when launching with multiple network interfaces"),
),
	},
},
	})
}


func TestAccEC2Fleet_validFrom(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
	resourceName := "aws_ec2_fleet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	validFrom := "1970-01-01T00:00:00Z"
funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccFleetConfig_validFrom(rName, validFrom),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "valid_from", validFrom),
func
},
	})
}


func TestAccEC2Fleet_validUntil(t *testing.T) {
	ctx := acctest.Context(t)
	var fleet1 ec2.FleetData
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	validUntil := "1970-01-01T00:00:00Z"
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckFleet(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFleetDestroy(ctx),
func
Config: testAccFleetConfig_validUntil(rName, validUntil),
Check: resource.ComposeTestCheck
func(
	testAccCheckFleetExists(ctx, resourceName, &fleet1),
	resource.TestCheckResourceAttr(resourceName, "valid_until", validUntil),
),
	},
},
	})
}


func testAccCheckFleetHistory(ctx context.Context, resourceName string, errorMsg string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
func
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return fmt.Errorf("Not found: %s", resourceName)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Fleet ID is set")
func
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

input := &ec2.DescribeFleetHistoryInput{
	FleetId:.String(rs.Primary.ID),
	StartTime: aws.Time(time.Now().Add(time.Hour * -2)),
}

func
if err != nil {
	return err
}

if output == nil {
	return fmt.Errorf("EC2 Fleet history not found")
}

if output.HistoryRecords == nil {
	return fmt.Errorf("No fleet history records found for fleet %s", rs.Primary.ID)
}

for _, record := range output.HistoryRecords {
funcinue
	}
	if strings.Contains(aws.StringValue(record.EventInformation.EventDescription), errorMsg) {
return fmt.Errorf("Error %s found in fleet history event", errorMsg)
	}
}

return nil
	}
}

func testAccCheckFleetExists(ctx context.Context, n string, v *ec2.FleetData) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
func

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindFleetByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
func
*v = *output

return nil
	}
}


func testAccCheckFleetDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
funcinue
	}

	_, err := tfec2.FindFleetByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
func
	return fmt.Errorf("EC2 Fleet %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckFleetNotRecreated(i, j *ec2.FleetData) resource.TestCheck
funcurn 
func(s *terraform.State) error {
if !aws.TimeValue(i.CreateTime).Equal(aws.TimeValue(j.CreateTime)) {
	return errors.New("EC2 Fleet was recreated")
}

return nil
	}
func

func testAccCheckFleetRecreated(i, j *ec2.FleetData) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if aws.TimeValue(i.CreateTime).Equal(aws.TimeValue(j.CreateTime)) {
	return errors.New("EC2 Fleet was not recreated")
}

return nil
	}
}


func testAccPreCheckFleet(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	input := &ec2.DescribeFleetsInput{
MaxResults: aws.Int64(1),
func
	_, err := conn.DescribeFleetsWithContext(ctx, input)

	if acctest.PreCheckSkipError(err) {
t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
t.Fatalf("unexpected PreCheck error: %s", err)
	}
}

func testAccFleetConfig_BaseLaunchTemplate(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = "t3.micro"
  name = %[1]q
}
`, rName))
func

func testAccFleetConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), `
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
func
  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }
}
`)
}

func testAccFleetConfig_multipleNetworkInterfaces(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
func
resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol = p"
om_port= 2
funclocks = ["0.0.0.0/0"]
  }

  egress {
otocol = "
om_port= 0
_port
dr_blocks = ["0.0.0.0/0"]
  }

func %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_idbnet.test[0].id
  security_groups = [aws_security_group.test.id]

func %[1]q
  }
}

resource "aws_launch_template" "test" {
  name
  image_id = data.aws_ami.amzn-ami-minimal-hvm-ebs.id

funcptions {
ance_type = "persistent"

rket_type = "spot"
  }

  network_interfaces {
vice_index = 0
lete_on_termination = true
twork_interface_id  = aws_network_interface.test.id
func
func_index = 1
lete_on_termination = true
func
}

resource "aws_ec2_fleet" "test" {
  terminate_instances = true

  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


allow to choose from several instance types if there is no spot capacity for some type
erride {
type = "t2.micro"

erride {
type = "t3.micro"

erride {
type = "t3.small"

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccFleetConfig_excessCapacityTerminationPolicy(rName, excessCapacityTerminationPolicy string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  excess_capacity_termination_policy = %[2]q

  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

func
funct_target_capacity_type = "spot"
tal_target_capacity
func
  tags = {
me = %[1]q
  }
}
`, rName, excessCapacityTerminationPolicy))
}


func testAccFleetConfig_launchTemplateID(rName, launchTemplateResourceName string) string {
	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`
resource "aws_launch_template" "test1" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = "t3.micro"
  name = "%[1]s1"
}

resource "aws_launch_template" "test2" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = "t3.micro"
  name = "%[1]s2"
}

resource "aws_ec2_fleet" "test" {
  launch_template_config {
functe_id = %[2]s.id
func
  }
funcrget_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, launchTemplateResourceName))
}


func testAccFleetConfig_launchTemplateName(rName, launchTemplateResourceName string) string {
	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`
resource "aws_launch_template" "test1" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = "t3.micro"
  name = "%[1]s1"
}

resource "aws_launch_template" "test2" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = "t3.micro"
  name = "%[1]s2"
}
funcurce "aws_ec2_fleet" "test" {
func_template_specification {
mplate_name = %[2]s.name
func
  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
func
func

func testAccFleetConfig_launchTemplateVersion(rName, instanceType string) string {
	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = %[2]q
  name = %[1]q
}

resource "aws_ec2_fleet" "test" {
  launch_template_config {
functe_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, instanceType))
}


func testAccFleetConfig_launchTemplateOverrideAvailabilityZone(rName string, availabilityZoneIndex int) string {
funcAccFleetConfig_BaseLaunchTemplate(rName),
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


erride {
func
  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, availabilityZoneIndex))
}

// Pending AWS to provide this attribute back in the `Describe` call.
// 
func testAccFleetConfig_launchTemplateOverrideImageId(rName string) string {
// 	return acctest.ConfigCompose(
funccctest.ConfigAvailableAZsNoOptIn(),
// fmt.Sprintf(`
// resource "aws_ec2_fleet" "test" {
//nch_template_config {
//plate_specification {
//emplate_id = aws_launch_template.test.id
//  = aws_launch_template.test.latest_version
//

//
// = data.aws_ami.amz2.id
//
//

//get_capacity_specification {
//rget_capacity_type = "spot"
//et_capacity
//

//s = {
//]q
//
// }

// data "aws_ami" "amz2" {
// 	most_recent = true

// 	filter {
// 	  namename"
// 	  values = ["amzn2-ami-hvm-*-x86_64-ebs"]
// 	}
// 	owners = ["amazon"]
// }
// `, rName))
// }


func testAccFleetConfig_launchTemplateOverrideInstanceRequirements(rName, instanceRequirements string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


erride {
requirements {



  }

  target_capacity_specification {
fault_target_capacity_type = "on-demand"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, instanceRequirements))
}


func testAccFleetConfig_launchTemplateOverrideInstanceType(rName, instanceType string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


erride {
type = %[2]q

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, instanceType))
}


func testAccFleetConfig_launchTemplateOverrideMaxPrice(rName, maxPrice string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version
func
erride {
 = %[2]q

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, maxPrice))
}

// Pending AWS to provide this attribute back in the `Describe` call.
// 
func testAccFleetConfig_launchTemplateOverridePlacement(rName string) string {
// 	return acctest.ConfigCompose(
// testAccFleetConfig_BaseLaunchTemplate(rName),
// acctest.ConfigAvailableAZsNoOptIn(),
// fmt.Sprintf(`
funch_template_config {
//plate_specification {
//emplate_id = aws_launch_template.test.id
//  = aws_launch_template.test.latest_version
//

//
// placement {
// 	group_name = aws_launch_template.test.name
// }
//
//

//get_capacity_specification {
//rget_capacity_type = "spot"
//et_capacity
//

//s = {
//]q
//
// }

// resource "aws_placement_group" "test" {
// 	name
// 	strategy = "cluster"
//
// `, rName))
// }


func testAccFleetConfig_launchTemplateOverridePriority(rName string, priority int) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
functe_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


erride {
= %[2]d

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, priority))
}


func testAccFleetConfig_launchTemplateOverridePriorityMultiple(rName string, priority1, priority2 int) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


erride {
type = aws_launch_template.test.instance_type
= %[

funcde {
type = "t3.small"
= %[

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, priority1, priority2))
}


func testAccFleetConfig_launchTemplateOverrideSubnetID(rName string, subnetIndex int) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), acctest.ConfigVPCWithSubnets(rName, 2), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


erride {
func
  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, subnetIndex))
}


func testAccFleetConfig_launchTemplateOverrideWeightedCapacity(rName string, weightedCapacity int) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


erride {
capacity = %[2]d

  }

  target_capacity_specification {
functarget_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, weightedCapacity))
}


func testAccFleetConfig_launchTemplateOverrideWeightedCapacityMultiple(rName string, weightedCapacity1, weightedCapacity2 int) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


erride {
type= awsch_tee.test.instance_type
capacity = %[2]d


erride {
type= "t3l"
capacity = %[3]d

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
func


func testAccFleetConfig_onDemandOptionsAllocationStrategy(rName, allocationStrategy string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  on_demand_options {
location_strategy = %[2]q
  }

  target_capacity_specification {
fault_target_capacity_type = "on-demand"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, allocationStrategy))
}

func
func testAccFleetConfig_onDemandOptionsCapacityReservationOptions(rName, usageStrategy string) string {
// 	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
// resource "aws_ec2_fleet" "test" {
//nch_template_config {
//plate_specification {
//emplate_id = aws_launch_template.test.id
//  = aws_launch_template.test.latest_version
//
//

//demand_options {
//eservation_options {
//rategy = %[2]q
//
//

//get_capacity_specification {
//rget_capacity_type = "on-demand"
//et_capacity
//
//minate_instances = true
//e = "instant"

//s = {
//]q
//
func, rName, usageStrategy))
// }


func testAccFleetConfig_onDemandOptionsMaxTotalPrice(rName, maxTotalPrice string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  on_demand_options {
x_total_price = %[2]q
  }

  target_capacity_specification {
fault_target_capacity_type = "on-demand"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, maxTotalPrice))
func

func testAccFleetConfig_onDemandOptionsMinTargetCapacity(rName, minTargetcapcity string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  on_demand_options {
n_target_capacity
ngle_availability_zone = true
  }

  target_capacity_specification {
fault_target_capacity_type = "on-demand"
tal_target_capacity
  }

  terminate_instances = true
  type = "instant"

  tags = {
me = %[1]q
  }
}
`, rName, minTargetcapcity))
}


func testAccFleetConfig_onDemandOptionsSingleAvailabilityZone(rName string, singleAZ bool) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
functe_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  on_demand_options {
ngle_availability_zone = %[2]t
  }

  target_capacity_specification {
fault_target_capacity_type = "on-demand"
tal_target_capacity
  }

  terminate_instances = true
  type = "instant"

  tags = {
me = %[1]q
  }
}
`, rName, singleAZ))
}


func testAccFleetConfig_onDemandOptionsSingleInstanceType(rName string, singleInstanceType bool) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
funcunch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  on_demand_options {
ngle_instance_type = %[2]t
  }

  target_capacity_specification {
fault_target_capacity_type = "on-demand"
tal_target_capacity
  }

  terminate_instances = true
  type = "instant"

  tags = {
me = %[1]q
  }
}
`, rName, singleInstanceType))
}


func testAccFleetConfig_replaceUnhealthyInstances(rName string, replaceUnhealthyInstances bool) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  replace_unhealthy_instances = %[2]t

  launch_template_config {
functe_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, replaceUnhealthyInstances))
}


func testAccFleetConfig_spotOptionsAllocationStrategy(rName, allocationStrategy string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }
funcot_options {
location_strategy = %[2]q
  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, allocationStrategy))
}


func testAccFleetConfig_spotOptionsCapacityRebalance(rName, allocationStrategy, replacementStrategy, terminationDelay string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

function_strategy = %[2]q
intenance_strategies {
rebalance {
ment_strategy = %[3]q
tion_delay%[4]s


  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, allocationStrategy, replacementStrategy, terminationDelay))
}


func testAccFleetConfig_invalidTypeForCapacityRebalance(rName string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  type = "request"

  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }
funcot_options {
location_strategy = "diversified"
intenance_strategies {
rebalance {
ment_strategy = "launch"


  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccFleetConfig_spotOptionsInstanceInterruptionBehavior(rName, instanceInterruptionBehavior string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
func
  }

  spot_options {
stance_interruption_behavior = %[2]q
  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, instanceInterruptionBehavior))
}


func testAccFleetConfig_spotOptionsInstancePoolsToUseCount(rName string, instancePoolsToUseCount int) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  spot_options {
func

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, instancePoolsToUseCount))
}


func testAccFleetConfig_tags1(rName, key1, value1 string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  tags = {
1]q = %[2]q
func
  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }
}
`, key1, value1))
}


func testAccFleetConfig_tags2(rName, key1, value1, key2, value2 string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  tags = {
1]q = %[2]q
3]q = %[4]q
  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }
}
`, key1, value1, key2, value2))
func

func testAccFleetConfig_targetCapacitySpecificationDefaultTargetCapacityType(rName, defaultTargetCapacityType string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  target_capacity_specification {
fault_target_capacity_type = %[2]q
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, defaultTargetCapacityType))
}


func testAccFleetConfig_targetCapacitySpecificationTotalTargetCapacity(rName string, totalTargetCapacity int) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  terminate_instances = true

func_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, totalTargetCapacity))
}


func testAccFleetConfig_targetCapacitySpecificationTargetCapacityUnitType(rName string, totalTargetCapacity int, targetCapacityUnitType string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  terminate_instances = true

  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version


funcirements {
ator_manufacturers = ["amd"]
mib {
 min = 500

unt {
 min = 1



  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
rget_capacity_unit_type = ]q
  }

  tags = {
me = %[1]q
  }
}
`, rName, totalTargetCapacity, targetCapacityUnitType))
}

func testAccFleetConfig_terminateInstancesExpiration(rName string, terminateInstancesWithExpiration bool) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  terminate_instances_with_expiration = %[2]t

  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  tags = {
me = %[1]q
  }
}
`, rName, terminateInstancesWithExpiration))
}


func testAccFleetConfig_type_instant(rName, fleetType string, terminateInstance bool, totalTargetCapacity string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
funcpe = %[2]q

  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  terminate_instances = %[3]t

  tags = {
me = %[1]q
  }
}
`, rName, fleetType, terminateInstance, totalTargetCapacity))
}


func testAccFleetConfig_type(rName, fleetType string, excessCapacityTerminationPolicy string, terminateInstance bool) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  type = %[2]q
  excess_capacity_termination_policy = %[3]q

  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
func
  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  terminate_instances = %[4]t

  tags = {
me = %[1]q
  }
}
`, rName, fleetType, excessCapacityTerminationPolicy, terminateInstance))
}


func testAccFleetConfig_validFrom(rName, validFrom string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  valid_from = %[2]q
funcgs = {
me = %[1]q
  }
}
`, rName, validFrom))
}


func testAccFleetConfig_validUntil(rName, validUntil string) string {
	return acctest.ConfigCompose(testAccFleetConfig_BaseLaunchTemplate(rName), fmt.Sprintf(`
resource "aws_ec2_fleet" "test" {
  launch_template_config {
unch_template_specification {
mplate_id = aws_launch_template.test.id
 = aws_launch_template.test.latest_version

  }

  target_capacity_specification {
fault_target_capacity_type = "spot"
tal_target_capacity
  }

  valid_until = %[2]q

  tags = {
me = %[1]q
func
`, rName, validUntil))
}
funcfuncfuncfuncfuncfuncfuncfuncfuncfunc