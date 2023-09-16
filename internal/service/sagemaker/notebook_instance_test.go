// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
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
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "accelerator_types.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "additional_code_repositories.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "default_code_repository", ""),
					resource.TestCheckResourceAttr(resourceName, "direct_internet_access", "Enabled"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_service_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_service_configuration.0.minimum_instance_metadata_service_version", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_type", "ml.t2.medium"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "platform_identifier", "notebook-al1-v1"),
					resource.TestCheckResourceAttrPair(resourceName, "role_arn", "aws_iam_role.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "root_access", "Enabled"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttr(resourceName, "volume_size", "5"),
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

func TestAccSageMakerNotebookInstance_imds(t *testing.T) {
functesting.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_imds(rName, "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_service_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_service_configuration.0.minimum_instance_metadata_service_version", "2"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_imds(rName, "1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_service_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_service_configuration.0.minimum_instance_metadata_service_version", "1"),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_update(t *testing.T) {
	ctx := acctest.Context(t)
funcSkip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "instance_type", "ml.t2.medium"),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_update(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "instance_type", "ml.m4.xlarge"),
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

func TestAccSageMakerNotebookInstance_volumeSize(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
func

	var notebook1, notebook2, notebook3 sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var resourceName = "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook1),
					resource.TestCheckResourceAttr(resourceName, "volume_size", "5"),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_volume(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook2),
					resource.TestCheckResourceAttr(resourceName, "volume_size", "8"),
					testAccCheckNotebookInstanceNotRecreated(&notebook1, &notebook2),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook3),
					resource.TestCheckResourceAttr(resourceName, "volume_size", "5"),
					testAccCheckNotebookInstanceRecreated(&notebook2, &notebook3),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_lifecycleName(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
func
	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance.test"
	sagemakerLifecycleConfigResourceName := "aws_sagemaker_notebook_instance_lifecycle_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_lifecycleName(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttrPair(resourceName, "lifecycle_config_name", sagemakerLifecycleConfigResourceName, "name"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "lifecycle_config_name", ""),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_lifecycleName(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttrPair(resourceName, "lifecycle_config_name", sagemakerLifecycleConfigResourceName, "name"),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_tags(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
func notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
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
				Config: testAccNotebookInstanceConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_kms(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_kms(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttrPair(resourceName, "kms_key_id", "aws_kms_key.test", "id"),
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

func TestAccSageMakerNotebookInstance_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
funcourceName := "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceNotebookInstance(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_Root_access(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_rootAccess(rName, "Disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "root_access", "Disabled"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_rootAccess(rName, "Enabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "root_access", "Enabled"),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_Platform_identifier(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance.test"
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_platformIdentifier(rName, "notebook-al2-v1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "platform_identifier", "notebook-al2-v1"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_platformIdentifier(rName, "notebook-al1-v1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "platform_identifier", "notebook-al1-v1"),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_DirectInternet_access(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance.test"

funceCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_directInternetAccess(rName, "Disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "direct_internet_access", "Disabled"),
					resource.TestCheckResourceAttrPair(resourceName, "subnet_id", "aws_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestMatchResourceAttr(resourceName, "network_interface_id", regexache.MustCompile("eni-.*")),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_directInternetAccess(rName, "Enabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "direct_internet_access", "Enabled"),
					resource.TestCheckResourceAttrPair(resourceName, "subnet_id", "aws_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestMatchResourceAttr(resourceName, "network_interface_id", regexache.MustCompile("eni-.*")),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_DefaultCode_repository(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var resourceName = "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
funcrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_defaultCodeRepository(rName, "https://github.com/hashicorp/terraform-provider-aws.git"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "default_code_repository", "https://github.com/hashicorp/terraform-provider-aws.git"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "default_code_repository", ""),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_defaultCodeRepository(rName, "https://github.com/hashicorp/terraform-provider-aws.git"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "default_code_repository", "https://github.com/hashicorp/terraform-provider-aws.git"),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_AdditionalCode_repositories(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var resourceName = "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
funcotoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_additionalCodeRepository1(rName, "https://github.com/hashicorp/terraform-provider-aws.git"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "additional_code_repositories.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "additional_code_repositories.*", "https://github.com/hashicorp/terraform-provider-aws.git"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "additional_code_repositories.#", "0"),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_additionalCodeRepository2(rName, "https://github.com/hashicorp/terraform-provider-aws.git", "https://github.com/hashicorp/terraform.git"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "additional_code_repositories.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "additional_code_repositories.*", "https://github.com/hashicorp/terraform-provider-aws.git"),
					resource.TestCheckTypeSetElemAttr(resourceName, "additional_code_repositories.*", "https://github.com/hashicorp/terraform.git"),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_additionalCodeRepository1(rName, "https://github.com/hashicorp/terraform-provider-aws.git"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "additional_code_repositories.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "additional_code_repositories.*", "https://github.com/hashicorp/terraform-provider-aws.git"),
				),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_DefaultCodeRepository_sageMakerRepo(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var resourceName = "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
funceckDestroy:CheckNotebookInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_defaultCodeRepositoryRepo(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttrPair(resourceName, "default_code_repository", "aws_sagemaker_code_repository.test", "code_repository_name"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "default_code_repository", ""),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_defaultCodeRepositoryRepo(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttrPair(resourceName, "default_code_repository", "aws_sagemaker_code_repository.test", "code_repository_name")),
			},
		},
	})
}

func TestAccSageMakerNotebookInstance_acceleratorTypes(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var notebook sagemaker.DescribeNotebookInstanceOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var resourceName = "aws_sagemaker_notebook_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funceps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceConfig_acceleratorType(rName, "ml.eia2.medium"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "accelerator_types.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "accelerator_types.*", "ml.eia2.medium"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebookInstanceConfig_acceleratorType(rName, "ml.eia2.large"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "accelerator_types.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "accelerator_types.*", "ml.eia2.large"),
				),
			},
			{
				Config: testAccNotebookInstanceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceExists(ctx, resourceName, &notebook),
					resource.TestCheckResourceAttr(resourceName, "accelerator_types.#", "0"),
				),
			},
		},
	})
}

func testAccCheckNotebookInstanceDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_sagemaker_notebook_instance" {
				continue
			}

			_, err := tfsagemaker.FindNotebookInstanceByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

funcreturn err
			}func
			return fmt.Errorf("SageMaker Notebook Instance %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckNotebookInstanceExists(ctx context.Context, n string, v *sagemaker.DescribeNotebookInstanceOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SageMaker Notebook Instance ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		output, err := tfsagemaker.FindNotebookInstanceByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
func
		*v = *func
		return nil
	}
}

func testAccCheckNotebookInstanceNotRecreated(i, j *sagemaker.DescribeNotebookInstanceOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if !aws.TimeValue(i.CreationTime).Equal(aws.TimeValue(j.CreationTime)) {
			return errors.New("SageMaker Notebook Instance was recreated")
		}

		return nil
	}
}

