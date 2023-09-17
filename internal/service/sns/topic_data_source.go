// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package snsimport (
	"context"
	"log"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)// @SDKDataSource("aws_sns_topic")
func DataSourceTopic() *schema.Resource {
	return &schema.Resource{
ReadWithoutTimeout: dataSourceTopicRead,Schema: map[string]*schema.Schema{
"arn": {
Type:schema.TypeString,
Computed: true,
},
"name": {
Type:schema.TypeString,
Required: true,
},	}
}func dataSourceTopicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).SNSConn(ctx)	resourceArn := ""
	name := d.Get("name").(string)	err := conn.ListTopicsPagesWithContext(ctx, &sns.ListTopicsInput{}, func(page *sns.ListTopicsOutput, lastPage bool) bool {
 page == nil {
return !lastPage
o_, topic := range page.Topics {
topicArn := aws.StringValue(topic.TopicArn)
arn, err := arn.Parse(topicArn)if err != nil {
log.Printf("[ERROR] %s", err)continue
}if arn.Resource == name {
resourceArn = topicArnbreak
}
ern !lastPage
	})	if err != nil {
turn diag.Errorf("listing SNS Topics: %s", err)
	}	if resourceArn == "" {
turn diag.Errorf("no matching SNS Topic found")
	}	d.SetId(resourceArn)
	d.Set("arn", resourceArn)	return nil
}
