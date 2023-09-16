// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)


func(t *testing.T) {
func
	cases := []struct {
inputterface{}
expected string
	}{
{
	input:cp",
	expected: "tcp",
},
{
	input:
	expected: "",
},
{
	input:7",
	expected: "udp",
},
{
	input:ll",
	expected: "-1",
},
{
	input:1",
	expected: "-1",
},
{
	input:,
	expected: "",
},
{
	input:",
	expected: "icmp",
},
{
	input:cmp",
	expected: "icmp",
},
{
	input:
	expected: "",
},
{
	input:cmpv6",
	expected: "icmpv6",
},
{
	input:8",
	expected: "icmpv6",
},
{
	input:,
	expected: "",
},
	}
	for _, c := range cases {
result := tfec2.ProtocolState
func(c.input)
if result != c.expected {
func
	}
}


func TestProtocolForValue(t *testing.T) {
	t.Parallel()

functring
expected string
	}{
{
	input:cp",
	expected: "tcp",
},
{
	input:",
	expected: "tcp",
},
{
	input:dp",
	expected: "udp",
},
{
	input:7",
	expected: "udp",
},
{
	input:ll",
	expected: "-1",
},
{
	input:1",
	expected: "-1",
},
{
	input:Cp",
	expected: "tcp",
},
{
	input:",
	expected: "tcp",
},
{
	input:Dp",
	expected: "udp",
},
{
	input:7",
	expected: "udp",
},
{
	input:LL",
	expected: "-1",
},
{
	input:cMp",
	expected: "icmp",
},
{
	input:",
	expected: "icmp",
},
{
	input:cMpv6",
	expected: "icmpv6",
},
{
	input:8",
	expected: "icmpv6",
},
	}

	for _, c := range cases {
result := tfec2.ProtocolForValue(c.input)
if result != c.expected {
	t.Errorf("Error matching protocol, expected (%s), got (%s)", c.expected, result)
}
	}
}


func calcSecurityGroupChecksum(rules []interface{}) int {
	sum := 0
	for _, rule := range rules {
sum += tfec2.SecurityGroupRuleHash(rule)
funcurn sum
}


func TestSecurityGroupExpandCollapseRules(t *testing.T) {
	t.Parallel()

	expected_compact_list := []interface{}{
map[string]interface{}{
funcom_port":(443),
	"to_port":
	"description": "block with description",
	"self":
	"cidr_blocks": []interface{}{
"10.0.0.1/32",
"10.0.0.2/32",
"10.0.0.3/32",
	},
},
map[string]interface{}{
	"protocol":cp",
	"from_port":(443),
	"to_port":
	"description": "block with another description",
	"self":
	"cidr_blocks": []interface{}{
"192.168.0.1/32",
"192.168.0.2/32",
	},
},
map[string]interface{}{
	"protocol":1",
	"from_port":(8000),
	"to_port":
	"description": "",
	"self":
	"ipv6_cidr_blocks": []interface{}{
"fd00::1/128",
"fd00::2/128",
	},
	"security_groups": schema.NewSet(schema.HashString, []interface{}{
"sg-11111",
"sg-22222",
"sg-33333",
	}),
},
map[string]interface{}{
	"protocol":dp",
	"from_port":(10000),
	"to_port":,
	"description": "",
	"self":
	"prefix_list_ids": []interface{}{
"pl-111111",
"pl-222222",
	},
},
	}

	expected_expanded_list := []interface{}{
map[string]interface{}{
	"protocol":cp",
	"from_port":(443),
	"to_port":
	"description": "block with description",
	"self":
},
map[string]interface{}{
	"protocol":cp",
	"from_port":(443),
	"to_port":
	"description": "block with description",
	"self":
	"cidr_blocks": []interface{}{
"10.0.0.1/32",
	},
},
map[string]interface{}{
	"protocol":cp",
	"from_port":(443),
	"to_port":
	"description": "block with description",
	"self":
	"cidr_blocks": []interface{}{
"10.0.0.2/32",
	},
},
map[string]interface{}{
	"protocol":cp",
	"from_port":(443),
	"to_port":
	"description": "block with description",
	"self":
	"cidr_blocks": []interface{}{
"10.0.0.3/32",
	},
},
map[string]interface{}{
	"protocol":cp",
	"from_port":(443),
	"to_port":
	"description": "block with another description",
	"self":
	"cidr_blocks": []interface{}{
"192.168.0.1/32",
	},
},
map[string]interface{}{
	"protocol":cp",
	"from_port":(443),
	"to_port":
	"description": "block with another description",
	"self":
	"cidr_blocks": []interface{}{
"192.168.0.2/32",
	},
},
map[string]interface{}{
	"protocol":1",
	"from_port":(8000),
	"to_port":
	"description": "",
	"self":
	"ipv6_cidr_blocks": []interface{}{
"fd00::1/128",
	},
},
map[string]interface{}{
	"protocol":1",
	"from_port":(8000),
	"to_port":
	"description": "",
	"self":
	"ipv6_cidr_blocks": []interface{}{
"fd00::2/128",
	},
},
map[string]interface{}{
	"protocol":1",
	"from_port":(8000),
	"to_port":
	"description": "",
	"self":
	"security_groups": schema.NewSet(schema.HashString, []interface{}{
"sg-11111",
	}),
},
map[string]interface{}{
	"protocol":1",
	"from_port":(8000),
	"to_port":
	"description": "",
	"self":
	"security_groups": schema.NewSet(schema.HashString, []interface{}{
"sg-22222",
	}),
},
map[string]interface{}{
	"protocol":1",
	"from_port":(8000),
	"to_port":
	"description": "",
	"self":
	"security_groups": schema.NewSet(schema.HashString, []interface{}{
"sg-33333",
	}),
},
map[string]interface{}{
	"protocol":dp",
	"from_port":(10000),
	"to_port":,
	"description": "",
	"self":
	"prefix_list_ids": []interface{}{
"pl-111111",
	},
},
map[string]interface{}{
	"protocol":dp",
	"from_port":(10000),
	"to_port":,
	"description": "",
	"self":
	"prefix_list_ids": []interface{}{
"pl-222222",
	},
},
	}

	expected_compact_set := schema.NewSet(tfec2.SecurityGroupRuleHash, expected_compact_list)
	actual_expanded_list := tfec2.SecurityGroupExpandRules(expected_compact_set).List()

	if calcSecurityGroupChecksum(expected_expanded_list) != calcSecurityGroupChecksum(actual_expanded_list) {
t.Fatalf("error matching expanded set for tfec2.SecurityGroupExpandRules()")
	}

	actual_collapsed_list := tfec2.SecurityGroupCollapseRules("ingress", expected_expanded_list)

	if calcSecurityGroupChecksum(expected_compact_list) != calcSecurityGroupChecksum(actual_collapsed_list) {
t.Fatalf("error matching collapsed set for tfec2.SecurityGroupCollapseRules()")
	}
}


func TestSecurityGroupIPPermGather(t *testing.T) {
	t.Parallel()

	raw := []*ec2.IpPermission{
{
	IpProtocol: aws.String("tcp"),
funcort:int64(-1)),
	IpRanges:ec2.IpRange{{CidrIp: aws.String("0.0.0.0/0")}},
	UserIdGroupPairs: []*ec2.UserIdGroupPair{
{
	GroupId:("sg-11111"),
	Description: aws.String("desc"),
},
	},
},
{
	IpProtocol: aws.String("tcp"),
	FromPort:.Int64(80),
	ToPort:80),
	UserIdGroupPairs: []*ec2.UserIdGroupPair{
// VPC
{
	GroupId: aws.String("sg-22222"),
},
	},
},
{
	IpProtocol: aws.String("tcp"),
	FromPort:.Int64(443),
	ToPort:443),
	UserIdGroupPairs: []*ec2.UserIdGroupPair{
{
	UserId:s.String("amazon-elb"),
	GroupId:.String("sg-d2c979d3"),
	GroupName: aws.String("amazon-elb-sg"),
},
	},
},
{
	IpProtocol: aws.String("-1"),
	FromPort:.Int64(0),
	ToPort:0),
	PrefixListIds: []*ec2.PrefixListId{
{
	PrefixListId: aws.String("pl-12345678"),
	Description:  aws.String("desc"),
},
	},
	UserIdGroupPairs: []*ec2.UserIdGroupPair{
// VPC
{
	GroupId: aws.String("sg-22222"),
},
	},
},
	}

	local := []map[string]interface{}{
{
	"protocol":cp",
	"from_port":64(1),
	"to_port":
	"cidr_blocks": []string{"0.0.0.0/0"},
	"self":
	"description": "desc",
},
{
	"protocol":  "tcp",
	"from_port": int64(80),
	"to_port":64(80),
	"security_groups": schema.NewSet(schema.HashString, []interface{}{
"sg-22222",
	}),
},
{
	"protocol":
	"from_port":,
	"to_port":int64(0),
	"prefix_list_ids": []string{"pl-12345678"},
	"security_groups": schema.NewSet(schema.HashString, []interface{}{
"sg-22222",
	}),
	"description": "desc",
},
	}

	out := tfec2.SecurityGroupIPPermGather("sg-11111", raw, aws.String("12345"))
	for _, i := range out {
// loop and match rules, because the ordering is not guarneteed
for _, l := range local {
	if i["from_port"] == l["from_port"] {
if i["to_port"] != l["to_port"] {
	t.Fatalf("to_port does not match")
}

if _, ok := i["cidr_blocks"]; ok {
	if !reflect.DeepEqual(i["cidr_blocks"], l["cidr_blocks"]) {
t.Fatalf("error matching cidr_blocks")
	}
}

if _, ok := i["security_groups"]; ok {
	outSet := i["security_groups"].(*schema.Set)
	localSet := l["security_groups"].(*schema.Set)

	if !outSet.Equal(localSet) {
t.Fatalf("Security Group sets are not equal")
	}
}
	}
}
	}
}


