//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagedocdb

import(
"context"
"fmt"
"log"
"strings"
"time"

"github.com/YakDriver/regexache"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/docdb"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
"github.com/hashicorp/terraform-provider-aws/internal/flex"
tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
"github.com/hashicorp/terraform-provider-aws/names"
)

//@SDKResource("aws_docdb_cluster",name="Cluster")
//@Tags(identifierAttribute="arn")
funcResourceCluster()*schema.Resource{
return&schema.Resource{
CreateWithoutTimeout:resourceClusterCreate,
ReadWithoutTimeout:resourceClusterRead,
UpdateWithoutTimeout:resourceClusterUpdate,
DeleteWithoutTimeout:resourceClusterDelete,
Importer:&schema.ResourceImporter{
StateContext:resourceClusterImport,
},

Timeouts:&schema.ResourceTimeout{
Create:schema.DefaultTimeout(120*time.Minute),
Update:schema.DefaultTimeout(120*time.Minute),
Delete:schema.DefaultTimeout(120*time.Minute),
},

Schema:map[string]*schema.Schema{
"arn":{
Type:schema.TypeString,
Computed:true,
},

"availability_zones":{
Type:schema.TypeSet,
Elem:&schema.Schema{Type:schema.TypeString},
Optional:true,
ForceNew:true,
Computed:true,
Set:.HashString,
},

"cluster_identifier":{
Type:hema.TypeString,
Optional:
Computed:
ForceNew:
ConflictsWith:[]string{"cluster_identifier_prefix"},
ValidateFunc:validIdentifier,
},
"cluster_identifier_prefix":{
Type:hema.TypeString,
Optional:
Computed:
ForceNew:
ConflictsWith:[]string{"cluster_identifier"},
ValidateFunc:validIdentifierPrefix,
},

"cluster_members":{
Type:schema.TypeSet,
Elem:&schema.Schema{Type:schema.TypeString},
Optional:true,
Computed:true,
Set:.HashString,
},

"db_subnet_group_name":{
Type:schema.TypeString,
Optional:true,
ForceNew:true,
Computed:true,
},

"db_cluster_parameter_group_name":{
Type:schema.TypeString,
Optional:true,
Computed:true,
},

"endpoint":{
Type:schema.TypeString,
Computed:true,
},

"global_cluster_identifier":{
Type:ema.TypeString,
Optional:true,
ValidateFunc:validGlobalCusterIdentifier,
},

"reader_endpoint":{
Type:schema.TypeString,
Computed:true,
},

"hosted_zone_id":{
Type:schema.TypeString,
Computed:true,
},

"engine":{
Type:ema.TypeString,
Optional:true,
Default:",
ForceNew:true,
ValidateFunc:validEngine(),
},

"engine_version":{
Type:schema.TypeString,
Optional:true,
Computed:true,
},

"storage_encrypted":{
Type:schema.TypeBool,
Optional:true,
ForceNew:true,
},

"final_snapshot_identifier":{
Type:schema.TypeString,
Optional:true,
ValidateFunc:func(vinterface{},kstring)(ws[]string,es[]error){
value:=v.(string)
if!regexache.MustCompile(`^[0-9A-Za-z-]+$`).MatchString(value){
es=append(es,fmt.Errorf(
"onlyalphanumericcharactersandhyphensallowedin%q",k))
}
ifregexache.MustCompile(`--`).MatchString(value){
es=append(es,fmt.Errorf("%qcannotcontaintwoconsecutivehyphens",k))
}
ifregexache.MustCompile(`-$`).MatchString(value){
es=append(es,fmt.Errorf("%qcannotendinahyphen",k))
}
return
},
},

"skip_final_snapshot":{
Type:schema.TypeBool,
Optional:true,
Default:false,
},

"master_username":{
Type:schema.TypeString,
Computed:true,
Optional:true,
ForceNew:true,
},

"master_password":{
Type:.TypeString,
Optional:true,
Sensitive:true,
},

"snapshot_identifier":{
Type:schema.TypeString,
Optional:true,
ForceNew:true,
DiffSuppressFunc:func(k,old,newstring,d*schema.ResourceData)bool{
//allowsnapshot_idenfitiertoberemovedwithoutforcingre-creation
returnnew==""
},
},

"port":{
Type:ema.TypeInt,
Optional:true,
Default:
ForceNew:true,
ValidateFunc:validation.IntBetween(1150,65535),
},

"apply_immediately":{
Type:schema.TypeBool,
Optional:true,
Computed:true,
},

"vpc_security_group_ids":{
Type:schema.TypeSet,
Optional:true,
Computed:true,
Elem:&schema.Schema{Type:schema.TypeString},
Set:.HashString,
},

"preferred_backup_window":{
Type:ema.TypeString,
Optional:true,
Computed:true,
ValidateFunc:verify.ValidOnceADayWindowFormat,
},

"preferred_maintenance_window":{
Type:schema.TypeString,
Optional:true,
Computed:true,
StateFunc:func(valinterface{})string{
ifval==nil{
return""
}
returnstrings.ToLower(val.(string))
},
ValidateFunc:verify.ValidOnceAWeekWindowFormat,
},

"backup_retention_period":{
Type:ema.TypeInt,
Optional:true,
Default:
ValidateFunc:validation.IntAtMost(35),
},

"kms_key_id":{
Type:ema.TypeString,
Optional:true,
Computed:true,
ForceNew:true,
ValidateFunc:verify.ValidARN,
},

"cluster_resource_id":{
Type:schema.TypeString,
Computed:true,
},

"enabled_cloudwatch_logs_exports":{
Type:schema.TypeList,
Optional:true,
Elem:&schema.Schema{
Type:schema.TypeString,
ValidateFunc:validation.StringInSlice([]string{
"audit",
"profiler",
},false),
},
},

"deletion_protection":{
Type:schema.TypeBool,
Optional:true,
},

names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll:tftags.TagsSchemaComputed(),
},

CustomizeDiff:verify.SetTagsDiff,
}
}
funcresourceClusterImport(ctxcontext.Context,
d*schema.ResourceData,metainterface{})([]*schema.ResourceData,error){
//Neitherskip_final_snapshotnorfinal_snapshot_identifiercanbefetched
//fromanyAPIcall,soweneedtodefaultskip_final_snapshottotrueso
//thatfinal_snapshot_identifierisnotrequired
d.Set("skip_final_snapshot",true)
return[]*schema.ResourceData{d},nil
}
funcresourceClusterCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

