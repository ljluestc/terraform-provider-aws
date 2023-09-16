// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	t.Parallel()

	type testCase struct {
plannedValue  types.String
currentValue  types.String
expectedValue types.String
expectErrorl
	}
	tests := map[string]testCase{
"planned name, current number (equivalent)": {
	plannedValue:  types.StringValue("icmp"),
	currentValue:  types.StringValue("1"),
	expectedValue: types.StringValue("1"),
},
"planned number, current name (equivalent)": {
	plannedValue:  types.StringValue("1"),
	currentValue:  types.StringValue("icmp"),
	expectedValue: types.StringValue("icmp"),
},
"planned name, current number (not equivalent)": {
	plannedValue:  types.StringValue("icmp"),
	currentValue:  types.StringValue("2"),
	expectedValue: types.StringValue("icmp"),
},
	}

	for name, test := range tests {
name, test := name, test
t.Run(name, 
func(t *testing.T) {
func
	request := planmodifier.StringRequest{
Path:t("test"),
PlanValue:  test.plannedValue,
StateValue: test.currentValue,
	}
	response := planmodifier.StringResponse{
PlanValue: request.PlanValue,
	}
	tfec2.NormalizeIPProtocol().PlanModifyString(ctx, request, &response)

	if !response.Diagnostics.HasError() && test.expectError {
t.Fatal("expected error, got no error")
	}

	if response.Diagnostics.HasError() && !test.expectError {
t.Fatalf("got unexpected error: %s", response.Diagnostics)
	}

	if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}
})
	}
}


func TestAccVPCSecurityGroupIngressRule_basic(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "cidr_ipv4", "10.0.0.0/8"),
funcource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "80"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "tcp"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckNoResourceAttr(resourceName, "referenced_security_group_id"),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "8080"),
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


func TestAccVPCSecurityGroupIngressRule_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	acctest.CheckFrameworkResourceDisappears(ctx, acctest.Provider, tfec2.ResourceSecurityGroupIngressRule, resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func TestAccVPCSecurityGroupIngressRule_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
Config: testAccVPCSecurityGroupIngressRuleConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_tags_computed(t *testing.T) {
func v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCSecurityGroupIngressRuleConfig_computed(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "tags.eip"),
),
	},
func
}


func TestAccVPCSecurityGroupIngressRule_DefaultTags_providerOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
func
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags2("providerkey1", "providervalue1", "providerkey2", "providervalue2"),
	testAccVPCSecurityGroupIngressRuleConfig_basic(rName),
),
Check: resource.ComposeTestCheck
functAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey2", "providervalue2"),
),
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey1", "value1"),
	testAccVPCSecurityGroupIngressRuleConfig_basic(rName),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "value1"),
func
},
	})
}


func TestAccVPCSecurityGroupIngressRule_DefaultTags_updateToProviderOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
func
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("key1", "value1"),
	testAccVPCSecurityGroupIngressRuleConfig_basic(rName),
),
Check: resource.ComposeTestCheck
functAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_DefaultTags_updateToResourceOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("key1", "value1"),
	testAccVPCSecurityGroupIngressRuleConfig_basic(rName),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
),
func
Config: testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
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

func TestAccVPCSecurityGroupIngressRule_DefaultTagsProviderAndResource_nonOverlappingTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
	testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "resourcekey1", "resourcevalue1"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
	testAccVPCSecurityGroupIngressRuleConfig_tags2(rName, "resourcekey1", "resourcevalue1", "resourcekey2", "resourcevalue2"),
),
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "3"),
	resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags.resourcekey2", "resourcevalue2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey2", "resourcevalue2"),
),
	},
