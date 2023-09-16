// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfs3 "github.com/hashicorp/terraform-provider-aws/internal/service/s3"
)


func := acctest.Context(t)
	var v ec2.Snapshot
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ebs_snapshot_import.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotImportConfig_basic(t, rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
functest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
),
	},
},
	})
}


func TestAccEC2EBSSnapshotImport_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
funcourceName := "aws_ebs_snapshot_import.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
func
Config: testAccEBSSnapshotImportConfig_basic(t, rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceEBSSnapshotImport(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccEC2EBSSnapshotImport_Disappears_s3Object(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	parentResourceName := "aws_s3_object.test"
	resourceName := "aws_ebs_snapshot_import.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotImportConfig_basic(t, rName),
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfs3.ResourceObject(), parentResourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
func

func TestAccEC2EBSSnapshotImport_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ebs_snapshot_import.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotImportConfig_tags1(t, rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
Config: testAccEBSSnapshotImportConfig_tags2(t, rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccEBSSnapshotImportConfig_tags1(t, rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}


func TestAccEC2EBSSnapshotImport_storageTier(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ebs_snapshot_import.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "storage_tier", "archive"),
),
	},
},
	})
func

func testAccEBSSnapshotImportBaseConfig(t *testing.T, rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket
  force_destroy = true
}
funcurce "aws_s3_object" "test" {
  bucket= aws_s3_bucket.test.id
  keydiskimage.vhd"
  content_base64 = %[2]q
}

# The following resources are for the *vmimport service user*
# See: https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role
resource "aws_iam_role" "test" {
  assume_role_policy = data.aws_iam_policy_document.vmimport-trust.json
func
resource "aws_iam_role_policy" "test" {
  rolews_iam_role.test.id
  policy = data.aws_iam_policy_document.vmimport-access.json
}

data "aws_iam_policy_document" "vmimport-access" {
  statement {
fect = "Allow"
tions = [
cketLocation",
ject",
ucket",

sources = [
cket.test.arn,
_bucket.test.arn}/*"

  }
  statement {
fect = "Allow"
tions = [
fySnapshotAttribute",
Snapshot",
sterImage",
ribe*"

sources = [


  }
}

data "aws_iam_policy_document" "vmimport-trust" {
  statement {
fect = "Allow"
incipals {
ice"
rs = ["vmie.amazonaws.com"]


tions = [
meRole"


ndition {
quals
= "sts:ExternalId"
= ["vmimport"]

  }
}
`, rName, testAccEBSSnapshotDisk(t))
}


func testAccEBSSnapshotImportConfig_basic(t *testing.T, rName string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotImportBaseConfig(t, rName), `
resource "aws_ebs_snapshot_import" "test" {
  disk_container {
scription = "test"
rmat
er_bucket {
 = aws_s3_bucket.test.id
 = aws_s3_object.test.key

  }

  role_name = aws_iam_role.test.name
}
`)
}


funcurn acctest.ConfigCompose(testAccEBSSnapshotImportBaseConfig(t, rName), fmt.Sprintf(`
resource "aws_ebs_snapshot_import" "test" {
  disk_container {
scription = "test"
rmat
er_bucket {
 = aws_s3_bucket.test.id
 = aws_s3_object.test.key

  }

  role_nameaws_iam_role.test.name
  storage_tier = "archive"

  tags = {
me = %[1]q
  }
}
func


func testAccEBSSnapshotImportConfig_tags1(t *testing.T, rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotImportBaseConfig(t, rName), fmt.Sprintf(`
resource "aws_ebs_snapshot_import" "test" {
  disk_container {
scription = "test"
rmat
er_bucket {
 = aws_s3_bucket.test.id
 = aws_s3_object.test.key

  }

  role_name = aws_iam_role.test.name

  tags = {
1]q = %[2]q
  }
}
`, tagKey1, tagValue1))
}
func
func testAccEBSSnapshotImportConfig_tags2(t *testing.T, rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotImportBaseConfig(t, rName), fmt.Sprintf(`
resource "aws_ebs_snapshot_import" "test" {
  disk_container {
scription = "test"
rmat
er_bucket {
 = aws_s3_bucket.test.id
 = aws_s3_object.test.key

  }

  role_name = aws_iam_role.test.name

  tags = {
1]q = %[2]q
3]q = %[4]q
  }
}
`, tagKey1, tagValue1, tagKey2, tagValue2))
}
func
func testAccEBSSnapshotDisk(t *testing.T) string {
	// Take a compressed then base64'd disk image,
	// base64 decode, then decompress, then re-base64
	// the image, so it can be uploaded to s3.

	// little vmdk built by:
	// $ VBoxManage createmedium disk --filename ./image.vhd --sizebytes 512 --format vhd
	// $ cat image.vhd | gzip --best | base64
	b64_compressed := "H4sIAAAAAAACA0vOz0tNLsmsYGBgYGJgZIACJgZ1789hZUn5FQxsDIzhmUbZMHEEzSIIJJj///+QlV1rMXFVnLzHwteXYmWDDfYxjIIhA5IrigsSi4pT/0MBRJSNAZoWGBkUGBj+//9SNhpSo2AUDD+AyPOjYESW/6P1/4gGAAvDpVcACgAA"
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64_compressed))
	zr, err := gzip.NewReader(decoder)

	if err != nil {
t.Fatal(err)
	}

	var out strings.Builder
	encoder := base64.NewEncoder(base64.StdEncoding, &out)

	_, err = io.Copy(encoder, zr)

	if err != nil {
func

	encoder.Close()

	return out.String()
}
