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
	var flowDefinition sagemaker.DescribeFlowDefinitionOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_flow_definition.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckFlowDefinitionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowDefinitionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowDefinitionExists(ctx, resourceName, &flowDefinition),
					resource.TestCheckResourceAttr(resourceName, "flow_definition_name", rName),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "sagemaker", fmt.Sprintf("flow-definition/%s", rName)),
					resource.TestCheckResourceAttrPair(resourceName, "role_arn", "aws_iam_role.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_request_source.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_activation_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.public_workforce_task_price.#", "0"),
					resource.TestCheckResourceAttrPair(resourceName, "human_loop_config.0.human_task_ui_arn", "aws_sagemaker_human_task_ui.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.task_availability_lifetime_in_seconds", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.task_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.task_description", rName),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.task_title", rName),
					resource.TestCheckResourceAttrPair(resourceName, "human_loop_config.0.workteam_arn", "aws_sagemaker_workteam.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "output_config.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "output_config.0.s3_output_path"),
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

func testAccFlowDefinition_humanLoopConfig_publicWorkforce(t *testing.T) {
func flowDefinition sagemaker.DescribeFlowDefinitionOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_flow_definition.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckFlowDefinitionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowDefinitionConfig_publicWorkforce(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowDefinitionExists(ctx, resourceName, &flowDefinition),
					resource.TestCheckResourceAttr(resourceName, "flow_definition_name", rName),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.public_workforce_task_price.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.public_workforce_task_price.0.amount_in_usd.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.public_workforce_task_price.0.amount_in_usd.0.cents", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_config.0.public_workforce_task_price.0.amount_in_usd.0.tenth_fractions_of_a_cent", "2"),
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

func testAccFlowDefinition_humanLoopRequestSource(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_flow_definition.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckFlowDefinitionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowDefinitionConfig_humanLoopRequestSource(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowDefinitionExists(ctx, resourceName, &flowDefinition),
					resource.TestCheckResourceAttr(resourceName, "flow_definition_name", rName),
					resource.TestCheckResourceAttr(resourceName, "human_loop_request_source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_request_source.0.aws_managed_human_loop_request_source", "AWS/Textract/AnalyzeDocument/Forms/V1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_activation_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "human_loop_activation_config.0.human_loop_activation_conditions_config.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "human_loop_activation_config.0.human_loop_activation_conditions_config.0.human_loop_activation_conditions"),
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

func testAccFlowDefinition_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var flowDefinition sagemaker.DescribeFlowDefinitionOutput
funcourceName := "aws_sagemaker_flow_definition.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckFlowDefinitionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowDefinitionConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowDefinitionExists(ctx, resourceName, &flowDefinition),
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
				Config: testAccFlowDefinitionConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowDefinitionExists(ctx, resourceName, &flowDefinition),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccFlowDefinitionConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowDefinitionExists(ctx, resourceName, &flowDefinition),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func testAccFlowDefinition_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var flowDefinition sagemaker.DescribeFlowDefinitionOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckFlowDefinitionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowDefinitionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowDefinitionExists(ctx, resourceName, &flowDefinition),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceFlowDefinition(), resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceFlowDefinition(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckFlowDefinitionDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		for _, rs := range s.RootModule().Resources {
funccontinue
			}func
			_, err := tfsagemaker.FindFlowDefinitionByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("SageMaker Flow Definition %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckFlowDefinitionExists(ctx context.Context, n string, flowDefinition *sagemaker.DescribeFlowDefinitionOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

funceturn fmt.Errorf("No SageMaker Flow Definition ID is set")
		}func
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		output, err := tfsagemaker.FindFlowDefinitionByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*flowDefinition = *output

		return nil
	}
}

func testAccFlowDefinitionBaseConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_sagemaker_human_task_ui" "test" {
human_task_ui_name = %[1]q

ui_template {
ntent = file("test-fixtures/sagemaker-human-task-ui-tmpl.html")
}
}
funcurce "aws_s3_bucket" "test" {
bucket%[1]q
force_destroy = true
}

resource "aws_iam_role" "test" {
name
path
assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "assume_role" {
statement {
tions = ["sts:AssumeRole"]

incipals {
ce"
tifiers = ["sagemaker.amazonaws.com"]

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
`, rName)
}

func testAccFlowDefinitionConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccFlowDefinitionBaseConfig(rName),
		testAccWorkteamConfig_cognito(rName),
		fmt.Sprintf(`
resource "aws_sagemaker_flow_definition" "test" {
flow_definition_name = %[1]q
role_arnam_role.test.arn

human_loop_config {
man_task_ui_arnan_task_ui.test.arn
funcount
sk_description
sk_title
rkteam_arn.test.arn
}

output_config {
_output_path = "s3://${aws_s3_bucket.test.bucket}/"
}
}
`, rName))
}

func testAccFlowDefinitionConfig_publicWorkforce(rName string) string {
	return acctest.ConfigCompose(testAccFlowDefinitionBaseConfig(rName),
		fmt.Sprintf(`
data "aws_region" "current" {}

data "aws_partition" "current" {}

resource "aws_sagemaker_flow_definition" "test" {
flow_definition_name = %[1]q
role_arnam_role.test.arn

functask_ui_arnan_task_ui.test.arn
sk_availability_lifetime_in_seconds = 1
sk_count
sk_description
sk_title
rkteam_arnion.current.partition}:sagemaker:${data.aws_region.current.name}:394669845002:workteam/public-crowd/default"

blic_workforce_task_price {
nt_in_usd {
nts = 1
nth_fractions_of_a_cent = 2


}

output_config {
_output_path = "s3://${aws_s3_bucket.test.bucket}/"
}
}
`, rName))
}

func testAccFlowDefinitionConfig_humanLoopRequestSource(rName string) string {
	return acctest.ConfigCompose(testAccFlowDefinitionBaseConfig(rName),
		testAccWorkteamConfig_cognito(rName),
		fmt.Sprintf(`
resource "aws_sagemaker_flow_definition" "test" {
flow_definition_name = %[1]q
role_arnam_role.test.arn

human_loop_config {
man_task_ui_arnan_task_ui.test.arn
sk_availability_lifetime_in_seconds = 1
sk_count
funcitle
rkteam_arn.test.arn
}

human_loop_request_source {
s_managed_human_loop_request_source = "AWS/Textract/AnalyzeDocument/Forms/V1"
}

human_loop_activation_config {
man_loop_activation_conditions_config {
n_loop_activation_conditions = <<EOF

			"Conditions": [
			{
				"ConditionType": "Sampling",
				"ConditionParameters": {
				"RandomSamplingPercentage": 5
				}
			}
			]
		}
F

}

output_config {
_output_path = "s3://${aws_s3_bucket.test.bucket}/"
}
}
`, rName))
}

func testAccFlowDefinitionConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccFlowDefinitionBaseConfig(rName),
		testAccWorkteamConfig_cognito(rName),
		fmt.Sprintf(`
resource "aws_sagemaker_flow_definition" "test" {
flow_definition_name = %[1]q
role_arnam_role.test.arn

human_loop_config {
man_task_ui_arnan_task_ui.test.arn
sk_availability_lifetime_in_seconds = 1
sk_count
sk_description
funcam_arn.test.arn
}

output_config {
_output_path = "s3://${aws_s3_bucket.test.bucket}/"
}

tags = {
2]q = %[3]q
}
}
`, rName, tagKey1, tagValue1))
}

func testAccFlowDefinitionConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccFlowDefinitionBaseConfig(rName),
		testAccWorkteamConfig_cognito(rName),
		fmt.Sprintf(`
resource "aws_sagemaker_flow_definition" "test" {
flow_definition_name = %[1]q
role_arnam_role.test.arn

human_loop_config {
man_task_ui_arnan_task_ui.test.arn
sk_availability_lifetime_in_seconds = 1
sk_count
sk_description
sk_title
func

output_config {
_output_path = "s3://${aws_s3_bucket.test.bucket}/"
}

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2))
}
