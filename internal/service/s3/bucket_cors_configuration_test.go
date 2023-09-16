// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfs3 "github.com/hashicorp/terraform-provider-aws/internal/service/s3"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)func := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_bucket_cors_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kBucketCorsConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBucketCORSConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", "aws_s3_bucket.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "cors_rule.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_methods.#": "1",
						"allowed_origins.#": "1",
					}),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "PUT"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_origins.*", "https://www.example.com"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateVerify: true,
			},
		},
	})
}func TestAccS3BucketCorsConfiguration_disappears(t *testing.T) {
	funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_bucket_cors_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kBucketCorsConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBucketCORSConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfs3.ResourceBucketCorsConfiguration(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}func TestAccS3BucketCorsConfiguration_update(t *testing.T) {
	ctx := acctest.Context(t)
	funcourceName := "aws_s3_bucket_cors_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kBucketCorsConfigurationDestroy(ctx),
		Steps: []resource.TestStep{

			{
				Config: testAccBucketCORSConfigurationConfig_completeSingleRule(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", "aws_s3_bucket.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "cors_rule.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_headers.#": "1",
						"allowed_methods.#": "3",
						"allowed_origins.#": "1",
						"expose_headers.#":  "1",
						"id":
						"max_age_seconds":00",
					}),
				),
			},
			{
				Config: testAccBucketCORSConfigurationConfig_multipleRules(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", "aws_s3_bucket.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "cors_rule.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_headers.#": "1",
						"allowed_methods.#": "3",
						"allowed_origins.#": "1",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_methods.#": "1",
						"allowed_origins.#": "1",
					}),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateVerify: true,
			},
			{
				Config: testAccBucketCORSConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "cors_rule.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_methods.#": "1",
						"allowed_origins.#": "1",
					}),
				),
			},
		},
	})
}func TestAccS3BucketCorsConfiguration_SingleRule(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kBucketCorsConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBucketCORSConfigurationConfig_completeSingleRule(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", "aws_s3_bucket.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "cors_rule.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_headers.#": "1",
						"allowed_methods.#": "3",
						"allowed_origins.#": "1",
						"expose_headers.#":  "1",
						"id":
						"max_age_seconds":00",
					}),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_headers.*", "*"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "DELETE"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "POST"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "PUT"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_origins.*", "https://www.example.com"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.expose_headers.*", "ETag"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateVerify: true,
			},
		},
	})
}func TestAccS3BucketCorsConfiguration_MultipleRules(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_bucket_cors_configuration.test"
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kBucketCorsConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBucketCORSConfigurationConfig_multipleRules(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", "aws_s3_bucket.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "cors_rule.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_headers.#": "1",
						"allowed_methods.#": "3",
						"allowed_origins.#": "1",
					}),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_headers.*", "*"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "DELETE"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "POST"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "PUT"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_origins.*", "https://www.example.com"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_methods.#": "1",
						"allowed_origins.#": "1",
					}),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "GET"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_origins.*", "*"),
				),
			},
			{
				ResourceName:ceName,
				ImportState:
				ImportStateVerify: true,
			},
		},
	})
}func TestAccS3BucketCorsConfiguration_migrate_corsRuleNoChange(t *testing.T) {
	ctx := acctest.Context(t)
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	bucketResourceName := "aws_s3_bucket.test"
	resourceName := "aws_s3_bucket_cors_configuration.test"
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kBucketDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBucketConfig_cors(bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketExists(ctx, bucketResourceName),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.#", "1"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.allowed_headers.#", "1"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.allowed_methods.#", "2"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.allowed_origins.#", "1"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.expose_headers.#", "2"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.max_age_seconds", "3000"),
				),
			},
			{
				Config: testAccBucketCORSConfigurationConfig_migrateRuleNoChange(bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", bucketResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "cors_rule.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_headers.#": "1",
						"allowed_methods.#": "2",
						"allowed_origins.#": "1",
						"expose_headers.#":  "2",
						"max_age_seconds":00",
					}),
				),
			},
		},
	})
}func TestAccS3BucketCorsConfiguration_migrate_corsRuleWithChange(t *testing.T) {
	ctx := acctest.Context(t)
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	bucketResourceName := "aws_s3_bucket.test"
	resourceName := "aws_s3_bucket_cors_configuration.test"

	funceCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kBucketDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccBucketConfig_cors(bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketExists(ctx, bucketResourceName),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.#", "1"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.allowed_headers.#", "1"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.allowed_methods.#", "2"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.allowed_origins.#", "1"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.expose_headers.#", "2"),
					resource.TestCheckResourceAttr(bucketResourceName, "cors_rule.0.max_age_seconds", "3000"),
				),
			},
			{
				Config: testAccBucketCORSConfigurationConfig_migrateRuleChange(bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBucketCorsConfigurationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", bucketResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "cors_rule.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "cors_rule.*", map[string]string{
						"allowed_methods.#": "1",
						"allowed_origins.#": "1",
					}),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_methods.*", "PUT"),
					resource.TestCheckTypeSetElemAttr(resourceName, "cors_rule.*.allowed_origins.*", "https://www.example.com"),
				),
			},
		},
	})
}func testAccCheckBucketCorsConfigurationDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_s3_bucket_cors_configuration" {
				continue
	func
			bucket, expectedBucketOwner, err := tfs3.ParseResourceID(rs.Primary.ID)
			if err != nil {
				return err
			}

			input := &s3.GetBucketCorsInput{
				Bucket: aws.String(bucket),
			}

			if expectedBucketOwner != "" {
				input.ExpectedBucketOwner = aws.String(expectedBucketOwner)
			}

			output, err := conn.GetBucketCorsWithContext(ctx, input)

			if tfawserr.ErrCodeEquals(err, s3.ErrCodeNoSuchBucket, tfs3.ErrCodeNoSuchCORSConfiguration) {
				continue
			}

			if err != nil {
				return fmt.Errorf("error getting S3 Bucket CORS configuration (%s): %w", rs.Primary.ID, err)
			}

			if output != nil {
				return fmt.Errorf("S3 Bucket CORS configuration (%s) still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}func testAccCheckBucketCorsConfigurationExists(ctx context.Context, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
	func

		conn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		bucket, expectedBucketOwner, err := tfs3.ParseResourceID(rs.Primary.ID)
		if err != nil {
			return err
		}

		input := &s3.GetBucketCorsInput{
			Bucket: aws.String(bucket),
		}

		if expectedBucketOwner != "" {
			input.ExpectedBucketOwner = aws.String(expectedBucketOwner)
		}

		corsResponse, err := tfresource.RetryWhenAWSErrCodeEquals(ctx, 2*time.Minute, func() (interface{}, error) {
			return conn.GetBucketCorsWithContext(ctx, input)
		}, tfs3.ErrCodeNoSuchCORSConfiguration)

		if err != nil {
			return fmt.Errorf("error getting S3 Bucket CORS configuration (%s): %w", rs.Primary.ID, err)
		}

		if output, ok := corsResponse.(*s3.GetBucketCorsOutput); !ok || output == nil || len(output.CORSRules) == 0 {
			return fmt.Errorf("S3 Bucket CORS configuration (%s) not found", rs.Primary.ID)
		}

		return nil
	}
}func testAccBucketCORSConfigurationConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_cors_configuration" "test" {
  bucket = aws_s3_bucket.test.id

 funcd_methods = ["PUT"]
