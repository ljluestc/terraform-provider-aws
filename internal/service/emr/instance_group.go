// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package emrimport (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)const (
	instanceGroupCreateTimeout = 30 * time.Minute
	instanceGroupUpdateTimeout = 30 * time.Minute
)// @SDKResource("aws_emr_instance_group")func ResourceInstanceGroup() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourceInstanceGroupCreate,
adWithoutTimeout:resourceInstanceGroupRead,
dateWithoutTimeout: resourceInstanceGroupUpdate,
leteWithoutTimeout: resourceInstanceGroupDelete,
porter: &schema.ResourceImporter{
StateContext:func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
idParts := strings.Split(d.Id(), "/")
if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
return nil, fmt.Errorf("Unexpected format of ID (%q), expected cluster-id/ig-id", d.Id())
}
clusterID := idParts[0]
resourceID := idParts[1]
d.Set("cluster_id", clusterID)
d.SetId(resourceID)
return []*schema.ResourceData{d}, nil
},hema: map[string]*schema.Schema{
"autoscaling_policy": {
Type:schema.TypeString,
Optional:true,
DiffSuppress
func: verify.SuppressEquivalentJSONDiffs,
Validate
func:validation.StringIsJSON,
},
"bid_price": {
Type:schema.TypeString,
Optional: true,
ForceNew: true,
},
"cluster_id": {
Type:schema.TypeString,
Required: true,
ForceNew: true,
},
"configurations_json": {
Type:schema.TypeString,
Optional:true,
Validate
func:validation.StringIsJSON,
DiffSuppress
func: verify.SuppressEquivalentJSONDiffs,
State
func:func(v interface{}) string {
json, _ := structure.NormalizeJsonString(v)
return json
},
},
"ebs_config": {
Type:schema.TypeSet,
Optional: true,
Computed: true,
ForceNew: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"iops": {
	Type:schema.TypeInt,
	Optional: true,
	ForceNew: true,
},
"size": {
	Type:schema.TypeInt,
	Required: true,
	ForceNew: true,
},
"type": {
	Type:schema.TypeString,
	Required:true,
	ForceNew:true,
	Validate
func: validEBSVolumeType(),
},
"volumes_per_instance": {
	Type:schema.TypeInt,
	Optional: true,
	ForceNew: true,
	Default:  1,
},
},
},
Set: resourceClusterEBSHashConfig,
},
"ebs_optimized": {
Type:schema.TypeBool,
Optional: true,
ForceNew: true,
},
"instance_count": {
Type:schema.TypeInt,
Optional: true,
Computed: true,
},
"instance_type": {
Type:schema.TypeString,
Required: true,
ForceNew: true,
},
"name": {
Type:schema.TypeString,
Optional: true,
ForceNew: true,
},
"running_instance_count": {
Type:schema.TypeInt,
Computed: true,
},
"status": {
Type:schema.TypeString,
Computed: true,
},	}
}func resourceInstanceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EMRConn(ctx)	instanceRole := emr.InstanceGroupTypeTask
	groupConfig := &emr.InstanceGroupConfig{
sConfiguration: readEBSConfig(d),
stanceRole:aws.String(instanceRole),
stanceType:aws.String(d.Get("instance_type").(string)),
me:aws.String(d.Get("name").(string)),
	}	if v, ok := d.GetOk("autoscaling_policy"); ok {
r autoScalingPolicy *emr.AutoScalingPolicyifrr := json.Unmarshal([]byte(v.(string)), &autoScalingPolicy); err != nil {
return sdkdiag.AppendErrorf(diags, "[DEBUG] error parsing Auto Scaling Policy %s", err)oupConfig.AutoScalingPolicy = autoScalingPolicy
	}	if v, ok := d.GetOk("configurations_json"); ok {
fo, err := structure.NormalizeJsonString(v)
 err != nil {
return sdkdiag.AppendErrorf(diags, "configurations_json contains an invalid JSON: %s", err)oupConfig.Configurations, err = expandConfigurationJSON(info)
 err != nil {
return sdkdiag.AppendErrorf(diags, "reading EMR configurations_json: %s", err)	}	if v, ok := d.GetOk("instance_count"); ok {
oupConfig.InstanceCount = aws.Int64(int64(v.(int)))
	} else {
oupConfig.InstanceCount = aws.Int64(1)
	}	groupConfig.Market = aws.String(emr.MarketTypeOnDemand)
	if v, ok := d.GetOk("bid_price"); ok {
oupConfig.BidPrice = aws.String(v.(string))
oupConfig.Market = aws.String(emr.MarketTypeSpot)
	}	params := &emr.AddInstanceGroupsInput{
stanceGroups: []*emr.InstanceGroupConfig{groupConfig},
bFlowId: aws.String(d.Get("cluster_id").(string)),
	}	resp, err := conn.AddInstanceGroupsWithContext(ctx, params)
	if err != nil {
turn sdkdiag.AppendErrorf(diags, "creating EMR Instance Group: %s", err)
	}	if resp == nil || len(resp.InstanceGroupIds) == 0 {
turn sdkdiag.AppendErrorf(diags, "creating EMR Instance Group: empty response")
	}
	d.SetId(aws.StringValue(resp.InstanceGroupIds[0]))	if err := waitForInstanceGroupStateRunning(ctx, conn, d.Get("cluster_id").(string), d.Id(), instanceGroupCreateTimeout); err != nil {
turn sdkdiag.AppendErrorf(diags, "creating EMR Instance Group (%s): waiting for completion: %s", d.Id(), err)
	}	return append(diags, resourceInstanceGroupRead(ctx, d, meta)...)
}func resourceInstanceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EMRConn(ctx)	ig, err := FetchInstanceGroup(ctx, conn, d.Get("cluster_id").(string), d.Id())	if tfresource.NotFound(err) {
g.Printf("[DEBUG] EMR Instance Group (%s) not found, removing", d.Id())
SetId("")
turn diags
	}	if err != nil {
 tfawserr.ErrMessageContains(err, emr.ErrCodeInvalidRequestException, "is not valid") {
log.Printf("[DEBUG] EMR Cluster corresponding to Instance Group (%s) not found, removing", d.Id())
d.SetId("")
return diagsturn sdkdiag.AppendErrorf(diags, "reading EMR Instance Group (%s): %s", d.Id(), err)
	}	if ig.Status != nil {
itch aws.StringValue(ig.Status.State) {
se emr.InstanceGroupStateTerminating:
fallthrough
se emr.InstanceGroupStateTerminated:
log.Printf("[DEBUG] EMR Instance Group (%s) terminated, removing", d.Id())
d.SetId("")
return diags	}	switch {
	case len(ig.Configurations) > 0:
nfigOut, err := flattenConfigurationJSON(ig.Configurations)
 err != nil {
return sdkdiag.AppendErrorf(diags, "reading EMR instance group configurations: %s", err) err := d.Set("configurations_json", configOut); err != nil {
return sdkdiag.AppendErrorf(diags, "setting EMR configurations_json for instance group (%s): %s", d.Id(), err)	default:
Set("configurations_json", "")
	}	autoscalingPolicyString, err := flattenAutoScalingPolicyDescription(ig.AutoScalingPolicy)	if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading EMR Instance Group (%s): %s", d.Id(), err)
	}	d.Set("autoscaling_policy", autoscalingPolicyString)	d.Set("bid_price", ig.BidPrice)
	if err := d.Set("ebs_config", flattenEBSConfig(ig.EbsBlockDevices)); err != nil {
turn sdkdiag.AppendErrorf(diags, "setting ebs_config: %s", err)
	}
	d.Set("ebs_optimized", ig.EbsOptimized)
	d.Set("instance_count", ig.RequestedInstanceCount)
	d.Set("instance_type", ig.InstanceType)
	d.Set("name", ig.Name)
	d.Set("running_instance_count", ig.RunningInstanceCount)	if ig.Status != nil {
Set("status", ig.Status.State)
	}	return diags
}func resourceInstanceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EMRConn(ctx)	log.Printf("[DEBUG] Modify EMR task group")
	if d.HasChanges("instance_count", "configurations_json") {
stanceGroupModifyConfig := emr.InstanceGroupModifyConfig{
InstanceGroupId: aws.String(d.Id()),
f.HasChange("instance_count") {
instanceCount := d.Get("instance_count").(int)
instanceGroupModifyConfig.InstanceCount = aws.Int64(int64(instanceCount)) d.HasChange("configurations_json") {
if v, ok := d.GetOk("configurations_json"); ok {
info, err := structure.NormalizeJsonString(v)
if err != nil {
return sdkdiag.AppendErrorf(diags, "configurations_json contains an invalid JSON: %s", err)
}
instanceGroupModifyConfig.Configurations, err = expandConfigurationJSON(info)
if err != nil {
return sdkdiag.AppendErrorf(diags, "reading EMR configurations_json: %s", err)
}
}rams := &emr.ModifyInstanceGroupsInput{
InstanceGroups: []*emr.InstanceGroupModifyConfig{
&instanceGroupModifyConfig,
},
,rr := conn.ModifyInstanceGroupsWithContext(ctx, params)
 err != nil {
return sdkdiag.AppendErrorf(diags, "modifying EMR Instance Group (%s): %s", d.Id(), err)
frr := waitForInstanceGroupStateRunning(ctx, conn, d.Get("cluster_id").(string), d.Id(), instanceGroupUpdateTimeout); err != nil {
return sdkdiag.AppendErrorf(diags, "waiting for EMR Instance Group (%s) modification: %s", d.Id(), err)	}	if d.HasChange("autoscaling_policy") {
r autoScalingPolicy *emr.AutoScalingPolicyifrr := json.Unmarshal([]byte(d.Get("autoscaling_policy").(string)), &autoScalingPolicy); err != nil {
return sdkdiag.AppendErrorf(diags, "parsing EMR Auto Scaling Policy JSON for update: %s", err)
uutoScalingPolicy := &emr.PutAutoScalingPolicyInput{
ClusterId:aws.String(d.Get("cluster_id").(string)),
AutoScalingPolicy: autoScalingPolicy,
InstanceGroupId:aws.String(d.Id()),
f, err := conn.PutAutoScalingPolicyWithContext(ctx, putAutoScalingPolicy); err != nil {
return sdkdiag.AppendErrorf(diags, "updating autoscaling policy for instance group %q: %s", d.Id(), err)	}	return append(diags, resourceInstanceGroupRead(ctx, d, meta)...)
}func resourceInstanceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EMRConn(ctx)	log.Printf("[WARN] AWS EMR Instance Group does not support DELETE; resizing cluster to zero before removing from state")
	params := &emr.ModifyInstanceGroupsInput{
stanceGroups: []*emr.InstanceGroupModifyConfig{
{
InstanceGroupId: aws.String(d.Id()),
InstanceCount:aws.Int64(0),
},	}	if _, err := conn.ModifyInstanceGroupsWithContext(ctx, params); err != nil {
turn sdkdiag.AppendErrorf(diags, "draining EMR Instance Group (%s): %s", d.Id(), err)
	}
	return diags
}func instanceGroupStateRefresh(ctx context.Context, conn *emr.EMR, clusterID, groupID string) retry.StateRefresh
func {
	returnfunc() (interface{}, string, error) {
, err := FetchInstanceGroup(ctx, conn, clusterID, groupID)
 err != nil {
return nil, "Not Found", err
fg.Status == nil || ig.Status.State == nil {
log.Printf("[WARN] ERM Instance Group found, but without state")
return nil, "Undefined", fmt.Errorf("Undefined EMR Cluster Instance Group state")
ern ig, *ig.Status.State, nil
	}
}func FetchInstanceGroup(ctx context.Context, conn *emr.EMR, clusterID, groupID string) (*emr.InstanceGroup, error) {
	input := &emr.ListInstanceGroupsInput{ClusterId: aws.String(clusterID)}	var groups []*emr.InstanceGroup
	err := conn.ListInstanceGroupsPagesWithContext(ctx, input,func(page *emr.ListInstanceGroupsOutput, lastPage bool) bool {
oups = append(groups, page.InstanceGroups...)rern !lastPage
	})	if err != nil {
turn nil, fmt.Errorf("unable to retrieve EMR Cluster (%q): %s", clusterID, err)
	}	if len(groups) == 0 {
turn nil, fmt.Errorf("no instance groups found for EMR Cluster (%s)", clusterID)
	}	var ig *emr.InstanceGroup
	for _, group := range groups {
 groupID == aws.StringValue(group.Id) {
ig = group
break	}	if ig == nil {
turn nil, &retry.NotFoundError{}
	}	return ig, nil
}// readEBSConfig populates an emr.EbsConfiguration structfunc readEBSConfig(d *schema.ResourceData) *emr.EbsConfiguration {
	result := &emr.EbsConfiguration{}
	if v, ok := d.GetOk("ebs_optimized"); ok {
sult.EbsOptimized = aws.Bool(v.(bool))
	}	ebsConfigs := make([]*emr.EbsBlockDeviceConfig, 0)
	if rawConfig, ok := d.GetOk("ebs_config"); ok {
nfigList := rawConfig.(*schema.Set).List()
r _, config := range configList {
conf := config.(map[string]interface{})
ebs := &emr.EbsBlockDeviceConfig{}
volumeSpec := &emr.VolumeSpecification{
SizeInGB:aws.Int64(int64(conf["size"].(int))),
VolumeType: aws.String(conf["type"].(string)),
}
if v, ok := conf["iops"].(int); ok && v != 0 {
volumeSpec.Iops = aws.Int64(int64(v))
}
if v, ok := conf["volumes_per_instance"].(int); ok && v != 0 {
ebs.VolumesPerInstance = aws.Int64(int64(v))
}
ebs.VolumeSpecification = volumeSpec
ebsConfigs = append(ebsConfigs, ebs)	}
	result.EbsBlockDeviceConfigs = ebsConfigs
	return result
}func waitForInstanceGroupStateRunning(ctx context.Context, conn *emr.EMR, clusterID string, instanceGroupID string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
nding: []string{
emr.InstanceGroupStateBootstrapping,
emr.InstanceGroupStateProvisioning,
emr.InstanceGroupStateReconfiguring,
emr.InstanceGroupStateResizing,rget:[]string{emr.InstanceGroupStateRunning},
fresh:instanceGroupStateRefresh(ctx, conn, clusterID, instanceGroupID),
meout:timeout,
lay: 10 * time.Second,
nTimeout: 3 * time.Second,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}
