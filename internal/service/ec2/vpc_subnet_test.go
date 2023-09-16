// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
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
	var v ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttrSet(resourceName, "availability_zone"),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone_id"),
	resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.1.1.0/24"),
	resource.TestCheckResourceAttr(resourceName, "customer_owned_ipv4_pool", ""),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_a_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "ipv6_native", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_customer_owned_ip_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_hostname_type_on_launch", "ip-name"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
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


func TestAccVPCSubnet_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
func
Config: testAccVPCSubnetConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCSubnetConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
func
Config: testAccVPCSubnetConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2")),
	},
},
	})
func

func TestAccVPCSubnet_DefaultTags_providerOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
	testAccVPCSubnetConfig_basic(rName),
funck: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
),
	},
	{
ResourceName:ame,
ImportState:
func
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags2("providerkey1", "providervalue1", "providerkey2", "providervalue2"),
	testAccVPCSubnetConfig_basic(rName),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey2", "providervalue2"),
),
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey1", "value1"),
func
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "value1"),
),
	},
},
	})
}


func := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCSubnetConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
func
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("key1", "value1"),
	testAccVPCSubnetConfig_basic(rName),
),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
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

func TestAccVPCSubnet_DefaultTags_updateToResourceOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("key1", "value1"),
	testAccVPCSubnetConfig_basic(rName),
funck: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
),
	},
funcig: testAccVPCSubnetConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
),
	},
	{
funcrtState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCSubnet_DefaultTagsProviderAndResource_nonOverlappingTag(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
	testAccVPCSubnetConfig_tags1(rName, "resourcekey1", "resourcevalue1"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
funcource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
	testAccVPCSubnetConfig_tags2(rName, "resourcekey1", "resourcevalue1", "resourcekey2", "resourcevalue2"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
funcource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags.resourcekey2", "resourcevalue2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey2", "resourcevalue2"),
),
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("providerkey2", "providervalue2"),
	testAccVPCSubnetConfig_tags1(rName, "resourcekey3", "resourcevalue3"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.resourcekey3", "resourcevalue3"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey2", "providervalue2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey3", "resourcevalue3"),
func
},
	})
}


func TestAccVPCSubnet_DefaultTagsProviderAndResource_overlappingTag(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("overlapkey1", "providervalue1"),
	testAccVPCSubnetConfig_tags1(rName, "overlapkey1", "resourcevalue1"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags2("overlapkey1", "providervalue1", "overlapkey2", "providervalue2"),
	testAccVPCSubnetConfig_tags2(rName, "overlapkey1", "resourcevalue1", "overlapkey2", "resourcevalue2"),
funck: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags.overlapkey2", "resourcevalue2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey2", "resourcevalue2"),
),
	},
funcig: acctest.ConfigCompose(
	acctest.ConfigDefaultTags_Tags1("overlapkey1", "providervalue1"),
	testAccVPCSubnetConfig_tags1(rName, "overlapkey1", "resourcevalue2"),
),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue2"),
),
	},
},
	})
}


func := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
functAccCheckSubnetUpdateTags(ctx, &subnet, nil, map[string]string{"defaultkey1": "defaultvalue1"}),
),
ExpectNonEmptyPlan: true,
	},
	{
Config: acctest.ConfigCompose(
	acctest.ConfigDefaultAndIgnoreTagsKeyPrefixes1("defaultkey1", "defaultvalue1", "defaultkey"),
	testAccVPCSubnetConfig_tags1(rName, "key1", "value1"),
),
PlanOnly: true,
	},
	{
Config: acctest.ConfigCompose(
functAccVPCSubnetConfig_tags1(rName, "key1", "value1"),
),
PlanOnly: true,
	},
},
	})
}

functtributes are correctly determined when the provider-level default_tags block
// is left unused and resource tags are only known at apply time, thereby
// eliminating "Inconsistent final plan" errors
// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/18366

func TestAccVPCSubnet_updateTagsKnownAtApply(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_tagsComputedFromDataSource1("key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
),
	},
	{
Config: testAccVPCSubnetConfig_tagsComputedFromDataSource2("key1", "value1", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}


func TestAccVPCSubnet_ignoreTags(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	testAccCheckSubnetUpdateTags(ctx, &subnet, nil, map[string]string{"ignorekey1": "ignorevalue1"}),
),
ExpectNonEmptyPlan: true,
func
Config:test.ConfigCompose(acctest.ConfigIgnoreTagsKeyPrefixes1("ignorekey"), testAccVPCSubnetConfig_basic(rName)),
PlanOnly: true,
	},
	{
Config:test.ConfigCompose(acctest.ConfigIgnoreTagsKeys("ignorekey1"), testAccVPCSubnetConfig_basic(rName)),
PlanOnly: true,
	},
},
	})
}


func TestAccVPCSubnet_ipv6(t *testing.T) {
	ctx := acctest.Context(t)
	var before, after ec2.Subnet
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
func
Config: testAccVPCSubnetConfig_ipv6(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &before),
	testAccCheckSubnetIPv6BeforeUpdate(&before),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCSubnetConfig_ipv6UpdateAssignv6OnCreation(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &after),
	testAccCheckSubnetIPv6AfterUpdate(&after),
),
	},
	{
Config: testAccVPCSubnetConfig_ipv6UpdateV6CIDR(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &after),
	testAccCheckSubnetNotRecreated(t, &before, &after),
),
func
	})
}


func TestAccVPCSubnet_enableIPv6(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
func
Config: testAccVPCSubnetConfig_prev6(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "false"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
Config: testAccVPCSubnetConfig_ipv6(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttrSet(resourceName, "ipv6_cidr_block"),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "true"),
),
func
Config: testAccVPCSubnetConfig_prev6(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "false"),
),
	},
},
func


func TestAccVPCSubnet_availabilityZoneID(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone"),
	resource.TestCheckResourceAttrPair(resourceName, "availability_zone_id", "data.aws_availability_zones.available", "zone_ids.0"),
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

func TestAccVPCSubnet_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceSubnet(), resourceName),
funcctNonEmptyPlan: true,
	},
},
	})
}


func TestAccVPCSubnet_customerOwnedIPv4Pool(t *testing.T) {
func subnet ec2.Subnet
	coipDataSourceName := "data.aws_ec2_coip_pool.test"
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_customerOwnedv4Pool(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttrPair(resourceName, "customer_owned_ipv4_pool", coipDataSourceName, "pool_id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
func
}


func TestAccVPCSubnet_mapCustomerOwnedIPOnLaunch(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCSubnetConfig_mapCustomerOwnedOnLaunch(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "map_customer_owned_ip_on_launch", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
	})
}


func TestAccVPCSubnet_mapPublicIPOnLaunch(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCSubnetConfig_mapPublicOnLaunch(rName, true),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCSubnetConfig_mapPublicOnLaunch(rName, false),
Check: resource.ComposeTestCheck
functAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "false"),
),
	},
	{
Config: testAccVPCSubnetConfig_mapPublicOnLaunch(rName, true),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "true"),
),
	},
},
	})
}


func := acctest.Context(t)
	var v ec2.Subnet
	outpostDataSourceName := "data.aws_outposts_outpost.test"
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_outpost(rName),
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrPair(resourceName, "outpost_arn", outpostDataSourceName, "arn"),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccVPCSubnet_enableDNS64(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_enableDNS64(rName, true),
Check: resource.ComposeTestCheck
functAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
Config: testAccVPCSubnetConfig_enableDNS64(rName, false),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "false"),
),
	},
	{
Config: testAccVPCSubnetConfig_enableDNS64(rName, true),
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "true"),
),
	},
},
	})
}

func TestAccVPCSubnet_ipv4ToIPv6(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_ipv4ToIPv6Before(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
),
	},
funcig: testAccVPCSubnetConfig_ipv4ToIPv6After(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "true"),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "true"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "true"),
	resource.TestCheckResourceAttrSet(resourceName, "ipv6_cidr_block"),
func
},
	})
}


func TestAccVPCSubnet_enableLNIAtDeviceIndex(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_enableLniAtDeviceIndex(rName, 1),
Check: resource.ComposeTestCheck
functAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "enable_lni_at_device_index", "1"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
Config: testAccVPCSubnetConfig_enableLniAtDeviceIndex(rName, 1),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "enable_lni_at_device_index", "1"),
),
	},
	{
Config: testAccVPCSubnetConfig_enableLniAtDeviceIndex(rName, 1),
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "enable_lni_at_device_index", "1"),
),
	},
},
	})
}
func
func TestAccVPCSubnet_privateDNSNameOptionsOnLaunch(t *testing.T) {
	ctx := acctest.Context(t)
	var subnet ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCSubnetConfig_privateDNSNameOptionsOnLaunch(rName, true, true, "resource-name"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "true"),
funcource.TestCheckResourceAttr(resourceName, "private_dns_hostname_type_on_launch", "resource-name"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCSubnetConfig_privateDNSNameOptionsOnLaunch(rName, false, true, "ip-name"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &subnet),
funcource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_a_record_on_launch", "true"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_hostname_type_on_launch", "ip-name"),
),
	},
	{
Config: testAccVPCSubnetConfig_privateDNSNameOptionsOnLaunch(rName, true, false, "resource-name"),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "true"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_a_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_hostname_type_on_launch", "resource-name"),
),
	},
},
	})
}
func
func TestAccVPCSubnet_ipv6Native(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
	resourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSubnetDestroy(ctx),
func
Config: testAccVPCSubnetConfig_ipv6Native(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "true"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_native", "true"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}

func testAccCheckSubnetIPv6BeforeUpdate(subnet *ec2.Subnet) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if subnet.Ipv6CidrBlockAssociationSet == nil {
	return fmt.Errorf("Expected IPV6 CIDR Block Association")
}

funcurn fmt.Errorf("bad AssignIpv6AddressOnCreation: %t", aws.BoolValue(subnet.AssignIpv6AddressOnCreation))
}

return nil
	}
}


func {
	return 
func(s *terraform.State) error {
if aws.BoolValue(subnet.AssignIpv6AddressOnCreation) {
	return fmt.Errorf("bad AssignIpv6AddressOnCreation: %t", aws.BoolValue(subnet.AssignIpv6AddressOnCreation))
}

return nil
	}
}


func testAccCheckSubnetNotRecreated(t *testing.T, before, after *ec2.Subnet) resource.TestCheck
func {
	return 
funcws.StringValue(before.SubnetId) != aws.StringValue(after.SubnetId) {
	t.Fatalf("Expected SubnetIDs not to change, but both got before: %s and after: %s",
aws.StringValue(before.SubnetId), aws.StringValue(after.SubnetId))
}
return nil
	}
}


func testAccCheckSubnetDestroy(ctx context.Context) resource.TestCheck
funcurn 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_subnet" {
continue
	}

	_, err := tfec2.FindSubnetByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
func

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Subnet %s still exists", rs.Primary.ID)
}
funcrn nil
	}
}


func testAccCheckSubnetExists(ctx context.Context, n string, v *ec2.Subnet) resource.TestCheck
func {
	return 
funcok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Subnet ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindSubnetByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

func
func
}
func
func testAccCheckSubnetUpdateTags(ctx context.Context, subnet *ec2.Subnet, oldTags, newTags map[string]string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

return tfec2.UpdateTags(ctx, conn, aws.StringValue(subnet.SubnetId), oldTags, newTags)
	}
}


func testAccVPCSubnetConfig_basic(rName string) string {
	return fmt.Sprintf(`
funcdr_block = "10.1.0.0/16"
funcgs = {
me = %[1]q
func

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_idtest.id
}
`, rName)
}


funcurn fmt.Sprintf(`
funcdr_block = "10.1.0.0/16"

func %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_idtest.id

  tags = {
2]q = %[3]q
func
func

func testAccVPCSubnetConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_idtest.id

  tags = {
2]q = %[3]q
4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}

const testAccSubnetComputedTagsBaseConfig = `
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"
  tagstags
func
func = aws_vpc.test.id
}
funcurce "aws_subnet" "test" {
  cidr_block = cidrsubnet(aws_vpc.test.cidr_block, 8, 0)
  vpc_idtest.id
  tagsws_vpc.test.tags
}
`


func testAccVPCSubnetConfig_tagsComputedFromDataSource1(tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(
testAccSubnetComputedTagsBaseConfig,
fmt.Sprintf(`
locals {
  tags = {
 = %q
  }
}
`, tagKey1, tagValue1))
}


func testAccVPCSubnetConfig_tagsComputedFromDataSource2(tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(
testAccSubnetComputedTagsBaseConfig,
fmt.Sprintf(`
funcgs = {
funcq
  }
funcagKey1, tagValue1, tagKey2, tagValue2))
}


func testAccVPCSubnetConfig_prev6(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
func
  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block = "10.10.1.0/24"
  vpc_idtest.id

  tags = {
me = %[1]q
  }
}
`, rName)
}


funcurn fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_id  = aws_vpc.test.id
  ipv6_cidr_block  = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)
  assign_ipv6_address_on_creation = true

  tags = {
me = %[1]q
  }
}
`, rName)
}
func
func testAccVPCSubnetConfig_ipv6UpdateAssignv6OnCreation(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_id  = aws_vpc.test.id
  ipv6_cidr_block  = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)
  assign_ipv6_address_on_creation = false

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSubnetConfig_ipv6UpdateV6CIDR(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block.1.0/24"
  vpc_id  = aws_vpc.test.id
funcsign_ipv6_address_on_creation = false

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSubnetConfig_availabilityZoneID(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_vpc" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block  = "10.1.1.0/24"
  vpc_id= aws_vpc.test.id
  availability_zone_id = data.aws_availability_zones.available.zone_ids[0]

  tags = {
me = %[1]q
  }
funcName))
}


func testAccVPCSubnetConfig_customerOwnedv4Pool(rName string) string {
	return fmt.Sprintf(`
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "test" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}

data "aws_ec2_local_gateway_route_tables" "test" {
  filter {
me= "post-arn"
lues = [data.aws_outposts_outpost.test.arn]
  }
}

data "aws_ec2_coip_pools" "test" {
  # Filtering by Local Gateway Route Table ID is documented but not working in EC2 API.
  # If there are multiple Outposts in the test account, this lookup can
  # be misaligned and cause downstream resource errors.
funcfilter {
  #e= "p-pool.local-gateway-route-table-id"
  #ues = [tolist(data.aws_ec2_local_gateway_route_tables.test.ids)[0]]
  # }
}

data "aws_ec2_coip_pool" "test" {
  pool_id = tolist(data.aws_ec2_coip_pools.test.pool_ids)[0]
}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone= data.aws_outposts_outpost.test.availability_zone
  cidr_blockbnet(aws_vpc.test.cidr_block, 8, 0)
  customer_owned_ipv4_poolaws_ec2_coip_pool.test.pool_id
  map_customer_owned_ip_on_launch = true
  outpost_arns_outposts_outpost.test.arn
  vpc_id  = aws_vpc.test.id
funcgs = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSubnetConfig_mapCustomerOwnedOnLaunch(rName string, mapCustomerOwnedIpOnLaunch bool) string {
	return fmt.Sprintf(`
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "test" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}

data "aws_ec2_local_gateway_route_tables" "test" {
  filter {
me= "post-arn"
lues = [data.aws_outposts_outpost.test.arn]
  }
}

data "aws_ec2_coip_pools" "test" {
  # Filtering by Local Gateway Route Table ID is documented but not working in EC2 API.
funcbe misaligned and cause downstream resource errors.
  #
  # filter {
  #e= "p-pool.local-gateway-route-table-id"
  #ues = [tolist(data.aws_ec2_local_gateway_route_tables.test.ids)[0]]
  # }
}

data "aws_ec2_coip_pool" "test" {
  pool_id = tolist(data.aws_ec2_coip_pools.test.pool_ids)[0]
}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone= data.aws_outposts_outpost.test.availability_zone
  cidr_blockbnet(aws_vpc.test.cidr_block, 8, 0)
  customer_owned_ipv4_poolaws_ec2_coip_pool.test.pool_id
  map_customer_owned_ip_on_launch = %[2]t
funcc_id  = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
`, rName, mapCustomerOwnedIpOnLaunch)
}


func testAccVPCSubnetConfig_mapPublicOnLaunch(rName string, mapPublicIpOnLaunch bool) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_blocket(aws_vpc.test.cidr_block, 8, 0)
  map_public_ip_on_launch = %[2]t
func
  tags = {
me = %[1]q
  }
}
`, rName, mapPublicIpOnLaunch)
}


func testAccVPCSubnetConfig_enableDNS64(rName string, enableDns64 bool) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_blockbnet(aws_vpc.test.cidr_block, 8, 0)
  enable_dns64
  vpc_id  = aws_vpc.test.id
  ipv6_cidr_block  = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)
  assign_ipv6_address_on_creation = true

  tags = {
me = %[1]q
  }
}
`, rName, enableDns64)
}


func testAccVPCSubnetConfig_enableLniAtDeviceIndex(rName string, deviceIndex int) string {
	return fmt.Sprintf(`


data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "test" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}

resource "aws_vpc" "test" {
  cidr_block = "10.10.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
funcdr_block  = cidrsubnet(aws_vpc.test.cidr_block, 8, 0)
  enable_lni_at_device_index = %[2]d
  outpost_arn = data.aws_outposts_outpost.test.arn
  vpc_id.test.id

  tags = {
me = %[1]q
  }
}
`, rName, deviceIndex)
}


func testAccVPCSubnetConfig_privateDNSNameOptionsOnLaunch(rName string, enableDnsAAAA, enableDnsA bool, hostnameType string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_blockbnet(aws_vpc.test.cidr_block, 8, 0)
  vpc_id  = aws_vpc.test.id
  ipv6_cidr_block  = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)
  assign_ipv6_address_on_creation = true

  enable_resource_name_dns_aaaa_record_on_launch = %[2]t
  enable_resource_name_dns_a_record_on_launch%[3]t
  private_dns_hostname_type_on_launch[4]q

  tags = {
me = %[1]q
  }
}
`, rName, enableDnsAAAA, enableDnsA, hostnameType)
}


func testAccVPCSubnetConfig_ipv6Native(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
funcv6_cidr_block  = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)
  assign_ipv6_address_on_creation = true
  ipv6_native

  enable_resource_name_dns_aaaa_record_on_launch = true

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSubnetConfig_outpost(rName string) string {
	return fmt.Sprintf(`
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "test" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}

resource "aws_vpc" "test" {
func
  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_outposts_outpost.test.availability_zone
  cidr_block.1.0/24"
  outpost_arnws_outposts_outpost.test.arn
  vpc_idws_vpc.test.id

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSubnetConfig_ipv4ToIPv6Before(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
  assign_generated_ipv6_cidr_block = false

func %[1]q
  }
}

resource "aws_subnet" "test" {
  assign_ipv6_address_on_creation = false
  cidr_blockbnet(aws_vpc.test.cidr_block, 8, 1)
  enable_dns64
  enable_resource_name_dns_aaaa_record_on_launch = false
  ipv6_cidr_block  = null
  vpc_id  = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCSubnetConfig_ipv4ToIPv6After(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block0.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  assign_ipv6_address_on_creation = true
funcable_dns64
  enable_resource_name_dns_aaaa_record_on_launch = true
  ipv6_cidr_block  = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)
  vpc_id  = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
`, rName)
}
funcfuncfuncfunc