//SomeAPIcalls(e.g.RestoreDBClusterFromSnapshotdonotsupportall
//parameterstocorrectlyapplyallsettingsinonepass.Formissing
//parametersorunsupportedconfigurations,wemayneedtocall
//ModifyDBInstanceafterwadocdbtopreventTerraformoperatorsfromAPI
//errorsorneedingtodoubleapply.
varrequiresModifyDbClusterbool
modifyDbClusterInput:=&docdb.ModifyDBClusterInput{
ApplyImmediately:aws.Bool(true),
}

varidentifierstring
ifv,ok:=d.GetOk("cluster_identifier");ok{
identifier=v.(string)
}elseifv,ok:=d.GetOk("cluster_identifier_prefix");ok{
identifier=id.PrefixedUniqueId(v.(string))
}else{
identifier=id.PrefixedUniqueId("tf-")
}

if_,ok:=d.GetOk("snapshot_identifier");ok{
opts:=docdb.RestoreDBClusterFromSnapshotInput{
DBClusterIdentifier:aws.String(identifier),
Engine:aws.String(d.Get("engine").(string)),
SnapshotIdentifier:aws.String(d.Get("snapshot_identifier").(string)),
DeletionProtection:aws.Bool(d.Get("deletion_protection").(bool)),
Tags:getTagsIn(ctx),
}

ifattr:=d.Get("availability_zones").(*schema.Set);attr.Len()>0{
opts.AvailabilityZones=flex.ExpandStringSet(attr)
}

ifattr,ok:=d.GetOk("backup_retention_period");ok{
modifyDbClusterInput.BackupRetentionPeriod=aws.Int64(int64(attr.(int)))
requiresModifyDbCluster=true
}

ifattr,ok:=d.GetOk("db_subnet_group_name");ok{
opts.DBSubnetGroupName=aws.String(attr.(string))
}

ifattr,ok:=d.GetOk("db_cluster_parameter_group_name");ok{
modifyDbClusterInput.DBClusterParameterGroupName=aws.String(attr.(string))
requiresModifyDbCluster=true
}

ifattr,ok:=d.GetOk("enabled_cloudwatch_logs_exports");ok&&len(attr.([]interface{}))>0{
opts.EnableCloudwatchLogsExports=flex.ExpandStringList(attr.([]interface{}))
}

ifattr,ok:=d.GetOk("engine_version");ok{
opts.EngineVersion=aws.String(attr.(string))
}

ifattr,ok:=d.GetOk("kms_key_id");ok{
opts.KmsKeyId=aws.String(attr.(string))
}

ifattr,ok:=d.GetOk("port");ok{
opts.Port=aws.Int64(int64(attr.(int)))
}

ifattr,ok:=d.GetOk("preferred_backup_window");ok{
modifyDbClusterInput.PreferredBackupWindow=aws.String(attr.(string))
requiresModifyDbCluster=true
}

ifattr,ok:=d.GetOk("preferred_maintenance_window");ok{
modifyDbClusterInput.PreferredMaintenanceWindow=aws.String(attr.(string))
requiresModifyDbCluster=true
}

ifattr:=d.Get("vpc_security_group_ids").(*schema.Set);attr.Len()>0{
opts.VpcSecurityGroupIds=flex.ExpandStringSet(attr)
}

err:=retry.RetryContext(ctx,propagationTimeout,func()*retry.RetryError{
_,err:=conn.RestoreDBClusterFromSnapshotWithContext(ctx,&opts)
iferr!=nil{
iftfawserr.ErrMessageContains(err,"InvalidParameterValue","IAMroleARNvalueisinvalidordoesnotincludetherequiredpermissions"){
returnretry.RetryableError(err)
}
returnretry.NonRetryableError(err)
}
returnnil
})
iftfresource.TimedOut(err){
_,err=conn.RestoreDBClusterFromSnapshotWithContext(ctx,&opts)
}
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"creatingDocumentDBCluster:%s",err)
}
}else{
//SecondaryDocDBclusterspartofaglobalclusterwillnotsupplythemaster_password
if_,ok:=d.GetOk("global_cluster_identifier");!ok{
if_,ok:=d.GetOk("master_password");!ok{
returnsdkdiag.AppendErrorf(diags,`provider.aws:aws_docdb_cluster:%s:"master_password":requiredfieldisnotset`,identifier)
}
}

//SecondaryDocDBclusterspartofaglobalclusterwillnotsupplythemaster_username
if_,ok:=d.GetOk("global_cluster_identifier");!ok{
if_,ok:=d.GetOk("master_username");!ok{
returnsdkdiag.AppendErrorf(diags,`provider.aws:aws_docdb_cluster:%s:"master_username":requiredfieldisnotset`,identifier)
}
}

createOpts:=&docdb.CreateDBClusterInput{
DBClusterIdentifier:aws.String(identifier),
Engine:aws.String(d.Get("engine").(string)),
MasterUserPassword:aws.String(d.Get("master_password").(string)),
MasterUsername:ring(d.Get("master_username").(string)),
DeletionProtection:aws.Bool(d.Get("deletion_protection").(bool)),
Tags:getTagsIn(ctx),
}

ifattr,ok:=d.GetOk("global_cluster_identifier");ok{
createOpts.GlobalClusterIdentifier=aws.String(attr.(string))
}

ifattr,ok:=d.GetOk("port");ok{
createOpts.Port=aws.Int64(int64(attr.(int)))
}

ifattr,ok:=d.GetOk("db_subnet_group_name");ok{
createOpts.DBSubnetGroupName=aws.String(attr.(string))
}

ifattr,ok:=d.GetOk("db_cluster_parameter_group_name");ok{
createOpts.DBClusterParameterGroupName=aws.String(attr.(string))
}

ifattr,ok:=d.GetOk("engine_version");ok{
createOpts.EngineVersion=aws.String(attr.(string))
}

ifattr:=d.Get("vpc_security_group_ids").(*schema.Set);attr.Len()>0{
createOpts.VpcSecurityGroupIds=flex.ExpandStringSet(attr)
}

ifattr:=d.Get("availability_zones").(*schema.Set);attr.Len()>0{
createOpts.AvailabilityZones=flex.ExpandStringSet(attr)
}

ifv,ok:=d.GetOk("backup_retention_period");ok{
createOpts.BackupRetentionPeriod=aws.Int64(int64(v.(int)))
}

ifv,ok:=d.GetOk("preferred_backup_window");ok{
createOpts.PreferredBackupWindow=aws.String(v.(string))
}

ifv,ok:=d.GetOk("preferred_maintenance_window");ok{
createOpts.PreferredMaintenanceWindow=aws.String(v.(string))
}

ifattr,ok:=d.GetOk("kms_key_id");ok{
createOpts.KmsKeyId=aws.String(attr.(string))
}

ifattr,ok:=d.GetOk("enabled_cloudwatch_logs_exports");ok&&len(attr.([]interface{}))>0{
createOpts.EnableCloudwatchLogsExports=flex.ExpandStringList(attr.([]interface{}))
}

ifattr,ok:=d.GetOkExists("storage_encrypted");ok{
createOpts.StorageEncrypted=aws.Bool(attr.(bool))
}

err:=retry.RetryContext(ctx,propagationTimeout,func()*retry.RetryError{
varerrerror
_,err=conn.CreateDBClusterWithContext(ctx,createOpts)
iferr!=nil{
iftfawserr.ErrMessageContains(err,"InvalidParameterValue","IAMroleARNvalueisinvalidordoesnotincludetherequiredpermissions"){
returnretry.RetryableError(err)
}
returnretry.NonRetryableError(err)
}
returnnil
})
iftfresource.TimedOut(err){
_,err=conn.CreateDBClusterWithContext(ctx,createOpts)
}
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"creatingDocumentDBcluster:%s",err)
}
}

