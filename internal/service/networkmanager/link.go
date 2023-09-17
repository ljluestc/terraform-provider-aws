// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package networkmanagerimport (
	"context"
	"fmt"
	"log"
	"strings"
	"time"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)// @SDKResource("aws_networkmanager_link", name="Link")
// @Tags(identifierAttribute="arn")
func ResourceLink() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourceLinkCreate,
adWithoutTimeout:resourceLinkRead,
dateWithoutTimeout: resourceLinkUpdate,
leteWithoutTimeout: resourceLinkDelete,Imrter: &schema.ResourceImporter{
StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
parsedARN, err := arn.Parse(d.Id())if err != nil {
return nil, fmt.Errorf("parsing ARN (%s): %w", d.Id(), err)
}// See https://docs.aws.amazon.com/service-authorization/latest/reference/list_networkmanager.html#networkmanager-resources-for-iam-policies.
resourceParts := strings.Split(parsedARN.Resource, "/")if actual, expected := len(resourceParts), 3; actual < expected {
return nil, fmt.Errorf("expected at least %d resource parts in ARN (%s), got: %d", expected, d.Id(), actual)
}d.SetId(resourceParts[2])
d.Set("global_network_id", resourceParts[1])return []*schema.ResourceData{d}, nil
},
CuomizeDiff: verify.SetTagsDiff,Timets: &schema.ResourceTimeout{
Create: schema.DefaultTimeout(10 * time.Minute),
Update: schema.DefaultTimeout(10 * time.Minute),
Delete: schema.DefaultTimeout(10 * time.Minute),
Scma: map[string]*schema.Schema{
"arn": {
Type:schema.TypeString,
Computed: true,
},
"bandwidth": {
Type:schema.TypeList,
Required: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"download_speed": {
Type:schema.TypeInt,
Optional: true,
},
"upload_speed": {
Type:schema.TypeInt,
Optional: true,
},
},
},
},
"description": {
Type:schema.TypeString,
Optional:true,
ValidateFunc: validation.StringLenBetween(0, 256),
},
"global_network_id": {
Type:schema.TypeString,
Required: true,
ForceNew: true,
},
"provider_name": {
Type:schema.TypeString,
Optional:true,
ValidateFunc: validation.StringLenBetween(0, 128),
},
"site_id": {
Type:schema.TypeString,
Required: true,
ForceNew: true,
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll: tftags.TagsSchemaComputed(),
"type": {
Type:schema.TypeString,
Optional:true,
ValidateFunc: validation.StringLenBetween(0, 128),
},	}
}func resourceLinkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).NetworkManagerConn(ctx)	globalNetworkID := d.Get("global_network_id").(string)
	input := &networkmanager.CreateLinkInput{
obalNetworkId: aws.String(globalNetworkID),
teId: aws.String(d.Get("site_id").(string)),
gs:getTagsIn(ctx),
	}	if v, ok := d.GetOk("bandwidth"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
put.Bandwidth = expandBandwidth(v.([]interface{})[0].(map[string]interface{}))
	}	if v, ok := d.GetOk("description"); ok {
put.Description = aws.String(v.(string))
	}	if v, ok := d.GetOk("provider_name"); ok {
put.Provider = aws.String(v.(string))
	}	if v, ok := d.GetOk("type"); ok {
put.Type = aws.String(v.(string))
	}	log.Printf("[DEBUG] Creating Network Manager Link: %s", input)
	output, err := conn.CreateLinkWithContext(ctx, input)	if err != nil {
turn diag.Errorf("creating Network Manager Link: %s", err)
	}	d.SetId(aws.StringValue(output.Link.LinkId))	if _, err := waitLinkCreated(ctx, conn, globalNetworkID, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
turn diag.Errorf("waiting for Network Manager Link (%s) create: %s", d.Id(), err)
	}	return resourceLinkRead(ctx, d, meta)
}func resourceLinkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).NetworkManagerConn(ctx)	globalNetworkID := d.Get("global_network_id").(string)
	link, err := FindLinkByTwoPartKey(ctx, conn, globalNetworkID, d.Id())	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] Network Manager Link %s not found, removing from state", d.Id())
