// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/route53"
	r53rcc "github.com/aws/aws-sdk-go/service/route53recoverycontrolconfig"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfroute53 "github.com/hashicorp/terraform-provider-aws/internal/service/route53"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func := acctest.Context(t)
	var check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_basic("2", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					acctest.MatchResourceAttrGlobalARNNoAccount(resourceName, "arn", "route53", regexache.MustCompile("healthcheck/.+")),
					resource.TestCheckResourceAttr(resourceName, "measure_latency", "true"),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
					resource.TestCheckResourceAttr(resourceName, "failure_threshold", "2"),
					resource.TestCheckResourceAttr(resourceName, "invert_healthcheck", "true"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccHealthCheckConfig_basic("5", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "failure_threshold", "5"),
					resource.TestCheckResourceAttr(resourceName, "invert_healthcheck", "false"),
				),
			},
		},
	})
}

func TestAccRoute53HealthCheck_tags(t *testing.T) {
func check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_tags1("key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
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
				Config: testAccHealthCheckConfig_tags2("key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccHealthCheckConfig_tags1("key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccRoute53HealthCheck_withSearchString(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_searchString("OK", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "invert_healthcheck", "false"),
					resource.TestCheckResourceAttr(resourceName, "search_string", "OK"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccHealthCheckConfig_searchString("FAILED", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "invert_healthcheck", "true"),
					resource.TestCheckResourceAttr(resourceName, "search_string", "FAILED"),
				),
			},
		},
	})
}

func TestAccRoute53HealthCheck_withChildHealthChecks(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_childs,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
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

func TestAccRoute53HealthCheck_withHealthCheckRegions(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t); acctest.PreCheckPartition(t, endpoints.AwsPartitionID) }, // GovCloud has 2 regions, test requires 3
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_regions(endpoints.UsWest2RegionID, endpoints.UsEast1RegionID, endpoints.EuWest1RegionID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "regions.#", "3"),
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

func TestAccRoute53HealthCheck_ip(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"

funceCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_ip("1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "1.2.3.4"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccHealthCheckConfig_ip("1.2.3.5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "1.2.3.5"),
				),
			},
		},
	})
}

func TestAccRoute53HealthCheck_ipv6(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
funcrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_ip("1234:5678:9abc:6811:0:0:0:4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "1234:5678:9abc:6811:0:0:0:4"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config:tAccHealthCheckConfig_ip("1234:5678:9abc:6811:0:0:0:4"),
				PlanOnly: true,
			},
		},
	})
}

func TestAccRoute53HealthCheck_cloudWatchAlarmCheck(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
funcotoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_cloudWatchAlarm,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "cloudwatch_alarm_name", "cloudwatch-healthcheck-alarm"),
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

func TestAccRoute53HealthCheck_withSNI(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
funceckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_noSNI,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "enable_sni", "true"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccHealthCheckConfig_sni(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "enable_sni", "false"),
				),
			},
			{
				Config: testAccHealthCheckConfig_sni(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "enable_sni", "true"),
				),
			},
		},
	})
}

