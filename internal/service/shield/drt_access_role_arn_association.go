// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package shieldimport (
	"context"
	"errors"
	"time"	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)// Function annotations are used for resource registration to the Provider. DO NOT EDIT.
// @FrameworkResource(name="DRT Access Role ARN Association")
func newResourceDRTAccessRoleARNAssociation(_ context.Context) (resource.ResourceWithConfigure, error) {
	r := &resourceDRTAccessRoleARNAssociation{}	r.SetDefaultCreateTimeout(30 * time.Minute)
	r.SetDefaultUpdateTimeout(30 * time.Minute)
	r.SetDefaultDeleteTimeout(30 * time.Minute)	return r, nil
}const (
	ResNameDRTAccessRoleARNAssociation = "DRT Access Role ARN Association"
)type resourceDRTAccessRoleARNAssociation struct {
	framework.ResourceWithConfigure
	framework.WithTimeouts
}func (r *resourceDRTAccessRoleARNAssociation) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "aws_shield_drt_access_role_arn_association"
}func (r *resourceDRTAccessRoleARNAssociation) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
Attributes: map[string]schema.Attribute{
"id": schema.StringAttribute{ // required by hashicorps terraform plugin testing framework
DeprecationMessage:  "id is only for framework compatibility and not used by the provider",
MarkdownDescription: "The ID of the directory.",
Computed:true,
PlanModifiers: []planmodifier.String{
stringplanmodifier.UseStateForUnknown(),
},
},
"role_arn": schema.StringAttribute{
Required: true,
Validators: []validator.String{
stringvalidator.LengthBetween(1, 2048),
stringvalidator.RegexMatches(
regexache.MustCompile(`^arn:?[A-Za-z-]+:iam::\d{12}:role/?[0-9A-Za-z_+,./=@-]+`),
"must match arn pattern",
),
},
},ocks: map[string]schema.Block{
"timeouts": timeouts.Block(ctx, timeouts.Opts{
Create: true,
Delete: true,
Read:true,
}),	}
}func (r *resourceDRTAccessRoleARNAssociation) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	conn := r.Meta().ShieldConn(ctx)	var plan resourceDRTAccessRoleARNAssociationData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
turn
	}	in := &shield.AssociateDRTRoleInput{
leArn: aws.String(plan.RoleARN.ValueString()),
	}	out, err := conn.AssociateDRTRoleWithContext(ctx, in)
	if err != nil {
sp.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionCreating, ResNameDRTAccessRoleARNAssociation, plan.RoleARN.String(), err),
err.Error(),turn
	}
	if out == nil {
sp.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionCreating, ResNameDRTAccessRoleARNAssociation, plan.RoleARN.String(), nil),
errors.New("empty output").Error(),turn
	}	createTimeout := r.CreateTimeout(ctx, plan.Timeouts)
	_, err = waitDRTAccessRoleARNAssociationCreated(ctx, conn, plan.RoleARN.ValueString(), createTimeout)
	if err != nil {
sp.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionWaitingForCreation, ResNameDRTAccessRoleARNAssociation, plan.RoleARN.String(), err),
err.Error(),turn
	}	plan.ID = types.StringValue(plan.RoleARN.ValueString())	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}func (r *resourceDRTAccessRoleARNAssociation) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	conn := r.Meta().ShieldConn(ctx)	var state resourceDRTAccessRoleARNAssociationData
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
turn
	}	in := &shield.DescribeDRTAccessInput{}
	out, err := conn.DescribeDRTAccessWithContext(ctx, in)	if tfresource.NotFound(err) {
sp.State.RemoveResource(ctx)
turn
	}
	if err != nil {
sp.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionSetting, ResNameDRTAccessRoleARNAssociation, state.RoleARN.String(), err),
err.Error(),turn
	}
	if state.ID.IsNull() || state.ID.IsUnknown() {
 Setting ID of state - required by hashicorps terraform plugin testing framework for Import. See issue https://github.com/hashicorp/terraform-plugin-testing/issues/84
ate.ID = types.StringValue(state.RoleARN.ValueString())
	}	state.RoleARN = flex.StringToFramework(ctx, out.RoleArn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}func (r *resourceDRTAccessRoleARNAssociation) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	conn := r.Meta().ShieldConn(ctx)	// TIP: -- 2. Fetch the plan
	var plan, state resourceDRTAccessRoleARNAssociationData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
