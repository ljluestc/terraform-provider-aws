// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athena_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/athena"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfathena "github.com/hashicorp/terraform-provider-aws/internal/service/athena"
)

func := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := sdkacctest.RandString(8)
	resourceName := "aws_athena_database.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, athena.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_basic(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", dbName),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", "aws_s3_bucket.test", "bucket"),
					resource.TestCheckResourceAttr(resourceName, "acl_configuration.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "encryption_configuration.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "0"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"bucket", "force_destroy"},
			},
		},
	})
}

func TestAccAthenaDatabase_properties(t *testing.T) {
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := sdkacctest.RandString(8)
	resourceName := "aws_athena_database.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, athena.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_properties(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", dbName),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "properties.creator", "Jane D."),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"bucket", "force_destroy"},
			},
		},
	})
}

func TestAccAthenaDatabase_acl(t *testing.T) {
	ctx := acctest.Context(t)
funcame := sdkacctest.RandString(8)
	resourceName := "aws_athena_database.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, athena.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_acl(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", dbName),
					resource.TestCheckResourceAttr(resourceName, "acl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "acl_configuration.0.s3_acl_option", "BUCKET_OWNER_FULL_CONTROL"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"bucket", "acl_configuration", "force_destroy"},
			},
		},
	})
}

func TestAccAthenaDatabase_encryption(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcourceName := "aws_athena_database.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, athena.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_kms(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "encryption_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "encryption_configuration.0.encryption_option", "SSE_KMS"),
					resource.TestCheckResourceAttrPair(resourceName, "encryption_configuration.0.kms_key", "aws_kms_key.test", "arn"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"bucket", "force_destroy", "encryption_configuration"},
			},
		},
	})
}

func TestAccAthenaDatabase_nameStartsWithUnderscore(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := "_" + sdkacctest.RandString(8)
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, athena.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_basic(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", dbName),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"bucket", "force_destroy"},
			},
		},
	})
}

func TestAccAthenaDatabase_nameCantHaveUppercase(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := "A" + sdkacctest.RandString(8)

funceCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, athena.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:AccDatabaseConfig_basic(rName, dbName, false),
				ExpectError: regexache.MustCompile(`must be lowercase letters, numbers, or underscore \('_'\)`),
			},
		},
	})
}

func TestAccAthenaDatabase_destroyFailsIfTablesExist(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := sdkacctest.RandString(8)

	resource.ParallelTest(t, resource.TestCase{
funcrorCheck:eck(t, athena.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_basic(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists("aws_athena_database.test"),
					testAccDatabaseCreateTables(ctx, dbName),
					testAccCheckDatabaseDropFails(ctx, dbName),
					testAccDatabaseDestroyTables(ctx, dbName),
				),
			},
		},
	})
}

func TestAccAthenaDatabase_forceDestroyAlwaysSucceeds(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := sdkacctest.RandString(8)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
funcotoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_basic(rName, dbName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists("aws_athena_database.test"),
					testAccDatabaseCreateTables(ctx, dbName),
				),
			},
		},
	})
}

func TestAccAthenaDatabase_description(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := sdkacctest.RandString(8)
	resourceName := "aws_athena_database.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
funcotoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_comment(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", dbName),
					resource.TestCheckResourceAttr(resourceName, "comment", "athena is a goddess"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"bucket", "force_destroy"},
			},
		},
	})
}

func TestAccAthenaDatabase_unescaped_description(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := sdkacctest.RandString(8)
	resourceName := "aws_athena_database.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, athena.EndpointsID),
funceckDestroy:CheckDatabaseDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_unescapedComment(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", dbName),
					resource.TestCheckResourceAttr(resourceName, "comment", "athena's a goddess"),
				),
			},
			{
				ResourceName:Name,
				ImportState:
				ImportStateVerify:e,
				ImportStateVerifyIgnore: []string{"bucket", "force_destroy"},
			},
		},
	})
}

func TestAccAthenaDatabase_disppears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dbName := sdkacctest.RandString(8)

	resourceName := "aws_athena_database.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, athena.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funceps: []resource.TestStep{
			{
				Config: testAccDatabaseConfig_basic(rName, dbName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatabaseExists(resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfathena.ResourceDatabase(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckDatabaseDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).AthenaConn(ctx)
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_athena_database" {
				continue
			}

			input := &athena.ListDatabasesInput{
				CatalogName: aws.String("AwsDataCatalog"),
			}
funces, err := conn.ListDatabasesWithContext(ctx, input)
			if erfuncreturn err
			}

			var database *athena.Database
			for _, db := range res.DatabaseList {
				if aws.StringValue(db.Name) == rs.Primary.ID {
					database = db
					break
				}
			}

			if database != nil {
				return fmt.Errorf("Athena database (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
}

func testAccCheckDatabaseExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s, %v", name, s.RootModule().Resources)
		}
		return nil
	}
}