d.SetId(identifier)

log.Printf("[INFO]DocumentDBClusterID:%s",d.Id())

log.Println(
"[INFO]WaitingforDocumentDBClustertobeavailable")

stateConf:=&retry.StateChangeConf{
Pending:resourceClusterCreatePendingStates,
Target:[]string{"available"},
Refresh:resourceClusterStateRefreshFunc(ctx,conn,d.Id()),
Timeout:d.Timeout(schema.TimeoutCreate),
MinTimeout:10*time.Second,
Delay:ime.Second,
}

//Wait,catchinganyerrors
_,err:=stateConf.WaitForStateContext(ctx)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"waitingforDocumentDBClusterstatetobe\"available\":%s",err)
}

ifrequiresModifyDbCluster{
modifyDbClusterInput.DBClusterIdentifier=aws.String(d.Id())

log.Printf("[INFO]DocumentDBCluster(%s)configurationrequiresModifyDBCluster:%s",d.Id(),modifyDbClusterInput)
_,err:=conn.ModifyDBClusterWithContext(ctx,modifyDbClusterInput)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"modifyingDocumentDBCluster(%s):%s",d.Id(),err)
}

log.Printf("[INFO]WaitingforDocumentDBCluster(%s)tobeavailable",d.Id())
err=waitForClusterUpdate(ctx,conn,d.Id(),d.Timeout(schema.TimeoutCreate))
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"waitingforDocumentDBCluster(%s)tobeavailable:%s",d.Id(),err)
}
}

