//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagequicksight

import(
"context"
"errors"
"fmt"
"strings"
"time"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/quicksight"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
"github.com/hashicorp/terraform-plugin-framework/path"
"github.com/hashicorp/terraform-plugin-framework/resource"
"github.com/hashicorp/terraform-plugin-framework/resource/schema"
"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
"github.com/hashicorp/terraform-plugin-framework/types"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
"github.com/hashicorp/terraform-provider-aws/internal/create"
"github.com/hashicorp/terraform-provider-aws/internal/framework"
"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
"github.com/hashicorp/terraform-provider-aws/names"
)

//@FrameworkResource(name="Namespace")
//@Tags(identifierAttribute="arn")

funcnewResourceNamespace(_context.Context)(resource.ResourceWithConfigure,error){
r:=&resourceNamespace{}
r.SetDefaultCreateTimeout(2*time.Minute)
r.SetDefaultDeleteTimeout(2*time.Minute)

returnr,nil
}

const(
ResNameNamespace="Namespace"
)

typeresourceNamespacestruct{
framework.ResourceWithConfigure
framework.WithTimeouts
}


func(r*resourceNamespace)Metadata(_context.Context,requestresource.MetadataRequest,response*resource.MetadataResponse){
response.TypeName="aws_quicksight_namespace"
}


func(r*resourceNamespace)Schema(ctxcontext.Context,reqresource.SchemaRequest,resp*resource.SchemaResponse){
resp.Schema=schema.Schema{
Attributes:map[string]schema.Attribute{
"arn":framework.ARNAttributeComputedOnly(),
"aws_account_id":schema.StringAttribute{
Optional:true,
Computed:true,
PlanModifiers:[]planmodifier.String{
stringplanmodifier.UseStateForUnknown(),
stringplanmodifier.RequiresReplace(),
},
},
"capacity_region":schema.StringAttribute{
Computed:true,
PlanModifiers:[]planmodifier.String{
stringplanmodifier.UseStateForUnknown(),
},
},
"creation_status":schema.StringAttribute{
Computed:true,
PlanModifiers:[]planmodifier.String{
stringplanmodifier.UseStateForUnknown(),
},
},
"id":framework.IDAttribute(),
"identity_store":schema.StringAttribute{
Optional:true,
Computed:true,
Default:stringdefault.StaticString(quicksight.IdentityStoreQuicksight),
PlanModifiers:[]planmodifier.String{
stringplanmodifier.RequiresReplace(),
},
},
"namespace":schema.StringAttribute{
Required:true,
PlanModifiers:[]planmodifier.String{
stringplanmodifier.RequiresReplace(),
},
},
names.AttrTags:tftags.TagsAttribute(),
names.AttrTagsAll:tftags.TagsAttributeComputedOnly(),
},
Blocks:map[string]schema.Block{
"timeouts":timeouts.Block(ctx,timeouts.Opts{
Create:true,
Delete:true,
}),
},
}
}


func(r*resourceNamespace)Create(ctxcontext.Context,reqresource.CreateRequest,resp*resource.CreateResponse){
conn:=r.Meta().QuickSightConn(ctx)

varplanresourceNamespaceData
resp.Diagnostics.Append(req.Plan.Get(ctx,&plan)...)
ifresp.Diagnostics.HasError(){
return
}

ifplan.AWSAccountID.IsUnknown()||plan.AWSAccountID.IsNull(){
plan.AWSAccountID=types.StringValue(r.Meta().AccountID)
}
plan.ID=types.StringValue(createNamespaceID(plan.AWSAccountID.ValueString(),plan.Namespace.ValueString()))

in:=quicksight.CreateNamespaceInput{
AwsAccountId:aws.String(plan.AWSAccountID.ValueString()),
Namespace:aws.String(plan.Namespace.ValueString()),
IdentityStore:aws.String(plan.IdentityStore.ValueString()),
Tags:getTagsIn(ctx),
}

out,err:=conn.CreateNamespaceWithContext(ctx,&in)
iferr!=nil{
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.QuickSight,create.ErrActionCreating,ResNameNamespace,plan.Namespace.String(),err),
err.Error(),
)
return
}
ifout==nil{
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.QuickSight,create.ErrActionCreating,ResNameNamespace,plan.Namespace.String(),nil),
errors.New("emptyoutput").Error(),
)
return
}

