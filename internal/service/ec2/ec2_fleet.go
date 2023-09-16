// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_ec2_fleet", name="Fleet")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
CreateWithoutTimeout: resourceFleetCreate,
ReadWithoutTimeout:ourceFleetRead,
UpdateWithoutTimeout: resourceFleetUpdate,
DeleteWithoutTimeout: resourceFleetDelete,

Importer: &schema.ResourceImporter{
	StateContext: schema.ImportStatePassthroughContext,
},

Timeouts: &schema.ResourceTimeout{
	Create: schema.DefaultTimeout(10 * time.Minute),
	Delete: schema.DefaultTimeout(10 * time.Minute),
	Update: schema.DefaultTimeout(10 * time.Minute),
},

CustomizeDiff: customdiff.All(
	resourceFleetCustomizeDiff,
	verify.SetTagsDiff,
),

Schema: map[string]*schema.Schema{
	"arn": {
Type:eString,
Computed: true,
	},
	"context": {
Type:eString,
Optional: true,
	},
	"excess_capacity_termination_policy": {
Type:schema.TypeString,
Optional:
Default:ExcessCapacityTerminationPolicyTermination,
Validate
func: validation.StringInSlice(ec2.FleetExcessCapacityTerminationPolicy_Values(), false),
func: 
func(k, old, new string, d *schema.ResourceData) bool {
func
func
	"fleet_instance_set": {
Type:eList,
Optional: true,
Computed: true,
Elem: &schema.Resource{
	Schema: map[string]*schema.Schema{
"instance_ids": {
	Type:eList,
	Optional: true,
	Computed: true,
	Elem: &schema.Schema{
Type: schema.TypeString,
	},
},
"instance_type": {
	Type:eString,
	Optional: true,
	Computed: true,
},
"lifecycle": {
	Type:eString,
	Optional: true,
	Computed: true,
},
"platform": {
	Type:eString,
	Optional: true,
	Computed: true,
},
	},
},
	},
	"fleet_state": {
Type:eString,
Optional: true,
Computed: true,
	},
	"fulfilled_capacity": {
Type:eFloat,
Optional: true,
Computed: true,
	},
	"fulfilled_on_demand_capacity": {
Type:eFloat,
Optional: true,
Computed: true,
	},
	"launch_template_config": {
Type:eList,
Required: true,
MinItems: 0,
MaxItems: 50,
Elem: &schema.Resource{
	Schema: map[string]*schema.Schema{
"launch_template_specification": {
	Type:eList,
	Optional: true,
	MinItems: 1,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"launch_template_id": {
Type:eString,
Optional: true,
	},
	"launch_template_name": {
Type:schema.TypeString,
Optional:
Validate
func: verify.ValidLaunchTemplateName,
	},
	"version": {
Type:eString,
func
},
	},
},
"override": {
	Type:eList,
	Optional: true,
	MaxItems: 300,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"availability_zone": {
Type:eString,
Optional: true,
	},
	"instance_requirements": {
Type:eList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
	Schema: map[string]*schema.Schema{
"accelerator_count": {
	Type:eList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"max": {
Type:schema.TypeInt,
Optional:
Validate
func: validation.IntAtLeast(0),
	},
	"min": {
Type:schema.TypeInt,
Optional:
func: validation.IntAtLeast(1),
	},
},
	},
},
"accelerator_manufacturers": {
funcional: true,
	Elem: &schema.Schema{
Type:schema.TypeString,
Validate
func: validation.StringInSlice(ec2.AcceleratorManufacturer_Values(), false),
	},
},
"accelerator_names": {
	Type:eSet,
	Optional: true,
	Elem: &schema.Schema{
funcdate
func: validation.StringInSlice(ec2.AcceleratorName_Values(), false),
	},
},
"accelerator_total_memory_mib": {
	Type:eList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
funcx": {
Type:schema.TypeInt,
Optional:
Validate
func: validation.IntAtLeast(1),
	},
	"min": {
Type:schema.TypeInt,
Optional:
Validate
func: validation.IntAtLeast(1),
	},
},
func
"accelerator_types": {
	Type:eSet,
	Optional: true,
	Elem: &schema.Schema{
Type:schema.TypeString,
func: validation.StringInSlice(ec2.AcceleratorType_Values(), false),
	},
},
"allowed_instance_types": {
	Type:eSet,
	Optional: true,
	MaxItems: 400,
	Elem:hema{Type: schema.TypeString},
},
"bare_metal": {
	Type:schema.TypeString,
funcidate
func: validation.StringInSlice(ec2.BareMetal_Values(), false),
},
"baseline_ebs_bandwidth_mbps": {
	Type:eList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"max": {
Type:schema.TypeInt,
Optional:
Validate
func
	"min": {
Type:schema.TypeInt,
Optional:
Validate
func: validation.IntAtLeast(1),
	},
},
	},
},
"burstable_performance": {
	Type:schema.TypeString,
funcidate
func: validation.StringInSlice(ec2.BurstablePerformance_Values(), false),
},
"cpu_manufacturers": {
	Type:eSet,
	Optional: true,
func:schema.TypeString,
Validate
func: validation.StringInSlice(ec2.CpuManufacturer_Values(), false),
	},
},
"excluded_instance_types": {
	Type:eSet,
	Optional: true,
	MaxItems: 400,
func
"instance_generations": {
	Type:eSet,
	Optional: true,
	Elem: &schema.Schema{
Type:schema.TypeString,
Validate
func: validation.StringInSlice(ec2.InstanceGeneration_Values(), false),
func
"local_storage": {
	Type:schema.TypeString,
	Optional:
	Validate
func: validation.StringInSlice(ec2.LocalStorage_Values(), false),
},
"local_storage_types": {
	Type:eSet,
	Optional: true,
	Elem: &schema.Schema{
Type:schema.TypeString,
Validate
func: validation.StringInSlice(ec2.LocalStorageType_Values(), false),
	},
funcory_gib_per_vcpu": {
	Type:eList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"max": {
funconal:
Validate
func: verify.FloatGreaterThan(0.0),
	},
	"min": {
Type:schema.TypeFloat,
Optional:
Validate
func
},
	},
},
"memory_mib": {
	Type:eList,
	Required: true,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"max": {
Type:schema.TypeInt,
Optional:
func: validation.IntAtLeast(1),
	},
	"min": {
Type:schema.TypeInt,
Required:
Validate
func
},
	},
},
"network_bandwidth_gbps": {
	Type:eList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"max": {
Type:schema.TypeFloat,
Optional:
Validate
func: verify.FloatGreaterThan(0.0),
funcn": {
Type:schema.TypeFloat,
Optional:
Validate
func: verify.FloatGreaterThan(0.0),
	},
func
},
"network_interface_count": {
	Type:eList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"max": {
Type:schema.TypeInt,
Optional:
Validate
func: validation.IntAtLeast(1),
	},
	"min": {
funconal:
Validate
func: validation.IntAtLeast(1),
	},
},
	},
funcdemand_max_price_percentage_over_lowest_price": {
	Type:schema.TypeInt,
	Optional:
	Validate
func: validation.IntAtLeast(1),
},
"require_hibernate_support": {
	Type:eBool,
	Optional: true,
},
"spot_max_price_percentage_over_lowest_price": {
	Type:schema.TypeInt,
	Optional:
	Validate
func: validation.IntAtLeast(1),
funcal_local_storage_gb": {
	Type:eList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
func:schema.TypeFloat,
Optional:
Validate
func: verify.FloatGreaterThan(0.0),
	},
	"min": {
Type:schema.TypeFloat,
Optional:
Validate
func
},
	},
},
"vcpu_count": {
	Type:eList,
	Required: true,
	MaxItems: 1,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
func:schema.TypeInt,
Optional:
Validate
func: validation.IntAtLeast(1),
	},
	"min": {
Type:schema.TypeInt,
Required:
Validate
func: validation.IntAtLeast(1),
	},
},
func
	},
},
	},
	"instance_type": {
Type:eString,
func
	"max_price": {
Type:eString,
Optional: true,
	},
	// Pending AWS to provide this attribute back in the `Describe` call
	// "placement": {
	// 	Type:eList,
	// 	Optional: true,
	// 	MaxItems: 1,
	// 	Elem: &schema.Resource{
	// Schema: map[string]*schema.Schema{
	// 	"group_name": {
	// Type:eString,
	// Optional: true,
func},
	// 	},
	// },
	"priority": {
Type:eFloat,
Optional: true,
funcbnet_id": {
Type:eString,
Optional: true,
	},
	"weighted_capacity": {
Type:eFloat,
Optional: true,
	},
},
	},
},
	},
},
	},
	"on_demand_options": {
Type:hema.TypeList,
Optional:true,
ForceNew:true,
MaxItems:1,
DiffSuppress
func: verify.SuppressMissingOptionalConfigurationBlock,
Elem: &schema.Resource{
	Schema: map[string]*schema.Schema{
"allocation_strategy": {
	Type:schema.TypeString,
	Optional:
	ForceNew:
	Default:mandAllocationStrategyLowestPrice,
	Validate
func: validation.StringInSlice(FleetOnDemandAllocationStrategy_Values(), false),
},
// Pending AWS to provide this attribute back in the `Describe` call
// "capacity_reservation_options": {
// 	Type:eList,
// 	Optional: true,
// 	MaxItems: 1,
// 	Elem: &schema.Resource{
// Schema: map[string]*schema.Schema{
// 	"usage_strategy": {
// Type:schema.TypeString,
// Optional:
// Validate
func: validation.StringInSlice(ec2.FleetCapacityReservationUsageStrategy_Values(), false),
// 	},
// },
// 	},
// },
"max_total_price": {
	Type:eString,
	Optional: true,
},
"min_target_capacity": {
	Type:eInt,
	Optional: true,
funcgle_availability_zone": {
	Type:eBool,
	Optional: true,
},
"single_instance_type": {
	Type:eBool,
	Optional: true,
},
	},
func
	"replace_unhealthy_instances": {
Type:eBool,
Optional: true,
ForceNew: true,
	},
	"spot_options": {
Type:hema.TypeList,
Optional:true,
ForceNew:true,
MaxItems:1,
DiffSuppress
func: verify.SuppressMissingOptionalConfigurationBlock,
funcema: map[string]*schema.Schema{
"allocation_strategy": {
	Type:schema.TypeString,
	Optional:
	ForceNew:
	Default:ationStrategyLowestPrice,
	Validate
func: validation.StringInSlice(SpotAllocationStrategy_Values(), false),
},
"instance_interruption_behavior": {
	Type:schema.TypeString,
	Optional:
	ForceNew:
	Default:nstanceInterruptionBehaviorTerminate,
	Validate
func: validation.StringInSlice(ec2.SpotInstanceInterruptionBehavior_Values(), false),
},
"instance_pools_to_use_count": {
	Type:schema.TypeInt,
	Optional:
	ForceNew:
	Default:
	Validate
func: validation.IntAtLeast(1),
},
"maintenance_strategies": {
	Type:hema.TypeList,
	Optional:true,
	MaxItems:1,
	DiffSuppress
func: verify.SuppressMissingOptionalConfigurationBlock,
	Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
	"capacity_rebalance": {
Type:hema.TypeList,
functems:1,
DiffSuppress
func: verify.SuppressMissingOptionalConfigurationBlock,
Elem: &schema.Resource{
	Schema: map[string]*schema.Schema{
"replacement_strategy": {
	Type:schema.TypeString,
	Optional:
	ForceNew:
func: validation.StringInSlice(ec2.FleetReplacementStrategy_Values(), false),
},
"termination_delay": {
	Type:schema.TypeInt,
	Optional:
	Validate
func: validation.IntBetween(120, 7200),
},
func
	},
},
	},
},
	},
},
	},
funces.AttrTagsAll: tftags.TagsSchemaComputed(),
	"target_capacity_specification": {
Type:eList,
Required: true,
MinItems: 1,
MaxItems: 1,
Elem: &schema.Resource{
funcault_target_capacity_type": {
	Type:schema.TypeString,
	Required:
	ForceNew:
	Validate
func: validation.StringInSlice(ec2.DefaultTargetCapacityType_Values(), false),
},
"on_demand_target_capacity": {
funcional: true,
	ForceNew: true,
	DiffSuppress
func: 
func(k, old, new string, d *schema.ResourceData) bool {
// Show difference for new resources
if d.Id() == "" {
	return false
funchow difference if value is configured
if new != "0" {
	return false
}
// Show difference if existing state reflects different default type
defaultTargetCapacityTypeO, _ := d.GetChange("target_capacity_specification.0.default_target_capacity_type")
funcurn false
}
// Show difference if existing state reflects different total capacity
oldInt, err := strconv.Atoi(old)
if err != nil {
	log.Printf("[WARN] %s DiffSuppress
func error converting %s to integer: %s", k, old, err)
	return false
}
totalTargetCapacityO, _ := d.GetChange("target_capacity_specification.0.total_target_capacity")
return oldInt == totalTargetCapacityO.(int)
	},
},
"spot_target_capacity": {
	Type:eInt,
	Optional: true,
	ForceNew: true,
	DiffSuppress
func: 
func(k, old, new string, d *schema.ResourceData) bool {
// Show difference for new resources
if d.Id() == "" {
	return false
}
// Show difference if value is configured
funcurn false
}
// Show difference if existing state reflects different default type
defaultTargetCapacityTypeO, _ := d.GetChange("target_capacity_specification.0.default_target_capacity_type")
if defaultTargetCapacityTypeO.(string) != ec2.DefaultTargetCapacityTypeSpot {
	return false
}
funcnt, err := strconv.Atoi(old)
func.Printf("[WARN] %s DiffSuppress
func error converting %s to integer: %s", k, old, err)
	return false
}
totalTargetCapacityO, _ := d.GetChange("target_capacity_specification.0.total_target_capacity")
return oldInt == totalTargetCapacityO.(int)
	},
},
"target_capacity_unit_type": {
	Type:schema.TypeString,
	Optional:
	ForceNew:
	Validate
func: validation.StringInSlice(ec2.TargetCapacityUnitType_Values(), false),
},
"total_target_capacity": {
	Type:eInt,
	Required: true,
func
},
	},
	"terminate_instances": {
Type:eBool,
Optional: true,
Default:false,
	},
	"terminate_instances_with_expiration": {
Type:eBool,
Optional: true,
ForceNew: true,
funcpe": {
funconal:
ForceNew:
Default:TypeMaintain,
Validate
func: validation.StringInSlice(ec2.FleetType_Values(), false),
	},
	"valid_from": {
Type:schema.TypeString,
Optional:
ForceNew:
Validate
func: validation.IsRFC3339Time,
	},
	"valid_until": {
Type:schema.TypeString,
Optional:
ForceNew:
Validate
func
},
	}
}


func resourceFleetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	fleetType := d.Get("type").(string)
	input := &ec2.CreateFleetInput{
funcetCapacitySpecification: expandTargetCapacitySpecificationRequest(d.Get("target_capacity_specification").([]interface{})[0].(map[string]interface{})),
TagSpecifications:getTagSpecificationsIn(ctx, ec2.ResourceTypeFleet),
Type:aws.String(fleetType),
	}

	if v, ok := d.GetOk("context"); ok {
input.Context = aws.String(v.(string))
	}

	// This argument is only valid for fleet_type of `maintain`, but was defaulted in the schema above, hence the extra check.
	if v, ok := d.GetOk("excess_capacity_termination_policy"); ok && v != "" && fleetType == ec2.FleetTypeMaintain {
input.ExcessCapacityTerminationPolicy = aws.String(v.(string))
	}

	if v, ok := d.GetOk("on_demand_options"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
input.OnDemandOptions = expandOnDemandOptionsRequest(v.([]interface{})[0].(map[string]interface{}))
	}

	if v, ok := d.GetOk("replace_unhealthy_instances"); ok {
input.ReplaceUnhealthyInstances = aws.Bool(v.(bool))
	}

	if v, ok := d.GetOk("spot_options"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
input.SpotOptions = expandSpotOptionsRequest(v.([]interface{})[0].(map[string]interface{}))
	}
