// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package evidentlyimport (
"context"
fmt"
log"
strings"
time""ithub.com/YakDriver/regexache"
github.com/aws/aws-sdk-go/aws"
github.com/aws/aws-sdk-go/service/cloudwatchevidently"
github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
github.com/hashicorp/terraform-plugin-sdk/v2/diag"
github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
github.com/hashicorp/terraform-provider-aws/internal/conns"
github.com/hashicorp/terraform-provider-aws/internal/flex"
ftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
github.com/hashicorp/terraform-provider-aws/internal/tfresource"
github.com/hashicorp/terraform-provider-aws/internal/verify"
github.com/hashicorp/terraform-provider-aws/names"
)// @SDKResource("aws_evidently_launch", name="Launch")
// @Tags(identifierAttribute="arn")
func ResourceLaunch() *schema.Resource {
eturn &schema.Resource{
CreateWithoutTimeout: resourceLaunchCreate,
ReadWithoutTimeout:resourceLaunchRead,
UpdateWithoutTimeout: resourceLaunchUpdate,
DeleteWithoutTimeout: resourceLaunchDelete,Importer: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,
},Timeouts: &schema.ResourceTimeout{
Create: schema.DefaultTimeout(2 * time.Minute),
Update: schema.DefaultTimeout(2 * time.Minute),
Delete: schema.DefaultTimeout(2 * time.Minute),
},Schema: map[string]*schema.Schema{
"arn": {
Type: schema.TypeString,
Computed: true,
},
"created_time": {
Type: schema.TypeString,
Computed: true,
},
"description": {
Type:schema.TypeString,
Optional: true,
ValidateFunc: validation.StringLenBetween(0, 160),
},
"execution": {
Type: schema.TypeList,
Computed: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
ended_time": {
Type: schema.TypeString,
Computed: true,
,
started_time": {
Type: schema.TypeString,
Computed: true,
,
},
},
},
"groups": {
Type: schema.TypeList,
Required: true,
MinItems: 1,
MaxItems: 5,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
description": {
Type:schema.TypeString,
Optional: true,
ValidateFunc: validation.StringLenBetween(0, 160),
,
feature": {
Type: schema.TypeString,
Required: true,
ValidateFunc: validation.All(
validation.StringLenBetween(1, 127),
validation.StringMatch(regexache.MustCompile(`^[0-9A-Za-z_.-]*$`), "alphanumeric and can contain hyphens, underscores, and periods"),
),
,
name": {
Type: schema.TypeString,
Required: true,
ValidateFunc: validation.All(
validation.StringLenBetween(1, 127),
validation.StringMatch(regexache.MustCompile(`^[0-9A-Za-z_.-]*$`), "alphanumeric and can contain hyphens, underscores, and periods"),
),
,
variation": {
Type: schema.TypeString,
Required: true,
ValidateFunc: validation.All(
validation.StringLenBetween(1, 127),
validation.StringMatch(regexache.MustCompile(`^[0-9A-Za-z_.-]*$`), "alphanumeric and can contain hyphens, underscores, and periods"),
),
,
},
},
},
"last_updated_time": {
Type: schema.TypeString,
Computed: true,
},
"metric_monitors": {
Type: schema.TypeList,
Optional: true,
MinItems: 0,
MaxItems: 3,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
metric_definition": {
Type: schema.TypeList,
Required: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"entity_id_key": {
Type:schema.TypeString,
Required: true,
ValidateFunc: validation.StringLenBetween(1, 256),
},
"event_pattern": {
Type: schema.TypeString,
Optional: true,
ValidateFunc: validation.All(
alidation.StringLenBetween(0, 1024),
alidation.StringIsJSON,
),
DiffSuppressFunc: verify.SuppressEquivalentJSONDiffs,
StateFunc: func(v interface{}) string {
son, _ := structure.NormalizeJsonString(v)
eturn json
},
},
"name": {
Type:schema.TypeString,
Required: true,
ValidateFunc: validation.StringLenBetween(1, 255),
},
"unit_label": {
Type:schema.TypeString,
Optional: true,
ValidateFunc: validation.StringLenBetween(1, 256),
},
"value_key": {
Type:schema.TypeString,
Required: true,
ValidateFunc: validation.StringLenBetween(1, 256),
},
},
},
,
},
},
},
"name": {
Type: schema.TypeString,
Required: true,
ForceNew: true,
ValidateFunc: validation.All(
validation.StringLenBetween(1, 127),
validation.StringMatch(regexache.MustCompile(`^[0-9A-Za-z_.-]*$`), "alphanumeric and can contain hyphens, underscores, and periods"),
),
},
"project": {
Type: schema.TypeString,
Required: true,
ForceNew: true,
ValidateFunc: validation.All(
validation.StringLenBetween(0, 2048),
validation.StringMatch(regexache.MustCompile(`(^[0-9A-Za-z_.-]*$)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[0-9A-Za-z_.-]*)`), "name or arn of the project"),
),
DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
// case 1: User-defined string (old) is a name and is the suffix of API-returned string (new). Check non-empty old in resoure creation scenario
// case 2: after setting API-returned string.  User-defined string (new) is suffix of API-returned string (old)
return (strings.HasSuffix(new, old) && old != "") || strings.HasSuffix(old, new)
},
},
"randomization_salt": {
Type:schema.TypeString,
Optional: true,
ValidateFunc: validation.StringLenBetween(0, 127),
// Default: set to the launch name if not specified
DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
return old == d.Get("name").(string) && new == ""
},
},
"scheduled_splits_config": {
Type: schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
steps": {
Type: schema.TypeList,
Required: true,
MinItems: 1,
MaxItems: 6,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"group_weights": {
Type: schema.TypeMap,
Required: true,
ValidateDiagFunc: validation.AllDiag(
alidation.MapKeyLenBetween(1, 127),
alidation.MapKeyMatch(regexache.MustCompile(`^[0-9A-Za-z_.-]*$`), "alphanumeric and can contain hyphens, underscores, and periods"),
),
Elem: &schema.Schema{
ype:schema.TypeInt,
alidateFunc: validation.IntBetween(0, 100000),
},
},
"segment_overrides": {
Type: schema.TypeList,
Optional: true,
MinItems: 0,
MaxItems: 6,
Elem: &schema.Resource{
chema: map[string]*schema.Schema{
"evaluation_order": {
Type: schema.TypeInt,
Required: true,
},
"segment": {
Type: schema.TypeString,
Required: true,
ValidateFunc: validation.All(
validation.StringLenBetween(0, 2048),
validation.StringMatch(regexache.MustCompile(`(^[0-9A-Za-z_.-]*$)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:segment/[0-9A-Za-z._-]*)`), "name or arn of the segment"),
),
DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
// case 1: User-defined string (old) is a name and is the suffix of API-returned string (new). Check non-empty old in resoure creation scenario
// case 2: after setting API-returned string.  User-defined string (new) is suffix of API-returned string (old)
return (strings.HasSuffix(new, old) && old != "") || strings.HasSuffix(old, new)
},
},
"weights": {
Type: schema.TypeMap,
Required: true,
ValidateDiagFunc: validation.AllDiag(
validation.MapKeyLenBetween(1, 127),
validation.MapKeyMatch(regexache.MustCompile(`^[0-9A-Za-z_.-]*$`), "alphanumeric and can contain hyphens, underscores, and periods"),
),
Elem: &schema.Schema{
Type:schema.TypeInt,
ValidateFunc: validation.IntBetween(0, 100000),
},
},
,
},
},
"start_time": {
Type:schema.TypeString,
Required: true,
ValidateFunc: verify.ValidUTCTimestamp,
},
},
},
,
},
},
},
"status": {
Type: schema.TypeString,
Computed: true,
},
"status_reason": {
Type: schema.TypeString,
Computed: true,
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll: tftags.TagsSchemaComputed(),
"type": {
Type: schema.TypeString,
Computed: true,
},
},
CustomizeDiff: verify.SetTagsDiff,}func resourceLaunchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
onn := meta.(*conns.AWSClient).EvidentlyConn(ctx)nme := d.Get("name").(string)
roject := d.Get("project").(string)
nput := &cloudwatchevidently.CreateLaunchInput{
Name:aws.String(name),
Project: aws.String(project),
Groups:  expandGroups(d.Get("groups").([]interface{})),
Tags:getTagsIn(ctx),
i v, ok := d.GetOk("description"); ok {
input.Description = aws.String(v.(string))
i v, ok := d.GetOk("metric_monitors"); ok {
input.MetricMonitors = expandMetricMonitors(v.([]interface{}))
i v, ok := d.GetOk("randomization_salt"); ok {
input.RandomizationSalt = aws.String(v.(string))
i v, ok := d.GetOk("scheduled_splits_config"); ok {
input.ScheduledSplitsConfig = expandScheduledSplitsConfig(v.([]interface{}))
otput, err := conn.CreateLaunchWithContext(ctx, input)iferr != nil {
return diag.Errorf("creating CloudWatch Evidently Launch (%s) for Project (%s): %s", name, project, err)
/ the GetLaunch API call uses the Launch name and Project ARN
/ concat Launch name and Project Name or ARN to be used in Read for imports
.SetId(fmt.Sprintf("%s:%s", aws.StringValue(output.Launch.Name), aws.StringValue(output.Launch.Project)))i _, err := waitLaunchCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
return diag.Errorf("waiting for CloudWatch Evidently Launch (%s) for Project (%s) creation: %s", name, project, err)
rturn resourceLaunchRead(ctx, d, meta)
}func resourceLaunchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
onn := meta.(*conns.AWSClient).EvidentlyConn(ctx)lunchName, projectNameOrARN, err := LaunchParseID(d.Id())iferr != nil {
return diag.FromErr(err)
lunch, err := FindLaunchWithProjectNameorARN(ctx, conn, launchName, projectNameOrARN)if!d.IsNewResource() && tfresource.NotFound(err) {
log.Printf("[WARN] CloudWatch Evidently Launch (%s) not found, removing from state", d.Id())
d.SetId("")
return nil
i err != nil {
return diag.Errorf("reading CloudWatch Evidently Launch (%s) for Project (%s): %s", launchName, projectNameOrARN, err)
i err := d.Set("execution", flattenExecution(launch.Execution)); err != nil {
return diag.Errorf("setting execution: %s", err)
i err := d.Set("groups", flattenGroups(launch.Groups)); err != nil {
return diag.Errorf("setting groups: %s", err)
i err := d.Set("metric_monitors", flattenMetricMonitors(launch.MetricMonitors)); err != nil {
return diag.Errorf("setting metric_monitors: %s", err)
i err := d.Set("scheduled_splits_config", flattenScheduledSplitsDefinition(launch.ScheduledSplitsDefinition)); err != nil {
return diag.Errorf("setting scheduled_splits_config: %s", err)
dSet("arn", launch.Arn)
.Set("created_time", aws.TimeValue(launch.CreatedTime).Format(time.RFC3339))
.Set("description", launch.Description)
.Set("last_updated_time", aws.TimeValue(launch.LastUpdatedTime).Format(time.RFC3339))
.Set("name", launch.Name)
.Set("project", launch.Project)
.Set("randomization_salt", launch.RandomizationSalt)
.Set("status", launch.Status)
.Set("status_reason", launch.StatusReason)
.Set("type", launch.Type)stTagsOut(ctx, launch.Tags)reurn nil
}func resourceLaunchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
onn := meta.(*conns.AWSClient).EvidentlyConn(ctx)i d.HasChanges("description", "groups", "metric_monitors", "randomization_salt", "scheduled_splits_config") {
name := d.Get("name").(string)
project := d.Get("project").(string)input := &cloudwatchevidently.UpdateLaunchInput{
Description:  aws.String(d.Get("description").(string)),
Groups:expandGroups(d.Get("groups").([]interface{})),
Launch:aws.String(name),
Project:  aws.String(project),
MetricMonitors:expandMetricMonitors(d.Get("metric_monitors").([]interface{})),
RandomizationSalt: aws.String(d.Get("randomization_salt").(string)),
ScheduledSplitsConfig: expandScheduledSplitsConfig(d.Get("scheduled_splits_config").([]interface{})),
}_, err := conn.UpdateLaunchWithContext(ctx, input)if err != nil {
return diag.Errorf("updating CloudWatch Evidently Launch (%s) for Project (%s): %s", name, project, err)
}if _, err := waitLaunchUpdated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
return diag.Errorf("waiting for CloudWatch Evidently Launch (%s) for Project (%s) update: %s", name, project, err)
}
rturn resourceLaunchRead(ctx, d, meta)
}func resourceLaunchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
onn := meta.(*conns.AWSClient).EvidentlyConn(ctx)nme := d.Get("name").(string)
roject := d.Get("project").(string)lg.Printf("[DEBUG] Deleting CloudWatch Evidently Launch: %s", d.Id())
, err := conn.DeleteLaunchWithContext(ctx, &cloudwatchevidently.DeleteLaunchInput{
Launch:  aws.String(name),
Project: aws.String(project),
)i tfawserr.ErrCodeEquals(err, cloudwatchevidently.ErrCodeResourceNotFoundException) {
return nil
i err != nil {
return diag.Errorf("deleting CloudWatch Evidently Launch (%s) for Project (%s): %s", name, project, err)
i _, err := waitLaunchDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
return diag.Errorf("waiting for CloudWatch Evidently Launch (%s) for Project (%s) deletion: %s", name, project, err)
rturn nil
}func LaunchParseID(id string) (string, string, error) {
aunchName, projectNameOrARN, _ := strings.Cut(id, ":")i launchName == "" || projectNameOrARN == "" {
return "", "", fmt.Errorf("unexpected format of ID (%s), expected launchName:projectNameOrARN", id)
rturn launchName, projectNameOrARN, nil
}func expandGroups(tfMaps []interface{}) []*cloudwatchevidently.LaunchGroupConfig {
piObjects := make([]*cloudwatchevidently.LaunchGroupConfig, 0, len(tfMaps))fr _, tfMap := range tfMaps {
apiObjects = append(apiObjects, expandGroup(tfMap.(map[string]interface{})))
rturn apiObjects
}func expandGroup(tfMap map[string]interface{}) *cloudwatchevidently.LaunchGroupConfig {
piObject := &cloudwatchevidently.LaunchGroupConfig{
Feature:aws.String(tfMap["feature"].(string)),
Name:  aws.String(tfMap["name"].(string)),
Variation: aws.String(tfMap["variation"].(string)),
i v, ok := tfMap["description"]; ok {
apiObject.Description = aws.String(v.(string))
rturn apiObject
}func expandMetricMonitors(tfMaps []interface{}) []*cloudwatchevidently.MetricMonitorConfig {
piObjects := make([]*cloudwatchevidently.MetricMonitorConfig, 0, len(tfMaps))fr _, tfMap := range tfMaps {
apiObjects = append(apiObjects, expandMetricMonitor(tfMap.(map[string]interface{})))
rturn apiObjects
}func expandMetricMonitor(tfMap map[string]interface{}) *cloudwatchevidently.MetricMonitorConfig {
piObject := &cloudwatchevidently.MetricMonitorConfig{
MetricDefinition: expandMetricDefinition(tfMap["metric_definition"].([]interface{})),
rturn apiObject
}func expandMetricDefinition(tfList []interface{}) *cloudwatchevidently.MetricDefinitionConfig {
f len(tfList) == 0 || tfList[0] == nil {
return nil
tMap := tfList[0].(map[string]interface{})apObject := &cloudwatchevidently.MetricDefinitionConfig{
EntityIdKey: aws.String(tfMap["entity_id_key"].(string)),
Name:aws.String(tfMap["name"].(string)),
ValueKey:aws.String(tfMap["value_key"].(string)),
i v, ok := tfMap["event_pattern"]; ok && v != "" {
apiObject.EventPattern = aws.String(v.(string))
i v, ok := tfMap["unit_label"]; ok && v != "" {
apiObject.UnitLabel = aws.String(v.(string))
rturn apiObject
}func expandScheduledSplitsConfig(tfList []interface{}) *cloudwatchevidently.ScheduledSplitsLaunchConfig {
f len(tfList) == 0 || tfList[0] == nil {
return nil
tMap := tfList[0].(map[string]interface{})apObject := &cloudwatchevidently.ScheduledSplitsLaunchConfig{
Steps: expandSteps(tfMap["steps"].([]interface{})),
rturn apiObject
}func expandSteps(tfMaps []interface{}) []*cloudwatchevidently.ScheduledSplitConfig {
piObjects := make([]*cloudwatchevidently.ScheduledSplitConfig, 0, len(tfMaps))fr _, tfMap := range tfMaps {
apiObjects = append(apiObjects, expandStep(tfMap.(map[string]interface{})))
rturn apiObjects
}func expandStep(tfMap map[string]interface{}) *cloudwatchevidently.ScheduledSplitConfig {
, _ := time.Parse(time.RFC3339, tfMap["start_time"].(string))
tartTime := aws.Time(t)aiObject := &cloudwatchevidently.ScheduledSplitConfig{
GroupWeights: flex.ExpandInt64Map(tfMap["group_weights"].(map[string]interface{})),
SegmentOverrides: expandSegmentOverrides(tfMap["segment_overrides"].([]interface{})),
StartTime:startTime,
rturn apiObject
}func expandSegmentOverrides(tfMaps []interface{}) []*cloudwatchevidently.SegmentOverride {
piObjects := make([]*cloudwatchevidently.SegmentOverride, 0, len(tfMaps))fr _, tfMap := range tfMaps {
apiObjects = append(apiObjects, expandSegmentOverride(tfMap.(map[string]interface{})))
rturn apiObjects
}func expandSegmentOverride(tfMap map[string]interface{}) *cloudwatchevidently.SegmentOverride {
piObject := &cloudwatchevidently.SegmentOverride{
EvaluationOrder: aws.Int64(int64(tfMap["evaluation_order"].(int))),
Segment:aws.String(tfMap["segment"].(string)),
Weights:flex.ExpandInt64Map(tfMap["weights"].(map[string]interface{})),
rturn apiObject
}func flattenExecution(apiObjects *cloudwatchevidently.LaunchExecution) []interface{} {
f apiObjects == nil {
return nil
vlues := map[string]interface{}{}ifapiObjects.EndedTime != nil {
values["ended_time"] = aws.TimeValue(apiObjects.EndedTime).Format(time.RFC3339)
i apiObjects.StartedTime != nil {
values["started_time"] = aws.TimeValue(apiObjects.StartedTime).Format(time.RFC3339)
rturn []interface{}{values}
}func flattenGroups(apiObjects []*cloudwatchevidently.LaunchGroup) []interface{} {
f len(apiObjects) == 0 {
return nil
vr tfList []interface{}fo _, apiObject := range apiObjects {
if apiObject == nil {
continue
}tfList = append(tfList, flattenGroup(apiObject))
rturn tfList
}func flattenGroup(apiObject *cloudwatchevidently.LaunchGroup) map[string]interface{} {
f apiObject == nil {
return nil
tMap := map[string]interface{}{
"name": aws.StringValue(apiObject.Name),
fr feature, variation := range apiObject.FeatureVariations {
tfMap["feature"] = feature
tfMap["variation"] = aws.StringValue(variation)
i v := apiObject.Description; v != nil {
tfMap["description"] = aws.StringValue(v)
rturn tfMap
}func flattenMetricMonitors(apiObjects []*cloudwatchevidently.MetricMonitor) []interface{} {
f len(apiObjects) == 0 {
return nil
vr tfList []interface{}fo _, apiObject := range apiObjects {
if apiObject == nil {
continue
}tfList = append(tfList, flattenMetricMonitor(apiObject))
rturn tfList
}func flattenMetricMonitor(apiObject *cloudwatchevidently.MetricMonitor) map[string]interface{} {
f apiObject == nil {
return nil
tMap := map[string]interface{}{
"metric_definition": flattenMetricMonitorDefinition(apiObject.MetricDefinition),
rturn tfMap
}func flattenMetricMonitorDefinition(apiObject *cloudwatchevidently.MetricDefinition) []interface{} {
f apiObject == nil {
return nil
tMap := map[string]interface{}{
"entity_id_key": aws.StringValue(apiObject.EntityIdKey),
"name": aws.StringValue(apiObject.Name),
"value_key": aws.StringValue(apiObject.ValueKey),
i v := apiObject.EventPattern; v != nil {
tfMap["event_pattern"] = aws.StringValue(v)
i v := apiObject.UnitLabel; v != nil {
tfMap["unit_label"] = aws.StringValue(v)
rturn []interface{}{tfMap}
}func flattenScheduledSplitsDefinition(apiObject *cloudwatchevidently.ScheduledSplitsLaunchDefinition) []interface{} {
f apiObject == nil {
return nil
tMap := map[string]interface{}{
"steps": flattenSteps(apiObject.Steps),
rturn []interface{}{tfMap}
}func flattenSteps(apiObjects []*cloudwatchevidently.ScheduledSplit) []interface{} {
f len(apiObjects) == 0 {
return nil
vr tfList []interface{}fo _, apiObject := range apiObjects {
if apiObject == nil {
continue
}tfList = append(tfList, flattenStep(apiObject))
rturn tfList
}func flattenStep(apiObject *cloudwatchevidently.ScheduledSplit) map[string]interface{} {
f apiObject == nil {
return nil
tMap := map[string]interface{}{
"group_weights": aws.Int64ValueMap(apiObject.GroupWeights),
"start_time":aws.TimeValue(apiObject.StartTime).Format(time.RFC3339),
i v := apiObject.SegmentOverrides; v != nil {
tfMap["segment_overrides"] = flattenSegmentOverrides(v)
rturn tfMap
}func flattenSegmentOverrides(apiObjects []*cloudwatchevidently.SegmentOverride) []interface{} {
f len(apiObjects) == 0 {
return nil
vr tfList []interface{}fo _, apiObject := range apiObjects {
if apiObject == nil {
continue
}tfList = append(tfList, flattenSegmentOverride(apiObject))
rturn tfList
}func flattenSegmentOverride(apiObject *cloudwatchevidently.SegmentOverride) map[string]interface{} {
f apiObject == nil {
return nil
tMap := map[string]interface{}{
"evaluation_order": aws.Int64Value(apiObject.EvaluationOrder),
"segment": aws.StringValue(apiObject.Segment),
"weights": aws.Int64ValueMap(apiObject.Weights),
rturn tfMap
}
