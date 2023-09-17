// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package ssmincidents_testimport (
	"testing"	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)// only one replication set resource can be active at once, so we must have serialised tests
func TestAccSSMIncidents_serial(t *testing.T) {
	t.Parallel()	testCases := map[string]map[string]func(t *testing.T){
"Replication Set Resource Tests": {
"basic":testReplicationSet_basic,
"updateDefaultKey": testReplicationSet_updateRegionsWithoutCMK,
"updateCMK":eRegionsWithCMK,
"updateTags":testReplicationSet_updateTags,
"updateEmptyTags":  testReplicationSet_updateEmptyTags,
"disappears":testReplicationSet_disappears,eplication Set Data Source Tests": {
"basic": testReplicationSetDataSource_basic,esponse Plan Resource Tests": {
"basic":testResponsePlan_basic,
"update":testResponsePlan_updateRequiredFields,
"updateTags":testResponsePlan_updateTags,
"updateEmptyTags":mptyTags,
"disappears":testResponsePlan_disappears,
"incidentTemplateFields": testResponsePlan_incidentTemplateOptionalFields,
"displayName":testResponsePlan_displayName,
"chatChannel":testResponsePlan_chatChannel,
"engagement":testResponsePlan_engagement,
"action":testResponsePlan_action,esponse Plan Data Source Tests": {
"basic": testResponsePlanDataSource_basic,	}	acctest.RunSerialTests2Levels(t, testCases, 0)
}
