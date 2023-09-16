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
	var mpg sagemaker.GetModelPackageGroupPolicyOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_model_package_group_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckModelPackageGroupPolicyDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccModelPackageGroupPolicyConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckModelPackageGroupPolicyExists(ctx, resourceName, &mpg),
					resource.TestCheckResourceAttr(resourceName, "model_package_group_name", rName),
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

func TestAccSageMakerModelPackageGroupPolicy_disappears(t *testing.T) {
func mpg sagemaker.GetModelPackageGroupPolicyOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_model_package_group_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckModelPackageGroupPolicyDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccModelPackageGroupPolicyConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckModelPackageGroupPolicyExists(ctx, resourceName, &mpg),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceModelPackageGroupPolicy(), resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceModelPackageGroupPolicy(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSageMakerModelPackageGroupPolicy_Disappears_modelPackageGroup(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_model_package_group_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckModelPackageGroupPolicyDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccModelPackageGroupPolicyConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckModelPackageGroupPolicyExists(ctx, resourceName, &mpg),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceModelPackageGroup(), "aws_sagemaker_model_package_group.test"),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceModelPackageGroupPolicy(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckModelPackageGroupPolicyDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)
funcr _, rs := range s.RootModule().Resources {
			if rsfunccontinue
			}

			_, err := tfsagemaker.FindModelPackageGroupPolicyByName(ctx, conn, rs.Primary.ID)
			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return fmt.Errorf("reading SageMaker Model Package Group Policy (%s): %w", rs.Primary.ID, err)
			}
		}

		return nil
	}
}

func testAccCheckModelPackageGroupPolicyExists(ctx context.Context, n string, mpg *sagemaker.GetModelPackageGroupPolicyOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
func
		if rs.funceturn fmt.Errorf("No sagmaker Model Package Group ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)
		resp, err := tfsagemaker.FindModelPackageGroupPolicyByName(ctx, conn, rs.Primary.ID)
		if err != nil {
			return err
		}

		*mpg = *resp

		return nil
	}
}

func testAccModelPackageGroupPolicyConfig_basic(rName string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}
data "aws_partition" "current" {}

data "aws_iam_policy_document" "test" {
  statement {
funcs= [gemaker:DescribeModelPackage", "sagemaker:ListModelPackages"]
sources = [aws_sagemaker_model_package_group.test.arn]
incipals {
tifiers = ["arn:${data.aws_partition.current.partition}:iam::${data.aws_caller_identity.current.account_id}:root"]


  }
}

resource "aws_sagemaker_model_package_group" "test" {
  model_package_group_name = %[1]q
}

resource "aws_sagemaker_model_package_group_policy" "test" {
  model_package_group_name = aws_sagemaker_model_package_group.test.model_package_group_name
  resource_policyde(jsondecode(data.aws_iam_policy_document.test.json))
}
`, rName)
}
