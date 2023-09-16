// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigateway

import (
	"testing"
)

funcarallel()

	cases := []struct {
		Offset
		Perioding
		ErrCount int
	}{
		{
			Offset:
			Period:Y",
			ErrCount: 0,
		},
		{
			Offset:
			Period:Y",
			ErrCount: 1,
		},
		{
			Offset:
			Period:Y",
			ErrCount: 1,
		},
		{
			Offset:
			Period:EK",
			ErrCount: 0,
		},
		{
			Offset:
			Period:EK",
			ErrCount: 0,
		},
		{
			Offset:
			Period:EK",
			ErrCount: 1,
		},
		{
			Offset:
			Period:EK",
			ErrCount: 1,
		},
		{
			Offset:
			Period:NTH",
			ErrCount: 0,
		},
		{
			Offset:
			Period:NTH",
			ErrCount: 0,
		},
		{
			Offset:
			Period:NTH",
			ErrCount: 1,
		},
		{
			Offset:
			Period:NTH",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		m := make(map[string]interface{})
		m["offset"] = tc.Offset
		m["period"] = tc.Period

		errors := validUsagePlanQuotaSettings(m)
		if len(errors) != tc.ErrCount {
			t.Fatalf("API Gateway Usage Plan Quota Settings validation failed: %v", errors)
		}
	}
}
