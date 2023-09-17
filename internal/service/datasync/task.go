// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package datasyncimport (
	"context"
	"fmt"
	"log"
	"time"	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)// @SDKResource("aws_datasync_task", name="Task")
// @Tags(identifierAttribute="id")
func ResourceTask() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourceTaskCreate,
adWithoutTimeout:resourceTaskRead,
dateWithoutTimeout: resourceTaskUpdate,
leteWithoutTimeout: resourceTaskDelete,
porter: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,meouts: &schema.ResourceTimeout{
Create: schema.DefaultTimeout(5 * time.Minute),
Scma: map[string]*schema.Schema{
"arn": {
Type:schema.TypeString,
Computed: true,
},
"cloudwatch_log_group_arn": {
Type:schema.TypeString,
Optional:true,
ValidateFunc: verify.ValidARN,
},
"destination_location_arn": {
Type:schema.TypeString,
Required:true,
ForceNew:true,
ValidateFunc: verify.ValidARN,
},
"excludes": {
Type:schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"filter_type": {
	Type:schema.TypeString,
	Optional:true,
	ValidateFunc: validation.StringInSlice(datasync.FilterType_Values(), false),
},
"value": {
	Type:schema.TypeString,
	Optional: true,
},
},
},
},
"includes": {
Type:schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"filter_type": {
	Type:schema.TypeString,
	Optional:true,
	ValidateFunc: validation.StringInSlice(datasync.FilterType_Values(), false),
},
"value": {
	Type:schema.TypeString,
	Optional: true,
},
},
},
},
"name": {
Type:schema.TypeString,
Optional: true,
},
"options": {
Type:schema.TypeList,
Optional:true,
MaxItems:1,
DiffSuppressFunc: verify.SuppressMissingOptionalConfigurationBlock,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"atime": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.AtimeBestEffort,
	ValidateFunc: validation.StringInSlice(datasync.Atime_Values(), false),
},
"bytes_per_second": {
	Type:schema.TypeInt,
	Optional:true,
	Default: -1,
	ValidateFunc: validation.IntAtLeast(-1),
},
"gid": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.GidIntValue,
	ValidateFunc: validation.StringInSlice(datasync.Gid_Values(), false),
},
"log_level": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.LogLevelOff,
	ValidateFunc: validation.StringInSlice(datasync.LogLevel_Values(), false),
},
"mtime": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.MtimePreserve,
	ValidateFunc: validation.StringInSlice(datasync.Mtime_Values(), false),
},
"object_tags": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.ObjectTagsPreserve,
	ValidateFunc: validation.StringInSlice(datasync.ObjectTags_Values(), false),
},
"overwrite_mode": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.OverwriteModeAlways,
	ValidateFunc: validation.StringInSlice(datasync.OverwriteMode_Values(), false),
},
"posix_permissions": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.PosixPermissionsPreserve,
	ValidateFunc: validation.StringInSlice(datasync.PosixPermissions_Values(), false),
},
"preserve_deleted_files": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.PreserveDeletedFilesPreserve,
	ValidateFunc: validation.StringInSlice(datasync.PreserveDeletedFiles_Values(), false),
},
"preserve_devices": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.PreserveDevicesNone,
	ValidateFunc: validation.StringInSlice(datasync.PreserveDevices_Values(), false),
},
"security_descriptor_copy_flags": {
	Type:schema.TypeString,
	Optional:true,
	Computed:true,
	ValidateFunc: validation.StringInSlice(datasync.SmbSecurityDescriptorCopyFlags_Values(), false),
},
"task_queueing": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.TaskQueueingEnabled,
	ValidateFunc: validation.StringInSlice(datasync.TaskQueueing_Values(), false),
},
"transfer_mode": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.TransferModeChanged,
	ValidateFunc: validation.StringInSlice(datasync.TransferMode_Values(), false),
},
"uid": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.UidIntValue,
	ValidateFunc: validation.StringInSlice(datasync.Uid_Values(), false),
},
"verify_mode": {
	Type:schema.TypeString,
	Optional:true,
	Default: datasync.VerifyModePointInTimeConsistent,
	ValidateFunc: validation.StringInSlice(datasync.VerifyMode_Values(), false),
},
},
},
},
"schedule": {
Type:schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"schedule_expression": {
	Type:schema.TypeString,
	Required: true,
	ValidateFunc: validation.All(
lidation.StringLenBetween(1, 256),
lidation.StringMatch(regexache.MustCompile(`^[0-9A-Za-z_\s #()*+,/?^|-]*$`),
"Schedule expressions must have the following syntax: rate(<number>\\\\s?(minutes?|hours?|days?)), cron(<cron_expression>) or at(yyyy-MM-dd'T'HH:mm:ss)."),
	),
},
},
},
},
"source_location_arn": {
Type:schema.TypeString,
Required:true,
ForceNew:true,
ValidateFunc: verify.ValidARN,
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll: tftags.TagsSchemaComputed(),
CuomizeDiff: verify.SetTagsDiff,
	}
}func resourceTaskCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	input := &datasync.CreateTaskInput{
stinationLocationArn: aws.String(d.Get("destination_location_arn").(string)),
tions:expandOptions(d.Get("options").([]interface{})),
urceLocationArn: aws.String(d.Get("source_location_arn").(string)),
gs: getTagsIn(ctx),
	}	if v, ok := d.GetOk("cloudwatch_log_group_arn"); ok {
put.CloudWatchLogGroupArn = aws.String(v.(string))
	}	if v, ok := d.GetOk("excludes"); ok {
put.Excludes = expandFilterRules(v.([]interface{}))
	}	if v, ok := d.GetOk("includes"); ok {
put.Includes = expandFilterRules(v.([]interface{}))
	}	if v, ok := d.GetOk("name"); ok {
put.Name = aws.String(v.(string))
	}	if v, ok := d.GetOk("schedule"); ok {
put.Schedule = expandTaskSchedule(v.([]interface{}))
	}	output, err := conn.CreateTaskWithContext(ctx, input)	if err != nil {
turn sdkdiag.AppendErrorf(diags, "creating DataSync Task: %s", err)
	}	d.SetId(aws.StringValue(output.TaskArn))	if _, err := waitTaskAvailable(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
turn sdkdiag.AppendErrorf(diags, "waiting for DataSync Task (%s) creation: %s", d.Id(), err)
	}	return append(diags, resourceTaskRead(ctx, d, meta)...)
}func resourceTaskRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	output, err := FindTaskByARN(ctx, conn, d.Id())	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] DataSync Task (%s) not found, removing from state", d.Id())
SetId("")
turn diags
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading DataSync Task (%s): %s", d.Id(), err)
	}	d.Set("arn", output.TaskArn)
	d.Set("cloudwatch_log_group_arn", output.CloudWatchLogGroupArn)
	d.Set("destination_location_arn", output.DestinationLocationArn)
	if err := d.Set("excludes", flattenFilterRules(output.Excludes)); err != nil {
turn sdkdiag.AppendErrorf(diags, "setting excludes: %s", err)
	}
	if err := d.Set("includes", flattenFilterRules(output.Includes)); err != nil {
turn sdkdiag.AppendErrorf(diags, "setting includes: %s", err)
	}
	d.Set("name", output.Name)
	if err := d.Set("options", flattenOptions(output.Options)); err != nil {
turn sdkdiag.AppendErrorf(diags, "setting options: %s", err)
	}
	if err := d.Set("schedule", flattenTaskSchedule(output.Schedule)); err != nil {
turn sdkdiag.AppendErrorf(diags, "setting schedule: %s", err)
	}
	d.Set("source_location_arn", output.SourceLocationArn)	return diags
}func resourceTaskUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	if d.HasChangesExcept("tags", "tags_all") {
put := &datasync.UpdateTaskInput{
TaskArn: aws.String(d.Id()),
f.HasChanges("cloudwatch_log_group_arn") {
input.CloudWatchLogGroupArn = aws.String(d.Get("cloudwatch_log_group_arn").(string))
f.HasChanges("excludes") {
input.Excludes = expandFilterRules(d.Get("excludes").([]interface{}))
f.HasChanges("includes") {
input.Includes = expandFilterRules(d.Get("includes").([]interface{}))
f.HasChanges("name") {
input.Name = aws.String(d.Get("name").(string))
f.HasChanges("options") {
input.Options = expandOptions(d.Get("options").([]interface{}))
f.HasChanges("schedule") {
input.Schedule = expandTaskSchedule(d.Get("schedule").([]interface{}))
,rr := conn.UpdateTaskWithContext(ctx, input)if e != nil {
return sdkdiag.AppendErrorf(diags, "updating DataSync Task (%s): %s", d.Id(), err)	}	return append(diags, resourceTaskRead(ctx, d, meta)...)
}func resourceTaskDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	log.Printf("[DEBUG] Deleting DataSync Task: %s", d.Id())
	_, err := conn.DeleteTaskWithContext(ctx, &datasync.DeleteTaskInput{
skArn: aws.String(d.Id()),
	})	if tfawserr.ErrMessageContains(err, datasync.ErrCodeInvalidRequestException, "not found") {
turn diags
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "deleting DataSync Task (%s): %s", d.Id(), err)
	}	return diags
}func FindTaskByARN(ctx context.Context, conn *datasync.DataSync, arn string) (*datasync.DescribeTaskOutput, error) {
	input := &datasync.DescribeTaskInput{
skArn: aws.String(arn),
	}	output, err := conn.DescribeTaskWithContext(ctx, input)	if tfawserr.ErrMessageContains(err, datasync.ErrCodeInvalidRequestException, "not found") {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if err != nil {
turn nil, err
	}	if output == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output, nil
}func statusTask(ctx context.Context, conn *datasync.DataSync, arn string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
tput, err := FindTaskByARN(ctx, conn, arn)iffresource.NotFound(err) {
return nil, "", nil
frr != nil {
return nil, "", err
ern output, aws.StringValue(output.Status), nil
	}
}func waitTaskAvailable(ctx context.Context, conn *datasync.DataSync, arn string, timeout time.Duration) (*datasync.DescribeTaskOutput, error) {
	stateConf := &retry.StateChangeConf{
nding: []string{datasync.TaskStatusCreating, datasync.TaskStatusUnavailable},
rget:  []string{datasync.TaskStatusAvailable, datasync.TaskStatusRunning},
fresh: statusTask(ctx, conn, arn),
meout: timeout,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)	if output, ok := outputRaw.(*datasync.DescribeTaskOutput); ok {
 errorCode, errorDetail := aws.StringValue(output.ErrorCode), aws.StringValue(output.ErrorDetail); errorCode != "" && errorDetail != "" {
tfresource.SetLastError(err, fmt.Errorf("%s: %s", errorCode, errorDetail))
ern output, err
	}	return nil, err
}func flattenOptions(options *datasync.Options) []interface{} {
	if options == nil {
turn []interface{}{}
	}	m := map[string]interface{}{
time":Atime),
ytes_per_second": aws.Int64Value(options.BytesPerSecond),
id": aws.StringValue(options.Gid),
og_level":aws.StringValue(options.LogLevel),
time":Mtime),
bject_tags":  aws.StringValue(options.ObjectTags),
verwrite_mode":OverwriteMode),
osix_permissions":aws.StringValue(options.PosixPermissions),
reserve_deleted_files":aws.StringValue(options.PreserveDeletedFiles),
reserve_devices": aws.StringValue(options.PreserveDevices),
ecurity_descriptor_copy_flags": aws.StringValue(options.SecurityDescriptorCopyFlags),
ask_queueing":aws.StringValue(options.TaskQueueing),
ransfer_mode":aws.StringValue(options.TransferMode),
id": aws.StringValue(options.Uid),
erify_mode":  aws.StringValue(options.VerifyMode),
	}	return []interface{}{m}
}func expandOptions(l []interface{}) *datasync.Options {
	if len(l) == 0 || l[0] == nil {
turn nil
	}	m := l[0].(map[string]interface{})	options := &datasync.Options{
ime:aws.String(m["atime"].(string)),
d:aws.String(m["gid"].(string)),
gLevel:aws.String(m["log_level"].(string)),
ime:aws.String(m["mtime"].(string)),
jectTags:  aws.String(m["object_tags"].(string)),
erwriteMode:mode"].(string)),
eserveDeletedFiles: aws.String(m["preserve_deleted_files"].(string)),
eserveDevices: aws.String(m["preserve_devices"].(string)),
sixPermissions:aws.String(m["posix_permissions"].(string)),
skQueueing:aws.String(m["task_queueing"].(string)),
ansferMode:aws.String(m["transfer_mode"].(string)),
d:aws.String(m["uid"].(string)),
rifyMode:  aws.String(m["verify_mode"].(string)),
	}	if v, ok := m["bytes_per_second"].(int); ok && v != 0 {
tions.BytesPerSecond = aws.Int64(int64(v))
	}	if v, ok := m["security_descriptor_copy_flags"].(string); ok && v != "" {
tions.SecurityDescriptorCopyFlags = aws.String(v)
	}	return options
}func expandTaskSchedule(l []interface{}) *datasync.TaskSchedule {
	if len(l) == 0 || l[0] == nil {
turn nil
	}	m := l[0].(map[string]interface{})	schedule := &datasync.TaskSchedule{
heduleExpression: aws.String(m["schedule_expression"].(string)),
	}	return schedule
}func flattenTaskSchedule(schedule *datasync.TaskSchedule) []interface{} {
	if schedule == nil {
turn []interface{}{}
	}	m := map[string]interface{}{
chedule_expression": aws.StringValue(schedule.ScheduleExpression),
	}	return []interface{}{m}
}func expandFilterRules(l []interface{}) []*datasync.FilterRule {
	filterRules := []*datasync.FilterRule{}	for _, mRaw := range l {
 mRaw == nil {
continue:= mRaw.(map[string]interface{})
lterRule := &datasync.FilterRule{
FilterType: aws.String(m["filter_type"].(string)),
Value: aws.String(m["value"].(string)),lterRules = append(filterRules, filterRule)
	}	return filterRules
}func flattenFilterRules(filterRules []*datasync.FilterRule) []interface{} {
	l := []interface{}{}	for _, filterRule := range filterRules {
:= map[string]interface{}{
"filter_type": aws.StringValue(filterRule.FilterType),
"value":aws.StringValue(filterRule.Value),= append(l, m)
	}	return l
}
