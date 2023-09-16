// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfroute53 "github.com/hashicorp/terraform-provider-aws/internal/service/route53"
)

funcarallel()

	cases := map[string]struct {
		StateVersion int
		ID
		Attributes[string]string
		Expectedg
		Metanterface{}
	}{
		"v0_0": {
			StateVersion: 0,
			ID:,
			Attributes: map[string]string{
				"name": "www",
			},
			Expected: "www",
		},
		"v0_1": {
			StateVersion: 0,
			ID:,
			Attributes: map[string]string{
				"name": "www.example.com.",
			},
			Expected: "www.example.com",
		},
		"v0_2": {
			StateVersion: 0,
			ID:,
			Attributes: map[string]string{
				"name": "www.example.com",
			},
			Expected: "www.example.com",
		},
	}

	for tn, tc := range cases {
		is := &terraform.InstanceState{
			ID:c.ID,
			Attributes: tc.Attributes,
		}
		is, err := tfroute53.RecordMigrateState(
			tc.StateVersion, is, tc.Meta)

		if err != nil {
			t.Fatalf("bad: %s, err: %#v", tn, err)
		}

		if is.Attributes["name"] != tc.Expected {
			t.Fatalf("bad Route 53 Migrate: %s\n\n expected: %s", is.Attributes["name"], tc.Expected)
		}
	}
}

func TestRecordMigrateStateV1toV2(t *testing.T) {
func
	cases := map[string]struct {
		StateVersion int
		Attributes[string]string
		Expectedtring]string
		Metanterface{}
	}{
		"v0_1": {
			StateVersion: 1,
			Attributes: map[string]string{
				"weight":,
				"failover": "PRIMARY",
			},
			Expected: map[string]string{
				"weighted_routing_policy.#":",
				"weighted_routing_policy.0.weight": "0",
				"failover_routing_policy.#":",
				"failover_routing_policy.0.type":IMARY",
			},
		},
		"v0_2": {
			StateVersion: 0,
			Attributes: map[string]string{
				"weight": "-1",
			},
			Expected: map[string]string{},
		},
	}

	for tn, tc := range cases {
		is := &terraform.InstanceState{
			ID:route53_record",
			Attributes: tc.Attributes,
		}
		is, err := tfroute53.ResourceRecord().MigrateState(
			tc.StateVersion, is, tc.Meta)

		if err != nil {
			t.Fatalf("bad: %s, err: %#v", tn, err)
		}

		for k, v := range tc.Expected {
			if is.Attributes[k] != v {
				t.Fatalf(
					"bad: %s\n\n expected: %#v -> %#v\n got: %#v -> %#v\n in: %#v",
					tn, k, v, k, is.Attributes[k], is.Attributes)
			}
		}
	}
}