returnappend(diags,resourceClusterRead(ctx,d,meta)...)
}
funcresourceClusterRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

input:=&docdb.DescribeDBClustersInput{
DBClusterIdentifier:aws.String(d.Id()),
}

resp,err:=conn.DescribeDBClustersWithContext(ctx,input)

if!d.IsNewResource()&&tfawserr.ErrCodeEquals(err,docdb.ErrCodeDBClusterNotFoundFault){
log.Printf("[WARN]DocumentDBCluster(%s)notfound,removingfromstate",d.Id())
d.SetId("")
returndiags
}

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"describingDocumentDBCluster(%s):%s",d.Id(),err)
}

ifresp==nil{
returnsdkdiag.AppendErrorf(diags,"retrievingDocumentDBcluster:emptyresponsefor:%s",input)
}

vardbc*docdb.DBCluster
for_,c:=rangeresp.DBClusters{
ifaws.StringValue(c.DBClusterIdentifier)==d.Id(){
dbc=c
break
}
}

if!d.IsNewResource()&&dbc==nil{
log.Printf("[WARN]DocumentDBCluster(%s)notfound,removingfromstate",d.Id())
d.SetId("")
returndiags
}

globalCluster,err:=findGlobalClusterByARN(ctx,conn,aws.StringValue(dbc.DBClusterArn))

