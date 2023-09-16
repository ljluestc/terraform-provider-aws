// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker_test

import (
	"context"
	"fmt"
	"os"
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
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "domain_name", rName),
					resource.TestCheckResourceAttr(resourceName, "auth_mode", "IAM"),
					resource.TestCheckResourceAttr(resourceName, "app_network_access_type", "PublicInternetOnly"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_space_settings.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.execution_role", "aws_iam_role.test", "arn"),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`domain/.+`)),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttrPair(resourceName, "vpc_id", "aws_vpc.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "home_efs_file_system_id"),
					resource.TestCheckResourceAttr(resourceName, "domain_settings.#", "0"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_domainSettings(t *testing.T) {
func domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_domainSettings(rName, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "domain_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domain_settings.0.execution_role_identity_config", "DISABLED"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
			{
				Config: testAccDomainConfig_domainSettings(rName, "USER_PROFILE_NAME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "domain_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domain_settings.0.execution_role_identity_config", "USER_PROFILE_NAME"),
				),
			},
		},
	})
}

func testAccDomain_kms(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_kms(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttrPair(resourceName, "kms_key_id", "aws_kms_key.test", "key_id"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
funcourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
			{
				Config: testAccDomainConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccDomainConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func testAccDomain_securityGroup(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_securityGroup1(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.security_groups.#", "1"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
			{
				Config: testAccDomainConfig_securityGroup2(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.security_groups.#", "2"),
				),
			},
		},
	})
}

func testAccDomain_sharingSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"
funcource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_sharingSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.sharing_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.sharing_settings.0.notebook_output_option", "Allowed"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.sharing_settings.0.s3_kms_key_id", "aws_kms_key.test", "key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "default_user_settings.0.sharing_settings.0.s3_output_path"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_canvasAppSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

funceCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_canvasAppSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.canvas_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.canvas_app_settings.0.time_series_forecasting_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.canvas_app_settings.0.time_series_forecasting_settings.0.status", "DISABLED"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_modelRegisterSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
funcrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_modelRegisterSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.canvas_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.canvas_app_settings.0.model_register_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.canvas_app_settings.0.model_register_settings.0.status", "DISABLED"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_workspaceSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
funcotoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_workspaceSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.canvas_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.canvas_app_settings.0.workspace_settings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "default_user_settings.0.canvas_app_settings.0.workspace_settings.0.s3_artifact_path"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_tensorboardAppSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
funceckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_tensorBoardAppSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.tensor_board_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.tensor_board_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.tensor_board_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.micro"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_tensorboardAppSettingsWithImage(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funceps: []resource.TestStep{
			{
				Config: testAccDomainConfig_tensorBoardAppSettingsImage(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.tensor_board_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.tensor_board_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.tensor_board_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.micro"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.tensor_board_app_settings.0.default_resource_spec.0.sagemaker_image_arn", "aws_sagemaker_image.test", "arn"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_rSessionAppSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
func
				Config: testAccDomainConfig_rSessionAppSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.r_session_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.r_session_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.r_session_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.micro"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_rStudioServerProAppSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
funcConfig: testAccDomainConfig_rStudioServerProAppSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.r_studio_server_pro_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.r_studio_server_pro_app_settings.0.access_status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.r_studio_server_pro_app_settings.0.user_group", "R_STUDIO_ADMIN"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_kernelGatewayAppSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
funcCheck: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.micro"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_kernelGatewayAppSettings_lifecycleConfig(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_kernelGatewayAppSettingsLifecycle(rName),
func	testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.lifecycle_config_arns.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.micro"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.0.lifecycle_config_arn", "aws_sagemaker_studio_lifecycle_config.test", "arn"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_kernelGatewayAppSettings_customImage(t *testing.T) {
	ctx := acctest.Context(t)
	if os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE") == "" {
		t.Skip("Environment variable SAGEMAKER_IMAGE_VERSION_BASE_IMAGE is not set")
	}

	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"
	baseImage := os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE")

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funceps: []resource.TestStep{
			{
				Config: testAccDomainConfig_kernelGatewayAppSettingsCustomImage(rName, baseImage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.custom_image.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.custom_image.0.app_image_config_name", "aws_sagemaker_app_image_config.test", "app_image_config_name"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.custom_image.0.image_name", "aws_sagemaker_image.test", "image_name"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_kernelGatewayAppSettings_defaultResourceSpecAndCustomImage(t *testing.T) {
	ctx := acctest.Context(t)
	if os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE") == "" {
		t.Skip("Environment variable SAGEMAKER_IMAGE_VERSION_BASE_IMAGE is not set")
	}

	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"
	baseImage := os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE")

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
func
				Config: testAccDomainConfig_kernelGatewayAppSettingsDefaultResourceSpecAndCustomImage(rName, baseImage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.custom_image.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.0.sagemaker_image_version_arn", "aws_sagemaker_image_version.test", "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.custom_image.0.app_image_config_name", "aws_sagemaker_app_image_config.test", "app_image_config_name"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.custom_image.0.image_name", "aws_sagemaker_image.test", "image_name"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_jupyterServerAppSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_jupyterServerAppSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
func	resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.jupyter_server_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.jupyter_server_app_settings.0.default_resource_spec.0.instance_type", "system"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_jupyterServerAppSettings_code(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_jupyterServerAppSettingsCode(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.jupyter_server_app_settings.#", "1"),
func	resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.jupyter_server_app_settings.0.default_resource_spec.0.instance_type", "system"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.jupyter_server_app_settings.0.code_repository.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "default_user_settings.0.jupyter_server_app_settings.0.code_repository.*", map[string]string{
						"repository_url": "https://github.com/hashicorp/terraform-provider-aws.git",
					}),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceDomain(), resourceName),
				),
				ExpectNonEmptyPlan: true,
func
	})
}

