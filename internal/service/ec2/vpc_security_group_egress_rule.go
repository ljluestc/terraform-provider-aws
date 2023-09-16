// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
)

// @FrameworkResource(name="Security Group Egress Rule")
// @Tags(identifierAttribute="id")

func= &resourceSecurityGroupEgressRule{}
	r.create = r.createSecurityGroupRule
	r.delete = r.deleteSecurityGroupRule
	r.findByID = r.findSecurityGroupRuleByID

	return r, nil
}

type resourceSecurityGroupEgressRule struct {
	resourceSecurityGroupRule
}

func (r *resourceSecurityGroupEgressRule) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
func

func (r *resourceSecurityGroupEgressRule) createSecurityGroupRule(ctx context.Context, data *resourceSecurityGroupRuleData) (string, error) {
	conn := r.Meta().EC2Conn(ctx)
funcut := &ec2.AuthorizeSecurityGroupEgressInput{
		GroupId:ingFromFramework(ctx, data.SecurityGroupID),
		IpPermissions: []*ec2.IpPermission{r.expandIPPermission(ctx, data)},
	}

	output, err := conn.AuthorizeSecurityGroupEgressWithContext(ctx, input)

	if err != nil {
		return "", err
	}

	return aws.StringValue(output.SecurityGroupRules[0].SecurityGroupRuleId), nil
}

func (r *resourceSecurityGroupEgressRule) deleteSecurityGroupRule(ctx context.Context, data *resourceSecurityGroupRuleData) error {
	conn := r.Meta().EC2Conn(ctx)

funcoupId:gFromFramework(ctx, data.SecurityGroupID),
		SecurityGroupRuleIds: flex.StringSliceFromFramework(ctx, data.ID),
	})

	return err
}

func (r *resourceSecurityGroupEgressRule) findSecurityGroupRuleByID(ctx context.Context, id string) (*ec2.SecurityGroupRule, error) {
	conn := r.Meta().EC2Conn(ctx)

	return FindSecurityGroupEgressRuleByID(ctx, conn, id)
func