func testAccCheckNotebookInstanceRecreated(i, j *sagemaker.DescribeNotebookInstanceOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if aws.TimeValue(i.CreationTime).Equal(aws.TimeValue(j.CreationTime)) {
			return errors.New("SageMaker Notebook Instance was not recreated")
		}

		return nil
	}
}
func testAccNotebookInstanceBaseConfig(rName string) string {
	return funcurce "aws_iam_role" "test" {
name
path
assume_role_policy = data.aws_iam_policy_document.test.json
}

data "aws_iam_policy_document" "test" {
statement {
tions = ["sts:AssumeRole"]
funcpals {
ce"funcers = ["sagemaker.amazonaws.com"]

}
}
`, rName)
}

func testAccNotebookInstanceConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
funcme
role_arns_iam_role.test.arn
instance_type = "ml.t2.medium"
}
`, rName))
}

func testAccNotebookInstanceConfig_update(rName string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arns_iam_role.test.arn
instance_type = "ml.m4.xlarge"
}
`, rName))
}

func testAccNotebookInstanceConfig_lifecycleName(rName string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance_lifecycle_configuration" "test" {
name = %[1]q
func
resource "aws_sagemaker_notebook_instance" "test" {
instance_type "ml.t2.medium"
lifecycle_config_name = aws_sagemaker_notebook_instance_lifecycle_configuration.test.name
name
role_arniam_role.test.arn
}
`, rName))
}