SetId("")
turn nil
	}	if err != nil {
turn diag.Errorf("reading Network Manager Link (%s): %s", d.Id(), err)
	}	d.Set("arn", link.LinkArn)
	if link.Bandwidth != nil {
 err := d.Set("bandwidth", []interface{}{flattenBandwidth(link.Bandwidth)}); err != nil {
return diag.Errorf("setting bandwidth: %s", err)	} else {
Set("bandwidth", nil)
	}
	d.Set("description", link.Description)
	d.Set("global_network_id", link.GlobalNetworkId)
	d.Set("provider_name", link.Provider)
	d.Set("site_id", link.SiteId)
	d.Set("type", link.Type)	setTagsOut(ctx, link.Tags)	return nil
}func resourceLinkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).NetworkManagerConn(ctx)	if d.HasChangesExcept("tags", "tags_all") {
obalNetworkID := d.Get("global_network_id").(string)
put := &networkmanager.UpdateLinkInput{
Description:aws.String(d.Get("description").(string)),
GlobalNetworkId: aws.String(globalNetworkID),
LinkId: aws.String(d.Id()),
Provider:er_name").(string)),
Type:aws.String(d.Get("type").(string)),
f, ok := d.GetOk("bandwidth"); ok && len(v.([]interface{})) > 0 && v.([]interface{})[0] != nil {
input.Bandwidth = expandBandwidth(v.([]interface{})[0].(map[string]interface{}))
oPrintf("[DEBUG] Updating Network Manager Link: %s", input)
 err := conn.UpdateLinkWithContext(ctx, input)ifrr != nil {
return diag.Errorf("updating Network Manager Link (%s): %s", d.Id(), err)
f, err := waitLinkUpdated(ctx, conn, globalNetworkID, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
return diag.Errorf("waiting for Network Manager Link (%s) update: %s", d.Id(), err)	}	return resourceLinkRead(ctx, d, meta)
}func resourceLinkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).NetworkManagerConn(ctx)	globalNetworkID := d.Get("global_network_id").(string)	log.Printf("[DEBUG] Deleting Network Manager Link: %s", d.Id())
	_, err := conn.DeleteLinkWithContext(ctx, &networkmanager.DeleteLinkInput{
obalNetworkId: aws.String(globalNetworkID),
nkId: aws.String(d.Id()),
	})	if globalNetworkIDNotFoundError(err) || tfawserr.ErrCodeEquals(err, networkmanager.ErrCodeResourceNotFoundException) {
turn nil
	}	if err != nil {
turn diag.Errorf("deleting Network Manager Link (%s): %s", d.Id(), err)
	}	if _, err := waitLinkDeleted(ctx, conn, globalNetworkID, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
turn diag.Errorf("waiting for Network Manager Link (%s) delete: %s", d.Id(), err)
	}	return nil
}func FindLink(ctx context.Context, conn *networkmanager.NetworkManager, input *networkmanager.GetLinksInput) (*networkmanager.Link, error) {
	output, err := FindLinks(ctx, conn, input)	if err != nil {
turn nil, err
	}	if len(output) == 0 || output[0] == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	if count := len(output); count > 1 {
turn nil, tfresource.NewTooManyResultsError(count, input)
	}	return output[0], nil
}func FindLinks(ctx context.Context, conn *networkmanager.NetworkManager, input *networkmanager.GetLinksInput) ([]*networkmanager.Link, error) {
	var output []*networkmanager.Link	err := conn.GetLinksPagesWithContext(ctx, input, func(page *networkmanager.GetLinksOutput, lastPage bool) bool {
 page == nil {
return !lastPage
o_, v := range page.Links {
if v == nil {
continue
}output = append(output, v)
ern !lastPage
	})	if globalNetworkIDNotFoundError(err) {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if err != nil {
turn nil, err
	}	return output, nil
}func FindLinkByTwoPartKey(ctx context.Context, conn *networkmanager.NetworkManager, globalNetworkID, linkID string) (*networkmanager.Link, error) {
	input := &networkmanager.GetLinksInput{
obalNetworkId: aws.String(globalNetworkID),
nkIds:aws.StringSlice([]string{linkID}),
	}	output, err := FindLink(ctx, conn, input)	if err != nil {
turn nil, err
	}	// Eventual consistency check.
	if aws.StringValue(output.GlobalNetworkId) != globalNetworkID || aws.StringValue(output.LinkId) != linkID {
turn nil, &retry.NotFoundError{
LastRequest: input,	}	return output, nil
}func statusLinkState(ctx context.Context, conn *networkmanager.NetworkManager, globalNetworkID, linkID string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
tput, err := FindLinkByTwoPartKey(ctx, conn, globalNetworkID, linkID)iffresource.NotFound(err) {
return nil, "", nil
frr != nil {
return nil, "", err
ern output, aws.StringValue(output.State), nil
	}
}func waitLinkCreated(ctx context.Context, conn *networkmanager.NetworkManager, globalNetworkID, linkID string, timeout time.Duration) (*networkmanager.Link, error) {
	stateConf := &retry.StateChangeConf{
nding: []string{networkmanager.LinkStatePending},
rget:  []string{networkmanager.LinkStateAvailable},
meout: timeout,
fresh: statusLinkState(ctx, conn, globalNetworkID, linkID),
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)	if output, ok := outputRaw.(*networkmanager.Link); ok {
turn output, err
	}	return nil, err
}func waitLinkDeleted(ctx context.Context, conn *networkmanager.NetworkManager, globalNetworkID, linkID string, timeout time.Duration) (*networkmanager.Link, error) {
	stateConf := &retry.StateChangeConf{
nding: []string{networkmanager.LinkStateDeleting},
rget:  []string{},
meout: timeout,
fresh: statusLinkState(ctx, conn, globalNetworkID, linkID),
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)	if output, ok := outputRaw.(*networkmanager.Link); ok {
turn output, err
	}	return nil, err
}func waitLinkUpdated(ctx context.Context, conn *networkmanager.NetworkManager, globalNetworkID, linkID string, timeout time.Duration) (*networkmanager.Link, error) {
	stateConf := &retry.StateChangeConf{
nding: []string{networkmanager.LinkStateUpdating},
rget:  []string{networkmanager.LinkStateAvailable},
meout: timeout,
fresh: statusLinkState(ctx, conn, globalNetworkID, linkID),
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)	if output, ok := outputRaw.(*networkmanager.Link); ok {
turn output, err
	}	return nil, err
}func expandBandwidth(tfMap map[string]interface{}) *networkmanager.Bandwidth {
	if tfMap == nil {
turn nil
	}	apiObject := &networkmanager.Bandwidth{}	if v, ok := tfMap["download_speed"].(int); ok && v != 0 {
iObject.DownloadSpeed = aws.Int64(int64(v))
	}	if v, ok := tfMap["upload_speed"].(int); ok && v != 0 {
iObject.UploadSpeed = aws.Int64(int64(v))
	}	return apiObject
}func flattenBandwidth(apiObject *networkmanager.Bandwidth) map[string]interface{} {
	if apiObject == nil {
turn nil
	}	tfMap := map[string]interface{}{}	if v := apiObject.DownloadSpeed; v != nil {
Map["download_speed"] = aws.Int64Value(v)
	}	if v := apiObject.UploadSpeed; v != nil {
Map["upload_speed"] = aws.Int64Value(v)
	}	return tfMap
}