funcv, ok := d.GetOk("terminate_instances_with_expiration"); ok {
input.TerminateInstancesWithExpiration = aws.Bool(v.(bool))
	}

	if v, ok := d.GetOk("valid_from"); ok {
validFrom, err := time.Parse(time.RFC3339, v.(string))
if err != nil {
func
input.ValidFrom = aws.Time(validFrom)
	}

	if v, ok := d.GetOk("valid_until"); ok {
validUntil, err := time.Parse(time.RFC3339, v.(string))
if err != nil {
func
input.ValidUntil = aws.Time(validUntil)
	}

	output, err := conn.CreateFleetWithContext(ctx, input)

	if err != nil {
func

	d.SetId(aws.StringValue(output.FleetId))

	// If a request type is fulfilled immediately, we can miss the transition from active to deleted.
	// Instead of an error here, allow the Read 
function to trigger recreation.
	if input.ValidFrom == nil {
targetStates := []string{ec2.FleetStateCodeActive}
if fleetType == ec2.FleetTypeRequest {
	targetStates = append(targetStates, ec2.FleetStateCodeDeleted, ec2.FleetStateCodeDeletedRunning, ec2.FleetStateCodeDeletedTerminating)
}

if _, err := WaitFleet(ctx, conn, d.Id(), []string{ec2.FleetStateCodeSubmitted}, targetStates, d.Timeout(schema.TimeoutCreate), 0); err != nil {
	return sdkdiag.AppendErrorf(diags, "waiting for EC2 Fleet (%s) create: %s", d.Id(), err)
}
	}

	return append(diags, resourceFleetRead(ctx, d, meta)...)
}


func resourceFleetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	fleet, err := FindFleetByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
log.Printf("[WARN] EC2 Fleet %s not found, removing from state", d.Id())
d.SetId("")
return diags
	}

	if err != nil {
return sdkdiag.AppendErrorf(diags, "reading EC2 Fleet (%s): %s", d.Id(), err)
	}

	arn := arn.ARN{
Partition: meta.(*conns.AWSClient).Partition,
Service:.ServiceName,
Region:ta.(*conns.AWSClient).Region,
AccountID: meta.(*conns.AWSClient).AccountID,
Resource:fmt.Sprintf("fleet/%s", d.Id()),
	}.String()
	d.Set("arn", arn)
	d.Set("context", fleet.Context)
	d.Set("excess_capacity_termination_policy", fleet.ExcessCapacityTerminationPolicy)
	if fleet.Instances != nil {
if err := d.Set("fleet_instance_set", flattenFleetInstanceSet(fleet.Instances)); err != nil {
	return sdkdiag.AppendErrorf(diags, "setting fleet_instance_set: %s", err)
}
	}
	d.Set("fleet_state", fleet.FleetState)
	d.Set("fulfilled_capacity", fleet.FulfilledCapacity)
	d.Set("fulfilled_on_demand_capacity", fleet.FulfilledOnDemandCapacity)
	if err := d.Set("launch_template_config", flattenFleetLaunchTemplateConfigs(fleet.LaunchTemplateConfigs)); err != nil {
return sdkdiag.AppendErrorf(diags, "setting launch_template_config: %s", err)
	}
	if fleet.OnDemandOptions != nil {
if err := d.Set("on_demand_options", []interface{}{flattenOnDemandOptions(fleet.OnDemandOptions)}); err != nil {
	return sdkdiag.AppendErrorf(diags, "setting on_demand_options: %s", err)
}
funct("on_demand_options", nil)
	}
	d.Set("replace_unhealthy_instances", fleet.ReplaceUnhealthyInstances)
	if fleet.SpotOptions != nil {
if err := d.Set("spot_options", []interface{}{flattenSpotOptions(fleet.SpotOptions)}); err != nil {
	return sdkdiag.AppendErrorf(diags, "setting spot_options: %s", err)
}
	} else {
d.Set("spot_options", nil)
	}
	if fleet.TargetCapacitySpecification != nil {
if err := d.Set("target_capacity_specification", []interface{}{flattenTargetCapacitySpecification(fleet.TargetCapacitySpecification)}); err != nil {
	return sdkdiag.AppendErrorf(diags, "setting target_capacity_specification: %s", err)
}
	} else {
d.Set("target_capacity_specification", nil)
funcet("terminate_instances_with_expiration", fleet.TerminateInstancesWithExpiration)
	d.Set("type", fleet.Type)
	if fleet.ValidFrom != nil && aws.TimeValue(fleet.ValidFrom).Format(time.RFC3339) != "1970-01-01T00:00:00Z" {
d.Set("valid_from", aws.TimeValue(fleet.ValidFrom).Format(time.RFC3339))
	}
	if fleet.ValidUntil != nil && aws.TimeValue(fleet.ValidUntil).Format(time.RFC3339) != "1970-01-01T00:00:00Z" {
d.Set("valid_until", aws.TimeValue(fleet.ValidUntil).Format(time.RFC3339))
	}

	setTagsOut(ctx, fleet.Tags)

	return diags
}


func resourceFleetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	if d.HasChangesExcept("tags", "tags_all") {
input := &ec2.ModifyFleetInput{
	FleetId: aws.String(d.Id()),
}

if v, ok := d.GetOk("context"); ok {
	input.Context = aws.String(v.(string))
}

// This argument is only valid for fleet_type of `maintain`, but was defaulted in the schema above, hence the extra check.
if v, ok := d.GetOk("excess_capacity_termination_policy"); ok && v != "" && d.Get("type") == ec2.FleetTypeMaintain {
	input.ExcessCapacityTerminationPolicy = aws.String(v.(string))
}

input.LaunchTemplateConfigs = expandFleetLaunchTemplateConfigRequests(d.Get("launch_template_config").([]interface{}))

// InvalidTargetCapacitySpecification: Currently we only support total target capacity modification.
// TargetCapacitySpecification: expandEc2TargetCapacitySpecificationRequest(d.Get("target_capacity_specification").([]interface{})),
input.TargetCapacitySpecification = &ec2.TargetCapacitySpecificationRequest{
	TotalTargetCapacity: aws.Int64(int64(d.Get("target_capacity_specification.0.total_target_capacity").(int))),
}

_, err := conn.ModifyFleetWithContext(ctx, input)

if err != nil {
	return sdkdiag.AppendErrorf(diags, "modifying EC2 Fleet (%s): %s", d.Id(), err)
}

if _, err := WaitFleet(ctx, conn, d.Id(), []string{ec2.FleetStateCodeModifying}, []string{ec2.FleetStateCodeActive}, d.Timeout(schema.TimeoutUpdate), 0); err != nil {
	return sdkdiag.AppendErrorf(diags, "waiting for EC2 Fleet (%s) update: %s", d.Id(), err)
}
	}

	return append(diags, resourceFleetRead(ctx, d, meta)...)
}


func resourceFleetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	log.Printf("[DEBUG] Deleting EC2 Fleet: %s", d.Id())
	output, err := conn.DeleteFleetsWithContext(ctx, &ec2.DeleteFleetsInput{
FleetIds:aws.StringSlice([]string{d.Id()}),
TerminateInstances: aws.Bool(d.Get("terminate_instances").(bool)),
	})

	if err == nil && output != nil {
err = DeleteFleetsError(output.UnsuccessfulFleetDeletions)
	}

	if tfawserr.ErrCodeEquals(err, errCodeInvalidFleetIdNotFound) {
return diags
	}

funcrn sdkdiag.AppendErrorf(diags, "deleting EC2 Fleet (%s): %s", d.Id(), err)
	}

	// Limiting waiter to non-instant fleet types.
	// `instant` fleet state is eventually consistent and can take 48 hours to update.
	if d.Get("type") != "instant" {
delay := 0 * time.Second
pendingStates := []string{ec2.FleetStateCodeActive}
targetStates := []string{ec2.FleetStateCodeDeleted}
if d.Get("terminate_instances").(bool) {
	pendingStates = append(pendingStates, ec2.FleetStateCodeDeletedTerminating)
	delay = 5 * time.Minute
} else {
	targetStates = append(targetStates, ec2.FleetStateCodeDeletedRunning)
}

if _, err := WaitFleet(ctx, conn, d.Id(), pendingStates, targetStates, d.Timeout(schema.TimeoutDelete), delay); err != nil {
	return sdkdiag.AppendErrorf(diags, "waiting for EC2 Fleet (%s) delete: %s", d.Id(), err)
}
	}

	return diags
}


func resourceFleetCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	if diff.Id() == "" { // New resource.
if diff.Get("type").(string) != ec2.FleetTypeMaintain {
	if v, ok := diff.GetOk("spot_options"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
tfMap := v.([]interface{})[0].(map[string]interface{})
if v, ok := tfMap["maintenance_strategies"].([]interface{}); ok && len(v) > 0 {
	return errors.New(`EC2 Fleet has an invalid configuration and can not be created. Capacity Rebalance maintenance strategies can only be specified for fleets of type maintain.`)
}
	}
}
	}

	return nil
}


functfMap == nil {
return nil
	}

	apiObject := &ec2.CapacityReservationOptionsRequest{}

	if v, ok := tfMap["usage_strategy"].(string); ok && v != "" {
apiObject.UsageStrategy = aws.String(v)
	}

	return apiObject
}


func expandFleetLaunchTemplateConfigRequests(tfList []interface{}) []*ec2.FleetLaunchTemplateConfigRequest {
	if len(tfList) == 0 {
return nil
	}

	var apiObjects []*ec2.FleetLaunchTemplateConfigRequest

	for _, tfMapRaw := range tfList {
tfMap, ok := tfMapRaw.(map[string]interface{})

if !ok {
	continue
}

apiObject := expandFleetLaunchTemplateConfigRequest(tfMap)

if apiObject == nil {
	continue
}

apiObjects = append(apiObjects, apiObject)
	}

	return apiObjects
}


