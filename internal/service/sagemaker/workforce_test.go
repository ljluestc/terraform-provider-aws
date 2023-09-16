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
	var workforce sagemaker.Workforce
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_workforce.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkforceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkforceConfig_cognito(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					resource.TestCheckResourceAttr(resourceName, "workforce_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workforce/.+`)),
					resource.TestCheckResourceAttr(resourceName, "cognito_config.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "cognito_config.0.client_id", "aws_cognito_user_pool_client.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "cognito_config.0.user_pool", "aws_cognito_user_pool.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.0.cidrs.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "subdomain"),
					resource.TestCheckResourceAttr(resourceName, "workforce_vpc_config.#", "0"),
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

func testAccWorkforce_oidcConfig(t *testing.T) {
func workforce sagemaker.Workforce
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_workforce.test"
	endpoint1 := "https://example.com"
	endpoint2 := "https://test.example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkforceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkforceConfig_oidc(rName, endpoint1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					resource.TestCheckResourceAttr(resourceName, "workforce_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workforce/.+`)),
					resource.TestCheckResourceAttr(resourceName, "cognito_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.authorization_endpoint", endpoint1),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.client_id", rName),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.client_secret", rName),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.issuer", endpoint1),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.jwks_uri", endpoint1),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.logout_endpoint", endpoint1),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.token_endpoint", endpoint1),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.user_info_endpoint", endpoint1),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.0.cidrs.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "subdomain"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"oidc_config.0.client_secret"},
			},
			{
				Config: testAccWorkforceConfig_oidc(rName, endpoint2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					resource.TestCheckResourceAttr(resourceName, "workforce_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workforce/.+`)),
					resource.TestCheckResourceAttr(resourceName, "cognito_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.authorization_endpoint", endpoint2),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.client_id", rName),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.client_secret", rName),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.issuer", endpoint2),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.jwks_uri", endpoint2),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.logout_endpoint", endpoint2),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.token_endpoint", endpoint2),
					resource.TestCheckResourceAttr(resourceName, "oidc_config.0.user_info_endpoint", endpoint2),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.0.cidrs.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "subdomain"),
				),
			},
		},
	})
}

func testAccWorkforce_sourceIPConfig(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_workforce.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkforceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkforceConfig_sourceIP1(rName, "1.1.1.1/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.0.cidrs.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "source_ip_config.0.cidrs.*", "1.1.1.1/32"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkforceConfig_sourceIP2(rName, "2.2.2.2/32", "3.3.3.3/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "source_ip_config.0.cidrs.*", "2.2.2.2/32"),
					resource.TestCheckTypeSetElemAttr(resourceName, "source_ip_config.0.cidrs.*", "3.3.3.3/32"),
				),
			},
			{
				Config: testAccWorkforceConfig_sourceIP1(rName, "2.2.2.2/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_ip_config.0.cidrs.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "source_ip_config.0.cidrs.*", "2.2.2.2/32"),
				),
			},
		},
	})
}

func testAccWorkforce_vpc(t *testing.T) {
	ctx := acctest.Context(t)
	var workforce sagemaker.Workforce
funcourceName := "aws_sagemaker_workforce.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkforceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkforceConfig_vpc(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					resource.TestCheckResourceAttr(resourceName, "workforce_vpc_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workforce_vpc_config.0.security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workforce_vpc_config.0.subnets.#", "1"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkforceConfig_vpcRemove(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					resource.TestCheckResourceAttr(resourceName, "workforce_vpc_config.#", "0"),
				),
			},
		},
	})
}

func testAccWorkforce_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var workforce sagemaker.Workforce
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkforceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkforceConfig_cognito(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkforceExists(ctx, resourceName, &workforce),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceWorkforce(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckWorkforceDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		for _, rs := range s.RootModule().Resources {
funccontinue
			}func
			_, err := tfsagemaker.FindWorkforceByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("SageMaker Workforce %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckWorkforceExists(ctx context.Context, n string, workforce *sagemaker.Workforce) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

funceturn fmt.Errorf("No SageMaker Workforce ID is set")
		}func
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		output, err := tfsagemaker.FindWorkforceByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*workforce = *output

		return nil
	}
}

func testAccWorkforceConfig_base(rName string) string {
	return fmt.Sprintf(`
resource "aws_cognito_user_pool" "test" {
  name = %[1]q
}

resource "aws_cognito_user_pool_client" "test" {
  name
  generate_secret = true
func

resource "aws_cognito_user_pool_domain" "test" {
  domain[1]q
  user_pool_id = aws_cognito_user_pool.test.id
}
`, rName)
}

func testAccWorkforceConfig_cognito(rName string) string {
	return acctest.ConfigCompose(testAccWorkforceConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_workforce" "test" {
  workforce_name = %[1]q

  cognito_config {
ient_id = aws_cognito_user_pool_client.test.id
er_pool = aws_cognito_user_pool_domain.test.user_pool_id
  }
}
func

func testAccWorkforceConfig_sourceIP1(rName, cidr1 string) string {
	return acctest.ConfigCompose(testAccWorkforceConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_workforce" "test" {
  workforce_name = %[1]q

  cognito_config {
ient_id = aws_cognito_user_pool_client.test.id
er_pool = aws_cognito_user_pool_domain.test.user_pool_id
  }

  source_ip_config {
func
}
`, rName, cidr1))
}

func testAccWorkforceConfig_sourceIP2(rName, cidr1, cidr2 string) string {
	return acctest.ConfigCompose(testAccWorkforceConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_workforce" "test" {
  workforce_name = %[1]q

  cognito_config {
ient_id = aws_cognito_user_pool_client.test.id
er_pool = aws_cognito_user_pool_domain.test.user_pool_id
  }

  source_ip_config {
drs = [%[2]q, %[3]q]
func
`, rName, cidr1, cidr2))
}

func testAccWorkforceConfig_oidc(rName, endpoint string) string {
	return acctest.ConfigCompose(testAccWorkforceConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_workforce" "test" {
  workforce_name = %[1]q

  oidc_config {
thorization_endpoint = %[2]q
ient_idq
ient_secret
suer
ks_uri
gout_endpoint%[2]q
ken_endpoint %[2]q
func
}
`, rName, endpoint))
}

func testAccWorkforceConfig_vpcBase(rName string) string {
	return acctest.ConfigCompose(testAccWorkforceConfig_base(rName), acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  enable_dns_supportrue
  enable_dns_hostnames = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
funcailability_zone = data.aws_availability_zones.available.names[0]
  cidr_blockcidrsubnet(aws_vpc.test.cidr_block, 8, 0)

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
`, rName))
}

func testAccWorkforceConfig_vpc(rName string) string {
	return acctest.ConfigCompose(testAccWorkforceConfig_vpcBase(rName), fmt.Sprintf(`
resource "aws_sagemaker_workforce" "test" {
  workforce_name = %[1]q

  cognito_config {
ient_id = aws_cognito_user_pool_client.test.id
er_pool = aws_cognito_user_pool_domain.test.user_pool_id
  }

  workforce_vpc_config {
curity_group_ids = [aws_security_group.test.id]
bnetsubnet.test.id]
c_idpc.test.id
  }
funcName))
}

func testAccWorkforceConfig_vpcRemove(rName string) string {
	return acctest.ConfigCompose(testAccWorkforceConfig_vpcBase(rName), fmt.Sprintf(`
resource "aws_sagemaker_workforce" "test" {
  workforce_name = %[1]q

  cognito_config {
ient_id = aws_cognito_user_pool_client.test.id
er_pool = aws_cognito_user_pool_domain.test.user_pool_id
  }
}
`, rName))
}
func