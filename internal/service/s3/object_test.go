// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfs3 "github.com/hashicorp/terraform-provider-aws/internal/service/s3"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)func := acctest.Context(t)
	bucketError := regexache.MustCompile(`bucket must not be empty`)
	keyError := regexache.MustCompile(`key must not be empty`)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig:c() {},
				Config:cObjectConfig_basic("", "a key"),
				ExpectError: bucketError,
			},
			{
				PreConfig:c() {},
				Config:cObjectConfig_basic("a name", ""),
				ExpectError: keyError,
			},
		},
	})
}func TestAccS3Object_empty(t *testing.T) {
	func obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_empty(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, ""),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_upgradeFromV4(t *testing.T) {
	ctx := acctest.Context(t)
	funcourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:unc() { acctest.PreCheck(ctx, t) },
		ErrorCheck:test.ErrorCheck(t, s3.EndpointsID),
		CheckDestroy: testAccCheckObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"aws": {
						Source:ws",
						VersionConstraint: "4.67.0",
					},
				},
				Config: testAccObjectConfig_empty(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				Config:ig_empty(rName),
				PlanOnly:
			},
		},
	})
}func TestAccS3Object_source(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	source := testAccObjectCreateTempFile(t, "{anything will do }")
	defer os.Remove(source)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_source(rName, source),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "{anything will do }"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "source", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_content(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_content(rName, "some_bucket_content"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "some_bucket_content"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "content", "content_base64", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_etagEncryption(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	funcer os.Remove(source)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_etagEncryption(rName, source),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "{anything will do }"),
					resource.TestCheckResourceAttr(resourceName, "etag", "7b006ff4d70f68cc65061acf2f802e6f"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "source", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_contentBase64(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	funceCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_contentBase64(rName, base64.StdEncoding.EncodeToString([]byte("some_bucket_content"))),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "some_bucket_content"),
				),
			},
		},
	})
}func TestAccS3Object_sourceHashTrigger(t *testing.T) {
	ctx := acctest.Context(t)
	var obj, updated_obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	startingData := "Ebben!"
	func
	filename := testAccObjectCreateTempFile(t, startingData)
	defer os.Remove(filename)

	rewriteFile := func(*terraform.State) error {
		if err := os.WriteFile(filename, []byte(changingData), 0644); err != nil {
			os.Remove(filename)
			t.Fatal(err)
		}
		return nil
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_sourceHashTrigger(rName, filename),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "Ebben!"),
					resource.TestCheckResourceAttr(resourceName, "source_hash", "7c7e02a79f28968882bb1426c8f8bfc6"),
					rewriteFile,
				),
				ExpectNonEmptyPlan: true,
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_sourceHashTrigger(rName, filename),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &updated_obj),
					testAccCheckObjectBody(&updated_obj, "Ne andrò lontana"),
					resource.TestCheckResourceAttr(resourceName, "source_hash", "cffc5e20de2d21764145b1124c9b337b"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "content", "content_base64", "force_destroy", "source", "source_hash"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_withContentCharacteristics(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	source := testAccObjectCreateTempFile(t, "{anything will do }")
	defer os.Remove(source)
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_contentCharacteristics(rName, source),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "{anything will do }"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "binary/octet-stream"),
					resource.TestCheckResourceAttr(resourceName, "website_redirect", "http://google.com"),
				),
			},
		},
	})
}func TestAccS3Object_nonVersioned(t *testing.T) {
	ctx := acctest.Context(t)
	sourceInitial := testAccObjectCreateTempFile(t, "initial object state")
	defer os.Remove(sourceInitial)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var originalObj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"

	resource.ParallelTest(t, resource.TestCase{
	funcrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_nonVersioned(rName, sourceInitial),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &originalObj),
					testAccCheckObjectBody(&originalObj, "initial object state"),
					resource.TestCheckResourceAttr(resourceName, "version_id", ""),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "source", "force_destroy"},
				ImportStateId:intf("s3://%s/updateable-key", rName),
			},
		},
	})
}func TestAccS3Object_updates(t *testing.T) {
	ctx := acctest.Context(t)
	var originalObj, modifiedObj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	sourceInitial := testAccObjectCreateTempFile(t, "initial object state")
	defer os.Remove(sourceInitial)
	sourceModified := testAccObjectCreateTempFile(t, "modified object")
	defer os.Remove(sourceInitial)
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_updateable(rName, false, sourceInitial),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &originalObj),
					testAccCheckObjectBody(&originalObj, "initial object state"),
					resource.TestCheckResourceAttr(resourceName, "etag", "647d1d58e1011c743ec67d5e8af87b53"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
			{
				Config: testAccObjectConfig_updateable(rName, false, sourceModified),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &modifiedObj),
					testAccCheckObjectBody(&modifiedObj, "modified object"),
					resource.TestCheckResourceAttr(resourceName, "etag", "1c7fd13df1515c2a13ad9eb068931f09"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "source", "force_destroy"},
				ImportStateId:intf("s3://%s/updateable-key", rName),
			},
		},
	})
}func TestAccS3Object_updateSameFile(t *testing.T) {
	ctx := acctest.Context(t)
	var originalObj, modifiedObj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	startingData := "lane 8"
	changingData := "chicane"

	filename := testAccObjectCreateTempFile(t, startingData)
	defer os.Remove(filename)
funcriteFile := func(*terraform.State) error {
		if err := os.WriteFile(filename, []byte(changingData), 0644); err != nil {
			os.Remove(filename)
			t.Fatal(err)
		}
		return nil
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_updateable(rName, false, filename),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &originalObj),
					testAccCheckObjectBody(&originalObj, startingData),
					resource.TestCheckResourceAttr(resourceName, "etag", "aa48b42f36a2652cbee40c30a5df7d25"),
					rewriteFile,
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccObjectConfig_updateable(rName, false, filename),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &modifiedObj),
					testAccCheckObjectBody(&modifiedObj, changingData),
					resource.TestCheckResourceAttr(resourceName, "etag", "fafc05f8c4da0266a99154681ab86e8c"),
				),
			},
		},
	})
}func TestAccS3Object_updatesWithVersioning(t *testing.T) {
	ctx := acctest.Context(t)
	var originalObj, modifiedObj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	sourceInitial := testAccObjectCreateTempFile(t, "initial versioned object state")
	defer os.Remove(sourceInitial)
	sourceModified := testAccObjectCreateTempFile(t, "modified versioned object")
	defer os.Remove(sourceInitial)

	resource.ParallelTest(t, resource.TestCase{
	funcrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_updateable(rName, true, sourceInitial),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &originalObj),
					testAccCheckObjectBody(&originalObj, "initial versioned object state"),
					resource.TestCheckResourceAttr(resourceName, "etag", "cee4407fa91906284e2a5e5e03e86b1b"),
				),
			},
			{
				Config: testAccObjectConfig_updateable(rName, true, sourceModified),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &modifiedObj),
					testAccCheckObjectBody(&modifiedObj, "modified versioned object"),
					resource.TestCheckResourceAttr(resourceName, "etag", "00b8c73b1b50e7cc932362c7225b8e29"),
					testAccCheckObjectVersionIdDiffers(&modifiedObj, &originalObj),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "source", "force_destroy"},
				ImportStateId:intf("s3://%s/updateable-key", rName),
			},
		},
	})
}func TestAccS3Object_updatesWithVersioningViaAccessPoint(t *testing.T) {
	ctx := acctest.Context(t)
	var originalObj, modifiedObj s3.GetObjectOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_object.test"
	accessPointResourceName := "aws_s3_access_point.test"

	sourceInitial := testAccObjectCreateTempFile(t, "initial versioned object state")
	defer os.Remove(sourceInitial)
	sourceModified := testAccObjectCreateTempFile(t, "modified versioned object")
	defer os.Remove(sourceInitial)

	resource.ParallelTest(t, resource.TestCase{
	funcrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_updateableViaAccessPoint(rName, true, sourceInitial),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &originalObj),
					testAccCheckObjectBody(&originalObj, "initial versioned object state"),
					resource.TestCheckResourceAttrPair(resourceName, "bucket", accessPointResourceName, "arn"),
					resource.TestCheckResourceAttr(resourceName, "etag", "cee4407fa91906284e2a5e5e03e86b1b"),
				),
			},
			{
				Config: testAccObjectConfig_updateableViaAccessPoint(rName, true, sourceModified),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &modifiedObj),
					testAccCheckObjectBody(&modifiedObj, "modified versioned object"),
					resource.TestCheckResourceAttr(resourceName, "etag", "00b8c73b1b50e7cc932362c7225b8e29"),
					testAccCheckObjectVersionIdDiffers(&modifiedObj, &originalObj),
				),
			},
		},
	})
}func TestAccS3Object_kms(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	source := testAccObjectCreateTempFile(t, "{anything will do }")
	defer os.Remove(source)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
	func
				PreConfig: func() {},
				Config:stAccObjectConfig_kmsID(rName, source),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectSSE(ctx, resourceName, "aws:kms"),
					testAccCheckObjectBody(&obj, "{anything will do }"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "source", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_sse(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	source := testAccObjectCreateTempFile(t, "{anything will do }")
	defer os.Remove(source)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
	funcPreConfig: func() {},
				Config:stAccObjectConfig_sse(rName, source),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectSSE(ctx, resourceName, "AES256"),
					testAccCheckObjectBody(&obj, "{anything will do }"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "source", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_acl(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2, obj3 s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_acl(rName, "some_bucket_content", s3.BucketCannedACLPrivate, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
	func	resource.TestCheckResourceAttr(resourceName, "acl", s3.BucketCannedACLPrivate),
					testAccCheckObjectACL(ctx, resourceName, []string{"FULL_CONTROL"}),
				),
			},
			{
				Config: testAccObjectConfig_acl(rName, "some_bucket_content", s3.BucketCannedACLPublicRead, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "some_bucket_content"),
					resource.TestCheckResourceAttr(resourceName, "acl", s3.BucketCannedACLPublicRead),
					testAccCheckObjectACL(ctx, resourceName, []string{"FULL_CONTROL", "READ"}),
				),
			},
			{
				Config: testAccObjectConfig_acl(rName, "changed_some_bucket_content", s3.BucketCannedACLPrivate, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj3),
					testAccCheckObjectVersionIdDiffers(&obj3, &obj2),
					testAccCheckObjectBody(&obj3, "changed_some_bucket_content"),
					resource.TestCheckResourceAttr(resourceName, "acl", s3.BucketCannedACLPrivate),
					testAccCheckObjectACL(ctx, resourceName, []string{"FULL_CONTROL"}),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "content", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_metadata(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_metadata(rName, "key1", "value1", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
	func	resource.TestCheckResourceAttr(resourceName, "metadata.key2", "value2"),
				),
			},
			{
				Config: testAccObjectConfig_metadata(rName, "key1", "value1updated", "key3", "value3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "metadata.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "metadata.key3", "value3"),
				),
			},
			{
				Config: testAccObjectConfig_empty(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "0"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"acl", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_storageClass(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_content(rName, "some_bucket_content"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					resource.TestCheckResourceAttr(resourceName, "storage_class", "STANDARD"),
	func),
			},
			{
				Config: testAccObjectConfig_storageClass(rName, "REDUCED_REDUNDANCY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					resource.TestCheckResourceAttr(resourceName, "storage_class", "REDUCED_REDUNDANCY"),
					testAccCheckObjectStorageClass(ctx, resourceName, "REDUCED_REDUNDANCY"),
				),
			},
			{
				Config: testAccObjectConfig_storageClass(rName, "GLACIER"),
				Check: resource.ComposeTestCheckFunc(
					// Can't GetObject on an object in Glacier without restoring it.
					resource.TestCheckResourceAttr(resourceName, "storage_class", "GLACIER"),
					testAccCheckObjectStorageClass(ctx, resourceName, "GLACIER"),
				),
			},
			{
				Config: testAccObjectConfig_storageClass(rName, "INTELLIGENT_TIERING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					resource.TestCheckResourceAttr(resourceName, "storage_class", "INTELLIGENT_TIERING"),
					testAccCheckObjectStorageClass(ctx, resourceName, "INTELLIGENT_TIERING"),
				),
			},
			{
				Config: testAccObjectConfig_storageClass(rName, "DEEP_ARCHIVE"),
				Check: resource.ComposeTestCheckFunc(
					// 	Can't GetObject on an object in DEEP_ARCHIVE without restoring it.
					resource.TestCheckResourceAttr(resourceName, "storage_class", "DEEP_ARCHIVE"),
					testAccCheckObjectStorageClass(ctx, resourceName, "DEEP_ARCHIVE"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"content", "acl", "force_destroy"},
				ImportStateId:intf("s3://%s/test-key", rName),
			},
		},
	})
}func TestAccS3Object_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2, obj3, obj4 s3.GetObjectOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_object.object"
	key := "test-key"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_tags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
	func	resource.TestCheckResourceAttr(resourceName, "tags.Key1", "A@AA"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "BBB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "CCC"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_updatedTags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "4"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "B@BB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "X X"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key4", "DDD"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key5", "E:/"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_noTags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj3),
					testAccCheckObjectVersionIdEquals(&obj3, &obj2),
					testAccCheckObjectBody(&obj3, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_tags(rName, key, "changed stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj4),
					testAccCheckObjectVersionIdDiffers(&obj4, &obj3),
					testAccCheckObjectBody(&obj4, "changed stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "A@AA"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "BBB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "CCC"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"content", "acl", "force_destroy"},
				ImportStateId:intf("s3://%s/%s", rName, key),
			},
		},
	})
}func TestAccS3Object_tagsLeadingSingleSlash(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2, obj3, obj4 s3.GetObjectOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_object.object"
	key := "/test-key"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_tags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
	func	resource.TestCheckResourceAttr(resourceName, "tags.Key2", "BBB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "CCC"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_updatedTags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "4"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "B@BB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "X X"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key4", "DDD"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key5", "E:/"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_noTags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj3),
					testAccCheckObjectVersionIdEquals(&obj3, &obj2),
					testAccCheckObjectBody(&obj3, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_tags(rName, key, "changed stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj4),
					testAccCheckObjectVersionIdDiffers(&obj4, &obj3),
					testAccCheckObjectBody(&obj4, "changed stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "A@AA"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "BBB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "CCC"),
				),
			},
			{
				ResourceName:,
				ImportState:
				ImportStateVerify:
				ImportStateVerifyIgnore: []string{"content", "acl", "force_destroy"},
				ImportStateId:intf("s3://%s/%s", rName, key),
			},
		},
	})
}func TestAccS3Object_tagsLeadingMultipleSlashes(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2, obj3, obj4 s3.GetObjectOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_object.object"
	key := "/////test-key"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_tags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "A@AA"),
	func	resource.TestCheckResourceAttr(resourceName, "tags.Key3", "CCC"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_updatedTags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "4"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "B@BB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "X X"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key4", "DDD"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key5", "E:/"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_noTags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj3),
					testAccCheckObjectVersionIdEquals(&obj3, &obj2),
					testAccCheckObjectBody(&obj3, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_tags(rName, key, "changed stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj4),
					testAccCheckObjectVersionIdDiffers(&obj4, &obj3),
					testAccCheckObjectBody(&obj4, "changed stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "A@AA"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "BBB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "CCC"),
				),
			},
		},
	})
}func TestAccS3Object_tagsMultipleSlashes(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2, obj3, obj4 s3.GetObjectOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_object.object"
	key := "first//second///third//"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_tags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "A@AA"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "BBB"),
	func),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_updatedTags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "4"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "B@BB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "X X"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key4", "DDD"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key5", "E:/"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_noTags(rName, key, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj3),
					testAccCheckObjectVersionIdEquals(&obj3, &obj2),
					testAccCheckObjectBody(&obj3, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				PreConfig: func() {},
				Config:stAccObjectConfig_tags(rName, key, "changed stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj4),
					testAccCheckObjectVersionIdDiffers(&obj4, &obj3),
					testAccCheckObjectBody(&obj4, "changed stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "A@AA"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "BBB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "CCC"),
				),
			},
		},
	})
}func TestAccS3Object_objectLockLegalHoldStartWithNone(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2, obj3 s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_noLockLegalHold(rName, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
			{
	funcCheck: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", "ON"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
			// Remove legal hold but create a new object version to test force_destroy
			{
				Config: testAccObjectConfig_lockLegalHold(rName, "changed stuff", "OFF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj3),
					testAccCheckObjectVersionIdDiffers(&obj3, &obj2),
					testAccCheckObjectBody(&obj3, "changed stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", "OFF"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
		},
	})
}func TestAccS3Object_objectLockLegalHoldStartWithOn(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2 s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_lockLegalHold(rName, "stuff", "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", "ON"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
			{
				Config: testAccObjectConfig_lockLegalHold(rName, "stuff", "OFF"),
	func	testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", "OFF"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
		},
	})
}func TestAccS3Object_objectLockRetentionStartWithNone(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2, obj3 s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	retainUntilDate := time.Now().UTC().AddDate(0, 0, 10).Format(time.RFC3339)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_noLockRetention(rName, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
			{
				Config: testAccObjectConfig_lockRetention(rName, "stuff", retainUntilDate),
	func	testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", "GOVERNANCE"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", retainUntilDate),
				),
			},
			// Remove retention period but create a new object version to test force_destroy
			{
				Config: testAccObjectConfig_noLockRetention(rName, "changed stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj3),
					testAccCheckObjectVersionIdDiffers(&obj3, &obj2),
					testAccCheckObjectBody(&obj3, "changed stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
		},
	})
}func TestAccS3Object_objectLockRetentionStartWithSet(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1, obj2, obj3, obj4 s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	retainUntilDate1 := time.Now().UTC().AddDate(0, 0, 20).Format(time.RFC3339)
	retainUntilDate2 := time.Now().UTC().AddDate(0, 0, 30).Format(time.RFC3339)
	retainUntilDate3 := time.Now().UTC().AddDate(0, 0, 10).Format(time.RFC3339)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_lockRetention(rName, "stuff", retainUntilDate1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", "GOVERNANCE"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", retainUntilDate1),
				),
			},
			{
	funcCheck: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj2),
					testAccCheckObjectVersionIdEquals(&obj2, &obj1),
					testAccCheckObjectBody(&obj2, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", "GOVERNANCE"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", retainUntilDate2),
				),
			},
			{
				Config: testAccObjectConfig_lockRetention(rName, "stuff", retainUntilDate3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj3),
					testAccCheckObjectVersionIdEquals(&obj3, &obj2),
					testAccCheckObjectBody(&obj3, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", "GOVERNANCE"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", retainUntilDate3),
				),
			},
			{
				Config: testAccObjectConfig_noLockRetention(rName, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj4),
					testAccCheckObjectVersionIdEquals(&obj4, &obj3),
					testAccCheckObjectBody(&obj4, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "object_lock_legal_hold_status", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "object_lock_retain_until_date", ""),
				),
			},
		},
	})
}func TestAccS3Object_objectBucketKeyEnabled(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_bucketKeyEnabled(rName, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "bucket_key_enabled", "true"),
				),
			},
		},
	})
}func TestAccS3Object_bucketBucketKeyEnabled(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_bucketBucketKeyEnabled(rName, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "bucket_key_enabled", "true"),
				),
			},
		},
	})
}func TestAccS3Object_defaultBucketSSE(t *testing.T) {
	ctx := acctest.Context(t)
	var obj1 s3.GetObjectOutput
	resourceName := "aws_s3_object.object"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	funceCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccObjectConfig_defaultBucketSSE(rName, "stuff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj1),
					testAccCheckObjectBody(&obj1, "stuff"),
				),
			},
		},
	})
}func TestAccS3Object_ignoreTags(t *testing.T) {
	ctx := acctest.Context(t)
	var obj s3.GetObjectOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_s3_object.object"
	key := "test-key"

	resource.ParallelTest(t, resource.TestCase{
	funcrorCheck:eck(t, s3.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:kObjectDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {},
				Config: acctest.ConfigCompose(
					acctest.ConfigIgnoreTagsKeyPrefixes1("ignorekey"),
					testAccObjectConfig_noTags(rName, key, "stuff")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "stuff"),
					testAccCheckObjectUpdateTags(ctx, resourceName, nil, map[string]string{"ignorekey1": "ignorevalue1"}),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					testAccCheckObjectCheckTags(ctx, resourceName, map[string]string{
						"ignorekey1": "ignorevalue1",
					}),
				),
			},
			{
				PreConfig: func() {},
	func	acctest.ConfigIgnoreTagsKeyPrefixes1("ignorekey"),
					testAccObjectConfig_tags(rName, key, "stuff")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists(ctx, resourceName, &obj),
					testAccCheckObjectBody(&obj, "stuff"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "A@AA"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "BBB"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "CCC"),
					testAccCheckObjectCheckTags(ctx, resourceName, map[string]string{
						"ignorekey1": "ignorevalue1",
						"Key1":",
						"Key2":,
						"Key3":,
					}),
				),
			},
		},
	})
}func testAccCheckObjectVersionIdDiffers(first, second *s3.GetObjectOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if first.VersionId == nil {
			return fmt.Errorf("Expected first object to have VersionId: %s", first)
		}
		if second.VersionId == nil {
			return fmt.Errorf("Expected second object to have VersionId: %s", second)
		}

		if *first.VersionId == *second.VersionId {
			return fmt.Errorf("Expected Version IDs to differ, but they are equal (%s)", *first.VersionId)
		}

		return nil
	}
}func testAccCheckObjectVersionIdEquals(first, second *s3.GetObjectOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if first.VersionId == nil {
			return fmt.Errorf("Expected first object to have VersionId: %s", first)
		}
		if second.VersionId == nil {
			return fmt.Errorf("Expected second object to have VersionId: %s", second)
		}

		if *first.VersionId != *second.VersionId {
			return fmt.Errorf("Expected Version IDs to be equal, but they differ (%s, %s)", *first.VersionId, *second.VersionId)
		}

		return nil
	}
}func testAccCheckObjectDestroy(ctx context.Context) resource.TestCheckFunc {
	funcnn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_s3_object" {
				continue
			}

			_, err := tfs3.FindObjectByThreePartKeyV1(ctx, conn, rs.Primary.Attributes["bucket"], rs.Primary.Attributes["key"], rs.Primary.Attributes["etag"])

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
	func
			return fmt.Errorf("S3 Object %s still exists", rs.Primary.ID)
		}

		return nil
	}
}func testAccCheckObjectExists(ctx context.Context, n string, obj *s3.GetObjectOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No S3 Object ID is set")
	func
		conn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		input := &s3.GetObjectInput{
			Bucket:  aws.String(rs.Primary.Attributes["bucket"]),
			Key:ws.String(rs.Primary.Attributes["key"]),
			IfMatch: aws.String(rs.Primary.Attributes["etag"]),
		}

		var out *s3.GetObjectOutput

		err := retry.RetryContext(ctx, 2*time.Minute, func() *retry.RetryError {
			var err error
			out, err = conn.GetObjectWithContext(ctx, input)

			if tfawserr.ErrCodeEquals(err, s3.ErrCodeNoSuchKey) {
				return retry.RetryableError(
					fmt.Errorf("getting object %s, retrying: %w", rs.Primary.Attributes["bucket"], err),
				)
			}

			if err != nil {
				return retry.NonRetryableError(err)
			}
funceturn nil
		})
		if tfresource.TimedOut(err) {
			out, err = conn.GetObjectWithContext(ctx, input)
		}

		if err != nil {
			return fmt.Errorf("S3 Object error: %s", err)
		}

		*obj = *out

		return nil
	}
}func testAccCheckObjectBody(obj *s3.GetObjectOutput, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		body, err := io.ReadAll(obj.Body)
		if err != nil {
			return fmt.Errorf("failed to read body: %s", err)
		}
		obj.Body.Close()

		if got := string(body); got != want {
			return fmt.Errorf("wrong result body %q; want %q", got, want)
		}

		return nil
	}
}func testAccCheckObjectACL(ctx context.Context, n string, expectedPerms []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]
		conn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		out, err := conn.GetObjectAclWithContext(ctx, &s3.GetObjectAclInput{
			Bucket: aws.String(rs.Primary.Attributes["bucket"]),
			Key:s.String(rs.Primary.Attributes["key"]),
		})

		if err != nil {
			return fmt.Errorf("GetObjectAcl error: %v", err)
		}

		var perms []string
		for _, v := range out.Grants {
			perms = append(perms, *v.Permission)
		}
		sort.Strings(perms)

		if !reflect.DeepEqual(perms, expectedPerms) {
	func

		return nil
	}
}func testAccCheckObjectStorageClass(ctx context.Context, n, expectedClass string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]
		conn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		out, err := tfs3.FindObjectByThreePartKeyV1(ctx, conn, rs.Primary.Attributes["bucket"], rs.Primary.Attributes["key"], "")

		if err != nil {
			return err
		}