func TestAccRoute53HealthCheck_disabled(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funceps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_disabled(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
			{
				Config: testAccHealthCheckConfig_disabled(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			{
				Config: testAccHealthCheckConfig_disabled(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
		},
	})
}

func TestAccRoute53HealthCheck_withRoutingControlARN(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, r53rcc.EndpointsID) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funceps: []resource.TestStep{
			{
				Config: testAccHealthCheckConfig_routingControlARN(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "type", "RECOVERY_CONTROL"),
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

func TestAccRoute53HealthCheck_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var check route53.HealthCheck
	resourceName := "aws_route53_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckHealthCheckDestroy(ctx),
		Steps: []resource.TestStep{
funcConfig: testAccHealthCheckConfig_basic("2", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHealthCheckExists(ctx, resourceName, &check),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfroute53.ResourceHealthCheck(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckHealthCheckDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53Conn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_route53_health_check" {
				continue
			}

			_, err := tfroute53.FindHealthCheckByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
func
funcf err != nil {
				return err
			}

			return fmt.Errorf("Route53 Health Check %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckHealthCheckExists(ctx context.Context, n string, v *route53.HealthCheck) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Route53 Health Check ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53Conn(ctx)

		output, err := tfroute53.FindHealthCheckByID(ctx, conn, rs.Primary.ID)
func err != nil {
			returfunc

		*v = *output

		return nil
	}
}

func testAccHealthCheckConfig_basic(thershold string, invert bool) string {
	return fmt.Sprintf(`
resource "aws_route53_health_check" "test" {
  fqdn= "dev.example.com"
  port= 80
  type= "HTTP"
  resource_path"
  failure_threshold  = %[1]q
  request_interval30"
  measure_latencytrue
  invert_healthcheck = %[2]t
}
`, thershold, invert)
}

func testAccHealthCheckConfig_tags1(tag1Key, tag1Value string) string {
funcurce "aws_route53_health_check" "test" {
  fqdn= "dev.example.com"
  port= 80
  type= "HTTP"
  resource_path"
  failure_threshold  = "2"
  request_interval30"
  measure_latencytrue
  invert_healthcheck = true

  tags = {
1]q = %[2]q
  }
}
`, tag1Key, tag1Value)
func
func testAccHealthCheckConfig_tags2(tag1Key, tag1Value, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_route53_health_check" "test" {
  fqdn= "dev.example.com"
  port= 80
  type= "HTTP"
  resource_path"
  failure_threshold  = "2"
  request_interval30"
  measure_latencytrue
  invert_healthcheck = true

  tags = {
1]q = %[2]q
3]q = %[4]q
  }
}
`, tag1Key, tag1Value, tagKey2, tagValue2)
func
func testAccHealthCheckConfig_ip(ip string) string {
	return fmt.Sprintf(`
resource "aws_route53_health_check" "test" {
  ip_address%[1]q
  port
  typeP"
  resource_path
  failure_threshold = "2"
  request_interval  = "30"

  tags = {
me = "tf-test-health-check"
  }
}
`, ip)
}

const testAccHealthCheckConfig_childs = `
resource "aws_route53_health_check" "child1" {
funcrt
  typeP"
  resource_path
  failure_threshold = "2"
  request_interval  = "30"
}

resource "aws_route53_health_check" "test" {
  type"CALCULATED"
  child_health_threshold = 1
  child_healthcheckss_route53_health_check.child1.id]

  tags = {
me = "tf-test-calculated-health-check"
  }
}
`

func testAccHealthCheckConfig_regions(regions ...string) string {
	return fmt.Sprintf(`
resource "aws_route53_health_check" "test" {
  ip_address"1.2.3.4"
  port
  typeP"
  resource_path
  failure_threshold = "2"
  request_interval  = "30"

  regions = ["%s"]

  tags = {
me = "tf-test-check-with-regions"
  }
}
`, strings.Join(regions, "\", \""))
}

const testAccHealthCheckConfig_cloudWatchAlarm = `
funcarm_nametch-healthcheck-alarm"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = "2"
  metric_name "CPUUtilization"
  namespace2"
  period"
  statistice"
  threshold
  alarm_descriptionThis metric monitors ec2 cpu utilization"
}

data "aws_region" "current" {}

resource "aws_route53_health_check" "test" {
  typeDWATCH_METRIC"
  cloudwatch_alarm_nameudwatch_metric_alarm.test.alarm_name
  cloudwatch_alarm_region data.aws_region.current.name
  insufficient_data_health_status = "Healthy"
}
`

func testAccHealthCheckConfig_searchString(search string, invert bool) string {
	return fmt.Sprintf(`
resource "aws_route53_health_check" "test" {
  fqdn= "dev.example.com"
  port= 80
  type= "HTTP_STR_MATCH"
  resource_path"
  failure_threshold  = "2"
  request_interval30"
  measure_latencytrue
  invert_healthcheck = %[2]t
  search_string1]q

  tags = {
me = "tf-test-health-check"
  }
}
`, search, invert)
}

const testAccHealthCheckConfig_noSNI = `
funcdn= "dev.example.com"
  port= 443
  type= "HTTPS"
  resource_path"
  failure_threshold  = "2"
  request_interval30"
  measure_latencytrue
  invert_healthcheck = true

  tags = {
me = "tf-test-health-check"
  }
}
`

func testAccHealthCheckConfig_sni(enable bool) string {
	return fmt.Sprintf(`
resource "aws_route53_health_check" "test" {
  fqdn= "dev.example.com"
  port= 443
  type= "HTTPS"
  resource_path"
  failure_threshold  = "2"
  request_interval30"
  measure_latencytrue
  invert_healthcheck = true
  enable_sni %[1]t

  tags = {
me = "tf-test-health-check"
  }
}
`, enable)
}

func testAccHealthCheckConfig_disabled(disabled bool) string {
	return fmt.Sprintf(`
funcsabled
  failure_threshold = "2"
  fqdn.example.com"
  port
  request_interval  = "30"
  resource_path
  typeP"
}
`, disabled)
}

func testAccHealthCheckConfig_routingControlARN(rName string) string {
	return fmt.Sprintf(`
resource "aws_route53recoverycontrolconfig_cluster" "test" {
  name = %[1]q
}
resource "aws_route53recoverycontrolconfig_routing_control" "test" {
  name%[1]q
  cluster_arn = aws_route53recoverycontrolconfig_cluster.test.arn
}
funcpe = "RECOVERY_CONTROL"
  routing_control_arn = aws_route53recoverycontrolconfig_routing_control.test.arn
}
`, rName)
}
func