func TestExpandIPPerms(t *testing.T) {
	t.Parallel()

	hash := schema.HashString

	expanded := []interface{}{
map[string]interface{}{
funcom_port":
	"to_port":
	"cidr_blocks": []interface{}{"0.0.0.0/0"},
	"security_groups": schema.NewSet(hash, []interface{}{
"sg-11111",
"foo/sg-22222",
	}),
	"description": "desc",
},
map[string]interface{}{
	"protocol":  "icmp",
	"from_port": 1,
	"to_port":
	"self":
},
	}
	group := &ec2.SecurityGroup{
GroupId: aws.String("foo"),
VpcId:.String("bar"),
	}
	perms, err := tfec2.ExpandIPPerms(group, expanded)
	if err != nil {
t.Fatalf("error expanding perms: %v", err)
	}

	expected := []ec2.IpPermission{
{
	IpProtocol: aws.String("icmp"),
	FromPort:.Int64(1),
	ToPort:int64(-1)),
	IpRanges: []*ec2.IpRange{
{
	CidrIp:g("0.0.0.0/0"),
	Description: aws.String("desc"),
},
	},
	UserIdGroupPairs: []*ec2.UserIdGroupPair{
{
	UserId:g("foo"),
	GroupId:("sg-22222"),
	Description: aws.String("desc"),
},
{
	GroupId:("sg-11111"),
	Description: aws.String("desc"),
},
	},
},
{
	IpProtocol: aws.String("icmp"),
	FromPort:.Int64(1),
	ToPort:int64(-1)),
	UserIdGroupPairs: []*ec2.UserIdGroupPair{
{
	GroupId: aws.String("foo"),
},
	},
},
	}

	exp := expected[0]
	perm := perms[0]

	if aws.Int64Value(exp.FromPort) != aws.Int64Value(perm.FromPort) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.Int64Value(perm.FromPort),
	aws.Int64Value(exp.FromPort))
	}

	if aws.StringValue(exp.IpRanges[0].CidrIp) != aws.StringValue(perm.IpRanges[0].CidrIp) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.IpRanges[0].CidrIp),
	aws.StringValue(exp.IpRanges[0].CidrIp))
	}

	if aws.StringValue(exp.UserIdGroupPairs[0].UserId) != aws.StringValue(perm.UserIdGroupPairs[0].UserId) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.UserIdGroupPairs[0].UserId),
	aws.StringValue(exp.UserIdGroupPairs[0].UserId))
	}

	if aws.StringValue(exp.UserIdGroupPairs[0].GroupId) != aws.StringValue(perm.UserIdGroupPairs[0].GroupId) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.UserIdGroupPairs[0].GroupId),
	aws.StringValue(exp.UserIdGroupPairs[0].GroupId))
	}

	if aws.StringValue(exp.UserIdGroupPairs[1].GroupId) != aws.StringValue(perm.UserIdGroupPairs[1].GroupId) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.UserIdGroupPairs[1].GroupId),
	aws.StringValue(exp.UserIdGroupPairs[1].GroupId))
	}

	exp = expected[1]
	perm = perms[1]

	if aws.StringValue(exp.UserIdGroupPairs[0].GroupId) != aws.StringValue(perm.UserIdGroupPairs[0].GroupId) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.UserIdGroupPairs[0].GroupId),
	aws.StringValue(exp.UserIdGroupPairs[0].GroupId))
	}
}


func TestExpandIPPerms_NegOneProtocol(t *testing.T) {
	t.Parallel()

	hash := schema.HashString

	expanded := []interface{}{
map[string]interface{}{
	"protocol":1",
func_port":
	"cidr_blocks": []interface{}{"0.0.0.0/0"},
	"security_groups": schema.NewSet(hash, []interface{}{
"sg-11111",
"foo/sg-22222",
	}),
},
	}
	group := &ec2.SecurityGroup{
GroupId: aws.String("foo"),
VpcId:.String("bar"),
	}

	perms, err := tfec2.ExpandIPPerms(group, expanded)
	if err != nil {
t.Fatalf("error expanding perms: %v", err)
	}

	expected := []ec2.IpPermission{
{
	IpProtocol: aws.String("-1"),
	FromPort:.Int64(0),
	ToPort:0),
	IpRanges:ec2.IpRange{{CidrIp: aws.String("0.0.0.0/0")}},
	UserIdGroupPairs: []*ec2.UserIdGroupPair{
{
	UserId:  aws.String("foo"),
	GroupId: aws.String("sg-22222"),
},
{
	GroupId: aws.String("sg-11111"),
},
	},
},
	}

	exp := expected[0]
	perm := perms[0]

	if aws.Int64Value(exp.FromPort) != aws.Int64Value(perm.FromPort) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.Int64Value(perm.FromPort),
	aws.Int64Value(exp.FromPort))
	}

	if aws.StringValue(exp.IpRanges[0].CidrIp) != aws.StringValue(perm.IpRanges[0].CidrIp) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.IpRanges[0].CidrIp),
	aws.StringValue(exp.IpRanges[0].CidrIp))
	}

	if aws.StringValue(exp.UserIdGroupPairs[0].UserId) != aws.StringValue(perm.UserIdGroupPairs[0].UserId) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.UserIdGroupPairs[0].UserId),
	aws.StringValue(exp.UserIdGroupPairs[0].UserId))
	}

	// Now test the error case. This *should* error when either from_port
	// or to_port is not zero, but protocol is "-1".
	errorCase := []interface{}{
map[string]interface{}{
	"protocol":1",
	"from_port":
	"to_port":
	"cidr_blocks": []interface{}{"0.0.0.0/0"},
	"security_groups": schema.NewSet(hash, []interface{}{
"sg-11111",
"foo/sg-22222",
	}),
},
	}
	securityGroups := &ec2.SecurityGroup{
GroupId: aws.String("foo"),
VpcId:.String("bar"),
	}

	_, expandErr := tfec2.ExpandIPPerms(securityGroups, errorCase)
	if expandErr == nil {
t.Fatal("ExpandIPPerms should have errored!")
	}
}


func TestExpandIPPerms_AllProtocol(t *testing.T) {
	t.Parallel()

	hash := schema.HashString

	expanded := []interface{}{
map[string]interface{}{
	"protocol":ll",
	"from_port":
funcdr_blocks": []interface{}{"0.0.0.0/0"},
	"security_groups": schema.NewSet(hash, []interface{}{
"sg-11111",
"foo/sg-22222",
	}),
},
	}
	group := &ec2.SecurityGroup{
GroupId: aws.String("foo"),
VpcId:.String("bar"),
	}

	perms, err := tfec2.ExpandIPPerms(group, expanded)
	if err != nil {
t.Fatalf("error expanding perms: %v", err)
	}

	expected := []ec2.IpPermission{
{
	IpProtocol: aws.String("-1"),
	FromPort:.Int64(0),
	ToPort:0),
	IpRanges:ec2.IpRange{{CidrIp: aws.String("0.0.0.0/0")}},
	UserIdGroupPairs: []*ec2.UserIdGroupPair{
{
	UserId:  aws.String("foo"),
	GroupId: aws.String("sg-22222"),
},
{
	GroupId: aws.String("sg-11111"),
},
	},
},
	}

	exp := expected[0]
	perm := perms[0]

	if aws.Int64Value(exp.FromPort) != aws.Int64Value(perm.FromPort) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.Int64Value(perm.FromPort),
	aws.Int64Value(exp.FromPort))
	}

	if aws.StringValue(exp.IpRanges[0].CidrIp) != aws.StringValue(perm.IpRanges[0].CidrIp) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.IpRanges[0].CidrIp),
	aws.StringValue(exp.IpRanges[0].CidrIp))
	}

	if aws.StringValue(exp.UserIdGroupPairs[0].UserId) != aws.StringValue(perm.UserIdGroupPairs[0].UserId) {
t.Fatalf(
	"Got:\n\n%#v\n\nExpected:\n\n%#v\n",
	aws.StringValue(perm.UserIdGroupPairs[0].UserId),
	aws.StringValue(exp.UserIdGroupPairs[0].UserId))
	}

	// Now test the error case. This *should* error when either from_port
	// or to_port is not zero, but protocol is "all".
	errorCase := []interface{}{
map[string]interface{}{
	"protocol":ll",
	"from_port":
	"to_port":
	"cidr_blocks": []interface{}{"0.0.0.0/0"},
	"security_groups": schema.NewSet(hash, []interface{}{
"sg-11111",
"foo/sg-22222",
	}),
},
	}
	securityGroups := &ec2.SecurityGroup{
GroupId: aws.String("foo"),
VpcId:.String("bar"),
	}

	_, expandErr := tfec2.ExpandIPPerms(securityGroups, errorCase)
	if expandErr == nil {
t.Fatal("ExpandIPPerms should have errored!")
	}
}


