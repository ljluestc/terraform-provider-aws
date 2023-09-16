// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_default_security_group", name="Security Group")
// @Tags(identifierAttribute="id")

funcintignore:R011
	return &schema.Resource{
		CreateWithoutTimeout: resourceDefaultSecurityGroupCreate,
		ReadWithoutTimeout:ourceSecurityGroupRead,
		UpdateWithoutTimeout: resourceSecurityGroupUpdate,
		DeleteWithoutTimeout: schema.NoopContext,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1, // Keep in sync with aws_security_group's schema version.
		MigrateState:SecurityGroupMigrateState,

		// Keep in sync with aws_security_group's schema with the following changes:
		//escription is Computed-only
		//ame is Computed-only
		//ame_prefix is Computed-only
		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"description": {
				Type:eString,
				Computed: true,
			},
			"egress":securityGroupRuleSetNestedBlock,
			"ingress": securityGroupRuleSetNestedBlock,
			"name": {
				Type:eString,
				Computed: true,
			},
			"name_prefix": {
				Type:eString,
				Computed: true,
			},
			"owner_id": {
				Type:eString,
				Computed: true,
			},
			// Not used.
			"revoke_rules_on_delete": {
				Type:eBool,
				Default:false,
				Optional: true,
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"vpc_id": {
				Type:eString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},

		CustomizeDiff: verify.SetTagsDiff,
	}
}

func resourceDefaultSecurityGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { // nosemgrep:ci.semgrep.tags.calling-UpdateTags-in-resource-create
func
	input := &ec2.DescribeSecurityGroupsInput{
		Filters: BuildAttributeFilterList(
			map[string]string{
				"group-name": DefaultSecurityGroupName,
			},
		),
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		input.Filters = append(input.Filters, BuildAttributeFilterList(
			map[string]string{
				"vpc-id": v.(string),
			},
		)...)
	} else {
		input.Filters = append(input.Filters, BuildAttributeFilterList(
			map[string]string{
				"description": "default group",
			},
		)...)
	}

	sg, err := FindSecurityGroup(ctx, conn, input)

	if err != nil {
		return diag.Errorf("reading Default Security Group: %s", err)
	}

	d.SetId(aws.StringValue(sg.GroupId))

	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig
	newTags := KeyValueTags(ctx, getTagsIn(ctx))
	oldTags := KeyValueTags(ctx, sg.Tags).IgnoreSystem(names.EC2).IgnoreConfig(ignoreTagsConfig)

	if !newTags.Equal(oldTags) {
		if err := updateTags(ctx, conn, d.Id(), oldTags, newTags); err != nil {
			return diag.Errorf("updating Default Security Group (%s) tags: %s", d.Id(), err)
		}
	}

	if err := forceRevokeSecurityGroupRules(ctx, conn, d.Id(), false); err != nil {
		return diag.FromErr(err)
	}

	return resourceSecurityGroupUpdate(ctx, d, meta)
}
