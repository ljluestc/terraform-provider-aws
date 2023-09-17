// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package snsimport (
"context"
"errors""github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/sns"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
tfiam "github.com/hashicorp/terraform-provider-aws/internal/service/iam"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)func FindPlatformApplicationAttributesByARN(ctx context.Context, conn *sns.SNS, arn string) (map[string]string, error) {
input := &sns.GetPlatformApplicationAttributesInput{
PlatformApplicationArn: aws.String(arn),
}output, err := conn.GetPlatformApplicationAttributesWithContext(ctx, input)if tfawserr.ErrCodeEquals(err, sns.ErrCodeNotFoundException) {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,}if err != nil {
turn nil, err
}if output == nil || len(output.Attributes) == 0 {
turn nil, tfresource.NewEmptyResultError(input)
}return aws.StringValueMap(output.Attributes), nil
}// FindTopicAttributesByARN returns topic attributes, ensuring that any Policy field is populated with
// valid principals, i.e. the principal is either an AWS Account ID or an ARN
func FindTopicAttributesByARN(ctx context.Context, conn *sns.SNS, arn string) (map[string]string, error) {
var attributes map[string]string
err := tfresource.Retry(ctx, propagationTimeout, func() *retry.RetryError {
r err error
tributes, err = GetTopicAttributesByARN(ctx, conn, arn)
 err != nil {
return retry.NonRetryableError(err)
ad, err := tfiam.PolicyHasValidAWSPrincipals(attributes[TopicAttributeNamePolicy])
 err != nil {
return retry.NonRetryableError(err) !valid {
return retry.RetryableError(errors.New("contains invalid principals"))
ern nil
})return attributes, err
}// GetTopicAttributesByARN returns topic attributes without any validation. Any principals in a Policy field
// may contain Unique IDs instead of valid values. To ensure policies are valid, use FindTopicAttributesByARN
func GetTopicAttributesByARN(ctx context.Context, conn *sns.SNS, arn string) (map[string]string, error) {
input := &sns.GetTopicAttributesInput{
picArn: aws.String(arn),
}output, err := conn.GetTopicAttributesWithContext(ctx, input)if tfawserr.ErrCodeEquals(err, sns.ErrCodeNotFoundException) {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,}if err != nil {
turn nil, err
}if output == nil || len(output.Attributes) == 0 {
turn nil, tfresource.NewEmptyResultError(input)
}return aws.StringValueMap(output.Attributes), nil
}
