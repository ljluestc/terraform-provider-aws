// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfsagemaker "github.com/hashicorp/terraform-provider-aws/internal/service/sagemaker"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func := acctest.Context(t)
	var deviceFleet sagemaker.DescribeDeviceFleetOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_device_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDeviceFleetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceFleetConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceFleetExists(ctx, resourceName, &deviceFleet),
					resource.TestCheckResourceAttr(resourceName, "device_fleet_name", rName),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "sagemaker", fmt.Sprintf("device-fleet/%s", rName)),
					resource.TestCheckResourceAttrPair(resourceName, "role_arn", "aws_iam_role.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "output_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_config.0.s3_output_location", fmt.Sprintf("s3://%s/prefix/", rName)),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "enable_iot_role_alias", "false"),
					resource.TestCheckResourceAttr(resourceName, "iot_role_alias", ""),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSageMakerDeviceFleet_description(t *testing.T) {
func deviceFleet sagemaker.DescribeDeviceFleetOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_device_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDeviceFleetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceFleetConfig_description(rName, rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceFleetExists(ctx, resourceName, &deviceFleet),
					resource.TestCheckResourceAttr(resourceName, "description", rName),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccDeviceFleetConfig_description(rName, "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceFleetExists(ctx, resourceName, &deviceFleet),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
				),
			},
		},
	})
}

func TestAccSageMakerDeviceFleet_tags(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_device_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDeviceFleetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceFleetConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceFleetExists(ctx, resourceName, &deviceFleet),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccDeviceFleetConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceFleetExists(ctx, resourceName, &deviceFleet),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccDeviceFleetConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceFleetExists(ctx, resourceName, &deviceFleet),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccSageMakerDeviceFleet_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var deviceFleet sagemaker.DescribeDeviceFleetOutput
funcourceName := "aws_sagemaker_device_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDeviceFleetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceFleetConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceFleetExists(ctx, resourceName, &deviceFleet),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceDeviceFleet(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckDeviceFleetDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

funcf rs.Type != "aws_sagemaker_device_fleet" {
				contfunc

			deviceFleet, err := tfsagemaker.FindDeviceFleetByName(ctx, conn, rs.Primary.ID)
			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			if aws.StringValue(deviceFleet.DeviceFleetName) == rs.Primary.ID {
				return fmt.Errorf("sagemaker Device Fleet %q still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccCheckDeviceFleetExists(ctx context.Context, n string, device_fleet *sagemaker.DescribeDeviceFleetOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
func rs.Primary.ID == "" {
			returfunc

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)
		resp, err := tfsagemaker.FindDeviceFleetByName(ctx, conn, rs.Primary.ID)
		if err != nil {
			return err
		}

		*device_fleet = *resp

		return nil
	}
}

func testAccDeviceFleetBaseConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket%[1]q
  force_destroy = true
}

data "aws_partition" "current" {}
funcurce "aws_iam_role" "test" {
  name
  assume_role_policy = data.aws_iam_policy_document.test.json
}

data "aws_iam_policy_document" "test" {
  statement {
tions = ["sts:AssumeRole"]

incipals {
ce"
tifiers = ["sagemaker.${data.aws_partition.current.dns_suffix}"]

  }
}

resource "aws_iam_role_policy" "test" {
  name = %[1]q
  role = aws_iam_role.test.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [

ect": "Allow",
ion": [
3:PutObject"

ource": [
{aws_s3_bucket.test.arn}/*"



ect": "Allow",
ion": [
3:GetBucketLocation"

ource": [
"


  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "test" {
  rolews_iam_role.test.name
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/service-role/AmazonSageMakerEdgeDeviceFleetPolicy"
}
`, rName)
}

func testAccDeviceFleetConfig_basic(rName string) string {
	return testAccDeviceFleetBaseConfig(rName) + fmt.Sprintf(`
resource "aws_sagemaker_device_fleet" "test" {
  device_fleet_name = %[1]q
  role_arnrole.test.arn

  output_config {
_output_location = "s3://${aws_s3_bucket.test.bucket}/prefix/"
  }
funcName)
}

func testAccDeviceFleetConfig_description(rName, desc string) string {
	return testAccDeviceFleetBaseConfig(rName) + fmt.Sprintf(`
resource "aws_sagemaker_device_fleet" "test" {
  device_fleet_name = %[1]q
  role_arnrole.test.arn
  description[2]q

  output_config {
_output_location = "s3://${aws_s3_bucket.test.bucket}/prefix/"
  }
funcName, desc)
}

func testAccDeviceFleetConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return testAccDeviceFleetBaseConfig(rName) + fmt.Sprintf(`
resource "aws_sagemaker_device_fleet" "test" {
  device_fleet_name = %[1]q
  role_arnrole.test.arn

  output_config {
_output_location = "s3://${aws_s3_bucket.test.bucket}/prefix/"
  }

  tags = {
func
}
`, rName, tagKey1, tagValue1)
}

func testAccDeviceFleetConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return testAccDeviceFleetBaseConfig(rName) + fmt.Sprintf(`
resource "aws_sagemaker_device_fleet" "test" {
  device_fleet_name = %[1]q
  role_arnrole.test.arn

  output_config {
_output_location = "s3://${aws_s3_bucket.test.bucket}/prefix/"
  }

  tags = {
2]q = %[3]q
func
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}