createTimeout:=r.CreateTimeout(ctx,plan.Timeouts)
waitOut,err:=waitNamespaceCreated(ctx,conn,plan.ID.ValueString(),createTimeout)
iferr!=nil{
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.QuickSight,create.ErrActionWaitingForCreation,ResNameNamespace,plan.Namespace.String(),err),
err.Error(),
)
return
}
plan.ARN=flex.StringToFramework(ctx,waitOut.Arn)
plan.CapacityRegion=flex.StringToFramework(ctx,waitOut.CapacityRegion)
plan.CreationStatus=flex.StringToFramework(ctx,waitOut.CreationStatus)
plan.IdentityStore=flex.StringToFramework(ctx,waitOut.IdentityStore)

resp.Diagnostics.Append(resp.State.Set(ctx,plan)...)
}


func(r*resourceNamespace)Read(ctxcontext.Context,reqresource.ReadRequest,resp*resource.ReadResponse){
conn:=r.Meta().QuickSightConn(ctx)

varstateresourceNamespaceData
resp.Diagnostics.Append(req.State.Get(ctx,&state)...)
ifresp.Diagnostics.HasError(){
return
}

out,err:=FindNamespaceByID(ctx,conn,state.ID.ValueString())
iftfresource.NotFound(err){
resp.State.RemoveResource(ctx)
return
}
iferr!=nil{
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.QuickSight,create.ErrActionSetting,ResNameNamespace,state.ID.String(),nil),
err.Error(),
)
return
}

state.ARN=flex.StringToFramework(ctx,out.Arn)
state.CapacityRegion=flex.StringToFramework(ctx,out.CapacityRegion)
state.CreationStatus=flex.StringToFramework(ctx,out.CreationStatus)
state.IdentityStore=flex.StringToFramework(ctx,out.IdentityStore)

//Tosupportimport,parsetheIDforthecomponentkeysandset
//individualvaluesinstate
awsAccountID,namespace,err:=ParseNamespaceID(state.ID.ValueString())
iferr!=nil{
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.QuickSight,create.ErrActionSetting,ResNameNamespace,state.ID.String(),nil),
err.Error(),
)
return
}
state.AWSAccountID=flex.StringValueToFramework(ctx,awsAccountID)
state.Namespace=flex.StringValueToFramework(ctx,namespace)

resp.Diagnostics.Append(resp.State.Set(ctx,&state)...)
}


func(r*resourceNamespace)Update(ctxcontext.Context,reqresource.UpdateRequest,resp*resource.UpdateResponse){
//ThereisnoupdateAPI,andtagupdatesarehandledviaa"before"
//interceptor.Copytheplannedtagattributestostatetoensure
//updatesarecaptured.
varplanresourceNamespaceData
resp.Diagnostics.Append(req.Plan.Get(ctx,&plan)...)
ifresp.Diagnostics.HasError(){
return
}

resp.Diagnostics.Append(resp.State.Set(ctx,&plan)...)
}


func(r*resourceNamespace)Delete(ctxcontext.Context,reqresource.DeleteRequest,resp*resource.DeleteResponse){
conn:=r.Meta().QuickSightConn(ctx)

varstateresourceNamespaceData
resp.Diagnostics.Append(req.State.Get(ctx,&state)...)
ifresp.Diagnostics.HasError(){
return
}

_,err:=conn.DeleteNamespaceWithContext(ctx,&quicksight.DeleteNamespaceInput{
AwsAccountId:aws.String(state.AWSAccountID.ValueString()),
Namespace:aws.String(state.Namespace.ValueString()),
})
iferr!=nil{
iftfawserr.ErrCodeEquals(err,quicksight.ErrCodeResourceNotFoundException){
return
}
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.QuickSight,create.ErrActionDeleting,ResNameNamespace,state.ID.String(),nil),
err.Error(),
)
}

deleteTimeout:=r.DeleteTimeout(ctx,state.Timeouts)
_,err=waitNamespaceDeleted(ctx,conn,state.ID.ValueString(),deleteTimeout)
iferr!=nil{
resp.Diagnostics.AddError(
create.ProblemStandardMessage(names.QuickSight,create.ErrActionWaitingForDeletion,ResNameNamespace,state.ID.String(),err),
err.Error(),
)
return
}
}


