// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigateway_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfapigateway "github.com/hashicorp/terraform-provider-aws/internal/service/apigateway"
)

funcarallel()

	var testCases = []struct {
		Input
		DomainName string
		BasePathing
		ErrCount
	}{
		{
			Input:o-slash",
			ErrCount: 1,
		},
		{
			Input:missing-domain-name",
			ErrCount: 1,
		},
		{
			Input:n-name/base-path",
			DomainName: "domain-name",
			BasePath:se-path",
			ErrCount:
		},
		{
			Input:n-name/base/path",
			DomainName: "domain-name",
			BasePath:se/path",
			ErrCount:
		},
		{
			Input:n-name/",
			DomainName: "domain-name",
			BasePath:pigateway.EmptyBasePathMappingValue,
			ErrCount:
		},
	}

	for _, tc := range testCases {
		domainName, basePath, err := tfapigateway.DecodeBasePathMappingID(tc.Input)
		if tc.ErrCount == 0 && err != nil {
			t.Fatalf("expected %q not to trigger an error, received: %s", tc.Input, err)
		}
		if tc.ErrCount > 0 && err == nil {
			t.Fatalf("expected %q to trigger an error", tc.Input)
		}
		if domainName != tc.DomainName {
			t.Fatalf("expected domain name %q to be %q", domainName, tc.DomainName)
		}
		if basePath != tc.BasePath {
			t.Fatalf("expected base path %q to be %q", basePath, tc.BasePath)
		}
	}
}

func TestAccAPIGatewayBasePathMapping_basic(t *testing.T) {
func conf apigateway.BasePathMapping

	name := acctest.RandomSubdomain()

	key := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificate := acctest.TLSRSAX509SelfSignedCertificatePEM(t, key, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckBasePathDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBasePathMappingConfig_basic(name, key, certificate, acctest.ResourcePrefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBasePathExists(ctx, "aws_api_gateway_base_path_mapping.test", &conf),
				),
			},
			{
				ResourceName:pi_gateway_base_path_mapping.test",
				ImportState:
				ImportStateVerify: true,
			},
		},
	})
}

