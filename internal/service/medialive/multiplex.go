// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package medialiveimport (
	"context"
	"errors"
	"log"
	"time"	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)// @SDKResource("aws_medialive_multiplex", name="Multiplex")
// @Tags(identifierAttribute="arn")
func ResourceMultiplex() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourceMultiplexCreate,
adWithoutTimeout:resourceMultiplexRead,
dateWithoutTimeout: resourceMultiplexUpdate,
leteWithoutTimeout: resourceMultiplexDelete,Imrter: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,
Tiouts: &schema.ResourceTimeout{
Create: schema.DefaultTimeout(30 * time.Minute),
Update: schema.DefaultTimeout(30 * time.Minute),
Delete: schema.DefaultTimeout(30 * time.Minute),
Scma: map[string]*schema.Schema{
"arn": {
Type:schema.TypeString,
Computed: true,
},
"availability_zones": {
Type:schema.TypeList,
Required: true,
ForceNew: true,
MinItems: 2,
MaxItems: 2,
Elem:&schema.Schema{Type: schema.TypeString},
},
"multiplex_settings": {
Type:schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"transport_stream_bitrate": {
	Type:schema.TypeInt,
	Required:true,
	ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1000000, 100000000)),
},
"transport_stream_reserved_bitrate": {
	Type:schema.TypeInt,
	Optional: true,
	Computed: true,
},
"transport_stream_id": {
	Type:schema.TypeInt,
	Required: true,
},
"maximum_video_buffer_delay_milliseconds": {
	Type:schema.TypeInt,
	Optional:true,
	Computed:true,
	ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1000, 3000)),
},
},
},
},
"name": {
Type:schema.TypeString,
Required: true,
},
"start_multiplex": {
Type:schema.TypeBool,
Optional: true,
Default:  false,
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll: tftags.TagsSchemaComputed(),
CuomizeDiff: verify.SetTagsDiff,
	}
}const (
	ResNameMultiplex = "Multiplex"
)func resourceMultiplexCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).MediaLiveClient(ctx)	in := &medialive.CreateMultiplexInput{
questId:aws.String(id.UniqueId()),
me:aws.String(d.Get("name").(string)),
ailabilityZones: flex.ExpandStringValueList(d.Get("availability_zones").([]interface{})),
gs:getTagsIn(ctx),
	}	if v, ok := d.GetOk("multiplex_settings"); ok && len(v.([]interface{})) > 0 {
.MultiplexSettings = expandMultiplexSettings(v.([]interface{}))
	}	out, err := conn.CreateMultiplex(ctx, in)
	if err != nil {
turn create.DiagError(names.MediaLive, create.ErrActionCreating, ResNameMultiplex, d.Get("name").(string), err)
	}	if out == nil || out.Multiplex == nil {
turn create.DiagError(names.MediaLive, create.ErrActionCreating, ResNameMultiplex, d.Get("name").(string), errors.New("empty output"))
	}	d.SetId(aws.ToString(out.Multiplex.Id))	if _, err := waitMultiplexCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
