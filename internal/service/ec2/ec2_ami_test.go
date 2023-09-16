// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"strings"
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
	var ami ec2.Image
	resourceName := "aws_ami.test"
	snapshotResourceName := "aws_ebs_snapshot.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMIConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
functest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "ec2", regexache.MustCompile(`image/ami-.+`)),
	resource.TestCheckResourceAttr(resourceName, "description", ""),
	resource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ebs_block_device.*", map[string]string{
"delete_on_termination": "true",
"device_name":"/dev/sda1",
"encrypted":alse",
"iops":,
"throughput":,
"volume_size":"8",
"outpost_arn":"",
"volume_type":"standard",
	}),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "ebs_block_device.*.snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "ena_support", "true"),
	resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "hypervisor", "xen"),
	resource.TestCheckResourceAttr(resourceName, "image_type", "machine"),
	resource.TestCheckResourceAttr(resourceName, "imds_support", ""),
	resource.TestCheckResourceAttr(resourceName, "kernel_id", ""),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "platform_details", "Linux/UNIX"),
	resource.TestCheckResourceAttr(resourceName, "ramdisk_id", ""),
	resource.TestCheckResourceAttr(resourceName, "root_device_name", "/dev/sda1"),
	resource.TestCheckResourceAttrPair(resourceName, "root_snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "sriov_net_support", "simple"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tpm_support", ""),
	resource.TestCheckResourceAttr(resourceName, "usage_operation", "RunInstances"),
	resource.TestCheckResourceAttr(resourceName, "virtualization_type", "hvm"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
	},
},
	})
}


func TestAccEC2AMI_deprecateAt(t *testing.T) {
	ctx := acctest.Context(t)
	var ami ec2.Image
funcpshotResourceName := "aws_ebs_snapshot.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	deprecateAt := "2027-10-15T13:17:00.000Z"
	deprecateAtUpdated := "2028-10-15T13:17:00.000Z"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
func
Config: testAccAMIConfig_deprecateAt(rName, deprecateAt),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "architecture", "x86_64"),
	acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "ec2", regexache.MustCompile(`image/ami-.+`)),
	resource.TestCheckResourceAttr(resourceName, "deprecation_time", deprecateAt),
funcource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ebs_block_device.*", map[string]string{
"delete_on_termination": "true",
"device_name":"/dev/sda1",
"encrypted":alse",
"iops":,
"throughput":,
"volume_size":"8",
"outpost_arn":"",
"volume_type":"standard",
	}),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "ebs_block_device.*.snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "ena_support", "true"),
	resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "kernel_id", ""),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "ramdisk_id", ""),
	resource.TestCheckResourceAttr(resourceName, "root_device_name", "/dev/sda1"),
	resource.TestCheckResourceAttrPair(resourceName, "root_snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "sriov_net_support", "simple"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "virtualization_type", "hvm"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
	},
	{
Config: testAccAMIConfig_deprecateAt(rName, deprecateAtUpdated),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "architecture", "x86_64"),
	acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "ec2", regexache.MustCompile(`image/ami-.+`)),
	resource.TestCheckResourceAttr(resourceName, "deprecation_time", deprecateAtUpdated),
	resource.TestCheckResourceAttr(resourceName, "description", ""),
funcource.TestCheckTypeSetElemNestedAttrs(resourceName, "ebs_block_device.*", map[string]string{
"delete_on_termination": "true",
"device_name":"/dev/sda1",
"encrypted":alse",
"iops":,
"throughput":,
"volume_size":"8",
"outpost_arn":"",
"volume_type":"standard",
	}),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "ebs_block_device.*.snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "ena_support", "true"),
	resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "kernel_id", ""),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "ramdisk_id", ""),
	resource.TestCheckResourceAttr(resourceName, "root_device_name", "/dev/sda1"),
	resource.TestCheckResourceAttrPair(resourceName, "root_snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "sriov_net_support", "simple"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "virtualization_type", "hvm"),
),
	},
},
	})
}


func TestAccEC2AMI_description(t *testing.T) {
	ctx := acctest.Context(t)
	var ami ec2.Image
	resourceName := "aws_ami.test"
	snapshotResourceName := "aws_ebs_snapshot.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	desc := sdkacctest.RandomWithPrefix("desc")
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMIConfig_desc(rName, desc),
Check: resource.ComposeAggregateTestCheck
functAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "architecture", "x86_64"),
	acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "ec2", regexache.MustCompile(`image/ami-.+`)),
	resource.TestCheckResourceAttr(resourceName, "description", desc),
	resource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ebs_block_device.*", map[string]string{
