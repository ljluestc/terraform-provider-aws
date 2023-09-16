// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/YakDriver/regexache"
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
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_name(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "capacity_reservation_specification.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "cpu_options.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "credit_specification.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "default_version", "1"),
	resource.TestCheckResourceAttr(resourceName, "description", ""),
	resource.TestCheckResourceAttr(resourceName, "disable_api_stop", "false"),
	resource.TestCheckResourceAttr(resourceName, "disable_api_termination", "false"),
	resource.TestCheckResourceAttr(resourceName, "ebs_optimized", ""),
	resource.TestCheckResourceAttr(resourceName, "elastic_gpu_specifications.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "elastic_inference_accelerator.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "enclave_options.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "hibernation_options.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "iam_instance_profile.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "image_id", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_initiated_shutdown_behavior", ""),
	resource.TestCheckResourceAttr(resourceName, "instance_market_options.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "instance_type", ""),
	resource.TestCheckResourceAttr(resourceName, "kernel_id", ""),
	resource.TestCheckResourceAttr(resourceName, "key_name", ""),
	resource.TestCheckResourceAttr(resourceName, "latest_version", "1"),
	resource.TestCheckResourceAttr(resourceName, "license_specification.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "maintenance_options.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "monitoring.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	resource.TestCheckResourceAttr(resourceName, "name_prefix", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "placement.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name_options.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "ram_disk_id", ""),
	resource.TestCheckResourceAttr(resourceName, "security_group_names.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tag_specifications.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "user_data", ""),
	resource.TestCheckResourceAttr(resourceName, "vpc_security_group_ids.#", "0"),
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


func TestAccEC2LaunchTemplate_Name_generated(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
func
Config: testAccLaunchTemplateConfig_nameGenerated(),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	acctest.CheckResourceAttrNameGenerated(resourceName, "name"),
	resource.TestCheckResourceAttr(resourceName, "name_prefix", "terraform-"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_Name_prefix(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_namePrefix("tf-acc-test-prefix-"),
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	acctest.CheckResourceAttrNameFromPrefix(resourceName, "name", "tf-acc-test-prefix-"),
	resource.TestCheckResourceAttr(resourceName, "name_prefix", "tf-acc-test-prefix-"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var launchTemplate ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_name(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &launchTemplate),
func
ExpectNonEmptyPlan: true,
	},
},
	})
}


func := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
func
Config: testAccLaunchTemplateConfig_blockDeviceMappingsEBS(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.device_name", "/dev/xvda"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.encrypted", ""),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.throughput", "0"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_size", "15"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_type", ""),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_BlockDeviceMappingsEBS_deleteOnTermination(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_blockDeviceMappingsEBSDeleteOnTermination(rName, true),
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.device_name", "/dev/xvda"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.delete_on_termination", "true"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_size", "15"),
),
func
Config: testAccLaunchTemplateConfig_blockDeviceMappingsEBSDeleteOnTermination(rName, false),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.device_name", "/dev/xvda"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_size", "15"),
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

func TestAccEC2LaunchTemplate_BlockDeviceMappingsEBS_gp3(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_blockDeviceMappingsEBSGP3(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.iops", "4000"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.throughput", "500"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_size", "15"),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_type", "gp3"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_ebsOptimized(rName, "true"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "true"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_ebsOptimized(rName, "false"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
func
	},
	{
Config: testAccLaunchTemplateConfig_ebsOptimized(rName, "\"true\""),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "true"),
func
	{
Config: testAccLaunchTemplateConfig_ebsOptimized(rName, "\"false\""),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "false"),
),
	},
	{
Config: testAccLaunchTemplateConfig_ebsOptimized(rName, "\"\""),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "ebs_optimized", ""),
),
	},
},
	})
}


func := acctest.Context(t)
	var template1 ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_elasticInferenceAccelerator(rName, "eia1.medium"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template1),
funcource.TestCheckResourceAttr(resourceName, "elastic_inference_accelerator.0.type", "eia1.medium"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_elasticInferenceAccelerator(rName, "eia1.large"),
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template1),
	resource.TestCheckResourceAttr(resourceName, "elastic_inference_accelerator.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "elastic_inference_accelerator.0.type", "eia1.large"),
),
	},
},
	})
func