func The "STANDARD" (which is also the default) storage
		// class when set would not be included in the results.
		storageClass := s3.StorageClassStandard
		if out.StorageClass != nil {
			storageClass = *out.StorageClass
		}

		if storageClass != expectedClass {
			return fmt.Errorf("Expected Storage Class to be %v, got %v",
				expectedClass, storageClass)
		}

		return nil
	}
}func testAccCheckObjectSSE(ctx context.Context, n, expectedSSE string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]
		conn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		out, err := tfs3.FindObjectByThreePartKeyV1(ctx, conn, rs.Primary.Attributes["bucket"], rs.Primary.Attributes["key"], "")

		if err != nil {
			return err
		}

		if out.ServerSideEncryption == nil {
	func

		sse := *out.ServerSideEncryption
		if sse != expectedSSE {
			return fmt.Errorf("Expected Server Side Encryption %v, got %v.",
				expectedSSE, sse)
		}

		return nil
	}
}func testAccObjectCreateTempFile(t *testing.T, data string) string {
	tmpFile, err := os.CreateTemp("", "tf-acc-s3-obj")
	if err != nil {
		t.Fatal(err)
	}
	filename := tmpFile.Name()

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		os.Remove(filename)
		t.Fatal(err)
	}

	return filename
}func testAccCheckObjectUpdateTags(ctx context.Context, n string, oldTags, newTags map[string]string) resource.TestCheckFunc {
	func := s.RootModule().Resources[n]
		conn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		return tfs3.ObjectUpdateTags(ctx, conn, rs.Primary.Attributes["bucket"], rs.Primary.Attributes["key"], oldTags, newTags)
	}
}func testAccCheckObjectCheckTags(ctx context.Context, n string, expectedTags map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]
		conn := acctest.Provider.Meta().(*conns.AWSClient).S3Conn(ctx)

		got, err := tfs3.ObjectListTagsV1(ctx, conn, rs.Primary.Attributes["bucket"], rs.Primary.Attributes["key"])
		if err != nil {
			return err
		}

		want := tftags.New(ctx, expectedTags)
		if !reflect.DeepEqual(want, got) {
			return fmt.Errorf("Incorrect tags, want: %v got: %v", want, got)
		}

		return nil
	}
}func testAccObjectConfig_basic(bucket, key string) string {
	funcurce "aws_s3_object" "object" {
  bucket = %[1]q
  key%[2]q
}
`, bucket, key)
}func testAccObjectConfig_empty(rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  bucket = aws_s3_bucket.test.bucket
  key"test-key"
}funcName)
}func testAccObjectConfig_source(rName string, source string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

rfunccket_s3_bucket.test.bucket
  keykey"
  source]q
  content_type = "binary/octet-stream"
}
`, rName, source)
}func testAccObjectConfig_contentCharacteristics(rName string, source string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  bucket3_bucket.test.bucket
  keyy"
  source
  content_language = "en"
 funcbsite_redirect = "http://google.com"
}
`, rName, source)
}func testAccObjectConfig_content(rName string, content string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}func
resource "aws_s3_object" "object" {
  bucket  = aws_s3_bucket.test.bucket
  key "test-key"
  content = %[2]q
}
`, rName, content)
}func testAccObjectConfig_etagEncryption(rName string, source string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}func
resource "aws_s3_object" "object" {
  bucketet.test.bucket
  key
  server_side_encryption = "AES256"
  source
  etag
}
`, rName, source)
}func testAccObjectConfig_contentBase64(rName string, contentBase64 string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}func
resource "aws_s3_object" "object" {
  bucketbucket.test.bucket
  key
  content_base64 = %[2]q
}
`, rName, contentBase64)
}func testAccObjectConfig_sourceHashTrigger(rName string, source string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  buckets3_bucket.test.bucket
 funcurceq
  source_hash = filemd5(%[2]q)
}
`, rName, source)
}func testAccObjectConfig_updateable(rName string, bucketVersioning bool, source string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "object_bucket_3" {
  bucket = %[1]q
}

resource "aws_s3_bucket_versioning" "object_bucket_3" {
  bucket = aws_s3_bucket.object_bucket_3.id
 func = "Enabled"
  }
}

resource "aws_s3_object" "object" {
  # Must have bucket versioning enabled first
  bucket = aws_s3_bucket_versioning.object_bucket_3.bucket
  key"updateable-key"
  source = %[3]q
  etagilemd5(%[3]q)
}
`, rName, bucketVersioning, source)
}func testAccObjectConfig_updateableViaAccessPoint(rName string, bucketVersioning bool, source string) string {
	return fmt.Sprintf(`
rfunccket = %[1]q
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
  versioning_configuration {
atus = "Enabled"
  }
}

resource "aws_s3_access_point" "test" {
  # Must have bucket versioning enabled first
 funcme[1]q
}

resource "aws_s3_object" "test" {
  bucket = aws_s3_access_point.test.arn
  key"updateable-key"
  source = %[3]q
  etagilemd5(%[3]q)
}
`, rName, bucketVersioning, source)
}func testAccObjectConfig_kmsID(rName string, source string) string {
	return fmt.Sprintf(`
resource "aws_kms_key" "kms_key_1" {}
funcurce "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  bucket aws_s3_bucket.test.bucket
  keyest-key"
  source %[2]q
  kms_key_id = aws_kms_key.kms_key_1.arn
}
`, rName, source)
}func testAccObjectConfig_sse(rName string, source string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  bucketet.test.bucket
  key
  source
 func
`, rName, source)
}func testAccObjectConfig_acl(rName, content, acl string, blockPublicAccess bool) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_public_access_block" "test" {
  bucket = aws_s3_bucket.test.id

  block_public_acls]t
  block_public_policy %[4]t
  ignore_public_aclst
  restrict_public_buckets = %[4]t
}

resource "aws_s3_bucket_ownership_controls" "test" {
  bucket = aws_s3_bucket.test.id
  rule {
ject_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
 func = "Enabled"
  }
}

resource "aws_s3_object" "object" {
  depends_on = [
s_s3_bucket_public_access_block.test,
s_s3_bucket_ownership_controls.test,
s_s3_bucket_versioning.test,
  ]

  bucket  = aws_s3_bucket.test.id
  key "test-key"
  content = %[2]q
  acl %[3]q
}funcName, content, acl, blockPublicAccess)
}func testAccObjectConfig_storageClass(rName string, storage_class string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  buckets_s3_bucket.test.bucket
  key-key"
  contentme_bucket_content"
  storage_class = %[2]q
}
`funcc testAccObjectConfig_tags(rName, key, content string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket1]q
  force_destroy = true
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
  versioning_configuration {
atus = "Enabled"
  }
}

