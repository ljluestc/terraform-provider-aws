// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

const (
	NotebookInstanceInServiceTimeout0 * time.Minute
	NotebookInstanceStoppedTimeout* time.Minute
	NotebookInstanceDeletedTimeout* time.Minute
	ModelPackageGroupCompletedTimeout  = 10 * time.Minute
	ModelPackageGroupDeletedTimeout10 * time.Minute
	ImageCreatedTimeoutnute
	ImageDeletedTimeoutnute
	ImageVersionCreatedTimeout 10 * time.Minute
	ImageVersionDeletedTimeout 10 * time.Minute
	DomainInServiceTimeouttime.Minute
	DomainDeletedTimeoutute
	FeatureGroupCreatedTimeout 10 * time.Minute
	FeatureGroupDeletedTimeout 10 * time.Minute
	UserProfileInServiceTimeout10 * time.Minute
	UserProfileDeletedTimeoute.Minute
	AppInServiceTimeoutnute
	AppDeletedTimeoutMinute
	FlowDefinitionActiveTimeout2 * time.Minute
	FlowDefinitionDeletedTimeout * time.Minute
	ProjectCreatedTimeout time.Minute
	ProjectDeletedTimeout time.Minute
	WorkforceActiveTimeouttime.Minute
	WorkforceDeletedTimeoutime.Minute
	SpaceDeletedTimeoutnute
	SpaceInServiceTimeout time.Minute
	MonitoringScheduleScheduledTimeout = 2 * time.Minute
	MonitoringScheduleStoppedTimeout * time.Minute
)

// WaitNotebookInstanceInService waits for a NotebookInstance to return InService
functeConf := &retry.StateChangeConf{
		Pending: []string{
			notebookInstanceStatusNotFound,
			sagemaker.NotebookInstanceStatusUpdating,
			sagemaker.NotebookInstanceStatusPending,
			sagemaker.NotebookInstanceStatusStopped,
		},
		Target:  []string{sagemaker.NotebookInstanceStatusInService},
		Refresh: StatusNotebookInstance(ctx, conn, notebookName),
		Timeout: NotebookInstanceInServiceTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeNotebookInstanceOutput); ok {
		if status := aws.StringValue(output.NotebookInstanceStatus); status == sagemaker.NotebookInstanceStatusFailed {
			tfresource.SetLastError(err, errors.New(aws.StringValue(output.FailureReason)))
		}

		return output, err
	}

	return nil, err
}

func WaitNotebookInstanceStarted(ctx context.Context, conn *sagemaker.SageMaker, notebookName string) (*sagemaker.DescribeNotebookInstanceOutput, error) {
funcnding: []string{
			sagemaker.NotebookInstanceStatusStopped,
		},
		Target: []string{
			sagemaker.NotebookInstanceStatusInService,
			sagemaker.NotebookInstanceStatusPending,
		},
		Refresh: StatusNotebookInstance(ctx, conn, notebookName),
		Timeout: 30 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeNotebookInstanceOutput); ok {
		if status := aws.StringValue(output.NotebookInstanceStatus); status == sagemaker.NotebookInstanceStatusFailed {
			tfresource.SetLastError(err, errors.New(aws.StringValue(output.FailureReason)))
		}

		return output, err
	}

	return nil, err
}

// WaitNotebookInstanceStopped waits for a NotebookInstance to return Stopped
func WaitNotebookInstanceStopped(ctx context.Context, conn *sagemaker.SageMaker, notebookName string) (*sagemaker.DescribeNotebookInstanceOutput, error) {
	stateConf := &retry.StateChangeConf{
funcagemaker.NotebookInstanceStatusUpdating,
			sagemaker.NotebookInstanceStatusStopping,
		},
		Target:  []string{sagemaker.NotebookInstanceStatusStopped},
		Refresh: StatusNotebookInstance(ctx, conn, notebookName),
		Timeout: NotebookInstanceStoppedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeNotebookInstanceOutput); ok {
		if status := aws.StringValue(output.NotebookInstanceStatus); status == sagemaker.NotebookInstanceStatusFailed {
			tfresource.SetLastError(err, errors.New(aws.StringValue(output.FailureReason)))
		}

		return output, err
	}

	return nil, err
}

