// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package emrimport (
	"testing"
)func TestValidCustomAMIID(t *testing.T) {
	t.Parallel()	cases := []struct {
Valuestring
rCount int
	}{Value:"ami-dbcf88b1", //lintignore:AWSAT002
ErrCount: 0,
Value:"vol-as7d65ash",
ErrCount: 1,	}	for _, tc := range cases {
 errors := validCustomAMIID(tc.Value, "custom_ami_id")ifen(errors) != tc.ErrCount {
t.Fatalf("Expected %d errors, got %d: %s", tc.ErrCount, len(errors), errors)	}
}