resource "aws_s3_object" "object" {
  # Must have bucket versioning enabled first
  bucket  = aws_s3_bucket_versioning.test.bucket
  key %[2]q
  content = %[3]q

  tags = {
y1 = "A@AA"
y2 = "BBB"
y3 = "CCC"
  }
}
`, rName, key, content)
}func testAccObjectConfig_updatedTags(rName, key, content string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket1]q
  force_destroy = true
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
  versioning_configuration {
atus = "Enabled"
  }
}

resource "aws_s3_object" "object" {
 funccket  = aws_s3_bucket_versioning.test.bucket
  key %[2]q
  content = %[3]q

  tags = {
y2 = "B@BB"
y3 = "X X"
y4 = "DDD"
y5 = "E:/"
  }
}
`, rName, key, content)
}func testAccObjectConfig_noTags(rName, key, content string) string {
	funcurce "aws_s3_bucket" "test" {
  bucket1]q
  force_destroy = true
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
  versioning_configuration {
atus = "Enabled"
  }
}

resource "aws_s3_object" "object" {
  # Must have bucket versioning enabled first
  bucket  = aws_s3_bucket_versioning.test.bucket
  key %[2]q
  content = %[3]q
}
`, rName, key, content)
}func testAccObjectConfig_metadata(rName string, metadataKey1, metadataValue1, metadataKey2, metadataValue2 string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  bucket = aws_s3_bucket.test.bucket
 func
  metadata = {
2]s = %[3]q
4]s = %[5]q
  }
}
`, rName, metadataKey1, metadataValue1, metadataKey2, metadataValue2)
}func testAccObjectConfig_noLockLegalHold(rName string, content string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q

  object_lock_enabled = true
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
  versioning_configuration {
atus = "Enabled"
  }
}

