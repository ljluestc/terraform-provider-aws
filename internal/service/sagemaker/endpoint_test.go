// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker_test

import (
	"context"
	"fmt"
	"testing"

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
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "sagemaker", fmt.Sprintf("endpoint/%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "endpoint_config_name", rName),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
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

func TestAccSageMakerEndpoint_endpointName(t *testing.T) {
functesting.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resourceName := "aws_sagemaker_endpoint.test"
	sagemakerEndpointConfigurationResourceName1 := "aws_sagemaker_endpoint_configuration.test"
	sagemakerEndpointConfigurationResourceName2 := "aws_sagemaker_endpoint_configuration.test2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "endpoint_config_name", sagemakerEndpointConfigurationResourceName1, "name"),
				),
			},
			{
				Config: testAccEndpointConfig_nameUpdate(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "endpoint_config_name", sagemakerEndpointConfigurationResourceName2, "name"),
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

func TestAccSageMakerEndpoint_tags(t *testing.T) {
	ctx := acctest.Context(t)
funcSkip("skipping long-running test in short mode")
	}

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_tags(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
				),
			},
			{
				Config: testAccEndpointConfig_tagsUpdate(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.bar", "baz"),
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

func TestAccSageMakerEndpoint_deploymentConfig(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
func

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_deploymentBasic(rName, "ALL_AT_ONCE", 60),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.auto_rollback_configuration.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.termination_wait_in_seconds", "0"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.type", "ALL_AT_ONCE"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.wait_interval_in_seconds", "60"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.canary_size.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.linear_step_size.#", "0"),
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

func TestAccSageMakerEndpoint_deploymentConfig_blueGreen(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
func
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_deploymentBlueGreen(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.auto_rollback_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.auto_rollback_configuration.0.alarms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.termination_wait_in_seconds", "0"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.type", "LINEAR"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.wait_interval_in_seconds", "60"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.canary_size.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.linear_step_size.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.linear_step_size.0.type", "INSTANCE_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.0.traffic_routing_configuration.0.linear_step_size.0.value", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.rolling_update_policy.#", "0"),
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

func TestAccSageMakerEndpoint_deploymentConfig_rolling(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_deploymentRolling(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.auto_rollback_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.auto_rollback_configuration.0.alarms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.blue_green_update_policy.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.rolling_update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.rolling_update_policy.0.wait_interval_in_seconds", "60"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.rolling_update_policy.0.maximum_batch_size.0.type", "CAPACITY_PERCENT"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config.0.rolling_update_policy.0.maximum_batch_size.0.value", "5"),
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

func TestAccSageMakerEndpoint_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
funcrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceEndpoint(), resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceEndpoint(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckEndpointDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_sagemaker_endpoint" {
				continue
func
			_, erfunc
			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("SageMaker Endpoint (%s) still exists", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckEndpointExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no SageMaker Endpoint ID is set")
func
		conn :func err := tfsagemaker.FindEndpointByName(ctx, conn, rs.Primary.ID)

		return err
	}
}

func testAccEndpointConfig_Base(rName string) string {
	return fmt.Sprintf(`
data "aws_iam_policy_document" "access" {
statement {
fect = "Allow"

tions = [
udwatch:PutMetricData",
s:CreateLogStream",
s:PutLogEvents",
s:CreateLogGroup",
funcAuthorizationToken",
:BatchCheckLayerAvailability",
:GetDownloadUrlForLayer",
:BatchGetImage",
GetObject",


sources = ["*"]
}
}

data "aws_partition" "current" {}

data "aws_iam_policy_document" "assume_role" {
statement {
tions = ["sts:AssumeRole"]

incipals {
ce"
tifiers = ["sagemaker.${data.aws_partition.current.dns_suffix}"]

}
}

resource "aws_iam_role" "test" {
name
path
assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy" "test" {
rolews_iam_role.test.name
policy = data.aws_iam_policy_document.access.json
}

resource "aws_s3_bucket" "test" {
bucket = %[1]q
}

resource "aws_s3_object" "test" {
bucket = aws_s3_bucket.test.id
key"model.tar.gz"
source = "test-fixtures/sagemaker-tensorflow-serving-test-model.tar.gz"
}

data "aws_sagemaker_prebuilt_ecr_image" "test" {
repository_name = "sagemaker-tensorflow-serving"
image_tag1.12-cpu"
}

resource "aws_sagemaker_model" "test" {
name
execution_role_arn = aws_iam_role.test.arn

primary_container {
age_sagemaker_prebuilt_ecr_image.test.registry_path
del_data_url = "https://${aws_s3_bucket.test.bucket_regional_domain_name}/${aws_s3_object.test.key}"
}

depends_on = [aws_iam_role_policy.test]
}

resource "aws_sagemaker_endpoint_configuration" "test" {
name = %[1]q

production_variants {
itial_instance_count = 2
itial_variant_weight = 1
stance_typeedium"
del_nameagemaker_model.test.name
riant_namet-1"
}
}
`, rName)
}

func testAccEndpointConfig_basic(rName string) string {
	return testAccEndpointConfig_Base(rName) + fmt.Sprintf(`
resource "aws_sagemaker_endpoint" "test" {
endpoint_config_name = aws_sagemaker_endpoint_configuration.test.name
name
}
`, rName)
}

func testAccEndpointConfig_nameUpdate(rName string) string {
	return testAccEndpointConfig_Base(rName) + fmt.Sprintf(`
resource "aws_sagemaker_endpoint_configuration" "test2" {
func
production_variants {
itial_instance_count = 1
itial_variant_weight = 1
stance_typeedium"
del_nameagemaker_model.test.name
riant_namet-1"
}
}
funcurce "aws_sagemaker_endpoint" "test" {
endpoint_config_name = aws_sagemaker_endpoint_configuration.test2.name
name
}
`, rName)
}

func testAccEndpointConfig_tags(rName string) string {
	return testAccEndpointConfig_Base(rName) + fmt.Sprintf(`
resource "aws_sagemaker_endpoint" "test" {
endpoint_config_name = aws_sagemaker_endpoint_configuration.test.name
name

tags = {
o = "bar"
}
}
`, rName)
}

func testAccEndpointConfig_tagsUpdate(rName string) string {
funcurce "aws_sagemaker_endpoint" "test" {
endpoint_config_name = aws_sagemaker_endpoint_configuration.test.name
name

tags = {
r = "baz"
}
}
`, rName)
}

func testAccEndpointConfig_deploymentBasic(rName, tType string, wait int) string {
	return testAccEndpointConfig_Base(rName) + fmt.Sprintf(`
funcdpoint_config_name = aws_sagemaker_endpoint_configuration.test.name
name

deployment_config {
ue_green_update_policy {
fic_routing_configuration {
pe = %[
it_interval_in_seconds = %[3]d


}
}
`, rName, tType, wait)
func
func testAccEndpointConfig_deploymentBlueGreen(rName string) string {
	return testAccEndpointConfig_Base(rName) + fmt.Sprintf(`
resource "aws_cloudwatch_metric_alarm" "test" {
alarm_name
comparison_operatorGreaterThanOrEqualToThreshold"
evaluation_periods"2"
metric_nameon"
namespace
period
statistic
threshold
alarm_description "This metric monitors ec2 cpu utilization"
insufficient_data_actions = []

dimensions = {
stanceId = "i-abc123"
}
func
resource "aws_sagemaker_endpoint" "test" {
endpoint_config_name = aws_sagemaker_endpoint_configuration.test.name
name

deployment_config {
ue_green_update_policy {
fic_routing_configuration {
pe = "L
it_interval_in_seconds = "60"

near_step_size {
NSTANCE_COUNT"





to_rollback_configuration {
ms {
arm_name = aws_cloudwatch_metric_alarm.test.alarm_name


}
}
`, rName)
}

func testAccEndpointConfig_deploymentRolling(rName string) string {
	return testAccEndpointConfig_Base(rName) + fmt.Sprintf(`
resource "aws_cloudwatch_metric_alarm" "test" {
alarm_name
comparison_operatorGreaterThanOrEqualToThreshold"
evaluation_periods"2"
metric_nameon"
namespace
period
statistic
threshold
alarm_description "This metric monitors ec2 cpu utilization"
insufficient_data_actions = []

dimensions = {
stanceId = "i-abc123"
}
}
funcurce "aws_sagemaker_endpoint" "test" {
endpoint_config_name = aws_sagemaker_endpoint_configuration.test.name
name

deployment_config {
to_rollback_configuration {
ms {
arm_name = aws_cloudwatch_metric_alarm.test.alarm_name



lling_update_policy {
_interval_in_seconds = 60

mum_batch_size {
pe= "CAPACITY_PERCENT"
lue = 5


}
}
`, rName)
}
