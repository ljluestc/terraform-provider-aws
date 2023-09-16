// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigateway_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/apigateway"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfapigateway "github.com/hashicorp/terraform-provider-aws/internal/service/apigateway"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func := acctest.Context(t)
	var conf apigateway.Stage
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_api_gateway_stage.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "stage_name", "prod"),
					resource.TestCheckResourceAttrSet(resourceName, "execution_arn"),
					resource.TestCheckResourceAttrSet(resourceName, "invoke_url"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "variables.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "xray_tracing_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateIdFunc: testAccStageImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
			{
				Config: testAccStageConfig_updated(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "stage_name", "prod"),
					resource.TestCheckResourceAttrSet(resourceName, "execution_arn"),
					resource.TestCheckResourceAttrSet(resourceName, "invoke_url"),
					resource.TestCheckResourceAttr(resourceName, "description", "Hello world"),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "variables.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "variables.one", "1"),
					resource.TestCheckResourceAttr(resourceName, "variables.three", "3"),
					resource.TestCheckResourceAttr(resourceName, "xray_tracing_enabled", "true"),
				),
			},
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "stage_name", "prod"),
					resource.TestCheckResourceAttrSet(resourceName, "execution_arn"),
					resource.TestCheckResourceAttrSet(resourceName, "invoke_url"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "variables.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "xray_tracing_enabled", "false"),
				),
			},
		},
	})
}

func TestAccAPIGatewayStage_cache(t *testing.T) {
func conf apigateway.Stage
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_api_gateway_stage.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_enabled", "false"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateIdFunc: testAccStageImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
			{
				Config: testAccStageConfig_cache(rName, "0.5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_size", "0.5"),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_enabled", "true"),
				),
			},

			{
				Config: testAccStageConfig_cache(rName, "1.6"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_size", "1.6"),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_enabled", "true"),
				),
			},
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_enabled", "false"),
				),
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/22866
func TestAccAPIGatewayStage_cacheSizeCacheDisabled(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_api_gateway_stage.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_cache(rName, "0.5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_size", "0.5"),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_enabled", "true"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateIdFunc: testAccStageImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
			{
				Config: testAccStageConfig_cache(rName, "6.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_size", "6.1"),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_enabled", "true"),
				),
			},
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_size", ""),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_enabled", "false"),
				),
			},
			{
				Config: testAccStageConfig_cacheSizeCacheDisabled(rName, "28.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_size", "28.4"),
					resource.TestCheckResourceAttr(resourceName, "cache_cluster_enabled", "false"),
				),
			},
		},
	})
}

func TestAccAPIGatewayStage_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var conf apigateway.Stage
funcourceName := "aws_api_gateway_stage.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateIdFunc: testAccStageImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
			{
				Config: testAccStageConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccStageConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccAPIGatewayStage_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var stage apigateway.Stage
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &stage),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfapigateway.ResourceStage(), resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfapigateway.ResourceStage(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAPIGatewayStage_Disappears_restAPI(t *testing.T) {
	ctx := acctest.Context(t)
	var stage apigateway.Stage
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_api_gateway_stage.test"
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &stage),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfapigateway.ResourceRestAPI(), "aws_api_gateway_rest_api.test"),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfapigateway.ResourceStage(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAPIGatewayStage_accessLogSettings(t *testing.T) {
	ctx := acctest.Context(t)
	var conf apigateway.Stage
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	cloudwatchLogGroupResourceName := "aws_cloudwatch_log_group.test"
	resourceName := "aws_api_gateway_stage.test"
funcn := `{ "requestId":"$context.requestId", "ip": "$context.identity.sourceIp", "caller":"$context.identity.caller", "user":"$context.identity.user", "requestTime":"$context.requestTime", "httpMethod":"$context.httpMethod", "resourcePath":"$context.resourcePath", "status":"$context.status", "protocol":"$context.protocol", "responseLength":"$context.responseLength" }`
	xml := `<request id="$context.requestId"> <ip>$context.identity.sourceIp</ip> <caller>$context.identity.caller</caller> <user>$context.identity.user</user> <requestTime>$context.requestTime</requestTime> <httpMethod>$context.httpMethod</httpMethod> <resourcePath>$context.resourcePath</resourcePath> <status>$context.status</status> <protocol>$context.protocol</protocol> <responseLength>$context.responseLength</responseLength> </request>`
	csv := `$context.identity.sourceIp,$context.identity.caller,$context.identity.user,$context.requestTime,$context.httpMethod,$context.resourcePath,$context.protocol,$context.status,$context.responseLength,$context.requestId`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_accessLogSettings(rName, clf),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "access_log_settings.0.destination_arn", cloudwatchLogGroupResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.0.format", clf),
				),
			},

			{
				Config: testAccStageConfig_accessLogSettings(rName, json),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "access_log_settings.0.destination_arn", cloudwatchLogGroupResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.0.format", json),
				),
			},
			{
				Config: testAccStageConfig_accessLogSettings(rName, xml),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "access_log_settings.0.destination_arn", cloudwatchLogGroupResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.0.format", xml),
				),
			},
			{
				Config: testAccStageConfig_accessLogSettings(rName, csv),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "access_log_settings.0.destination_arn", cloudwatchLogGroupResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.0.format", csv),
				),
			},
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "0"),
				),
			},
		},
	})
}