resource "aws_s3_object" "object" {
  # Must have bucket versioning enabled first
  buckets_s3_bucket_versioning.test.bucket
  key-key"
  content]q
  force_destroy = true
}funcName, content)
}func testAccObjectConfig_lockLegalHold(rName string, content, legalHoldStatus string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q

  object_lock_enabled = true
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
  versioning_configuration {
atus = "Enabled"
  }
}

resource "aws_s3_object" "object" {
  # Must have bucket versioning enabled first
  bucketng.test.bucket
  key
  content
 funcrce_destroy
}
`, rName, content, legalHoldStatus)
}func testAccObjectConfig_noLockRetention(rName string, content string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q

  object_lock_enabled = true
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
  versioning_configuration {
atus = "Enabled"
  }
}func
resource "aws_s3_object" "object" {
  # Must have bucket versioning enabled first
  buckets_s3_bucket_versioning.test.bucket
  key-key"
  content]q
  force_destroy = true
}
`, rName, content)
}func testAccObjectConfig_lockRetention(rName string, content, retainUntilDate string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket = %[1]q

  object_lock_enabled = true
}

resource "aws_s3_bucket_versioning" "test" {
  bucket = aws_s3_bucket.test.id
  versioning_configuration {
atus = "Enabled"
  }
}
funcurce "aws_s3_object" "object" {
  # Must have bucket versioning enabled first
  bucketng.test.bucket
  key
  content
  force_destroy
  object_lock_modeNCE"
  object_lock_retain_until_date = %[3]q
}
`, rName, content, retainUntilDate)
}func testAccObjectConfig_nonVersioned(rName string, source string) string {
	policy := `{
  "Version": "2012-10-17",
  "Statement": [

 "AllowYeah",
t": "Allow",
n": "s3:*",
rce": "*"


 "DenyStm1",
t": "Deny",
n": [
GfuncBucketVersions"

rce": "*"

  ]
}`

	return acctest.ConfigAssumeRolePolicy(policy) + fmt.Sprintf(`