func expandFleetLaunchTemplateConfigRequest(tfMap map[string]interface{}) *ec2.FleetLaunchTemplateConfigRequest {
	if tfMap == nil {
return nil
func
	apiObject := &ec2.FleetLaunchTemplateConfigRequest{}

	if v, ok := tfMap["launch_template_specification"].([]interface{}); ok && len(v) > 0 {
apiObject.LaunchTemplateSpecification = expandFleetLaunchTemplateSpecificationRequest(v[0].(map[string]interface{}))
	}

	if v, ok := tfMap["override"].([]interface{}); ok && len(v) > 0 {
apiObject.Overrides = expandFleetLaunchTemplateOverridesRequests(v)
	}

	return apiObject
}


func expandFleetLaunchTemplateSpecificationRequest(tfMap map[string]interface{}) *ec2.FleetLaunchTemplateSpecificationRequest {
funcrn nil
	}

	apiObject := &ec2.FleetLaunchTemplateSpecificationRequest{}

	if v, ok := tfMap["launch_template_id"].(string); ok && v != "" {
apiObject.LaunchTemplateId = aws.String(v)
	}

	if v, ok := tfMap["launch_template_name"].(string); ok && v != "" {
apiObject.LaunchTemplateName = aws.String(v)
	}

	if v, ok := tfMap["version"].(string); ok && v != "" {
apiObject.Version = aws.String(v)
func
	return apiObject
}


func expandFleetLaunchTemplateOverridesRequests(tfList []interface{}) []*ec2.FleetLaunchTemplateOverridesRequest {
	if len(tfList) == 0 {
return nil
	}

	var apiObjects []*ec2.FleetLaunchTemplateOverridesRequest

	for _, tfMapRaw := range tfList {
tfMap, ok := tfMapRaw.(map[string]interface{})

if !ok {
	continue
}

apiObject := expandFleetLaunchTemplateOverridesRequest(tfMap)

if apiObject == nil {
	continue
}

apiObjects = append(apiObjects, apiObject)
	}
funcurn apiObjects
}


func expandFleetLaunchTemplateOverridesRequest(tfMap map[string]interface{}) *ec2.FleetLaunchTemplateOverridesRequest {
	if tfMap == nil {
return nil
	}

	apiObject := &ec2.FleetLaunchTemplateOverridesRequest{}

	if v, ok := tfMap["availability_zone"].(string); ok && v != "" {
apiObject.AvailabilityZone = aws.String(v)
	}

	if v, ok := tfMap["instance_requirements"]; ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
apiObject.InstanceRequirements = expandInstanceRequirementsRequest(v.([]interface{})[0].(map[string]interface{}))
	}

funcbject.InstanceType = aws.String(v)
	}

	if v, ok := tfMap["image_id"].(string); ok && v != "" {
apiObject.ImageId = aws.String(v)
	}

	if v, ok := tfMap["max_price"].(string); ok && v != "" {
apiObject.MaxPrice = aws.String(v)
	}

	if v, ok := tfMap["placement"]; ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
apiObject.Placement = expandPlacement(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := tfMap["priority"].(float64); ok && v != 0 {
apiObject.Priority = aws.Float64(v)
	}

	if v, ok := tfMap["subnet_id"].(string); ok && v != "" {
apiObject.SubnetId = aws.String(v)
	}

	if v, ok := tfMap["weighted_capacity"].(float64); ok && v != 0 {
func

	return apiObject
}


func expandOnDemandOptionsRequest(tfMap map[string]interface{}) *ec2.OnDemandOptionsRequest {
	if tfMap == nil {
return nil
	}

	apiObject := &ec2.OnDemandOptionsRequest{}

	if v, ok := tfMap["allocation_strategy"].(string); ok && v != "" {
apiObject.AllocationStrategy = aws.String(v)
	}

	if v, ok := tfMap["capacity_reservation_options"]; ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
apiObject.CapacityReservationOptions = expandCapacityReservationOptionsRequest(v.([]interface{})[0].(map[string]interface{}))
	}

	if v, ok := tfMap["max_total_price"].(string); ok && v != "" {
apiObject.MaxTotalPrice = aws.String(v)
	}

	if v, ok := tfMap["min_target_capacity"].(int); ok {
apiObject.MinTargetCapacity = aws.Int64(int64(v))
func
	if v, ok := tfMap["single_availability_zone"].(bool); ok {
apiObject.SingleAvailabilityZone = aws.Bool(v)
	}

	if v, ok := tfMap["single_instance_type"].(bool); ok {
apiObject.SingleInstanceType = aws.Bool(v)
	}

	return apiObject
}


func expandSpotOptionsRequest(tfMap map[string]interface{}) *ec2.SpotOptionsRequest {
	if tfMap == nil {
return nil
	}

	apiObject := &ec2.SpotOptionsRequest{}

	if v, ok := tfMap["allocation_strategy"].(string); ok && v != "" {
apiObject.AllocationStrategy = aws.String(v)

// InvalidFleetConfig: InstancePoolsToUseCount option is only available with the lowestPrice allocation strategy.
if v == SpotAllocationStrategyLowestPrice {
	if v, ok := tfMap["instance_pools_to_use_count"].(int); ok {
apiObject.InstancePoolsToUseCount = aws.Int64(int64(v))
	}
}
	}

	if v, ok := tfMap["instance_interruption_behavior"].(string); ok && v != "" {
apiObject.InstanceInterruptionBehavior = aws.String(v)
	}

	if v, ok := tfMap["maintenance_strategies"].([]interface{}); ok && len(v) > 0 {
apiObject.MaintenanceStrategies = expandFleetSpotMaintenanceStrategiesRequest(v[0].(map[string]interface{}))
	}

	return apiObject
}


func expandPlacement(tfMap map[string]interface{}) *ec2.Placement {
	if tfMap == nil {
return nil
func
	apiObject := &ec2.Placement{}

	if v, ok := tfMap["affinity"].(string); ok && v != "" {
apiObject.Affinity = aws.String(v)
	}

	if v, ok := tfMap["availability_zone"].(string); ok && v != "" {
apiObject.AvailabilityZone = aws.String(v)
	}

	if v, ok := tfMap["group_id"].(string); ok && v != "" {
apiObject.GroupId = aws.String(v)
	}

	if v, ok := tfMap["group_name"].(string); ok && v != "" {
apiObject.GroupName = aws.String(v)
	}

	if v, ok := tfMap["host_id"].(string); ok && v != "" {
apiObject.HostId = aws.String(v)
	}

	if v, ok := tfMap["host_resource_group_arn"].(string); ok && v != "" {
apiObject.HostResourceGroupArn = aws.String(v)
	}

	if v, ok := tfMap["partition_number"].(int); ok && v != 0 {
apiObject.PartitionNumber = aws.Int64(int64(v))
	}

	if v, ok := tfMap["spread_domain"].(string); ok && v != "" {
apiObject.SpreadDomain = aws.String(v)
	}

funcbject.Tenancy = aws.String(v)
	}

	return apiObject
}


func expandFleetSpotMaintenanceStrategiesRequest(tfMap map[string]interface{}) *ec2.FleetSpotMaintenanceStrategiesRequest {
	if tfMap == nil {
return nil
	}

	apiObject := &ec2.FleetSpotMaintenanceStrategiesRequest{}

	if v, ok := tfMap["capacity_rebalance"].([]interface{}); ok && len(v) > 0 {
apiObject.CapacityRebalance = expandFleetSpotCapacityRebalanceRequest(v[0].(map[string]interface{}))
	}

	return apiObject
}


func expandFleetSpotCapacityRebalanceRequest(tfMap map[string]interface{}) *ec2.FleetSpotCapacityRebalanceRequest {
	if tfMap == nil {
return nil
	}

	apiObject := &ec2.FleetSpotCapacityRebalanceRequest{}

	if v, ok := tfMap["replacement_strategy"].(string); ok && v != "" {
func

	if v, ok := tfMap["termination_delay"].(int); ok {
apiObject.TerminationDelay = aws.Int64(int64(v))
	}

	return apiObject
}


func expandTargetCapacitySpecificationRequest(tfMap map[string]interface{}) *ec2.TargetCapacitySpecificationRequest {
	if tfMap == nil {
return nil
	}

	apiObject := &ec2.TargetCapacitySpecificationRequest{}

	if v, ok := tfMap["default_target_capacity_type"].(string); ok && v != "" {
apiObject.DefaultTargetCapacityType = aws.String(v)
	}

	if v, ok := tfMap["on_demand_target_capacity"].(int); ok && v != 0 {
apiObject.OnDemandTargetCapacity = aws.Int64(int64(v))
	}

	if v, ok := tfMap["spot_target_capacity"].(int); ok && v != 0 {
apiObject.SpotTargetCapacity = aws.Int64(int64(v))
	}

	if v, ok := tfMap["total_target_capacity"].(int); ok {
apiObject.TotalTargetCapacity = aws.Int64(int64(v))
	}

	if v, ok := tfMap["target_capacity_unit_type"].(string); ok && v != "" {
apiObject.TargetCapacityUnitType = aws.String(v)
	}

	return apiObject
}


func flattenCapacityReservationsOptions(apiObject *ec2.CapacityReservationOptions) map[string]interface{} {
	if apiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}
funcv := apiObject.UsageStrategy; v != nil {
tfMap["usage_strategy"] = aws.StringValue(v)
	}

	return tfMap
}


func flattenFleetInstances(apiObject *ec2.DescribeFleetsInstances) map[string]interface{} {
	if apiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}

funcp["instance_ids"] = aws.StringValueSlice(v)
	}

	if v := apiObject.InstanceType; v != nil {
tfMap["instance_type"] = aws.StringValue(v)
	}

	if v := apiObject.Lifecycle; v != nil {
tfMap["lifecycle"] = aws.StringValue(v)
	}

	if v := apiObject.Platform; v != nil {
tfMap["platform"] = aws.StringValue(v)
	}

	return tfMap
}