funcig: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey2", "providervalue2"),
	testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "resourcekey3", "resourcevalue3"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.resourcekey3", "resourcevalue3"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey2", "providervalue2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey3", "resourcevalue3"),
),
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_DefaultTagsProviderAndResource_overlappingTag(t *testing.T) {
func v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("overlapkey1", "providervalue1"),
	testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "overlapkey1", "resourcevalue1"),
),
Check: resource.ComposeTestCheck
functAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags2("overlapkey1", "providervalue1", "overlapkey2", "providervalue2"),
func
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags.overlapkey2", "resourcevalue2"),
funcource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey2", "resourcevalue2"),
),
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("overlapkey1", "providervalue1"),
	testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "overlapkey1", "resourcevalue2"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue2"),
),
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_DefaultTagsProviderAndResource_duplicateTag(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	var v ec2.SecurityGroupRule
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("overlapkey", "overlapvalue"),
	testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "overlapkey", "overlapvalue"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
),
	},
func
}


func TestAccVPCSecurityGroupIngressRule_updateTagsKnownAtApply(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_tagsComputedFromDataSource1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
),
	},
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_tagsComputedFromDataSource2(rName, "key1", "value1", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_defaultAndIgnoreTags(t *testing.T) {
func v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	testAccCheckSecurityGroupIngressRuleUpdateTags(ctx, &v, nil, map[string]string{"defaultkey1": "defaultvalue1"}),
funcctNonEmptyPlan: true,
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultAndIgnoreTagsKeyPrefixes1("defaultkey1", "defaultvalue1", "defaultkey"),
	testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "key1", "value1"),
),
PlanOnly: true,
	},
funcig: acctest.ConfigCompose(
	acctest.ConfigDefaultAndIgnoreTagsKeys1("defaultkey1", "defaultvalue1"),
	testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, "key1", "value1"),
),
PlanOnly: true,
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_ignoreTags(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	testAccCheckSecurityGroupIngressRuleUpdateTags(ctx, &v, nil, map[string]string{"ignorekey1": "ignorevalue1"}),
),
ExpectNonEmptyPlan: true,
	},
	{
functest.ConfigIgnoreTagsKeyPrefixes1("ignorekey"),
	testAccVPCSecurityGroupIngressRuleConfig_basic(rName),
),
PlanOnly: true,
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigIgnoreTagsKeys("ignorekey1"),
	testAccVPCSecurityGroupIngressRuleConfig_basic(rName),
),
PlanOnly: true,
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_cidrIPv4(t *testing.T) {
	ctx := acctest.Context(t)
	var v1, v2 ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_cidrIPv4(rName),
Check: resource.ComposeAggregateTestCheck
func(
funcource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "cidr_ipv4", "0.0.0.0/0"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv6"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "53"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "udp"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckNoResourceAttr(resourceName, "referenced_security_group_id"),
funcource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "53"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_cidrIPv4Updated(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v2),
	testAccCheckSecurityGroupRuleNotRecreated(&v2, &v1),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "cidr_ipv4", "10.0.0.0/16"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv6"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "-1"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "1"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckNoResourceAttr(resourceName, "referenced_security_group_id"),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
func
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_cidrIPv6(t *testing.T) {
func v1, v2 ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_cidrIPv6(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v1),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv4"),
	resource.TestCheckResourceAttr(resourceName, "cidr_ipv6", "2001:db8:85a3::/64"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "80"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "tcp"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckNoResourceAttr(resourceName, "referenced_security_group_id"),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "8080"),
),
	},
	{
ResourceName:ame,
ImportState:
func
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_cidrIPv6Updated(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v2),
	testAccCheckSecurityGroupRuleNotRecreated(&v2, &v1),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv4"),
	resource.TestCheckResourceAttr(resourceName, "cidr_ipv6", "2001:db8:85a3:2::/64"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckNoResourceAttr(resourceName, "from_port"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "icmpv6"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckNoResourceAttr(resourceName, "referenced_security_group_id"),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckNoResourceAttr(resourceName, "to_port"),
),
	},
},
func


func TestAccVPCSecurityGroupIngressRule_description(t *testing.T) {
	ctx := acctest.Context(t)
	var v1, v2 ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v1),
	resource.TestCheckResourceAttr(resourceName, "description", "description1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_description(rName, "description2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v2),
	testAccCheckSecurityGroupRuleNotRecreated(&v2, &v1),
	resource.TestCheckResourceAttr(resourceName, "description", "description2"),
),
	},
},
	})
}
func
func TestAccVPCSecurityGroupIngressRule_prefixListID(t *testing.T) {
	ctx := acctest.Context(t)
	var v1, v2 ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"
	vpcEndpoint1ResourceName := "aws_vpc_endpoint.test1"
	vpcEndpoint2ResourceName := "aws_vpc_endpoint.test2"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_prefixListID(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v1),
funcource.TestCheckNoResourceAttr(resourceName, "cidr_ipv4"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv6"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "80"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "tcp"),
	resource.TestCheckResourceAttrPair(resourceName, "prefix_list_id", vpcEndpoint1ResourceName, "prefix_list_id"),
	resource.TestCheckNoResourceAttr(resourceName, "referenced_security_group_id"),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
funcource.TestCheckResourceAttr(resourceName, "to_port", "8080"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
funcig: testAccVPCSecurityGroupIngressRuleConfig_prefixListIDUpdated(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v2),
	testAccCheckSecurityGroupRuleNotRecreated(&v2, &v1),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv4"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv6"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "80"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "tcp"),
	resource.TestCheckResourceAttrPair(resourceName, "prefix_list_id", vpcEndpoint2ResourceName, "prefix_list_id"),
	resource.TestCheckNoResourceAttr(resourceName, "referenced_security_group_id"),
funcource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "8080"),
),
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_referencedSecurityGroupID(t *testing.T) {
	ctx := acctest.Context(t)
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"
	securityGroup1ResourceName := "aws_security_group.test"
	securityGroup2ResourceName := "aws_security_group.test1"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_referencedSecurityGroupID(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v1),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv4"),
funcource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "80"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "tcp"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttrPair(resourceName, "referenced_security_group_id", securityGroup1ResourceName, "id"),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "8080"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_referencedSecurityGroupIDUpdated(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v2),
	testAccCheckSecurityGroupRuleNotRecreated(&v2, &v1),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv4"),
funcource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "80"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "tcp"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckResourceAttrPair(resourceName, "referenced_security_group_id", securityGroup2ResourceName, "id"),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "8080"),
),
	},
},
	})
}


