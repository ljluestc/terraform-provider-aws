//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packagemedialiveimport(
	"context"
	"errors"
	"fmt"
	"log"
	"time"	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)//@SDKResource("aws_medialive_channel",name="Channel")
//@Tags(identifierAttribute="arn")
funcResourceChannel()*schema.Resource{
	return&schema.Resource{
CreateWithoutTimeout:resourceChannelCreate,
adWithoutTimeout:resourceChannelRead,
dateWithoutTimeout:resourceChannelUpdate,
leteWithoutTimeout:resourceChannelDelete,Imrter:&schema.ResourceImporter{
StateContext:schema.ImportStatePassthroughContext,
Tiouts:&schema.ResourceTimeout{
Create:schema.DefaultTimeout(15*time.Minute),
Update:schema.DefaultTimeout(15*time.Minute),
Delete:schema.DefaultTimeout(15*time.Minute),
ScmaFunc:func()map[string]*schema.Schema{
returnmap[string]*schema.Schema{
"arn":{
Type:schema.TypeString,
Computed:true,
},
"cdi_input_specification":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"resolution":{
pe:schema.TypeString,
quired:true,
lidateDiagFunc:enum.Validate[types.CdiInputResolution](),
	},
},
},
},
"channel_class":{
Type:schema.TypeString,
Required:true,
ForceNew:true,
ValidateDiagFunc:enum.Validate[types.ChannelClass](),
},
"channel_id":{
Type:schema.TypeString,
Computed:true,
},
"destinations":{
Type:schema.TypeSet,
Required:true,
MinItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"id":{
pe:schema.TypeString,
quired:true,
	},
	"media_package_settings":{
pe:schema.TypeSet,
tional:true,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"channel_id":{
Type:schema.TypeString,
Required:true,
},
},	},
	"multiplex_settings":{
pe:schema.TypeList,
tional:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"multiplex_id":{
Type:schema.TypeString,
Required:true,
},
"program_name":{
Type:schema.TypeString,
Required:true,
},
},	},
	"settings":{
pe:schema.TypeSet,
tional:true,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"password_param":{
Type:schema.TypeString,
Optional:true,
},
"stream_name":{
Type:schema.TypeString,
Optional:true,
},
"url":{
Type:schema.TypeString,
Optional:true,
},
"username":{
Type:schema.TypeString,
Optional:true,
},
},	},
},
},
},
"encoder_settings":func()*schema.Schema{
returnchannelEncoderSettingsSchema()
}(),
"input_attachments":{
Type:schema.TypeSet,
Required:true,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"automatic_input_failover_settings":{
pe:schema.TypeList,
tional:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"secondary_input_id":{
Type:schema.TypeString,
Required:true,
},
"error_clear_time_msec":{
Type:schema.TypeInt,
Optional:true,
},
"failover_condition":{
Type:schema.TypeSet,
Optional:true,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"failover_condition_settings":{
pe:schema.TypeList,
tional:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"audio_silence_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"audio_selector_name":{
pe:schema.TypeString,
quired:true,
	},
	"audio_silence_threshold_msec":{
pe:schema.TypeInt,
tional:true,
	},
},
},
},
"input_loss_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"input_loss_threshold_msec":{
pe:schema.TypeInt,
tional:true,
	},
},
},
},
"video_black_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"black_detect_threshold":{
pe:schema.TypeFloat,
tional:true,
	},
	"video_black_threshold_msec":{
pe:schema.TypeInt,
tional:true,
	},
},
},
},
},	},
},
},
},
"input_preference":{
Type:schema.TypeString,
Optional:true,
ValidateDiagFunc:enum.Validate[types.InputPreference](),
},
},	},
	"input_attachment_name":{
pe:schema.TypeString,
quired:true,
	},
	"input_id":{
pe:schema.TypeString,
quired:true,
	},
	"input_settings":{
pe:schema.TypeList,
tional:true,
mputed:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"audio_selector":{
Type:schema.TypeList,
Optional:true,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"name":{
pe:schema.TypeString,
quired:true,
	},
	"selector_settings":{
pe:schema.TypeList,
tional:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"audio_hls_rendition_selection":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"group_id":{
pe:schema.TypeString,
quired:true,
	},
	"name":{
pe:schema.TypeString,
quired:true,
	},
},
},
},
"audio_language_selection":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"language_code":{
pe:schema.TypeString,
quired:true,
	},
	"language_selection_policy":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.AudioLanguageSelectionPolicy](),
	},
},
},
},
"audio_pid_selection":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"pid":{
pe:schema.TypeInt,
quired:true,
	},
},
},
},
"audio_track_selection":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"dolby_e_decode":{
pe:schema.TypeList,
tional:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"program_selection":{
Type:schema.TypeString,
Required:true,
ValidateDiagFunc:enum.Validate[types.DolbyEProgramSelection](),
},
},	},
	"tracks":{
pe:schema.TypeSet,
quired:true,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"track":{
Type:schema.TypeInt,
Required:true,
},
},	},
},
},
},
},	},
},
},
},
"caption_selector":{
Type:schema.TypeList,
Optional:true,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"name":{
pe:schema.TypeString,
quired:true,
	},
	"language_code":{
pe:schema.TypeString,
tional:true,
	},
	"selector_settings":{
pe:schema.TypeList,
tional:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"ancillary_source_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"source_ancillary_channel_number":{
pe:schema.TypeInt,
tional:true,
	},
},
},
},
"arib_source_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{},//noexportedelementsinthislist
},
},
"dvb_sub_source_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"ocr_language":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.DvbSubOcrLanguage](),
	},
	"pid":{
pe:schema.TypeInt,
tional:true,
lidateFunc:validation.IntAtLeast(1),
	},
},
},
},
"embedded_source_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"convert_608_to_708":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.EmbeddedConvert608To708](),
	},
	"scte20_detection":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.EmbeddedScte20Detection](),
	},
	"source_608_channel_number":{
pe:schema.TypeInt,
tional:true,
	},
},
},
},
"scte20_source_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"convert_608_to_708":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.Scte20Convert608To708](),
	},
	"source_608_channel_number":{
pe:schema.TypeInt,
tional:true,
	},
},
},
},
"scte27_source_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"ocr_language":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.Scte27OcrLanguage](),
	},
	"pid":{
pe:schema.TypeInt,
tional:true,
	},
},
},
},
"teletext_source_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"output_rectangle":{
pe:schema.TypeList,
tional:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"height":{
Type:schema.TypeFloat,
Required:true,
},
"left_offset":{
Type:schema.TypeFloat,
Required:true,
},
"top_offset":{
Type:schema.TypeFloat,
Required:true,
},
"width":{
Type:schema.TypeFloat,
Required:true,
},
},	},
	"page_number":{
pe:schema.TypeString,
tional:true,
	},
},
},
},
},	},
},
},
},
"deblock_filter":{
Type:schema.TypeString,
Optional:true,
ValidateDiagFunc:enum.Validate[types.InputDeblockFilter](),
},
"denoise_filter":{
Type:schema.TypeString,
Optional:true,
ValidateDiagFunc:enum.Validate[types.InputDenoiseFilter](),
},
"filter_strength":{
Type:schema.TypeInt,
Optional:true,
ValidateDiagFunc:validation.ToDiagFunc(validation.IntBetween(1,5)),
},
"input_filter":{
Type:schema.TypeString,
Optional:true,
Computed:true,
ValidateDiagFunc:enum.Validate[types.InputFilter](),
},
"network_input_settings":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"hls_input_settings":{
pe:schema.TypeList,
tional:true,
xItems:1,
em:&schema.Resource{
Schema:map[string]*schema.Schema{
"bandwidth":{
Type:schema.TypeInt,
Optional:true,
},
"buffer_segments":{
Type:schema.TypeInt,
Optional:true,
},
"retries":{
Type:schema.TypeInt,
Optional:true,
},
"retry_interval":{
Type:schema.TypeInt,
Optional:true,
},
"scte35_source":{
Type:schema.TypeString,
Optional:true,
ValidateDiagFunc:enum.Validate[types.HlsScte35SourceType](),
},
},	},
	"server_validation":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.NetworkInputServerValidation](),
	},
},
},
},
"scte35_pid":{
Type:schema.TypeInt,
Optional:true,
},
"smpte2038_data_preference":{
Type:schema.TypeString,
Optional:true,
ValidateDiagFunc:enum.Validate[types.Smpte2038DataPreference](),
},
"source_end_behavior":{
Type:schema.TypeString,
Optional:true,
ValidateDiagFunc:enum.Validate[types.InputSourceEndBehavior](),
},
"video_selector":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"color_space":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.VideoSelectorColorSpace](),
	},
	//TODOimplementcolor_space_settings
	"color_space_usage":{
pe:schema.TypeString,
tional:true,
lidateDiagFunc:enum.Validate[types.VideoSelectorColorSpaceUsage](),
	},
	//TODOimplementselector_settings
},
},
},
},	},
},
},
},
"input_specification":{
Type:schema.TypeList,
Required:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"codec":{
pe:schema.TypeString,
quired:true,
lidateDiagFunc:enum.Validate[types.InputCodec](),
	},
	"maximum_bitrate":{
pe:schema.TypeString,
quired:true,
lidateDiagFunc:enum.Validate[types.InputMaximumBitrate](),
	},
	"input_resolution":{
pe:schema.TypeString,
quired:true,
lidateDiagFunc:enum.Validate[types.InputResolution](),
	},
},
},
},
"log_level":{
Type:schema.TypeString,
Optional:true,
Computed:true,
ValidateDiagFunc:enum.Validate[types.LogLevel](),
},
"maintenance":{
Type:schema.TypeList,
Optional:true,
Computed:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"maintenance_day":{
pe:schema.TypeString,
quired:true,
lidateDiagFunc:enum.Validate[types.MaintenanceDay](),
	},
	"maintenance_start_time":{
pe:schema.TypeString,
quired:true,
	},
},
},
},
"name":{
Type:schema.TypeString,
Required:true,
},
"role_arn":{
Type:schema.TypeString,
Optional:true,
ValidateDiagFunc:validation.ToDiagFunc(verify.ValidARN),
},
"start_channel":{
Type:schema.TypeBool,
Optional:true,
Default:false,
},
"vpc":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
ForceNew:true,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
	"availability_zones":{
pe:schema.TypeList,
mputed:true,
em:&schema.Schema{Type:schema.TypeString},
	},
	"public_address_allocation_ids":{
pe:schema.TypeList,
quired:true,
em:&schema.Schema{Type:schema.TypeString},
	},
	"security_group_ids":{
pe:schema.TypeList,
tional:true,
mputed:true,
xItems:5,
em:&schema.Schema{Type:schema.TypeString},
	},
	"subnet_ids":{
pe:schema.TypeList,
quired:true,
em:&schema.Schema{Type:schema.TypeString},
	},
},
},
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll:tftags.TagsSchemaComputed(),
}
CuomizeDiff:verify.SetTagsDiff,
	}
}const(
	ResNameChannel="Channel"
)funcresourceChannelCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).MediaLiveClient(ctx)	in:=&medialive.CreateChannelInput{
me:ring(d.Get("name").(string)),
questId:aws.String(id.UniqueId()),
gs:sIn(ctx),
	}	ifv,ok:=d.GetOk("cdi_input_specification");ok&&len(v.([]interface{}))>0{
.CdiInputSpecification=expandChannelCdiInputSpecification(v.([]interface{}))
	}
	ifv,ok:=d.GetOk("channel_class");ok{
.ChannelClass=types.ChannelClass(v.(string))
	}
	ifv,ok:=d.GetOk("destinations");ok&&v.(*schema.Set).Len()>0{
.Destinations=expandChannelDestinations(v.(*schema.Set).List())
	}
	ifv,ok:=d.GetOk("encoder_settings");ok&&len(v.([]interface{}))>0{
.EncoderSettings=expandChannelEncoderSettings(v.([]interface{}))
	}
	ifv,ok:=d.GetOk("input_attachments");ok&&v.(*schema.Set).Len()>0{
.InputAttachments=expandChannelInputAttachments(v.(*schema.Set).List())
	}
	ifv,ok:=d.GetOk("input_specification");ok&&len(v.([]interface{}))>0{
.InputSpecification=expandChannelInputSpecification(v.([]interface{}))
	}
	ifv,ok:=d.GetOk("maintenance");ok&&len(v.([]interface{}))>0{
.Maintenance=expandChannelMaintenanceCreate(v.([]interface{}))
	}
	ifv,ok:=d.GetOk("role_arn");ok{
.RoleArn=aws.String(v.(string))
	}
	ifv,ok:=d.GetOk("vpc");ok&&len(v.([]interface{}))>0{
.Vpc=expandChannelVPC(v.([]interface{}))
	}	out,err:=conn.CreateChannel(ctx,in)
	iferr!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionCreating,ResNameChannel,d.Get("name").(string),err)
	}	ifout==nil||out.Channel==nil{
turncreate.DiagError(names.MediaLive,create.ErrActionCreating,ResNameChannel,d.Get("name").(string),errors.New("emptyoutput"))
	}	d.SetId(aws.ToString(out.Channel.Id))	if_,err:=waitChannelCreated(ctx,conn,d.Id(),d.Timeout(schema.TimeoutCreate));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionWaitingForCreation,ResNameChannel,d.Id(),err)
	}	ifd.Get("start_channel").(bool){
err:=startChannel(ctx,conn,d.Timeout(schema.TimeoutCreate),d.Id());err!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionCreating,ResNameChannel,d.Get("name").(string),err)	}	returnresourceChannelRead(ctx,d,meta)
}funcresourceChannelRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).MediaLiveClient(ctx)	out,err:=FindChannelByID(ctx,conn,d.Id())	if!d.IsNewResource()&&tfresource.NotFound(err){
g.Printf("[WARN]MediaLiveChannel(%s)notfound,removingfromstate",d.Id())
SetId("")
turnnil
	}	iferr!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionReading,ResNameChannel,d.Id(),err)
	}	d.Set("arn",out.Arn)
	d.Set("name",out.Name)
	d.Set("channel_class",out.ChannelClass)
	d.Set("channel_id",out.Id)
	d.Set("log_level",out.LogLevel)
	d.Set("role_arn",out.RoleArn)	iferr:=d.Set("cdi_input_specification",flattenChannelCdiInputSpecification(out.CdiInputSpecification));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("input_attachments",flattenChannelInputAttachments(out.InputAttachments));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("destinations",flattenChannelDestinations(out.Destinations));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("encoder_settings",flattenChannelEncoderSettings(out.EncoderSettings));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("input_specification",flattenChannelInputSpecification(out.InputSpecification));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("maintenance",flattenChannelMaintenance(out.Maintenance));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("vpc",flattenChannelVPC(out.Vpc));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}	returnnil
}funcresourceChannelUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).MediaLiveClient(ctx)	ifd.HasChangesExcept("tags","tags_all","start_channel"){
:=&medialive.UpdateChannelInput{
ChannelId:aws.String(d.Id()),
fHasChange("name"){
in.Name=aws.String(d.Get("name").(string))
fHasChange("cdi_input_specification"){
in.CdiInputSpecification=expandChannelCdiInputSpecification(d.Get("cdi_input_specification").([]interface{}))
fHasChange("destinations"){
in.Destinations=expandChannelDestinations(d.Get("destinations").(*schema.Set).List())
fHasChange("encoder_settings"){
in.EncoderSettings=expandChannelEncoderSettings(d.Get("encoder_settings").([]interface{}))
fHasChange("input_attachments"){
in.InputAttachments=expandChannelInputAttachments(d.Get("input_attachments").(*schema.Set).List())
fHasChange("input_specification"){
in.InputSpecification=expandChannelInputSpecification(d.Get("input_specification").([]interface{}))
fHasChange("log_level"){
in.LogLevel=types.LogLevel(d.Get("log_level").(string))
fHasChange("maintenance"){
in.Maintenance=expandChannelMaintenanceUpdate(d.Get("maintenance").([]interface{}))
fHasChange("role_arn"){
in.RoleArn=aws.String(d.Get("role_arn").(string))
hnel,err:=FindChannelByID(ctx,conn,d.Id())ifer=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
fannel.State==types.ChannelStateRunning{
iferr:=stopChannel(ctx,conn,d.Timeout(schema.TimeoutUpdate),d.Id());err!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
}
uerr:=conn.UpdateChannel(ctx,in)
err!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
ferr:=waitChannelUpdated(ctx,conn,aws.ToString(out.Channel.Id),d.Timeout(schema.TimeoutUpdate));err!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionWaitingForUpdate,ResNameChannel,d.Id(),err)	}	ifd.Get("start_channel").(bool){
err:=startChannel(ctx,conn,d.Timeout(schema.TimeoutUpdate),d.Id());err!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Get("name").(string),err)	}	ifd.HasChange("start_channel"){
annel,err:=FindChannelByID(ctx,conn,d.Id())ifr!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
wchd.Get("start_channel").(bool){
setrue:
ifchannel.State==types.ChannelStateIdle{
iferr:=startChannel(ctx,conn,d.Timeout(schema.TimeoutUpdate),d.Id());err!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
}
}
fault:
ifchannel.State==types.ChannelStateRunning{
iferr:=stopChannel(ctx,conn,d.Timeout(schema.TimeoutUpdate),d.Id());err!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
}
}	}	returnresourceChannelRead(ctx,d,meta)
}funcresourceChannelDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).MediaLiveClient(ctx)	log.Printf("[INFO]DeletingMediaLiveChannel%s",d.Id())	channel,err:=FindChannelByID(ctx,conn,d.Id())	iferr!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionDeleting,ResNameChannel,d.Id(),err)
	}	ifchannel.State==types.ChannelStateRunning{
err:=stopChannel(ctx,conn,d.Timeout(schema.TimeoutDelete),d.Id());err!=nil{
returncreate.DiagError(names.MediaLive,create.ErrActionDeleting,ResNameChannel,d.Id(),err)	}	_,err=conn.DeleteChannel(ctx,&medialive.DeleteChannelInput{
annelId:aws.String(d.Id()),
	})	iferr!=nil{
rnfe*types.NotFoundException
errors.As(err,&nfe){
returnnil
erncreate.DiagError(names.MediaLive,create.ErrActionDeleting,ResNameChannel,d.Id(),err)
	}	if_,err:=waitChannelDeleted(ctx,conn,d.Id(),d.Timeout(schema.TimeoutDelete));err!=nil{
turncreate.DiagError(names.MediaLive,create.ErrActionWaitingForDeletion,ResNameChannel,d.Id(),err)
	}	returnnil
}funcstartChannel(ctxcontext.Context,conn*medialive.Client,timeouttime.Duration,idstring)error{
	_,err:=conn.StartChannel(ctx,&medialive.StartChannelInput{
annelId:aws.String(id),
	})	iferr!=nil{
turnfmt.Errorf("startingMedialiveChannel(%s):%s",id,err)
	}	_,err=waitChannelStarted(ctx,conn,id,timeout)	iferr!=nil{
turnfmt.Errorf("waitingforMedialiveChannel(%s)start:%s",id,err)
	}	returnnil
}funcstopChannel(ctxcontext.Context,conn*medialive.Client,timeouttime.Duration,idstring)error{
	_,err:=conn.StopChannel(ctx,&medialive.StopChannelInput{
annelId:aws.String(id),
	})	iferr!=nil{
turnfmt.Errorf("stoppingMedialiveChannel(%s):%s",id,err)
	}	_,err=waitChannelStopped(ctx,conn,id,timeout)	iferr!=nil{
turnfmt.Errorf("waitingforMedialiveChannel(%s)stop:%s",id,err)
	}	returnnil
}funcwaitChannelCreated(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
nding:enum.Slice(types.ChannelStateCreating),
rget:enum.Slice(types.ChannelStateIdle),
fresh:statusChannel(ctx,conn,id),
meout:timeout,
tFoundChecks:20,
ntinuousTargetOccurence:2,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
turnout,err
	}	returnnil,err
}funcwaitChannelUpdated(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
nding:enum.Slice(types.ChannelStateUpdating),
rget:enum.Slice(types.ChannelStateIdle),
fresh:statusChannel(ctx,conn,id),
meout:timeout,
tFoundChecks:20,
ntinuousTargetOccurence:2,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
turnout,err
	}	returnnil,err
}funcwaitChannelDeleted(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
nding:enum.Slice(types.ChannelStateDeleting),
rget:[]string{},
fresh:statusChannel(ctx,conn,id),
meout:timeout,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
turnout,err
	}	returnnil,err
}funcwaitChannelStarted(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
nding:enum.Slice(types.ChannelStateStarting),
rget:enum.Slice(types.ChannelStateRunning),
fresh:statusChannel(ctx,conn,id),
meout:timeout,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
turnout,err
	}	returnnil,err
}funcwaitChannelStopped(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
nding:enum.Slice(types.ChannelStateStopping),
rget:enum.Slice(types.ChannelStateIdle),
fresh:statusChannel(ctx,conn,id),
meout:timeout,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
turnout,err
	}	returnnil,err
}funcstatusChannel(ctxcontext.Context,conn*medialive.Client,idstring)retry.StateRefreshFunc{
	returnfunc()(interface{},string,error){
t,err:=FindChannelByID(ctx,conn,id)
tfresource.NotFound(err){
returnnil,"",nil
fr!=nil{
returnnil,"",err
ernout,string(out.State),nil
	}
}funcFindChannelByID(ctxcontext.Context,conn*medialive.Client,idstring)(*medialive.DescribeChannelOutput,error){
	in:=&medialive.DescribeChannelInput{
annelId:aws.String(id),
	}
	out,err:=conn.DescribeChannel(ctx,in)
	iferr!=nil{
rnfe*types.NotFoundException
errors.As(err,&nfe){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:in,
}
ernnil,err
	}	ifout==nil{
turnnil,tfresource.NewEmptyResultError(in)
	}	//ChannelcanstillbefoundwithastateofDELETED.
	//Setresultasnotfoundwhenthestateisdeleted.
	ifout.State==types.ChannelStateDeleted{
turnnil,&retry.NotFoundError{
LastResponse:string(types.ChannelStateDeleted),
LastRequest:in,	}	returnout,nil
}funcexpandChannelInputAttachments(tfList[]interface{})[]types.InputAttachment{
	varattachments[]types.InputAttachment
	for_,v:=rangetfList{
ok:=v.(map[string]interface{})
!ok{
continue
atypes.InputAttachment
v,ok:=m["input_attachment_name"].(string);ok{
a.InputAttachmentName=aws.String(v)v,ok:=m["input_id"].(string);ok{
a.InputId=aws.String(v)v,ok:=m["input_settings"].([]interface{});ok&&len(v)>0{
a.InputSettings=expandInputAttachmentInputSettings(v)v,ok:=m["automatic_input_failover_settings"].([]interface{});ok&&len(v)>0{
a.AutomaticInputFailoverSettings=expandInputAttachmentAutomaticInputFailoverSettings(v)
tchments=append(attachments,a)
	}	returnattachments
}funcexpandInputAttachmentInputSettings(tfList[]interface{})*types.InputSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.InputSettings
	ifv,ok:=m["audio_selector"].([]interface{});ok&&len(v)>0{
t.AudioSelectors=expandInputAttachmentInputSettingsAudioSelectors(v)
	}
	ifv,ok:=m["caption_selector"].([]interface{});ok&&len(v)>0{
t.CaptionSelectors=expandInputAttachmentInputSettingsCaptionSelectors(v)
	}
	ifv,ok:=m["deblock_filter"].(string);ok&&v!=""{
t.DeblockFilter=types.InputDeblockFilter(v)
	}
	ifv,ok:=m["denoise_filter"].(string);ok&&v!=""{
t.DenoiseFilter=types.InputDenoiseFilter(v)
	}
	ifv,ok:=m["filter_strength"].(int);ok{
t.FilterStrength=int32(v)
	}
	ifv,ok:=m["input_filter"].(string);ok&&v!=""{
t.InputFilter=types.InputFilter(v)
	}
	ifv,ok:=m["network_input_settings"].([]interface{});ok&&len(v)>0{
t.NetworkInputSettings=expandInputAttachmentInputSettingsNetworkInputSettings(v)
	}
	ifv,ok:=m["scte35_pid"].(int);ok{
t.Scte35Pid=int32(v)
	}
	ifv,ok:=m["smpte2038_data_preference"].(string);ok&&v!=""{
t.Smpte2038DataPreference=types.Smpte2038DataPreference(v)
	}
	ifv,ok:=m["source_end_behavior"].(string);ok&&v!=""{
t.SourceEndBehavior=types.InputSourceEndBehavior(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsAudioSelectors(tfList[]interface{})[]types.AudioSelector{
	varas[]types.AudioSelector
	for_,v:=rangetfList{
ok:=v.(map[string]interface{})
!ok{
continue
atypes.AudioSelector
v,ok:=m["name"].(string);ok&&v!=""{
a.Name=aws.String(v)v,ok:=m["selector_settings"].([]interface{});ok&&len(v)>0{
a.SelectorSettings=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettings(v)
sppend(as,a)
	}	returnas
}funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettings(tfList[]interface{})*types.AudioSelectorSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AudioSelectorSettings
	ifv,ok:=m["audio_hls_rendition_selection"].([]interface{});ok&&len(v)>0{
t.AudioHlsRenditionSelection=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioHlsRenditionSelection(v)
	}
	ifv,ok:=m["audio_language_selection"].([]interface{});ok&&len(v)>0{
t.AudioLanguageSelection=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioLanguageSelection(v)
	}
	ifv,ok:=m["audio_pid_selection"].([]interface{});ok&&len(v)>0{
t.AudioPidSelection=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioPidSelection(v)
	}
	ifv,ok:=m["audio_track_selection"].([]interface{});ok&&len(v)>0{
t.AudioTrackSelection=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelection(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioHlsRenditionSelection(tfList[]interface{})*types.AudioHlsRenditionSelection{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AudioHlsRenditionSelection
	ifv,ok:=m["group_id"].(string);ok&&len(v)>0{
t.GroupId=aws.String(v)
	}
	ifv,ok:=m["name"].(string);ok&&len(v)>0{
t.Name=aws.String(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioLanguageSelection(tfList[]interface{})*types.AudioLanguageSelection{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AudioLanguageSelection
	ifv,ok:=m["language_code"].(string);ok&&len(v)>0{
t.LanguageCode=aws.String(v)
	}
	ifv,ok:=m["language_selection_policy"].(string);ok&&len(v)>0{
t.LanguageSelectionPolicy=types.AudioLanguageSelectionPolicy(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioPidSelection(tfList[]interface{})*types.AudioPidSelection{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AudioPidSelection
	ifv,ok:=m["pid"].(int);ok{
t.Pid=int32(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelection(tfList[]interface{})*types.AudioTrackSelection{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AudioTrackSelection
	ifv,ok:=m["tracks"].(*schema.Set);ok&&v.Len()>0{
t.Tracks=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionTracks(v.List())
	}
	ifv,ok:=m["dolby_e_decode"].([]interface{});ok&&len(v)>0{
t.DolbyEDecode=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionDolbyEDecode(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionTracks(tfList[]interface{})[]types.AudioTrack{
	iflen(tfList)==0{
turnnil
	}	varout[]types.AudioTrack
	for_,v:=rangetfList{
ok:=v.(map[string]interface{})
!ok{
continue
atypes.AudioTrack
v,ok:=m["track"].(int);ok{
o.Track=int32(v)
uappend(out,o)
	}	returnout
}funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionDolbyEDecode(tfList[]interface{})*types.AudioDolbyEDecode{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AudioDolbyEDecode
	ifv,ok:=m["program_selection"].(string);ok&&v!=""{
t.ProgramSelection=types.DolbyEProgramSelection(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsCaptionSelectors(tfList[]interface{})[]types.CaptionSelector{
	iflen(tfList)==0{
turnnil
	}	varout[]types.CaptionSelector
	for_,v:=rangetfList{
ok:=v.(map[string]interface{})
!ok{
continue
atypes.CaptionSelector
v,ok:=m["name"].(string);ok&&v!=""{
o.Name=aws.String(v)v,ok:=m["language_code"].(string);ok&&v!=""{
o.LanguageCode=aws.String(v)v,ok:=m["selector_settings"].([]interface{});ok&&len(v)>0{
o.SelectorSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettings(v)
uappend(out,o)
	}	returnout
}funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettings(tfList[]interface{})*types.CaptionSelectorSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.CaptionSelectorSettings
	ifv,ok:=m["ancillary_source_settings"].([]interface{});ok&&len(v)>0{
t.AncillarySourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsAncillarySourceSettings(v)
	}
	ifv,ok:=m["arib_source_settings"].([]interface{});ok&&len(v)>0{
t.AribSourceSettings=&types.AribSourceSettings{}//noexportedfields
	}
	ifv,ok:=m["dvb_sub_source_settings"].([]interface{});ok&&len(v)>0{
t.DvbSubSourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsDvbSubSourceSettings(v)
	}
	ifv,ok:=m["embedded_source_settings"].([]interface{});ok&&len(v)>0{
t.EmbeddedSourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsEmbeddedSourceSettings(v)
	}
	ifv,ok:=m["scte20_source_settings"].([]interface{});ok&&len(v)>0{
t.Scte20SourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsScte20SourceSettings(v)
	}
	ifv,ok:=m["scte27_source_settings"].([]interface{});ok&&len(v)>0{
t.Scte27SourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsScte27SourceSettings(v)
	}
	ifv,ok:=m["teletext_source_settings"].([]interface{});ok&&len(v)>0{
t.TeletextSourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettings(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsAncillarySourceSettings(tfList[]interface{})*types.AncillarySourceSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AncillarySourceSettings
	ifv,ok:=m["source_ancillary_channel_number"].(int);ok{
t.SourceAncillaryChannelNumber=int32(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsDvbSubSourceSettings(tfList[]interface{})*types.DvbSubSourceSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.DvbSubSourceSettings
	ifv,ok:=m["ocr_language"].(string);ok&&v!=""{
t.OcrLanguage=types.DvbSubOcrLanguage(v)
	}
	ifv,ok:=m["pid"].(int);ok{
t.Pid=int32(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsEmbeddedSourceSettings(tfList[]interface{})*types.EmbeddedSourceSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.EmbeddedSourceSettings
	ifv,ok:=m["convert_608_to_708"].(string);ok&&v!=""{
t.Convert608To708=types.EmbeddedConvert608To708(v)
	}
	ifv,ok:=m["scte20_detection"].(string);ok&&v!=""{
t.Scte20Detection=types.EmbeddedScte20Detection(v)
	}
	ifv,ok:=m["source_608_channel_number"].(int);ok{
t.Source608ChannelNumber=int32(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsScte20SourceSettings(tfList[]interface{})*types.Scte20SourceSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.Scte20SourceSettings
	ifv,ok:=m["convert_608_to_708"].(string);ok&&v!=""{
t.Convert608To708=types.Scte20Convert608To708(v)
	}
	ifv,ok:=m["source_608_channel_number"].(int);ok{
t.Source608ChannelNumber=int32(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsScte27SourceSettings(tfList[]interface{})*types.Scte27SourceSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.Scte27SourceSettings
	ifv,ok:=m["ocr_language"].(string);ok&&v!=""{
t.OcrLanguage=types.Scte27OcrLanguage(v)
	}
	ifv,ok:=m["pid"].(int);ok{
t.Pid=int32(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettings(tfList[]interface{})*types.TeletextSourceSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.TeletextSourceSettings
	ifv,ok:=m["output_rectangle"].([]interface{});ok&&len(v)>0{
t.OutputRectangle=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettingsOutputRectangle(v)
	}
	ifv,ok:=m["page_number"].(string);ok&&v!=""{
t.PageNumber=aws.String(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettingsOutputRectangle(tfList[]interface{})*types.CaptionRectangle{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.CaptionRectangle
	ifv,ok:=m["height"].(float32);ok{
t.Height=float64(v)
	}
	ifv,ok:=m["left_offset"].(float32);ok{
t.LeftOffset=float64(v)
	}
	ifv,ok:=m["top_offset"].(float32);ok{
t.TopOffset=float64(v)
	}
	ifv,ok:=m["width"].(float32);ok{
t.Width=float64(v)
	}	return&out
}funcexpandInputAttachmentInputSettingsNetworkInputSettings(tfList[]interface{})*types.NetworkInputSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.NetworkInputSettings
	ifv,ok:=m["hls_input_settings"].([]interface{});ok&&len(v)>0{
t.HlsInputSettings=expandNetworkInputSettingsHLSInputSettings(v)
	}
	ifv,ok:=m["server_validation"].(string);ok&&v!=""{
t.ServerValidation=types.NetworkInputServerValidation(v)
	}	return&out
}funcexpandNetworkInputSettingsHLSInputSettings(tfList[]interface{})*types.HlsInputSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.HlsInputSettings
	ifv,ok:=m["bandwidth"].(int);ok{
t.Bandwidth=int32(v)
	}
	ifv,ok:=m["buffer_segments"].(int);ok{
t.BufferSegments=int32(v)
	}
	ifv,ok:=m["retries"].(int);ok{
t.Retries=int32(v)
	}
	ifv,ok:=m["retry_interval"].(int);ok{
t.RetryInterval=int32(v)
	}
	ifv,ok:=m["scte35_source"].(string);ok&&v!=""{
t.Scte35Source=types.HlsScte35SourceType(v)
	}	return&out
}funcexpandInputAttachmentAutomaticInputFailoverSettings(tfList[]interface{})*types.AutomaticInputFailoverSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AutomaticInputFailoverSettings
	ifv,ok:=m["secondary_input_id"].(string);ok&&v!=""{
t.SecondaryInputId=aws.String(v)
	}
	ifv,ok:=m["error_clear_time_msec"].(int);ok{
t.ErrorClearTimeMsec=int32(v)
	}
	ifv,ok:=m["failover_conditions"].(*schema.Set);ok&&v.Len()>0{
t.FailoverConditions=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditions(v.List())
	}
	ifv,ok:=m["input_preference"].(string);ok&&v!=""{
t.InputPreference=types.InputPreference(v)
	}	return&out
}funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditions(tfList[]interface{})[]types.FailoverCondition{
	iflen(tfList)==0{
turnnil
	}	varout[]types.FailoverCondition
	for_,v:=rangetfList{
ok:=v.(map[string]interface{})
!ok{
continue
atypes.FailoverCondition
v,ok:=m["failover_condition_settings"].([]interface{});ok&&len(v)>0{
o.FailoverConditionSettings=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettings(v)
uappend(out,o)
	}	returnout
}funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettings(tfList[]interface{})*types.FailoverConditionSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.FailoverConditionSettings
	ifv,ok:=m["audio_silence_settings"].([]interface{});ok&&len(v)>0{
t.AudioSilenceSettings=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsAudioSilenceSettings(v)
	}
	ifv,ok:=m["input_loss_settings"].([]interface{});ok&&len(v)>0{
t.InputLossSettings=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsInputLossSettings(v)
	}
	ifv,ok:=m["video_black_settings"].([]interface{});ok&&len(v)>0{
t.VideoBlackSettings=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsVideoBlackSettings(v)
	}	return&out
}funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsAudioSilenceSettings(tfList[]interface{})*types.AudioSilenceFailoverSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.AudioSilenceFailoverSettings
	ifv,ok:=m["audio_selector_name"].(string);ok&&v!=""{
t.AudioSelectorName=aws.String(v)
	}
	ifv,ok:=m["audio_silence_threshold_msec"].(int);ok{
t.AudioSilenceThresholdMsec=int32(v)
	}	return&out
}funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsInputLossSettings(tfList[]interface{})*types.InputLossFailoverSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.InputLossFailoverSettings
	ifv,ok:=m["input_loss_threshold_msec"].(int);ok{
t.InputLossThresholdMsec=int32(v)
	}	return&out
}funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsVideoBlackSettings(tfList[]interface{})*types.VideoBlackFailoverSettings{
	iftfList==nil{
turnnil
	}	m:=tfList[0].(map[string]interface{})	varouttypes.VideoBlackFailoverSettings
	ifv,ok:=m["black_detect_threshold"].(float32);ok{
t.BlackDetectThreshold=float64(v)
	}
	ifv,ok:=m["video_black_threshold_msec"].(int);ok{
t.VideoBlackThresholdMsec=int32(v)
	}	return&out
}funcflattenChannelInputAttachments(tfList[]types.InputAttachment)[]interface{}{
	iflen(tfList)==0{
turnnil
	}	varout[]interface{}	for_,item:=rangetfList{
=map[string]interface{}{
"input_id":ToString(item.InputId),
"input_attachment_name":aws.ToString(item.InputAttachmentName),
"input_settings":flattenInputAttachmentsInputSettings(item.InputSettings),
"automatic_input_failover_settings":flattenInputAttachmentAutomaticInputFailoverSettings(item.AutomaticInputFailoverSettings),
uappend(out,m)
	}	returnout
}funcflattenInputAttachmentsInputSettings(in*types.InputSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
udio_selector":flattenInputAttachmentsInputSettingsAudioSelectors(in.AudioSelectors),
aption_selector":flattenInputAttachmentsInputSettingsCaptionSelectors(in.CaptionSelectors),
eblock_filter":string(in.DeblockFilter),
enoise_filter":string(in.DenoiseFilter),
ilter_strength":int(in.FilterStrength),
nput_filter":string(in.InputFilter),
etwork_input_settings":flattenInputAttachmentsInputSettingsNetworkInputSettings(in.NetworkInputSettings),
cte35_pid":n.Scte35Pid),
mpte2038_data_preference":string(in.Smpte2038DataPreference),
ource_end_behavior":g(in.SourceEndBehavior),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsAudioSelectors(tfList[]types.AudioSelector)[]interface{}{
	iflen(tfList)==0{
turnnil
	}	varout[]interface{}	for_,v:=rangetfList{
=map[string]interface{}{
"name":aws.ToString(v.Name),
"selector_settings":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettings(v.SelectorSettings),
uappend(out,m)
	}	returnout
}funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettings(in*types.AudioSelectorSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
udio_hls_rendition_selection":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioHlsRenditionSelection(in.AudioHlsRenditionSelection),
udio_language_selection":nInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioLanguageSelection(in.AudioLanguageSelection),
udio_pid_selection":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioPidSelection(in.AudioPidSelection),
udio_track_selection":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelection(in.AudioTrackSelection),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioHlsRenditionSelection(in*types.AudioHlsRenditionSelection)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
roup_id":aws.ToString(in.GroupId),
ame":aws.ToString(in.Name),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioLanguageSelection(in*types.AudioLanguageSelection)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
anguage_code":aws.ToString(in.LanguageCode),
anguage_selection_policy":string(in.LanguageSelectionPolicy),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioPidSelection(in*types.AudioPidSelection)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
id":int(in.Pid),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelection(in*types.AudioTrackSelection)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
olby_e_decode":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionDolbyEDecode(in.DolbyEDecode),
racks":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionTracks(in.Tracks),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionDolbyEDecode(in*types.AudioDolbyEDecode)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
rogram_selection":string(in.ProgramSelection),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionTracks(tfList[]types.AudioTrack)[]interface{}{
	iflen(tfList)==0{
turnnil
	}	varout[]interface{}	for_,v:=rangetfList{
=map[string]interface{}{
"track":int(v.Track),
uappend(out,m)
	}	returnout
}funcflattenInputAttachmentsInputSettingsCaptionSelectors(tfList[]types.CaptionSelector)[]interface{}{
	iflen(tfList)==0{
turnnil
	}	varout[]interface{}	for_,v:=rangetfList{
=map[string]interface{}{
"name":aws.ToString(v.Name),
"language_code":aws.ToString(v.LanguageCode),
"selector_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettings(v.SelectorSettings),
uappend(out,m)
	}	returnout
}funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettings(in*types.CaptionSelectorSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
ncillary_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsAncillarySourceSettings(in.AncillarySourceSettings),
rib_source_settings":rface{}{},//attributehasnoexportedfields
vb_sub_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsDvbSubSourceSettings(in.DvbSubSourceSettings),
mbedded_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsEmbeddedSourceSettings(in.EmbeddedSourceSettings),
cte20_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsScte20SourceSettings(in.Scte20SourceSettings),
cte27_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsScte27SourceSettings(in.Scte27SourceSettings),
eletext_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettings(in.TeletextSourceSettings),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsAncillarySourceSettings(in*types.AncillarySourceSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
ource_ancillary_channel_number":int(in.SourceAncillaryChannelNumber),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsDvbSubSourceSettings(in*types.DvbSubSourceSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
cr_language":string(in.OcrLanguage),
id":int(in.Pid),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsEmbeddedSourceSettings(in*types.EmbeddedSourceSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
onvert_608_to_708":ng(in.Convert608To708),
cte20_detection":string(in.Scte20Detection),
ource_608_channel_number":int(in.Source608ChannelNumber),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsScte20SourceSettings(in*types.Scte20SourceSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
onvert_608_to_708":ng(in.Convert608To708),
ource_608_channel_number":int(in.Source608ChannelNumber),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsScte27SourceSettings(in*types.Scte27SourceSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
cr_language":string(in.OcrLanguage),
id":int(in.Pid),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettings(in*types.TeletextSourceSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
utput_rectangle":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettingsOutputRectangle(in.OutputRectangle),
age_number":String(in.PageNumber),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettingsOutputRectangle(in*types.CaptionRectangle)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
eight":2(in.Height),
eft_offset":float32(in.LeftOffset),
op_offset":float32(in.TopOffset),
idth":32(in.Width),
	}	return[]interface{}{m}
}funcflattenInputAttachmentsInputSettingsNetworkInputSettings(in*types.NetworkInputSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
ls_input_settings":flattenNetworkInputSettingsHLSInputSettings(in.HlsInputSettings),
erver_validation":string(in.ServerValidation),
	}	return[]interface{}{m}
}funcflattenNetworkInputSettingsHLSInputSettings(in*types.HlsInputSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
andwidth":n.Bandwidth),
uffer_segments":int(in.BufferSegments),
etries":int(in.Retries),
etry_interval":int(in.RetryInterval),
cte35_source":string(in.Scte35Source),
	}	return[]interface{}{m}
}funcflattenInputAttachmentAutomaticInputFailoverSettings(in*types.AutomaticInputFailoverSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
econdary_input_id":aws.ToString(in.SecondaryInputId),
rror_clear_time_msec":int(in.ErrorClearTimeMsec),
ailover_conditions":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditions(in.FailoverConditions),
nput_preference":(in.InputPreference),
	}	return[]interface{}{m}
}funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditions(tfList[]types.FailoverCondition)[]interface{}{
	iflen(tfList)==0{
turnnil
	}	varout[]interface{}	for_,item:=rangetfList{
=map[string]interface{}{
"failover_condition_settings":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettings(item.FailoverConditionSettings),
uappend(out,m)
	}
	returnout
}funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettings(in*types.FailoverConditionSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
udio_silence_settings":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsAudioSilenceSettings(in.AudioSilenceSettings),
nput_loss_settings":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsInputLossSettings(in.InputLossSettings),
ideo_black_settings":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsVideoBlackSettings(in.VideoBlackSettings),
	}	return[]interface{}{m}
}funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsAudioSilenceSettings(in*types.AudioSilenceFailoverSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
udio_selector_name":aws.ToString(in.AudioSelectorName),
udio_silence_threshold_msec":int(in.AudioSilenceThresholdMsec),
	}	return[]interface{}{m}
}funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsInputLossSettings(in*types.InputLossFailoverSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
nput_loss_threshold_msec":int(in.InputLossThresholdMsec),
	}	return[]interface{}{m}
}funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsVideoBlackSettings(in*types.VideoBlackFailoverSettings)[]interface{}{
	ifin==nil{
turnnil
	}	m:=map[string]interface{}{
lack_detect_threshold":float32(in.BlackDetectThreshold),
ideo_black_threshold_msec":int(in.VideoBlackThresholdMsec),
	}	return[]interface{}{m}
}funcexpandChannelCdiInputSpecification(tfList[]interface{})*types.CdiInputSpecification{
	iftfList==nil{
turnnil
	}
	m:=tfList[0].(map[string]interface{})	spec:=&types.CdiInputSpecification{}
	ifv,ok:=m["resolution"].(string);ok&&v!=""{
ec.Resolution=types.CdiInputResolution(v)
	}	returnspec
}funcflattenChannelCdiInputSpecification(apiObject*types.CdiInputSpecification)[]interface{}{
	ifapiObject==nil{
turnnil
	}	m:=map[string]interface{}{
esolution":string(apiObject.Resolution),
	}	return[]interface{}{m}
}funcexpandChannelDestinations(tfList[]interface{})[]types.OutputDestination{
	iftfList==nil{
turnnil
	}	vardestinations[]types.OutputDestination
	for_,v:=rangetfList{
ok:=v.(map[string]interface{})
!ok{
continue
atypes.OutputDestination
v,ok:=m["id"].(string);ok{
d.Id=aws.String(v)v,ok:=m["media_package_settings"].(*schema.Set);ok&&v.Len()>0{
d.MediaPackageSettings=expandChannelDestinationsMediaPackageSettings(v.List())v,ok:=m["multiplex_settings"].([]interface{});ok&&len(v)>0{
d.MultiplexSettings=expandChannelDestinationsMultiplexSettings(v)v,ok:=m["settings"].(*schema.Set);ok&&v.Len()>0{
d.Settings=expandChannelDestinationsSettings(v.List())
einations=append(destinations,d)
	}	returndestinations
}funcexpandChannelDestinationsMediaPackageSettings(tfList[]interface{})[]types.MediaPackageOutputDestinationSettings{
	iftfList==nil{
turnnil
	}	varsettings[]types.MediaPackageOutputDestinationSettings
	for_,v:=rangetfList{
ok:=v.(map[string]interface{})
!ok{
continue
atypes.MediaPackageOutputDestinationSettings
v,ok:=m["channel_id"].(string);ok{
s.ChannelId=aws.String(v)
eings=append(settings,s)
	}	returnsettings
}funcexpandChannelDestinationsMultiplexSettings(tfList[]interface{})*types.MultiplexProgramChannelDestinationSettings{
	iftfList==nil{
turnnil
	}
	m:=tfList[0].(map[string]interface{})	settings:=&types.MultiplexProgramChannelDestinationSettings{}
	ifv,ok:=m["multiplex_id"].(string);ok&&v!=""{
ttings.MultiplexId=aws.String(v)
	}
	ifv,ok:=m["program_name"].(string);ok&&v!=""{
ttings.ProgramName=aws.String(v)
	}	returnsettings
}funcexpandChannelDestinationsSettings(tfList[]interface{})[]types.OutputDestinationSettings{
	iftfList==nil{
turnnil
	}	varsettings[]types.OutputDestinationSettings
	for_,v:=rangetfList{
ok:=v.(map[string]interface{})
!ok{
continue
atypes.OutputDestinationSettings
v,ok:=m["password_param"].(string);ok{
s.PasswordParam=aws.String(v)v,ok:=m["stream_name"].(string);ok{
s.StreamName=aws.String(v)v,ok:=m["url"].(string);ok{
s.Url=aws.String(v)v,ok:=m["username"].(string);ok{
s.Username=aws.String(v)
eings=append(settings,s)
	}	returnsettings
}funcflattenChannelDestinations(apiObject[]types.OutputDestination)[]interface{}{
	ifapiObject==nil{
turnnil
	}	vartfList[]interface{}
	for_,v:=rangeapiObject{
=map[string]interface{}{
"id":aws.ToString(v.Id),
"media_package_settings":flattenChannelDestinationsMediaPackageSettings(v.MediaPackageSettings),
"multiplex_settings":flattenChannelDestinationsMultiplexSettings(v.MultiplexSettings),
"settings":nChannelDestinationsSettings(v.Settings),
fst=append(tfList,m)
	}	returntfList
}funcflattenChannelDestinationsMediaPackageSettings(apiObject[]types.MediaPackageOutputDestinationSettings)[]interface{}{
	ifapiObject==nil{
turnnil
	}	vartfList[]interface{}
	for_,v:=rangeapiObject{
=map[string]interface{}{
"channel_id":aws.ToString(v.ChannelId),
fst=append(tfList,m)
	}	returntfList
}funcflattenChannelDestinationsMultiplexSettings(apiObject*types.MultiplexProgramChannelDestinationSettings)[]interface{}{
	ifapiObject==nil{
turnnil
	}	m:=map[string]interface{}{
ultiplex_id":aws.ToString(apiObject.MultiplexId),
rogram_name":aws.ToString(apiObject.ProgramName),
	}	return[]interface{}{m}
}funcflattenChannelDestinationsSettings(apiObject[]types.OutputDestinationSettings)[]interface{}{
	ifapiObject==nil{
turnnil
	}	vartfList[]interface{}
	for_,v:=rangeapiObject{
=map[string]interface{}{
"password_param":aws.ToString(v.PasswordParam),
"stream_name":aws.ToString(v.StreamName),
"url":aws.ToString(v.Url),
"username":oString(v.Username),
fst=append(tfList,m)
	}	returntfList
}funcexpandChannelInputSpecification(tfList[]interface{})*types.InputSpecification{
	iftfList==nil{
turnnil
	}
	m:=tfList[0].(map[string]interface{})	spec:=&types.InputSpecification{}
	ifv,ok:=m["codec"].(string);ok&&v!=""{
ec.Codec=types.InputCodec(v)
	}
	ifv,ok:=m["maximum_bitrate"].(string);ok&&v!=""{
ec.MaximumBitrate=types.InputMaximumBitrate(v)
	}
	ifv,ok:=m["input_resolution"].(string);ok&&v!=""{
ec.Resolution=types.InputResolution(v)
	}	returnspec
}funcflattenChannelInputSpecification(apiObject*types.InputSpecification)[]interface{}{
	ifapiObject==nil{
turnnil
	}	m:=map[string]interface{}{
odec":string(apiObject.Codec),
aximum_bitrate":string(apiObject.MaximumBitrate),
nput_resolution":string(apiObject.Resolution),
	}	return[]interface{}{m}
}funcexpandChannelMaintenanceCreate(tfList[]interface{})*types.MaintenanceCreateSettings{
	iftfList==nil{
turnnil
	}
	m:=tfList[0].(map[string]interface{})	settings:=&types.MaintenanceCreateSettings{}
	ifv,ok:=m["maintenance_day"].(string);ok&&v!=""{
ttings.MaintenanceDay=types.MaintenanceDay(v)
	}
	ifv,ok:=m["maintenance_start_time"].(string);ok&&v!=""{
ttings.MaintenanceStartTime=aws.String(v)
	}	returnsettings
}funcexpandChannelMaintenanceUpdate(tfList[]interface{})*types.MaintenanceUpdateSettings{
	iftfList==nil{
turnnil
	}
	m:=tfList[0].(map[string]interface{})	settings:=&types.MaintenanceUpdateSettings{}
	ifv,ok:=m["maintenance_day"].(string);ok&&v!=""{
ttings.MaintenanceDay=types.MaintenanceDay(v)
	}
	ifv,ok:=m["maintenance_start_time"].(string);ok&&v!=""{
ttings.MaintenanceStartTime=aws.String(v)
	}
	//NOTE:Thisfieldisonlyavailableintheupdatestruct.Toallowuserstosetascheduled
	//dateonupdate,itmaybeworthaddingtothebaseschema.
	//ifv,ok:=m["maintenance_scheduled_date"].(string);ok&&v!=""{
	//	settings.MaintenanceScheduledDate=aws.String(v)
	//}	returnsettings
}funcflattenChannelMaintenance(apiObject*types.MaintenanceStatus)[]interface{}{
	ifapiObject==nil{
turnnil
	}	m:=map[string]interface{}{
aintenance_day":ng(apiObject.MaintenanceDay),
aintenance_start_time":aws.ToString(apiObject.MaintenanceStartTime),
	}	return[]interface{}{m}
}funcexpandChannelVPC(tfList[]interface{})*types.VpcOutputSettings{
	iftfList==nil{
turnnil
	}
	m:=tfList[0].(map[string]interface{})	settings:=&types.VpcOutputSettings{}
	ifv,ok:=m["security_group_ids"].([]string);ok&&len(v)>0{
ttings.SecurityGroupIds=v
	}
	ifv,ok:=m["subnet_ids"].([]string);ok&&len(v)>0{
ttings.SubnetIds=v
	}
	ifv,ok:=m["public_address_allocation_ids"].([]string);ok&&len(v)>0{
ttings.PublicAddressAllocationIds=v
	}	returnsettings
}funcflattenChannelVPC(apiObject*types.VpcOutputSettingsDescription)[]interface{}{
	ifapiObject==nil{
turnnil
	}	m:=map[string]interface{}{
ecurity_group_ids":flex.FlattenStringValueList(apiObject.SecurityGroupIds),
ubnet_ids":flex.FlattenStringValueList(apiObject.SubnetIds),
public_address_allocation_idsisnotincludedintheoutputstruct
	}	return[]interface{}{m}
}
