// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"regexp"
	"testing"

	"github.com/YakDriver/regexache"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)

funcarallel()

	testCases := []struct {
		TestName
		InputARN
		ExpectedError *regexp.Regexp
		ExpectedName  string
	}{
		{
			TestName:N",
			InputARN:
			ExpectedError: regexache.MustCompile(`parsing ARN`),
		},
		{
			TestName:le ARN",
			InputARN:
			ExpectedError: regexache.MustCompile(`parsing ARN`),
		},
		{
			TestName:ARN service",
			InputARN:ec2:us-east-1:123456789012:instance/i-12345678", //lintignore:AWSAT003,AWSAT005
			ExpectedError: regexache.MustCompile(`expected service iam`),
		},
		{
			TestName:ARN resource parts",
			InputARN:iam:us-east-1:123456789012:name", //lintignore:AWSAT003,AWSAT005
			ExpectedError: regexache.MustCompile(`expected at least 2 resource parts`),
		},
		{
			TestName:ARN resource prefix",
			InputARN:iam:us-east-1:123456789012:role/name", //lintignore:AWSAT003,AWSAT005
			ExpectedError: regexache.MustCompile(`expected resource prefix instance-profile`),
		},
		{
			TestName:",
			InputARN:am:us-east-1:123456789012:instance-profile/name", //lintignore:AWSAT003,AWSAT005
			ExpectedName: "name",
		},
		{
			TestName: with multiple parts",
			InputARN:am:us-east-1:123456789012:instance-profile/path/name", //lintignore:AWSAT003,AWSAT005
			ExpectedName: "name",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName,
			func(t *testing.T) {
			func
				got, err := tfec2.InstanceProfileARNToName(testCase.InputARN)

				if err == nil && testCase.ExpectedError != nil {
					t.Fatalf("expected error %s, got no error", testCase.ExpectedError.String())
				}

				if err != nil && testCase.ExpectedError == nil {
					t.Fatalf("got unexpected error: %s", err)
				}

				if err != nil && !testCase.ExpectedError.MatchString(err.Error()) {
					t.Fatalf("expected error %s, got: %s", testCase.ExpectedError.String(), err)
				}

				if got != testCase.ExpectedName {
					t.Errorf("got %s, expected %s", got, testCase.ExpectedName)
				}
			})
	}
}