"delete_on_termination": "true",
"device_name":"/dev/sda1",
funcs":,
"throughput":,
"volume_size":"8",
"outpost_arn":"",
"volume_type":"standard",
	}),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "ebs_block_device.*.snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "ena_support", "true"),
	resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "kernel_id", ""),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "ramdisk_id", ""),
	resource.TestCheckResourceAttr(resourceName, "root_device_name", "/dev/sda1"),
	resource.TestCheckResourceAttrPair(resourceName, "root_snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "sriov_net_support", "simple"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "virtualization_type", "hvm"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
	},
	{
Config: testAccAMIConfig_desc(rName, descUpdated),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "architecture", "x86_64"),
	acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "ec2", regexache.MustCompile(`image/ami-.+`)),
	resource.TestCheckResourceAttr(resourceName, "description", descUpdated),
	resource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ebs_block_device.*", map[string]string{
"delete_on_termination": "true",
"device_name":"/dev/sda1",
"encrypted":alse",
funcoughput":,
"volume_size":"8",
"outpost_arn":"",
"volume_type":"standard",
	}),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "ebs_block_device.*.snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "ena_support", "true"),
	resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "kernel_id", ""),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "ramdisk_id", ""),
	resource.TestCheckResourceAttr(resourceName, "root_device_name", "/dev/sda1"),
	resource.TestCheckResourceAttrPair(resourceName, "root_snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "sriov_net_support", "simple"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "virtualization_type", "hvm"),
),
	},
},
	})
}


func TestAccEC2AMI_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var ami ec2.Image
	resourceName := "aws_ami.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccAMIConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceAMI(), resourceName),
),
func
},
	})
}


func TestAccEC2AMI_ephemeralBlockDevices(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_ami.test"
	snapshotResourceName := "aws_ebs_snapshot.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccAMIConfig_ephemeralBlockDevices(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "architecture", "x86_64"),
	acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "ec2", regexache.MustCompile(`image/ami-.+`)),
	resource.TestCheckResourceAttr(resourceName, "description", ""),
	resource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ebs_block_device.*", map[string]string{
funcice_name":"/dev/sda1",
"encrypted":alse",
"iops":,
"throughput":,
"volume_size":"8",
"outpost_arn":"",
"volume_type":"standard",
	}),
funcource.TestCheckResourceAttr(resourceName, "ena_support", "true"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ephemeral_block_device.*", map[string]string{
"device_name":"/dev/sdb",
"virtual_name": "ephemeral0",
	}),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ephemeral_block_device.*", map[string]string{
"device_name":"/dev/sdc",
"virtual_name": "ephemeral1",
	}),
	resource.TestCheckResourceAttr(resourceName, "kernel_id", ""),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "ramdisk_id", ""),
	resource.TestCheckResourceAttr(resourceName, "root_device_name", "/dev/sda1"),
	resource.TestCheckResourceAttrPair(resourceName, "root_snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "sriov_net_support", "simple"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "virtualization_type", "hvm"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
	},
},
	})
}


func TestAccEC2AMI_gp3BlockDevice(t *testing.T) {
	ctx := acctest.Context(t)
	var ami ec2.Image
	resourceName := "aws_ami.test"
	snapshotResourceName := "aws_ebs_snapshot.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMIConfig_gp3BlockDevice(rName),
Check: resource.ComposeAggregateTestCheck
functAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "architecture", "x86_64"),
	acctest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "ec2", regexache.MustCompile(`image/ami-.+`)),
	resource.TestCheckResourceAttr(resourceName, "description", ""),
	resource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "2"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ebs_block_device.*", map[string]string{
"delete_on_termination": "true",
"device_name":"/dev/sda1",
"encrypted":alse",
funcoughput":,
"volume_size":"8",
"outpost_arn":"",
"volume_type":"standard",
	}),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "ebs_block_device.*.snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ebs_block_device.*", map[string]string{
"delete_on_termination": "false",
funcrypted":rue",
"iops":0",
"throughput":0",
"volume_size":"10",
"outpost_arn":"",
"volume_type":"gp3",
	}),
	resource.TestCheckResourceAttr(resourceName, "ena_support", "false"),
	resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "kernel_id", ""),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "ramdisk_id", ""),
	resource.TestCheckResourceAttr(resourceName, "root_device_name", "/dev/sda1"),
	resource.TestCheckResourceAttrPair(resourceName, "root_snapshot_id", snapshotResourceName, "id"),
	resource.TestCheckResourceAttr(resourceName, "sriov_net_support", "simple"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "virtualization_type", "hvm"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
	},
},
	})
}


func TestAccEC2AMI_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var ami ec2.Image
	resourceName := "aws_ami.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMIConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
func
Config: testAccAMIConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
func
	{
Config: testAccAMIConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}


