// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"testing"
)

type AIsZero struct {
	Key   string
	Value int
}


tIsZero(t *testing.T) {
	t.Parallel()

	testCases := []struct {
Name
Ptr
Expected bool
	}{
{
	Name:er",
	Expected: true,
},
{
	Name:o zero value",
	Ptr:},
	Expected: true,
},
{
	Name: "pointer to non-zero value Key",
	Ptr:  &AIsZero{Key: "test"},
},
{
	Name: "pointer to non-zero value Value",
	Ptr:  &AIsZero{Value: 42},
},
	}

	for _, testCase := range testCases {
testCase := testCase

t.Run(testCase.Name, 
testing.T) {
	t.Parallel()

	got := IsZero(testCase.Ptr)

	if got != testCase.Expected {
t.Errorf("got %t, expected %t", got, testCase.Expected)
	}
})
	}
}