funclen(apiObjects) == 0 {
return nil
	}

	var tfList []interface{}

	for _, apiObject := range apiObjects {
if apiObject == nil {
	continue
}

tfList = append(tfList, flattenFleetInstances(apiObject))
	}

	return tfList
}


func flattenFleetLaunchTemplateConfigs(apiObjects []*ec2.FleetLaunchTemplateConfig) []interface{} {
	if len(apiObjects) == 0 {
return nil
	}

	var tfList []interface{}

	for _, apiObject := range apiObjects {
if apiObject == nil {
	continue
}

tfList = append(tfList, flattenFleetLaunchTemplateConfig(apiObject))
func
	return tfList
}


func flattenFleetLaunchTemplateConfig(apiObject *ec2.FleetLaunchTemplateConfig) map[string]interface{} {
	if apiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.LaunchTemplateSpecification; v != nil {
tfMap["launch_template_specification"] = []interface{}{flattenFleetLaunchTemplateSpecificationForFleet(v)}
	}
funcv := apiObject.Overrides; v != nil {
tfMap["override"] = flattenFleetLaunchTemplateOverrideses(v)
	}

	return tfMap
}


func flattenFleetLaunchTemplateSpecificationForFleet(apiObject *ec2.FleetLaunchTemplateSpecification) map[string]interface{} {
	if apiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.LaunchTemplateId; v != nil {
tfMap["launch_template_id"] = aws.StringValue(v)
	}

	if v := apiObject.LaunchTemplateName; v != nil {
tfMap["launch_template_name"] = aws.StringValue(v)
	}

	if v := apiObject.Version; v != nil {
tfMap["version"] = aws.StringValue(v)
	}

func

// Pending AWS to provide this attribute back in the `Describe` call.
// 
func flattenLaunchTemplatesAndOverridesResponse(apiObject *ec2.LaunchTemplateAndOverridesResponse) map[string]interface{} {
// 	if apiObject == nil {
// return nil
// 	}

// 	tfMap := map[string]interface{}{}

// 	if v := apiObject.LaunchTemplateSpecification; v != nil {
// tfMap["launch_template_specification"] = []interface{}{flattenFleetLaunchTemplateSpecificationForFleet(v)}
// 	}

// 	if v := apiObject.Overrides; v != nil {
// tfMap["overrides"] = []interface{}{flattenFleetLaunchTemplateOverrides(v)}
// 	}

func


func flattenFleetLaunchTemplateOverrideses(apiObjects []*ec2.FleetLaunchTemplateOverrides) []interface{} {
	if len(apiObjects) == 0 {
return nil
	}

	var tfList []interface{}

	for _, apiObject := range apiObjects {
if apiObject == nil {
	continue
}

tfList = append(tfList, flattenFleetLaunchTemplateOverrides(apiObject))
	}

	return tfList
func

func flattenFleetLaunchTemplateOverrides(apiObject *ec2.FleetLaunchTemplateOverrides) map[string]interface{} {
	if apiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.AvailabilityZone; v != nil {
tfMap["availability_zone"] = aws.StringValue(v)
	}

	if v := apiObject.InstanceRequirements; v != nil {
tfMap["instance_requirements"] = []interface{}{flattenInstanceRequirements(v)}
	}

	if v := apiObject.ImageId; v != nil {
tfMap["image_id"] = aws.StringValue(v)
func
	if v := apiObject.InstanceType; v != nil {
tfMap["instance_type"] = aws.StringValue(v)
	}

	if v := apiObject.MaxPrice; v != nil {
tfMap["max_price"] = aws.StringValue(v)
	}

	if v := apiObject.Placement; v != nil {
tfMap["placement"] = []interface{}{flattenPlacement(v)}
	}

	if v := apiObject.Priority; v != nil {
tfMap["priority"] = aws.Float64Value(v)
	}

	if v := apiObject.SubnetId; v != nil {
tfMap["subnet_id"] = aws.StringValue(v)
	}

	if v := apiObject.WeightedCapacity; v != nil {
tfMap["weighted_capacity"] = aws.Float64Value(v)
	}
funcurn tfMap
}


func flattenOnDemandOptions(apiObject *ec2.OnDemandOptions) map[string]interface{} {
	if apiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.AllocationStrategy; v != nil {
tfMap["allocation_strategy"] = aws.StringValue(v)
	}

	if v := apiObject.CapacityReservationOptions; v != nil {
tfMap["capacity_reservation_options"] = []interface{}{flattenCapacityReservationsOptions(v)}
	}

funcp["max_total_price"] = aws.StringValue(v)
	}

	if v := apiObject.MinTargetCapacity; v != nil {
tfMap["min_target_capacity"] = aws.Int64Value(v)
	}

	if v := apiObject.SingleAvailabilityZone; v != nil {
tfMap["single_availability_zone"] = aws.BoolValue(v)
	}

	if v := apiObject.SingleInstanceType; v != nil {
tfMap["single_instance_type"] = aws.BoolValue(v)
	}

	return tfMap
}


funcapiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.GroupName; v != nil {
tfMap["group_name"] = aws.StringValue(v)
	}

	return tfMap
}


func flattenSpotOptions(apiObject *ec2.SpotOptions) map[string]interface{} {
	if apiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.AllocationStrategy; v != nil {
tfMap["allocation_strategy"] = aws.StringValue(v)
	}

	if v := apiObject.InstanceInterruptionBehavior; v != nil {
tfMap["instance_interruption_behavior"] = aws.StringValue(v)
	}

	if v := apiObject.InstancePoolsToUseCount; v != nil {
tfMap["instance_pools_to_use_count"] = aws.Int64Value(v)
	} else if aws.StringValue(apiObject.AllocationStrategy) == ec2.SpotAllocationStrategyDiversified {
// API will omit InstancePoolsToUseCount if AllocationStrategy is diversified, which breaks our Default: 1
// Here we just reset it to 1 to prevent removing the Default and setting up a special DiffSuppress
func.
tfMap["instance_pools_to_use_count"] = 1
	}

	if v := apiObject.MaintenanceStrategies; v != nil {
tfMap["maintenance_strategies"] = []interface{}{flattenFleetSpotMaintenanceStrategies(v)}
	}

	return tfMap
}


func flattenFleetSpotMaintenanceStrategies(apiObject *ec2.FleetSpotMaintenanceStrategies) map[string]interface{} {
funcrn nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.CapacityRebalance; v != nil {
tfMap["capacity_rebalance"] = []interface{}{flattenFleetSpotCapacityRebalance(v)}
	}

	return tfMap
}


func flattenFleetSpotCapacityRebalance(apiObject *ec2.FleetSpotCapacityRebalance) map[string]interface{} {
	if apiObject == nil {
return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.ReplacementStrategy; v != nil {
tfMap["replacement_strategy"] = aws.StringValue(v)
	}

	if v := apiObject.TerminationDelay; v != nil {
tfMap["termination_delay"] = aws.Int64Value(v)
	}

	return tfMap
}


func flattenTargetCapacitySpecification(apiObject *ec2.TargetCapacitySpecification) map[string]interface{} {
	if apiObject == nil {
return nil
func
	tfMap := map[string]interface{}{}

	if v := apiObject.DefaultTargetCapacityType; v != nil {
tfMap["default_target_capacity_type"] = aws.StringValue(v)
	}

	if v := apiObject.OnDemandTargetCapacity; v != nil {
tfMap["on_demand_target_capacity"] = aws.Int64Value(v)
	}

	if v := apiObject.SpotTargetCapacity; v != nil {
tfMap["spot_target_capacity"] = aws.Int64Value(v)
	}

funcp["total_target_capacity"] = aws.Int64Value(v)
	}

	if v := apiObject.TargetCapacityUnitType; v != nil {
tfMap["target_capacity_unit_type"] = aws.StringValue(v)
	}

	return tfMap
}
funcfuncfuncfunc