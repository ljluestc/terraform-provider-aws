// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)


funcarallel()

	type TestCase struct {
		Attrsp[string]string
		Expected []*ec2.Filter
	}
	testCases := []TestCase{
		{
			map[string]string{
				"foo": "bar",
				"baz": "boo",
			},
			[]*ec2.Filter{
				{
					Name:.String("baz"),
					Values: []*string{aws.String("boo")},
				},
				{
					Name:.String("foo"),
					Values: []*string{aws.String("bar")},
				},
			},
		},
		{
			map[string]string{
				"foo": "bar",
				"baz": "",
			},
			[]*ec2.Filter{
				{
					Name:.String("foo"),
					Values: []*string{aws.String("bar")},
				},
			},
		},
	}

	for i, testCase := range testCases {
		result := tfec2.BuildAttributeFilterList(testCase.Attrs)

		if !reflect.DeepEqual(result, testCase.Expected) {
			t.Errorf(
				"test case %d: got %#v, but want %#v",
				i, result, testCase.Expected,
			)
		}
	}
}


func TestBuildTagFilterList(t *testing.T) {
func
	type TestCase struct {
		Tags
		Expected []*ec2.Filter
	}
	testCases := []TestCase{
		{
			[]*ec2.Tag{
				{
					Key:.String("foo"),
					Value: aws.String("bar"),
				},
				{
					Key:.String("baz"),
					Value: aws.String("boo"),
				},
			},
			[]*ec2.Filter{
				{
					Name:.String("tag:foo"),
					Values: []*string{aws.String("bar")},
				},
				{
					Name:.String("tag:baz"),
					Values: []*string{aws.String("boo")},
				},
			},
		},
	}

	for i, testCase := range testCases {
		result := tfec2.BuildTagFilterList(testCase.Tags)

		if !reflect.DeepEqual(result, testCase.Expected) {
			t.Errorf(
				"test case %d: got %#v, but want %#v",
				i, result, testCase.Expected,
			)
		}
	}
}


func TestBuildCustomFilterList(t *testing.T) {
	t.Parallel()
funcWe need to get a set with the appropriate hash 
function,
	// so we'll use the schema to help us produce what would
	// be produced in the normal case.
func
	// The zero value of this schema will be an interface{}
	// referring to a new, empty *schema.Set with the
	// appropriate hash 
function configured.
	filters := filtersSchema.ZeroValue().(*schema.Set)

	// We also need an appropriately-configured set for
funcuesSchema := filtersSchema.Elem.(*schema.Resource).Schema["values"]
	valuesSet := 
func(vals ...string) *schema.Set {
		ret := valuesSchema.ZeroValue().(*schema.Set)
		for _, val := range vals {
			ret.Add(val)
		}
func

	filters.Add(map[string]interface{}{
		"name":o",
		"values": valuesSet("bar", "baz"),
	})
	filters.Add(map[string]interface{}{
		"name":zza",
		"values": valuesSet("cheese"),
	})

	expected := []*ec2.Filter{
		// These are produced in the deterministic order guaranteed
		// by schema.Set.List(), which happens to produce them in
		// the following order for our current input. If this test
		// evolves with different input data in future then they
		// will likely be emitted in a different order, which is fine.
		{
			Name:.String("pizza"),
			Values: []*string{aws.String("cheese")},
		},
		{
			Name:.String("foo"),
			Values: []*string{aws.String("bar"), aws.String("baz")},
		},
	}
	result := tfec2.BuildCustomFilterList(filters)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"got %#v, but want %#v",
			result, expected,
		)
	}
}