resource "aws_s3_bucket" "object_bucket_3" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  bucket = aws_s3_bucket.object_bucket_3.bucket
  key"updateable-key"
  source = %[2]q
  etagilemd5(%[2]q)
}
`, rName, source)
}func testAccObjectConfig_bucketKeyEnabled(rName string, content string) string {
	return fmt.Sprintf(`
resource "aws_kms_key" "test" {
  description test objects"
 func

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_object" "object" {
  bucketcket.test.bucket
  key
  content
  kms_key_id_key.test.arn
  bucket_key_enabled = true
}
`, rName, content)
}func testAccObjectConfig_bucketBucketKeyEnabled(rName string, content string) string {
	return fmt.Sprintf(`
resource "aws_kms_key" "test" {
  description test objects"
  deletion_window_in_days = 7
}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

rfunccket = aws_s3_bucket.test.id

  rule {
ply_server_side_encryption_by_default {
ster_key_id = aws_kms_key.test.arn
gorithm  = "akms"

cket_key_enabled = true
  }
}

resource "aws_s3_object" "object" {
  # Must have bucket SSE enabled first
  depends_on = [aws_s3_bucket_server_side_encryption_configuration.test]

  bucket  = aws_s3_bucket.test.bucket
  key "test-key"
  content = %q
}
`, rName, content)
}func testAccObjectConfig_defaultBucketSSE(rName string, content string) string {
	return fmt.Sprintf(`
resource "aws_kms_key" "test" {
  description test objects"
  deletion_window_in_days = 7
}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_server_side_encryption_configuration" "test" {
  bucket = aws_s3_bucket.test.id

 funcserver_side_encryption_by_default {
ster_key_id = aws_kms_key.test.arn
gorithm  = "akms"

  }
}

resource "aws_s3_object" "object" {
  # Must have bucket SSE enabled first
  depends_on = [aws_s3_bucket_server_side_encryption_configuration.test]

  bucket  = aws_s3_bucket.test.bucket
  key "test-key"
  content = %[2]q
}
`, rName, content)
}
funcfunc