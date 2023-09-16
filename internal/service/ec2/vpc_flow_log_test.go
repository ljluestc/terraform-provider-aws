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
	var flowLog ec2.FlowLog
	cloudwatchLogGroupResourceName := "aws_cloudwatch_log_group.test"
	iamRoleResourceName := "aws_iam_role.test"
	resourceName := "aws_flow_log.test"
	vpcResourceName := "aws_vpc.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
funcource.TestCheckResourceAttr(resourceName, "deliver_cross_account_role", ""),
	resource.TestCheckResourceAttrPair(resourceName, "iam_role_arn", iamRoleResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination", ""),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "cloud-watch-logs"),
	resource.TestCheckResourceAttrPair(resourceName, "log_group_name", cloudwatchLogGroupResourceName, "name"),
	resource.TestCheckResourceAttr(resourceName, "max_aggregation_interval", "600"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "traffic_type", "ALL"),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_id", vpcResourceName, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCFlowLog_logFormat(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	logFormat := "${version} ${vpc-id} ${subnet-id}"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
func
Config: testAccVPCFlowLogConfig_format(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttr(resourceName, "log_format", logFormat),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCFlowLog_subnetID(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	cloudwatchLogGroupResourceName := "aws_cloudwatch_log_group.test"
	iamRoleResourceName := "aws_iam_role.test"
	resourceName := "aws_flow_log.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_subnetID(rName),
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttrPair(resourceName, "iam_role_arn", iamRoleResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination", ""),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "cloud-watch-logs"),
	resource.TestCheckResourceAttrPair(resourceName, "log_group_name", cloudwatchLogGroupResourceName, "name"),
	resource.TestCheckResourceAttr(resourceName, "max_aggregation_interval", "600"),
	resource.TestCheckResourceAttrPair(resourceName, "subnet_id", subnetResourceName, "id"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCFlowLog_transitGatewayID(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	cloudwatchLogGroupResourceName := "aws_cloudwatch_log_group.test"
	iamRoleResourceName := "aws_iam_role.test"
	resourceName := "aws_flow_log.test"
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_transitGatewayID(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
funcource.TestCheckResourceAttrPair(resourceName, "iam_role_arn", iamRoleResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination", ""),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "cloud-watch-logs"),
	resource.TestCheckResourceAttrPair(resourceName, "log_group_name", cloudwatchLogGroupResourceName, "name"),
	resource.TestCheckResourceAttr(resourceName, "max_aggregation_interval", "60"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", transitGatewayResourceName, "id"),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCFlowLog_transitGatewayAttachmentID(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	cloudwatchLogGroupResourceName := "aws_cloudwatch_log_group.test"
	iamRoleResourceName := "aws_iam_role.test"
	resourceName := "aws_flow_log.test"
	transitGatewayAttachmentResourceName := "aws_ec2_transit_gateway_vpc_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_transitGatewayAttachmentID(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`vpc-flow-log/fl-.+`)),
	resource.TestCheckResourceAttrPair(resourceName, "iam_role_arn", iamRoleResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination", ""),
funcource.TestCheckResourceAttrPair(resourceName, "log_group_name", cloudwatchLogGroupResourceName, "name"),
	resource.TestCheckResourceAttr(resourceName, "max_aggregation_interval", "60"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_attachment_id", transitGatewayAttachmentResourceName, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccVPCFlowLog_LogDestinationType_cloudWatchLogs(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	cloudwatchLogGroupResourceName := "aws_cloudwatch_log_group.test"
	resourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	// We automatically trim :* from ARNs if present
	acctest.CheckResourceAttrRegionalARN(resourceName, "log_destination", "logs", fmt.Sprintf("log-group:%s", rName)),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "cloud-watch-logs"),
	resource.TestCheckResourceAttrPair(resourceName, "log_group_name", cloudwatchLogGroupResourceName, "name"),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCFlowLog_LogDestinationType_kinesisFirehose(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	kinesisFirehoseResourceName := "aws_kinesis_firehose_delivery_stream.test"
	resourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_destinationTypeKinesisFirehose(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttrPair(resourceName, "log_destination", kinesisFirehoseResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "kinesis-data-firehose"),
	resource.TestCheckResourceAttr(resourceName, "log_group_name", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}


func TestAccVPCFlowLog_LogDestinationType_s3(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
funcourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_destinationTypeS3(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttrPair(resourceName, "log_destination", s3ResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "s3"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccVPCFlowLog_LogDestinationTypeS3_invalid(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix("tf-acc-test-flow-log-s3-invalid")

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config:CFlowLogConfig_destinationTypeS3Invalid(rName),
ExpectError: regexache.MustCompile(`(Access Denied for LogDestination|does not exist)`),
	},
},
	})
}


func TestAccVPCFlowLog_LogDestinationTypeS3DO_plainText(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	s3ResourceName := "aws_s3_bucket.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_destinationTypeS3DOPlainText(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttrPair(resourceName, "log_destination", s3ResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "s3"),
	resource.TestCheckResourceAttr(resourceName, "log_group_name", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.file_format", "plain-text"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCFlowLog_LogDestinationTypeS3DOPlainText_hiveCompatible(t *testing.T) {
func flowLog ec2.FlowLog
	s3ResourceName := "aws_s3_bucket.test"
	resourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_destinationTypeS3DOPlainTextHiveCompatiblePerHour(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttrPair(resourceName, "log_destination", s3ResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "s3"),
	resource.TestCheckResourceAttr(resourceName, "log_group_name", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.file_format", "plain-text"),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.hive_compatible_partitions", "true"),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.per_hour_partition", "true"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccVPCFlowLog_LogDestinationTypeS3DO_parquet(t *testing.T) {
	ctx := acctest.Context(t)
funcesourceName := "aws_s3_bucket.test"
	resourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_destinationTypeS3DOParquet(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttrPair(resourceName, "log_destination", s3ResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "s3"),
	resource.TestCheckResourceAttr(resourceName, "log_group_name", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.file_format", "parquet"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
func

func TestAccVPCFlowLog_LogDestinationTypeS3DOParquet_hiveCompatible(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	s3ResourceName := "aws_s3_bucket.test"
	resourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_destinationTypeS3DOParquetHiveCompatible(rName),
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttrPair(resourceName, "log_destination", s3ResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "s3"),
	resource.TestCheckResourceAttr(resourceName, "log_group_name", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.file_format", "parquet"),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.hive_compatible_partitions", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}

func TestAccVPCFlowLog_LogDestinationTypeS3DOParquetHiveCompatible_perHour(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	s3ResourceName := "aws_s3_bucket.test"
	resourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_destinationTypeS3DOParquetHiveCompatiblePerHour(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttrPair(resourceName, "log_destination", s3ResourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "log_destination_type", "s3"),
	resource.TestCheckResourceAttr(resourceName, "log_group_name", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.file_format", "parquet"),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.hive_compatible_partitions", "true"),
	resource.TestCheckResourceAttr(resourceName, "destination_options.0.per_hour_partition", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCFlowLog_LogDestinationType_maxAggregationInterval(t *testing.T) {
func flowLog ec2.FlowLog
	resourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_maxAggregationInterval(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttr(resourceName, "max_aggregation_interval", "60"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCFlowLog_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var flowLog ec2.FlowLog
	resourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckFlowLogDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCFlowLogConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
func
},
	})
}


func TestAccVPCFlowLog_disappears(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_flow_log.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCFlowLogConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckFlowLogExists(ctx, resourceName, &flowLog),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceFlowLog(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}

func testAccCheckFlowLogExists(ctx context.Context, n string, v *ec2.FlowLog) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
func

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindFlowLogByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output
funcrn nil
	}
}


func testAccCheckFlowLogDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_flow_log" {
continue
	}

	_, err := tfec2.FindFlowLogByID(ctx, conn, rs.Primary.ID)
functfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("Flow Log %s still exists", rs.Primary.ID)
}

func
func

funcurn acctest.ConfigVPCWithSubnets(rName, 1)
}


func testAccVPCFlowLogConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

 "Allow",
l": {
e": [
 "ec2.${data.aws_partition.current.dns_suffix}"


 [
sumeRole"

func
func
}
funcurce "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_flow_log" "test" {
iam_role_arnws_iam_role.test.arn
log_group_name = aws_cloudwatch_log_group.test.name
traffic_typeALL"
vpc_id= aws_vpc.test.id
}
`, rName))
}


func testAccVPCFlowLogConfig_destinationTypeCloudWatchLogs(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

func{
e": [
 "ec2.${data.aws_partition.current.dns_suffix}"


funcRole"


]
}
EOF
}

resource "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_flow_log" "test" {
iam_role_arn= aws_iam_role.test.arn
log_destinationudwatch_log_group.test.arn
log_destination_type = "cloud-watch-logs"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_destinationTypeS3(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket
force_destroy = true
}

resource "aws_flow_log" "test" {
log_destinationbucket.test.arn
log_destination_type = "s3"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id

tags = {
func
}
`, rName))
}


func testAccVPCFlowLogConfig_destinationTypeS3Invalid(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_flow_log" "test" {
log_destinationdata.aws_partition.current.partition}:s3:::does-not-exist"
log_destination_type = "s3"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_destinationTypeS3DOPlainText(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket
force_destroy = true
}

resource "aws_flow_log" "test" {
log_destinationbucket.test.arn
log_destination_type = "s3"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id
destination_options {
le_format = "plain-text"
}

tags = {
me = %[1]q
}
}
`, rName))
}
func
func testAccVPCFlowLogConfig_destinationTypeS3DOPlainTextHiveCompatiblePerHour(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket
force_destroy = true
}

resource "aws_flow_log" "test" {
log_destinationbucket.test.arn
log_destination_type = "s3"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id
destination_options {
le_format = "plain-text"
ve_compatible_partitions = true
r_hour_partition= true
}

tags = {
me = %[1]q
func
`, rName))
}


func testAccVPCFlowLogConfig_destinationTypeS3DOParquet(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket
force_destroy = true
}

resource "aws_flow_log" "test" {
log_destinationbucket.test.arn
log_destination_type = "s3"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id
destination_options {
func

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_destinationTypeS3DOParquetHiveCompatible(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket
force_destroy = true
}

resource "aws_flow_log" "test" {
log_destinationbucket.test.arn
log_destination_type = "s3"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id
destination_options {
le_format = "parquet"
func

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_destinationTypeS3DOParquetHiveCompatiblePerHour(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
bucket
force_destroy = true
}

resource "aws_flow_log" "test" {
log_destinationbucket.test.arn
log_destination_type = "s3"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id
destination_options {
le_format = "parquet"
ve_compatible_partitions = true
r_hour_partition= true
func
tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_subnetID(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

 "Allow",
l": {
e": [
func

 [
sumeRole"


]
}
EOF
}

resource "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_flow_log" "test" {
iam_role_arnws_iam_role.test.arn
log_group_name = aws_cloudwatch_log_group.test.name
subnet_idnet.test[0].id
traffic_typeALL"

tags = {
me = %[1]q
}
}
func


func testAccVPCFlowLogConfig_format(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

 "Allow",
l": {
e": [
 "ec2.${data.aws_partition.current.dns_suffix}"


 [
sumeRole"


]
func
}

resource "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_s3_bucket" "test" {
bucket
force_destroy = true
}

resource "aws_flow_log" "test" {
log_destinationbucket.test.arn
log_destination_type = "s3"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id
log_format= "$${version} $${vpc-id} $${subnet-id}"

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

 "Allow",
l": {
e": [
 "ec2.${data.aws_partition.current.dns_suffix}"


funcRole"


]
}
EOF
}

resource "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_flow_log" "test" {
iam_role_arnws_iam_role.test.arn
log_group_name = aws_cloudwatch_log_group.test.name
traffic_typeALL"
vpc_id= aws_vpc.test.id

tags = {
2]q = %[3]q
}
}
`, rName, tagKey1, tagValue1))
}


func testAccVPCFlowLogConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

 "Allow",
l": {
e": [
 "ec2.${data.aws_partition.current.dns_suffix}"


 [
sumeRole"


]
}
EOF
func
resource "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_flow_log" "test" {
iam_role_arnws_iam_role.test.arn
log_group_name = aws_cloudwatch_log_group.test.name
traffic_typeALL"
vpc_id= aws_vpc.test.id

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2))
}


func testAccVPCFlowLogConfig_maxAggregationInterval(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

 "Allow",
l": {
e": [
 "ec2.${data.aws_partition.current.dns_suffix}"


 [
sumeRole"


]
}
EOF
func
resource "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_flow_log" "test" {
iam_role_arnws_iam_role.test.arn
log_group_name = aws_cloudwatch_log_group.test.name
traffic_typeALL"
vpc_id= aws_vpc.test.id

max_aggregation_interval = 60

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_transitGatewayID(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
tags = {
me = %[1]q
}
}

data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

 "Allow",
l": {
e": [
 "ec2.${data.aws_partition.current.dns_suffix}"


 [
func

]
}
EOF
}

resource "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_flow_log" "test" {
iam_role_arnaws_iam_role.test.arn
log_group_name= aws_cloudwatch_log_group.test.name
max_aggregation_interval = 60
transit_gateway_id2_transit_gateway.test.id

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_transitGatewayAttachmentID(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
tags = {
me = %[1]q
}
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
transit_gateway_id = aws_ec2_transit_gateway.test.id
vpc_idaws_vpc.test.id
subnet_ids= aws_subnet.test[*].id

tags = {
me = %[1]q
}
}

data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
name = %[1]q
funcsume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [

 "Allow",
l": {
e": [
 "ec2.${data.aws_partition.current.dns_suffix}"


 [
sumeRole"


]
}
EOF
}

resource "aws_cloudwatch_log_group" "test" {
name = %[1]q
}

resource "aws_flow_log" "test" {
iam_role_arnws_iam_role.test.arn
log_group_name = aws_cloudwatch_log_group.test.name
max_aggregation_interval
transit_gateway_attachment_id = aws_ec2_transit_gateway_vpc_attachment.test.id

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccVPCFlowLogConfig_destinationTypeKinesisFirehose(rName string) string {
	return acctest.ConfigCompose(testAccFlowLogConfig_base(rName), fmt.Sprintf(`
resource "aws_flow_log" "test" {
log_destinationesis_firehose_delivery_stream.test.arn
log_destination_type = "kinesis-data-firehose"
traffic_type= "ALL"
vpc_id= aws_vpc.test.id

tags = {
me = %[1]q
}
}

funcme
destination = "extended_s3"

extended_s3_configuration {
le_arn= aiam_role.test.arn
cket_arn = aws_s3_bucket.bucket.arn
}

tags = {
ogDeliveryEnabled" = "true"
}
}

resource "aws_s3_bucket" "bucket" {
bucket = %[1]q
}

resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
"Version":"2012-10-17",
"Statement": [

"sts:AssumeRole",
l":{
e":"firehose.amazonaws.com"

"Allow",


]
}
EOF
}

resource "aws_iam_role_policy" "test" {
name = %[1]q
role = aws_iam_role.test.id

policy = <<EOF
{
"Version":"2012-10-17",
"Statement":[

 [
reateLogDelivery",
eleteLogDelivery",
istLogDeliveries",
etLogDelivery",
se:TagDeliveryStream"

"Allow",
":"*"

]
}
EOF
}
`, rName))
func