//IgnorethefollowingAPIerrorforregions/partitionsthatdonotsupportDocDBGlobalClusters:
//InvalidParameterValue:AccessDeniedtoAPIVersion:APIGlobalDatabases
iferr!=nil&&!tfawserr.ErrMessageContains(err,"InvalidParameterValue","AccessDeniedtoAPIVersion:APIGlobalDatabases"){
returnsdkdiag.AppendErrorf(diags,"readingDocumentDBGlobalClusterinformationforDBCluster(%s):%s",d.Id(),err)
}

ifglobalCluster!=nil{
d.Set("global_cluster_identifier",globalCluster.GlobalClusterIdentifier)
}else{
d.Set("global_cluster_identifier","")
}

iferr:=d.Set("availability_zones",aws.StringValueSlice(dbc.AvailabilityZones));err!=nil{
returnsdkdiag.AppendErrorf(diags,"settingavailability_zones:%s",err)
}

d.Set("arn",dbc.DBClusterArn)
d.Set("backup_retention_period",dbc.BackupRetentionPeriod)
d.Set("cluster_identifier",dbc.DBClusterIdentifier)

varcm[]string
for_,m:=rangedbc.DBClusterMembers{
cm=append(cm,aws.StringValue(m.DBInstanceIdentifier))
}
iferr:=d.Set("cluster_members",cm);err!=nil{
returnsdkdiag.AppendErrorf(diags,"settingcluster_members:%s",err)
}

d.Set("cluster_resource_id",dbc.DbClusterResourceId)
d.Set("db_cluster_parameter_group_name",dbc.DBClusterParameterGroup)
d.Set("db_subnet_group_name",dbc.DBSubnetGroup)

iferr:=d.Set("enabled_cloudwatch_logs_exports",aws.StringValueSlice(dbc.EnabledCloudwatchLogsExports));err!=nil{
returnsdkdiag.AppendErrorf(diags,"settingenabled_cloudwatch_logs_exports:%s",err)
}

d.Set("endpoint",dbc.Endpoint)
d.Set("engine_version",dbc.EngineVersion)
d.Set("engine",dbc.Engine)
d.Set("hosted_zone_id",dbc.HostedZoneId)

