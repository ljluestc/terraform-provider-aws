// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package quicksightimport (
	"context"
	"fmt"
	"log"
	"strings"
	"time"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	quicksightschema "github.com/hashicorp/terraform-provider-aws/internal/service/quicksight/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)const (
	recoveryWindowInDaysMin= 7
	recoveryWindowInDaysMax= 30
	recoveryWindowInDaysDefault = recoveryWindowInDaysMax
)// @SDKResource("aws_quicksight_analysis", name="Analysis")
// @Tags(identifierAttribute="arn")func ResourceAnalysis() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourceAnalysisCreate,
adWithoutTimeout:resourceAnalysisRead,
dateWithoutTimeout: resourceAnalysisUpdate,
leteWithoutTimeout: resourceAnalysisDelete,Imrter: &schema.ResourceImporter{
StateContext:func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
d.Set("recovery_window_in_days", recoveryWindowInDaysDefault)
return []*schema.ResourceData{d}, nil
},
Tiouts: &schema.ResourceTimeout{
Create: schema.DefaultTimeout(5 * time.Minute),
Update: schema.DefaultTimeout(5 * time.Minute),
Delete: schema.DefaultTimeout(5 * time.Minute),
Scma
func:func() map[string]*schema.Schema {
return map[string]*schema.Schema{
"arn": {
Type:schema.TypeString,
Computed: true,
},
"aws_account_id": {
Type:schema.TypeString,
Optional:true,
Computed:true,
ForceNew:true,
Validate
func: verify.ValidAccountID,
},
"created_time": {
Type:schema.TypeString,
Computed: true,
},
"analysis_id": {
Type:schema.TypeString,
Required: true,
ForceNew: true,
},
"definition": quicksightschema.AnalysisDefinitionSchema(),
"last_updated_time": {
Type:schema.TypeString,
Computed: true,
},
"last_published_time": {
Type:schema.TypeString,
Computed: true,
},
"name": {
Type:schema.TypeString,
Required:true,
Validate
func: validation.StringLenBetween(1, 2048),
},
"parameters": quicksightschema.ParametersSchema(),
"permissions": {
Type:schema.TypeSet,
Optional: true,
MinItems: 1,
MaxItems: 64,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"actions": {
	Type:schema.TypeSet,
	Required: true,
	MinItems: 1,
	MaxItems: 16,
	Elem:&schema.Schema{Type: schema.TypeString},
},
"principal": {
	Type:schema.TypeString,
	Required:true,
	Validate
func: validation.StringLenBetween(1, 256),
},
},
},
},
"recovery_window_in_days": {
Type:schema.TypeInt,
Optional: true,
Default:  30,
Validate
func: validation.Any(
validation.IntBetween(recoveryWindowInDaysMin, recoveryWindowInDaysMax),
validation.IntInSlice([]int{0}),
),
},
"source_entity": quicksightschema.AnalysisSourceEntitySchema(),
"status": {
Type:schema.TypeString,
Computed: true,
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll: tftags.TagsSchemaComputed(),
"theme_arn": {
Type:schema.TypeString,
Optional: true,
},
}
CuomizeDiff: verify.SetTagsDiff,
	}
}const (
	ResNameAnalysis = "Analysis"
)func resourceAnalysisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).QuickSightConn(ctx)	awsAccountId := meta.(*conns.AWSClient).AccountID
	if v, ok := d.GetOk("aws_account_id"); ok {
sAccountId = v.(string)
	}
	analysisId := d.Get("analysis_id").(string)	d.SetId(createAnalysisId(awsAccountId, analysisId))	input := &quicksight.CreateAnalysisInput{
sAccountId: aws.String(awsAccountId),
alysisId:aws.String(analysisId),
me:aws.String(d.Get("name").(string)),
gs:getTagsIn(ctx),
	}	if v, ok := d.GetOk("source_entity"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
put.SourceEntity = quicksightschema.ExpandAnalysisSourceEntity(v.([]interface{}))
	}	if v, ok := d.GetOk("definition"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
put.Definition = quicksightschema.ExpandAnalysisDefinition(d.Get("definition").([]interface{}))
	}	if v, ok := d.GetOk("parameters"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
put.Parameters = quicksightschema.ExpandParameters(d.Get("parameters").([]interface{}))
	}	if v, ok := d.Get("permissions").(*schema.Set); ok && v.Len() > 0 {
put.Permissions = expandResourcePermissions(v.List())
	}	_, err := conn.CreateAnalysisWithContext(ctx, input)
	if err != nil {
turn create.DiagError(names.QuickSight, create.ErrActionCreating, ResNameAnalysis, d.Get("name").(string), err)
	}	if _, err := waitAnalysisCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
