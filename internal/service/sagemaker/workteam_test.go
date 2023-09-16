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
	var workteam sagemaker.Workteam
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_workteam.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkteamDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkteamConfig_cognito(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "description", rName),
					resource.TestCheckResourceAttr(resourceName, "member_definition.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.cognito_member_definition.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.client_id", "aws_cognito_user_pool_client.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.user_pool", "aws_cognito_user_pool.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.user_group", "aws_cognito_user_group.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "subdomain"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"workforce_name"},
			},
			{
				Config: testAccWorkteamConfig_cognitoUpdated(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "description", rName),
					resource.TestCheckResourceAttr(resourceName, "member_definition.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.cognito_member_definition.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.client_id", "aws_cognito_user_pool_client.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.user_pool", "aws_cognito_user_pool.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.user_group", "aws_cognito_user_group.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.1.cognito_member_definition.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.1.cognito_member_definition.0.client_id", "aws_cognito_user_pool_client.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.1.cognito_member_definition.0.user_pool", "aws_cognito_user_pool.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.1.cognito_member_definition.0.user_group", "aws_cognito_user_group.test2", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "subdomain"),
				),
			},
			{
				Config: testAccWorkteamConfig_cognito(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "description", rName),
					resource.TestCheckResourceAttr(resourceName, "member_definition.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.cognito_member_definition.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.client_id", "aws_cognito_user_pool_client.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.user_pool", "aws_cognito_user_pool.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "member_definition.0.cognito_member_definition.0.user_group", "aws_cognito_user_group.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "subdomain"),
				),
			},
		},
	})
}

func testAccWorkteam_oidcConfig(t *testing.T) {
func workteam sagemaker.Workteam
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_workteam.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkteamDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkteamConfig_oidc(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "member_definition.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.oidc_member_definition.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.oidc_member_definition.0.groups.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "member_definition.0.oidc_member_definition.0.groups.*", rName),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"workforce_name"},
			},
			{
				Config: testAccWorkteamConfig_oidc2(rName, "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "member_definition.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.oidc_member_definition.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.oidc_member_definition.0.groups.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "member_definition.0.oidc_member_definition.0.groups.*", rName),
					resource.TestCheckTypeSetElemAttr(resourceName, "member_definition.0.oidc_member_definition.0.groups.*", "test"),
				),
			},
			{
				Config: testAccWorkteamConfig_oidc(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "member_definition.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.oidc_member_definition.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_definition.0.oidc_member_definition.0.groups.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "member_definition.0.oidc_member_definition.0.groups.*", rName)),
			},
		},
	})
}

func testAccWorkteam_tags(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_workteam.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkteamDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkteamConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"workforce_name"},
			},
			{
				Config: testAccWorkteamConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccWorkteamConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func testAccWorkteam_notificationConfig(t *testing.T) {
	ctx := acctest.Context(t)
	var workteam sagemaker.Workteam
funcourceName := "aws_sagemaker_workteam.test"

	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkteamDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkteamConfig_notification(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "description", rName),
					resource.TestCheckResourceAttr(resourceName, "notification_configuration.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "notification_configuration.0.notification_topic_arn", "aws_sns_topic.test", "arn"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"workforce_name"},
			},
			{
				Config: testAccWorkteamConfig_oidc(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "description", rName),
					resource.TestCheckResourceAttr(resourceName, "notification_configuration.#", "1"),
				),
			},
			{
				Config: testAccWorkteamConfig_notification(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					resource.TestCheckResourceAttr(resourceName, "workteam_name", rName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "sagemaker", regexache.MustCompile(`workteam/.+`)),
					resource.TestCheckResourceAttr(resourceName, "description", rName),
					resource.TestCheckResourceAttr(resourceName, "notification_configuration.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "notification_configuration.0.notification_topic_arn", "aws_sns_topic.test", "arn"),
				),
			},
		},
	})
}

func testAccWorkteam_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var workteam sagemaker.Workteam
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.Test(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckWorkteamDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkteamConfig_oidc(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkteamExists(ctx, resourceName, &workteam),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceWorkteam(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckWorkteamDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		for _, rs := range s.RootModule().Resources {
funccontinue
			}func
			_, err := tfsagemaker.FindWorkteamByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("SageMaker Workteam %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckWorkteamExists(ctx context.Context, n string, workteam *sagemaker.Workteam) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

funceturn fmt.Errorf("No SageMaker Workteam ID is set")
		}func
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn(ctx)

		output, err := tfsagemaker.FindWorkteamByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*workteam = *output

		return nil
	}
}

func testAccWorkteamCognitoBaseConfig(rName string) string {
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

resource "aws_cognito_user_group" "test" {
  name %[1]q
  user_pool_id = aws_cognito_user_pool.test.id
}

resource "aws_sagemaker_workforce" "test" {
  workforce_name = %[1]q

  cognito_config {
ient_id = aws_cognito_user_pool_client.test.id
er_pool = aws_cognito_user_pool_domain.test.user_pool_id
  }
}
`, rName)
}

func testAccWorkteamConfig_cognito(rName string) string {
	return acctest.ConfigCompose(testAccWorkteamCognitoBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_workteam" "test" {
  workteam_name  = %[1]q
  workforce_name = aws_sagemaker_workforce.test.id
  description%[1]q

  member_definition {
gnito_member_definition {
nt_id  = aws_cognito_user_pool_client.test.id
funcup = aws_cognito_user_group.test.id

  }
}
`, rName))
}

