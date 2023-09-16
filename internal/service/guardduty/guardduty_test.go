// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guardduty_test

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)
func TestAccGuardDuty_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Detector": {
			"basic":
			"datasources_s3logs":3logs,
			"datasources_kubernetes_audit_logs": testAccDetector_datasources_kubernetes_audit_logs,
			"datasources_malware_protection": testAccDetector_datasources_malware_protection,
			"datasources_all":s_all,
			"tags":
			"datasource_basic":asic,
			"datasource_id":
		},
		"Filter": {
			"basic":lter_basic,
			"update":ter_update,
			"tags":ilter_tags,
			"disappears": testAccFilter_disappears,
		},
		"InviteAccepter": {
			"basic": testAccInviteAccepter_basic,
		},
		"IPSet": {
			"basic": testAccIPSet_basic,
			"tags":  testAccIPSet_tags,
		},
		"OrganizationAdminAccount": {
			"basic": testAccOrganizationAdminAccount_basic,
		},
		"OrganizationConfiguration": {
			"basic":
			"autoEnableOrganizationMembers": testAccOrganizationConfiguration_autoEnableOrganizationMembers,
			"s3Logs":ogs,
			"kubernetes":tes,
			"malwareProtection":onConfiguration_malwareprotection,
		},
		"ThreatIntelSet": {
			"basic": testAccThreatIntelSet_basic,
			"tags":  testAccThreatIntelSet_tags,
		},
		"Member": {
			"basic":sic,
			"inviteOnUpdate":ber_invite_onUpdate,
			"inviteDisassociate": testAccMember_invite_disassociate,
			"invitationMessage":  testAccMember_invitationMessage,
		},
		"PublishingDestination": {
			"basic":blishingDestination_basic,
			"disappears": testAccPublishingDestination_disappears,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
func testAccMemberFromEnv(t *testing.T) (string, string) {
	accountID := os.Getenv("AWS_GUARDDUTY_MEMBER_ACCOUNT_ID")
	if accountID == "" {
		t.Skip(
			"Environment variable AWS_GUARDDUTY_MEMBER_ACCOUNT_ID is not set. " +
				"To properly test inviting GuardDuty member accounts, " +
				"a valid AWS account ID must be provided.")
	}
	email := os.Getenv("AWS_GUARDDUTY_MEMBER_EMAIL")
	if email == "" {
		t.Skip(
			"Environment variable AWS_GUARDDUTY_MEMBER_EMAIL is not set. " +
				"To properly test inviting GuardDuty member accounts, " +
				"a valid email associated with the AWS_GUARDDUTY_MEMBER_ACCOUNT_ID must be provided.")
	}
	return accountID, email
}

// testAccPreCheckDetectorExists verifies the current account has a single active
// GuardDuty detector configured.
func testAccPreCheckDetectorExists(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).GuardDutyConn(ctx)

	out, err := conn.ListDetectorsWithContext(ctx, &guardduty.ListDetectorsInput{})
	if out == nil || len(out.DetectorIds) == 0 {
		t.Skip("this AWS account must have an existing GuardDuty detector configured")
	}
	if len(out.DetectorIds) > 1 {
		t.Skipf("this AWS account must have a single existing GuardDuty detector configured. Found %d.", len(out.DetectorIds))
	}

	if err != nil {
		t.Fatalf("listing GuardDuty Detectors: %s", err)
	}
}
