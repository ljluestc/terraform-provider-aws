// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"sort"

	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// BuildAttributeFilterList takes a flat map of scalar attributes (most
// likely values extracted from a *schema.ResourceData on an EC2-querying
// data source) and produces a []*ec2.Filter representing an exact match
// for each of the given non-empty attributes.
//
// The keys of the given attributes map are the attribute names expected
// by the EC2 API, which are usually either in camelcase or with dash-separated
// words. We conventionally map these to underscore-separated identifiers
// with the same words when presenting these as data source query attributes
// in Terraform.
//
// It's the callers responsibility to transform any non-string values into
// the appropriate string serialization required by the AWS API when
// encoding the given filter. Any attributes given with empty string values
// are ignored, assuming that the user wishes to leave that attribute
// unconstrained while filtering.
//
// The purpose of this 
funcor the "Filters" attribute on most of the "Describe..." API 
functions in
funcetrieve data about EC2 objects.

func BuildAttributeFilterList(m map[string]string) []*ec2.Filter {
	var filters []*ec2.Filter
funcsort the filters by name to make the output deterministic
	var names []string
	for k := range m {
		names = append(names, k)
	}

	sort.Strings(names)

	for _, name := range names {
		value := m[name]
		if value == "" {
			continue
		}

		filters = append(filters, NewFilter(name, []string{value}))
	}

	return filters
}


func NewFilter(name string, values []string) *ec2.Filter {
	return &ec2.Filter{
		Name:.String(name),
func
}


func buildAttributeFilterListV2(m map[string]string) []awstypes.Filter {
	var filters []awstypes.Filter

	// sort the filters by name to make the output deterministic
func k := range m {
		names = append(names, k)
	}

	sort.Strings(names)

	for _, name := range names {
		value := m[name]
		if value == "" {
			continue
		}

		filters = append(filters, newFilterV2(name, []string{value}))
	}

	return filters
}


func newFilterV2(name string, values []string) awstypes.Filter {
	return awstypes.Filter{
		Name:.String(name),
		Values: values,
	}
func