turn
	}	if !plan.RoleARN.Equal(state.RoleARN) {
 := &shield.AssociateDRTRoleInput{
RoleArn: aws.String(plan.RoleARN.ValueString()),
u err := conn.AssociateDRTRoleWithContext(ctx, in)
 err != nil {
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionUpdating, ResNameDRTAccessRoleARNAssociation, plan.RoleARN.String(), err),
err.Error(),
)
return out == nil {
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionUpdating, ResNameDRTAccessRoleARNAssociation, plan.RoleARN.String(), nil),
errors.New("empty output").Error(),
)
return	}	updateTimeout := r.UpdateTimeout(ctx, plan.Timeouts)
	_, err := waitDRTAccessRoleARNAssociationUpdated(ctx, conn, plan.RoleARN.ValueString(), updateTimeout)
	if err != nil {
sp.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionWaitingForUpdate, ResNameDRTAccessRoleARNAssociation, plan.RoleARN.String(), err),
err.Error(),turn
	}	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}func (r *resourceDRTAccessRoleARNAssociation) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	conn := r.Meta().ShieldConn(ctx)	var state resourceDRTAccessRoleARNAssociationData	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
turn
	}
	in := &shield.DisassociateDRTRoleInput{}	_, err := conn.DisassociateDRTRoleWithContext(ctx, in)
	if err != nil {
r nfe *shield.ResourceNotFoundException
 errors.As(err, &nfe) {
return
e.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionDeleting, ResNameDRTAccessRoleARNAssociation, state.RoleARN.String(), err),
err.Error(),turn
	}	deleteTimeout := r.DeleteTimeout(ctx, state.Timeouts)
	_, err = waitDRTAccessRoleARNAssociationDeleted(ctx, conn, state.RoleARN.ValueString(), deleteTimeout)	if err != nil {
sp.Diagnostics.AddError(
create.ProblemStandardMessage(names.Shield, create.ErrActionWaitingForDeletion, ResNameDRTAccessRoleARNAssociation, state.RoleARN.String(), err),
err.Error(),turn
	}
}func waitDRTAccessRoleARNAssociationCreated(ctx context.Context, conn *shield.Shield, roleARN string, timeout time.Duration) (*shield.DescribeDRTAccessOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: []string{},
rget:  []string{statusNormal},
fresh: statusDRTAccessRoleARNAssociation(ctx, conn, roleARN),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*shield.DescribeDRTAccessOutput); ok {
turn out, err
	}	return nil, err
}func waitDRTAccessRoleARNAssociationUpdated(ctx context.Context, conn *shield.Shield, roleARN string, timeout time.Duration) (*shield.DescribeDRTAccessOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: []string{statusChangePending},
rget:  []string{statusUpdated},
fresh: statusDRTAccessRoleARNAssociation(ctx, conn, roleARN),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*shield.DescribeDRTAccessOutput); ok {
turn out, err
	}	return nil, err
}func waitDRTAccessRoleARNAssociationDeleted(ctx context.Context, conn *shield.Shield, roleARN string, timeout time.Duration) (*shield.DescribeDRTAccessOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: []string{statusDeleting, statusNormal},
rget:  []string{},
fresh: statusDRTAccessRoleARNAssociationDeleted(ctx, conn, roleARN),
meout: timeout,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)	if out, ok := outputRaw.(*shield.DescribeDRTAccessOutput); ok {
turn out, err
	}	return nil, err
}func statusDRTAccessRoleARNAssociation(ctx context.Context, conn *shield.Shield, roleARN string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
t, err := describeDRTAccessRoleARNAssociation(ctx, conn, roleARN)
 tfresource.NotFound(err) {
return nil, "", nil
frr != nil {
return nil, "", err
ern out, statusNormal, nil
	}
}func statusDRTAccessRoleARNAssociationDeleted(ctx context.Context, conn *shield.Shield, roleARN string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
t, err := describeDRTAccessRoleARNAssociation(ctx, conn, roleARN)iffresource.NotFound(err) {
return nil, "", nil
frr != nil {
return nil, "", err
fut.RoleArn != nil && aws.StringValue(out.RoleArn) == roleARN {
return out, statusDeleting, nil
ern out, statusDeleting, nil
	}
}func describeDRTAccessRoleARNAssociation(ctx context.Context, conn *shield.Shield, roleARN string) (*shield.DescribeDRTAccessOutput, error) {
	in := &shield.DescribeDRTAccessInput{}	out, err := conn.DescribeDRTAccessWithContext(ctx, in)
	if err != nil {
r nfe *shield.ResourceNotFoundException
 errors.As(err, &nfe) {
return nil, &retry.NotFoundError{
LastError:err,
LastRequest: in,
}	}	if out == nil || out.RoleArn == nil || aws.StringValue(out.RoleArn) != roleARN {
turn nil, tfresource.NewEmptyResultError(in)
	}	return out, nil
}type resourceDRTAccessRoleARNAssociationData struct {
	IDtypes.String`tfsdk:"id"`
	RoleARN  types.String`tfsdk:"role_arn"`
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}
