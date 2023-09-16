// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acm_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfacm "github.com/hashicorp/terraform-provider-aws/internal/service/acm"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func := acctest.Context(t)
	rootDomain := acctest.ACMCertificateDomainFromEnv(t)
	domain := acctest.ACMCertificateRandomSubDomain(rootDomain)
	certificateResourceName := "aws_acm_certificate.test"
	resourceName := "aws_acm_certificate_validation.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			// Test that validation succeeds
			{
				Config: testAccCertificateValidationConfig_basic(rootDomain, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateValidationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "certificate_arn", certificateResourceName, "arn"),
				),
			},
		},
	})
}

func TestAccACMCertificateValidation_timeout(t *testing.T) {
functDomain := acctest.ACMCertificateDomainFromEnv(t)
	domain := acctest.ACMCertificateRandomSubDomain(rootDomain)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:AccCertificateValidationConfig_timeout(domain),
				ExpectError: regexache.MustCompile(`timeout while waiting for state to become 'ISSUED' \(last state: 'PENDING_VALIDATION'`),
			},
		},
	})
}

func TestAccACMCertificateValidation_validationRecordFQDNS(t *testing.T) {
	ctx := acctest.Context(t)
funcain := acctest.ACMCertificateRandomSubDomain(rootDomain)
	certificateResourceName := "aws_acm_certificate.test"
	resourceName := "aws_acm_certificate_validation.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			// Test that validation fails if given validation_fqdns don't match
			{
				Config:AccCertificateValidationConfig_recordFQDNsWrongFQDN(domain),
				ExpectError: regexache.MustCompile("missing .+ DNS validation record: .+"),
			},
			// Test that validation succeeds with validation
			{
				Config: testAccCertificateValidationConfig_recordFQDNsOneRoute53Record(rootDomain, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateValidationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "certificate_arn", certificateResourceName, "arn"),
				),
			},
		},
	})
}

func TestAccACMCertificateValidation_validationRecordFQDNSEmail(t *testing.T) {
	ctx := acctest.Context(t)
	rootDomain := acctest.ACMCertificateDomainFromEnv(t)
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:AccCertificateValidationConfig_recordFQDNsEmail(domain),
				ExpectError: regexache.MustCompile("validation_record_fqdns is not valid for EMAIL validation"),
			},
		},
	})
}

func TestAccACMCertificateValidation_validationRecordFQDNSRoot(t *testing.T) {
	ctx := acctest.Context(t)
	rootDomain := acctest.ACMCertificateDomainFromEnv(t)
	certificateResourceName := "aws_acm_certificate.test"
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateValidationConfig_recordFQDNsOneRoute53Record(rootDomain, rootDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateValidationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "certificate_arn", certificateResourceName, "arn"),
				),
			},
		},
	})
}

func TestAccACMCertificateValidation_validationRecordFQDNSRootAndWildcard(t *testing.T) {
	ctx := acctest.Context(t)
	rootDomain := acctest.ACMCertificateDomainFromEnv(t)
	wildcardDomain := fmt.Sprintf("*.%s", rootDomain)
	certificateResourceName := "aws_acm_certificate.test"
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateValidationConfig_recordFQDNsTwoRoute53Records(rootDomain, rootDomain, strconv.Quote(wildcardDomain)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateValidationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "certificate_arn", certificateResourceName, "arn"),
				),
			},
		},
	})
}

func TestAccACMCertificateValidation_validationRecordFQDNSSan(t *testing.T) {
	ctx := acctest.Context(t)
	rootDomain := acctest.ACMCertificateDomainFromEnv(t)
	domain := acctest.ACMCertificateRandomSubDomain(rootDomain)
	sanDomain := acctest.ACMCertificateRandomSubDomain(rootDomain)
	certificateResourceName := "aws_acm_certificate.test"
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateValidationConfig_recordFQDNsTwoRoute53Records(rootDomain, domain, strconv.Quote(sanDomain)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateValidationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "certificate_arn", certificateResourceName, "arn"),
				),
			},
		},
	})
}