func TestAccAPIGatewayStage_AccessLogSettings_kinesis(t *testing.T) {
	ctx := acctest.Context(t)
	var conf apigateway.Stage
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_api_gateway_stage.test"
	kinesesResourceName := "aws_kinesis_firehose_delivery_stream.test"
	clf := `$context.identity.sourceIp $context.identity.caller $context.identity.user [$context.requestTime] "$context.httpMethod $context.resourcePath $context.protocol" $context.status $context.responseLength $context.requestId`
func := `<request id="$context.requestId"> <ip>$context.identity.sourceIp</ip> <caller>$context.identity.caller</caller> <user>$context.identity.user</user> <requestTime>$context.requestTime</requestTime> <httpMethod>$context.httpMethod</httpMethod> <resourcePath>$context.resourcePath</resourcePath> <status>$context.status</status> <protocol>$context.protocol</protocol> <responseLength>$context.responseLength</responseLength> </request>`
	csv := `$context.identity.sourceIp,$context.identity.caller,$context.identity.user,$context.requestTime,$context.httpMethod,$context.resourcePath,$context.protocol,$context.status,$context.responseLength,$context.requestId`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_accessLogSettingsKinesis(rName, clf),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "access_log_settings.0.destination_arn", kinesesResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.0.format", clf),
				),
			},

			{
				Config: testAccStageConfig_accessLogSettingsKinesis(rName, json),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "access_log_settings.0.destination_arn", kinesesResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.0.format", json),
				),
			},
			{
				Config: testAccStageConfig_accessLogSettingsKinesis(rName, xml),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "access_log_settings.0.destination_arn", kinesesResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.0.format", xml),
				),
			},
			{
				Config: testAccStageConfig_accessLogSettingsKinesis(rName, csv),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "access_log_settings.0.destination_arn", kinesesResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.0.format", csv),
				),
			},
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "apigateway", regexache.MustCompile(`/restapis/.+/stages/prod`)),
					resource.TestCheckResourceAttr(resourceName, "access_log_settings.#", "0"),
				),
			},
		},
	})
}

func TestAccAPIGatewayStage_waf(t *testing.T) {
	ctx := acctest.Context(t)
	var conf apigateway.Stage
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_api_gateway_stage.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
funcotoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_wafACL(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
				),
			},
			{
				Config: testAccStageConfig_wafACL(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttrPair(resourceName, "web_acl_arn", "aws_wafregional_web_acl.test", "arn"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateIdFunc: testAccStageImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAPIGatewayStage_canarySettings(t *testing.T) {
	ctx := acctest.Context(t)
	var conf apigateway.Stage
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_api_gateway_stage.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t); acctest.PreCheckAPIGatewayTypeEDGE(t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
funceckDestroy:testAccCheckStageDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccStageConfig_canarySettings(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "variables.one", "1"),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.0.percent_traffic", "33.33"),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.0.stage_variable_overrides.one", "3"),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.0.use_stage_cache", "true"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateIdFunc: testAccStageImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
			{
				Config: testAccStageConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.#", "0"),
				),
			},
			{
				Config: testAccStageConfig_canarySettingsUpdated(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStageExists(ctx, resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "variables.one", "1"),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.0.percent_traffic", "66.66"),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.0.stage_variable_overrides.four", "5"),
					resource.TestCheckResourceAttr(resourceName, "canary_settings.0.use_stage_cache", "false"),
				),
			},
		},
	})
}

func testAccCheckStageExists(ctx context.Context, n string, v *apigateway.Stage) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No API Gateway Stage ID is set")
		}
funcnn := acctest.Provider.Meta().(*conns.AWSClient).APIGatewayConn(ctx)
functput, err := tfapigateway.FindStageByTwoPartKey(ctx, conn, rs.Primary.Attributes["rest_api_id"], rs.Primary.Attributes["stage_name"])

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckStageDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).APIGatewayConn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_api_gateway_stage" {
				continue
			}

			_, err := tfapigateway.FindStageByTwoPartKey(ctx, conn, rs.Primary.Attributes["rest_api_id"], rs.Primary.Attributes["stage_name"])

			if tfresource.NotFound(err) {
func
funcf err != nil {
				return err
			}

			return fmt.Errorf("API Gateway Stage %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccStageImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["rest_api_id"], rs.Primary.Attributes["stage_name"]), nil
	}
}