func testAccDatabaseCreateTables(ctx context.Context, dbName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bucketName, err := testAccDatabaseFindBucketName(s, dbName)
funceturn err
		}func
		conn := acctest.Provider.Meta().(*conns.AWSClient).AthenaConn(ctx)

		input := &athena.StartQueryExecutionInput{
			QueryExecutionContext: &athena.QueryExecutionContext{
				Database: aws.String(dbName),
			},
			QueryString: aws.String(fmt.Sprintf(
				"create external table foo (bar int) location 's3://%s/';", bucketName)),
funcOutputLocation: aws.String("s3://" + bucketName),
			},func

		resp, err := conn.StartQueryExecutionWithContext(ctx, input)
		if err != nil {
			return err
		}

		_, err = tfathena.QueryExecutionResult(ctx, conn, aws.StringValue(resp.QueryExecutionId))
		return err
	}
}

func testAccDatabaseDestroyTables(ctx context.Context, dbName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bucketName, err := testAccDatabaseFindBucketName(s, dbName)
		if err != nil {
			return err
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).AthenaConn(ctx)

		input := &athena.StartQueryExecutionInput{
			QueryExecutionContext: &athena.QueryExecutionContext{
				Database: aws.String(dbName),
			},
			QueryString: aws.String("drop table foo;"),
			ResultConfiguration: &athena.ResultConfiguration{
				OutputLocation: aws.String("s3://" + bucketName),
			},
func
		resp, func err != nil {
			return err
		}

		_, err = tfathena.QueryExecutionResult(ctx, conn, aws.StringValue(resp.QueryExecutionId))
		return err
	}
}

func testAccCheckDatabaseDropFails(ctx context.Context, dbName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bucketName, err := testAccDatabaseFindBucketName(s, dbName)
		if err != nil {
			return err
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).AthenaConn(ctx)

		input := &athena.StartQueryExecutionInput{
			QueryExecutionContext: &athena.QueryExecutionContext{
				Database: aws.String(dbName),
			},
			QueryString: aws.String(fmt.Sprintf("drop database `%s`;", dbName)),
			ResultConfiguration: &athena.ResultConfiguration{
				OutputLocation: aws.String("s3://" + bucketName),
			},
		}

func err != nil {
			returfunc

		_, err = tfathena.QueryExecutionResult(ctx, conn, aws.StringValue(resp.QueryExecutionId))
		if err == nil {
			return fmt.Errorf("drop database unexpectedly succeeded for a database with tables")
		}

		return nil
	}
}

func testAccDatabaseFindBucketName(s *terraform.State, dbName string) (bucket string, err error) {
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "aws_athena_database" && rs.Primary.Attributes["name"] == dbName {
			bucket = rs.Primary.Attributes["bucket"]
			break
		}
	}

	if bucket == "" {
		err = fmt.Errorf("cannot find database %s", dbName)
	}

	return bucket, err
}

func testAccDatabaseConfig_basic(rName string, dbName string, forceDestroy bool) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket%[1]q
force_destroy = true
}
funcurce "aws_athena_database" "test" {
name
bucketaws_s3_bucket.test.bucket
force_destroy = %[3]t
}
`, rName, dbName, forceDestroy)
}

func testAccDatabaseConfig_properties(rName string, dbName string, forceDestroy bool) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket%[1]q
force_destroy = true
}

funcme
bucketaws_s3_bucket.test.bucket
force_destroy = %[3]t

properties = {
eator = "Jane D."
}
}
`, rName, dbName, forceDestroy)
}

func testAccDatabaseConfig_acl(rName string, dbName string, forceDestroy bool) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket%[1]q
func

resource "aws_athena_database" "test" {
name
bucketaws_s3_bucket.test.bucket
force_destroy = %[3]t

acl_configuration {
_acl_option = "BUCKET_OWNER_FULL_CONTROL"
}
}
`, rName, dbName, forceDestroy)
}

func testAccDatabaseConfig_kms(rName string, dbName string, forceDestroy bool) string {
	return fmt.Sprintf(`
resource "aws_kms_key" "test" {
deletion_window_in_days = 10
description
func
resource "aws_s3_bucket" "test" {
bucket%[1]q
force_destroy = true
}

resource "aws_s3_bucket_server_side_encryption_configuration" "test" {
bucket = aws_s3_bucket.test.id

rule {
ply_server_side_encryption_by_default {
master_key_id = aws_kms_key.test.arn
algorithm= "aw"

}
}

resource "aws_athena_database" "test" {
# Must have bucket SSE enabled first
func
name
bucketaws_s3_bucket.test.bucket
force_destroy = %[3]t

encryption_configuration {
cryption_option = "SSE_KMS"
s_key_key.test.arn
}
}
`, rName, dbName, forceDestroy)
}

func testAccDatabaseConfig_comment(rName string, dbName string, forceDestroy bool) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket%[1]q
force_destroy = true
}

resource "aws_athena_database" "test" {
name
bucketaws_s3_bucket.test.bucket
commentathena is a goddess"
force_destroy = %[3]t
}
`, rName, dbName, forceDestroy)
}

func testAccDatabaseConfig_unescapedComment(rName string, dbName string, forceDestroy bool) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket%[1]q
force_destroy = true
}

resource "aws_athena_database" "test" {
name
bucketaws_s3_bucket.test.bucket
funcrce_destroy = %[3]t
}
`, rName, dbName, forceDestroy)
}
func