funcurn acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arns_iam_role.test.arn
instance_type = "ml.t2.medium"

tags = {
2]q = %[3]q
}
}
func

func testAccNotebookInstanceConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arns_iam_role.test.arn
instance_type = "ml.t2.medium"

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2))
func
func testAccNotebookInstanceConfig_rootAccess(rName string, rootAccess string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arns_iam_role.test.arn
instance_type = "ml.t2.medium"
root_access[2]q
}
`, rName, rootAccess))
}

func testAccNotebookInstanceConfig_platformIdentifier(rName string, platformIdentifier string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
funcme
role_arnm_role.test.arn
instance_typeml.t2.medium"
platform_identifier = %[2]q
}
`, rName, platformIdentifier))
}
func testAccNotebookInstanceConfig_directInternetAccess(rName string, directInternetAccess string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arntest.arn
instance_typeedium"
security_groups[aws_security_group.test.id]
subnet_idsubnet.test.id
func

resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
funcdr_block = "10.0.0.0/24"

tags = {
me = %[1]q
}
}

resource "aws_security_group" "test" {
vpc_id = aws_vpc.test.id

func %[1]q
}
}
`, rName, directInternetAccess))
}

func testAccNotebookInstanceConfig_volume(rName string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arns_iam_role.test.arn
instance_type = "ml.t2.medium"
volume_size
}
`, rName))
}

func testAccNotebookInstanceConfig_defaultCodeRepository(rName string, defaultCodeRepository string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arn.test.arn
instance_typemedium"
default_code_repository = %[2]q
}
`, rName, defaultCodeRepository))
}

func testAccNotebookInstanceConfig_additionalCodeRepository1(rName, repo1 string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arn.arn
instance_typem"
additional_code_repositories = ["%[2]s"]
}
`, rName, repo1))
}
func testAccNotebookInstanceConfig_additionalCodeRepository2(rName, repo1, repo2 string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
role_arn.arn
instance_typem"
additional_code_repositories = ["%[2]s", "%[3]s"]
}
`, rName, repo1, repo2))
}

funcurn acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_code_repository" "test" {
code_repository_name = %[1]q

git_config {
pository_url = "https://github.com/hashicorp/terraform-provider-aws.git"
}
}

resource "aws_sagemaker_notebook_instance" "test" {
name
funcstance_typemedium"
default_code_repository = aws_sagemaker_code_repository.test.code_repository_name
}
`, rName))
}

func testAccNotebookInstanceConfig_kms(rName string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_kms_key" "test" {
description = %[1]q

func
"Version": "2012-10-17",
"Id": "kms-tf-1",
"Statement": [

": "Enable IAM User Permissions",
ect": "Allow",
ncipal": {
WS": "*"

ion": "kms:*",
func
]
}
POLICY
}

resource "aws_sagemaker_notebook_instance" "test" {
name
role_arns_iam_role.test.arn
instance_type = "ml.t2.medium"
kms_key_idaws_kms_key.test.id
}
`, rName))
}

func testAccNotebookInstanceConfig_imds(rName, version string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
name
funcstance_type = "ml.t2.medium"
instance_metadata_service_configuration {
nimum_instance_metadata_service_version = %[2]q
}
}
`, rName, version))
}

func testAccNotebookInstanceConfig_acceleratorType(rName, acceleratorType string) string {
	return acctest.ConfigCompose(testAccNotebookInstanceBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance" "test" {
nameq
role_arnrole.test.arn
instance_type.t2.xlarge"
accelerator_types = [%[2]q]
}
`, rName, acceleratorType))
}
funcfunc