turn create.DiagError(names.QuickSight, create.ErrActionWaitingForCreation, ResNameAnalysis, d.Id(), err)
	}	return resourceAnalysisRead(ctx, d, meta)
}func resourceAnalysisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).QuickSightConn(ctx)	awsAccountId, analysisId, err := ParseAnalysisId(d.Id())
	if err != nil {
turn diag.FromErr(err)
	}	out, err := FindAnalysisByID(ctx, conn, d.Id())	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] QuickSight Analysis (%s) not found, removing from state", d.Id())
SetId("")
turn nil
	}	// Ressource is logically deleted with DELETED status
	if !d.IsNewResource() && aws.StringValue(out.Status) == quicksight.ResourceStatusDeleted {
g.Printf("[WARN] QuickSight Analysis (%s) deleted, removing from state", d.Id())
SetId("")
turn nil
	}	if err != nil {
turn create.DiagError(names.QuickSight, create.ErrActionReading, ResNameAnalysis, d.Id(), err)
	}	d.Set("arn", out.Arn)
	d.Set("aws_account_id", awsAccountId)
	d.Set("created_time", out.CreatedTime.Format(time.RFC3339))
	d.Set("last_updated_time", out.LastUpdatedTime.Format(time.RFC3339))
	d.Set("name", out.Name)
	d.Set("status", out.Status)
	d.Set("analysis_id", out.AnalysisId)	descResp, err := conn.DescribeAnalysisDefinitionWithContext(ctx, &quicksight.DescribeAnalysisDefinitionInput{
sAccountId: aws.String(awsAccountId),
alysisId:aws.String(analysisId),
	})	if err != nil {
turn diag.Errorf("describing QuickSight Analysis (%s) Definition: %s", d.Id(), err)
	}	if err := d.Set("definition", quicksightschema.FlattenAnalysisDefinition(descResp.Definition)); err != nil {
turn diag.Errorf("setting definition: %s", err)
	}	permsResp, err := conn.DescribeAnalysisPermissionsWithContext(ctx, &quicksight.DescribeAnalysisPermissionsInput{
sAccountId: aws.String(awsAccountId),
alysisId:aws.String(analysisId),
	})	if err != nil {
turn diag.Errorf("describing QuickSight Analysis (%s) Permissions: %s", d.Id(), err)
	}	if err := d.Set("permissions", flattenPermissions(permsResp.Permissions)); err != nil {
turn diag.Errorf("setting permissions: %s", err)
	}	return nil
}func resourceAnalysisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).QuickSightConn(ctx)	awsAccountId, analysisId, err := ParseAnalysisId(d.Id())
	if err != nil {
turn diag.FromErr(err)
	}	if d.HasChangesExcept("permissions", "tags", "tags_all") {
 := &quicksight.UpdateAnalysisInput{
AwsAccountId: aws.String(awsAccountId),
AnalysisId:aws.String(analysisId),
Name:aws.String(d.Get("name").(string)),
,reatedFromEntity := d.GetOk("source_entity")
 createdFromEntity {
in.SourceEntity = quicksightschema.ExpandAnalysisSourceEntity(d.Get("source_entity").([]interface{}))
else {
in.Definition = quicksightschema.ExpandAnalysisDefinition(d.Get("definition").([]interface{}))
f, ok := d.GetOk("parameters"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
in.Parameters = quicksightschema.ExpandParameters(d.Get("parameters").([]interface{}))
oPrintf("[DEBUG] Updating QuickSight Analysis (%s): %#v", d.Id(), in)
 err := conn.UpdateAnalysisWithContext(ctx, in)
 err != nil {
return create.DiagError(names.QuickSight, create.ErrActionUpdating, ResNameAnalysis, d.Id(), err)
f, err := waitAnalysisUpdated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
return create.DiagError(names.QuickSight, create.ErrActionWaitingForUpdate, ResNameAnalysis, d.Id(), err)	}	if d.HasChange("permissions") {
aw, nraw := d.GetChange("permissions")
:= oraw.(*schema.Set)
:= nraw.(*schema.Set)toant, toRevoke := DiffPermissions(o.List(), n.List())para := &quicksight.UpdateAnalysisPermissionsInput{
AwsAccountId: aws.String(awsAccountId),
AnalysisId:aws.String(analysisId),
fen(toGrant) > 0 {
params.GrantPermissions = toGrant
fen(toRevoke) > 0 {
params.RevokePermissions = toRevoke
,rr = conn.UpdateAnalysisPermissionsWithContext(ctx, params)if e != nil {
return diag.Errorf("updating QuickSight Analysis (%s) permissions: %s", analysisId, err)	}	return resourceAnalysisRead(ctx, d, meta)
}func resourceAnalysisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).QuickSightConn(ctx)	awsAccountId, analysisId, err := ParseAnalysisId(d.Id())
	if err != nil {
turn diag.FromErr(err)
	}	input := &quicksight.DeleteAnalysisInput{
sAccountId: aws.String(awsAccountId),
alysisId:aws.String(analysisId),
	}	recoveryWindowInDays := d.Get("recovery_window_in_days").(int)
	if recoveryWindowInDays == 0 {
put.ForceDeleteWithoutRecovery = aws.Bool(true)
	} else {
put.RecoveryWindowInDays = aws.Int64(int64(recoveryWindowInDays))
	}	log.Printf("[INFO] Deleting QuickSight Analysis %s", d.Id())
	_, err = conn.DeleteAnalysisWithContext(ctx, input)	if tfawserr.ErrCodeEquals(err, quicksight.ErrCodeResourceNotFoundException) {
turn nil
	}	if err != nil {
turn create.DiagError(names.QuickSight, create.ErrActionDeleting, ResNameAnalysis, d.Id(), err)
	}	return nil
}func FindAnalysisByID(ctx context.Context, conn *quicksight.QuickSight, id string) (*quicksight.Analysis, error) {
	awsAccountId, analysisId, err := ParseAnalysisId(id)
	if err != nil {
turn nil, err
	}	descOpts := &quicksight.DescribeAnalysisInput{
sAccountId: aws.String(awsAccountId),
alysisId:aws.String(analysisId),
	}	out, err := conn.DescribeAnalysisWithContext(ctx, descOpts)	if tfawserr.ErrCodeEquals(err, quicksight.ErrCodeResourceNotFoundException) {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: descOpts,	}	if err != nil {
turn nil, err
	}	if out == nil || out.Analysis == nil {
turn nil, tfresource.NewEmptyResultError(descOpts)
	}	return out.Analysis, nil
}func ParseAnalysisId(id string) (string, string, error) {
	parts := strings.SplitN(id, ",", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
turn "", "", fmt.Errorf("unexpected format of ID (%s), expected AWS_ACCOUNT_ID,ANALYSIS_ID", id)
	}
	return parts[0], parts[1], nil
}func createAnalysisId(awsAccountID, analysisId string) string {
	return fmt.Sprintf("%s,%s", awsAccountID, analysisId)
}
