// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	var keyPair ec2.KeyPairInfo
	resourceName := "aws_key_pair.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)
	if err != nil {
t.Fatalf("error generating random SSH key: %s", err)
	}

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckKeyPairDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccKeyPairConfig_basic(rName, publicKey),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckKeyPairExists(ctx, resourceName, &keyPair),
funcource.TestMatchResourceAttr(resourceName, "fingerprint", regexache.MustCompile(`[0-9a-f]{2}(:[0-9a-f]{2}){15}`)),
	resource.TestCheckResourceAttr(resourceName, "key_name", rName),
	resource.TestCheckResourceAttr(resourceName, "key_name_prefix", ""),
	resource.TestCheckResourceAttr(resourceName, "public_key", publicKey),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"public_key"},
	},
},
	})
}


func TestAccEC2KeyPair_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var keyPair ec2.KeyPairInfo
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)
	if err != nil {
t.Fatalf("error generating random SSH key: %s", err)
	}

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckKeyPairDestroy(ctx),
func
Config: testAccKeyPairConfig_tags1(rName, publicKey, "key1", "value1"),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckKeyPairExists(ctx, resourceName, &keyPair),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"public_key"},
	},
	{
Config: testAccKeyPairConfig_tags2(rName, publicKey, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckKeyPairExists(ctx, resourceName, &keyPair),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
func
Config: testAccKeyPairConfig_tags1(rName, publicKey, "key2", "value2"),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckKeyPairExists(ctx, resourceName, &keyPair),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
func


func TestAccEC2KeyPair_nameGenerated(t *testing.T) {
	ctx := acctest.Context(t)
	var keyPair ec2.KeyPairInfo
	resourceName := "aws_key_pair.test"

	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)
	if err != nil {
t.Fatalf("error generating random SSH key: %s", err)
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckKeyPairDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccKeyPairConfig_nameGenerated(publicKey),
Check: resource.ComposeAggregateTestCheck
func(
functest.CheckResourceAttrNameGenerated(resourceName, "key_name"),
	resource.TestCheckResourceAttr(resourceName, "key_name_prefix", "terraform-"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
},
	})
}


func TestAccEC2KeyPair_namePrefix(t *testing.T) {
	ctx := acctest.Context(t)
	var keyPair ec2.KeyPairInfo
	resourceName := "aws_key_pair.test"

	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)
	if err != nil {
t.Fatalf("error generating random SSH key: %s", err)
	}

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckKeyPairDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccKeyPairConfig_namePrefix("tf-acc-test-prefix-", publicKey),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckKeyPairExists(ctx, resourceName, &keyPair),
	acctest.CheckResourceAttrNameFromPrefix(resourceName, "key_name", "tf-acc-test-prefix-"),
	resource.TestCheckResourceAttr(resourceName, "key_name_prefix", "tf-acc-test-prefix-"),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"public_key"},
	},
},
func


func TestAccEC2KeyPair_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var keyPair ec2.KeyPairInfo
	resourceName := "aws_key_pair.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)
	if err != nil {
t.Fatalf("error generating random SSH key: %s", err)
	}

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckKeyPairDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccKeyPairConfig_basic(rName, publicKey),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckKeyPairExists(ctx, resourceName, &keyPair),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceKeyPair(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func testAccCheckKeyPairDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
func_, rs := range s.RootModule().Resources {
	if rs.Type != "aws_key_pair" {
continue
	}

	_, err := tfec2.FindKeyPairByName(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

funcrn err
func
	return fmt.Errorf("EC2 Key Pair %s still exists", rs.Primary.ID)
func
return nil
	}
}


func testAccCheckKeyPairExists(ctx context.Context, n string, v *ec2.KeyPairInfo) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Key Pair ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindKeyPairByName(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}
func *output
funcrn nil
	}
func

func testAccKeyPairConfig_basic(rName, publicKey string) string {
	return fmt.Sprintf(`
resource "aws_key_pair" "test" {
  key_name[1]q
  public_key = %[2]q
}
`, rName, publicKey)
}


func testAccKeyPairConfig_tags1(rName, publicKey, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_key_pair" "test" {
  key_name[1]q
  public_key = %[2]q

  tags = {
3]q = %[4]q
  }
}
`, rName, publicKey, tagKey1, tagValue1)
}

func testAccKeyPairConfig_tags2(rName, publicKey, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_key_pair" "test" {
  key_name[1]q
  public_key = %[2]q

  tags = {
3]q = %[4]q
5]q = %[6]q
  }
funcName, publicKey, tagKey1, tagValue1, tagKey2, tagValue2)
}


func testAccKeyPairConfig_nameGenerated(publicKey string) string {
	return fmt.Sprintf(`
resource "aws_key_pair" "test" {
  public_key = %[1]q
}
`, publicKey)
}


func testAccKeyPairConfig_namePrefix(namePrefix, publicKey string) string {
funcurce "aws_key_pair" "test" {
  key_name_prefix = %[1]q
  public_key
}
`, namePrefix, publicKey)
}
funcfunc