func TestAccEC2LaunchTemplate_NetworkInterfaces_deleteOnTermination(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_networkInterfacesDeleteOnTermination(rName, "true"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.security_groups.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.delete_on_termination", "true"),
func
	{
Config: testAccLaunchTemplateConfig_networkInterfacesDeleteOnTermination(rName, "false"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.security_groups.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.delete_on_termination", "false"),
),
	},
funcig: testAccLaunchTemplateConfig_networkInterfacesDeleteOnTermination(rName, "\"\""),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.security_groups.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.delete_on_termination", ""),
),
func
Config: testAccLaunchTemplateConfig_networkInterfacesDeleteOnTermination(rName, "null"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.security_groups.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.delete_on_termination", ""),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}

func TestAccEC2LaunchTemplate_data(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_data(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "disable_api_stop"),
	resource.TestCheckResourceAttrSet(resourceName, "disable_api_termination"),
funcource.TestCheckResourceAttr(resourceName, "elastic_gpu_specifications.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "iam_instance_profile.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "image_id"),
	resource.TestCheckResourceAttrSet(resourceName, "instance_initiated_shutdown_behavior"),
	resource.TestCheckResourceAttr(resourceName, "instance_market_options.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "instance_type"),
	resource.TestCheckResourceAttrSet(resourceName, "kernel_id"),
	resource.TestCheckResourceAttrSet(resourceName, "key_name"),
	resource.TestCheckResourceAttr(resourceName, "maintenance_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "monitoring.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.security_groups.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.delete_on_termination", ""),
	resource.TestCheckResourceAttr(resourceName, "placement.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "ram_disk_id"),
	resource.TestCheckResourceAttr(resourceName, "tag_specifications.#", "5"),
	resource.TestCheckResourceAttr(resourceName, "vpc_security_group_ids.#", "1"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccEC2LaunchTemplate_description(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_description(rName, "Test Description 1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "description", "Test Description 1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_description(rName, "Test Description 2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "description", "Test Description 2"),
),
	},
},
	})
}


func TestAccEC2LaunchTemplate_update(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
funcourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_asgBasic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "default_version", "1"),
	resource.TestCheckResourceAttr(resourceName, "latest_version", "1"),
funcource.TestCheckResourceAttr(asgResourceName, "launch_template.0.version", "1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_asgUpdate(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttr(resourceName, "latest_version", "2"),
	resource.TestCheckResourceAttr(asgResourceName, "launch_template.#", "1"),
	resource.TestCheckResourceAttr(asgResourceName, "launch_template.0.version", "2"),
),
	},
},
	})
}


func := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
func
Config: testAccLaunchTemplateConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func TestAccEC2LaunchTemplate_CapacityReservation_preference(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_capacityReservationPreference(rName, "open"),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "capacity_reservation_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "capacity_reservation_specification.0.capacity_reservation_preference", "open"),
	resource.TestCheckResourceAttr(resourceName, "capacity_reservation_specification.0.capacity_reservation_target.#", "0"),
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

func TestAccEC2LaunchTemplate_CapacityReservation_target(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_capacityReservationTarget(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "capacity_reservation_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "capacity_reservation_specification.0.capacity_reservation_target.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "capacity_reservation_specification.0.capacity_reservation_preference", ""),
	resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_specification.0.capacity_reservation_target.0.capacity_reservation_id"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2LaunchTemplate_cpuOptions(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcginalCoreCount := 2
	updatedCoreCount := 3
	originalThreadsPerCore := 2
	updatedThreadsPerCore := 1

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_cpuOptions(rName, ec2.AmdSevSnpSpecificationEnabled, originalCoreCount, originalThreadsPerCore),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resName, &template),
funcource.TestCheckResourceAttr(resName, "cpu_options.0.core_count", strconv.Itoa(originalCoreCount)),
	resource.TestCheckResourceAttr(resName, "cpu_options.0.threads_per_core", strconv.Itoa(originalThreadsPerCore)),
),
	},
	{
ResourceName:
ImportState:
ImportStateVerify: true,
func
Config: testAccLaunchTemplateConfig_cpuOptions(rName, ec2.AmdSevSnpSpecificationDisabled, updatedCoreCount, updatedThreadsPerCore),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resName, &template),
	resource.TestCheckResourceAttr(resName, "cpu_options.0.amd_sev_snp", ec2.AmdSevSnpSpecificationDisabled),
	resource.TestCheckResourceAttr(resName, "cpu_options.0.core_count", strconv.Itoa(updatedCoreCount)),
	resource.TestCheckResourceAttr(resName, "cpu_options.0.threads_per_core", strconv.Itoa(updatedThreadsPerCore)),
func
},
	})
}


func TestAccEC2LaunchTemplate_CreditSpecification_nonBurstable(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccLaunchTemplateConfig_creditSpecification(rName, "m1.small", "standard"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"credit_specification"},
	},
},
func