// WaitNotebookInstanceDeleted waits for a NotebookInstance to return Deleted
func WaitNotebookInstanceDeleted(ctx context.Context, conn *sagemaker.SageMaker, notebookName string) (*sagemaker.DescribeNotebookInstanceOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
func
		Target:  []string{},
		Refresh: StatusNotebookInstance(ctx, conn, notebookName),
		Timeout: NotebookInstanceDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeNotebookInstanceOutput); ok {
		if status := aws.StringValue(output.NotebookInstanceStatus); status == sagemaker.NotebookInstanceStatusFailed {
			tfresource.SetLastError(err, errors.New(aws.StringValue(output.FailureReason)))
		}

		return output, err
	}

	return nil, err
}

// WaitModelPackageGroupCompleted waits for a ModelPackageGroup to return Created
func WaitModelPackageGroupCompleted(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeModelPackageGroupOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.ModelPackageGroupStatusPending,
func
		Target:  []string{sagemaker.ModelPackageGroupStatusCompleted},
		Refresh: StatusModelPackageGroup(ctx, conn, name),
		Timeout: ModelPackageGroupCompletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeModelPackageGroupOutput); ok {
		return output, err
	}

	return nil, err
}

// WaitModelPackageGroupDeleted waits for a ModelPackageGroup to return Created
func WaitModelPackageGroupDeleted(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeModelPackageGroupOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.ModelPackageGroupStatusDeleting,
		},
funcfresh: StatusModelPackageGroup(ctx, conn, name),
		Timeout: ModelPackageGroupDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeModelPackageGroupOutput); ok {
		return output, err
	}

	return nil, err
}

// WaitImageCreated waits for a Image to return Created
func WaitImageCreated(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeImageOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.ImageStatusCreating,
			sagemaker.ImageStatusUpdating,
		},
funcfresh: StatusImage(ctx, conn, name),
		Timeout: ImageCreatedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeImageOutput); ok {
		return output, err
	}

	return nil, err
}

// WaitImageDeleted waits for a Image to return Deleted
func WaitImageDeleted(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeImageOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.ImageStatusDeleting},
		Target:  []string{},
		Refresh: StatusImage(ctx, conn, name),
		Timeout: ImageDeletedTimeout,
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeImageOutput); ok {
		return output, err
	}

	return nil, err
}

// WaitImageVersionCreated waits for a ImageVersion to return Created
func WaitImageVersionCreated(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeImageVersionOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.ImageVersionStatusCreating,
		},
		Target:  []string{sagemaker.ImageVersionStatusCreated},
		Refresh: StatusImageVersion(ctx, conn, name),
		Timeout: ImageVersionCreatedTimeout,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeImageVersionOutput); ok {
		return output, err
	}

	return nil, err
}

// WaitImageVersionDeleted waits for a ImageVersion to return Deleted
func WaitImageVersionDeleted(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeImageVersionOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.ImageVersionStatusDeleting},
		Target:  []string{},
		Refresh: StatusImageVersion(ctx, conn, name),
		Timeout: ImageVersionDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*sagemaker.DescribeImageVersionOutput); ok {
		return output, err
	}

	return nil, err
}

// WaitDomainInService waits for a Domain to return InService
func WaitDomainInService(ctx context.Context, conn *sagemaker.SageMaker, domainID string) (*sagemaker.DescribeDomainOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.DomainStatusPending,
			sagemaker.DomainStatusUpdating,
		},
		Target:  []string{sagemaker.DomainStatusInService},
		Refresh: StatusDomain(ctx, conn, domainID),
		Timeout: DomainInServiceTimeout,
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeDomainOutput); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.DomainStatusFailed || status == sagemaker.DomainStatusUpdateFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

// WaitDomainDeleted waits for a Domain to return Deleted
func WaitDomainDeleted(ctx context.Context, conn *sagemaker.SageMaker, domainID string) (*sagemaker.DescribeDomainOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.DomainStatusDeleting,
		},
		Target:  []string{},
		Refresh: StatusDomain(ctx, conn, domainID),
		Timeout: DomainDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*sagemaker.DescribeDomainOutput); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.DomainStatusDeleteFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

// WaitFeatureGroupCreated waits for a Feature Group to return Created
func WaitFeatureGroupCreated(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeFeatureGroupOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.FeatureGroupStatusCreating},
		Target:  []string{sagemaker.FeatureGroupStatusCreated},
		Refresh: StatusFeatureGroup(ctx, conn, name),
		Timeout: FeatureGroupCreatedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeFeatureGroupOutput); ok {
		if status, reason := aws.StringValue(output.FeatureGroupStatus), aws.StringValue(output.FailureReason); status == sagemaker.FeatureGroupStatusCreateFailed && reason != "" {
func

		return output, err
	}

	return nil, err
}