func TestFlattenSecurityGroups(t *testing.T) {
	t.Parallel()

	cases := []struct {
ownerId  *string
pairs*ec2.UserIdGroupPair
expected []*tfec2.GroupIdentifier
	}{
// simple, no user id included (we ignore it mostly)
{
funcrs: []*ec2.UserIdGroupPair{
{
	GroupId: aws.String("sg-12345"),
},
	},
	expected: []*tfec2.GroupIdentifier{
{
	GroupId: aws.String("sg-12345"),
},
	},
},
{
	ownerId: aws.String("user1234"),
	pairs: []*ec2.UserIdGroupPair{
{
	GroupId: aws.String("sg-12345"),
	UserId:  aws.String("user1234"),
},
	},
	expected: []*tfec2.GroupIdentifier{
{
	GroupId: aws.String("sg-12345"),
},
	},
},
{
	ownerId: aws.String("user1234"),
	pairs: []*ec2.UserIdGroupPair{
{
	GroupId: aws.String("sg-12345"),
	UserId:  aws.String("user4321"),
},
	},
	expected: []*tfec2.GroupIdentifier{
{
	GroupId: aws.String("user4321/sg-12345"),
},
	},
},

// include description
{
	ownerId: aws.String("user1234"),
	pairs: []*ec2.UserIdGroupPair{
{
	GroupId:("sg-12345"),
	Description: aws.String("desc"),
},
	},
	expected: []*tfec2.GroupIdentifier{
{
	GroupId:("sg-12345"),
	Description: aws.String("desc"),
},
	},
},
	}

	for _, c := range cases {
out := tfec2.FlattenSecurityGroups(c.pairs, c.ownerId)
if !reflect.DeepEqual(out, c.expected) {
	t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
}
	}
}


func TestAccVPCSecurityGroup_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_name(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`security-group/.+$`)),
	resource.TestCheckResourceAttr(resourceName, "description", "Managed by Terraform"),
funcource.TestCheckResourceAttr(resourceName, "ingress.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	resource.TestCheckResourceAttr(resourceName, "name_prefix", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "revoke_rules_on_delete", "false"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_id", "aws_vpc.test", "id"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}


func TestAccVPCSecurityGroup_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceSecurityGroup(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func TestAccVPCSecurityGroup_noVPC(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_noVPC(rName),
Check: resource.ComposeAggregateTestCheck
func(
funcource.TestCheckResourceAttrPair(resourceName, "vpc_id", "data.aws_vpc.default", "id"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
func
Config:tAccVPCSecurityGroupConfig_defaultVPC(rName),
PlanOnly: true,
	},
	{
Config:tAccVPCSecurityGroupConfig_noVPC(rName),
PlanOnly: true,
	},
func
}


func TestAccVPCSecurityGroup_nameGenerated(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_nameGenerated(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	acctest.CheckResourceAttrNameGenerated(resourceName, "name"),
	resource.TestCheckResourceAttr(resourceName, "name_prefix", id.UniqueIdPrefix),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
func

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/17017

func TestAccVPCSecurityGroup_nameTerraformPrefix(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_name(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	resource.TestCheckResourceAttr(resourceName, "name_prefix", ""),
),
	},
	{
funcrtState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}

func TestAccVPCSecurityGroup_namePrefix(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_namePrefix(rName, "tf-acc-test-prefix-"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	acctest.CheckResourceAttrNameFromPrefix(resourceName, "name", "tf-acc-test-prefix-"),
	resource.TestCheckResourceAttr(resourceName, "name_prefix", "tf-acc-test-prefix-"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
},
	})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/17017

func TestAccVPCSecurityGroup_namePrefixTerraform(t *testing.T) {
func group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_namePrefix(rName, "terraform-test"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	acctest.CheckResourceAttrNameFromPrefix(resourceName, "name", "terraform-test"),
	resource.TestCheckResourceAttr(resourceName, "name_prefix", "terraform-test"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
func


func TestAccVPCSecurityGroup_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
	{
Config: testAccVPCSecurityGroupConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCSecurityGroupConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
functAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}
func
func TestAccVPCSecurityGroup_allowAll(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
func
	})
}


func TestAccVPCSecurityGroup_sourceSecurityGroup(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_sourceSecurityGroup(rName),
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
},
	})
}


func TestAccVPCSecurityGroup_ipRangeAndSecurityGroupWithSameRules(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_ipRangeAndSecurityGroupWithSameRules(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
func


func TestAccVPCSecurityGroup_ipRangesWithSameRules(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
func
Config: testAccVPCSecurityGroupConfig_ipRangesWithSameRules(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
),
	},
	{
funcrtState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}

func TestAccVPCSecurityGroup_egressMode(t *testing.T) {
	ctx := acctest.Context(t)
	var securityGroup1, securityGroup2, securityGroup3 ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkACLDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_egressModeBlocks(rName),
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &securityGroup1),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "2"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
	{
Config: testAccVPCSecurityGroupConfig_egressModeNoBlocks(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &securityGroup2),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "2"),
func
	{
Config: testAccVPCSecurityGroupConfig_egressModeZeroed(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &securityGroup3),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "0"),
),
	},
},
	})
}


func TestAccVPCSecurityGroup_ingressMode(t *testing.T) {
func securityGroup1, securityGroup2, securityGroup3 ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckNetworkACLDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_ingressModeBlocks(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &securityGroup1),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "2"),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
	{
Config: testAccVPCSecurityGroupConfig_ingressModeNoBlocks(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &securityGroup2),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "2"),
),
func
Config: testAccVPCSecurityGroupConfig_ingressModeZeroed(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &securityGroup3),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "0"),
),
	},
func
}


func TestAccVPCSecurityGroup_ruleGathering(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_ruleGathering(rName),
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "3"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"description": for all ipv6",
"from_port": "0",
func6_cidr_blocks.0": "::/0",
"prefix_list_ids.#":  "0",
"protocol":  "-1",
"security_groups.#":  "0",
"self":"false",
"to_port":,
	}),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":0",
"description": for all ipv4",
"from_port": "0",
"ipv6_cidr_blocks.#": "0",
"prefix_list_ids.#":  "0",
funcurity_groups.#":  "0",
"self":"false",
"to_port":,
	}),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "5"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":0.0/16",
funcm_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":",
	}),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"description":s from all ipv6",
func6_cidr_blocks.#": "1",
"ipv6_cidr_blocks.0": "::/0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":",
	}),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
funcr_blocks.0":/24",
"cidr_blocks.1":/24",
"description":s from 10.0.0.0/16",
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
func
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":/24",
"cidr_blocks.1":/24",
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"true",
"to_port":",
	}),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}

// This test should fail to destroy the Security Groups and VPC, due to a
// dependency cycle added outside of terraform's management. There is a sweeper
// 'aws_vpc' and 'aws_security_group' that cleans these up, however, the test is
// written to allow Terraform to clean it up because we do go and revoke the
// cyclic rules that were added.

func TestAccVPCSecurityGroup_forceRevokeRulesTrue(t *testing.T) {
	ctx := acctest.Context(t)
	var primary ec2.SecurityGroup
	var secondary ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.primary"
	resourceName2 := "aws_security_group.secondary"

	// Add rules to create a cycle between primary and secondary. This prevents
	// Terraform/AWS from being able to destroy the groups
	testAddCycle := testAddRuleCycle(ctx, &primary, &secondary)
	// Remove the rules that created the cycle; Terraform/AWS can now destroy them
	testRemoveCycle := testRemoveRuleCycle(ctx, &primary, &secondary)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	// create the configuration with 2 security groups, then create a
	// dependency cycle such that they cannot be deleted
	{
Config: testAccVPCSecurityGroupConfig_revokeBase(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &primary),
	testAccCheckSecurityGroupExists(ctx, resourceName2, &secondary),
	testAddCycle,
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
	// Verify the DependencyViolation error by using a configuration with the
	// groups removed. Terraform tries to destroy them but cannot. Expect a
	// DependencyViolation error
	{
Config:CSecurityGroupConfig_revokeBaseRemoved(rName),
ExpectError: regexache.MustCompile("DependencyViolation"),
	},
	// Restore the config (a no-op plan) but also remove the dependencies
	// between the groups with testRemoveCycle
	{
Config: testAccVPCSecurityGroupConfig_revokeBase(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &primary),
	testAccCheckSecurityGroupExists(ctx, resourceName2, &secondary),
	testRemoveCycle,
),
	},
	// Again try to apply the config with the sgs removed; it should work
	{
Config: testAccVPCSecurityGroupConfig_revokeBaseRemoved(rName),
	},
	////
	// now test with revoke_rules_on_delete
funccreate the configuration with 2 security groups, then create a
	// dependency cycle such that they cannot be deleted. In this
	// configuration, each Security Group has `revoke_rules_on_delete`
	// specified, and should delete with no issue
	{
Config: testAccVPCSecurityGroupConfig_revokeTrue(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &primary),
	testAccCheckSecurityGroupExists(ctx, resourceName2, &secondary),
	testAddCycle,
),
	},
	// Again try to apply the config with the sgs removed; it should work,
	// because we've told the SGs to forcefully revoke their rules first
	{
func
},
	})
}


