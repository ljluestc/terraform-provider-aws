---
subcategory: "CloudFront"
layout: "aws"
page_title: "AWS: aws_cloudfront_realtime_log_config"
description: |-
  Provides a CloudFront real-time log configuration resource.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_cloudfront_realtime_log_config

Provides a CloudFront real-time log configuration resource.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_cloudfront_realtime_log_config import DataAwsCloudfrontRealtimeLogConfig
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsCloudfrontRealtimeLogConfig(self, "example",
            name="example"
        )
```

## Argument Reference

This data source supports the following arguments:

* `name` - (Required) Unique name to identify this real-time log configuration.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
* `endpoint` - (Required) Amazon Kinesis data streams where real-time log data is sent.
* `fields` - (Required) Fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
* `sampling_rate` - (Required) Sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.

The `endpoint` object supports the following:

* `kinesis_stream_config` - (Required) Amazon Kinesis data stream configuration.
* `stream_type` - (Required) Type of data stream where real-time log data is sent. The only valid value is `Kinesis`.

The `kinesis_stream_config` object supports the following:

* `role_arn` - (Required) ARN of an [IAM role](iam_role.html) that CloudFront can use to send real-time log data to the Kinesis data stream.
See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-iam-role) for more information.
* `stream_arn` - (Required) ARN of the [Kinesis data stream](kinesis_stream.html).

<!-- cache-key: cdktf-0.20.8 input-6e46dcbfc5e16d8ff6f2789ee3bd99879502463156b4531588dacfee314c12ca -->