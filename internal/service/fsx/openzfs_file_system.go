//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packagefsximport(
	"context"
	"fmt"
	"log"
	"time"	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
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
	"golang.org/x/exp/slices"
)//@SDKResource("aws_fsx_openzfs_file_system",name="OpenZFSFileSystem")
//@Tags(identifierAttribute="arn")
funcResourceOpenZFSFileSystem()*schema.Resource{
	return&schema.Resource{
CreateWithoutTimeout:resourceOpenZFSFileSystemCreate,
ReadWithoutTimeout:resourceOpenZFSFileSystemRead,
UpdateWithoutTimeout:resourceOpenZFSFileSystemUpdate,
DeleteWithoutTimeout:resourceOpenZFSFileSystemDelete,Importer:&schema.ResourceImporter{
	StateContext:schema.ImportStatePassthroughContext,
},Timeouts:&schema.ResourceTimeout{
	Create:schema.DefaultTimeout(60*time.Minute),
	Update:schema.DefaultTimeout(60*time.Minute),
	Delete:schema.DefaultTimeout(60*time.Minute),
},Schema:map[string]*schema.Schema{
	"arn":{
Type:schema.TypeString,
Computed:true,
	},
	"automatic_backup_retention_days":{
Type:schema.TypeInt,
Optional:true,
Default:0,
ValidateFunc:validation.IntBetween(0,90),
	},
	"backup_id":{
Type:schema.TypeString,
Optional:true,
ForceNew:true,
	},
	"copy_tags_to_backups":{
Type:schema.TypeBool,
Optional:true,
Default:false,
	},
	"copy_tags_to_volumes":{
Type:schema.TypeBool,
Optional:true,
Default:false,
	},
	"daily_automatic_backup_start_time":{
Type:schema.TypeString,
Optional:true,
Computed:true,
ValidateFunc:validation.All(
	validation.StringLenBetween(5,5),
	validation.StringMatch(regexache.MustCompile(`^([01]\d|2[0-3]):?([0-5]\d)$`),"mustbeintheformatHH:MM"),
),
	},
	"deployment_type":{
Type:schema.TypeString,
Required:true,
ForceNew:true,
ValidateFunc:validation.StringInSlice(fsx.OpenZFSDeploymentType_Values(),false),
	},
	"disk_iops_configuration":{
Type:schema.TypeList,
Optional:true,
Computed:true,
MaxItems:1,
Elem:&schema.Resource{
	Schema:map[string]*schema.Schema{
"iops":{
	Type:schema.TypeInt,
	Optional:true,
	Computed:true,
},
"mode":{
	Type:schema.TypeString,
	Optional:true,
	Default:fsx.DiskIopsConfigurationModeAutomatic,
	ValidateFunc:validation.StringInSlice(fsx.DiskIopsConfigurationMode_Values(),false),
},
	},
},
	},
	"dns_name":{
Type:schema.TypeString,
Computed:true,
	},
	"endpoint_ip_address_range":{
Type:schema.TypeString,
Optional:true,
Computed:true,
ForceNew:true,
	},
	"kms_key_id":{
Type:schema.TypeString,
Optional:true,
Computed:true,
ForceNew:true,
ValidateFunc:verify.ValidARN,
	},
	"network_interface_ids":{
Type:schema.TypeList,
Computed:true,
Elem:&schema.Schema{Type:schema.TypeString},
	},
	"owner_id":{
Type:schema.TypeString,
Computed:true,
	},
	"preferred_subnet_id":{
Type:schema.TypeString,
Optional:true,
ForceNew:true,
	},
	"root_volume_configuration":{
Type:schema.TypeList,
Optional:true,
Computed:true,
MaxItems:1,
Elem:&schema.Resource{
	Schema:map[string]*schema.Schema{
"copy_tags_to_snapshots":{
	Type:schema.TypeBool,
	Optional:true,
	ForceNew:true,
},
"data_compression_type":{
	Type:schema.TypeString,
	Optional:true,
	ValidateFunc:validation.StringInSlice(fsx.OpenZFSDataCompressionType_Values(),false),
},
"nfs_exports":{
	Type:schema.TypeList,
	Optional:true,
	MaxItems:1,
	Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"client_configurations":{
Type:schema.TypeSet,
Required:true,
MaxItems:25,
Elem:&schema.Resource{
	Schema:map[string]*schema.Schema{
"clients":{
	Type:schema.TypeString,
	Required:true,
	ValidateFunc:validation.All(
validation.StringLenBetween(1,128),
validation.StringMatch(regexache.MustCompile(`^[-~]{1,128}$`),"mustbeeitherIPAddressorCIDR"),
	),
},
"options":{
	Type:schema.TypeList,
	Required:true,
	MinItems:1,
	MaxItems:20,
	Elem:&schema.Schema{
Type:schema.TypeString,
ValidateFunc:validation.StringLenBetween(1,128),
	},
},
	},
},
	},
},
	},
},
"read_only":{
	Type:schema.TypeBool,
	Optional:true,
	Computed:true,
},
"record_size_kib":{
	Type:schema.TypeInt,
	Optional:true,
	Default:128,
	ValidateFunc:validation.IntInSlice([]int{4,8,16,32,64,128,256,512,1024}),
},
"user_and_group_quotas":{
	Type:schema.TypeSet,
	Optional:true,
	Computed:true,
	MaxItems:100,
	Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"id":{
Type:schema.TypeInt,
Required:true,
ValidateFunc:validation.IntBetween(0,2147483647),
	},
	"storage_capacity_quota_gib":{
Type:schema.TypeInt,
Required:true,
ValidateFunc:validation.IntBetween(0,2147483647),
	},
	"type":{
Type:schema.TypeString,
Required:true,
ValidateFunc:validation.StringInSlice(fsx.OpenZFSQuotaType_Values(),false),
	},
},
	},
},
	},
},
	},
	"root_volume_id":{
Type:schema.TypeString,
Computed:true,
	},
	"route_table_ids":{
Type:schema.TypeSet,
Optional:true,
Computed:true,
MaxItems:50,
Elem:&schema.Schema{Type:schema.TypeString},
	},
	"security_group_ids":{
Type:schema.TypeSet,
Optional:true,
ForceNew:true,
MaxItems:50,
Elem:&schema.Schema{Type:schema.TypeString},
	},
	"storage_capacity":{
Type:schema.TypeInt,
Optional:true,
ValidateFunc:validation.IntBetween(64,512*1024),
	},
	"storage_type":{
Type:schema.TypeString,
Optional:true,
ForceNew:true,
Default:fsx.StorageTypeSsd,
ValidateFunc:validation.StringInSlice(fsx.StorageType_Values(),false),
	},
	"subnet_ids":{
Type:schema.TypeList,
Required:true,
ForceNew:true,
MinItems:1,
Elem:&schema.Schema{Type:schema.TypeString},
	},
	names.AttrTags:tftags.TagsSchema(),
	names.AttrTagsAll:tftags.TagsSchemaComputed(),
	"throughput_capacity":{
Type:schema.TypeInt,
Required:true,
	},
	"vpc_id":{
Type:schema.TypeString,
Computed:true,
	},
	"weekly_maintenance_start_time":{
Type:schema.TypeString,
Optional:true,
Computed:true,
ValidateFunc:validation.All(
	validation.StringLenBetween(7,7),
	validation.StringMatch(regexache.MustCompile(`^[1-7]:([01]\d|2[0-3]):?([0-5]\d)$`),"mustbeintheformatd:HH:MM"),
),
	},
},CustomizeDiff:customdiff.All(
	verify.SetTagsDiff,
	validateDiskConfigurationIOPS,
	func(_context.Context,d*schema.ResourceDiff,metainterface{})error{
var(
	singleAZ1ThroughputCapacityValues=[]int{64,128,256,512,1024,2048,3072,4096}
	singleAZ2AndMultiAZ1ThroughputCapacityValues=[]int{160,320,640,1280,2560,3840,5120,7680,10240}
)switchdeploymentType,throughputCapacity:=d.Get("deployment_type").(string),d.Get("throughput_capacity").(int);deploymentType{
casefsx.OpenZFSDeploymentTypeSingleAz1:
	if!slices.Contains(singleAZ1ThroughputCapacityValues,throughputCapacity){
returnfmt.Errorf("%disnotavalidvaluefor`throughput_capacity`when`deployment_type`is%q.Validvalues:%v",throughputCapacity,deploymentType,singleAZ1ThroughputCapacityValues)
	}
casefsx.OpenZFSDeploymentTypeSingleAz2,fsx.OpenZFSDeploymentTypeMultiAz1:
	if!slices.Contains(singleAZ2AndMultiAZ1ThroughputCapacityValues,throughputCapacity){
returnfmt.Errorf("%disnotavalidvaluefor`throughput_capacity`when`deployment_type`is%q.Validvalues:%v",throughputCapacity,deploymentType,singleAZ2AndMultiAZ1ThroughputCapacityValues)
	}
	//default:
	//Allowvalidationtopassforunknown/newtypes.
}returnnil
	},
),
	}
}funcvalidateDiskConfigurationIOPS(_context.Context,d*schema.ResourceDiff,metainterface{})error{
	deploymentType:=d.Get("deployment_type").(string)	ifdiskConfiguration,ok:=d.GetOk("disk_iops_configuration");ok{
iflen(diskConfiguration.([]interface{}))>0{
	m:=diskConfiguration.([]interface{})[0].(map[string]interface{})	ifv,ok:=m["iops"].(int);ok{
ifdeploymentType==fsx.OpenZFSDeploymentTypeSingleAz1{
	ifv<0||v>160000{
returnfmt.Errorf("expecteddisk_iops_configuration.0.iopstobeintherange(0-160000)whendeployment_type(%s),got%d",fsx.OpenZFSDeploymentTypeSingleAz1,v)
	}
}elseifdeploymentType==fsx.OpenZFSDeploymentTypeSingleAz2{
	ifv<0||v>350000{
returnfmt.Errorf("expecteddisk_iops_configuration.0.iopstobeintherange(0-350000)whendeployment_type(%s),got%d",fsx.OpenZFSDeploymentTypeSingleAz2,v)
	}
}
	}
}
	}	returnnil
}funcresourceOpenZFSFileSystemCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).FSxConn(ctx)	inputC:=&fsx.CreateFileSystemInput{
ClientRequestToken:aws.String(id.UniqueId()),
FileSystemType:aws.String(fsx.FileSystemTypeOpenzfs),
OpenZFSConfiguration:&fsx.CreateFileSystemOpenZFSConfiguration{
	DeploymentType:aws.String(d.Get("deployment_type").(string)),
	AutomaticBackupRetentionDays:aws.Int64(int64(d.Get("automatic_backup_retention_days").(int))),
},
StorageCapacity:aws.Int64(int64(d.Get("storage_capacity").(int))),
StorageType:aws.String(d.Get("storage_type").(string)),
SubnetIds:flex.ExpandStringList(d.Get("subnet_ids").([]interface{})),
Tags:getTagsIn(ctx),
	}
	inputB:=&fsx.CreateFileSystemFromBackupInput{
ClientRequestToken:aws.String(id.UniqueId()),
OpenZFSConfiguration:&fsx.CreateFileSystemOpenZFSConfiguration{
	DeploymentType:aws.String(d.Get("deployment_type").(string)),
	AutomaticBackupRetentionDays:aws.Int64(int64(d.Get("automatic_backup_retention_days").(int))),
},
StorageType:aws.String(d.Get("storage_type").(string)),
SubnetIds:flex.ExpandStringList(d.Get("subnet_ids").([]interface{})),
Tags:getTagsIn(ctx),
	}	ifv,ok:=d.GetOk("copy_tags_to_backups");ok{
inputC.OpenZFSConfiguration.CopyTagsToBackups=aws.Bool(v.(bool))
inputB.OpenZFSConfiguration.CopyTagsToBackups=aws.Bool(v.(bool))
	}	ifv,ok:=d.GetOk("copy_tags_to_volumes");ok{
inputC.OpenZFSConfiguration.CopyTagsToVolumes=aws.Bool(v.(bool))
inputB.OpenZFSConfiguration.CopyTagsToVolumes=aws.Bool(v.(bool))
	}	ifv,ok:=d.GetOk("daily_automatic_backup_start_time");ok{
inputC.OpenZFSConfiguration.DailyAutomaticBackupStartTime=aws.String(v.(string))
inputB.OpenZFSConfiguration.DailyAutomaticBackupStartTime=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("disk_iops_configuration");ok{
inputC.OpenZFSConfiguration.DiskIopsConfiguration=expandOpenZFSFileDiskIopsConfiguration(v.([]interface{}))
inputB.OpenZFSConfiguration.DiskIopsConfiguration=expandOpenZFSFileDiskIopsConfiguration(v.([]interface{}))
	}	ifv,ok:=d.GetOk("endpoint_ip_address_range");ok{
inputC.OpenZFSConfiguration.EndpointIpAddressRange=aws.String(v.(string))
inputB.OpenZFSConfiguration.EndpointIpAddressRange=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("kms_key_id");ok{
inputC.KmsKeyId=aws.String(v.(string))
inputB.KmsKeyId=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("preferred_subnet_id");ok{
inputC.OpenZFSConfiguration.PreferredSubnetId=aws.String(v.(string))
inputB.OpenZFSConfiguration.PreferredSubnetId=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("root_volume_configuration");ok{
inputC.OpenZFSConfiguration.RootVolumeConfiguration=expandOpenZFSRootVolumeConfiguration(v.([]interface{}))
inputB.OpenZFSConfiguration.RootVolumeConfiguration=expandOpenZFSRootVolumeConfiguration(v.([]interface{}))
	}	ifv,ok:=d.GetOk("route_table_ids");ok{
inputC.OpenZFSConfiguration.RouteTableIds=flex.ExpandStringSet(v.(*schema.Set))
inputB.OpenZFSConfiguration.RouteTableIds=flex.ExpandStringSet(v.(*schema.Set))
	}	ifv,ok:=d.GetOk("security_group_ids");ok{
inputC.SecurityGroupIds=flex.ExpandStringSet(v.(*schema.Set))
inputB.SecurityGroupIds=flex.ExpandStringSet(v.(*schema.Set))
	}	ifv,ok:=d.GetOk("throughput_capacity");ok{
inputC.OpenZFSConfiguration.ThroughputCapacity=aws.Int64(int64(v.(int)))
inputB.OpenZFSConfiguration.ThroughputCapacity=aws.Int64(int64(v.(int)))
	}	ifv,ok:=d.GetOk("weekly_maintenance_start_time");ok{
inputC.OpenZFSConfiguration.WeeklyMaintenanceStartTime=aws.String(v.(string))
inputB.OpenZFSConfiguration.WeeklyMaintenanceStartTime=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("backup_id");ok{
backupID:=v.(string)
inputB.BackupId=aws.String(backupID)output,err:=conn.CreateFileSystemFromBackupWithContext(ctx,inputB)iferr!=nil{
	returnsdkdiag.AppendErrorf(diags,"creatingFSxforOpenZFSFileSystemfrombackup(%s):%s",backupID,err)
}d.SetId(aws.StringValue(output.FileSystem.FileSystemId))
	}else{
output,err:=conn.CreateFileSystemWithContext(ctx,inputC)iferr!=nil{
	returnsdkdiag.AppendErrorf(diags,"creatingFSxforOpenZFSFileSystem:%s",err)
}d.SetId(aws.StringValue(output.FileSystem.FileSystemId))
	}	if_,err:=waitFileSystemCreated(ctx,conn,d.Id(),d.Timeout(schema.TimeoutCreate));err!=nil{
returnsdkdiag.AppendErrorf(diags,"waitingforFSxforOpenZFSFileSystem(%s)create:%s",d.Id(),err)
	}	returnappend(diags,resourceOpenZFSFileSystemRead(ctx,d,meta)...)
}funcresourceOpenZFSFileSystemRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).FSxConn(ctx)	filesystem,err:=FindOpenZFSFileSystemByID(ctx,conn,d.Id())	if!d.IsNewResource()&&tfresource.NotFound(err){
log.Printf("[WARN]FSxforOpenZFSFileSystem(%s)notfound,removingfromstate",d.Id())
d.SetId("")
returndiags
	}	iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"readingFSxforOpenZFSFileSystem(%s):%s",d.Id(),err)
	}	openZFSConfig:=filesystem.OpenZFSConfiguration	d.Set("arn",filesystem.ResourceARN)
	d.Set("automatic_backup_retention_days",openZFSConfig.AutomaticBackupRetentionDays)
	d.Set("copy_tags_to_backups",openZFSConfig.CopyTagsToBackups)
	d.Set("copy_tags_to_volumes",openZFSConfig.CopyTagsToVolumes)
	d.Set("daily_automatic_backup_start_time",openZFSConfig.DailyAutomaticBackupStartTime)
	d.Set("deployment_type",openZFSConfig.DeploymentType)
	iferr:=d.Set("disk_iops_configuration",flattenOpenZFSFileDiskIopsConfiguration(openZFSConfig.DiskIopsConfiguration));err!=nil{
returnsdkdiag.AppendErrorf(diags,"settingdisk_iops_configuration:%s",err)
	}
	d.Set("dns_name",filesystem.DNSName)
	d.Set("endpoint_ip_address_range",openZFSConfig.EndpointIpAddressRange)
	d.Set("kms_key_id",filesystem.KmsKeyId)
	d.Set("network_interface_ids",aws.StringValueSlice(filesystem.NetworkInterfaceIds))
	d.Set("owner_id",filesystem.OwnerId)
	d.Set("preferred_subnet_id",openZFSConfig.PreferredSubnetId)
	rootVolumeID:=aws.StringValue(openZFSConfig.RootVolumeId)
	d.Set("root_volume_id",rootVolumeID)
	d.Set("route_table_ids",aws.StringValueSlice(openZFSConfig.RouteTableIds))
	d.Set("storage_capacity",filesystem.StorageCapacity)
	d.Set("storage_type",filesystem.StorageType)
	d.Set("subnet_ids",aws.StringValueSlice(filesystem.SubnetIds))
	d.Set("throughput_capacity",openZFSConfig.ThroughputCapacity)
	d.Set("vpc_id",filesystem.VpcId)
	d.Set("weekly_maintenance_start_time",openZFSConfig.WeeklyMaintenanceStartTime)	setTagsOut(ctx,filesystem.Tags)	rootVolume,err:=FindVolumeByID(ctx,conn,rootVolumeID)	iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"readingFSxforOpenZFSFileSystem(%s)rootvolume(%s):%s",d.Id(),rootVolumeID,err)
	}	iferr:=d.Set("root_volume_configuration",flattenOpenZFSRootVolumeConfiguration(rootVolume));err!=nil{
returnsdkdiag.AppendErrorf(diags,"settingroot_volume_configuration:%s",err)
	}	returndiags
}funcresourceOpenZFSFileSystemUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).FSxConn(ctx)	ifd.HasChangesExcept("tags","tags_all"){
input:=&fsx.UpdateFileSystemInput{
	ClientRequestToken:aws.String(id.UniqueId()),
	FileSystemId:aws.String(d.Id()),
	OpenZFSConfiguration:&fsx.UpdateFileSystemOpenZFSConfiguration{},
}ifd.HasChange("automatic_backup_retention_days"){
	input.OpenZFSConfiguration.AutomaticBackupRetentionDays=aws.Int64(int64(d.Get("automatic_backup_retention_days").(int)))
}ifd.HasChange("copy_tags_to_backups"){
	input.OpenZFSConfiguration.CopyTagsToBackups=aws.Bool(d.Get("copy_tags_to_backups").(bool))
}ifd.HasChange("copy_tags_to_volumes"){
	input.OpenZFSConfiguration.CopyTagsToVolumes=aws.Bool(d.Get("copy_tags_to_volumes").(bool))
}ifd.HasChange("daily_automatic_backup_start_time"){
	input.OpenZFSConfiguration.DailyAutomaticBackupStartTime=aws.String(d.Get("daily_automatic_backup_start_time").(string))
}ifd.HasChange("disk_iops_configuration"){
	input.OpenZFSConfiguration.DiskIopsConfiguration=expandOpenZFSFileDiskIopsConfiguration(d.Get("disk_iops_configuration").([]interface{}))
}ifd.HasChange("route_table_ids"){
	o,n:=d.GetChange("route_table_ids")
	os,ns:=o.(*schema.Set),n.(*schema.Set)
	add,del:=flex.ExpandStringValueSet(ns.Difference(os)),flex.ExpandStringValueSet(os.Difference(ns))	iflen(add)>0{
input.OpenZFSConfiguration.AddRouteTableIds=aws.StringSlice(add)
	}
	iflen(del)>0{
input.OpenZFSConfiguration.RemoveRouteTableIds=aws.StringSlice(del)
	}
}ifd.HasChange("storage_capacity"){
	input.StorageCapacity=aws.Int64(int64(d.Get("storage_capacity").(int)))
}ifd.HasChange("throughput_capacity"){
	input.OpenZFSConfiguration.ThroughputCapacity=aws.Int64(int64(d.Get("throughput_capacity").(int)))
}ifd.HasChange("weekly_maintenance_start_time"){
	input.OpenZFSConfiguration.WeeklyMaintenanceStartTime=aws.String(d.Get("weekly_maintenance_start_time").(string))
}startTime:=time.Now()
_,err:=conn.UpdateFileSystemWithContext(ctx,input)iferr!=nil{
	returnsdkdiag.AppendErrorf(diags,"updatingFSxforOpenZFSFileSystem(%s):%s",d.Id(),err)
}if_,err:=waitFileSystemUpdated(ctx,conn,d.Id(),startTime,d.Timeout(schema.TimeoutUpdate));err!=nil{
	returnsdkdiag.AppendErrorf(diags,"waitingforFSxforOpenZFSFileSystem(%s)update:%s",d.Id(),err)
}if_,err:=waitAdministrativeActionCompleted(ctx,conn,d.Id(),fsx.AdministrativeActionTypeFileSystemUpdate,d.Timeout(schema.TimeoutUpdate));err!=nil{
	returnsdkdiag.AppendErrorf(diags,"waitingforFSxforOpenZFSFileSystem(%s)administrativeaction(%s)complete:%s",d.Id(),fsx.AdministrativeActionTypeFileSystemUpdate,err)
}ifd.HasChange("root_volume_configuration"){
	rootVolumeID:=d.Get("root_volume_id").(string)
	input:=&fsx.UpdateVolumeInput{
ClientRequestToken:aws.String(id.UniqueId()),
OpenZFSConfiguration:expandOpenZFSUpdateRootVolumeConfiguration(d.Get("root_volume_configuration").([]interface{})),
VolumeId:aws.String(rootVolumeID),
	}	_,err:=conn.UpdateVolumeWithContext(ctx,input)	iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"updatingFSxforOpenZFSRootVolume(%s):%s",rootVolumeID,err)
	}	if_,err:=waitVolumeUpdated(ctx,conn,rootVolumeID,d.Timeout(schema.TimeoutUpdate));err!=nil{
returnsdkdiag.AppendErrorf(diags,"waitingforFSxforOpenZFSRootVolume(%s)update:%s",rootVolumeID,err)
	}	if_,err:=waitAdministrativeActionCompleted(ctx,conn,d.Id(),fsx.AdministrativeActionTypeVolumeUpdate,d.Timeout(schema.TimeoutUpdate));err!=nil{
returnsdkdiag.AppendErrorf(diags,"waitingforFSxforOpenZFSFileSystem(%s)administrativeaction(%s)complete:%s",d.Id(),fsx.AdministrativeActionTypeVolumeUpdate,err)
	}
}
	}	returnappend(diags,resourceOpenZFSFileSystemRead(ctx,d,meta)...)
}funcresourceOpenZFSFileSystemDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).FSxConn(ctx)	log.Printf("[DEBUG]DeletingFSxforOpenZFSFileSystem:%s",d.Id())
	_,err:=conn.DeleteFileSystemWithContext(ctx,&fsx.DeleteFileSystemInput{
FileSystemId:aws.String(d.Id()),
	})	iftfawserr.ErrCodeEquals(err,fsx.ErrCodeFileSystemNotFound){
returndiags
	}	iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"deletingFSxforOpenZFSFileSystem(%s):%s",d.Id(),err)
	}	if_,err:=waitFileSystemDeleted(ctx,conn,d.Id(),d.Timeout(schema.TimeoutDelete));err!=nil{
returnsdkdiag.AppendErrorf(diags,"waitingforFSxforOpenZFSFileSystem(%s)delete:%s",d.Id(),err)
	}	returndiags
}funcexpandOpenZFSFileDiskIopsConfiguration(cfg[]interface{})*fsx.DiskIopsConfiguration{
	iflen(cfg)<1{
returnnil
	}	conf:=cfg[0].(map[string]interface{})	out:=fsx.DiskIopsConfiguration{}	ifv,ok:=conf["mode"].(string);ok&&len(v)>0{
out.Mode=aws.String(v)
	}	ifv,ok:=conf["iops"].(int);ok{
out.Iops=aws.Int64(int64(v))
	}	return&out
}funcexpandOpenZFSRootVolumeConfiguration(cfg[]interface{})*fsx.OpenZFSCreateRootVolumeConfiguration{
	iflen(cfg)<1{
returnnil
	}	conf:=cfg[0].(map[string]interface{})	out:=fsx.OpenZFSCreateRootVolumeConfiguration{}	ifv,ok:=conf["copy_tags_to_snapshots"].(bool);ok{
out.CopyTagsToSnapshots=aws.Bool(v)
	}	ifv,ok:=conf["data_compression_type"].(string);ok{
out.DataCompressionType=aws.String(v)
	}	ifv,ok:=conf["read_only"].(bool);ok{
out.ReadOnly=aws.Bool(v)
	}	ifv,ok:=conf["record_size_kib"].(int);ok{
out.RecordSizeKiB=aws.Int64(int64(v))
	}	ifv,ok:=conf["user_and_group_quotas"];ok{
out.UserAndGroupQuotas=expandOpenZFSUserAndGroupQuotas(v.(*schema.Set).List())
	}	ifv,ok:=conf["nfs_exports"].([]interface{});ok{
out.NfsExports=expandOpenZFSNFSExports(v)
	}	return&out
}funcexpandOpenZFSUpdateRootVolumeConfiguration(cfg[]interface{})*fsx.UpdateOpenZFSVolumeConfiguration{
	iflen(cfg)<1{
returnnil
	}	conf:=cfg[0].(map[string]interface{})	out:=fsx.UpdateOpenZFSVolumeConfiguration{}	ifv,ok:=conf["data_compression_type"].(string);ok{
out.DataCompressionType=aws.String(v)
	}	ifv,ok:=conf["read_only"].(bool);ok{
out.ReadOnly=aws.Bool(v)
	}	ifv,ok:=conf["record_size_kib"].(int);ok{
out.RecordSizeKiB=aws.Int64(int64(v))
	}	ifv,ok:=conf["user_and_group_quotas"];ok{
out.UserAndGroupQuotas=expandOpenZFSUserAndGroupQuotas(v.(*schema.Set).List())
	}	ifv,ok:=conf["nfs_exports"].([]interface{});ok{
out.NfsExports=expandOpenZFSNFSExports(v)
	}	return&out
}funcexpandOpenZFSUserAndGroupQuotas(cfg[]interface{})[]*fsx.OpenZFSUserOrGroupQuota{
	quotas:=[]*fsx.OpenZFSUserOrGroupQuota{}	for_,quota:=rangecfg{
expandedQuota:=expandOpenZFSUserAndGroupQuota(quota.(map[string]interface{}))
ifexpandedQuota!=nil{
	quotas=append(quotas,expandedQuota)
}
	}	returnquotas
}funcexpandOpenZFSUserAndGroupQuota(confmap[string]interface{})*fsx.OpenZFSUserOrGroupQuota{
	iflen(conf)<1{
returnnil
	}	out:=fsx.OpenZFSUserOrGroupQuota{}	ifv,ok:=conf["id"].(int);ok{
out.Id=aws.Int64(int64(v))
	}	ifv,ok:=conf["storage_capacity_quota_gib"].(int);ok{
out.StorageCapacityQuotaGiB=aws.Int64(int64(v))
	}	ifv,ok:=conf["type"].(string);ok{
out.Type=aws.String(v)
	}	return&out
}funcexpandOpenZFSNFSExports(cfg[]interface{})[]*fsx.OpenZFSNfsExport{
	exports:=[]*fsx.OpenZFSNfsExport{}	for_,export:=rangecfg{
expandedExport:=expandOpenZFSNFSExport(export.(map[string]interface{}))
ifexpandedExport!=nil{
	exports=append(exports,expandedExport)
}
	}	returnexports
}funcexpandOpenZFSNFSExport(cfgmap[string]interface{})*fsx.OpenZFSNfsExport{
	out:=fsx.OpenZFSNfsExport{}	ifv,ok:=cfg["client_configurations"];ok{
out.ClientConfigurations=expandOpenZFSClinetConfigurations(v.(*schema.Set).List())
	}	return&out
}funcexpandOpenZFSClinetConfigurations(cfg[]interface{})[]*fsx.OpenZFSClientConfiguration{
	configurations:=[]*fsx.OpenZFSClientConfiguration{}	for_,configuration:=rangecfg{
expandedConfiguration:=expandOpenZFSClientConfiguration(configuration.(map[string]interface{}))
ifexpandedConfiguration!=nil{
	configurations=append(configurations,expandedConfiguration)
}
	}	returnconfigurations
}funcexpandOpenZFSClientConfiguration(confmap[string]interface{})*fsx.OpenZFSClientConfiguration{
	out:=fsx.OpenZFSClientConfiguration{}	ifv,ok:=conf["clients"].(string);ok&&len(v)>0{
out.Clients=aws.String(v)
	}	ifv,ok:=conf["options"].([]interface{});ok{
out.Options=flex.ExpandStringList(v)
	}	return&out
}funcflattenOpenZFSFileDiskIopsConfiguration(rs*fsx.DiskIopsConfiguration)[]interface{}{
	ifrs==nil{
return[]interface{}{}
	}	m:=make(map[string]interface{})
	ifrs.Mode!=nil{
m["mode"]=aws.StringValue(rs.Mode)
	}
	ifrs.Iops!=nil{
m["iops"]=aws.Int64Value(rs.Iops)
	}	return[]interface{}{m}
}funcflattenOpenZFSRootVolumeConfiguration(rs*fsx.Volume)[]interface{}{
	ifrs==nil{
return[]interface{}{}
	}	m:=make(map[string]interface{})
	ifrs.OpenZFSConfiguration.CopyTagsToSnapshots!=nil{
m["copy_tags_to_snapshots"]=aws.BoolValue(rs.OpenZFSConfiguration.CopyTagsToSnapshots)
	}
	ifrs.OpenZFSConfiguration.DataCompressionType!=nil{
m["data_compression_type"]=aws.StringValue(rs.OpenZFSConfiguration.DataCompressionType)
	}
	ifrs.OpenZFSConfiguration.NfsExports!=nil{
m["nfs_exports"]=flattenOpenZFSFileNFSExports(rs.OpenZFSConfiguration.NfsExports)
	}
	ifrs.OpenZFSConfiguration.ReadOnly!=nil{
m["read_only"]=aws.BoolValue(rs.OpenZFSConfiguration.ReadOnly)
	}
	ifrs.OpenZFSConfiguration.RecordSizeKiB!=nil{
m["record_size_kib"]=aws.Int64Value(rs.OpenZFSConfiguration.RecordSizeKiB)
	}
	ifrs.OpenZFSConfiguration.UserAndGroupQuotas!=nil{
m["user_and_group_quotas"]=flattenOpenZFSFileUserAndGroupQuotas(rs.OpenZFSConfiguration.UserAndGroupQuotas)
	}	return[]interface{}{m}
}funcflattenOpenZFSFileNFSExports(rs[]*fsx.OpenZFSNfsExport)[]map[string]interface{}{
	exports:=make([]map[string]interface{},0)	for_,export:=rangers{
ifexport!=nil{
	cfg:=make(map[string]interface{})
	cfg["client_configurations"]=flattenOpenZFSClientConfigurations(export.ClientConfigurations)
	exports=append(exports,cfg)
}
	}	iflen(exports)>0{
returnexports
	}	returnnil
}funcflattenOpenZFSClientConfigurations(rs[]*fsx.OpenZFSClientConfiguration)[]map[string]interface{}{
	configurations:=make([]map[string]interface{},0)	for_,configuration:=rangers{
ifconfiguration!=nil{
	cfg:=make(map[string]interface{})
	cfg["clients"]=aws.StringValue(configuration.Clients)
	cfg["options"]=flex.FlattenStringList(configuration.Options)
	configurations=append(configurations,cfg)
}
	}	iflen(configurations)>0{
returnconfigurations
	}	returnnil
}funcflattenOpenZFSFileUserAndGroupQuotas(rs[]*fsx.OpenZFSUserOrGroupQuota)[]map[string]interface{}{
	quotas:=make([]map[string]interface{},0)	for_,quota:=rangers{
ifquota!=nil{
	cfg:=make(map[string]interface{})
	cfg["id"]=aws.Int64Value(quota.Id)
	cfg["storage_capacity_quota_gib"]=aws.Int64Value(quota.StorageCapacityQuotaGiB)
	cfg["type"]=aws.StringValue(quota.Type)
	quotas=append(quotas,cfg)
}
	}	iflen(quotas)>0{
returnquotas
	}	returnnil
}funcFindOpenZFSFileSystemByID(ctxcontext.Context,conn*fsx.FSx,idstring)(*fsx.FileSystem,error){
	output,err:=findFileSystemByIDAndType(ctx,conn,id,fsx.FileSystemTypeOpenzfs)	iferr!=nil{
returnnil,err
	}	ifoutput.OpenZFSConfiguration==nil{
returnnil,tfresource.NewEmptyResultError(nil)
	}	returnoutput,nil
}funcFindVolumeByID(ctxcontext.Context,conn*fsx.FSx,idstring)(*fsx.Volume,error){
	input:=&fsx.DescribeVolumesInput{
VolumeIds:aws.StringSlice([]string{id}),
	}	returnfindVolume(ctx,conn,input)
}funcfindVolume(ctxcontext.Context,conn*fsx.FSx,input*fsx.DescribeVolumesInput)(*fsx.Volume,error){
	output,err:=findVolumes(ctx,conn,input)	iferr!=nil{
returnnil,err
	}	returntfresource.AssertSinglePtrResult(output)
}funcfindVolumes(ctxcontext.Context,conn*fsx.FSx,input*fsx.DescribeVolumesInput)([]*fsx.Volume,error){
	varoutput[]*fsx.Volume	err:=conn.DescribeVolumesPagesWithContext(ctx,input,func(page*fsx.DescribeVolumesOutput,lastPagebool)bool{
ifpage==nil{
	return!lastPage
}for_,v:=rangepage.Volumes{
	ifv!=nil{
output=append(output,v)
	}
}return!lastPage
	})	iftfawserr.ErrCodeEquals(err,fsx.ErrCodeVolumeNotFound){
returnnil,&retry.NotFoundError{
	LastError:err,
	LastRequest:input,
}
	}	iferr!=nil{
returnnil,err
	}	returnoutput,nil
}