func TestAccVPCSecurityGroup_forceRevokeRulesFalse(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
func
	var primary ec2.SecurityGroup
	var secondary ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.primary"
	resourceName2 := "aws_security_group.secondary"

	// Add rules to create a cycle between primary and secondary. This prevents
	// Terraform/AWS from being able to destroy the groups
	testAddCycle := testAddRuleCycle(ctx, &primary, &secondary)
	// Remove the rules that created the cycle; Terraform/AWS can now destroy them
	testRemoveCycle := testRemoveRuleCycle(ctx, &primary, &secondary)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	// create the configuration with 2 security groups, then create a
	// dependency cycle such that they cannot be deleted. These Security
	// Groups are configured to explicitly not revoke rules on delete,
	// `revoke_rules_on_delete = false`
funcig: testAccVPCSecurityGroupConfig_revokeFalse(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &primary),
	testAccCheckSecurityGroupExists(ctx, resourceName2, &secondary),
	testAddCycle,
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
	// Verify the DependencyViolation error by using a configuration with the
	// groups removed, and the Groups not configured to revoke their ruls.
	// Terraform tries to destroy them but cannot. Expect a
	// DependencyViolation error
	{
Config:CSecurityGroupConfig_revokeBaseRemoved(rName),
func
	// Restore the config (a no-op plan) but also remove the dependencies
	// between the groups with testRemoveCycle
	{
Config: testAccVPCSecurityGroupConfig_revokeFalse(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &primary),
	testAccCheckSecurityGroupExists(ctx, resourceName2, &secondary),
	testRemoveCycle,
),
	},
	// Again try to apply the config with the sgs removed; it should work
	{
Config: testAccVPCSecurityGroupConfig_revokeBaseRemoved(rName),
	},
func
}


func TestAccVPCSecurityGroup_change(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_basic(rName),
Check: resource.ComposeTestCheck
functAccCheckSecurityGroupExists(ctx, resourceName, &group),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
	{
Config: testAccVPCSecurityGroupConfig_changed(rName),
Check: resource.ComposeTestCheck
functAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":/8",
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "2"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":/8",
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
funcr_blocks.#":
"cidr_blocks.0":0",
"cidr_blocks.1":/8",
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
),
	},
},
	})
func

func TestAccVPCSecurityGroup_ipv6(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_ipv6(rName),
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "name", rName),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "1",
"ipv6_cidr_blocks.0": "::/0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
func
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "1",
"ipv6_cidr_blocks.0": "::/0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}


func TestAccVPCSecurityGroup_self(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

	checkSelf := 
func(s *terraform.State) (err error) {
if len(group.IpPermissions) > 0 &&
	len(group.IpPermissions[0].UserIdGroupPairs) > 0 &&
	aws.StringValue(group.IpPermissions[0].UserIdGroupPairs[0].GroupId) == aws.StringValue(group.GroupId) {
	return nil
}

return fmt.Errorf("Security Group does not contain \"self\" rule: %#v", group)
	}

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_self(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"protocol":  "tcp",
"from_port": "80",
"to_port":00",
"self":
	}),
	checkSelf,
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}


func TestAccVPCSecurityGroup_vpc(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_vpc(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"protocol":
"from_port":
"to_port":
"cidr_blocks.#": "1",
"cidr_blocks.0": "10.0.0.0/8",
	}),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
funcm_port":
"to_port":
"cidr_blocks.#": "1",
"cidr_blocks.0": "10.0.0.0/8",
	}),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_id", "aws_vpc.test", "id"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}


func TestAccVPCSecurityGroup_vpcNegOneIngress(t *testing.T) {
func group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_vpcNegativeOneIngress(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"protocol":
"from_port":
"to_port":
"cidr_blocks.#": "1",
"cidr_blocks.0": "10.0.0.0/8",
	}),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_id", "aws_vpc.test", "id"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
func
	})
}


func TestAccVPCSecurityGroup_vpcProtoNumIngress(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
funcourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
func
Config: testAccVPCSecurityGroupConfig_vpcProtocolNumberIngress(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"protocol":
"from_port":
"to_port":
"cidr_blocks.#": "1",
"cidr_blocks.0": "10.0.0.0/8",
	}),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_id", "aws_vpc.test", "id"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}


func TestAccVPCSecurityGroup_multiIngress(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test1"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}


func TestAccVPCSecurityGroup_vpcAllEgress(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_vpcAllEgress(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"protocol":
"from_port":
"to_port":
"cidr_blocks.#": "1",
"cidr_blocks.0": "10.0.0.0/8",
	}),
func
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
func
}


func TestAccVPCSecurityGroup_ruleDescription(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_ruleDescription(rName, "Egress description", "Ingress description"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
funcr_blocks.0":/8",
"description": description",
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"prefix_list_ids.#":  "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
func
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":/8",
"description":s description",
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
funcurity_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
	// Change just the rule descriptions.
	{
Config: testAccVPCSecurityGroupConfig_ruleDescription(rName, "New egress description", "New ingress description"),
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":/8",
"description":ress description",
"from_port": "80",
funcfix_list_ids.#":  "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
funcr_blocks.0":/8",
"description":gress description",
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
),
	},
	// Remove just the rule descriptions.
	{
Config: testAccVPCSecurityGroupConfig_emptyRuleDescription(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":8",
"description":
"from_port":"80",
"protocol": "tcp",
funcf":
"to_port":  "8000",
	}),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":8",
"description":
functocol": "tcp",
"security_groups.#": "0",
"self":
"to_port":  "8000",
	}),
),
	},
},
func


func TestAccVPCSecurityGroup_defaultEgressVPC(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_defaultEgress(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "0"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}

// Testing drift detection with groups containing the same port and types

func TestAccVPCSecurityGroup_driftComplex(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_security_group.test1"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_driftComplex(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "3"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":/8",
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"prefix_list_ids.#":  "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":0/8",
"description":
"from_port": "80",
funcfix_list_ids.#":  "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "3"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":/8",
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":0/8",
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
),
	},
funcurceName: resourceName,
ImportState:  true,
// In rules with cidr_block drift, import only creates a single ingress
// rule with the cidr_blocks de-normalized. During subsequent apply, its
// normalized to create the 2 ingress rules seen in checks above.
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete", "ingress", "egress"},
	},
func
}


func TestAccVPCSecurityGroup_invalidCIDRBlock(t *testing.T) {
	ctx := acctest.Context(t)
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config:CSecurityGroupConfig_invalidIngressCIDR,
ExpectError: regexache.MustCompile("invalid CIDR address: 1.2.3.4/33"),
	},
	{
Config:CSecurityGroupConfig_invalidEgressCIDR,
ExpectError: regexache.MustCompile("invalid CIDR address: 1.2.3.4/33"),
	},
	{
Config:CSecurityGroupConfig_invalidIPv6IngressCIDR,
ExpectError: regexache.MustCompile("invalid CIDR address: ::/244"),
	},
	{
Config:CSecurityGroupConfig_invalidIPv6EgressCIDR,
func
},
	})
}


func TestAccVPCSecurityGroup_cidrAndGroups(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_security_group.test1"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_combinedCIDRAndGroups(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}


func TestAccVPCSecurityGroup_ingressWithCIDRAndSGsVPC(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test1"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_ingressWithCIDRAndSGs(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":/8",
"description":
"from_port": "80",
"ipv6_cidr_blocks.#": "0",
"prefix_list_ids.#":  "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":00",
	}),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "2"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ingress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":0.1/32",
"description":
"from_port": "22",
"ipv6_cidr_blocks.#": "0",
"protocol":  "tcp",
"security_groups.#":  "0",
"self":"false",
"to_port":",
	}),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
func


func TestAccVPCSecurityGroup_egressWithPrefixList(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_prefixListEgress(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete"},
	},
},
	})
}


func := acctest.Context(t)
	var group ec2.SecurityGroup
	resourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_prefixListIngress(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
func
},
	})
}


func TestAccVPCSecurityGroup_ipv4AndIPv6Egress(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_security_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCSecurityGroupConfig_ipv4andIPv6Egress(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "2"),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"cidr_blocks.0":0",
"description":
"from_port": "0",
"ipv6_cidr_blocks.#": "0",
"prefix_list_ids.#":  "0",
"protocol":  "-1",
"security_groups.#":  "0",
"self":"false",
"to_port":,
	}),
	resource.TestCheckTypeSetElemNestedAttrs(resourceName, "egress.*", map[string]string{
"cidr_blocks.#":
"description":
"from_port": "0",
"ipv6_cidr_blocks.#": "1",
"ipv6_cidr_blocks.0": "::/0",
"prefix_list_ids.#":  "0",
"protocol":  "-1",
"security_groups.#":  "0",
"self":"false",
"to_port":,
	}),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "0"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"revoke_rules_on_delete", "egress"},
	},
func
}


func TestAccVPCSecurityGroup_failWithDiffMismatch(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCSecurityGroupConfig_failWithDiffMismatch(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "2"),
),
	},
},
	})
}

var ruleLimit int

// testAccSecurityGroup_ruleLimit sets the global "ruleLimit" and is only called once
// but does not run in parallel slowing down tests. It cannot run in parallel since
func
func testAccSecurityGroup_ruleLimit(t *testing.T) {
	ctx := acctest.Context(t)
	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	// get limit
	{
Config: testAccVPCSecurityGroupConfig_getLimit(),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupRuleLimit("data.aws_servicequotas_service_quota.test", &ruleLimit),
),
func
},
	})
}


func TestAccVPCSecurityGroup_RuleLimit_exceededAppend(t *testing.T) {
	ctx := acctest.Context(t)
	if ruleLimit == 0 {
testAccSecurityGroup_ruleLimit(t)
	}

	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	// create a valid SG just under the limit
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	testAccCheckSecurityGroupRuleCount(ctx, &group, 0, ruleLimit),
	resource.TestCheckResourceAttr(resourceName, "egress.#", strconv.Itoa(ruleLimit)),
),
	},
	// append a rule to step over the limit