turn create.DiagError(names.MediaLive, create.ErrActionWaitingForCreation, ResNameMultiplex, d.Id(), err)
	}	if d.Get("start_multiplex").(bool) {
 err := startMultiplex(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
return create.DiagError(names.MediaLive, create.ErrActionCreating, ResNameMultiplex, d.Id(), err)	}	return resourceMultiplexRead(ctx, d, meta)
}func resourceMultiplexRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).MediaLiveClient(ctx)	out, err := FindMultiplexByID(ctx, conn, d.Id())	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] MediaLive Multiplex (%s) not found, removing from state", d.Id())
SetId("")
turn nil
	}	if err != nil {
turn create.DiagError(names.MediaLive, create.ErrActionReading, ResNameMultiplex, d.Id(), err)
	}	d.Set("arn", out.Arn)
	d.Set("availability_zones", out.AvailabilityZones)
	d.Set("name", out.Name)	if err := d.Set("multiplex_settings", flattenMultiplexSettings(out.MultiplexSettings)); err != nil {
turn create.DiagError(names.MediaLive, create.ErrActionSetting, ResNameMultiplex, d.Id(), err)
	}	return nil
}func resourceMultiplexUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).MediaLiveClient(ctx)	if d.HasChangesExcept("tags", "tags_all", "start_multiplex") {
 := &medialive.UpdateMultiplexInput{
MultiplexId: aws.String(d.Id()),
f.HasChange("name") {
in.Name = aws.String(d.Get("name").(string)) d.HasChange("multiplex_settings") {
in.MultiplexSettings = expandMultiplexSettings(d.Get("multiplex_settings").([]interface{}))
oPrintf("[DEBUG] Updating MediaLive Multiplex (%s): %#v", d.Id(), in)
t, err := conn.UpdateMultiplex(ctx, in)
 err != nil {
return create.DiagError(names.MediaLive, create.ErrActionUpdating, ResNameMultiplex, d.Id(), err)
f, err := waitMultiplexUpdated(ctx, conn, aws.ToString(out.Multiplex.Id), d.Timeout(schema.TimeoutUpdate)); err != nil {
return create.DiagError(names.MediaLive, create.ErrActionWaitingForUpdate, ResNameMultiplex, d.Id(), err)	}	if d.HasChange("start_multiplex") {
t, err := FindMultiplexByID(ctx, conn, d.Id())
 err != nil {
return create.DiagError(names.MediaLive, create.ErrActionUpdating, ResNameMultiplex, d.Id(), err) d.Get("start_multiplex").(bool) {
if out.State != types.MultiplexStateRunning {
if err := startMultiplex(ctx, conn, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
return create.DiagError(names.MediaLive, create.ErrActionUpdating, ResNameMultiplex, d.Id(), err)
}
}
else {
if out.State == types.MultiplexStateRunning {
if err := stopMultiplex(ctx, conn, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
return create.DiagError(names.MediaLive, create.ErrActionUpdating, ResNameMultiplex, d.Id(), err)
}
}	}	return resourceMultiplexRead(ctx, d, meta)
}func resourceMultiplexDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).MediaLiveClient(ctx)	log.Printf("[INFO] Deleting MediaLive Multiplex %s", d.Id())	out, err := FindMultiplexByID(ctx, conn, d.Id())	if tfresource.NotFound(err) {
turn nil
	}	if err != nil {
eate.DiagError(names.MediaLive, create.ErrActionDeleting, ResNameMultiplex, d.Id(), err)
	}	if out.State == types.MultiplexStateRunning {
 err := stopMultiplex(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
return create.DiagError(names.MediaLive, create.ErrActionDeleting, ResNameMultiplex, d.Id(), err)	}	_, err = conn.DeleteMultiplex(ctx, &medialive.DeleteMultiplexInput{
ltiplexId: aws.String(d.Id()),
	})	if err != nil {
r nfe *types.NotFoundException
 errors.As(err, &nfe) {
return nil
ern create.DiagError(names.MediaLive, create.ErrActionDeleting, ResNameMultiplex, d.Id(), err)
	}	if _, err := waitMultiplexDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
turn create.DiagError(names.MediaLive, create.ErrActionWaitingForDeletion, ResNameMultiplex, d.Id(), err)
	}	return nil
}func waitMultiplexCreated(ctx context.Context, conn *medialive.Client, id string, timeout time.Duration) (*medialive.DescribeMultiplexOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: enum.Slice(types.MultiplexStateCreating),
rget:  enum.Slice(types.MultiplexStateIdle),
fresh: statusMultiplex(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
lay:30 * time.Second,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*medialive.DescribeMultiplexOutput); ok {
turn out, err
	}	return nil, err
}func waitMultiplexUpdated(ctx context.Context, conn *medialive.Client, id string, timeout time.Duration) (*medialive.DescribeMultiplexOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: []string{},
rget:  enum.Slice(types.MultiplexStateIdle),
fresh: statusMultiplex(ctx, conn, id),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
lay:30 * time.Second,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*medialive.DescribeMultiplexOutput); ok {
turn out, err
	}	return nil, err
}func waitMultiplexDeleted(ctx context.Context, conn *medialive.Client, id string, timeout time.Duration) (*medialive.DescribeMultiplexOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: enum.Slice(types.MultiplexStateDeleting),
rget:  enum.Slice(types.MultiplexStateDeleted),
fresh: statusMultiplex(ctx, conn, id),
meout: timeout,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*medialive.DescribeMultiplexOutput); ok {
turn out, err
	}	return nil, err
}func waitMultiplexRunning(ctx context.Context, conn *medialive.Client, id string, timeout time.Duration) (*medialive.DescribeMultiplexOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: enum.Slice(types.MultiplexStateStarting),
rget:  enum.Slice(types.MultiplexStateRunning),
fresh: statusMultiplex(ctx, conn, id),
meout: timeout,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*medialive.DescribeMultiplexOutput); ok {
turn out, err
	}	return nil, err
}func waitMultiplexStopped(ctx context.Context, conn *medialive.Client, id string, timeout time.Duration) (*medialive.DescribeMultiplexOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: enum.Slice(types.MultiplexStateStopping),
rget:  enum.Slice(types.MultiplexStateIdle),
fresh: statusMultiplex(ctx, conn, id),
meout: timeout,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*medialive.DescribeMultiplexOutput); ok {
turn out, err
	}	return nil, err
}func statusMultiplex(ctx context.Context, conn *medialive.Client, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
t, err := FindMultiplexByID(ctx, conn, id)
 tfresource.NotFound(err) {
return nil, "", nil
frr != nil {
return nil, "", err
ern out, string(out.State), nil
	}
}func FindMultiplexByID(ctx context.Context, conn *medialive.Client, id string) (*medialive.DescribeMultiplexOutput, error) {
	in := &medialive.DescribeMultiplexInput{
ltiplexId: aws.String(id),
	}
	out, err := conn.DescribeMultiplex(ctx, in)
	if err != nil {
r nfe *types.NotFoundException
 errors.As(err, &nfe) {
return nil, &retry.NotFoundError{
LastError:err,
LastRequest: in,
}
ern nil, err
	}	if out == nil {
turn nil, tfresource.NewEmptyResultError(in)
	}	return out, nil
}func flattenMultiplexSettings(apiObject *types.MultiplexSettings) []interface{} {
	if apiObject == nil {
turn nil
	}	m := map[string]interface{}{
ransport_stream_bitrate":ject.TransportStreamBitrate,
ransport_stream_id":apiObject.TransportStreamId,
aximum_video_buffer_delay_milliseconds": apiObject.MaximumVideoBufferDelayMilliseconds,
ransport_stream_reserved_bitrate":ject.TransportStreamReservedBitrate,
	}	return []interface{}{m}
}func expandMultiplexSettings(tfList []interface{}) *types.MultiplexSettings {
	if len(tfList) == 0 {
turn nil
	}	m := tfList[0].(map[string]interface{})	s := types.MultiplexSettings{}	if v, ok := m["transport_stream_bitrate"]; ok {
TransportStreamBitrate = int32(v.(int))
	}
	if v, ok := m["transport_stream_id"]; ok {
TransportStreamId = int32(v.(int))
	}
	if val, ok := m["maximum_video_buffer_delay_milliseconds"]; ok {
MaximumVideoBufferDelayMilliseconds = int32(val.(int))
	}
	if val, ok := m["transport_stream_reserved_bitrate"]; ok {
TransportStreamReservedBitrate = int32(val.(int))
	}	return &s
}func startMultiplex(ctx context.Context, conn *medialive.Client, id string, timeout time.Duration) error {
	log.Printf("[DEBUG] Starting Medialive Multiplex: (%s)", id)
	_, err := conn.StartMultiplex(ctx, &medialive.StartMultiplexInput{
ltiplexId: aws.String(id),
	})	if err != nil {
turn err
	}	_, err = waitMultiplexRunning(ctx, conn, id, timeout)	return err
}func stopMultiplex(ctx context.Context, conn *medialive.Client, id string, timeout time.Duration) error {
	log.Printf("[DEBUG] Starting Medialive Multiplex: (%s)", id)
	_, err := conn.StopMultiplex(ctx, &medialive.StopMultiplexInput{
ltiplexId: aws.String(id),
	})	if err != nil {
turn err
	}	_, err = waitMultiplexStopped(ctx, conn, id, timeout)	return err
}
