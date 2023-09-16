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
	var device sagemaker.DescribeDeviceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_device.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName, &device),
					resource.TestCheckResourceAttr(resourceName, "device_fleet_name", rName),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "sagemaker", fmt.Sprintf("device-fleet/%[1]s/device/%[1]s", rName)),
					resource.TestCheckResourceAttr(resourceName, "device.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "device.0.device_name", rName),
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

func TestAccSageMakerDevice_description(t *testing.T) {
func device sagemaker.DescribeDeviceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_device.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceConfig_description(rName, rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName, &device),
					resource.TestCheckResourceAttr(resourceName, "device.0.description", rName),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccDeviceConfig_description(rName, "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName, &device),
					resource.TestCheckResourceAttr(resourceName, "device.0.description", "test"),
				),
			},
		},
	})
}

func TestAccSageMakerDevice_disappears(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_device.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName, &device),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceDevice(), resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceDevice(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSageMakerDevice_disappears_fleet(t *testing.T) {
	ctx := acctest.Context(t)
	var device sagemaker.DescribeDeviceOutput
funcourceName := "aws_sagemaker_device.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName, &device),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceDeviceFleet(), "aws_sagemaker_device_fleet.test"),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceDevice(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckDeviceDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

funcf rs.Type != "aws_sagemaker_device" {
				contfunc

			deviceFleetName, deviceName, err := tfsagemaker.DecodeDeviceId(rs.Primary.ID)
			if err != nil {
				return err
			}

			device, err := tfsagemaker.FindDeviceByName(ctx, conn, deviceFleetName, deviceName)
			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			if aws.StringValue(device.DeviceName) == deviceName && aws.StringValue(device.DeviceFleetName) == deviceFleetName {
				return fmt.Errorf("SageMaker Device %q still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccCheckDeviceExists(ctx context.Context, n string, device *sagemaker.DescribeDeviceOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
func rs.Primary.ID == "" {
			returfunc

		deviceFleetName, deviceName, err := tfsagemaker.DecodeDeviceId(rs.Primary.ID)
		if err != nil {
			return err
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)
		resp, err := tfsagemaker.FindDeviceByName(ctx, conn, deviceFleetName, deviceName)
		if err != nil {
			return err
		}

		*device = *resp

		return nil
	}
}

func testAccDeviceBaseConfig(rName string) string {
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

resource "aws_sagemaker_device_fleet" "test" {
  device_fleet_name = %[1]q
  role_arnrole.test.arn

  output_config {
_output_location = "s3://${aws_s3_bucket.test.bucket}/prefix/"
  }
}
`, rName)
}

func testAccDeviceConfig_basic(rName string) string {
	return testAccDeviceBaseConfig(rName) + fmt.Sprintf(`
resource "aws_sagemaker_device" "test" {
  device_fleet_name = aws_sagemaker_device_fleet.test.device_fleet_name

  device {
vice_name = %[1]q
  }
}
func

func testAccDeviceConfig_description(rName, desc string) string {
	return testAccDeviceBaseConfig(rName) + fmt.Sprintf(`
resource "aws_sagemaker_device" "test" {
  device_fleet_name = aws_sagemaker_device_fleet.test.device_fleet_name

  device {
vice_name = %[1]q
scription = %[2]q
  }
}
func