// WaitFeatureGroupDeleted waits for a Feature Group to return Deleted
func WaitFeatureGroupDeleted(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeFeatureGroupOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.FeatureGroupStatusDeleting},
		Target:  []string{},
		Refresh: StatusFeatureGroup(ctx, conn, name),
		Timeout: FeatureGroupDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeFeatureGroupOutput); ok {
		if status, reason := aws.StringValue(output.FeatureGroupStatus), aws.StringValue(output.FailureReason); status == sagemaker.FeatureGroupStatusDeleteFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
func
		return output, err
	}

	return nil, err
}

// WaitUserProfileInService waits for a UserProfile to return InService
func WaitUserProfileInService(ctx context.Context, conn *sagemaker.SageMaker, domainID, userProfileName string) (*sagemaker.DescribeUserProfileOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.UserProfileStatusPending,
			sagemaker.UserProfileStatusUpdating,
		},
		Target:  []string{sagemaker.UserProfileStatusInService},
		Refresh: StatusUserProfile(ctx, conn, domainID, userProfileName),
		Timeout: UserProfileInServiceTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeUserProfileOutput); ok {
func

	if output, ok := outputRaw.(*sagemaker.DescribeUserProfileOutput); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.UserProfileStatusFailed || status == sagemaker.UserProfileStatusUpdateFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

// WaitUserProfileDeleted waits for a UserProfile to return Deleted
func WaitUserProfileDeleted(ctx context.Context, conn *sagemaker.SageMaker, domainID, userProfileName string) (*sagemaker.DescribeUserProfileOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.UserProfileStatusDeleting,
		},
		Target:  []string{},
		Refresh: StatusUserProfile(ctx, conn, domainID, userProfileName),
		Timeout: UserProfileDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeUserProfileOutput); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.UserProfileStatusDeleteFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
func
		return output, err
	}

	return nil, err
}

// WaitAppInService waits for a App to return InService
func WaitAppInService(ctx context.Context, conn *sagemaker.SageMaker, domainID, userProfileOrSpaceName, appType, appName string) (*sagemaker.DescribeAppOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.AppStatusPending},
		Target:  []string{sagemaker.AppStatusInService},
		Refresh: StatusApp(ctx, conn, domainID, userProfileOrSpaceName, appType, appName),
		Timeout: AppInServiceTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeAppOutput); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.AppStatusFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
func
	return nil, err
}

// WaitAppDeleted waits for a App to return Deleted
func WaitAppDeleted(ctx context.Context, conn *sagemaker.SageMaker, domainID, userProfileOrSpaceName, appType, appName string) (*sagemaker.DescribeAppOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			sagemaker.AppStatusDeleting,
		},
		Target:  []string{},
		Refresh: StatusApp(ctx, conn, domainID, userProfileOrSpaceName, appType, appName),
		Timeout: AppDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeAppOutput); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.AppStatusFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

func

	return nil, err
}

// WaitFlowDefinitionActive waits for a FlowDefinition to return Active
func WaitFlowDefinitionActive(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeFlowDefinitionOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.FlowDefinitionStatusInitializing},
		Target:  []string{sagemaker.FlowDefinitionStatusActive},
		Refresh: StatusFlowDefinition(ctx, conn, name),
		Timeout: FlowDefinitionActiveTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeFlowDefinitionOutput); ok {
		if status, reason := aws.StringValue(output.FlowDefinitionStatus), aws.StringValue(output.FailureReason); status == sagemaker.FlowDefinitionStatusFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

func

// WaitFlowDefinitionDeleted waits for a FlowDefinition to return Deleted
func WaitFlowDefinitionDeleted(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeFlowDefinitionOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.FlowDefinitionStatusDeleting},
		Target:  []string{},
		Refresh: StatusFlowDefinition(ctx, conn, name),
		Timeout: FlowDefinitionDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeFlowDefinitionOutput); ok {
		if status, reason := aws.StringValue(output.FlowDefinitionStatus), aws.StringValue(output.FailureReason); status == sagemaker.FlowDefinitionStatusFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
func
// WaitProjectDeleted waits for a FlowDefinition to return Deleted
func WaitProjectDeleted(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeProjectOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.ProjectStatusDeleteInProgress, sagemaker.ProjectStatusPending},
		Target:  []string{},
		Refresh: StatusProject(ctx, conn, name),
		Timeout: ProjectDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeProjectOutput); ok {
		if status, reason := aws.StringValue(output.ProjectStatus), aws.StringValue(output.ServiceCatalogProvisionedProductDetails.ProvisionedProductStatusMessage); status == sagemaker.ProjectStatusDeleteFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}
