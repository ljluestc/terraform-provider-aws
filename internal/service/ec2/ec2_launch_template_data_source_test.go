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
	dataSourceName := "data.aws_launch_template.test"
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateDataSourceConfig_name(rName),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrPair(resourceName, "arn", dataSourceName, "arn"),
funcource.TestCheckResourceAttrPair(resourceName, "capacity_reservation_specification.#", dataSourceName, "capacity_reservation_specification.#"),
	resource.TestCheckResourceAttrPair(resourceName, "cpu_options.#", dataSourceName, "cpu_options.#"),
	resource.TestCheckResourceAttrPair(resourceName, "cpu_options.0.amd_sev_snp", dataSourceName, "cpu_options.0.amd_sev_snp"),
	resource.TestCheckResourceAttrPair(resourceName, "cpu_options.0.core_count", dataSourceName, "cpu_options.0.core_count"),
	resource.TestCheckResourceAttrPair(resourceName, "cpu_options.0.threads_per_core", dataSourceName, "cpu_options.0.threads_per_core"),
	resource.TestCheckResourceAttrPair(resourceName, "credit_specification.#", dataSourceName, "credit_specification.#"),
	resource.TestCheckResourceAttrPair(resourceName, "default_version", dataSourceName, "default_version"),
	resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "description"),
	resource.TestCheckResourceAttrPair(resourceName, "disable_api_stop", dataSourceName, "disable_api_stop"),
	resource.TestCheckResourceAttrPair(resourceName, "disable_api_termination", dataSourceName, "disable_api_termination"),
	resource.TestCheckResourceAttrPair(resourceName, "ebs_optimized", dataSourceName, "ebs_optimized"),
	resource.TestCheckResourceAttrPair(resourceName, "elastic_gpu_specifications.#", dataSourceName, "elastic_gpu_specifications.#"),
	resource.TestCheckResourceAttrPair(resourceName, "elastic_inference_accelerator.#", dataSourceName, "elastic_inference_accelerator.#"),
	resource.TestCheckResourceAttrPair(resourceName, "enclave_options.#", dataSourceName, "enclave_options.#"),
	resource.TestCheckResourceAttrPair(resourceName, "hibernation_options.#", dataSourceName, "hibernation_options.#"),
	resource.TestCheckResourceAttrPair(resourceName, "iam_instance_profile.#", dataSourceName, "iam_instance_profile.#"),
	resource.TestCheckResourceAttrPair(resourceName, "image_id", dataSourceName, "image_id"),
	resource.TestCheckResourceAttrPair(resourceName, "instance_initiated_shutdown_behavior", dataSourceName, "instance_initiated_shutdown_behavior"),
	resource.TestCheckResourceAttrPair(resourceName, "instance_market_options.#", dataSourceName, "instance_market_options.#"),
	resource.TestCheckResourceAttrPair(resourceName, "instance_requirements.#", dataSourceName, "instance_requirements.#"),
	resource.TestCheckResourceAttrPair(resourceName, "instance_type", dataSourceName, "instance_type"),
	resource.TestCheckResourceAttrPair(resourceName, "kernel_id", dataSourceName, "kernel_id"),
	resource.TestCheckResourceAttrPair(resourceName, "key_name", dataSourceName, "key_name"),
	resource.TestCheckResourceAttrPair(resourceName, "latest_version", dataSourceName, "latest_version"),
	resource.TestCheckResourceAttrPair(resourceName, "license_specification.#", dataSourceName, "license_specification.#"),
	resource.TestCheckResourceAttrPair(resourceName, "maintenance_options.#", dataSourceName, "maintenance_options.#"),
	resource.TestCheckResourceAttrPair(resourceName, "metadata_options.#", dataSourceName, "metadata_options.#"),
	resource.TestCheckResourceAttrPair(resourceName, "monitoring.#", dataSourceName, "monitoring.#"),
	resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "name"),
	resource.TestCheckResourceAttrPair(resourceName, "network_interfaces.#", dataSourceName, "network_interfaces.#"),
	resource.TestCheckResourceAttrPair(resourceName, "placement.#", dataSourceName, "placement.#"),
	resource.TestCheckResourceAttrPair(resourceName, "private_dns_name_options.#", dataSourceName, "private_dns_name_options.#"),
	resource.TestCheckResourceAttrPair(resourceName, "ram_disk_id", dataSourceName, "ram_disk_id"),
	resource.TestCheckResourceAttrPair(resourceName, "security_group_names.#", dataSourceName, "security_group_names.#"),
	resource.TestCheckResourceAttrPair(resourceName, "tag_specifications.#", dataSourceName, "tag_specifications.#"),
	resource.TestCheckResourceAttrPair(resourceName, "tags.%", dataSourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(resourceName, "user_data", dataSourceName, "user_data"),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_security_group_ids.#", dataSourceName, "vpc_security_group_ids.#"),
),
	},
},
	})
}


func TestAccEC2LaunchTemplateDataSource_id(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
func
Config: testAccLaunchTemplateDataSourceConfig_id(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "name"),
),
	},
func
}


func TestAccEC2LaunchTemplateDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_launch_template.test"
	resourceName := "aws_launch_template.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateDataSourceConfig_filter(rName),
func(
	resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "name"),
),
	},
},
	})
}
func
func TestAccEC2LaunchTemplateDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_launch_template.test"
	resourceName := "aws_launch_template.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckLaunchTemplateDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccLaunchTemplateDataSourceConfig_tags(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "id"),
funcource.TestCheckResourceAttrPair(resourceName, "tags.%", dataSourceName, "tags.%"),
),
	},
},
	})
}


funcurn acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_launch_template" "test" {
  image_ids_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  name = %[1]q

  block_device_mappings {
func
s {

t  = 500
ze = 15
pe = "gp3"

  }

  elastic_inference_accelerator {
pe = "eia1.medium"
  }

  elastic_gpu_specifications {
pe = "test"
  }

  iam_instance_profile {
me = "test"
  }

  instance_initiated_shutdown_behavior = "terminate"

  instance_market_options {
rket_type = "spot"
  }

  maintenance_options {
to_recovery = "disabled"
  }

  disable_api_stop
  disable_api_termination = true
  ebs_optimized  = false

  kernel_id = "aki-a12bc3de"
  key_name  = "test"

  placement {
ailability_zone = data.aws_availability_zones.available.names[0]
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

  tags = {
me = %[1]q
  }

  capacity_reservation_specification {
pacity_reservation_preference = "open"
  }

  cpu_options {
re_count
reads_per_core = 2
  }

  credit_specification {
u_credits = "unlimited"
  }
}

data "aws_launch_template" "test" {
  name = aws_launch_template.test.name
}
`, rName))
}


func testAccLaunchTemplateDataSourceConfig_id(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q
}

data "aws_launch_template" "test" {
  id = aws_launch_template.test.id
}
`, rName)
}


funcurn fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q
}

data "aws_launch_template" "test" {
  filter {
me= "nch-template-name"
lues = [aws_launch_template.test.name]
  }
}
`, rName)
}
func
func testAccLaunchTemplateDataSourceConfig_tags(rName string) string {
	return fmt.Sprintf(`
resource "aws_launch_template" "test" {
  name = %[1]q

  tags = {
me = %[1]q
  }
}

data "aws_launch_template" "test" {
  tags = {
me = aws_launch_template.test.tags["Name"]
  }
}
func