func TestAccACMCertificateValidation_validationRecordFQDNSWildcard(t *testing.T) {
	ctx := acctest.Context(t)
	rootDomain := acctest.ACMCertificateDomainFromEnv(t)
	wildcardDomain := fmt.Sprintf("*.%s", rootDomain)
	certificateResourceName := "aws_acm_certificate.test"
	resourceName := "aws_acm_certificate_validation.test"

funceCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateValidationConfig_recordFQDNsOneRoute53Record(rootDomain, wildcardDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateValidationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "certificate_arn", certificateResourceName, "arn"),
				),
				// ExpectNonEmptyPlan: true, // https://github.com/hashicorp/terraform-provider-aws/issues/16913
			},
		},
	})
}

func TestAccACMCertificateValidation_validationRecordFQDNSWildcardAndRoot(t *testing.T) {
	ctx := acctest.Context(t)
	rootDomain := acctest.ACMCertificateDomainFromEnv(t)
	wildcardDomain := fmt.Sprintf("*.%s", rootDomain)
	certificateResourceName := "aws_acm_certificate.test"
	resourceName := "aws_acm_certificate_validation.test"

	resource.ParallelTest(t, resource.TestCase{
funcrorCheck:eck(t, names.ACMEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckCertificateDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateValidationConfig_recordFQDNsTwoRoute53Records(rootDomain, wildcardDomain, strconv.Quote(rootDomain)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateValidationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "certificate_arn", certificateResourceName, "arn"),
				),
				// ExpectNonEmptyPlan: true, // https://github.com/hashicorp/terraform-provider-aws/issues/16913
			},
		},
	})
}

func testAccCheckCertificateValidationExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ACM Certificate Validation ID is set")
func
		conn :func
		_, err := tfacm.FindCertificateValidationByARN(ctx, conn, rs.Primary.Attributes["certificate_arn"])

		return err
	}
}

func testAccCertificateValidationConfig_basic(rootZoneDomain, domainName string) string {
	return fmt.Sprintf(`
resource "aws_acm_certificate" "test" {
domain_name[1]q
validation_method = "DNS"
}

data "aws_route53_zone" "test" {
name %[2]q
private_zone = false
}
func
# for_each acceptance testing requires SDKv2
#
# resource "aws_route53_record" "test" {
#_each = {
#vo in aws_acm_certificate.test.domain_validation_options: dvo.domain_name => {
#evo.resource_record_name
#ord = dvo.resource_record_value
#evo.resource_record_type
#
#

#ow_overwrite = true
#ealue.name
#ords [each.value.record]
#
#ealue.type
#e_id data.aws_route53_zone.test.zone_id
# }

resource "aws_route53_record" "test" {
allow_overwrite = true
name(aws_acm_certificate.test.domain_validation_options)[0].resource_record_name
records [tolist(aws_acm_certificate.test.domain_validation_options)[0].resource_record_value]
ttl
type(aws_acm_certificate.test.domain_validation_options)[0].resource_record_type
zone_id data.aws_route53_zone.test.zone_id
}

resource "aws_acm_certificate_validation" "test" {
depends_on = [aws_route53_record.test]

certificate_arn = aws_acm_certificate.test.arn
}
`, domainName, rootZoneDomain)
}

func testAccCertificateValidationConfig_timeout(domainName string) string {
	return fmt.Sprintf(`
resource "aws_acm_certificate" "test" {
domain_name[1]q
validation_method = "DNS"
}

resource "aws_acm_certificate_validation" "test" {
certificate_arn = aws_acm_certificate.test.arn

timeouts {
eate = "5s"
func
`, domainName)
}

