// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package glueimport (
	"context"
	"log"
	"time"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)// @SDKResource("aws_glue_partition_index")
func ResourcePartitionIndex() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourcePartitionIndexCreate,
adWithoutTimeout:resourcePartitionIndexRead,
leteWithoutTimeout: resourcePartitionIndexDelete,
porter: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,
Scma: map[string]*schema.Schema{
"catalog_id": {
Type:schema.TypeString,
ForceNew: true,
Optional: true,
Computed: true,
},
"database_name": {
Type:schema.TypeString,
ForceNew:true,
Required:true,
ValidateFunc: validation.StringLenBetween(1, 255),
},
"table_name": {
Type:schema.TypeString,
ForceNew:true,
Required:true,
ValidateFunc: validation.StringLenBetween(1, 255),
},
"partition_index": {
Type:schema.TypeList,
Required: true,
ForceNew: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"index_name": {
	Type:schema.TypeString,
	Optional: true,
	ForceNew: true,
},
"index_status": {
	Type:schema.TypeString,
	Computed: true,
},
"keys": {
	Type:schema.TypeList,
	Optional: true,
	ForceNew: true,
	Elem:&schema.Schema{Type: schema.TypeString},
},
},
},
},
Tiouts: &schema.ResourceTimeout{
Create: schema.DefaultTimeout(10 * time.Minute),
Delete: schema.DefaultTimeout(10 * time.Minute),	}
}func resourcePartitionIndexCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).GlueConn(ctx)
	catalogID := createCatalogID(d, meta.(*conns.AWSClient).AccountID)
	dbName := d.Get("database_name").(string)
	tableName := d.Get("table_name").(string)	input := &glue.CreatePartitionIndexInput{
talogId: aws.String(catalogID),
tabaseName:aws.String(dbName),
bleName: aws.String(tableName),
rtitionIndex: expandPartitionIndex(d.Get("partition_index").([]interface{})),
	}	log.Printf("[DEBUG] Creating Glue Partition Index: %#v", input)
	_, err := conn.CreatePartitionIndexWithContext(ctx, input)
	if err != nil {
turn sdkdiag.AppendErrorf(diags, "creating Glue Partition Index: %s", err)
	}	d.SetId(createPartitionIndexID(catalogID, dbName, tableName, aws.StringValue(input.PartitionIndex.IndexName)))	if _, err := waitPartitionIndexCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
turn sdkdiag.AppendErrorf(diags, "while waiting for Glue Partition Index (%s) to become available: %s", d.Id(), err)
	}	return append(diags, resourcePartitionIndexRead(ctx, d, meta)...)
}func resourcePartitionIndexRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).GlueConn(ctx)	catalogID, dbName, tableName, _, err := readPartitionIndexID(d.Id())
	if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading Glue Partition Index (%s): %s", d.Id(), err)
	}	log.Printf("[DEBUG] Reading Glue Partition Index: %s", d.Id())
	partition, err := FindPartitionIndexByName(ctx, conn, d.Id())
	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] Glue Partition Index (%s) not found, removing from state", d.Id())
SetId("")
turn diags
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading Glue Partition Index (%s): %s", d.Id(), err)
	}	d.Set("table_name", tableName)
	d.Set("catalog_id", catalogID)
	d.Set("database_name", dbName)	if err := d.Set("partition_index", []map[string]interface{}{flattenPartitionIndex(partition)}); err != nil {
turn sdkdiag.AppendErrorf(diags, "setting partition_index: %s", err)
	}	return diags
}func resourcePartitionIndexDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).GlueConn(ctx)	catalogID, dbName, tableName, partIndex, err := readPartitionIndexID(d.Id())
	if err != nil {
turn sdkdiag.AppendErrorf(diags, "deleting Glue Partition Index: %s", err)
	}	log.Printf("[DEBUG] Deleting Glue Partition Index: %s", d.Id())
	_, err = conn.DeletePartitionIndexWithContext(ctx, &glue.DeletePartitionIndexInput{
talogId:aws.String(catalogID),
bleName:aws.String(tableName),
tabaseName: aws.String(dbName),
dexName:aws.String(partIndex),
	})
	if err != nil {
 tfawserr.ErrCodeEquals(err, glue.ErrCodeEntityNotFoundException) {
return diagsturn sdkdiag.AppendErrorf(diags, "deleting Glue Partition Index: %s", err)
	}	if _, err := waitPartitionIndexDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
turn sdkdiag.AppendErrorf(diags, "while waiting for Glue Partition Index (%s) to be deleted: %s", d.Id(), err)
	}	return diags
}func expandPartitionIndex(l []interface{}) *glue.PartitionIndex {
	if len(l) == 0 || l[0] == nil {
turn nil
	}	s := l[0].(map[string]interface{})
	parIndex := &glue.PartitionIndex{}	if v, ok := s["keys"].([]interface{}); ok && len(v) > 0 {
rIndex.Keys = flex.ExpandStringList(v)
	}	if v, ok := s["index_name"].(string); ok && v != "" {
rIndex.IndexName = aws.String(v)
	}	return parIndex
}
