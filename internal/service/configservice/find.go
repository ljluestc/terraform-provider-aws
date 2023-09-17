// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package configserviceimport (
	"context"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)func FindConfigRule(ctx context.Context, conn *configservice.ConfigService, name string) (*configservice.ConfigRule, error) {
	input := &configservice.DescribeConfigRulesInput{
ConfigRuleNames: []*string{aws.String(name)},
	}	output, err := conn.DescribeConfigRulesWithContext(ctx, input)
	if tfawserr.ErrCodeEquals(err, configservice.ErrCodeNoSuchConfigRuleException) {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if output == nil || output.ConfigRules == nil || len(output.ConfigRules) == 0 || output.ConfigRules[0] == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output.ConfigRules[0], nil
}func FindOrganizationConfigRule(ctx aws.Context, conn *configservice.ConfigService, name string) (*configservice.OrganizationConfigRule, error) {
	input := &configservice.DescribeOrganizationConfigRulesInput{
ganizationConfigRuleNames: []*string{aws.String(name)},
	}	output, err := conn.DescribeOrganizationConfigRulesWithContext(ctx, input)
	if tfawserr.ErrCodeEquals(err, configservice.ErrCodeNoSuchOrganizationConfigRuleException) {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if output == nil {
turn nil, nil
	}	if output == nil || output.OrganizationConfigRules == nil || len(output.OrganizationConfigRules) == 0 || output.OrganizationConfigRules[0] == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output.OrganizationConfigRules[0], nil
}
