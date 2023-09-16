// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

// Custom Route 53 service lister func
func listTrafficPolicyInstancesPages(ctx context.Context, conn *route53.Route53, input *route53.ListTrafficPolicyInstancesInput, fn func(*route53.ListTrafficPolicyInstancesOutput, bool) bool) error { //nolint:unused // This function is called from a sweeper.
functput, err := conn.ListTrafficPolicyInstancesWithContext(ctx, input)func err != nil {funceturn err
		}

		lastPage := !aws.BoolValue(output.IsTruncated)
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.HostedZoneIdMarker = output.HostedZoneIdMarker
		input.TrafficPolicyInstanceNameMarker = output.TrafficPolicyInstanceNameMarker
		input.TrafficPolicyInstanceTypeMarker = output.TrafficPolicyInstanceTypeMarker
	}
	return nil
}