func testAccDomain_defaultUserSettingsUpdated(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "domain_name", rName),
					resource.TestCheckResourceAttr(resourceName, "auth_mode", "IAM"),
					resource.TestCheckResourceAttr(resourceName, "app_network_access_type", "PublicInternetOnly"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
func	resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.execution_role", "aws_iam_role.test", "arn"),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`domain/.+`)),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttrPair(resourceName, "vpc_id", "aws_vpc.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "home_efs_file_system_id"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
			{
				Config: testAccDomainConfig_sharingSettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.sharing_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.sharing_settings.0.notebook_output_option", "Allowed"),
					resource.TestCheckResourceAttrPair(resourceName, "default_user_settings.0.sharing_settings.0.s3_kms_key_id", "aws_kms_key.test", "key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "default_user_settings.0.sharing_settings.0.s3_output_path"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
		},
	})
}

func testAccDomain_spaceSettingsKernelGatewayAppSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var domain sagemaker.DescribeDomainOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDomainDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDomainConfig_defaultSpaceKernelGatewayAppSettings(rName, "ml.t3.micro"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.micro"),
					resource.TestCheckResourceAttr(resourceName, "default_space_settings.#", "1"),
func	resource.TestCheckResourceAttr(resourceName, "default_space_settings.0.kernel_gateway_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_space_settings.0.kernel_gateway_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.micro"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"retention_policy"},
			},
			{
				Config: testAccDomainConfig_defaultSpaceKernelGatewayAppSettings(rName, "ml.t3.small"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDomainExists(ctx, resourceName, &domain),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_user_settings.0.kernel_gateway_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.small"),
					resource.TestCheckResourceAttr(resourceName, "default_space_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_space_settings.0.kernel_gateway_app_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_space_settings.0.kernel_gateway_app_settings.0.default_resource_spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_space_settings.0.kernel_gateway_app_settings.0.default_resource_spec.0.instance_type", "ml.t3.small"),
				),
			},
		},
	})
}

func testAccCheckDomainDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_sagemaker_domain" {
				continue
			}

			domain, err := tfsagemaker.FindDomainByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return fmt.Errorf("reading SageMaker Domain (%s): %w", rs.Primary.ID, err)
			}

			domainArn := aws.StringValue(domain.DomainArn)
			domainID, err := tfsagemaker.DecodeDomainID(domainArn)
			if err != nil {
func
funcf domainID == rs.Primary.ID {
				return fmt.Errorf("sagemaker domain %q still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccCheckDomainExists(ctx context.Context, n string, codeRepo *sagemaker.DescribeDomainOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sagmaker domain ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)
		resp, err := tfsagemaker.FindDomainByName(ctx, conn, rs.Primary.ID)
		if err != nil {
			return err
		}

		*codeRepo = *resp

		return nil
	}
}

func testAccDomainConfig_base(rName string) string {
func "aws_partition" "current" {}
funcurce "aws_iam_role" "test" {
name
path
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

resource "aws_iam_role_policy_attachment" "test" {
rolews_iam_role.test.name
policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/AmazonSageMakerFullAccess"
}
`, rName))
func
func testAccDomainConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn
}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_domainSettings(rName, config string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

funcion_role = aws_iam_role.test.arn
}

domain_settings {
ecution_role_identity_config = %[2]q
}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName, config))
}

func testAccDomainConfig_kms(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_kms_key" "test" {
description
deletion_window_in_days = 7
func
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id
kms_key_id= aws_kms_key.test.key_id

default_user_settings {
ecution_role = aws_iam_role.test.arn
}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_securityGroup1(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_security_group" "test" {
count = 1
funcme = "%[1]s-${count.index}"

tags = {
me = %[1]q
}
}

resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role= aws_iam_role.test.arn
curity_groups = aws_security_group.test[*].id
}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

funcurn acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_security_group" "test" {
count = 2

name = "%[1]s-${count.index}"

tags = {
me = %[1]q
}
}

resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role= aws_iam_role.test.arn
curity_groups = aws_security_group.test[*].id
}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_tags1(rName, tagKey1, tagValue1 string) string {
funcurce "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn
}

retention_policy {
me_efs_file_system = "Delete"
}

tags = {
2]q = %[3]q
}
}
`, rName, tagKey1, tagValue1))
}

func testAccDomainConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
func

retention_policy {
me_efs_file_system = "Delete"
}

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2))
}

func testAccDomainConfig_sharingSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_kms_key" "test" {
description
deletion_window_in_days = 7
}


resource "aws_s3_bucket" "test" {
funcrce_destroy = true
}

resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

aring_settings {
book_output_option = "Allowed"
ms_key_id= awsest.key_id
utput_path/${aws_s3_bucket.test.bucket}/sharing"

}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
func
func testAccDomainConfig_canvasAppSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