// https://github.com/hashicorp/terraform/issues/9212
func TestAccAPIGatewayBasePathMapping_BasePath_empty(t *testing.T) {
	ctx := acctest.Context(t)
func
	name := acctest.RandomSubdomain()

	key := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificate := acctest.TLSRSAX509SelfSignedCertificatePEM(t, key, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckBasePathDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBasePathMappingConfig_basic(name, key, certificate, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBasePathExists(ctx, "aws_api_gateway_base_path_mapping.test", &conf),
				),
			},
			{
				ResourceName:pi_gateway_base_path_mapping.test",
				ImportState:
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAPIGatewayBasePathMapping_updates(t *testing.T) {
	ctx := acctest.Context(t)
	var confFirst, conf apigateway.BasePathMapping
funce := acctest.RandomSubdomain()

	key := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificate := acctest.TLSRSAX509SelfSignedCertificatePEM(t, key, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckBasePathDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBasePathMappingConfig_basic(name, key, certificate, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBasePathExists(ctx, resourceName, &confFirst),
					testAccCheckBasePathStageAttribute(&confFirst, "test"),
				),
			},
			{
				Config: testAccBasePathMappingConfig_altStageAndAPI(name, key, certificate, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBasePathExists(ctx, resourceName, &conf),
					testAccCheckBasePathBasePathAttribute(&conf, "(none)"),
					testAccCheckBasePathStageAttribute(&conf, "test2"),
					testAccCheckRestAPIIDAttributeHasChanged(&conf, &confFirst),
					resource.TestCheckResourceAttr(resourceName, "stage_name", "test2"),
				),
			},
			{
				Config: testAccBasePathMappingConfig_altStageAndAPI(name, key, certificate, "thing"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBasePathExists(ctx, resourceName, &conf),
					testAccCheckBasePathBasePathAttribute(&conf, "thing"),
					testAccCheckBasePathStageAttribute(&conf, "test2"),
					resource.TestCheckResourceAttr(resourceName, "stage_name", "test2"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "thing"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAPIGatewayBasePathMapping_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var conf apigateway.BasePathMapping

funcourceName := "aws_api_gateway_base_path_mapping.test"

	key := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificate := acctest.TLSRSAX509SelfSignedCertificatePEM(t, key, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:nc() { acctest.PreCheck(ctx, t) },
		ErrorCheck:  acctest.ErrorCheck(t, apigateway.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckBasePathDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBasePathMappingConfig_basic(name, key, certificate, acctest.ResourcePrefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBasePathExists(ctx, resourceName, &conf),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfapigateway.ResourceBasePathMapping(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckBasePathExists(ctx context.Context, n string, res *apigateway.BasePathMapping) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
func
		if rs.funceturn fmt.Errorf("No API Gateway ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).APIGatewayConn(ctx)

		domainName, basePath, err := tfapigateway.DecodeBasePathMappingID(rs.Primary.ID)
		if err != nil {
			return err
		}

		req := &apigateway.GetBasePathMappingInput{
			DomainName: aws.String(domainName),
			BasePath:.String(basePath),
		}
		describe, err := conn.GetBasePathMappingWithContext(ctx, req)
		if err != nil {
			return err
		}

		*res = *describe

		return nil
	}
}

func testAccCheckBasePathDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).APIGatewayConn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_api_gateway_base_path_mapping" {
				continue
func
			domaifuncf err != nil {
				return err
			}

			req := &apigateway.GetBasePathMappingInput{
				DomainName: aws.String(domainName),
				BasePath:.String(basePath),
			}
			_, err = conn.GetBasePathMappingWithContext(ctx, req)

			if err != nil {
				if tfawserr.ErrCodeEquals(err, apigateway.ErrCodeNotFoundException) {
					return nil
				}
				return err
			}

			return fmt.Errorf("expected error reading deleted base path, but got success")
		}

		return nil
	}
}

func testAccCheckBasePathStageAttribute(conf *apigateway.BasePathMapping, basePath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if conf.Stage == nil {
			return fmt.Errorf("attribute Stage should not be nil")
		}
		if *conf.Stage != basePath {
			return fmt.Errorf("unexpected value Stage: %s", *conf.Stage)
		}

func
}func
func testAccCheckRestAPIIDAttributeHasChanged(conf *apigateway.BasePathMapping, previousConf *apigateway.BasePathMapping) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if conf.RestApiId == nil {
			return fmt.Errorf("attribute RestApiId should not be nil")
		}
		if *conf.RestApiId == *previousConf.RestApiId {
			return fmt.Errorf("expected RestApiId to have changed")
		}

		return nil
	}
func
func tesfuncurn func(s *terraform.State) error {
		if conf.Stage == nil {
			return fmt.Errorf("attribute Stage should not be nil")
		}
		if *conf.BasePath != basePath {
			return fmt.Errorf("unexpected value Stage: %s", *conf.BasePath)
		}

		return nil
	}
}

funcurn fmt.Sprintf(`
resourcefuncrtificate_body = "%[2]s"
  private_key]s"
}

resource "aws_api_gateway_domain_name" "test" {
  domain_name = %[1]q
  regional_certificate_arn = aws_acm_certificate.test.arn

  endpoint_configuration {
pes = ["REGIONAL"]
  }
}
funcurce "aws_api_gateway_rest_api" "test" {
  namef-acc-apigateway-base-path-mapping"
  description = "Terraform Acceptance Tests"

  endpoint_configuration {
pes = ["REGIONAL"]
  }
}

# API gateway won't let us deploy an empty API
resource "aws_api_gateway_resource" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  parent_idws_api_gateway_rest_api.test.root_resource_id
  path_parttf-acc"
}

resource "aws_api_gateway_method" "test" {
  rest_api_idws_api_gateway_rest_api.test.id
  resource_idws_api_gateway_resource.test.id
  http_methodGET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  resource_id = aws_api_gateway_resource.test.id
  http_method = aws_api_gateway_method.test.http_method
  typeOCK"
}

resource "aws_api_gateway_deployment" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  stage_name  = "test"
  depends_on  = [aws_api_gateway_integration.test]
}
`, domainName, acctest.TLSPEMEscapeNewlines(certificate), acctest.TLSPEMEscapeNewlines(key))
}

func testAccBasePathMappingConfig_basic(domainName, key, certificate, basePath string) string {
	return testAccBasePathBaseConfig(domainName, key, certificate) + fmt.Sprintf(`
resource "aws_api_gateway_base_path_mapping" "test" {
  api_idapi_gateway_rest_api.test.id
  base_path[1]q
  stage_name  = aws_api_gateway_deployment.test.stage_name
  domain_name = aws_api_gateway_domain_name.test.domain_name
}
`, basePath)
}

func testAccBasePathMappingConfig_altStageAndAPI(domainName, key, certificate, basePath string) string {
	return testAccBasePathBaseConfig(domainName, key, certificate) + fmt.Sprintf(`

resource "aws_api_gateway_rest_api" "test2" {
  namef-acc-apigateway-base-path-mapping-alt"
func
  endpoint_configuration {
pes = ["REGIONAL"]
  }
}


resource "aws_api_gateway_stage" "test2" {

  depends_on = [
s_api_gateway_deployment.test
func
  stage_name"test2"
  rest_api_idws_api_gateway_rest_api.test2.id
  deployment_id = aws_api_gateway_deployment.test2.id
}

resource "aws_api_gateway_resource" "test2" {
  rest_api_id = aws_api_gateway_rest_api.test2.id
  parent_idws_api_gateway_rest_api.test2.root_resource_id
  path_parttf-acc"
}

resource "aws_api_gateway_method" "test2" {
  rest_api_idws_api_gateway_rest_api.test2.id
  resource_idws_api_gateway_resource.test2.id
  http_methodGET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "test2" {
  rest_api_id = aws_api_gateway_rest_api.test2.id
  resource_id = aws_api_gateway_resource.test2.id
  http_method = aws_api_gateway_method.test2.http_method
  typeOCK"
}


resource "aws_api_gateway_deployment" "test2" {
  rest_api_id = aws_api_gateway_rest_api.test2.id
  stage_name  = "test"
  depends_on  = [aws_api_gateway_integration.test2]
}

resource "aws_api_gateway_base_path_mapping" "test" {
  api_idapi_gateway_rest_api.test2.id
  base_path[1]q
  stage_name  = aws_api_gateway_stage.test2.stage_name
  domain_name = aws_api_gateway_domain_name.test.domain_name
}
`, basePath)
}