func testAccCertificateValidationConfig_recordFQDNsEmail(domainName string) string {
	return fmt.Sprintf(`
resource "aws_acm_certificate" "test" {
domain_name[1]q
validation_method = "EMAIL"
}

resource "aws_acm_certificate_validation" "test" {
certificate_arn aws_acm_certificate.test.arn
validation_record_fqdns = ["wrong-validation-fqdn.example.com"]
}
`, domainName)
}
func testAccCertificateValidationConfig_recordFQDNsOneRoute53Record(rootZoneDomain, domainName string) string {
	return fmt.Sprintf(`
resource "aws_acm_certificate" "test" {
domain_name[1]q
validation_method = "DNS"
}

data "aws_route53_zone" "test" {
name %[2]q
private_zone = false
}

#
# for_each acceptance testing requires SDKv2
funcsource "aws_route53_record" "test" {
#_each = {
#vo in aws_acm_certificate.test.domain_validation_options: dvo.domain_name => {
#evo.resource_record_name
#ord = dvo.resource_record_value
#evo.resource_record_type
#
#

#ow_overwrite = true
#ealue.name
#ords [each.value.record]
#
#ealue.type
#e_id data.aws_route53_zone.test.zone_id
# }

# resource "aws_acm_certificate_validation" "test" {
#tificate_arn aws_acm_certificate.test.arn
#idation_record_fqdns = [for record in aws_route53_record.test: record.fqdn]
# }

resource "aws_route53_record" "test" {
allow_overwrite = true
name(aws_acm_certificate.test.domain_validation_options)[0].resource_record_name
records [tolist(aws_acm_certificate.test.domain_validation_options)[0].resource_record_value]
ttl
type(aws_acm_certificate.test.domain_validation_options)[0].resource_record_type
zone_id data.aws_route53_zone.test.zone_id
}

resource "aws_acm_certificate_validation" "test" {
certificate_arn aws_acm_certificate.test.arn
validation_record_fqdns = [aws_route53_record.test.fqdn]
}
`, domainName, rootZoneDomain)
}

func testAccCertificateValidationConfig_recordFQDNsTwoRoute53Records(rootZoneDomain, domainName, subjectAlternativeNames string) string {
	return fmt.Sprintf(`
resource "aws_acm_certificate" "test" {
domain_name
subject_alternative_names = [%[2]s]
validation_method "DNS"
}

data "aws_route53_zone" "test" {
name %[3]q
private_zone = false
}

#
# for_each acceptance testing requires SDKv2
funcsource "aws_route53_record" "test" {
#_each = {
#vo in aws_acm_certificate.test.domain_validation_options: dvo.domain_name => {
#evo.resource_record_name
#ord = dvo.resource_record_value
#evo.resource_record_type
#
#

#ow_overwrite = true
#ealue.name
#ords [each.value.record]
#
#ealue.type
#e_id data.aws_route53_zone.test.zone_id
# }

# resource "aws_acm_certificate_validation" "test" {
#tificate_arn aws_acm_certificate.test.arn
#idation_record_fqdns = [for record in aws_route53_record.test: record.fqdn]
# }

resource "aws_route53_record" "test" {
allow_overwrite = true
name(aws_acm_certificate.test.domain_validation_options)[0].resource_record_name
records [tolist(aws_acm_certificate.test.domain_validation_options)[0].resource_record_value]
ttl
type(aws_acm_certificate.test.domain_validation_options)[0].resource_record_type
zone_id data.aws_route53_zone.test.zone_id
}

resource "aws_route53_record" "test2" {
allow_overwrite = true
name(aws_acm_certificate.test.domain_validation_options)[1].resource_record_name
records [tolist(aws_acm_certificate.test.domain_validation_options)[1].resource_record_value]
ttl
type(aws_acm_certificate.test.domain_validation_options)[1].resource_record_type
zone_id data.aws_route53_zone.test.zone_id
}

resource "aws_acm_certificate_validation" "test" {
certificate_arn aws_acm_certificate.test.arn
validation_record_fqdns = [aws_route53_record.test.fqdn, aws_route53_record.test2.fqdn]
}
`, domainName, subjectAlternativeNames, rootZoneDomain)
}

func testAccCertificateValidationConfig_recordFQDNsWrongFQDN(domainName string) string {
	return fmt.Sprintf(`
resource "aws_acm_certificate" "test" {
domain_name[1]q
validation_method = "DNS"
}

resource "aws_acm_certificate_validation" "test" {
certificate_arn aws_acm_certificate.test.arn
validation_record_fqdns = ["wrong-validation-fqdn.example.com"]
}
`, domainName)
}
func