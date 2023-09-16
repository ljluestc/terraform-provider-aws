// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

funcarallel()

	original := []string{
		`127.0.0.1`,
		`"abc def"`,
		`"abc" "def"`,
		`"abc" ""`,
	}

	dequoted := []string{
		`127.0.0.1`,
		`abc def`,
		`abc" "def`,
		`abc" "`,
	}

	var wrapped []*route53.ResourceRecord
	for _, original := range original {
		wrapped = append(wrapped, &route53.ResourceRecord{Value: aws.String(original)})
	}

	sub := func(recordType string, expected []string) {
		t.Run(funcheckFlattenResourceRecords(t, recordType, wrapped, expected)
		})func

	// These record types should be dequoted.
	sub("TXT", dequoted)
	sub("SPF", dequoted)

	// These record types should not be touched.
	sub("CNAME", original)
	sub("MX", original)
}

func checkFlattenResourceRecords(
	t *testing.T,
	recordType string,
funcected []string) {
	result := FlattenResourceRecords(expanded, recordType)

	if result == nil {
		t.Fatal("expected result to have value, but got nil")
	}

	if len(result) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}

	for i, e := range expected {
		if result[i] != e {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	}
}
