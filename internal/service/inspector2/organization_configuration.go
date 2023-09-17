// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package inspector2import (
	"context"
	"fmt"
	"log"
	"time"	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)// @SDKResource("aws_inspector2_organization_configuration")
func ResourceOrganizationConfiguration() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourceOrganizationConfigurationCreate,
adWithoutTimeout:resourceOrganizationConfigurationRead,
dateWithoutTimeout: resourceOrganizationConfigurationUpdate,
leteWithoutTimeout: resourceOrganizationConfigurationDelete,Tiouts: &schema.ResourceTimeout{
Create: schema.DefaultTimeout(5 * time.Minute),
Update: schema.DefaultTimeout(5 * time.Minute),
Delete: schema.DefaultTimeout(5 * time.Minute),
Scma: map[string]*schema.Schema{
"auto_enable": {
Type:schema.TypeList,
Required: true,
MaxItems: 1,
MinItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"ec2": {
	Type:schema.TypeBool,
	Required: true,
},
"ecr": {
	Type:schema.TypeBool,
	Required: true,
},
"lambda": {
	Type:schema.TypeBool,
	Optional: true,
	Default:  false,
},
},
},
},
"max_account_limit_reached": {
Type:schema.TypeBool,
Computed: true,
},	}
}const (
	ResNameOrganizationConfiguration = "Organization Configuration"
	orgConfigMutex = "f14b54d7-2b10-58c2-9c1b-c48260a4825d"
)func resourceOrganizationConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId(meta.(*conns.AWSClient).AccountID)
	return resourceOrganizationConfigurationUpdate(ctx, d, meta)
}func resourceOrganizationConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).Inspector2Client(ctx)	out, err := conn.DescribeOrganizationConfiguration(ctx, &inspector2.DescribeOrganizationConfigurationInput{})	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] Inspector2 OrganizationConfiguration (%s) not found, removing from state", d.Id())
SetId("")
turn nil
	}	if err != nil {
turn create.DiagError(names.Inspector2, create.ErrActionReading, ResNameOrganizationConfiguration, d.Id(), err)
	}	if err := d.Set("auto_enable", []interface{}{flattenAutoEnable(out.AutoEnable)}); err != nil {
turn create.DiagError(names.Inspector2, create.ErrActionSetting, ResNameOrganizationConfiguration, d.Id(), err)
	}	d.Set("max_account_limit_reached", out.MaxAccountLimitReached)	return nil
}func resourceOrganizationConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).Inspector2Client(ctx)	update := false	in := &inspector2.UpdateOrganizationConfigurationInput{}	if d.HasChanges("auto_enable") {
.AutoEnable = expandAutoEnable(d.Get("auto_enable").([]interface{})[0].(map[string]interface{}))
date = true
	}	if !update {
turn nil
	}	conns.GlobalMutexKV.Lock(orgConfigMutex)
	defer conns.GlobalMutexKV.Unlock(orgConfigMutex)	log.Printf("[DEBUG] Updating Inspector2 Organization Configuration (%s): %#v", d.Id(), in)
	_, err := conn.UpdateOrganizationConfiguration(ctx, in)
	if err != nil {
turn create.DiagError(names.Inspector2, create.ErrActionUpdating, ResNameOrganizationConfiguration, d.Id(), err)
	}	if err := waitOrganizationConfigurationUpdated(ctx, conn, d.Get("auto_enable.0.ec2").(bool), d.Get("auto_enable.0.ecr").(bool), d.Get("auto_enable.0.lambda").(bool), d.Timeout(schema.TimeoutUpdate)); err != nil {
turn create.DiagError(names.Inspector2, create.ErrActionWaitingForUpdate, ResNameOrganizationConfiguration, d.Id(), err)
	}	return resourceOrganizationConfigurationRead(ctx, d, meta)
}func resourceOrganizationConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).Inspector2Client(ctx)	conns.GlobalMutexKV.Lock(orgConfigMutex)
	defer conns.GlobalMutexKV.Unlock(orgConfigMutex)	in := &inspector2.UpdateOrganizationConfigurationInput{
toEnable: &types.AutoEnable{
Ec2:aws.Bool(false),
Ecr:aws.Bool(false),
Lambda: aws.Bool(false),	}	log.Printf("[DEBUG] Setting Inspector2 Organization Configuration (%s): %#v", d.Id(), in)
	_, err := conn.UpdateOrganizationConfiguration(ctx, in)
	if err != nil {
turn create.DiagError(names.Inspector2, create.ErrActionUpdating, ResNameOrganizationConfiguration, d.Id(), err)
	}	if err := waitOrganizationConfigurationUpdated(ctx, conn, false, false, false, d.Timeout(schema.TimeoutUpdate)); err != nil {
turn create.DiagError(names.Inspector2, create.ErrActionWaitingForUpdate, ResNameOrganizationConfiguration, d.Id(), err)
	}	return nil
}func waitOrganizationConfigurationUpdated(ctx context.Context, conn *inspector2.Client, ec2, ecr, lambda bool, timeout time.Duration) error {
	needle := fmt.Sprintf("%t:%t:%t", ec2, ecr, lambda)	all := []string{
t.Sprintf("%t:%t:%t", false, false, false),
t.Sprintf("%t:%t:%t", false, true, false),
t.Sprintf("%t:%t:%t", false, false, true),
t.Sprintf("%t:%t:%t", false, true, true),
t.Sprintf("%t:%t:%t", true, false, false),
t.Sprintf("%t:%t:%t", true, false, true),
t.Sprintf("%t:%t:%t", true, true, false),
t.Sprintf("%t:%t:%t", true, true, true),
	}	for i, v := range all {
 v == needle {
all = append(all[:i], all[i+1:]...)
break	}	stateConf := &retry.StateChangeConf{
nding: all,
rget:  []string{needle},
fresh: statusOrganizationConfiguration(ctx, conn),
meout: timeout,
tFoundChecks:20,
ntinuousTargetOccurence: 2,
nTimeout:time.Second * 5,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}func statusOrganizationConfiguration(ctx context.Context, conn *inspector2.Client) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
t, err := conn.DescribeOrganizationConfiguration(ctx, &inspector2.DescribeOrganizationConfigurationInput{})
 tfresource.NotFound(err) {
return nil, "", nil
frr != nil {
return nil, "", err
ern out, fmt.Sprintf("%t:%t:%t", aws.ToBool(out.AutoEnable.Ec2), aws.ToBool(out.AutoEnable.Ecr), aws.ToBool(out.AutoEnable.Lambda)), nil
	}
}func flattenAutoEnable(apiObject *types.AutoEnable) map[string]interface{} {
	if apiObject == nil {
turn nil
	}	m := map[string]interface{}{}	if v := apiObject.Ec2; v != nil {
"ec2"] = aws.ToBool(v)
	}	if v := apiObject.Ecr; v != nil {
"ecr"] = aws.ToBool(v)
	}	if v := apiObject.Lambda; v != nil {
"lambda"] = aws.ToBool(v)
	}	return m
}func expandAutoEnable(tfMap map[string]interface{}) *types.AutoEnable {
	if tfMap == nil {
turn nil
	}	a := &types.AutoEnable{}	if v, ok := tfMap["ec2"].(bool); ok {
Ec2 = aws.Bool(v)
	}	if v, ok := tfMap["ecr"].(bool); ok {
Ecr = aws.Bool(v)
	}	if v, ok := tfMap["lambda"].(bool); ok {
Lambda = aws.Bool(v)
	}	return a
}