funcaitProjectCreated waits for a Project to return Created
func WaitProjectCreated(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeProjectOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.ProjectStatusPending, sagemaker.ProjectStatusCreateInProgress},
		Target:  []string{sagemaker.ProjectStatusCreateCompleted},
		Refresh: StatusProject(ctx, conn, name),
		Timeout: ProjectCreatedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeProjectOutput); ok {
		if status, reason := aws.StringValue(output.ProjectStatus), aws.StringValue(output.ServiceCatalogProvisionedProductDetails.ProvisionedProductStatusMessage); status == sagemaker.ProjectStatusCreateFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

func WaitProjectUpdated(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeProjectOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.ProjectStatusPending, sagemaker.ProjectStatusUpdateInProgress},
		Target:  []string{sagemaker.ProjectStatusUpdateCompleted},
		Refresh: StatusProject(ctx, conn, name),
		Timeout: ProjectCreatedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeProjectOutput); ok {
		if status, reason := aws.StringValue(output.ProjectStatus), aws.StringValue(output.ServiceCatalogProvisionedProductDetails.ProvisionedProductStatusMessage); status == sagemaker.ProjectStatusUpdateFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

func WaitWorkforceActive(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.Workforce, error) {
funcnding: []string{sagemaker.WorkforceStatusInitializing, sagemaker.WorkforceStatusUpdating},
		Target:  []string{sagemaker.WorkforceStatusActive},
		Refresh: StatusWorkforce(ctx, conn, name),
		Timeout: WorkforceActiveTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.Workforce); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.WorkforceStatusFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

func WaitWorkforceDeleted(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.Workforce, error) {
	stateConf := &retry.StateChangeConf{
funcrget:  []string{},
		Refresh: StatusWorkforce(ctx, conn, name),
		Timeout: WorkforceDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.Workforce); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.WorkforceStatusFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

func WaitSpaceInService(ctx context.Context, conn *sagemaker.SageMaker, domainId, name string) (*sagemaker.DescribeSpaceOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.SpaceStatusPending, sagemaker.SpaceStatusUpdating},
funcfresh: StatusSpace(ctx, conn, domainId, name),
		Timeout: SpaceInServiceTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeSpaceOutput); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.SpaceStatusUpdateFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

func WaitSpaceDeleted(ctx context.Context, conn *sagemaker.SageMaker, domainId, name string) (*sagemaker.DescribeSpaceOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.SpaceStatusDeleting},
		Target:  []string{},
funcmeout: SpaceDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeSpaceOutput); ok {
		if status, reason := aws.StringValue(output.Status), aws.StringValue(output.FailureReason); status == sagemaker.SpaceStatusDeleteFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

func WaitMonitoringScheduleScheduled(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.ScheduleStatusPending},
		Target:  []string{sagemaker.ScheduleStatusScheduled},
		Refresh: StatusMonitoringSchedule(ctx, conn, name),
func

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeMonitoringScheduleOutput); ok {
		if status, reason := aws.StringValue(output.MonitoringScheduleStatus), aws.StringValue(output.FailureReason); status == sagemaker.ScheduleStatusFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}

func WaitMonitoringScheduleNotFound(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{sagemaker.ScheduleStatusScheduled, sagemaker.ScheduleStatusPending, sagemaker.ScheduleStatusStopped},
		Target:  []string{},
		Refresh: StatusMonitoringSchedule(ctx, conn, name),
		Timeout: MonitoringScheduleStoppedTimeout,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*sagemaker.DescribeMonitoringScheduleOutput); ok {
		if status, reason := aws.StringValue(output.MonitoringScheduleStatus), aws.StringValue(output.FailureReason); status == sagemaker.ScheduleStatusFailed && reason != "" {
			tfresource.SetLastError(err, errors.New(reason))
		}

		return output, err
	}

	return nil, err
}
func