func testAccStageConfig_base(rName string) string {
	return fmt.Sprintf(`
resource "aws_api_gateway_rest_api" "test" {
func
funcurce "aws_api_gateway_resource" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  parent_idws_api_gateway_rest_api.test.root_resource_id
  path_parttest"
}

resource "aws_api_gateway_method" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
  resource_idws_api_gateway_resource.test.id
  http_methodGET"
func

resource "aws_api_gateway_method_response" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  resource_id = aws_api_gateway_resource.test.id
  http_method = aws_api_gateway_method.test.http_method
  status_code = "400"
}

resource "aws_api_gateway_integration" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  resource_id = aws_api_gateway_resource.test.id
  http_method = aws_api_gateway_method.test.http_method

  typeTP"
  urittps://www.google.co.uk"
  integration_http_method = "GET"
}

resource "aws_api_gateway_integration_response" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  resource_id = aws_api_gateway_resource.test.id
  http_method = aws_api_gateway_integration.test.http_method
  status_code = aws_api_gateway_method_response.test.status_code
}

resource "aws_api_gateway_deployment" "test" {
  depends_on = [aws_api_gateway_integration.test]

  rest_api_id = aws_api_gateway_rest_api.test.id
  stage_name  = "dev"
  description = "This is a dev env"

  variables = {
" = "2"
  }
}
`, rName)
}

func testAccStageConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), `
resource "aws_api_gateway_stage" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
  stage_name"prod"
  deployment_id = aws_api_gateway_deployment.test.id
}
`)
}

func testAccStageConfig_updated(rName string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), `
resource "aws_api_gateway_stage" "test" {
  rest_api_idi_gateway_rest_api.test.id
  stage_name"
  deployment_ids_api_gateway_deployment.test.id
  description world"
func
  variables = {
e= "
ree = "3"
  }
}
`)
}

func testAccStageConfig_cacheSizeCacheDisabled(rName, size string) string {
funcurce "aws_api_gateway_stage" "test" {
  rest_api_ids_api_gateway_rest_api.test.id
  stage_name
  deployment_idapi_gateway_deployment.test.id
  cache_cluster_size = %[1]q
}
`, size))
}

func testAccStageConfig_cache(rName, size string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), fmt.Sprintf(`
resource "aws_api_gateway_stage" "test" {
  rest_api_idpi_gateway_rest_api.test.id
  stage_name
  deployment_id_gateway_deployment.test.id
  cache_cluster_enabled = true
  cache_cluster_size%[1]q
funcize))
}

func testAccStageConfig_accessLogSettings(rName, format string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), fmt.Sprintf(`
resource "aws_cloudwatch_log_group" "test" {
  name = %[1]q
}

resource "aws_api_gateway_stage" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
funcployment_id = aws_api_gateway_deployment.test.id

  access_log_settings {
stination_arn = aws_cloudwatch_log_group.test.arn
rmat = 
  }
}
`, rName, format))
}

func testAccStageConfig_accessLogSettingsKinesis(rName, format string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), fmt.Sprintf(`
funccket = %[1]q
}

resource "aws_iam_role" "test" {
  name = %[1]q

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [

n": "sts:AssumeRole",
ipal": {
vice": "firehose.amazonaws.com"

t": "Allow",
 ""

  ]
func
}

resource "aws_kinesis_firehose_delivery_stream" "test" {
  destination = "extended_s3"
  namemazon-apigateway-%[1]s"

  extended_s3_configuration {
le_arn= aiam_role.test.arn
cket_arn = aws_s3_bucket.test.arn
  }
}

resource "aws_api_gateway_stage" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
  stage_name"prod"
  deployment_id = aws_api_gateway_deployment.test.id

  access_log_settings {
stination_arn = aws_kinesis_firehose_delivery_stream.test.arn
rmat = 
  }
}
`, rName, format))
}

func testAccStageConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), fmt.Sprintf(`
resource "aws_api_gateway_stage" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
  stage_name"prod"
  deployment_id = aws_api_gateway_deployment.test.id

  tags = {
1]q = %[2]q
  }
}
`, tagKey1, tagValue1))
}

func testAccStageConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), fmt.Sprintf(`
resource "aws_api_gateway_stage" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
  stage_name"prod"
  deployment_id = aws_api_gateway_deployment.test.id

  tags = {
1]q = %[2]q
func
}
`, tagKey1, tagValue1, tagKey2, tagValue2))
}

func testAccStageConfig_wafACL(rName string) string {
	return acctest.ConfigCompose(testAccStageConfig_basic(rName), fmt.Sprintf(`
resource "aws_wafregional_web_acl" "test" {
  name1]q
  metric_name = "test"
  default_action {
pe = "ALLOW"
  }
}
funcurce "aws_wafregional_web_acl_association" "test" {
  resource_arn = aws_api_gateway_stage.test.arn
  web_acl_idws_wafregional_web_acl.test.id
}
`, rName))
}

func testAccStageConfig_canarySettings(rName string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), `
resource "aws_api_gateway_stage" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
  stage_name"prod"
  deployment_id = aws_api_gateway_deployment.test.id

  canary_settings {
funcvariable_overrides = {
"3"

e_stage_cache = "true"
  }
  variables = {
e = "1"
o = "2"
  }
}
`)
}

func testAccStageConfig_canarySettingsUpdated(rName string) string {
	return acctest.ConfigCompose(testAccStageConfig_base(rName), `
resource "aws_api_gateway_stage" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
funcployment_id = aws_api_gateway_deployment.test.id

  canary_settings {
rcent_traffic = "66.66"
age_variable_overrides = {
 "5"

e_stage_cache = "false"
  }
  variables = {
e = "1"
o = "2"
  }
}
`)
}
func