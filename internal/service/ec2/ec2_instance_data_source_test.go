// Copyright (c) HashiCorp, Inc.// SPDX-License-Identifier: MPL-2.0package ec2_testimport (	"fmt"	"testing"	"github.com/aws/aws-sdk-go/service/ec2"	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"	"github.com/hashicorp/terraform-plugin-testing/helper/resource"	"github.com/hashicorp/terraform-provider-aws/internal/acctest")
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_basic(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_tags(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "ami", resourceName, "ami"),	resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_azUser(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_gp2IopsDevice(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "ami", resourceName, "ami"),	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.#", resourceName, "root_block_device.#"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.0.volume_size", resourceName, "root_block_device.0.volume_size"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.0.volume_type", resourceName, "root_block_device.0.volume_type"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.0.device_name", resourceName, "root_block_device.0.device_name"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.0.iops", resourceName, "root_block_device.0.iops"),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_gp3ThroughputDevice(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_blockDevices(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "ami", resourceName, "ami"),	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.#", resourceName, "root_block_device.#"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.0.volume_size", resourceName, "root_block_device.0.volume_size"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.0.volume_type", resourceName, "root_block_device.0.volume_type"),	resource.TestCheckResourceAttrPair(datasourceName, "root_block_device.0.device_name", resourceName, "root_block_device.0.device_name"),	resource.TestCheckResourceAttrPair(datasourceName, "ebs_block_device.#", resourceName, "ebs_block_device.#"),	//resource.TestCheckResourceAttrPair(datasourceName, "ephemeral_block_device.#", resourceName, "ephemeral_block_device.#"),	// ephemeral block devices don't get saved properly due to API limitations, so this can't actually be tested right now),	},},	})}// Test to verify that ebs_block_device kms_key_id does not elicit a panic
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_ebsKMSKeyID(rName),	},},	})}// Test to verify that root_block_device kms_key_id does not elicit a panic
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_rootBlockDeviceKMSKeyID(rName),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_rootStore(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_privateIP(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "ami", resourceName, "ami"),	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_name_options.#", resourceName, "private_dns_name_options.#"),	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_name_options.0.enable_resource_name_dns_aaaa_record", resourceName, "private_dns_name_options.0.enable_resource_name_dns_aaaa_record"),	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_name_options.0.enable_resource_name_dns_a_record", resourceName, "private_dns_name_options.0.enable_resource_name_dns_a_record"),	resource.TestCheckResourceAttrPair(datasourceName, "private_dns_name_options.0.hostname_type", resourceName, "private_dns_name_options.0.hostname_type"),	resource.TestCheckResourceAttrPair(datasourceName, "private_ip", resourceName, "private_ip"),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_secondaryPrivateIPs(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_ipv6Addresses(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "ami", resourceName, "ami"),	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),	resource.TestCheckResourceAttrPair(datasourceName, "ipv6_addresses.#", resourceName, "ipv6_address_count"),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_keyPair(rName, publicKey),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_vpc(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "ami", resourceName, "ami"),	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),	resource.TestCheckResourceAttrPair(datasourceName, "user_data", resourceName, "user_data"),	resource.TestCheckResourceAttrPair(datasourceName, "associate_public_ip_address", resourceName, "associate_public_ip_address"),	resource.TestCheckResourceAttrPair(datasourceName, "tenancy", resourceName, "tenancy"),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_placementGroup(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_securityGroups(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "ami", resourceName, "ami"),	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),	resource.TestCheckResourceAttrPair(datasourceName, "user_data", resourceName, "user_data"),	resource.TestCheckResourceAttrPair(datasourceName, "vpc_security_group_ids.#", resourceName, "vpc_security_group_ids.#"),	resource.TestCheckResourceAttrPair(datasourceName, "security_groups.#", resourceName, "security_groups.#"),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_vpcSecurityGroups(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_GetPasswordData_trueToFalse(t *testing.T) {	ctx := acctest.Context(t)	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)	if err != nil {t.Fatalf("error generating random SSH key: %s", err)	}	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttr(datasourceName, "get_password_data", "true"),	resource.TestCheckResourceAttrSet(datasourceName, "password_data"),),	},	{Config: testAccInstanceDataSourceConfig_getPassword(rName, publicKey, false),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_GetPasswordData_falseToTrue(t *testing.T) {	ctx := acctest.Context(t)	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)	if err != nil {t.Fatalf("error generating random SSH key: %s", err)	}	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttr(datasourceName, "get_password_data", "false"),	resource.TestCheckNoResourceAttr(datasourceName, "password_data"),),	},	{Config: testAccInstanceDataSourceConfig_getPassword(rName, publicKey, true),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_getUserData(t *testing.T) {	ctx := acctest.Context(t)	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttr(datasourceName, "get_user_data", "true"),	resource.TestCheckResourceAttr(datasourceName, "user_data_base64", "IyEvYmluL2Jhc2gKCmVjaG8gImhlbGxvIHdvcmxkIgo="),),	},	{Config: testAccInstanceDataSourceConfig_getUser(rName, false),Check: resource.ComposeTestCheck
func(	resource.TestCheckResourceAttr(datasourceName, "get_user_data", "true"),	resource.TestCheckResourceAttr(datasourceName, "user_data_base64", "IyEvYmluL2Jhc2gKCmVjaG8gImhlbGxvIHdvcmxkIgo="),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_getUserNoUser(rName, true),Check: resource.ComposeTestCheck
func(	resource.TestCheckResourceAttr(datasourceName, "get_user_data", "false"),	resource.TestCheckNoResourceAttr(datasourceName, "user_data_base64"),	resource.TestCheckResourceAttrPair(datasourceName, "user_data_base64", resourceName, "user_data_base64"),),	},	{Config: testAccInstanceDataSourceConfig_getUserNoUser(rName, true),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_autoRecovery(t *testing.T) {	ctx := acctest.Context(t)	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttr(datasourceName, "maintenance_options.#", "1"),	resource.TestCheckResourceAttr(datasourceName, "maintenance_options.0.auto_recovery", "default"),),	},	{Config: testAccInstanceDataSourceConfig_autoRecovery(rName, "disabled"),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_creditSpecification(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),	resource.TestCheckResourceAttrPair(datasourceName, "credit_specification.#", resourceName, "credit_specification.#"),	resource.TestCheckResourceAttrPair(datasourceName, "credit_specification.0.cpu_credits", resourceName, "credit_specification.0.cpu_credits"),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_metaOptions(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_enclaveOptions(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "enclave_options.#", resourceName, "enclave_options.#"),	resource.TestCheckResourceAttrPair(datasourceName, "enclave_options.0.enabled", resourceName, "enclave_options.0.enabled"),),	},},	})}
func() { acctest.PreCheck(ctx, t) },ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,Steps: []resource.TestStep{	{Config: testAccInstanceDataSourceConfig_blockDeviceTags(rName),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_disableAPIStopTermination(t *testing.T) {	ctx := acctest.Context(t)	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttr(datasourceName, "disable_api_stop", "true"),	resource.TestCheckResourceAttr(datasourceName, "disable_api_termination", "true"),),	},	{Config: testAccInstanceDataSourceConfig_disableAPIStopTermination(rName, false),Check: resource.ComposeTestCheck
func TestAccEC2InstanceDataSource_timeout(t *testing.T) {	ctx := acctest.Context(t)	resourceName := "aws_instance.test"	datasourceName := "data.aws_instance.test"	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{PreCheck:  
func(	resource.TestCheckResourceAttrPair(datasourceName, "ami", resourceName, "ami"),	resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),	resource.TestCheckResourceAttrPair(datasourceName, "instance_type", resourceName, "instance_type"),	resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),),	},},	})}// Lookup based on InstanceID
func testAccInstanceDataSourceConfig_tags(rName string) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`resource "aws_instance" "test" {  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id  instance_type = "t2.small"  tags = {me TestSeed = "%[2]d"  }}data "aws_instance" "test" {  instance_tags = { Na awsance.tags["Name"] TestS = "%[2]d"  }}`, rName, sdkacctest.RandInt()))}// filter on tag, populate more attributes
func testAccInstanceDataSourceConfig_gp2IOPSDevice(rName string) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`resource "aws_instance" "test" {  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id  instance_type = "t3.medium"  root_block_device {lume_type = "gp2" voe_size = 11  }  tags = { Name [1]q  }}data "aws_instance" "test" {  instance_id = aws_instance.test.id}`, rName))}// GP3ThroughputDevice
func testAccInstanceDataSourceConfig_blockDevices(rName string) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`resource "aws_instance" "test" {  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id  instance_type = "t3.medium"  root_block_device {lume_type = "gp2" voe_size = 11  }  ebs_block_device { devicame = "/dev/sdb" volume_s = 9  }  ebs_block_device { device_name"/dev/sdc" volume_size = volume_type = "ioiops}  # Encrypted eblock device  ebs_block_device { device_name = "/dev/sddolume_size = 12 encryptedrue  }  ephemeralock_devicedevice_name  = "/dev/sde" virtual_n = "ephemeral0"  }  ebs_bl_device { device_name = "/dev/sdf" volume_size = 10lume_type = "gp3" through  = 300  }  tags  Name = %[1]q  }}dataws_instance" "test" {  instancd = aws_instance.test.id}`, rName))}
func testAccInstanceDataSourceConfig_rootBlockDeviceKMSKeyID(rName string) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`resource "aws_kms_key" "test" {  deletion_window_in_days = 7}resource "aws_instance" "test" {  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id  instance_type = "t3.medium"  root_block_device {crypted= t kms_kid  = aws_kms_key.test.arn volume_t = "gp2" volume_size11  }  tags = { Name = %[1]q  ata "aws_instance" "test" {  instance_id = aws_instance.test.id}`, rName))}
func testAccInstanceDataSourceConfig_privateIP(rName string) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),testAccInstanceVPCConfig(rName, false, 1),fmt.Sprintf(`resource "aws_instance" "test" {  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id  instance_type = "t2.micro"  subnet_idet.test.id  private_ip"10.1.1.42"  tags = { Na= %[1]q  }}data "aws_instance" "test" {  instance_id = aws_instance.test.id}`, rName))}
func testAccInstanceDataSourceConfig_ipv6Addresses(rName string) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),testAccInstanceVPCIPv6Config(rName),fmt.Sprintf(`resource "aws_instance" "test" {  ami = data.aws_ami.amzn-ami-minimal-hvm-ebs.id  instance_typero"  subnet_id = aws_subnet.test.id  ipv6_address_count = 1  tags = {me = %[1]q  }}data "aws_instance" "test" {  instance_id = aws_instance.test.id}`, rName))}
func testAccInstanceDataSourceConfig_vpc(rName string) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),testAccInstanceVPCConfig(rName, false, 1),fmt.Sprintf(`resource "aws_instance" "test" {  ami = data.aws_ami.amzn-ami-minimal-hvm-ebs.id  instance_type= "t2.small"  subnet_idaws_subnet.test.id  associate_public_ip_address = true  #tenancyted"  # pre-encoded base64 data  user_data = "3dc39dda39be1205215e776bad998da361a5955d"  tags = { Na= %[1]q  }}data "aws_instance" "test" {  instance_id = aws_instance.test.id}`, rName))}
func testAccInstanceDataSourceConfig_securityGroups(rName string) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`resource "aws_security_group" "test" {  name = %[1]q  ingress {otocol  = "icmp" frport = -1 to_po -1 selfags  Name = %[1]q  esource "aws_instance" "test" {  ami = data.aws_ami.amami-minimal-hvm-ebs.id  instance_type= "t2.small"  securitroups = [aws_security_group.test.name]  user_data  = "with-cter's"  tags = { Name = %[1]q  }}data "ainstance" "test" {  instance_id = aws_instance.test.id}`, rName))}
func testAccInstanceDataSourceConfig_getPassword(rName, publicKey string, val bool) string {	return acctest.ConfigCompose(testAccLatestWindowsServer2016CoreAMIConfig(), fmt.Sprintf(`resource "aws_key_pair" "test" {  key_name[1]q  public_key = %[2]q}resource "aws_instance" "test" {  ami  = data.aws_ami.win2016core-ami.id  instance_type = "t2.medium"  key_name_pair.test.key_name  tags = { Na= %[1]q  }}data "aws_instance" "test" {  instance_id = aws_instance.test.id  get_password_data = %[3]t}`, rName, publicKey, val))}
func testAccInstanceDataSourceConfig_getUserNoUser(rName string, getUserData bool) string {	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),testAccInstanceVPCConfig(rName, false, 1),
funcurce "aws_instance" "test" {
funcstance_type = "t2.micro"
func
func %[1]q
func
func "aws_instance" "test" {
funcstance_idws_instance.test.id
funcName, getUserData))
func
func testAccInstanceDataSourceConfig_autoRecovery(rName string, val string) string {
funcest.ConfigLatestAmazonLinuxHVMEBSAMI(),
funcSprintf(`
funci  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
funcbnet_idet.test.id
funcintenance_options {
func
funcgs = {
func
func
funcstance_id = aws_instance.test.id
funcName, val))
func
func testAccInstanceDataSourceConfig_creditSpecification(rName string) string {
funcest.ConfigLatestAmazonLinuxHVMEBSAMI(),
funcSprintf(`
funci  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
funcbnet_idet.test.id
funcedit_specification {
func
funcgs = {
func
func
funcstance_id = aws_instance.test.id
funcName))
func
func testAccInstanceDataSourceConfig_metaOptions(rName string) string {
funcest.ConfigLatestAmazonLinuxHVMEBSAMI(),
funcest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
funcurce "aws_instance" "test" {
funcstance_type = data.aws_ec2_instance_type_offering.available.instance_type
func
func %[1]q
func
funcndpoint= "enabled"
funcut_response_hop_limit = 2
func
func
funcstance_id = aws_instance.test.id
funcName))
func
func testAccInstanceDataSourceConfig_enclaveOptions(rName string) string {
funcest.ConfigLatestAmazonLinuxHVMEBSAMI(),
funcest.AvailableEC2InstanceTypeForRegion("c5a.xlarge", "c5.xlarge"),
funcurce "aws_instance" "test" {
funcstance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_idet.test.id

  tags = {
me = %[1]q
  }

  enclave_options {
abled = true
  }
}

data "aws_instance" "test" {
  instance_id = aws_instance.test.id
}
`, rName))
}


func testAccInstanceDataSourceConfig_blockDeviceTags(rName string) string {
funcest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type

  tags = {
me = %[1]q
  }

  ebs_block_device {
vice_name = "/dev/xvdc"
lume_size = 10

gs = {
%[1]q
"SapereAude"

  }

  root_block_device {
gs = {
%[1]q
"VincitQuiSeVincit"

func

data "aws_instance" "test" {
  instance_id = aws_instance.test.id
}
`, rName))
}


func testAccInstanceDataSourceConfig_disableAPIStopTermination(rName string, disableApiStopTermination bool) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
testAccInstanceVPCConfig(rName, false, 1),
fmt.Sprintf(`
resource "aws_instance" "test" {
  amis_ami.amzn-ami-minimal-hvm-ebs.id
  disable_api_stop
  disable_api_termination = %[2]t
  instance_type  = "t2.micro"
  subnet_id= aws_subnet.test.id

  tags = {
me = %[1]q
  }
}

funcstance_id = aws_instance.test.id
}
`, rName, disableApiStopTermination))
}


func testAccInstanceDataSourceConfig_timeout(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
fmt.Sprintf(`
resource "aws_instance" "test" {
  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = "t2.small"

  tags = {
me = %[1]q
  }
}

data "aws_instance" "test" {
  filter {
me= "tance-id"
lues = [aws_instance.test.id]
  }

  timeouts {
ad = "60m"
  }
}
`, rName))
funcfuncfuncfunc