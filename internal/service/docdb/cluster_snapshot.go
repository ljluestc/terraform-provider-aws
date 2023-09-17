//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagedocdb

import(
"context"
"fmt"
"log"
"time"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/docdb"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
"github.com/hashicorp/terraform-provider-aws/internal/flex"
)

//@SDKResource("aws_docdb_cluster_snapshot")
funcResourceClusterSnapshot()*schema.Resource{
return&schema.Resource{
CreateWithoutTimeout:resourceClusterSnapshotCreate,
ReadWithoutTimeout:resourceClusterSnapshotRead,
DeleteWithoutTimeout:resourceClusterSnapshotDelete,
Importer:&schema.ResourceImporter{
StateContext:schema.ImportStatePassthroughContext,
},

Timeouts:&schema.ResourceTimeout{
Create:schema.DefaultTimeout(20*time.Minute),
},

Schema:map[string]*schema.Schema{
"db_cluster_snapshot_identifier":{
Type:ema.TypeString,
ValidateFunc:validClusterSnapshotIdentifier,
Required:true,
ForceNew:true,
},
"db_cluster_identifier":{
Type:ema.TypeString,
ValidateFunc:validClusterIdentifier,
Required:true,
ForceNew:true,
},

"availability_zones":{
Type:schema.TypeList,
Elem:&schema.Schema{Type:schema.TypeString},
Computed:true,
},
"db_cluster_snapshot_arn":{
Type:schema.TypeString,
Computed:true,
},
"storage_encrypted":{
Type:schema.TypeBool,
Computed:true,
},
"engine":{
Type:schema.TypeString,
Computed:true,
},
"engine_version":{
Type:schema.TypeString,
Computed:true,
},
"kms_key_id":{
Type:schema.TypeString,
Computed:true,
},
"port":{
Type:schema.TypeInt,
Computed:true,
},
"source_db_cluster_snapshot_arn":{
Type:schema.TypeString,
Computed:true,
},
"snapshot_type":{
Type:schema.TypeString,
Computed:true,
},
"status":{
Type:schema.TypeString,
Computed:true,
},
"vpc_id":{
Type:schema.TypeString,
Computed:true,
},
},
}
}
funcresourceClusterSnapshotCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

params:=&docdb.CreateDBClusterSnapshotInput{
DBClusterIdentifier:.String(d.Get("db_cluster_identifier").(string)),
DBClusterSnapshotIdentifier:aws.String(d.Get("db_cluster_snapshot_identifier").(string)),
}

_,err:=conn.CreateDBClusterSnapshotWithContext(ctx,params)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"creatingDocumentDBClusterSnapshot:%s",err)
}
d.SetId(d.Get("db_cluster_snapshot_identifier").(string))

stateConf:=&retry.StateChangeConf{
Pending:[]string{"creating"},
Target:[]string{"available"},
Refresh:resourceClusterSnapshotStateRefreshFunc(ctx,d.Id(),conn),
Timeout:d.Timeout(schema.TimeoutCreate),
MinTimeout:10*time.Second,
Delay:me.Second,
}

//Wait,catchinganyerrors
_,err=stateConf.WaitForStateContext(ctx)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"waitingforDocumentDBClusterSnapshot%qtocreate:%s",d.Id(),err)
}

returnappend(diags,resourceClusterSnapshotRead(ctx,d,meta)...)
}
funcresourceClusterSnapshotRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

params:=&docdb.DescribeDBClusterSnapshotsInput{
DBClusterSnapshotIdentifier:aws.String(d.Id()),
}
resp,err:=conn.DescribeDBClusterSnapshotsWithContext(ctx,params)
iferr!=nil{
if!d.IsNewResource()&&tfawserr.ErrCodeEquals(err,docdb.ErrCodeDBClusterSnapshotNotFoundFault){
log.Printf("[WARN]DocumentDBClusterSnapshot(%s)notfound,removingfromstate",d.Id())
d.SetId("")
returndiags
}
returnsdkdiag.AppendErrorf(diags,"readingDocumentDBClusterSnapshot%q:%s",d.Id(),err)
}

if!d.IsNewResource()&&(resp==nil||len(resp.DBClusterSnapshots)==0||resp.DBClusterSnapshots[0]==nil||aws.StringValue(resp.DBClusterSnapshots[0].DBClusterSnapshotIdentifier)!=d.Id()){
log.Printf("[WARN]DocumentDBClusterSnapshot(%s)notfound,removingfromstate",d.Id())
d.SetId("")
returndiags
}

snapshot:=resp.DBClusterSnapshots[0]

iferr:=d.Set("availability_zones",flex.FlattenStringList(snapshot.AvailabilityZones));err!=nil{
returnsdkdiag.AppendErrorf(diags,"settingavailability_zones:%s",err)
}
d.Set("db_cluster_identifier",snapshot.DBClusterIdentifier)
d.Set("db_cluster_snapshot_arn",snapshot.DBClusterSnapshotArn)
d.Set("db_cluster_snapshot_identifier",snapshot.DBClusterSnapshotIdentifier)
d.Set("engine_version",snapshot.EngineVersion)
d.Set("engine",snapshot.Engine)
d.Set("kms_key_id",snapshot.KmsKeyId)
d.Set("port",snapshot.Port)
d.Set("snapshot_type",snapshot.SnapshotType)
d.Set("source_db_cluster_snapshot_arn",snapshot.SourceDBClusterSnapshotArn)
d.Set("status",snapshot.Status)
d.Set("storage_encrypted",snapshot.StorageEncrypted)
d.Set("vpc_id",snapshot.VpcId)

returndiags
}
funcresourceClusterSnapshotDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

params:=&docdb.DeleteDBClusterSnapshotInput{
DBClusterSnapshotIdentifier:aws.String(d.Id()),
}
_,err:=conn.DeleteDBClusterSnapshotWithContext(ctx,params)
iferr!=nil{
iftfawserr.ErrCodeEquals(err,docdb.ErrCodeDBClusterSnapshotNotFoundFault){
returndiags
}
returnsdkdiag.AppendErrorf(diags,"deletingDocumentDBClusterSnapshot%q:%s",d.Id(),err)
}

returndiags
}
funcresourceClusterSnapshotStateRefreshFunc(ctxcontext.Context,dbClusterSnapshotIdentifierstring,conn*docdb.DocDB)retry.StateRefreshFunc{
returnfunc()(interface{},string,error){
opts:=&docdb.DescribeDBClusterSnapshotsInput{
DBClusterSnapshotIdentifier:aws.String(dbClusterSnapshotIdentifier),
}

resp,err:=conn.DescribeDBClusterSnapshotsWithContext(ctx,opts)
iferr!=nil{
iftfawserr.ErrCodeEquals(err,docdb.ErrCodeDBClusterSnapshotNotFoundFault){
returnnil,"",nil
}
returnnil,"",fmt.Errorf("retrievingDocumentDBClusterSnapshots:%s",err)
}

ifresp==nil||len(resp.DBClusterSnapshots)==0||resp.DBClusterSnapshots[0]==nil{
returnnil,"",fmt.Errorf("Nosnapshotsreturnedfor%s",dbClusterSnapshotIdentifier)
}

snapshot:=resp.DBClusterSnapshots[0]

returnresp,aws.StringValue(snapshot.Status),nil
}
}