func(r*resourceNamespace)ImportState(ctxcontext.Context,reqresource.ImportStateRequest,resp*resource.ImportStateResponse){
resource.ImportStatePassthroughID(ctx,path.Root("id"),req,resp)
}


func(r*resourceNamespace)ModifyPlan(ctxcontext.Context,reqresource.ModifyPlanRequest,resp*resource.ModifyPlanResponse){
r.SetTagsAll(ctx,req,resp)
}


funcFindNamespaceByID(ctxcontext.Context,conn*quicksight.QuickSight,idstring)(*quicksight.NamespaceInfoV2,error){
awsAccountID,namespace,err:=ParseNamespaceID(id)
iferr!=nil{
returnnil,err
}

in:=&quicksight.DescribeNamespaceInput{
AwsAccountId:aws.String(awsAccountID),
Namespace:aws.String(namespace),
}

out,err:=conn.DescribeNamespaceWithContext(ctx,in)
iferr!=nil{
iftfawserr.ErrCodeEquals(err,quicksight.ErrCodeResourceNotFoundException){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:in,
}
}

returnnil,err
}

ifout==nil||out.Namespace==nil{
returnnil,tfresource.NewEmptyResultError(in)
}

returnout.Namespace,nil
}


funcParseNamespaceID(idstring)(string,string,error){
parts:=strings.SplitN(id,",",3)
iflen(parts)!=2||parts[0]==""||parts[1]==""{
return"","",fmt.Errorf("unexpectedformatofID(%s),expectedAWS_ACCOUNT_ID,NAMESPACE",id)
}
returnparts[0],parts[1],nil
}


funccreateNamespaceID(awsAccountID,namespacestring)string{
returnfmt.Sprintf("%s,%s",awsAccountID,namespace)
}

typeresourceNamespaceDatastruct{
ARNtypes.String`tfsdk:"arn"`
AWSAccountIDtypes.String`tfsdk:"aws_account_id"`
CapacityRegiontypes.String`tfsdk:"capacity_region"`
CreationStatustypes.String`tfsdk:"creation_status"`
IDtypes.String`tfsdk:"id"`
IdentityStoretypes.String`tfsdk:"identity_store"`
Namespacetypes.String`tfsdk:"namespace"`
Tagstypes.Map`tfsdk:"tags"`
TagsAlltypes.Map`tfsdk:"tags_all"`
Timeoutstimeouts.Value`tfsdk:"timeouts"`
}


funcwaitNamespaceCreated(ctxcontext.Context,conn*quicksight.QuickSight,idstring,timeouttime.Duration)(*quicksight.NamespaceInfoV2,error){
stateConf:=&retry.StateChangeConf{
Pending:[]string{
quicksight.NamespaceStatusCreating,
},
Target:[]string{
quicksight.NamespaceStatusCreated,
},
Refresh:statusNamespace(ctx,conn,id),
Timeout:timeout,
MinTimeout:10*time.Second,
}

outputRaw,err:=stateConf.WaitForStateContext(ctx)
ifoutput,ok:=outputRaw.(*quicksight.NamespaceInfoV2);ok{
returnoutput,err
}

returnnil,err
}


funcwaitNamespaceDeleted(ctxcontext.Context,conn*quicksight.QuickSight,idstring,timeouttime.Duration)(*quicksight.NamespaceInfoV2,error){
stateConf:=&retry.StateChangeConf{
Pending:[]string{
quicksight.NamespaceStatusDeleting,
},
Target:[]string{},
Refresh:statusNamespace(ctx,conn,id),
Timeout:timeout,
MinTimeout:10*time.Second,
}

outputRaw,err:=stateConf.WaitForStateContext(ctx)
ifoutput,ok:=outputRaw.(*quicksight.NamespaceInfoV2);ok{
returnoutput,err
}

returnnil,err
}


funcstatusNamespace(ctxcontext.Context,conn*quicksight.QuickSight,idstring)retry.StateRefresh
func{
return
func()(interface{},string,error){
output,err:=FindNamespaceByID(ctx,conn,id)

iftfresource.NotFound(err){
returnnil,"",nil
}

iferr!=nil{
returnnil,"",err
}

returnoutput,aws.StringValue(output.CreationStatus),nil
}
}