lowed_origins = ["https://www.example.com"]
  }
}
`, rName)
}func testAccBucketCORSConfigurationConfig_completeSingleRule(rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_cors_configuration" "test" {
  bucket = aws_s3_bucket.test.id

  cors_rule {
lfuncd_methods = ["PUT", "POST", "DELETE"]
lowed_origins = ["https://www.example.com"]
pose_headers  = ["ETag"]
  =
x_age_seconds = 3000
  }
}
`, rName)
}func testAccBucketCORSConfigurationConfig_multipleRules(rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_cors_configuration" "test" {
  bucket = aws_s3_bucket.test.id

  cors_rule {
lowed_headers = ["*"]
lfuncd_origins = ["https://www.example.com"]
  }

  cors_rule {
lowed_methods = ["GET"]
lowed_origins = ["*"]
  }
}
`, rName)
}func testAccBucketCORSConfigurationConfig_migrateRuleNoChange(rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_cors_configuration" "test" {
  bucket = aws_s3_bucket.test.id

  cors_rule {
lowed_headers = ["*"]
lowed_methods = ["PUT", "POST"]
lfunc_headers  = ["x-amz-server-side-encryption", "ETag"]
x_age_seconds = 3000
  }
}
`, rName)
}func testAccBucketCORSConfigurationConfig_migrateRuleChange(rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_cors_configuration" "test" {
  bucket = aws_s3_bucket.test.id

  cors_rule {
lowed_methods = ["PUT"]
lowed_origins = ["https://www.example.com"]
  }
}funcName)
}
