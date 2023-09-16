// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadmin_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ssoadmin"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfssoadmin "github.com/hashicorp/terraform-provider-aws/internal/service/ssoadmin"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckPermissionSetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionSetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
funcource.TestCheckResourceAttr(resourceName, "session_duration", "PT1H"),
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


func TestAccSSOAdminPermissionSet_tags(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set.test"
func
	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckPermissionSetDestroy(ctx),
func
Config: testAccPermissionSetConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
func
ResourceName:ceName,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccPermissionSetConfig_tags2(rName, "key1", "updatedvalue1", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "updatedvalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
func
ResourceName:ceName,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccPermissionSetConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccSSOAdminPermissionSet_updateDescription(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckPermissionSetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionSetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "description", ""),
),
	},
	{
Config: testAccPermissionSetConfig_updateDescription(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
func
	},
	{
ResourceName:ceName,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccSSOAdminPermissionSet_updateRelayState(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckPermissionSetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccPermissionSetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "relay_state", ""),
),
	},
funcig: testAccPermissionSetConfig_updateRelayState(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "relay_state", "https://example.com"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckPermissionSetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionSetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
func
	},
	{
Config: testAccPermissionSetConfig_updateSessionDuration(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
func
	},
	{
ResourceName:ceName,
ImportState:
ImportStateVerify: true,
	},
},
func

// TestAccSSOAdminPermissionSet_RelayState_updateSessionDuration validates
// the resource's unchanged values (primarily relay_state) after updating the session_duration argument
// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/17411

func TestAccSSOAdminPermissionSet_RelayState_updateSessionDuration(t *testing.T) {
funcourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckPermissionSetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionSetConfig_relayState(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "description", rName),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	resource.TestCheckResourceAttr(resourceName, "relay_state", "https://example.com"),
func
	},
	{
Config: testAccPermissionSetConfig_relayStateUpdateSessionDuration(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
funcource.TestCheckResourceAttr(resourceName, "name", rName),
	resource.TestCheckResourceAttr(resourceName, "relay_state", "https://example.com"),
	resource.TestCheckResourceAttr(resourceName, "session_duration", "PT2H"),
),
	},
	{
ResourceName:ceName,
ImportState:
func
},
	})
}


func TestAccSSOAdminPermissionSet_mixedPolicyAttachments(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckPermissionSetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionSetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
),
	},
	{
Config: testAccPermissionSetConfig_mixedPolicyAttachments(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSOAdminPermissionSetExists(ctx, resourceName),
func
	{
ResourceName:ceName,
ImportState:
ImportStateVerify: true,
	},
},
func


func testAccCheckPermissionSetDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).SSOAdminConn(ctx)
func_, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ssoadmin_permission_set" {
continue
	}

	permissionSetARN, instanceARN, err := tfssoadmin.ParseResourceID(rs.Primary.ID)
	if err != nil {
func

	_, err = tfssoadmin.FindPermissionSet(ctx, conn, permissionSetARN, instanceARN)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("SSO Permission Set %s still exists", rs.Primary.ID)
}
funcrn nil
func

func testAccCheckSOAdminPermissionSetExists(ctx context.Context, n string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

permissionSetARN, instanceARN, err := tfssoadmin.ParseResourceID(rs.Primary.ID)
if err != nil {
	return err
}

conn := acctest.Provider.Meta().(*conns.AWSClient).SSOAdminConn(ctx)

_, err = tfssoadmin.FindPermissionSet(ctx, conn, permissionSetARN, instanceARN)

return err
	}
}


func testAccPermissionSetConfig_basic(rName string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  name= %[1]q
  instance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]
}
func
func
func testAccPermissionSetConfig_updateDescription(rName string) string {
func "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  name= %[1]q
  description  = %[1]q
  instance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]
}
`, rName)
}


func testAccPermissionSetConfig_updateRelayState(rName string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  name= %[1]q
  instance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]
  relay_state  = "https://example.com"
}
func


func testAccPermissionSetConfig_updateSessionDuration(rName string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  name%[1]q
  instance_arn tolist(data.aws_ssoadmin_instances.test.arns)[0]
  session_duration = "PT2H"
}
func


func testAccPermissionSetConfig_relayState(rName string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  descriptionq
  name%[1]q
  instance_arn tolist(data.aws_ssoadmin_instances.test.arns)[0]
  relay_stateps://example.com"
  session_duration = "PT1H"
funcName)
}


func testAccPermissionSetConfig_relayStateUpdateSessionDuration(rName string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  descriptionq
  name%[1]q
  instance_arn tolist(data.aws_ssoadmin_instances.test.arns)[0]
  relay_stateps://example.com"
func
`, rName)
}


func testAccPermissionSetConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  name= %[1]q
  instance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]

func= %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}


func testAccPermissionSetConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  name= %[1]q
  instance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]

func= %[3]q
4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}


func testAccPermissionSetConfig_mixedPolicyAttachments(rName string) string {
	return fmt.Sprintf(`
data "aws_partition" "current" {}

data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
funcstance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]
}

resource "aws_ssoadmin_managed_policy_attachment" "test" {
  instance_arn_ssoadmin_permission_set.test.instance_arn
  managed_policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/AlexaForBusinessDeviceSetup"
  permission_set_arn = aws_ssoadmin_permission_set.test.arn
}

data "aws_iam_policy_document" "test" {
  statement {
d = "1"

tions = [
stAllMyBuckets",

funcces = [
{data.aws_partition.current.partition}:s3:::*",

  }
}
resource "aws_ssoadmin_permission_set_inline_policy" "test" {
  inline_policy.aws_iam_policy_document.test.json
  instance_arn_ssoadmin_permission_set.test.instance_arn
  permission_set_arn = aws_ssoadmin_permission_set.test.arn
}
`, rName)
}
func