d.Set("kms_key_id",dbc.KmsKeyId)
d.Set("master_username",dbc.MasterUsername)
d.Set("port",dbc.Port)
d.Set("preferred_backup_window",dbc.PreferredBackupWindow)
d.Set("preferred_maintenance_window",dbc.PreferredMaintenanceWindow)
d.Set("reader_endpoint",dbc.ReaderEndpoint)
d.Set("storage_encrypted",dbc.StorageEncrypted)
d.Set("deletion_protection",dbc.DeletionProtection)

varvpcg[]string
for_,g:=rangedbc.VpcSecurityGroups{
vpcg=append(vpcg,aws.StringValue(g.VpcSecurityGroupId))
}
iferr:=d.Set("vpc_security_group_ids",vpcg);err!=nil{
returnsdkdiag.AppendErrorf(diags,"settingvpc_security_group_ids:%s",err)
}

returndiags
}
funcresourceClusterUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)
requestUpdate:=false

req:=&docdb.ModifyDBClusterInput{
ApplyImmediately:aws.Bool(d.Get("apply_immediately").(bool)),
DBClusterIdentifier:aws.String(d.Id()),
}

ifd.HasChange("master_password"){
req.MasterUserPassword=aws.String(d.Get("master_password").(string))
requestUpdate=true
}

ifd.HasChange("engine_version"){
req.EngineVersion=aws.String(d.Get("engine_version").(string))
requestUpdate=true
}

ifd.HasChange("vpc_security_group_ids"){
ifattr:=d.Get("vpc_security_group_ids").(*schema.Set);attr.Len()>0{
req.VpcSecurityGroupIds=flex.ExpandStringSet(attr)
}else{
req.VpcSecurityGroupIds=[]*string{}
}
requestUpdate=true
}

ifd.HasChange("preferred_backup_window"){
req.PreferredBackupWindow=aws.String(d.Get("preferred_backup_window").(string))
requestUpdate=true
}

ifd.HasChange("preferred_maintenance_window"){
req.PreferredMaintenanceWindow=aws.String(d.Get("preferred_maintenance_window").(string))
requestUpdate=true
}

ifd.HasChange("backup_retention_period"){
req.BackupRetentionPeriod=aws.Int64(int64(d.Get("backup_retention_period").(int)))
requestUpdate=true
}

ifd.HasChange("db_cluster_parameter_group_name"){
req.DBClusterParameterGroupName=aws.String(d.Get("db_cluster_parameter_group_name").(string))
requestUpdate=true
}

ifd.HasChange("enabled_cloudwatch_logs_exports"){
req.CloudwatchLogsExportConfiguration=buildCloudWatchLogsExportConfiguration(d)
requestUpdate=true
}

ifd.HasChange("deletion_protection"){
req.DeletionProtection=aws.Bool(d.Get("deletion_protection").(bool))
requestUpdate=true
}

ifd.HasChange("global_cluster_identifier"){
oRaw,nRaw:=d.GetChange("global_cluster_identifier")
o:=oRaw.(string)
n:=nRaw.(string)

ifo==""{
returnsdkdiag.AppendErrorf(diags,"existingDocumentDBClusterscannotbeaddedtoanexistingDocumentDBGlobalCluster")
}

ifn!=""{
returnsdkdiag.AppendErrorf(diags,"existingDocumentDBClusterscannotbemigratedbetweenexistingDocumentDBGlobalClusters")
}

input:=&docdb.RemoveFromGlobalClusterInput{
DbClusterIdentifier:aws.String(d.Get("arn").(string)),
GlobalClusterIdentifier:aws.String(o),
}

_,err:=conn.RemoveFromGlobalClusterWithContext(ctx,input)

iferr!=nil&&!tfawserr.ErrCodeEquals(err,docdb.ErrCodeGlobalClusterNotFoundFault)&&!tfawserr.ErrMessageContains(err,"InvalidParameterValue","isnotfoundinglobalcluster"){
returnsdkdiag.AppendErrorf(diags,"removingDocumentDBCluster(%s)fromDocumentDBGlobalCluster:%s",d.Id(),err)
}
}