func TestAccEC2AMI_outpost(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_ami.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccAMIConfig_outpost(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckTypeSetElemAttrPair(resourceName, "ebs_block_device.*.outpost_arn", " data.aws_outposts_outpost.test", "arn"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
funcnage_ebs_snapshots",
},
	},
},
	})
}


func := acctest.Context(t)
	var ami ec2.Image
	resourceName := "aws_ami.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMIConfig_boot(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "boot_mode", "uefi"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
func
	})
}


func TestAccEC2AMI_tpmSupport(t *testing.T) {
	ctx := acctest.Context(t)
	var ami ec2.Image
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
func
Config: testAccAMIConfig_tpmSupport(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "tpm_support", "v2.0"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
	},
},
	})
func

func TestAccEC2AMI_imdsSupport(t *testing.T) {
	ctx := acctest.Context(t)
	var ami ec2.Image
	resourceName := "aws_ami.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMIConfig_imdsSupport(rName),
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	resource.TestCheckResourceAttr(resourceName, "imds_support", "v2.0"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"manage_ebs_snapshots",
},
	},
},
	})
}


func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for n, rs := range s.RootModule().Resources {
	// The configuration may contain aws_ami data sources.
	// Ignore them.
funcinue
	}

	if rs.Type != "aws_ami" {
continue
	}

	_, err := tfec2.FindImageByID(ctx, conn, rs.Primary.ID)
functfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 AMI %s still exists", rs.Primary.ID)
}

// Check for managed EBS snapshots.
return testAccCheckEBSSnapshotDestroy(ctx)(s)
	}
}


func testAccCheckAMIExists(ctx context.Context, n string, v *ec2.Image) resource.TestCheck
funcurn 
funcok := s.RootModule().Resources[n]
if !ok {
func

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 AMI ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindImageByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccAMIConfig_base(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
resource "aws_ebs_volume" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
size

tags = {
me = %[1]q
}
}
funcurce "aws_ebs_snapshot" "test" {
func
tags = {
func
}
`, rName))
}


func testAccAMIConfig_basic(rName string) string {
	return acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
ena_support= true
name = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
}
}
`, rName))
}


funcurn acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
ena_support= true
name = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"
deprecation_time%[2]q

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
}
}
`, rName, deprecateAt))
}


func testAccAMIConfig_desc(rName, desc string) string {
	return acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
funcme = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"
description= %[2]q

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
}
}
`, rName, desc))
}


func testAccAMIConfig_ephemeralBlockDevices(rName string) string {
	return acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
funcme = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
}

ephemeral_block_device {
vice_name= "/dev/sdb"
rtual_name = "ephemeral0"
}

ephemeral_block_device {
vice_name= "/dev/sdc"
rtual_name = "ephemeral1"
}
}
`, rName))
func

func testAccAMIConfig_gp3BlockDevice(rName string) string {
	return acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
ena_support= false
name = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
}

ebs_block_device {
lete_on_termination = false
vice_name= "/dev/sdb"
func1
roughput= 5
lume_size= 10
lume_type= "gp3"
}
}
`, rName))
}


func testAccAMIConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
ena_support= true
name = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
}

tags = {
2]q = %[3]q
}
}
func


func testAccAMIConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
ena_support= true
name = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
}

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2))
}


func testAccAMIConfig_outpost(rName string) string {
	return acctest.ConfigCompose(
funcSprintf(`
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "test" {
id = tolist(data.aws_outposts_outposts.test.ids)[0]
}

resource "aws_ami" "test" {
ena_support= true
name = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
tpost_arn = data.aws_outposts_outpost.test.arn
}
}
`, rName))
}


funcurn acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
ena_support= true
name = %[1]q
root_device_name"/dev/sda1"
virtualization_type = "hvm"
boot_mode= "uefi"

ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
}
}
`, rName))
}


func testAccAMIConfig_tpmSupport(rName string) string {
	return acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
funcme = %[1]q
root_device_name"/dev/xvda"
virtualization_type = "hvm"
boot_mode= "uefi"
tpm_support= "v2.0"

ebs_block_device {
vice_name = "/dev/xvda"
apshot_id = aws_ebs_snapshot.test.id
}
}
`, rName))
}


func testAccAMIConfig_imdsSupport(rName string) string {
	return acctest.ConfigCompose(
testAccAMIConfig_base(rName),
fmt.Sprintf(`
resource "aws_ami" "test" {
name = %[1]q
root_device_name"/dev/xvda"
virtualization_type = "hvm"
boot_mode= "uefi"
imds_support"

func_name = "/dev/xvda"
apshot_id = aws_ebs_snapshot.test.id
}
}
`, rName))
}
funcfunc