func TestAccVPCSecurityGroupIngressRule_ReferencedSecurityGroupID_peerVPC(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

funcheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckAlternateAccount(t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv4"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv6"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "80"),
funcource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestMatchResourceAttr(resourceName, "referenced_security_group_id", regexache.MustCompile("^[0-9]{12}/sg-[0-9a-z]{17}$")),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "8080"),
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


func TestAccVPCSecurityGroupIngressRule_updateSourceType(t *testing.T) {
	ctx := acctest.Context(t)
	var v1, v2 ec2.SecurityGroupRule
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_security_group_ingress_rule.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupIngressRuleDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupIngressRuleConfig_cidrIPv4(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v1),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "cidr_ipv4", "0.0.0.0/0"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv6"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "53"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "udp"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
	resource.TestCheckNoResourceAttr(resourceName, "referenced_security_group_id"),
	resource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
funck: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupIngressRuleExists(ctx, resourceName, &v2),
	testAccCheckSecurityGroupRuleRecreated(&v2, &v1),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckNoResourceAttr(resourceName, "cidr_ipv4"),
	resource.TestCheckResourceAttr(resourceName, "cidr_ipv6", "2001:db8:85a3::/64"),
	resource.TestCheckNoResourceAttr(resourceName, "description"),
	resource.TestCheckResourceAttr(resourceName, "from_port", "80"),
	resource.TestCheckResourceAttr(resourceName, "ip_protocol", "tcp"),
	resource.TestCheckNoResourceAttr(resourceName, "prefix_list_id"),
funcource.TestCheckResourceAttrSet(resourceName, "security_group_rule_id"),
	resource.TestCheckNoResourceAttr(resourceName, "tags"),
	resource.TestCheckResourceAttr(resourceName, "to_port", "8080"),
),
	},
},
	})
}


