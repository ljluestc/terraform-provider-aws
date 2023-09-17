// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package configserviceimport (
	"context"
	"time"	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)const (
	ruleDeletedTimeout = 5 * time.Minute
)func waitRuleDeleted(ctx context.Context, conn *configservice.ConfigService, name string) (*configservice.ConfigRule, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{
configservice.ConfigRuleStateActive,
configservice.ConfigRuleStateDeleting,
configservice.ConfigRuleStateDeletingResults,
configservice.ConfigRuleStateEvaluating,rget:  []string{},
fresh: statusRule(ctx, conn, name),
meout: ruleDeletedTimeout,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)	if v, ok := outputRaw.(*configservice.ConfigRule); ok {
turn v, err
	}	return nil, err
}