funcig:CSecurityGroupConfig_ruleLimit(rName, 0, ruleLimit+1),
ExpectError: regexache.MustCompile("RulesPerSecurityGroupLimitExceeded"),
	},
	{
PreConfig: 
func() {
	// should have the original rules still
	err := testSecurityGroupRuleCount(ctx, aws.StringValue(group.GroupId), 0, ruleLimit)
	if err != nil {
t.Fatalf("PreConfig check failed: %s", err)
	}
},
// running the original config again now should restore the rules
Config: testAccVPCSecurityGroupConfig_ruleLimit(rName, 0, ruleLimit),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	testAccCheckSecurityGroupRuleCount(ctx, &group, 0, ruleLimit),
	resource.TestCheckResourceAttr(resourceName, "egress.#", strconv.Itoa(ruleLimit)),
),
	},
},
	})
}


func TestAccVPCSecurityGroup_RuleLimit_cidrBlockExceededAppend(t *testing.T) {
	ctx := acctest.Context(t)
	if ruleLimit == 0 {
testAccSecurityGroup_ruleLimit(t)
	}

	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	// create a valid SG just under the limit
	{
Config: testAccVPCSecurityGroupConfig_cidrBlockRuleLimit(rName, 0, ruleLimit),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	testAccCheckSecurityGroupRuleCount(ctx, &group, 0, 1),
func
	// append a rule to step over the limit
	{
Config:CSecurityGroupConfig_cidrBlockRuleLimit(rName, 0, ruleLimit+1),
ExpectError: regexache.MustCompile("RulesPerSecurityGroupLimitExceeded"),
	},
	{
PreConfig: 
funcshould have the original cidr blocks still in 1 rule
	err := testSecurityGroupRuleCount(ctx, aws.StringValue(group.GroupId), 0, 1)
	if err != nil {
t.Fatalf("PreConfig check failed: %s", err)
	}

	id := aws.StringValue(group.GroupId)

	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	match, err := tfec2.FindSecurityGroupByID(ctx, conn, id)
	if tfresource.NotFound(err) {
t.Fatalf("PreConfig check failed: Security Group (%s) not found: %s", id, err)
	}
	if err != nil {
t.Fatalf("PreConfig check failed: %s", err)
func
	if cidrCount := len(match.IpPermissionsEgress[0].IpRanges); cidrCount != ruleLimit {
t.Fatalf("PreConfig check failed: rule does not have previous IP ranges, has %d", cidrCount)
	}
funcunning the original config again now should restore the rules
Config: testAccVPCSecurityGroupConfig_cidrBlockRuleLimit(rName, 0, ruleLimit),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	testAccCheckSecurityGroupRuleCount(ctx, &group, 0, 1),
),
	},
},
func


func TestAccVPCSecurityGroup_RuleLimit_exceededPrepend(t *testing.T) {
	ctx := acctest.Context(t)
	if ruleLimit == 0 {
testAccSecurityGroup_ruleLimit(t)
	}

	var group ec2.SecurityGroup
funcourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	// create a valid SG just under the limit
	{
Config: testAccVPCSecurityGroupConfig_ruleLimit(rName, 0, ruleLimit),
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	testAccCheckSecurityGroupRuleCount(ctx, &group, 0, ruleLimit),
),
	},
	// prepend a rule to step over the limit
	{
Config:CSecurityGroupConfig_ruleLimit(rName, 1, ruleLimit+1),
ExpectError: regexache.MustCompile("RulesPerSecurityGroupLimitExceeded"),
func
PreConfig: 
func() {
	// should have the original rules still (limit - 1 because of the shift)
	err := testSecurityGroupRuleCount(ctx, aws.StringValue(group.GroupId), 0, ruleLimit-1)
	if err != nil {
t.Fatalf("PreConfig check failed: %s", err)
	}
},
// running the original config again now should restore the rules
Config: testAccVPCSecurityGroupConfig_ruleLimit(rName, 0, ruleLimit),
Check: resource.ComposeTestCheck
func(
functAccCheckSecurityGroupRuleCount(ctx, &group, 0, ruleLimit),
),
	},
},
	})
}


func TestAccVPCSecurityGroup_RuleLimit_exceededAllNew(t *testing.T) {
	ctx := acctest.Context(t)
funcAccSecurityGroup_ruleLimit(t)
	}

	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	// create a valid SG just under the limit
	{
Config: testAccVPCSecurityGroupConfig_ruleLimit(rName, 0, ruleLimit),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	testAccCheckSecurityGroupRuleCount(ctx, &group, 0, ruleLimit),
),
	},
	// add a rule to step over the limit with entirely new rules
funcig:CSecurityGroupConfig_ruleLimit(rName, 100, ruleLimit+1),
ExpectError: regexache.MustCompile("RulesPerSecurityGroupLimitExceeded"),
	},
	{
// all the rules should have been revoked and the add failed
PreConfig: 
func() {
	err := testSecurityGroupRuleCount(ctx, aws.StringValue(group.GroupId), 0, 0)
	if err != nil {
func
},
// running the original config again now should restore the rules
Config: testAccVPCSecurityGroupConfig_ruleLimit(rName, 0, ruleLimit),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	testAccCheckSecurityGroupRuleCount(ctx, &group, 0, ruleLimit),
),
	},
},
	})
func

func TestAccVPCSecurityGroup_rulesDropOnError(t *testing.T) {
	ctx := acctest.Context(t)
	var group ec2.SecurityGroup
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_security_group.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
	// Create a valid security group with some rules and make sure it exists
	{
Config: testAccVPCSecurityGroupConfig_rulesDropOnErrorInit(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
),
	},
	// Add a bad rule to trigger API error
	{
Config:CSecurityGroupConfig_rulesDropOnErrorAddBadRule(rName),
func
	// All originally added rules must survive. This will return non-empty plan if anything changed.
	{
Config:tAccVPCSecurityGroupConfig_rulesDropOnErrorInit(rName),
PlanOnly: true,
	},
},
	})
}

func problem seen in EMR and other services. The main gist is that a security
// group can have 0 rules and still have dependencies. Services, like EMR,
// create rules in security groups. If a 0-rule SG is listed as the source of
// a rule in another SG, it could not previously be deleted.

func TestAccVPCSecurityGroup_emrDependencyViolation(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
t.Skip("skipping long-running test in short mode")
	}

	var group ec2.SecurityGroup
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSecurityGroupDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCSecurityGroupConfig_emrLinkedRulesDestroy(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSecurityGroupExists(ctx, resourceName, &group),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`security-group/.+$`)),
	resource.TestCheckResourceAttr(resourceName, "egress.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "ingress.#", "1"),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "revoke_rules_on_delete", "true"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_id", "aws_vpc.test", "id"),
),
func
},
	})
}

// cycleIPPermForGroup returns an IpPermission struct with a configured
// UserIdGroupPair for the groupid given. Used in
// TestAccAWSSecurityGroup_forceRevokeRules_should_fail to create a cyclic rule
// between 2 security groups

func perm ec2.IpPermission
	perm.FromPort = aws.Int64(0)
	perm.ToPort = aws.Int64(0)
	perm.IpProtocol = aws.String("icmp")
	perm.UserIdGroupPairs = make([]*ec2.UserIdGroupPair, 1)
	perm.UserIdGroupPairs[0] = &ec2.UserIdGroupPair{
GroupId: aws.String(groupId),
	}
	return &perm
}
funcestAddRuleCycle returns a TestCheck
func to use at the end of a test, such
// that a Security Group Rule cyclic dependency will be created between the two
// Security Groups. A companion 
function, testRemoveRuleCycle, will undo this.

func testAddRuleCycle(ctx context.Context, primary, secondary *ec2.SecurityGroup) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if primary.GroupId == nil {
	return fmt.Errorf("Primary SG not set for TestAccAWSSecurityGroup_forceRevokeRules_should_fail")
funcecondary.GroupId == nil {
	return fmt.Errorf("Secondary SG not set for TestAccAWSSecurityGroup_forceRevokeRules_should_fail")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

// cycle from primary to secondary
perm1 := cycleIPPermForGroup(aws.StringValue(secondary.GroupId))
// cycle from secondary to primary
func
req1 := &ec2.AuthorizeSecurityGroupEgressInput{
	GroupId:GroupId,
	IpPermissions: []*ec2.IpPermission{perm1},
}
req2 := &ec2.AuthorizeSecurityGroupEgressInput{
	GroupId:y.GroupId,
	IpPermissions: []*ec2.IpPermission{perm2},
}

var err error
_, err = conn.AuthorizeSecurityGroupEgressWithContext(ctx, req1)
if err != nil {
func
_, err = conn.AuthorizeSecurityGroupEgressWithContext(ctx, req2)
if err != nil {
	return fmt.Errorf("Error authorizing secondary security group %s rules: %w", aws.StringValue(secondary.GroupId), err)
}
return nil
	}
}

funchat was added in testAddRuleCycle

func testRemoveRuleCycle(ctx context.Context, primary, secondary *ec2.SecurityGroup) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if primary.GroupId == nil {
	return fmt.Errorf("Primary SG not set for TestAccAWSSecurityGroup_forceRevokeRules_should_fail")
}
if secondary.GroupId == nil {
func

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
for _, sg := range []*ec2.SecurityGroup{primary, secondary} {
	var err error
	if sg.IpPermissions != nil {
req := &ec2.RevokeSecurityGroupIngressInput{
	GroupId:Id,
func

if _, err = conn.RevokeSecurityGroupIngressWithContext(ctx, req); err != nil {
	return fmt.Errorf("Error revoking default ingress rule for Security Group in testRemoveCycle (%s): %w", aws.StringValue(primary.GroupId), err)
}
	}

	if sg.IpPermissionsEgress != nil {
req := &ec2.RevokeSecurityGroupEgressInput{
funcermissions: sg.IpPermissionsEgress,
}

if _, err = conn.RevokeSecurityGroupEgressWithContext(ctx, req); err != nil {
	return fmt.Errorf("Error revoking default egress rule for Security Group in testRemoveCycle (%s): %w", aws.StringValue(sg.GroupId), err)
}
	}
}
return nil
	}
}


func testAccCheckSecurityGroupDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_security_group" {
continue
	}