func testAccCheckSecurityGroupRuleNotRecreated(i, j *ec2.SecurityGroupRule) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if aws.StringValue(i.SecurityGroupRuleId) != aws.StringValue(j.SecurityGroupRuleId) {
	return errors.New("VPC Security Group Rule was recreated")
}

return nil
	}
}


func testAccCheckSecurityGroupRuleRecreated(i, j *ec2.SecurityGroupRule) resource.TestCheck
func {
func(s *terraform.State) error {
if aws.StringValue(i.SecurityGroupRuleId) == aws.StringValue(j.SecurityGroupRuleId) {
	return errors.New("VPC Security Group Rule was not recreated")
}

return nil
	}
}
func
func testAccCheckSecurityGroupIngressRuleDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
funcinue
	}

	_, err := tfec2.FindSecurityGroupIngressRuleByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("VPC Security Group Ingress Rule still exists: %s", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckSecurityGroupIngressRuleExists(ctx context.Context, n string, v *ec2.SecurityGroupRule) resource.TestCheck
func {
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No VPC Security Group Ingress Rule ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindSecurityGroupIngressRuleByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output

return nil
func
func
func testAccCheckSecurityGroupIngressRuleUpdateTags(ctx context.Context, v *ec2.SecurityGroupRule, oldTags, newTags map[string]string) resource.TestCheck
funcurn 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

return tfec2.UpdateTags(ctx, conn, aws.StringValue(v.SecurityGroupRuleId), oldTags, newTags)
	}
}


func testAccVPCSecurityGroupRuleConfig_base(rName string) string {
funcurce "aws_vpc" "test" {
func
  tags = {
func
}

resource "aws_security_group" "test" {
  vpc_id = aws_vpc.test.id
  name[1]q

  tags = {
me = %[1]q
  }
funcName)
func

funcurn acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  cidr_ipv410.0.0.0/8"
  from_port0
  ip_protocol = "tcp"
  to_port
}
`)
}


func testAccVPCSecurityGroupIngressRuleConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), fmt.Sprintf(`
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  cidr_ipv410.0.0.0/8"
  from_port0
  ip_protocol = "tcp"
  to_port

  tags = {
1]q = %[2]q
  }
funcagKey1, tagValue1))
func

funcurn acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), `
resource "aws_eip" "test" {
  domain = "vpc"
}

resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  cidr_ipv410.0.0.0/8"
  from_port0
  ip_protocol = "tcp"
  to_port

  tags = {
p = aws_eip.test.public_ip
  }
}
`)
}


func testAccVPCSecurityGroupIngressRuleConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), fmt.Sprintf(`
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id
funcdr_ipv410.0.0.0/8"
func_protocol = "tcp"
  to_port
funcgs = {
1]q = %[2]q
3]q = %[4]q
  }
}
`, tagKey1, tagValue1, tagKey2, tagValue2))
}

func testAccVPCSecurityGroupIngressRuleConfig_computedTagsBase(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"
  tagstags
}

data "aws_vpc" "test" {
  id = aws_vpc.test.id
}

resource "aws_security_group" "test" {
  vpc_id = aws_vpc.test.id
  name[1]q

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id
funcdr_ipv410.0.0.0/8"
  from_port0
  ip_protocol = "tcp"
  to_port

  tags = data.aws_vpc.test.tags
}
`, rName)
}


func testAccVPCSecurityGroupIngressRuleConfig_tagsComputedFromDataSource1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupIngressRuleConfig_computedTagsBase(rName), fmt.Sprintf(`
locals {
func= %[2]q
  }
}
`, tagKey1, tagValue1))
}


func testAccVPCSecurityGroupIngressRuleConfig_tagsComputedFromDataSource2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupIngressRuleConfig_computedTagsBase(rName), fmt.Sprintf(`
locals {
  tags = {
1]q = %[2]q
3]q = %[4]q
  }
}
`, tagKey1, tagValue1, tagKey2, tagValue2))
}

func testAccVPCSecurityGroupIngressRuleConfig_cidrIPv4(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  cidr_ipv40.0.0.0/0"
  from_port3
  ip_protocol = "udp"
  to_port
}
`)
}


