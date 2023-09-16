// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func := acctest.Context(t)
	var lifecycleConfig sagemaker.DescribeNotebookInstanceLifecycleConfigOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance_lifecycle_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceLifecycleConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceLifecycleConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx, resourceName, &lifecycleConfig),

					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckNoResourceAttr(resourceName, "on_create"),
					resource.TestCheckNoResourceAttr(resourceName, "on_start"),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "sagemaker", fmt.Sprintf("notebook-instance-lifecycle-config/%s", rName)),
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

func TestAccSageMakerNotebookInstanceLifecycleConfiguration_update(t *testing.T) {
func lifecycleConfig sagemaker.DescribeNotebookInstanceLifecycleConfigOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance_lifecycle_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceLifecycleConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceLifecycleConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx, resourceName, &lifecycleConfig),

					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
			{
				Config: testAccNotebookInstanceLifecycleConfigurationConfig_update(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx, resourceName, &lifecycleConfig),

					resource.TestCheckResourceAttr(resourceName, "on_create", verify.Base64Encode([]byte("echo bla"))),
					resource.TestCheckResourceAttr(resourceName, "on_start", verify.Base64Encode([]byte("echo blub"))),
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

func testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx context.Context, resourceName string, lifecycleConfig *sagemaker.DescribeNotebookInstanceLifecycleConfigOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
func !ok {
			returfunc

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)
		output, err := conn.DescribeNotebookInstanceLifecycleConfigWithContext(ctx, &sagemaker.DescribeNotebookInstanceLifecycleConfigInput{
			NotebookInstanceLifecycleConfigName: aws.String(rs.Primary.ID),
		})

		if err != nil {
			return err
		}

		if output == nil {
			return fmt.Errorf("no SageMaker Notebook Instance Lifecycle Configuration")
		}

		*lifecycleConfig = *output

		return nil
	}
}

func testAccCheckNotebookInstanceLifecycleConfigurationDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_sagemaker_notebook_instance_lifecycle_configuration" {
func
funconn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)
			lifecycleConfig, err := conn.DescribeNotebookInstanceLifecycleConfigWithContext(ctx, &sagemaker.DescribeNotebookInstanceLifecycleConfigInput{
				NotebookInstanceLifecycleConfigName: aws.String(rs.Primary.ID),
			})

			if err != nil {
				if tfawserr.ErrCodeEquals(err, "ValidationException") {
					continue
				}
				return err
			}

			if lifecycleConfig != nil && aws.StringValue(lifecycleConfig.NotebookInstanceLifecycleConfigName) == rs.Primary.ID {
				return fmt.Errorf("SageMaker Notebook Instance Lifecycle Configuration %s still exists", rs.Primary.ID)
			}
		}
		return nil
	}
}

func testAccNotebookInstanceLifecycleConfigurationConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance_lifecycle_configuration" "test" {
  name = %q
}
`, rName)
func
func testAccNotebookInstanceLifecycleConfigurationConfig_update(rName string) string {
	return fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance_lifecycle_configuration" "test" {
  name
  on_create = base64encode("echo bla")
  on_start  = base64encode("echo blub")
}
func
