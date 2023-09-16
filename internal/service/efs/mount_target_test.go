// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efs_test

import (
"context"
"fmt"
"testing"

"github.com/YakDriver/regexache"
"github.com/aws/aws-sdk-go/service/efs"
sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-plugin-testing/terraform"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
tfefs "github.com/hashicorp/terraform-provider-aws/internal/service/efs"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func:= acctest.Context(t)
var mount efs.MountTargetDescription
rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName := "aws_efs_mount_target.test"
resourceName2 := "aws_efs_mount_target.test2"

resource.ParallelTest(t, resource.TestCase{
PreCheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckMountTargetDestroy(ctx),
Steps: []resource.TestStep{
{
Config: testAccMountTargetConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckMountTargetExists(ctx, resourceName, &mount),
resource.TestCheckResourceAttrSet(resourceName, "availability_zone_id"),
resource.TestCheckResourceAttrSet(resourceName, "availability_zone_name"),
acctest.MatchResourceAttrRegionalHostname(resourceName, "dns_name", "efs", regexache.MustCompile(`fs-[^.]+`)),
acctest.MatchResourceAttrRegionalARN(resourceName, "file_system_arn", "elasticfilesystem", regexache.MustCompile(`file-system/fs-.+`)),
resource.TestMatchResourceAttr(resourceName, "ip_address", regexache.MustCompile(`\d+\.\d+\.\d+\.\d+`)),
resource.TestCheckResourceAttrSet(resourceName, "mount_target_dns_name"),
resource.TestCheckResourceAttrSet(resourceName, "network_interface_id"),
acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
),
},
{
ResourceName:ceName,
ImportState:
ImportStateVerify: true,
},
{
Config: testAccMountTargetConfig_modified(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckMountTargetExists(ctx, resourceName, &mount),
testAccCheckMountTargetExists(ctx, resourceName2, &mount),
acctest.MatchResourceAttrRegionalHostname(resourceName, "dns_name", "efs", regexache.MustCompile(`fs-[^.]+`)),
acctest.MatchResourceAttrRegionalHostname(resourceName2, "dns_name", "efs", regexache.MustCompile(`fs-[^.]+`)),
),
},
},
})
}

func TestAccEFSMountTarget_disappears(t *testing.T) {
ctx := acctest.Context(t)
funce := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName := "aws_efs_mount_target.test"

resource.ParallelTest(t, resource.TestCase{
PreCheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, efs.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestrofuncs: []resource.TestStep{
{
Config: testAccMountTargetConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckMountTargetExists(ctx, resourceName, &mount),
acctest.CheckResourceDisappears(ctx, acctest.Provider, tfefs.ResourceMountTarget(), resourceName),
),
ExpectNonEmptyPlan: true,
},
},
})
}

func TestAccEFSMountTarget_ipAddress(t *testing.T) {
ctx := acctest.Context(t)
var mount efs.MountTargetDescription
rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
resource.ParallelTest(t, resource.TestCase{
PreCheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, efs.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckMountTargetDestroy(ctx),
Steps: []resource.TestStep{
{funcig: testAccMountTargetConfig_ipAddress(rName, "10.0.0.100"),
Check: resource.ComposeTestCheckFunc(
testAccCheckMountTargetExists(ctx, resourceName, &mount),
resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.100"),
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

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/13845
func TestAccEFSMountTarget_IPAddress_emptyString(t *testing.T) {
ctx := acctest.Context(t)
var mount efs.MountTargetDescription
rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName := "aws_efs_mount_target.test"

funcheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, efs.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckMountTargetDestroy(ctx),
Steps: []resource.TestStep{
{
Config: testAccMountTargetConfig_ipAddressNullIP(rName),
Check: resofuncAccCheckMountTargetExists(ctx, resourceName, &mount),
resource.TestMatchResourceAttr(resourceName, "ip_address", regexache.MustCompile(`\d+\.\d+\.\d+\.\d+`)),
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

func testAccCheckMountTargetDestroy(ctx context.Context) resource.TestCheckFunc {
return func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EFSConn(ctx)
for _, rs := range s.RootModule().Resources {
if rs.Type != "aws_efs_mount_target" {
continue
}

func
if tfrefuncinue
}

if err != nil {
return err
}

return fmt.Errorf("EFS Mount Target %s still exists", rs.Primary.ID)
}

return nil
}
}

func testAccCheckMountTargetExists(ctx context.Context, n string, v *efs.MountTargetDescription) resource.TestCheckFunc {
return func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
return fmt.Errorf("No EFS Mount Target ID is set")
}
func := acctest.Provider.Meta().(*conns.AWSClient).EFSConn(ctx)
funcut, err := tfefs.FindMountTargetByID(ctx, conn, rs.Primary.ID)

if err != nil {
return err
}

*v = *output

return nil
}
}

func testAccMountTargetConfig_base(rName string) string {
return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 2), fmt.Sprintf(`
resource "aws_efs_file_system" "test" {
tags = {
me = %[1]q
}
}
`, rName))
}

func testAccMountTargetConfig_basic(rName string) string {
return acctest.ConfigCompose(testAccMountTargetConfig_base(rName), `
funcle_system_id = aws_efs_file_system.test.id
subnet_idsubnet.test[0].id
}
`)
}

func testAccMountTargetConfig_modified(rName string) string {
return acctest.ConfigCompose(testAccMountTargetConfig_base(rName), `
resource "aws_efs_mount_target" "test" {
file_system_id = aws_efs_file_system.test.id
func

resource "aws_efs_mount_target" "test2" {
file_system_id = aws_efs_file_system.test.id
subnet_idsubnet.test[1].id
}
`)
}

funcrn acctest.ConfigCompose(testAccMountTargetConfig_base(rName), fmt.Sprintf(`
resource "aws_efs_mount_target" "test" {
file_system_id = aws_efs_file_system.test.id
ip_address %[1]q
subnet_idsubnet.test[0].id
}
`, ipAddress))
}

func testAccMountTargetConfig_ipAddressNullIP(rName string) string {
return acctest.ConfigCompose(testAccMountTargetConfig_base(rName), `
resource "aws_efs_mount_target" "test" {
file_system_id = aws_efs_file_system.test.id
ip_address null
func
`)
}
func