func TestAccEC2LaunchTemplate_CreditSpecification_t2(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_creditSpecification(rName, "t2.micro", "unlimited"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "credit_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "credit_specification.0.cpu_credits", "unlimited"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_CreditSpecification_t3(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_creditSpecification(rName, "t3.micro", "unlimited"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "credit_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "credit_specification.0.cpu_credits", "unlimited"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_CreditSpecification_t4g(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccLaunchTemplateConfig_creditSpecification(rName, "t4g.micro", "unlimited"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "credit_specification.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "credit_specification.0.cpu_credits", "unlimited"),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/6757

func TestAccEC2LaunchTemplate_IAMInstanceProfile_emptyBlock(t *testing.T) {
	ctx := acctest.Context(t)
	var template1 ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_iamInstanceProfileEmptyConfigurationBlock(rName),
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template1),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}
func
func TestAccEC2LaunchTemplate_networkInterface(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_networkInterface(rName),
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_carrier_ip_address", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_public_ip_address", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.delete_on_termination", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.description", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.device_index", "0"),
funcource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_address_count", "2"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_addresses.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_prefix_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_prefixes.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv6_address_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv6_addresses.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv6_prefix_count", "0"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv6_prefixes.#", "0"),
funcource.TestCheckResourceAttrSet(resourceName, "network_interfaces.0.network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.private_ip_address", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.security_groups.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.subnet_id", ""),
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


func TestAccEC2LaunchTemplate_networkInterfaceAddresses(t *testing.T) {
func template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_networkInterfaceAddresses(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_public_ip_address", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_addresses.#", "2"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2LaunchTemplate_networkInterfaceType(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.interface_type", "efa"),
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


func TestAccEC2LaunchTemplate_networkInterfaceCardIndex(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_networkInterfaceCardIndex(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttrSet(resourceName, "instance_type"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.network_card_index", "1"),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccEC2LaunchTemplate_networkInterfaceIPv4PrefixCount(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_networkInterfaceIPv4PrefixCount(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_prefix_count", "1"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}
func
func TestAccEC2LaunchTemplate_networkInterfaceIPv4Prefixes(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_networkInterfaceIPv4Prefixes(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_prefixes.#", "2"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_networkInterfaceIPv6PrefixCount(t *testing.T) {
func template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_networkInterfaceIPv6PrefixCount(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv6_prefix_count", "2"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2LaunchTemplate_networkInterfaceIPv6Prefixes(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv6_prefixes.#", "1"),
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

func TestAccEC2LaunchTemplate_associatePublicIPAddress(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_associatePublicIPAddress(rName, "true"),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interfaces.0.network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_public_ip_address", "true"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_address_count", "2"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_associatePublicIPAddress(rName, "false"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttrSet(resourceName, "network_interfaces.0.network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_public_ip_address", "false"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_address_count", "2"),
),
	},
	{
Config: testAccLaunchTemplateConfig_associatePublicIPAddress(rName, "null"),
Check: resource.ComposeTestCheck
functAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interfaces.0.network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_public_ip_address", ""),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_address_count", "2"),
),
	},
},
func


func TestAccEC2LaunchTemplate_associateCarrierIPAddress(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccLaunchTemplateConfig_associateCarrierIPAddress(rName, "true"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interfaces.0.network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_carrier_ip_address", "true"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_address_count", "2"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_associateCarrierIPAddress(rName, "false"),
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interfaces.0.network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_carrier_ip_address", "false"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv4_address_count", "2"),
),
	},
	{
Config: testAccLaunchTemplateConfig_associateCarrierIPAddress(rName, "null"),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interfaces.0.network_interface_id"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.associate_carrier_ip_address", ""),
func
	},
},
	})
}


func TestAccEC2LaunchTemplate_Placement_hostResourceGroupARN(t *testing.T) {
func template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_placementHostResourceGroupARN(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttrPair(resourceName, "placement.0.host_resource_group_arn", "aws_resourcegroups_group.test", "arn"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2LaunchTemplate_Placement_partitionNum(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_partition(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "placement.0.partition_number", "1"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_partition(rName, 2),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
func
	},
},
	})
}


func TestAccEC2LaunchTemplate_privateDNSNameOptions(t *testing.T) {
func template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_privateDNSNameOptions(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttr(resourceName, "private_dns_name_options.0.enable_resource_name_dns_aaaa_record", "true"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name_options.0.enable_resource_name_dns_a_record", "false"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_name_options.0.hostname_type", "resource-name"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2LaunchTemplate_NetworkInterface_ipv6Addresses(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_networkInterfaceIPv6Addresses(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2LaunchTemplate_NetworkInterface_ipv6AddressCount(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
func
Config: testAccLaunchTemplateConfig_ipv6Count(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "network_interfaces.0.ipv6_address_count", "1"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}
func
func TestAccEC2LaunchTemplate_instanceMarketOptions(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	asgResourceName := "aws_autoscaling_group.test"
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceMarketOptionsBasic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_market_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_market_options.0.spot_options.#", "1"),
	resource.TestCheckResourceAttr(asgResourceName, "launch_template.#", "1"),
	resource.TestCheckResourceAttr(asgResourceName, "launch_template.0.version", "1"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceMarketOptionsUpdate(rName),
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_market_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_market_options.0.spot_options.#", "1"),
	resource.TestCheckResourceAttr(asgResourceName, "launch_template.#", "1"),
	resource.TestCheckResourceAttr(asgResourceName, "launch_template.0.version", "2"),
),
	},
func
}


func TestAccEC2LaunchTemplate_instanceRequirements_memoryMiBAndVCPUCount(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_mib.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_mib.0.min", "500"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.vcpu_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.vcpu_count.0.min", "1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_mib {
00
000

t {



Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_mib.0.min", "500"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_mib.0.max", "4000"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.vcpu_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.vcpu_count.0.min", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.vcpu_count.0.max", "8"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_count {


func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_count.0.min", "1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
funcig: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_count {



b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_count.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_count.0.max", "4"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_count {


func

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_count.0.max", "0"),
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

func TestAccEC2LaunchTemplate_instanceRequirements_acceleratorManufacturers(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_manufacturers = ["amd"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
funcource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_manufacturers.*", "amd"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_manufacturers = ["amazon-web-services", "amd", "nvidia", "xilinx"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
funcource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_manufacturers.*", "amazon-web-services"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_manufacturers.*", "amd"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_manufacturers.*", "nvidia"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_manufacturers.*", "xilinx"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_instanceRequirements_acceleratorNames(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_names = ["a100"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_names.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_names.*", "a100"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
funccelerator_names = ["a100", "v100", "k80", "t4", "m60", "radeon-pro-v520", "vu9p"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_names.#", "7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_names.*", "a100"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_names.*", "v100"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_names.*", "k80"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_names.*", "t4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_names.*", "m60"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_names.*", "radeon-pro-v520"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_names.*", "vu9p"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}


func TestAccEC2LaunchTemplate_instanceRequirements_acceleratorTotalMemoryMiB(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_total_memory_mib {
000

b {
00

func

Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_total_memory_mib.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_total_memory_mib.0.min", "1000"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
funcig: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_total_memory_mib {
4000

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_total_memory_mib.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_total_memory_mib.0.max", "24000"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`accelerator_total_memory_mib {
000
4000

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_total_memory_mib.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_total_memory_mib.0.min", "1000"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_total_memory_mib.0.max", "24000"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccEC2LaunchTemplate_instanceRequirements_acceleratorTypes(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
funccelerator_types = ["fpga"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_types.*", "fpga"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
func
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.accelerator_types.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_types.*", "fpga"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_types.*", "gpu"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.accelerator_types.*", "inference"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccEC2LaunchTemplate_instanceRequirements_allowedInstanceTypes(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`allowed_instance_types = ["m4.large"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.allowed_instance_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.allowed_instance_types.*", "m4.large"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`allowed_instance_types = ["m4.large", "m5.*", "m6*"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.allowed_instance_types.#", "3"),
funcource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.allowed_instance_types.*", "m5.*"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.allowed_instance_types.*", "m6*"),
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


func TestAccEC2LaunchTemplate_instanceRequirements_bareMetal(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
func
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`bare_metal = "excluded"
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.bare_metal", "excluded"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`bare_metal = "included"
b {
00
func


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.bare_metal", "included"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
funcig: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`bare_metal = "required"
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.bare_metal", "required"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccEC2LaunchTemplate_instanceRequirements_baselineEBSBandwidthMbps(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`baseline_ebs_bandwidth_mbps {
0

func

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.baseline_ebs_bandwidth_mbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.baseline_ebs_bandwidth_mbps.0.min", "10"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`baseline_ebs_bandwidth_mbps {
0000

func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.baseline_ebs_bandwidth_mbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.baseline_ebs_bandwidth_mbps.0.max", "20000"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`baseline_ebs_bandwidth_mbps {
0
func
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.baseline_ebs_bandwidth_mbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.baseline_ebs_bandwidth_mbps.0.min", "10"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.baseline_ebs_bandwidth_mbps.0.max", "20000"),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccEC2LaunchTemplate_instanceRequirements_burstablePerformance(t *testing.T) {
func template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`burstable_performance = "excluded"
b {
00
func


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.burstable_performance", "excluded"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`burstable_performance = "included"
b {
000

func

Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.burstable_performance", "included"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`burstable_performance = "required"
b {
000

t {
func
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.burstable_performance", "required"),
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
func
func TestAccEC2LaunchTemplate_instanceRequirements_cpuManufacturers(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`cpu_manufacturers = ["amd"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.cpu_manufacturers.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.cpu_manufacturers.*", "amd"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`cpu_manufacturers = ["amazon-web-services", "amd", "intel"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.cpu_manufacturers.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.cpu_manufacturers.*", "amazon-web-services"),
funcource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.cpu_manufacturers.*", "intel"),
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


func TestAccEC2LaunchTemplate_instanceRequirements_excludedInstanceTypes(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`excluded_instance_types = ["t2.nano"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.excluded_instance_types.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.excluded_instance_types.*", "t2.nano"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`excluded_instance_types = ["t2.nano", "t3*", "t4g.*"]
b {
00
func


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.excluded_instance_types.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.excluded_instance_types.*", "t2.nano"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.excluded_instance_types.*", "t3*"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.excluded_instance_types.*", "t4g.*"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_instanceRequirements_instanceGenerations(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
func
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.instance_generations.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.instance_generations.*", "current"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
func
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.instance_generations.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.instance_generations.*", "current"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.instance_generations.*", "previous"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}

func TestAccEC2LaunchTemplate_instanceRequirements_localStorage(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage = "excluded"
func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.local_storage", "excluded"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage = "included"
b {
00
func


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.local_storage", "included"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage = "required"
b {
func
t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.local_storage", "required"),
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


func TestAccEC2LaunchTemplate_instanceRequirements_localStorageTypes(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage_types = ["hdd"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
funcource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.local_storage_types.*", "hdd"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`local_storage_types = ["hdd", "ssd"]
b {
00

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.local_storage_types.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.local_storage_types.*", "hdd"),
	resource.TestCheckTypeSetElemAttr(resourceName, "instance_requirements.0.local_storage_types.*", "ssd"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_instanceRequirements_memoryGiBPerVCPU(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_gib_per_vcpu {
.5

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_gib_per_vcpu.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_gib_per_vcpu.0.min", "0.5"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_gib_per_vcpu {
.5

b {
00

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_gib_per_vcpu.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_gib_per_vcpu.0.max", "9.5"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`memory_gib_per_vcpu {
.5
.5

b {
00

t {


Check: resource.ComposeTestCheck
functAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_gib_per_vcpu.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_gib_per_vcpu.0.min", "0.5"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.memory_gib_per_vcpu.0.max", "9.5"),
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


func TestAccEC2LaunchTemplate_instanceRequirements_networkBandwidthGbps(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_bandwidth_gbps {
.5

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_bandwidth_gbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_bandwidth_gbps.0.min", "1.5"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_bandwidth_gbps {
00

b {
00
func


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_bandwidth_gbps.0.max", "200"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_bandwidth_gbps {
.5
50

b {
func
t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_bandwidth_gbps.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_bandwidth_gbps.0.min", "2.5"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_bandwidth_gbps.0.max", "250"),
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

func TestAccEC2LaunchTemplate_instanceRequirements_networkInterfaceCount(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_interface_count {


b {
func
t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_interface_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_interface_count.0.min", "1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_interface_count {
0

b {
00

func

Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_interface_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_interface_count.0.max", "10"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`network_interface_count {

0

b {
00

func

Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_interface_count.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_interface_count.0.min", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.network_interface_count.0.max", "10"),
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


func TestAccEC2LaunchTemplate_instanceRequirements_onDemandMaxPricePercentageOverLowestPrice(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`on_demand_max_price_percentage_over_lowest_price = 50
b {
00

t {


Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.on_demand_max_price_percentage_over_lowest_price", "50"),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccEC2LaunchTemplate_instanceRequirements_requireHibernateSupport(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
func
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`require_hibernate_support = false
b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.require_hibernate_support", "false"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`require_hibernate_support = true
func

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.require_hibernate_support", "true"),
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


func TestAccEC2LaunchTemplate_instanceRequirements_spotMaxPricePercentageOverLowestPrice(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`spot_max_price_percentage_over_lowest_price = 75
b {
00

t {


Check: resource.ComposeTestCheck
functAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.spot_max_price_percentage_over_lowest_price", "75"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccEC2LaunchTemplate_instanceRequirements_totalLocalStorageGB(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`total_local_storage_gb {
.5

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.total_local_storage_gb.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.total_local_storage_gb.0.min", "0.5"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
Config: testAccLaunchTemplateConfig_instanceRequirements(sdkacctest.RandomWithPrefix(acctest.ResourcePrefix),
	`total_local_storage_gb {
0.5

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.total_local_storage_gb.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.total_local_storage_gb.0.max", "20.5"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
functal_local_storage_gb {
.5
0.5

b {
00

t {


Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.total_local_storage_gb.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.total_local_storage_gb.0.min", "0.5"),
	resource.TestCheckResourceAttr(resourceName, "instance_requirements.0.total_local_storage_gb.0.max", "20.5"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccEC2LaunchTemplate_licenseSpecification(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckIAMServiceLinkedRole(ctx, t, "/aws-service-role/license-manager.amazonaws.com")
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_licenseSpecification(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "license_specification.#", "1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2LaunchTemplate_metadataOptions(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_metadataOptions(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_endpoint", "enabled"),
funcource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_put_response_hop_limit", "2"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_protocol_ipv6", ""),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.instance_metadata_tags", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_metadataOptionsIPv6(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_endpoint", "enabled"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_tokens", "required"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_put_response_hop_limit", "2"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_protocol_ipv6", "enabled"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.instance_metadata_tags", ""),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_metadataOptionsInstanceTags(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_endpoint", "enabled"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_tokens", "required"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_put_response_hop_limit", "2"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_protocol_ipv6", "enabled"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
funcig: testAccLaunchTemplateConfig_metadataOptionsNoHTTPEndpoint(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_endpoint", "enabled"), //Setting any of the values in metadata options will set the http_endpoint to enabled, you will not see it via the Console, but will in the API for any instance made from the template
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_tokens", "required"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_put_response_hop_limit", "2"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.http_protocol_ipv6", "enabled"),
	resource.TestCheckResourceAttr(resourceName, "metadata_options.0.instance_metadata_tags", "enabled"),
),
	},
},
	})
}
func
func TestAccEC2LaunchTemplate_enclaveOptions(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_enclaveOptions(rName, true),
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "enclave_options.0.enabled", "true"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_enclaveOptions(rName, false),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "enclave_options.0.enabled", "false"),
),
	},
	{
Config: testAccLaunchTemplateConfig_enclaveOptions(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "enclave_options.0.enabled", "true"),
),
func
	})
}


func TestAccEC2LaunchTemplate_hibernation(t *testing.T) {
	ctx := acctest.Context(t)
	var template ec2.LaunchTemplate
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_hibernation(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "hibernation_options.0.configured", "true"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccLaunchTemplateConfig_hibernation(rName, false),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "hibernation_options.0.configured", "false"),
),
	},
	{
Config: testAccLaunchTemplateConfig_hibernation(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckLaunchTemplateExists(ctx, resourceName, &template),
	resource.TestCheckResourceAttr(resourceName, "hibernation_options.0.configured", "true"),
),
	},
},
	})
}
func
func TestAccEC2LaunchTemplate_defaultVersion(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_launch_template.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	description := "Test Description 1"
	descriptionNew := "Test Description 2"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_description(rName, description),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr(resourceName, "default_version", "1"),
	resource.TestCheckResourceAttr(resourceName, "latest_version", "1"),
),
	},
	// An updated config should cause a new version to be created
	// but keep the default_version unchanged if unset
	{
Config: testAccLaunchTemplateConfig_description(rName, descriptionNew),
func(
	resource.TestCheckResourceAttr(resourceName, "default_version", "1"),
	resource.TestCheckResourceAttr(resourceName, "latest_version", "2"),
),
	},
	// An updated config to set the default_version to an existing version
	// should not cause a new version to be created
	{
Config: testAccLaunchTemplateConfig_descriptionDefaultVersion(rName, descriptionNew, 2),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "latest_version", "2"),
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


func TestAccEC2LaunchTemplate_updateDefaultVersion(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	description := "Test Description 1"
	descriptionNew := "Test Description 2"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateConfig_description(rName, description),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "default_version", "1"),
	resource.TestCheckResourceAttr(resourceName, "latest_version", "1"),
func
	// Updating a field should create a new version but not update the default_version
	// if update_default_version is set to false
	{
Config: testAccLaunchTemplateConfig_configDescriptionUpdateDefaultVersion(rName, descriptionNew, false),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "default_version", "1"),
	resource.TestCheckResourceAttr(resourceName, "latest_version", "2"),
),
	},
	// Only updating the update_default_version to true should not create a new version
	// but update the template version to the latest available
	{
Config: testAccLaunchTemplateConfig_configDescriptionUpdateDefaultVersion(rName, descriptionNew, true),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "default_version", "2"),
func
	},
	// Updating a field should create a new version and update the default_version
	// if update_default_version is set to true
	{
Config: testAccLaunchTemplateConfig_configDescriptionUpdateDefaultVersion(rName, description, true),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "default_version", "3"),
	resource.TestCheckResourceAttr(resourceName, "latest_version", "3"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"update_default_version",
func
},
	})
}


func testAccCheckLaunchTemplateExists(ctx context.Context, n string, v *ec2.LaunchTemplate) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Launch Template ID is set")
}
func := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindLaunchTemplateByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_launch_template" {
continue
func
	_, err := tfec2.FindLaunchTemplateByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
func

	return fmt.Errorf("EC2 Launch Template %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccLaunchTemplateConfig_name(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
func
`, rName)
}


func testAccLaunchTemplateConfig_nameGenerated() string {
	return `
resource "aws_launch_template" "test" {}
func


func testAccLaunchTemplateConfig_namePrefix(namePrefix string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name_prefix = %[1]q
}
`, namePrefix)
}
func
func testAccLaunchTemplateConfig_ipv6Count(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
v6_address_count = 1
func
`, rName)
}


func testAccLaunchTemplateConfig_blockDeviceMappingsEBS(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
funcest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  name = %[1]q

  block_device_mappings {
vice_name = "/dev/xvda"

s {
ze = 15

func

# Creating an AutoScaling Group verifies the launch template
# ValidationError: You must use a valid fully-formed launch template. the encrypted flag cannot be specified since device /dev/sda1 has a snapshot specified.
resource "aws_autoscaling_group" "test" {
  availability_zones = [data.aws_availability_zones.available.names[0]]
  desired_capacity
  max_size  = 0
funcme= %[1]q

  launch_template {
nch_template.test.id
rsion = aws_launch_template.test.default_version
  }
}
`, rName))
}

func testAccLaunchTemplateConfig_blockDeviceMappingsEBSDeleteOnTermination(rName string, deleteOnTermination bool) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
func
  block_device_mappings {
vice_name = "/dev/xvda"

s {
_termination = %[2]t
ze  = 15

func

# Creating an AutoScaling Group verifies the launch template
# ValidationError: You must use a valid fully-formed launch template. the encrypted flag cannot be specified since device /dev/sda1 has a snapshot specified.
resource "aws_autoscaling_group" "test" {
  availability_zones = [data.aws_availability_zones.available.names[0]]
  desired_capacity
  max_size  = 0
  min_size  = 0
  name= %[1]q
funcunch_template {
nch_template.test.id
rsion = aws_launch_template.test.default_version
  }
}
`, rName, deleteOnTermination))
}


func testAccLaunchTemplateConfig_blockDeviceMappingsEBSGP3(rName string) string {
funcest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  name = %[1]q

  block_device_mappings {
vice_name = "/dev/xvda"

s {

t  = 500
func "gp3"

  }
}

# Creating an AutoScaling Group verifies the launch template
# ValidationError: You must use a valid fully-formed launch template. the encrypted flag cannot be specified since device /dev/sda1 has a snapshot specified.
resource "aws_autoscaling_group" "test" {
  availability_zones = [data.aws_availability_zones.available.names[0]]
funcx_size  = 0
  min_size  = 0
  name= %[1]q

  launch_template {
nch_template.test.id
rsion = aws_launch_template.test.default_version
  }
funcName))
}


func testAccLaunchTemplateConfig_networkInterfacesDeleteOnTermination(rName string, deleteOnTermination string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
functy_groupsa23bc45"]
lete_on_termination = %[2]s
  }
}
`, rName, deleteOnTermination)
}


func testAccLaunchTemplateConfig_ebsOptimized(rName, ebsOptimized string) string {
	return fmt.Sprintf(`
funcs_optimized = %[1]s
  name = %[2]q
}
`, ebsOptimized, rName)
}


func testAccLaunchTemplateConfig_elasticInferenceAccelerator(rName, elasticInferenceAcceleratorType string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
func
  elastic_inference_accelerator {
pe = %[2]q
  }
}
`, rName, elasticInferenceAcceleratorType)
}


func testAccLaunchTemplateConfig_data(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %q

  block_device_mappings {
vice_name = "test"
  }

funcecovery = "disabled"
func
  disable_api_stop
funcs_optimized  = false

  elastic_gpu_specifications {
pe = "test"
  }

  iam_instance_profile {
me = "test"
  }

  image_id3b456"
  instance_initiated_shutdown_behavior = "terminate"

  instance_market_options {
rket_type = "spot"
  }

  instance_type = "t2.micro"
  kernel_idbc3de"
  key_name

  monitoring {
abled = true
  }

funck_interface_id = "eni-123456ab"
func

funcbility_zone = data.aws_availability_zones.available.names[0]
  }

  ram_disk_idari-a12bc3de"
  vpc_security_group_ids = ["sg-12a3b45c"]

  tag_specifications {
source_type = "instance"

gs = {
est"

  }

  tag_specifications {
source_type = "volume"

gs = {
est"

  }

  tag_specifications {
source_type = "spot-instances-request"

gs = {
func
  }

  tag_specifications {
source_type = "elastic-gpu"

gs = {
est"

func
  tag_specifications {
source_type = "network-interface"

gs = {
est"

func
`, rName)) //lintignore:AWSAT002
}


func testAccLaunchTemplateConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q
funcgs = {
2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}


func testAccLaunchTemplateConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

func= %[3]q
4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}


func testAccLaunchTemplateConfig_capacityReservationPreference(rName string, preference string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  capacity_reservation_specification {
pacity_reservation_preference = %[2]q
  }
}
`, rName, preference)
}


func testAccLaunchTemplateConfig_capacityReservationTarget(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_capacity_reservation" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  instance_count1
  instance_platform = "Linux/UNIX"
  instance_typeo"

  tags = {
me = %[1]q
  }
}

resource "aws_launch_template" "test" {
  name = %[1]q

  capacity_reservation_specification {
funcrvation_id = aws_ec2_capacity_reservation.test.id

  }
}
`, rName))
}


func testAccLaunchTemplateConfig_cpuOptions(rName, amdSevSnp string, coreCount, threadsPerCore int) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  cpu_options {
d_sev_snp
re_count
reads_per_core = %[4]d
  }
}
`, rName, amdSevSnp, coreCount, threadsPerCore)
}


func testAccLaunchTemplateConfig_creditSpecification(rName, instanceType, cpuCredits string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  instance_type = %[1]q
  name = %[2]q

  credit_specification {
u_credits = %[3]q
  }
}
`, instanceType, rName, cpuCredits)
}


func testAccLaunchTemplateConfig_iamInstanceProfileEmptyConfigurationBlock(rName string) string {
	return fmt.Sprintf(`
funcme = %[1]q

  iam_instance_profile {}
}
`, rName)
}


func testAccLaunchTemplateConfig_licenseSpecification(rName string) string {
	return fmt.Sprintf(`
resource "aws_licensemanager_license_configuration" "test" {
  nameTest"
  license_counting_type = "vCPU"
}

resource "aws_launch_template" "test" {
  name = %[1]q

  license_specification {
cense_configuration_arn = aws_licensemanager_license_configuration.test.id
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_description(rName, description string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name
  description = %[2]q
}
`, rName, description)
}


func testAccLaunchTemplateConfig_networkInterface(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

func %[1]q
  }
}

resource "aws_subnet" "test" {
  vpc_idtest.id
  cidr_block = "10.1.0.0/24"

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id
funcgs = {
me = %[1]q
  }
}

resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
twork_interface_id = aws_network_interface.test.id
func
}
`, rName)
}


func testAccLaunchTemplateConfig_partition(rName string, partNum int) string {
	return fmt.Sprintf(`
resource "aws_placement_group" "test" {
  name
  strategy = "partition"
}

funcme = %[1]q

  placement {
oup_nameacement_group.test.name
rtition_number = %[2]d
  }
}
`, rName, partNum)
}


func testAccLaunchTemplateConfig_placementHostResourceGroupARN(rName string) string {
	return fmt.Sprintf(`
resource "aws_resourcegroups_group" "test" {
  name = %[1]q

  resource_query {
ery = jsonencode({
ypeFilters = ["AWS::EC2::Instance"]
s = [

 Key"Stage"
 Values = ["Test"]



  }
}

resource "aws_launch_template" "test" {
  name = %[1]q

  placement {
st_resource_group_arn = aws_resourcegroups_group.test.arn
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_privateDNSNameOptions(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  private_dns_name_options {
able_resource_name_dns_aaaa_record = true
able_resource_name_dns_a_record = se
stname_type= "resource-name"
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_networkInterfaceAddresses(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  vpc_idtest.id
  cidr_block = "10.1.0.0/24"

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

resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
twork_interface_id = aws_network_interface.test.id
v4_addresses.0.10", "10.1.0.11"]
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_associatePublicIPAddress(rName, associatePublicIPAddress string) string {
	return fmt.Sprintf(`
funcdr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  vpc_idtest.id
  cidr_block = "10.1.0.0/24"

  tags = {
me = %[1]q
func

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
funcate_public_ip_address = %[2]s
v4_address_count = 2
  }
}
`, rName, associatePublicIPAddress)
}


func testAccLaunchTemplateConfig_associateCarrierIPAddress(rName, associateCarrierIPAddress string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

func %[1]q
  }
}

resource "aws_subnet" "test" {
  vpc_idtest.id
  cidr_block = "10.1.0.0/24"

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

resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
twork_interface_id= aws_network_interface.test.id
funcddress_count  = 2
  }
}
`, rName, associateCarrierIPAddress)
}


func testAccLaunchTemplateConfig_networkInterfaceIPv6Addresses(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
v6_addresses = [
0:ffff:a01:5",
func
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_networkInterfaceTypeEFA(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
terface_type = "efa"
func
`, rName)
}


func testAccLaunchTemplateConfig_networkInterfaceCardIndex(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q
  instance_type = "p4d.24xlarge"

funck_card_index = 1
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_networkInterfaceIPv4PrefixCount(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
v4_prefix_count = 1
  }
}
`, rName)
}
func
func testAccLaunchTemplateConfig_networkInterfaceIPv4Prefixes(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
v4_prefixes = ["172.16.10.16/28", "172.16.10.32/28"]
  }
}
func


func testAccLaunchTemplateConfig_networkInterfaceIPv6PrefixCount(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
v6_prefix_count = 2
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_networkInterfaceIPv6Prefixes(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  network_interfaces {
v6_prefixes = ["2001:db8::/80"]
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_asgBasic(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  name = %[1]q
func
resource "aws_autoscaling_group" "test" {
  availability_zones = [data.aws_availability_zones.available.names[0]]
  desired_capacity
  max_size  = 0
  min_size  = 0
  name= %[1]q

  launch_template {
nch_template.test.id
rsion = aws_launch_template.test.latest_version
  }
}
`, rName))
}


func testAccLaunchTemplateConfig_asgUpdate(rName string) string {
	return acctest.ConfigCompose(
funcest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForRegion("t3.nano", "t2.nano"),
fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  name = %[1]q
}

resource "aws_autoscaling_group" "test" {
  availability_zones = [data.aws_availability_zones.available.names[0]]
  desired_capacity
  max_size  = 0
  min_size  = 0
  name= %[1]q

  launch_template {
nch_template.test.id
rsion = aws_launch_template.test.latest_version
  }
}
`, rName))
}


func testAccLaunchTemplateConfig_instanceMarketOptionsBasic(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
funcSprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  name = %[1]q

  instance_market_options {
rket_type = "spot"

ot_options {
ance_type = "one-time"

  }
}

funcailability_zones = [data.aws_availability_zones.available.names[0]]
  desired_capacity
  min_size  = 0
  max_size  = 0
  name= %[1]q

  launch_template {
nch_template.test.id
rsion = aws_launch_template.test.latest_version
  }
}
`, rName))
}


func testAccLaunchTemplateConfig_instanceMarketOptionsUpdate(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  name = %[1]q

  instance_market_options {
rket_type = "spot"

ot_options {
 = "0.5"
ance_type = "one-time"

  }
}

resource "aws_autoscaling_group" "test" {
  availability_zones = [data.aws_availability_zones.available.names[0]]
  desired_capacity
funcx_size  = 0
  name= %[1]q

  launch_template {
nch_template.test.id
rsion = aws_launch_template.test.latest_version
  }
}
`, rName))
}


func testAccLaunchTemplateConfig_instanceRequirements(rName, instanceRequirements string) string {
	return fmt.Sprintf(`
data "aws_ami" "test" {
  most_recent = true
  ownersn"]

  filter {
me= "e"
lues = ["amzn-ami-hvm-*-x86_64-gp2"]
  }
}

resource "aws_launch_template" "test" {
  name
  image_id = data.aws_ami.test.id

  instance_requirements {
2]s
  }
}
`, rName, instanceRequirements)
}


func testAccLaunchTemplateConfig_metadataOptions(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q
functadata_options {
tp_endpoint= "enabled"
tp_tokens  = "required"
tp_put_response_hop_limit = 2
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_metadataOptionsIPv6(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  metadata_options {
tp_endpoint= "enabled"
tp_tokens  = "required"
tp_put_response_hop_limit = 2
tp_protocol_ipv6 = "enabled"
  }
}
`, rName)
}


func testAccLaunchTemplateConfig_metadataOptionsInstanceTags(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  metadata_options {
tp_endpoint= "enabled"
tp_tokens  = "required"
tp_put_response_hop_limit = 2
tp_protocol_ipv6 = "enabled"
stance_metadata_tagsd"
  }
}
`, rName)
func
func testAccLaunchTemplateConfig_metadataOptionsNoHTTPEndpoint(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  metadata_options {
tp_tokens  = "required"
tp_put_response_hop_limit = 2
stance_metadata_tagsd"
  }
}
`, rName)
}

func testAccLaunchTemplateConfig_enclaveOptions(rName string, enabled bool) string {
funcurce "aws_launch_template" "test" {
  name = %[1]q

  enclave_options {
abled = %[2]t
  }
}
`, rName, enabled)
}


func testAccLaunchTemplateConfig_hibernation(rName string, enabled bool) string {
	return fmt.Sprintf(`
funcme = %[1]q

  hibernation_options {
nfigured = %[2]t
  }
}
`, rName, enabled)
}


func testAccLaunchTemplateConfig_descriptionDefaultVersion(rName, description string, version int) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name[1]q
funcfault_version = %[3]d
}
`, rName, description, version)
}


func testAccLaunchTemplateConfig_configDescriptionUpdateDefaultVersion(rName, description string, update bool) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name%[1]q
  description[2]q
  update_default_version = %[3]t
}
func
funcfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfunc