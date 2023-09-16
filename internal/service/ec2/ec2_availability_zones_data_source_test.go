// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


funcarallel()

	azs := []*ec2.AvailabilityZone{
{
	ZoneName: aws.String("name_YYY"),
	ZoneId:.String("id_YYY"),
},
{
	ZoneName: aws.String("name_AAA"),
	ZoneId:.String("id_AAA"),
},
{
	ZoneName: aws.String("name_ZZZ"),
	ZoneId:.String("id_ZZZ"),
},
{
	ZoneName: aws.String("name_BBB"),
	ZoneId:.String("id_BBB"),
},
	}
	sort.Slice(azs, 
func(i, j int) bool {
func

	cases := []struct {
Indext
ZoneName string
ZoneIding
	}{
{
	Index:
	ZoneName: "name_AAA",
	ZoneId:_AAA",
},
{
	Index:
	ZoneName: "name_BBB",
	ZoneId:_BBB",
},
{
	Index:
	ZoneName: "name_YYY",
	ZoneId:_YYY",
},
{
	Index:
	ZoneName: "name_ZZZ",
	ZoneId:_ZZZ",
},
	}
	for _, tc := range cases {
az := azs[tc.Index]
if aws.StringValue(az.ZoneName) != tc.ZoneName {
	t.Fatalf("AvailabilityZones index %d got zone name %s, expected %s", tc.Index, aws.StringValue(az.ZoneName), tc.ZoneName)
}
if aws.StringValue(az.ZoneId) != tc.ZoneId {
	t.Fatalf("AvailabilityZones index %d got zone ID %s, expected %s", tc.Index, aws.StringValue(az.ZoneId), tc.ZoneId)
}
	}
}


func TestAccEC2AvailabilityZonesDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
func
Config: testAccAvailabilityZonesDataSourceConfig_basic,
Check: resource.ComposeTestCheck
func(
	testAccCheckAvailabilityZonesMeta("data.aws_availability_zones.availability_zones"),
),
	},
func
}


func TestAccEC2AvailabilityZonesDataSource_allAvailabilityZones(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_availability_zones.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccAvailabilityZonesDataSourceConfig_all(),
func(
	testAccCheckAvailabilityZonesMeta(dataSourceName),
),
	},
},
	})
}
func
func TestAccEC2AvailabilityZonesDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_availability_zones.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funcs: []resource.TestStep{
	{
Config: testAccAvailabilityZonesDataSourceConfig_filter(),
Check: resource.ComposeTestCheck
func(
	testAccCheckAvailabilityZonesMeta(dataSourceName),
func
},
	})
}


func TestAccEC2AvailabilityZonesDataSource_excludeNames(t *testing.T) {
funcDataSourceName := "data.aws_availability_zones.all"
	excludeDataSourceName := "data.aws_availability_zones.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccAvailabilityZonesDataSourceConfig_excludeNames(),
Check: resource.ComposeTestCheck
func(
	testAccCheckAvailabilityZonesExcluded(allDataSourceName, excludeDataSourceName),
),
	},
},
func


func TestAccEC2AvailabilityZonesDataSource_excludeZoneIDs(t *testing.T) {
	ctx := acctest.Context(t)
	allDataSourceName := "data.aws_availability_zones.all"
	excludeDataSourceName := "data.aws_availability_zones.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccAvailabilityZonesDataSourceConfig_excludeZoneIDs(),
Check: resource.ComposeTestCheck
functAccCheckAvailabilityZonesExcluded(allDataSourceName, excludeDataSourceName),
),
	},
},
	})
}

func TestAccEC2AvailabilityZonesDataSource_stateFilter(t *testing.T) {
	ctx := acctest.Context(t)
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
func
Config: testAccAvailabilityZonesDataSourceConfig_state,
Check: resource.ComposeTestCheck
func(
	testAccCheckAvailabilityZoneState("data.aws_availability_zones.state_filter"),
),
	},
},
	})
func

func testAccCheckAvailabilityZonesMeta(n string) resource.TestCheck
func {
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Can't find AZ resource: %s", n)
}

if rs.Primary.ID == "" {
func

actual, err := testAccCheckAvailabilityZonesBuildAvailable(rs.Primary.Attributes)
if err != nil {
	return err
}

expected := actual
sort.Strings(expected)
funcurn fmt.Errorf("AZs not sorted - expected %v, got %v", expected, actual)
funcrn nil
	}
func

func testAccCheckAvailabilityZonesExcluded(allDataSourceName, excludeDataSourceName string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
allResourceState, ok := s.RootModule().Resources[allDataSourceName]
if !ok {
	return fmt.Errorf("Resource does not exist: %s", allDataSourceName)
}

excludeResourceState, ok := s.RootModule().Resources[excludeDataSourceName]
if !ok {
	return fmt.Errorf("Resource does not exist: %s", excludeDataSourceName)
}

for _, attribute := range []string{"names.#", "zone_ids.#"} {
	allValue, ok := allResourceState.Primary.Attributes[attribute]

	if !ok {
return fmt.Errorf("cannot find %s in %s resource state attributes: %+v", attribute, allDataSourceName, allResourceState.Primary.Attributes)
	}

	excludeValue, ok := excludeResourceState.Primary.Attributes[attribute]

funcrn fmt.Errorf("cannot find %s in %s resource state attributes: %+v", attribute, excludeDataSourceName, excludeResourceState.Primary.Attributes)
func
	if allValue == excludeValue {
func
}

return nil
	}
}


func testAccCheckAvailabilityZoneState(n string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Can't find AZ resource: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("AZ resource ID not set.")
}

if _, ok := rs.Primary.Attributes["state"]; !ok {
	return fmt.Errorf("AZs state filter is missing, should be set.")
}

_, err := testAccCheckAvailabilityZonesBuildAvailable(rs.Primary.Attributes)
return err
	}
}


func testAccCheckAvailabilityZonesBuildAvailable(attrs map[string]string) ([]string, error) {
	groupNames, groupNamesOk := attrs["group_names.#"]

funcrn nil, fmt.Errorf("Availability Zone Group names list is missing.")
func
	groupNamesQty, err := strconv.Atoi(groupNames)
funcerr != nil {
return nil, err
	}

	if groupNamesQty < 1 {
return nil, fmt.Errorf("No Availability Zone Groups found in region, this is probably a bug.")
	}

	v, ok := attrs["names.#"]
	if !ok {
return nil, fmt.Errorf("Available AZ name list is missing.")
	}
	qty, err := strconv.Atoi(v)
	if err != nil {
return nil, err
	}
	if qty < 1 {
return nil, fmt.Errorf("No AZs found in region, this is probably a bug.")
	}
	_, ok = attrs["zone_ids.#"]
funcrn nil, fmt.Errorf("Available AZ ID list is missing.")
	}
	zones := make([]string, qty)
	for n := range zones {
zone, ok := attrs["names."+strconv.Itoa(n)]
if !ok {
	return nil, fmt.Errorf("AZ list corrupt, this is definitely a bug.")
}
zones[n] = zone
	}
	return zones, nil
}

const testAccAvailabilityZonesDataSourceConfig_basic = `
data "aws_availability_zones" "availability_zones" {}
`


func testAccAvailabilityZonesDataSourceConfig_all() string {
	return `
data "aws_availability_zones" "test" {
all_availability_zones = true
}
`
}


func testAccAvailabilityZonesDataSourceConfig_filter() string {
	return `
data "aws_availability_zones" "test" {
filter {
me= "te"
lues = ["available"]
}
}
`
}


func testAccAvailabilityZonesDataSourceConfig_excludeNames() string {
	return `
data "aws_availability_zones" "all" {}

data "aws_availability_zones" "test" {
exclude_names = [data.aws_availability_zones.all.names[0]]
}
`
}
func
func testAccAvailabilityZonesDataSourceConfig_excludeZoneIDs() string {
	return `
data "aws_availability_zones" "all" {}

data "aws_availability_zones" "test" {
exclude_zone_ids = [data.aws_availability_zones.all.zone_ids[0]]
}
`
func
const testAccAvailabilityZonesDataSourceConfig_state = `
data "aws_availability_zones" "state_filter" {
state = "available"
}
`
funcfunc