func
	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return create.Error(names.EC2, create.ErrActionCheckingDestroyed, "Security Group", rs.Primary.ID, err)
	}

	return fmt.Errorf("VPC Security Group (%s) still exists.", rs.Primary.ID)
}

func
}


func testAccCheckSecurityGroupExists(ctx context.Context, n string, v *ec2.SecurityGroup) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
funcok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No VPC Security Group ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindSecurityGroupByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return create.Error(names.EC2, create.ErrActionCheckingExistence, "Security Group", rs.Primary.ID, err)
}

*v = *output

return nil
	}
}
func
func testAccCheckSecurityGroupRuleLimit(n string, v *int) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No Service Quotas ID is set")
}
funct, err := strconv.Atoi(rs.Primary.Attributes["value"])
if err != nil {
	return fmt.Errorf("converting value to int: %s", err)
func
*v = limit
funcrn nil
func

func testAccCheckSecurityGroupRuleCount(ctx context.Context, group *ec2.SecurityGroup, expectedIngressCount, expectedEgressCount int) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
id := aws.StringValue(group.GroupId)
return testSecurityGroupRuleCount(ctx, id, expectedIngressCount, expectedEgressCount)
	}
}


func testSecurityGroupRuleCount(ctx context.Context, id string, expectedIngressCount, expectedEgressCount int) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	group, err := tfec2.FindSecurityGroupByID(ctx, conn, id)
	if tfresource.NotFound(err) {
return fmt.Errorf("Security Group (%s) not found: %w", id, err)
	}
	if err != nil {
return create.Error(names.EC2, create.ErrActionChecking, "Security Group", id, err)
	}

	if actual := len(group.IpPermissions); actual != expectedIngressCount {
return fmt.Errorf("Security group ingress rule count %d does not match %d", actual, expectedIngressCount)
	}

	if actual := len(group.IpPermissionsEgress); actual != expectedEgressCount {
return fmt.Errorf("Security group egress rule count %d does not match %d", actual, expectedEgressCount)
	}

	return nil
}


func testAccVPCSecurityGroupConfig_name(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
func
funcurce "aws_security_group" "test" {
  name[1]q
func
`, rName)
}


func testAccVPCSecurityGroupConfig_nameGenerated(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  vpc_id = aws_vpc.test.id
}
`, rName)
}


func testAccVPCSecurityGroupConfig_namePrefix(rName, namePrefix string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name_prefix = %[2]q
  vpc_id.test.id
}
`, rName, namePrefix)
}
func
funcurn fmt.Sprintf(`
resource "aws_security_group" "test" {
func

data "aws_vpc" "default" {
  default = true
}
`, rName)
}


func testAccVPCSecurityGroupConfig_defaultVPC(rName string) string {
	return fmt.Sprintf(`
resource "aws_security_group" "test" {
  name[1]q
  vpc_id = data.aws_vpc.default.id
}

data "aws_vpc" "default" {
  default = true
}
`, rName)
}


func testAccVPCSecurityGroupConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
func
func %[1]q
  }
func
resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}


func testAccVPCSecurityGroupConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
func
func= %[3]q
4]q = %[5]q
func
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}


func testAccVPCSecurityGroupConfig_getLimit() string {
	return `
data "aws_servicequotas_service_quota" "test" {
  quota_nameInbound or outbound rules per security group"
  service_code = "vpc"
}
`
}


func testAccVPCSecurityGroupConfig_ruleLimit(rName string, egressStartIndex, egressRulesCount int) string {
	var egressRules strings.Builder
	for i := egressStartIndex; i < egressRulesCount+egressStartIndex; i++ {
fmt.Fprintf(&egressRules, `
  egress {
otocol = p"
om_port= "0 + %[1]d}"
funclocks = ["${cidrhost("10.1.0.0/16", %[1]d)}/32"]
func)
	}
funcurce "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}
funcurce "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  # egress rules to exhaust the limit
  %[2]s
}
`, rName, egressRules.String())
}


func testAccVPCSecurityGroupConfig_cidrBlockRuleLimit(rName string, egressStartIndex, egressRulesCount int) string {
	var cidrBlocks strings.Builder
	for i := egressStartIndex; i < egressRulesCount+egressStartIndex; i++ {
fmt.Fprintf(&cidrBlocks, `
"${cidrhost("10.1.0.0/16", %[1]d)}/32",
`, i)
	}

funcurce "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  egress {
otocol  = "tcp"
funct= "
cidr_blocks to exhaust the limit
dr_blocks = [
%[2]s

  }
}
`, rName, cidrBlocks.String())
}


func testAccVPCSecurityGroupConfig_emptyRuleDescription(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
func
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol = 
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
scription = ""
  }

  egress {
otocol = p"
om_port= 8
funclocks = ["10.0.0.0/8"]
scription = ""
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_ipv6(rName string) string {
funcurce "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol= "6"
funct = 8000
v6_cidr_blocks = ["::/0"]
  }

  egress {
otocol= "tcp"
om_port
_port = 8000
v6_cidr_blocks = ["::/0"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol = 
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}
func
func testAccVPCSecurityGroupConfig_revokeBaseRemoved(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}
func


func testAccVPCSecurityGroupConfig_revokeBase(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "primary" {
  name%[1]s-primary"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  timeouts {
lete = "2m"
  }
}

resource "aws_security_group" "secondary" {
  name%[1]s-secondary"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  timeouts {
lete = "2m"
func
`, rName)
}


func testAccVPCSecurityGroupConfig_revokeFalse(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "primary" {
  name%[1]s-primary"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  revoke_rules_on_delete = false
}

resource "aws_security_group" "secondary" {
  name%[1]s-secondary"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  revoke_rules_on_delete = false
}
`, rName)
}

func testAccVPCSecurityGroupConfig_revokeTrue(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "primary" {
  name%[1]s-primary"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  revoke_rules_on_delete = true
}

resource "aws_security_group" "secondary" {
  name%[1]s-secondary"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  revoke_rules_on_delete = true
}
`, rName)
}


func testAccVPCSecurityGroupConfig_changed(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["0.0.0.0/0", "10.0.0.0/8"]
  }

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
func

func testAccVPCSecurityGroupConfig_ruleDescription(rName, egressDescription, ingressDescription string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol = 
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
scription = %[2]q
  }

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
func

  tags = {
me = %[1]q
  }
}
`, rName, ingressDescription, egressDescription)
}


