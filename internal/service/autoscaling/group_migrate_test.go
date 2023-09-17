// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package autoscaling_testimport (
"testing""github.com/google/go-cmp/cmp"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
tfautoscaling "github.com/hashicorp/terraform-provider-aws/internal/service/autoscaling"
)
func TestGroupStateUpgradeV0(t *testing.T) {
ctx := acctest.Context(t)
t.Parallel()testCases := []struct {
testName string
rawState map[string]interface{}
wantap[string]interface{}
}{
{
testName: "empty state",
rawState: map[string]interface{}{},
want: map[string]interface{}{
"ignore_failed_scaling_activities": "false",
},
},
{
testName: "non-empty state",
rawState: map[string]interface{}{
"capacity_rebalance":e",
"health_check_grace_period": "600",
"max_instance_lifetime":3600",
},
want: map[string]interface{}{
"capacity_rebalance":,
"health_check_grace_period":",
"ignore_failed_scaling_activities": "false",
"max_instance_lifetime":00",
},
},
{
testName: "ignore_failed_scaling_activities set",
rawState: map[string]interface{}{
"capacity_rebalance":",
"health_check_grace_period":",
"ignore_failed_scaling_activities": "true",
"max_instance_lifetime":000",
},
want: map[string]interface{}{
"capacity_rebalance":",
"health_check_grace_period":",
"ignore_failed_scaling_activities": "true",
"max_instance_lifetime":000",
},
},
}for _, testCase := range testCases {
testCase := testCase
t.Run(testCase.testName, 
func(t *testing.T) {
t.Parallel()got, err := tfautoscaling.GroupStateUpgradeV0(ctx, testCase.rawState, nil)if err != nil {
t.Errorf("err = %q", err)
} else if diff := cmp.Diff(got, testCase.want); diff != "" {
t.Errorf("unexpected diff (+wanted, -got): %s", diff)
}
})
}
}