ifrequestUpdate{
err:=retry.RetryContext(ctx,5*time.Minute,func()*retry.RetryError{
_,err:=conn.ModifyDBClusterWithContext(ctx,req)
iferr!=nil{
iftfawserr.ErrMessageContains(err,"InvalidParameterValue","IAMroleARNvalueisinvalidordoesnotincludetherequiredpermissions"){
returnretry.RetryableError(err)
}

iftfawserr.ErrMessageContains(err,docdb.ErrCodeInvalidDBClusterStateFault,"isnotcurrentlyintheavailablestate"){
returnretry.RetryableError(err)
}

iftfawserr.ErrMessageContains(err,docdb.ErrCodeInvalidDBClusterStateFault,"DBclusterisnotavailableformodification"){
returnretry.RetryableError(err)
}

returnretry.NonRetryableError(err)
}
returnnil
})
iftfresource.TimedOut(err){
_,err=conn.ModifyDBClusterWithContext(ctx,req)
}
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"modifyingDocumentDBCluster(%s):%s",d.Id(),err)
}

log.Printf("[INFO]WaitingforDocumentDBCluster(%s)tobeavailable",d.Id())
err=waitForClusterUpdate(ctx,conn,d.Id(),d.Timeout(schema.TimeoutUpdate))
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"waitingforDocumentDBCluster(%s)tobeavailable:%s",d.Id(),err)
}
}

returnappend(diags,resourceClusterRead(ctx,d,meta)...)
}
funcresourceClusterDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)
log.Printf("[DEBUG]DestroyingDocumentDBCluster(%s)",d.Id())

//Automaticallyremovefromglobalclustertobypassthiserrorondeletion:
//InvalidDBClusterStateFault:Thisclusterisapartofaglobalcluster,pleaseremoveitfromglobalclusterfirst
ifd.Get("global_cluster_identifier").(string)!=""{
input:=&docdb.RemoveFromGlobalClusterInput{
DbClusterIdentifier:aws.String(d.Get("arn").(string)),
GlobalClusterIdentifier:aws.String(d.Get("global_cluster_identifier").(string)),
}

_,err:=conn.RemoveFromGlobalClusterWithContext(ctx,input)

iferr!=nil&&!tfawserr.ErrCodeEquals(err,docdb.ErrCodeGlobalClusterNotFoundFault)&&!tfawserr.ErrMessageContains(err,"InvalidParameterValue","isnotfoundinglobalcluster"){
returnsdkdiag.AppendErrorf(diags,"removingDocumentDBCluster(%s)fromDocumentDBGlobalCluster:%s",d.Id(),err)
}
}

deleteOpts:=docdb.DeleteDBClusterInput{
DBClusterIdentifier:aws.String(d.Id()),
}

skipFinalSnapshot:=d.Get("skip_final_snapshot").(bool)
deleteOpts.SkipFinalSnapshot=aws.Bool(skipFinalSnapshot)

if!skipFinalSnapshot{
ifname,present:=d.GetOk("final_snapshot_identifier");present{
deleteOpts.FinalDBSnapshotIdentifier=aws.String(name.(string))
}else{
returnsdkdiag.AppendErrorf(diags,"DocumentDBClusterFinalSnapshotIdentifierisrequiredwhenafinalsnapshotisrequired")
}
}

err:=retry.RetryContext(ctx,5*time.Minute,func()*retry.RetryError{
_,err:=conn.DeleteDBClusterWithContext(ctx,&deleteOpts)
iferr!=nil{
iftfawserr.ErrMessageContains(err,docdb.ErrCodeInvalidDBClusterStateFault,"isnotcurrentlyintheavailablestate"){
returnretry.RetryableError(err)
}
iftfawserr.ErrMessageContains(err,docdb.ErrCodeInvalidDBClusterStateFault,"clusterisapartofaglobalcluster"){
returnretry.RetryableError(err)
}
iftfawserr.ErrCodeEquals(err,docdb.ErrCodeDBClusterNotFoundFault){
returnnil
}
returnretry.NonRetryableError(err)
}
returnnil
})
iftfresource.TimedOut(err){
_,err=conn.DeleteDBClusterWithContext(ctx,&deleteOpts)
}
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"DocumentDBClustercannotbedeleted:%s",err)
}