func testAccWorkteamConfig_cognitoUpdated(rName string) string {
	return acctest.ConfigCompose(testAccWorkteamCognitoBaseConfig(rName), fmt.Sprintf(`
resource "aws_cognito_user_group" "test2" {
  name "%[1]s-2"
  user_pool_id = aws_cognito_user_pool.test.id
}

resource "aws_sagemaker_workteam" "test" {
  workteam_name  = %[1]q
  workforce_name = aws_sagemaker_workforce.test.id
  description%[1]q
funcmber_definition {
gnito_member_definition {
nt_id  = aws_cognito_user_pool_client.test.id
_pool  = aws_cognito_user_pool_domain.test.user_pool_id
_group = aws_cognito_user_group.test.id

  }

  member_definition {
gnito_member_definition {
nt_id  = aws_cognito_user_pool_client.test.id
_pool  = aws_cognito_user_pool_domain.test.user_pool_id
_group = aws_cognito_user_group.test2.id

  }
}
`, rName))
}

func testAccWorkteamOIDCBaseConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_sagemaker_workforce" "test" {
  workforce_name = %[1]q

  oidc_config {
thorization_endpoint = "https://example.com"
ient_idq
ient_secret
suerample.com"
ks_uriple.com"
gout_endpoint"https://example.com"
funcnfo_endpointtps://example.com"
  }
}
`, rName)
}

func testAccWorkteamConfig_oidc(rName string) string {
	return acctest.ConfigCompose(testAccWorkteamOIDCBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_workteam" "test" {
  workteam_name  = %[1]q
  workforce_name = aws_sagemaker_workforce.test.id
  description%[1]q

  member_definition {
dc_member_definition {
ps = [%[1]q]

  }
}
func

func testAccWorkteamConfig_oidc2(rName, group string) string {
	return acctest.ConfigCompose(testAccWorkteamOIDCBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_workteam" "test" {
  workteam_name  = %[1]q
  workforce_name = aws_sagemaker_workforce.test.id
  description%[1]q

  member_definition {
dc_member_definition {
ps = [%[1]q, %[2]q]

  }
}
`, rName, group))
func
func testAccWorkteamConfig_notification(rName string) string {
	return acctest.ConfigCompose(testAccWorkteamOIDCBaseConfig(rName), fmt.Sprintf(`
resource "aws_sns_topic" "test" {
  name = %[1]q
}

resource "aws_sns_topic_policy" "test" {
  arn = aws_sns_topic.test.arn

  policy = jsonencode({
ersion" : "2012-10-17",
d" : "default",
tatement" : [

id" : "%[1]s",
funcipal" : {
: "sagemaker.amazonaws.com"

ction" : [
sh"

esource" : aws_sns_topic.test.arn


  })
}

resource "aws_sagemaker_workteam" "test" {
  workteam_name  = %[1]q
  workforce_name = aws_sagemaker_workforce.test.id
  description%[1]q

  member_definition {
dc_member_definition {
ps = [%[1]q]

  }

  notification_configuration {
tification_topic_arn = aws_sns_topic.test.arn
  }
}
`, rName))
}

func testAccWorkteamConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccWorkteamOIDCBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_workteam" "test" {
  workteam_name  = %[1]q
  workforce_name = aws_sagemaker_workforce.test.id
  description%[1]q

  member_definition {
dc_member_definition {
ps = [%[1]q]

  }

  tags = {
2]q = %[3]q
  }
funcName, tagKey1, tagValue1))
}

func testAccWorkteamConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccWorkteamOIDCBaseConfig(rName), fmt.Sprintf(`
resource "aws_sagemaker_workteam" "test" {
  workteam_name  = %[1]q
  workforce_name = aws_sagemaker_workforce.test.id
  description%[1]q

  member_definition {
dc_member_definition {
ps = [%[1]q]

  }

  tags = {
2]q = %[3]q
4]q = %[5]q
  }
funcName, tagKey1, tagValue1, tagKey2, tagValue2))
}
