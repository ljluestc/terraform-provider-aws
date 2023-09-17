// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package quicksightimport (
"context"
"fmt"
"time""github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/awserr"
"github.com/aws/aws-sdk-go/service/quicksight"
"github.com/hashicorp/go-multierror"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)const (
iamPropagationTimeout= 2 * time.Minute
dataSourceCreateTimeout = 5 * time.Minute
dataSourceUpdateTimeout = 5 * time.Minute
)// waitCreated waits for a DataSource to return CREATION_SUCCESSFULfunc waitCreated(ctx context.Context, conn *quicksight.QuickSight, accountId, dataSourceId string) (*quicksight.DataSource, error) {
stateConf := &retry.StateChangeConf{
Pending: []string{quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusCreationSuccessful},
fresh: status(ctx, conn, accountId, dataSourceId),
meout: dataSourceCreateTimeout,
}outputRaw, err := stateConf.WaitForStateContext(ctx)if output, ok := outputRaw.(*quicksight.DataSource); ok {
 status, errorInfo := aws.StringValue(output.Status), output.ErrorInfo; status == quicksight.ResourceStatusCreationFailed && errorInfo != nil {
tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(errorInfo.Type), aws.StringValue(errorInfo.Message)))
ern output, err
}return nil, err
}// waitUpdated waits for a DataSource to return UPDATE_SUCCESSFULfunc waitUpdated(ctx context.Context, conn *quicksight.QuickSight, accountId, dataSourceId string) (*quicksight.DataSource, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusUpdateInProgress},
rget:  []string{quicksight.ResourceStatusUpdateSuccessful},
fresh: status(ctx, conn, accountId, dataSourceId),
meout: dataSourceUpdateTimeout,
}outputRaw, err := stateConf.WaitForStateContext(ctx)if output, ok := outputRaw.(*quicksight.DataSource); ok {
 status, errorInfo := aws.StringValue(output.Status), output.ErrorInfo; status == quicksight.ResourceStatusUpdateFailed && errorInfo != nil {
tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(errorInfo.Type), aws.StringValue(errorInfo.Message)))
ern output, err
}return nil, err
}func waitTemplateCreated(ctx context.Context, conn *quicksight.QuickSight, id string, timeout time.Duration) (*quicksight.Template, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusCreationSuccessful},
fresh: statusTemplate(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
}outputRaw, err := stateConf.WaitForStateContext(ctx)
if out, ok := outputRaw.(*quicksight.Template); ok {
 status, apiErrors := aws.StringValue(out.Version.Status), out.Version.Errors; status == quicksight.ResourceStatusCreationFailed && apiErrors != nil {
var errors *multierror.Errorfor _, apiError := range apiErrors {
if apiError == nil {
continue
}
errors = multierror.Append(errors, awserr.New(aws.StringValue(apiError.Type), aws.StringValue(apiError.Message), nil))
}
tfresource.SetLastError(err, errors)
ern out, err
}return nil, err
}func waitTemplateUpdated(ctx context.Context, conn *quicksight.QuickSight, id string, timeout time.Duration) (*quicksight.Template, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusUpdateInProgress, quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusUpdateSuccessful, quicksight.ResourceStatusCreationSuccessful},
fresh: statusTemplate(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
}outputRaw, err := stateConf.WaitForStateContext(ctx)
if out, ok := outputRaw.(*quicksight.Template); ok {
 status, apiErrors := aws.StringValue(out.Version.Status), out.Version.Errors; status == quicksight.ResourceStatusCreationFailed && apiErrors != nil {
var errors *multierror.Errorfor _, apiError := range apiErrors {
if apiError == nil {
continue
}
errors = multierror.Append(errors, awserr.New(aws.StringValue(apiError.Type), aws.StringValue(apiError.Message), nil))
}
tfresource.SetLastError(err, errors)
ern out, err
}return nil, err
}func waitDashboardCreated(ctx context.Context, conn *quicksight.QuickSight, id string, timeout time.Duration) (*quicksight.Dashboard, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusCreationSuccessful},
fresh: statusDashboard(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
}outputRaw, err := stateConf.WaitForStateContext(ctx)
if out, ok := outputRaw.(*quicksight.Dashboard); ok {
 status, apiErrors := aws.StringValue(out.Version.Status), out.Version.Errors; status == quicksight.ResourceStatusCreationFailed && apiErrors != nil {
var errors *multierror.Errorfor _, apiError := range apiErrors {
if apiError == nil {
continue
}
errors = multierror.Append(errors, awserr.New(aws.StringValue(apiError.Type), aws.StringValue(apiError.Message), nil))
}
tfresource.SetLastError(err, errors)
ern out, err
}return nil, err
}func waitDashboardUpdated(ctx context.Context, conn *quicksight.QuickSight, id string, timeout time.Duration) (*quicksight.Dashboard, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusUpdateInProgress, quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusUpdateSuccessful, quicksight.ResourceStatusCreationSuccessful},
fresh: statusDashboard(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
}outputRaw, err := stateConf.WaitForStateContext(ctx)
if out, ok := outputRaw.(*quicksight.Dashboard); ok {
 status, apiErrors := aws.StringValue(out.Version.Status), out.Version.Errors; status == quicksight.ResourceStatusCreationFailed && apiErrors != nil {
var errors *multierror.Errorfor _, apiError := range apiErrors {
if apiError == nil {
continue
}
errors = multierror.Append(errors, awserr.New(aws.StringValue(apiError.Type), aws.StringValue(apiError.Message), nil))
}
tfresource.SetLastError(err, errors)
ern out, err
}return nil, err
}func waitAnalysisCreated(ctx context.Context, conn *quicksight.QuickSight, id string, timeout time.Duration) (*quicksight.Analysis, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusCreationSuccessful},
fresh: statusAnalysis(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
}outputRaw, err := stateConf.WaitForStateContext(ctx)
if out, ok := outputRaw.(*quicksight.Analysis); ok {
 status, apiErrors := aws.StringValue(out.Status), out.Errors; status == quicksight.ResourceStatusCreationFailed && apiErrors != nil {
var errors *multierror.Errorfor _, apiError := range apiErrors {
if apiError == nil {
continue
}
errors = multierror.Append(errors, awserr.New(aws.StringValue(apiError.Type), aws.StringValue(apiError.Message), nil))
}
tfresource.SetLastError(err, errors)
ern out, err
}return nil, err
}func waitAnalysisUpdated(ctx context.Context, conn *quicksight.QuickSight, id string, timeout time.Duration) (*quicksight.Analysis, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusUpdateInProgress, quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusUpdateSuccessful, quicksight.ResourceStatusCreationSuccessful},
fresh: statusAnalysis(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
}outputRaw, err := stateConf.WaitForStateContext(ctx)
if out, ok := outputRaw.(*quicksight.Analysis); ok {
 status, apiErrors := aws.StringValue(out.Status), out.Errors; status == quicksight.ResourceStatusCreationFailed && apiErrors != nil {
var errors *multierror.Errorfor _, apiError := range apiErrors {
if apiError == nil {
continue
}
errors = multierror.Append(errors, awserr.New(aws.StringValue(apiError.Type), aws.StringValue(apiError.Message), nil))
}
tfresource.SetLastError(err, errors)
ern out, err
}return nil, err
}func waitThemeCreated(ctx context.Context, conn *quicksight.QuickSight, id string, timeout time.Duration) (*quicksight.Theme, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusCreationSuccessful},
fresh: statusTheme(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
}outputRaw, err := stateConf.WaitForStateContext(ctx)
if out, ok := outputRaw.(*quicksight.Theme); ok {
 status, apiErrors := aws.StringValue(out.Version.Status), out.Version.Errors; status == quicksight.ResourceStatusCreationFailed && apiErrors != nil {
var errors *multierror.Errorfor _, apiError := range apiErrors {
if apiError == nil {
continue
}
errors = multierror.Append(errors, awserr.New(aws.StringValue(apiError.Type), aws.StringValue(apiError.Message), nil))
}
tfresource.SetLastError(err, errors)
ern out, err
}return nil, err
}func waitThemeUpdated(ctx context.Context, conn *quicksight.QuickSight, id string, timeout time.Duration) (*quicksight.Theme, error) {
stateConf := &retry.StateChangeConf{
nding: []string{quicksight.ResourceStatusUpdateInProgress, quicksight.ResourceStatusCreationInProgress},
rget:  []string{quicksight.ResourceStatusUpdateSuccessful, quicksight.ResourceStatusCreationSuccessful},
fresh: statusTheme(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
}outputRaw, err := stateConf.WaitForStateContext(ctx)
if out, ok := outputRaw.(*quicksight.Theme); ok {
 status, apiErrors := aws.StringValue(out.Version.Status), out.Version.Errors; status == quicksight.ResourceStatusCreationFailed && apiErrors != nil {
var errors *multierror.Errorfor _, apiError := range apiErrors {
if apiError == nil {
continue
}
errors = multierror.Append(errors, awserr.New(aws.StringValue(apiError.Type), aws.StringValue(apiError.Message), nil))
}
tfresource.SetLastError(err, errors)
ern out, err
}return nil, err
}