nvas_app_settings {
_series_forecasting_settings {
atus = "DISABLED"


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_modelRegisterSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn
func_app_settings {
l_register_settings {
atus = "DISABLED"


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_workspaceSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket%[1]q
force_destroy = true
}

resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
func
default_user_settings {
ecution_role = aws_iam_role.test.arn

nvas_app_settings {
space_settings {
_artifact_path = "s3://${aws_s3_bucket.test.bucket}/path"


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_tensorBoardAppSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id
funcfault_user_settings {
ecution_role = aws_iam_role.test.arn

nsor_board_app_settings {
ult_resource_spec {
stance_type = "ml.t3.micro"


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_tensorBoardAppSettingsImage(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_image" "test" {
image_name = %[1]q
role_arnws_iam_role.test.arn
}

resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

funcion_role = aws_iam_role.test.arn

nsor_board_app_settings {
ult_resource_spec {
stance_type= ".micro"
gemaker_image_arn = aws_sagemaker_image.test.arn


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_jupyterServerAppSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

funcion_role = aws_iam_role.test.arn

pyter_server_app_settings {
ult_resource_spec {
stance_type = "system"


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_jupyterServerAppSettingsCode(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

pyter_server_app_settings {
_repository {
pository_url = "https://github.com/hashicorp/terraform-provider-aws.git"

funcresource_spec {
stance_type = "system"


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_rSessionAppSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

session_app_settings {
funcce_type = "ml.t3.micro"


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_rStudioServerProAppSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

studio_server_pro_app_settings {
ss_status = "ENABLED"
_group"R_STUDIO_ADMIN"

}

funcfs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_kernelGatewayAppSettings(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

rnel_gateway_app_settings {
ult_resource_spec {
stance_type = "ml.t3.micro"


}

retention_policy {
func
}
`, rName))
}

func testAccDomainConfig_defaultSpaceKernelGatewayAppSettings(rName, instance string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

rnel_gateway_app_settings {
ult_resource_spec {
stance_type = %[2]q


}

default_space_settings {
func
rnel_gateway_app_settings {
ult_resource_spec {
stance_type = %[2]q


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName, instance))
}

func testAccDomainConfig_kernelGatewayAppSettingsLifecycle(rName string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_studio_lifecycle_config" "test" {
studio_lifecycle_config_name]q
studio_lifecycle_config_app_type = "JupyterServer"
studio_lifecycle_config_content= base64encode("echo Hello")
}

resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
funcc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

rnel_gateway_app_settings {
ult_resource_spec {
stance_type.micro"
fecycle_config_arn = aws_sagemaker_studio_lifecycle_config.test.arn


cycle_config_arns = [aws_sagemaker_studio_lifecycle_config.test.arn]

}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName))
}

func testAccDomainConfig_kernelGatewayAppSettingsCustomImage(rName, baseImage string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_image" "test" {
image_name = %[1]q
role_arnws_iam_role.test.arn

depends_on = [aws_iam_role_policy_attachment.test]
}

resource "aws_sagemaker_app_image_config" "test" {
app_image_config_name = %[1]q

func_spec {
 = %[1]q

}
}

resource "aws_sagemaker_image_version" "test" {
image_name = aws_sagemaker_image.test.id
base_image = %[2]q
}

resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

rnel_gateway_app_settings {
om_image {
p_image_config_name = aws_sagemaker_app_image_config.test.app_image_config_name
age_name= aer_image_version.test.image_name


}

retention_policy {
me_efs_file_system = "Delete"
}
}
`, rName, baseImage))
}
func testAccDomainConfig_kernelGatewayAppSettingsDefaultResourceSpecAndCustomImage(rName, baseImage string) string {
	return acctest.ConfigCompose(testAccDomainConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_image" "test" {
image_name = %[1]q
role_arnws_iam_role.test.arn

depends_on = [aws_iam_role_policy_attachment.test]
}

resource "aws_sagemaker_app_image_config" "test" {
app_image_config_name = %[1]q

kernel_gateway_image_config {
rnel_spec {
 = %[1]q

}
}

resource "aws_sagemaker_image_version" "test" {
image_name = aws_sagemaker_image.test.id
base_image = %[2]q
}

resource "aws_sagemaker_domain" "test" {
domain_name = %[1]q
auth_modeIAM"
vpc_ids_vpc.test.id
subnet_ids= aws_subnet.test[*].id

default_user_settings {
ecution_role = aws_iam_role.test.arn

rnel_gateway_app_settings {
ult_resource_spec {
stance_type= "ml
gemaker_image_version_arn = aws_sagemaker_image_version.test.arn


om_image {
p_image_config_name = aws_sagemaker_app_image_config.test.app_image_config_name
age_name= aer_image_version.test.image_name


}

retention_policy {
me_efs_file_system = "Delete"
func
`, rName, baseImage))
}
