// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elb_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func TestAccELBCookieStickinessPolicy_basic(t *testing.T) {
	ctx := context.Background()
	var policy elb.DescribeLoadBalancerPoliciesOutput
	resourceName := "aws_lb_cookie_stickiness_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, elb.ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckLBCookieStickinessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLBCookieStickinessPolicyConfig_zeroExpiration,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBCookieStickinessPolicyExists(ctx, resourceName, &policy),
					resource.TestCheckResourceAttr(resourceName, "cookie_expiration_period", "0"),
					testAccCheckLBCookieStickinessPolicyExpirationPeriod(&policy, 0),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccELBCookieStickinessPolicy_positiveExpiration(t *testing.T) {
	ctx := context.Background()
	var policy elb.DescribeLoadBalancerPoliciesOutput
	resourceName := "aws_lb_cookie_stickiness_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, elb.ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckLBCookieStickinessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLBCookieStickinessPolicyConfig_positiveExpiration,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBCookieStickinessPolicyExists(ctx, resourceName, &policy),
					resource.TestCheckResourceAttr(resourceName, "cookie_expiration_period", "3600"),
					testAccCheckLBCookieStickinessPolicyExpirationPeriod(&policy, 3600),
				),
			},
		},
	})
}

func TestAccELBCookieStickinessPolicy_unsetExpiration(t *testing.T) {
	ctx := context.Background()
	var policy elb.DescribeLoadBalancerPoliciesOutput
	resourceName := "aws_lb_cookie_stickiness_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, elb.ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckLBCookieStickinessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLBCookieStickinessPolicyConfig_unsetExpiration,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBCookieStickinessPolicyExists(ctx, resourceName, &policy),
					resource.TestCheckResourceAttr(resourceName, "cookie_expiration_period", "0"),
					testAccCheckLBCookieStickinessPolicyExpirationPeriod(&policy, 0),
				),
			},
		},
	})
}

func testAccCheckLBCookieStickinessPolicyExists(ctx context.Context, n string, policy *elb.DescribeLoadBalancerPoliciesOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LBCookieStickinessPolicy ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).ELBV1Conn(ctx)
		parts, err := parseLBCookieStickinessPolicyID(rs.Primary.ID)
		if err != nil {
			return err
		}
		lbName, policyName := parts[0], parts[2]

		resp, err := conn.DescribeLoadBalancerPolicies(ctx, &elb.DescribeLoadBalancerPoliciesInput{
			LoadBalancerName: aws.String(lbName),
			PolicyNames:      []string{policyName},
		})
		if err != nil {
			return fmt.Errorf("error describing LBCookieStickinessPolicy: %w", err)
		}

		if len(resp.PolicyDescriptions) == 0 {
			return fmt.Errorf("LBCookieStickinessPolicy not found")
		}

		*policy = *resp
		return nil
	}
}

func testAccCheckLBCookieStickinessPolicyExpirationPeriod(policy *elb.DescribeLoadBalancerPoliciesOutput, expected int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, attr := range policy.PolicyDescriptions[0].PolicyAttributeDescriptions {
			if aws.ToString(attr.AttributeName) == "CookieExpirationPeriod" {
				if aws.ToInt64(attr.AttributeValue) != int64(expected) {
					return fmt.Errorf("Expected CookieExpirationPeriod %d, got %d", expected, aws.ToInt64(attr.AttributeValue))
				}
				return nil
			}
		}
		if expected == 0 {
			return nil // No CookieExpirationPeriod attribute expected for 0
		}
		return fmt.Errorf("CookieExpirationPeriod attribute not found")
	}
}

func testAccCheckLBCookieStickinessPolicyDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).ELBV1Conn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_lb_cookie_stickiness_policy" {
				continue
			}

			parts, err := parseLBCookieStickinessPolicyID(rs.Primary.ID)
			if err != nil {
				return err
			}
			lbName, policyName := parts[0], parts[2]

			_, err = conn.DescribeLoadBalancerPolicies(ctx, &elb.DescribeLoadBalancerPoliciesInput{
				LoadBalancerName: aws.String(lbName),
				PolicyNames:      []string{policyName},
			})
			if err != nil && !strings.Contains(err.Error(), "PolicyNotFound") {
				return fmt.Errorf("error checking LBCookieStickinessPolicy destroy: %w", err)
			}
		}
		return nil
	}
}

func parseLBCookieStickinessPolicyID(id string) ([]string, error) {
	parts := strings.Split(id, ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid ID format: %s (expected load_balancer:port:policy_name)", id)
	}
	return parts, nil
}

const testAccLBCookieStickinessPolicyProvider = `
provider "aws" {
  access_key = "test"
  secret_key = "test"
  region     = "us-east-1"
  endpoints {
    elasticloadbalancing = "http://localhost:4566"
    sts = "http://localhost:4566"
  }
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}
`

const testAccLBCookieStickinessPolicyConfig_zeroExpiration = testAccLBCookieStickinessPolicyProvider + `
resource "aws_elb" "test" {
  name               = "test-elb-zero"
  availability_zones = ["us-east-1a"]
  listener {
    instance_port      = 8088
    instance_protocol  = "http"
    lb_port            = 8088
    lb_protocol        = "http"
  }
}

resource "aws_lb_cookie_stickiness_policy" "test" {
  name                     = "test-policy-zero"
  load_balancer            = aws_elb.test.name
  lb_port                  = 8088
  cookie_expiration_period = 0
}
`

const testAccLBCookieStickinessPolicyConfig_positiveExpiration = testAccLBCookieStickinessPolicyProvider + `
resource "aws_elb" "test" {
  name               = "test-elb-pos"
  availability_zones = ["us-east-1a"]
  listener {
    instance_port      = 8088
    instance_protocol  = "http"
    lb_port            = 8088
    lb_protocol        = "http"
  }
}

resource "aws_lb_cookie_stickiness_policy" "test" {
  name                     = "test-policy-pos"
  load_balancer            = aws_elb.test.name
  lb_port                  = 8088
  cookie_expiration_period = 3600
}
`

const testAccLBCookieStickinessPolicyConfig_unsetExpiration = testAccLBCookieStickinessPolicyProvider + `
resource "aws_elb" "test" {
  name               = "test-elb-unset"
  availability_zones = ["us-east-1a"]
  listener {
    instance_port      = 8088
    instance_protocol  = "http"
    lb_port            = 8088
    lb_protocol        = "http"
  }
}

resource "aws_lb_cookie_stickiness_policy" "test" {
  name          = "test-policy-unset"
  load_balancer = aws_elb.test.name
  lb_port       = 8088
  # cookie_expiration_period intentionally omitted
}
`