func testAccVPCSecurityGroupConfig_self(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol  = "tcp"
om_port = 80
_port= 8
lf
  }

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_vpc(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
func

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_vpcNegativeOneIngress(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
func

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  ingress {
otocol = "
om_port= 0
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_vpcProtocolNumberIngress(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

funcol = "
om_port= 0
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_multiIngress(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test1" {
  name%[1]s-1"
  vpc_id = aws_vpc.test.id

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
func
}

resource "aws_security_group" "test2" {
  name%[1]s-2"
  vpc_id = aws_vpc.test.id

  ingress {
otocol = p"
om_port= 2
_port
dr_blocks = ["10.0.0.0/8"]
  }

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  ingress {
otocol
om_port
_port= 8000
curity_groups = [aws_security_group.test1.id]
  }

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
funcName)
}


func testAccVPCSecurityGroupConfig_vpcAllEgress(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  egress {
otocol = l"
om_port= 0
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_defaultEgress(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_driftComplex(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test2" {
funcc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test1" {
  name%[1]s-1"
  vpc_id = aws_vpc.test.id

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["206.0.0.0/8"]
  }

  ingress {
otocol
om_port
_port= 22
func

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["206.0.0.0/8"]
  }

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }

  egress {
otocol
om_port
_port= 22
curity_groups = [aws_security_group.test2.id]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}
funct testAccVPCSecurityGroupConfig_invalidIngressCIDR = `
resource "aws_security_group" "test" {
  ingress {
om_port= 0
_port
otocol = "
dr_blocks = ["1.2.3.4/33"]
  }
}
`

const testAccVPCSecurityGroupConfig_invalidEgressCIDR = `
resource "aws_security_group" "test" {
  egress {
om_port= 0
_port
otocol = "
dr_blocks = ["1.2.3.4/33"]
  }
}
`

const testAccVPCSecurityGroupConfig_invalidIPv6IngressCIDR = `
resource "aws_security_group" "test" {
  ingress {
om_port
_port = 0
otocol= "-1"
v6_cidr_blocks = ["::/244"]
  }
}
`

const testAccVPCSecurityGroupConfig_invalidIPv6EgressCIDR = `
resource "aws_security_group" "test" {
  egress {
om_port
_port = 0
otocol= "-1"
v6_cidr_blocks = ["::/244"]
  }
}
`


func testAccVPCSecurityGroupConfig_combinedCIDRAndGroups(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test2" {
  name%[1]s-2"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test3" {
  name%[1]s-3"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

funcme%[1]s-4"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test1" {
  name%[1]s-1"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  ingress {
om_port= 8
_port
otocol = p"
dr_blocks = ["10.0.0.0/16", "10.1.0.0/16", "10.7.0.0/16"]

curity_groups = [
ity_group.test2.id,
ity_group.test3.id,
ity_group.test4.id,

  }
}
func


func testAccVPCSecurityGroupConfig_ingressWithCIDRAndSGs(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test2" {
  name%[1]s-2"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test1" {
  name%[1]s-1"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
funcgress {
otocol  = "tcp"
om_port = "22"
_port= "

dr_blocks = [
0.1/32",

  }

  ingress {
otocol
om_port
_port= 8000
dr_blocks.0/8"]
curity_groups = [aws_security_group.test2.id]
  }

  egress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/8"]
  }
}
`, rName)
}

// fails to apply in one pass with the error "diffs didn't match during apply"
// GH-2027

func testAccVPCSecurityGroupConfig_failWithDiffMismatch(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test3" {
  vpc_id = aws_vpc.main.id
  name%[1]s-3"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test2" {
  vpc_id = aws_vpc.main.id
  name%[1]s-2"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test1" {
  vpc_id = aws_vpc.main.id
  name%[1]s-1"

  ingress {
om_port
_port= 22
otocol
curity_groups = [aws_security_group.test2.id]
  }

  ingress {
om_port
_port= 22
otocol
curity_groups = [aws_security_group.test3.id]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_allowAll(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group_rule" "allow_all-1" {
  typeess"
  from_port
  to_port
  protocol"tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.test.id
}

resource "aws_security_group_rule" "allow_all-2" {
  types"
  from_port = 65534
funcotocol  = "tcp"

  self
  security_group_id = aws_security_group.test.id
}
`, rName)
}


func testAccVPCSecurityGroupConfig_sourceSecurityGroup(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name%[1]s-1"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test2" {
  name%[1]s-2"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test3" {
  name%[1]s-3"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group_rule" "allow_test2" {
  types"
  from_port = 0
  to_port
  protocol  = "tcp"

  source_security_group_id = aws_security_group.test.id
  security_group_idecurity_group.test2.id
}

resource "aws_security_group_rule" "allow_test3" {
  types"
  from_port = 0
  to_port
  protocol  = "tcp"

funccurity_group_idecurity_group.test3.id
}
`, rName)
}


func testAccVPCSecurityGroupConfig_ipRangeAndSecurityGroupWithSameRules(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name%[1]s-1"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test2" {
  name%[1]s-2"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group_rule" "allow_security_group" {
  types"
  from_port = 0
  to_port
  protocol  = "tcp"

  source_security_group_id = aws_security_group.test2.id
  security_group_idecurity_group.test.id
}

resource "aws_security_group_rule" "allow_cidr_block" {
  types"
  from_port = 0
  to_port
  protocol  = "tcp"

  cidr_blocks.0.0/32"]
  security_group_id = aws_security_group.test.id
}

resource "aws_security_group_rule" "allow_ipv6_cidr_block" {
  types"
  from_port = 0
  to_port
func
  ipv6_cidr_blocks  = ["::/0"]
  security_group_id = aws_security_group.test.id
}
`, rName)
}


func testAccVPCSecurityGroupConfig_ipRangesWithSameRules(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group_rule" "allow_cidr_block" {
  types"
  from_port = 0
  to_port
  protocol  = "tcp"

  cidr_blocks.0.0/32"]
  security_group_id = aws_security_group.test.id
}

resource "aws_security_group_rule" "allow_ipv6_cidr_block" {
  types"
  from_port = 0
  to_port
  protocol  = "tcp"

  ipv6_cidr_blocks  = ["::/0"]
  security_group_id = aws_security_group.test.id
}
`, rName)
}


func testAccVPCSecurityGroupConfig_ipv4andIPv6Egress(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block.0.0/16"
func
  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  egress {
om_port= 0
_port
otocol = "
dr_blocks = ["0.0.0.0/0"]
  }

  egress {
om_port
_port = 0
otocol= "-1"
v6_cidr_blocks = ["::/0"]
  }

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_prefixListEgress(rName string) string {
	return fmt.Sprintf(`
data "aws_region" "current" {}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
func
resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_endpoint" "test" {
  vpc_id = aws_vpc.test.id
  service_name"com.amazonaws.${data.aws_region.current.name}.s3"
  route_table_ids = [aws_route_table.test.id]

  tags = {
me = %[1]q
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [

llowAll",
 "Allow",
l": "*",
 "*",
": "*"

  ]
}
POLICY
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  egress {
otocol
om_port
_port= 0
efix_list_ids = [aws_vpc_endpoint.test.prefix_list_id]
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_prefixListIngress(rName string) string {
	return fmt.Sprintf(`
data "aws_region" "current" {}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

func %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_endpoint" "test" {
  vpc_id = aws_vpc.test.id
  service_name"com.amazonaws.${data.aws_region.current.name}.s3"
  route_table_ids = [aws_route_table.test.id]

  tags = {
me = %[1]q
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [

llowAll",
 "Allow",
l": "*",
 "*",
": "*"

  ]
}
POLICY
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  ingress {
otocol
om_port
_port= 0
efix_list_ids = [aws_vpc_endpoint.test.prefix_list_id]
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_ruleGathering(rName string) string {
	return fmt.Sprintf(`
data "aws_region" "current" {}

resource "aws_vpc" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_endpoint" "test" {
  vpc_id = aws_vpc.test.id
  service_name"com.amazonaws.${data.aws_region.current.name}.s3"
  route_table_ids = [aws_route_table.test.id]

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [

llowAll",
 "Allow",
l": "*",
 "*",
": "*"

  ]
}
POLICY

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "source1" {
  name%[1]s-source1"
  vpc_id = aws_vpc.test.id
funcgs = {
me = %[1]q
  }
}

resource "aws_security_group" "source2" {
  name%[1]s-source2"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.0.0/24", "10.0.1.0/24"]
lf
  }

  ingress {
otocol = p"
om_port= 8
_port
dr_blocks = ["10.0.2.0/24", "10.0.3.0/24"]
scription = "ingress from 10.0.0.0/16"
  }
funcgress {
otocol = p"
om_port= 8
_port
dr_blocks = ["192.168.0.0/16"]
scription = "ingress from 192.168.0.0/16"
  }

  ingress {
otocol= "tcp"
om_port
_port = 80
v6_cidr_blocks = ["::/0"]
scriptions from all ipv6"
  }

  ingress {
otocol
om_port
_port= 80
curity_groups = [aws_security_group.source1.id, aws_security_group.source2.id]
scription from other security groups"
  }

  egress {
om_port= 0
_port
otocol = "
dr_blocks = ["0.0.0.0/0"]
scription = "egress for all ipv4"
  }

  egress {
om_port
_port = 0
otocol= "-1"
v6_cidr_blocks = ["::/0"]
scription for all ipv6"
  }

  egress {
om_port
_port= 0
otocol
efix_list_ids = [aws_vpc_endpoint.test.prefix_list_id]
scriptionfor vpc endpoints"
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_rulesDropOnErrorInit(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test_ref0" {
  name%[1]s-ref0"
func
  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test_ref1" {
  name%[1]s-ref1"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  ingress {
otocol  = "tcp"
om_port = "80"
_port= "
curity_groups = [
ity_group.test_ref0.id,
ity_group.test_ref1.id,

  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_rulesDropOnErrorAddBadRule(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test_ref0" {
  name%[1]s-ref0"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test_ref1" {
  name%[1]s-ref1"
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
funcurce "aws_security_group" "test" {
  name[1]q
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }

  ingress {
otocol  = "tcp"
om_port = "80"
_port= "
curity_groups = [
ity_group.test_ref0.id,
ity_group.test_ref1.id,
rmed", # non-existent rule to trigger API error

  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_egressModeBlocks(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name = %[1]q

  tags = {
me = %[1]q
  }

  vpc_id = aws_vpc.test.id

  egress {
dr_blocks = [aws_vpc.test.cidr_block]
om_port= 0
otocol = p"
_port
  }

  egress {
dr_blocks = [aws_vpc.test.cidr_block]
om_port= 0
otocol = p"
_port
  }
}
`, rName)
}


func testAccVPCSecurityGroupConfig_egressModeNoBlocks(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name = %[1]q

  tags = {
me = %[1]q
  }

  vpc_id = aws_vpc.test.id
}
`, rName)
}


func testAccVPCSecurityGroupConfig_egressModeZeroed(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name = %[1]q

  tags = {
me = %[1]q
  }

  egress = []

  vpc_id = aws_vpc.test.id
}
`, rName)
}


func testAccVPCSecurityGroupConfig_ingressModeBlocks(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name = %[1]q

  tags = {
me = %[1]q
  }

  vpc_id = aws_vpc.test.id

  ingress {
dr_blocks = [aws_vpc.test.cidr_block]
om_port= 0
otocol = p"
_port
  }

  ingress {
dr_blocks = [aws_vpc.test.cidr_block]
om_port= 0
otocol = p"
_port
func
`, rName)
}


func testAccVPCSecurityGroupConfig_ingressModeNoBlocks(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name = %[1]q

  tags = {
me = %[1]q
  }

  vpc_id = aws_vpc.test.id
}
`, rName)
}


func testAccVPCSecurityGroupConfig_ingressModeZeroed(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "test" {
  name = %[1]q

  tags = {
me = %[1]q
  }

  ingress = []

  vpc_id = aws_vpc.test.id
}
`, rName)
func
// testAccVPCSecurityGroupConfig_emrLinkedRulesDestroy is very involved but captures
// a problem seen in EMR and other contexts.

func testAccVPCSecurityGroupConfig_emrLinkedRulesDestroy(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptInDefaultExclude(),
fmt.Sprintf(`
# VPC
resource "aws_vpc" "main" {
  cidr_block = "10.1.0.0/16"
  tags = {
me = %[1]q
  }
}

# subnets
resource "aws_subnet" "private" {
  vpc_idws_vpc.main.id
  cidr_block.0.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
me = %[1]q
  }
}

resource "aws_security_group" "allow_ssh" {
  names-ssh"
  description = "ssh"
  vpc_id.main.id

  tags = {
me = "%[1]s-ssh"
  }
}

# internet gateway
resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id

  tags = {
me = %[1]q
  }
}

# elastic ip for NAT gateway
resource "aws_eip" "nat" {
  domain = "vpc"
  tags = {
me = %[1]q
func

# NAT gateway
resource "aws_nat_gateway" "nat" {
  allocation_id = aws_eip.nat.id
  subnet_idet.private.id

  tags = {
me = %[1]q
  }

  # To ensure proper ordering, it is recommended to add an explicit dependency
  # on the Internet Gateway for the VPC.
  depends_on = [aws_internet_gateway.gw]
}

# route tables
# add internet gateway
resource "aws_route_table" "test" {
  vpc_id = aws_vpc.main.id

  route {
dr_block = "0.0.0.0/0"
teway_id = aws_internet_gateway.gw.id
  }

  tags = {
me = %[1]q
  }
}

# route table for nat
resource "aws_route_table" "nat" {
  vpc_id = aws_vpc.main.id

  route {
dr_block/0"
func

  tags = {
me = %[1]q
  }
}

# associate nat route table with subnet
resource "aws_route_table_association" "nat" {
  subnet_idnet.private.id
  route_table_id = aws_route_table.nat.id
}

resource "aws_security_group" "allow_access" {
  name"%[1]s-allow-access"
  descriptionAllow inbound traffic"
  vpc_id  = aws_vpc.main.id
  revoke_rules_on_delete = true

  ingress {
om_port= 0
_port
otocol = "
func

  egress {
om_port= 0
_port
otocol = "
dr_blocks = ["0.0.0.0/0"]
  }

  lifecycle {
nore_changes = [



  }

  tags = {
me = "%[1]s-allow-access"
  }
}

resource "aws_security_group" "service_access" {
  name"%[1]s-service-access"
  descriptionAllow inbound traffic"
  vpc_id  = aws_vpc.main.id
func
  ingress {
om_port= 0
_port
otocol = "
dr_blocks = [aws_vpc.main.cidr_block]
  }

  ingress {
om_port
_port= 8443
otocol
dr_blocks.main.cidr_block]
curity_groups = [aws_security_group.allow_access.id]
  }

  ingress {
om_port
_port= 9443
otocol
dr_blocks.main.cidr_block]
curity_groups = [aws_security_group.allow_access.id]
  }

  egress {
om_port= 0
_port
otocol = "
dr_blocks = ["0.0.0.0/0"]
  }

  lifecycle {
nore_changes = [



  }
funcgs = {
me = "%[1]s-service-access"
  }
}

# IAM role for EMR Service
resource "aws_iam_role" "iam_emr_service_role" {
  name = "%[1]s-service-role"

  assume_role_policy = <<EOF
{
  "Version": "2008-10-17",
  "Statement": [

,
 "Allow",
l": {
e": "elasticmapreduce.amazonaws.com"

 "sts:AssumeRole"

  ]
}
func

resource "aws_iam_role_policy" "iam_emr_service_policy" {
  name = "%[1]s-service-policy"
  role = aws_iam_role.iam_emr_service_role.id

  policy = <<EOF
{
ersion": "2012-10-17",
tatement": [{
": "Allow",
ce": "*",
": [
2:AuthorizeSecurityGroupEgress",
2:AuthorizeSecurityGroupIngress",
2:CancelSpotInstanceRequests",
2:CreateNetworkInterface",
2:CreateSecurityGroup",
2:CreateTags",
2:DeleteNetworkInterface",
2:DeleteSecurityGroup",
2:DeleteTags",
2:DescribeAvailabilityZones",
2:DescribeAccountAttributes",
2:DescribeDhcpOptions",
2:DescribeInstanceStatus",
2:DescribeInstances",
funcscribeNetworkAcls",
2:DescribeNetworkInterfaces",
2:DescribePrefixLists",
2:DescribeRouteTables",
2:DescribeSecurityGroups",
2:DescribeSpotInstanceRequests",
2:DescribeSpotPriceHistory",
2:DescribeSubnets",
2:DescribeVpcAttribute",
2:DescribeVpcEndpoints",
2:DescribeVpcEndpointServices",
2:DescribeVpcs",
2:DetachNetworkInterface",
2:ModifyImageAttribute",
2:ModifyInstanceAttribute",
2:RequestSpotInstances",
2:RevokeSecurityGroupEgress",
2:RunInstances",
2:TerminateInstances",
2:DeleteVolume",
2:DescribeVolumeStatus",
2:DescribeVolumes",
2:DetachVolume",
m:GetRole",
m:GetRolePolicy",
m:ListInstanceProfiles",
m:ListRolePolicies",
m:PassRole",
:CreateBucket",
:Get*",
:List*",
b:BatchPutAttributes",
b:Select",
s:CreateQueue",
s:Delete*",
s:GetQueue*",
s:PurgeQueue",
s:ReceiveMessage"


}
EOF
}

# IAM Role for EC2 Instance Profile
resource "aws_iam_role" "iam_emr_profile_role" {
  name = "%[1]s-profile-role"

  assume_role_policy = <<EOF
{
  "Version": "2008-10-17",
  "Statement": [

,
 "Allow",
l": {
e": "ec2.amazonaws.com"

 "sts:AssumeRole"

  ]
}
EOF
}

resource "aws_iam_instance_profile" "emr_profile" {
  name = "%[1]s-profile"
  role = aws_iam_role.iam_emr_profile_role.name
}

resource "aws_iam_role_policy" "iam_emr_profile_policy" {
  name = "%[1]s-profile-policy"
  role = aws_iam_role.iam_emr_profile_role.id

  policy = <<EOF
{
ersion": "2012-10-17",
tatement": [{
": "Allow",
ce": "*",
": [
oudwatch:*",
namodb:*",
2:Describe*",
asticmapreduce:Describe*",
asticmapreduce:ListBootstrapActions",
asticmapreduce:ListClusters",
asticmapreduce:ListInstanceGroups",
asticmapreduce:ListInstances",
asticmapreduce:ListSteps",
nesis:CreateStream",
nesis:DeleteStream",
nesis:DescribeStream",
nesis:GetRecords",
nesis:GetShardIterator",
nesis:MergeShards",
nesis:PutRecord",
nesis:SplitShard",
s:Describe*",
:*",
b:*",
s:*",
s:*"


}
EOF
}

resource "aws_emr_cluster" "cluster" {
  name = %[1]q
  release_label = "emr-6.6.0"
  applications  = ["Spark"]

  additional_info = <<EOF
{
  "instanceAwsClientConfiguration": {
roxyPort": 8099,
roxyHost": "myproxy.example.com"
  }
}
EOF

  termination_protectionalse
  keep_job_flow_alive_when_no_steps = true

  ec2_attributes {
bnet_id = aws_subnet.private.id
stance_profile= aiam_instance_profile.emr_profile.arn
r_managed_master_security_group = aws_security_group.allow_access.id
r_managed_slave_security_group  = aws_security_group.allow_access.id
ditional_master_security_groups = aws_security_group.allow_ssh.id
ditional_slave_security_groups  = aws_security_group.allow_ssh.id
rvice_access_security_grouprity_group.service_access.id
  }

  master_instance_group {
stance_type = "c4.large"
  }

  core_instance_group {
stance_type  = "c4.large"
stance_count = 1

s_config {
40"
gp2"
er_instance = 1

  }

  ebs_root_volume_size = 100

  tags = {
le = "rolename"
v  = "env"
  }

  bootstrap_action {
th = "s3://elasticmapreduce/bootstrap-actions/run-if"
me = "runif"
gs = ["instance.isMaster=true", "echo running on master node"]
  }

  configurations_json = <<EOF
  [

cation": "hadoop-env",
ations": [

 "Classification": "export",
 "Properties": {
VA_HOME": "/usr/lib/jvm/java-1.8.0"
 }


es": {}


cation": "spark-env",
ations": [

 "Classification": "export",
 "Properties": {
VA_HOME": "/usr/lib/jvm/java-1.8.0"
 }


es": {}

  ]
EOF

  service_role = aws_iam_role.iam_emr_service_role.arn

  depends_on = [
s_route_table_association.nat,
s_iam_role_policy.iam_emr_service_policy,
s_iam_role_policy.iam_emr_profile_policy
  ]
}
`, rName))
}
