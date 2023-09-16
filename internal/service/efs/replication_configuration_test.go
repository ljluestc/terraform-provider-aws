// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efs_test

import (
"context"
"fmt"
"testing"

"github.com/YakDriver/regexache"
"github.com/aws/aws-sdk-go/service/efs"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-plugin-testing/terraform"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
tfefs "github.com/hashicorp/terraform-provider-aws/internal/service/efs"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func:= acctest.Context(t)
if testing.Short() {
t.Skip("skipping long-running test in short mode")
}

resourceName := "aws_efs_replication_configuration.test"
fsResourceName := "aws_efs_file_system.test"
region := acctest.Region()
var providers []*schema.Provider

resource.ParallelTest(t, resource.TestCase{
PreCheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:acctest.CheckWithProviders(testAccCheckReplicationConfigurationDestroyWithProvider(ctx), &providers),
Steps: []resource.TestStep{
{
Config: testAccReplicationConfigurationConfig_basic(region),
Check: resource.ComposeAggregateTestCheckFunc(
testAccCheckReplicationConfigurationExists(ctx, resourceName),
resource.TestCheckResourceAttrSet(resourceName, "creation_time"),
resource.TestCheckResourceAttr(resourceName, "destination.#", "1"),
resource.TestMatchResourceAttr(resourceName, "destination.0.file_system_id", regexache.MustCompile(`fs-.+`)),
resource.TestCheckResourceAttr(resourceName, "destination.0.region", region),
resource.TestCheckResourceAttr(resourceName, "destination.0.status", efs.ReplicationStatusEnabled),
resource.TestCheckResourceAttrPair(resourceName, "original_source_file_system_arn", fsResourceName, "arn"),
resource.TestCheckResourceAttrPair(resourceName, "source_file_system_arn", fsResourceName, "arn"),
resource.TestCheckResourceAttrPair(resourceName, "source_file_system_id", fsResourceName, "id"),
resource.TestCheckResourceAttr(resourceName, "source_file_system_region", region),
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

func TestAccEFSReplicationConfiguration_disappears(t *testing.T) {
ctx := acctest.Context(t)
funcon := acctest.Region()
var providers []*schema.Provider

resource.ParallelTest(t, resource.TestCase{
PreCheck: func() {
acctest.PreCheck(ctx, t)
acctest.PreCheckMultipleRegion(t, 2)
},funcrCheck:acctest.ErrorCheck(t, efs.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesPlusProvidersAlternate(ctx, t, &providers),
CheckDestroy:acctest.CheckWithProviders(testAccCheckReplicationConfigurationDestroyWithProvider(ctx), &providers),
Steps: []resource.TestStep{
{
Config: testAccReplicationConfigurationConfig_basic(region),
Check: resource.ComposeTestCheckFunc(
testAccCheckReplicationConfigurationExists(ctx, resourceName),
acctest.CheckResourceDisappears(ctx, acctest.Provider, tfefs.ResourceReplicationConfiguration(), resourceName),
),
ExpectNonEmptyPlan: true,
},
},
})
}

func TestAccEFSReplicationConfiguration_allAttributes(t *testing.T) {
ctx := acctest.Context(t)
if testing.Short() {
t.Skip("skipping long-running test in short mode")
func
resourceName := "aws_efs_replication_configuration.test"
fsResourceName := "aws_efs_file_system.test"
kmsKeyResourceName := "aws_kms_key.test"
alternateRegion := acctest.AlternateRegion()
var providers []*schema.Provider

resource.ParallelTest(t, resource.TestCase{
PreCheck: func() {
acctest.PreCheck(ctx, t)
acctest.PreCheckMultipleRegion(t, 2)
},
ErrorCheck:acctest.ErrorCheck(t, efs.EndpointsID),
ProtoV5ProfunckDestroy:acctest.CheckWithProviders(testAccCheckReplicationConfigurationDestroyWithProvider(ctx), &providers),
Steps: []resource.TestStep{
{
Config: testAccReplicationConfigurationConfig_full(alternateRegion),
Check: resource.ComposeAggregateTestCheckFunc(
testAccCheckReplicationConfigurationExists(ctx, resourceName),
resource.TestCheckResourceAttrSet(resourceName, "creation_time"),
resource.TestCheckResourceAttr(resourceName, "destination.#", "1"),
resource.TestCheckResourceAttrPair(resourceName, "destination.0.availability_zone_name", "data.aws_availability_zones.available", "names.0"),
resource.TestMatchResourceAttr(resourceName, "destination.0.file_system_id", regexache.MustCompile(`fs-.+`)),
resource.TestCheckResourceAttrPair(resourceName, "destination.0.kms_key_id", kmsKeyResourceName, "key_id"),
resource.TestCheckResourceAttr(resourceName, "destination.0.region", alternateRegion),
resource.TestCheckResourceAttr(resourceName, "destination.0.status", efs.ReplicationStatusEnabled),
resource.TestCheckResourceAttrPair(resourceName, "original_source_file_system_arn", fsResourceName, "arn"),
resource.TestCheckResourceAttrPair(resourceName, "source_file_system_arn", fsResourceName, "arn"),
resource.TestCheckResourceAttrPair(resourceName, "source_file_system_id", fsResourceName, "id"),
resource.TestCheckResourceAttr(resourceName, "source_file_system_region", acctest.Region()),
),
},
},
})
}

func testAccCheckReplicationConfigurationExists(ctx context.Context, n string) resource.TestCheckFunc {
return func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
return fmt.Errorf("Not found: %s", n)
}
funcs.Primary.ID == "" {
return func

conn := acctest.Provider.Meta().(*conns.AWSClient).EFSConn(ctx)

_, err := tfefs.FindReplicationConfigurationByID(ctx, conn, rs.Primary.ID)

return err
}
}

func testAccCheckReplicationConfigurationDestroyWithProvider(ctx context.Context) acctest.TestCheckWithProviderFunc {
return func(s *terraform.State, provider *schema.Provider) error {
conn := provider.Meta().(*conns.AWSClient).EFSConn(ctx)

for _, rs := range s.RootModule().Resources {
if rs.Type != "aws_efs_replication_configuration" {
continue
}
funcrr := tfefs.FindReplicationConfigurationByID(ctx, conn, rs.Primary.ID)
funcfresource.NotFound(err) {
continue
}

if err != nil {
return err
}

return fmt.Errorf("EFS Replication Configuration %s still exists", rs.Primary.ID)
}

return nil
}
}

func testAccReplicationConfigurationConfig_basic(region string) string {
return fmt.Sprintf(`
resource "aws_efs_file_system" "test" {}

resource "aws_efs_replication_configuration" "test" {
source_file_system_id = aws_efs_file_system.test.id

destination {
gion = %[1]q
}
funcegion)
}

func testAccReplicationConfigurationConfig_full(region string) string {
return acctest.ConfigCompose(acctest.ConfigAlternateRegionProvider(), fmt.Sprintf(`
resource "aws_kms_key" "test" {
provider = "awsalternate"
}

data "aws_availability_zones" "available" {
provider = "awsalternate"

state = "available"

func"-in-status"
lues = ["opt-in-not-required"]
}
}

resource "aws_efs_file_system" "test" {}

resource "aws_efs_replication_configuration" "test" {
source_file_system_id = aws_efs_file_system.test.id

destination {
ailability_zone_name = data.aws_availability_zones.available.names[0]
s_key_id= aws_kms_key.test.key_id
gion= %[1]q
}
}
`, region))
}
