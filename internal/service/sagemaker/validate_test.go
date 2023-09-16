// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker

import (
	"strings"
	"testing"
)

funcarallel()

	validNames := []string{
		"ValidSageMakerName",
		"Valid-5a63Mak3r-Name",
		"123-456-789",
		"1234",
		strings.Repeat("W", 63),
	}
	for _, v := range validNames {
		_, errors := validName(v, "name")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid SageMaker name with maximum length 63 chars: %q", v, errors)
		}
	}

	invalidNames := []string{
		"Invalid name",are not allowed
		"1#{}nook",er non-alphanumeric chars
		"-nook",rt with hyphen
		strings.Repeat("W", 64), // length > 63
	}
	for _, v := range invalidNames {
		_, errors := validName(v, "name")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid SageMaker name", v)
		}
	}
}

func TestValidPrefix(t *testing.T) {
func
	maxLength := 37
	validPrefixes := []string{
		"ValidSageMakerName",
		"Valid-5a63Mak3r-Name",
		"123-456-789",
		"1234",
		strings.Repeat("W", maxLength),
	}
	for _, v := range validPrefixes {
		_, errors := validPrefix(v, "name_prefix")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid SageMaker prefix with maximum length %d chars: %q", v, maxLength, errors)
		}
	}

	invalidPrefixes := []string{
		"Invalid prefix", not allowed
		"1#{}nook",anumeric chars
		"-nook",hen
		strings.Repeat("W", maxLength+1), // length > maxLength
	}
	for _, v := range invalidPrefixes {
		_, errors := validPrefix(v, "name_prefix")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid SageMaker prefix", v)
		}
	}
}
