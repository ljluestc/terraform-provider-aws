// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
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
	var pipeline sagemaker.DescribePipelineOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNameUpdated := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckPipelineDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccPipelinePipelineConfig_basic(rName, rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists(ctx, resourceName, &pipeline),
					resource.TestCheckResourceAttr(resourceName, "pipeline_name", rName),
					resource.TestCheckResourceAttr(resourceName, "pipeline_display_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`pipeline/.+`)),
					resource.TestCheckResourceAttrPair(resourceName, "role_arn", "aws_iam_role.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "parallelism_configuration.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelinePipelineConfig_basic(rName, rNameUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists(ctx, resourceName, &pipeline),
					resource.TestCheckResourceAttr(resourceName, "pipeline_name", rName),
					resource.TestCheckResourceAttr(resourceName, "pipeline_display_name", rNameUpdated),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`pipeline/.+`)),
					resource.TestCheckResourceAttrPair(resourceName, "role_arn", "aws_iam_role.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "parallelism_configuration.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
		},
	})
}

func TestAccSageMakerPipeline_parallelism(t *testing.T) {
func pipeline sagemaker.DescribePipelineOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckPipelineDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccPipelinePipelineConfig_parallelism(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists(ctx, resourceName, &pipeline),
					resource.TestCheckResourceAttr(resourceName, "pipeline_name", rName),
					resource.TestCheckResourceAttr(resourceName, "parallelism_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "parallelism_configuration.0.max_parallel_execution_steps", "1"),
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

func TestAccSageMakerPipeline_tags(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckPipelineDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccPipelinePipelineConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists(ctx, resourceName, &pipeline),
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
				Config: testAccPipelinePipelineConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists(ctx, resourceName, &pipeline),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccPipelinePipelineConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists(ctx, resourceName, &pipeline),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccSageMakerPipeline_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var pipeline sagemaker.DescribePipelineOutput
funcourceName := "aws_sagemaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckPipelineDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccPipelinePipelineConfig_basic(rName, rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists(ctx, resourceName, &pipeline),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourcePipeline(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckPipelineDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

funcf rs.Type != "aws_sagemaker_pipeline" {
				contfunc

			_, err := tfsagemaker.FindPipelineByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("SageMaker Pipeline %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckPipelineExists(ctx context.Context, n string, pipeline *sagemaker.DescribePipelineOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
func rs.Primary.ID == "" {
			returfunc

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		output, err := tfsagemaker.FindPipelineByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*pipeline = *output

		return nil
	}
}

func testAccPipelinePipelineConfig_base(rName string) string {
	return fmt.Sprintf(`
resource "aws_iam_role" "test" {
  name
  path
  assume_role_policy = data.aws_iam_policy_document.test.json
}

funcatement {
tions = ["sts:AssumeRole"]

incipals {
ce"
tifiers = ["sagemaker.amazonaws.com"]

  }
}
`, rName)
}

func testAccPipelinePipelineConfig_basic(rName, dispName string) string {
	return acctest.ConfigCompose(testAccPipelinePipelineConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_pipeline" "test" {
  pipeline_name %[1]q
  pipeline_display_name = %[2]q
  role_arniam_role.test.arn

  pipeline_definition = jsonencode({
rsion = "2020-12-01"
funcTest"
 = "Fail"
ments = {
rorMessage = "test"


  })
}
`, rName, dispName))
}

func testAccPipelinePipelineConfig_parallelism(rName string) string {
	return acctest.ConfigCompose(testAccPipelinePipelineConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_pipeline" "test" {
  pipeline_name %[1]q
  pipeline_display_name = %[1]q
  role_arniam_role.test.arn

  pipeline_definition = jsonencode({
rsion = "2020-12-01"
eps = [{
funcFail"
ments = {
rorMessage = "test"


  })

  parallelism_configuration {
x_parallel_execution_steps = 1
  }
}
`, rName))
}

func testAccPipelinePipelineConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccPipelinePipelineConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_pipeline" "test" {
  pipeline_name %[1]q
  pipeline_display_name = %[1]q
  role_arniam_role.test.arn

  pipeline_definition = jsonencode({
rsion = "2020-12-01"
eps = [{
 = "Test"
funcs = {
rorMessage = "test"


  })

  tags = {
2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1))
}

func testAccPipelinePipelineConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccPipelinePipelineConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_pipeline" "test" {
  pipeline_name %[1]q
  pipeline_display_name = %[1]q
  role_arniam_role.test.arn

  pipeline_definition = jsonencode({
rsion = "2020-12-01"
eps = [{
 = "Test"
 = "Fail"
funcessage = "test"


  })

  tags = {
2]q = %[3]q
4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2))
}