func testAccVPCSecurityGroupIngressRuleConfig_cidrIPv4Updated(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  cidr_ipv410.0.0.0/16"
  from_port1
  ip_protocol = "1"
func
`)
}


func testAccVPCSecurityGroupIngressRuleConfig_cidrIPv6(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  cidr_ipv62001:db8:85a3::/64"
  from_port0
  ip_protocol = "tcp"
  to_port
}
`)
}


funcurn acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  cidr_ipv62001:db8:85a3:2::/64"
  ip_protocol = "icmpv6"
}
`)
}


func testAccVPCSecurityGroupIngressRuleConfig_description(rName, description string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), fmt.Sprintf(`
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  cidr_ipv410.0.0.0/8"
  from_port0
  ip_protocol = "tcp"
  to_port

  description = %[1]q
}
`, description))
}


func testAccVPCSecurityGroupIngressRuleConfig_prefixListIDBase(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), fmt.Sprintf(`
data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test1" {
  vpc_idc.test.id
  service_name = "com.amazonaws.${data.aws_region.current.name}.s3"
funcgs = {
me = %[1]q
  }
}

resource "aws_vpc_endpoint" "test2" {
  vpc_idc.test.id
  service_name = "com.amazonaws.${data.aws_region.current.name}.dynamodb"

  tags = {
me = %[1]q
func
`, rName))
}


func testAccVPCSecurityGroupIngressRuleConfig_prefixListID(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupIngressRuleConfig_prefixListIDBase(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  prefix_list_id = aws_vpc_endpoint.test1.prefix_list_id
  from_port
func_port
}
`)
}


func testAccVPCSecurityGroupIngressRuleConfig_prefixListIDUpdated(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupIngressRuleConfig_prefixListIDBase(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  prefix_list_id = aws_vpc_endpoint.test2.prefix_list_id
  from_port
  ip_protocol"tcp"
func
`)
}


func testAccVPCSecurityGroupIngressRuleConfig_referencedSecurityGroupID(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), `
resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  referenced_security_group_id = aws_security_group.test.id
  from_port
  ip_protocoltcp"
  to_port
func
}


func testAccVPCSecurityGroupIngressRuleConfig_referencedSecurityGroupIDUpdated(rName string) string {
	return acctest.ConfigCompose(testAccVPCSecurityGroupRuleConfig_base(rName), fmt.Sprintf(`
resource "aws_security_group" "test1" {
  vpc_id = aws_vpc.test.id
  name%[1]s-1"

  tags = {
me = %[1]q
  }
}
funcurce "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  referenced_security_group_id = aws_security_group.test1.id
  from_port
  ip_protocoltcp"
  to_port
}
`, rName))
}


funcurn acctest.ConfigCompose(acctest.ConfigAlternateAccountProvider(), testAccVPCSecurityGroupRuleConfig_base(rName), fmt.Sprintf(`
resource "aws_vpc" "peer" {
  provider = "awsalternate"

  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "peer" {
  provider = "awsalternate"

  vpc_id = aws_vpc.peer.id
  name[1]q
funcgs = {
me = %[1]q
  }
}

data "aws_caller_identity" "peer" {
  provider = "awsalternate"
}

# Requester's side of the connection.
resource "aws_vpc_peering_connection" "test" {
  vpc_idpc.test.id
  peer_vpc_idws_vpc.peer.id
  peer_owner_id = data.aws_caller_identity.peer.account_id
  peer_region[2]q
  auto_acceptalse

  tags = {
me = %[1]q
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_accepter" "peer" {
  provider = "awsalternate"
funcc_peering_connection_id = aws_vpc_peering_connection.test.id
  auto_accept= true

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_security_group_ingress_rule" "test" {
  security_group_id = aws_security_group.test.id

  referenced_security_group_id = "${data.aws_caller_identity.peer.account_id}/${aws_security_group.peer.id}"
  from_port
  ip_protocoltcp"
func
  depends_on = [aws_vpc_peering_connection_accepter.peer]
}
`, rName, acctest.Region()))
}
funcfuncfunc