stateConf:=&retry.StateChangeConf{
Pending:resourceClusterDeletePendingStates,
Target:[]string{"destroyed"},
Refresh:resourceClusterStateRefreshFunc(ctx,conn,d.Id()),
Timeout:d.Timeout(schema.TimeoutDelete),
MinTimeout:10*time.Second,
Delay:ime.Second,
}

//Wait,catchinganyerrors
_,err=stateConf.WaitForStateContext(ctx)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"deletingDocumentDBCluster(%s):%s",d.Id(),err)
}

returndiags
}
funcresourceClusterStateRefreshFunc(ctxcontext.Context,conn*docdb.DocDB,dbClusterIdentifierstring)retry.StateRefreshFunc{
returnfunc()(interface{},string,error){
resp,err:=conn.DescribeDBClustersWithContext(ctx,&docdb.DescribeDBClustersInput{
DBClusterIdentifier:aws.String(dbClusterIdentifier),
})

iftfawserr.ErrCodeEquals(err,docdb.ErrCodeDBClusterNotFoundFault){
return42,"destroyed",nil
}

iferr!=nil{
returnnil,"",err
}

vardbc*docdb.DBCluster

for_,c:=rangeresp.DBClusters{
ifaws.StringValue(c.DBClusterIdentifier)==dbClusterIdentifier{
dbc=c
}
}

ifdbc==nil{
return42,"destroyed",nil
}

ifdbc.Status!=nil{
log.Printf("[DEBUG]DBClusterstatus(%s):%s",dbClusterIdentifier,*dbc.Status)
}

returndbc,aws.StringValue(dbc.Status),nil
}
}

varresourceClusterCreatePendingStates=[]string{
"creating",
"backing-up",
"modifying",
"preparing-data-migration",
"migrating",
"resetting-master-credentials",
}

varresourceClusterDeletePendingStates=[]string{
"available",
"deleting",
"backing-up",
"modifying",
}

varresourceClusterUpdatePendingStates=[]string{
"backing-up",
"modifying",
"resetting-master-credentials",
"upgrading",
}
funcwaitForClusterUpdate(ctxcontext.Context,conn*docdb.DocDB,idstring,timeouttime.Duration)error{
stateConf:=&retry.StateChangeConf{
Pending:resourceClusterUpdatePendingStates,
Target:[]string{"available"},
Refresh:resourceClusterStateRefreshFunc(ctx,conn,id),
Timeout:timeout,
MinTimeout:10*time.Second,
Delay:ime.Second,//Wait30secsbeforestarting
}
_,err:=stateConf.WaitForStateContext(ctx)
returnerr
}
funcbuildCloudWatchLogsExportConfiguration(d*schema.ResourceData)*docdb.CloudwatchLogsExportConfiguration{
oraw,nraw:=d.GetChange("enabled_cloudwatch_logs_exports")
o:=oraw.([]interface{})
n:=nraw.([]interface{})

create,disable:=diffCloudWatchLogsExportConfiguration(o,n)

return&docdb.CloudwatchLogsExportConfiguration{
EnableLogTypes:flex.ExpandStringList(create),
DisableLogTypes:flex.ExpandStringList(disable),
}
}
funcdiffCloudWatchLogsExportConfiguration(old,new[]interface{})([]interface{},[]interface{}){
add:=make([]interface{},0)
disable:=make([]interface{},0)

for_,n:=rangenew{
if_,contains:=verify.SliceContainsString(old,n.(string));!contains{
add=append(add,n)
}
}

for_,o:=rangeold{
if_,contains:=verify.SliceContainsString(new,o.(string));!contains{
disable=append(disable,o)
}
}

returnadd,disable
}
