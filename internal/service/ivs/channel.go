//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packageivsimport(
	"context"
	"errors"
	"log"
	"time"	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)//@SDKResource("aws_ivs_channel",name="Channel")
//@Tags(identifierAttribute="id")
funcResourceChannel()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceChannelCreate,
		ReadWithoutTimeout:resourceChannelRead,
		UpdateWithoutTimeout:resourceChannelUpdate,
		DeleteWithoutTimeout:resourceChannelDelete,		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},		Timeouts:&schema.ResourceTimeout{
			Create:schema.DefaultTimeout(5*time.Minute),
			Update:schema.DefaultTimeout(5*time.Minute),
			Delete:schema.DefaultTimeout(5*time.Minute),
		},		Schema:map[string]*schema.Schema{
			"arn":{
				Type:schema.TypeString,
				Computed:true,
			},
			"authorized":{
				Type:schema.TypeBool,
				Optional:true,
				Computed:true,
			},
			"ingest_endpoint":{
				Type:schema.TypeString,
				Computed:true,
			},
			"latency_mode":{
				Type:schema.TypeString,
				Optional:true,
				Computed:true,
				ValidateFunc:validation.StringInSlice(ivs.ChannelLatencyMode_Values(),false),
			},
			"name":{
				Type:schema.TypeString,
				Optional:true,
				Computed:true,
				ValidateFunc:validation.StringMatch(regexache.MustCompile(`^[0-9A-Za-z_-]{0,128}$`),"mustcontainonlyalphanumericcharacters,hyphen,orunderscoreandatmost128characters"),
			},
			"playback_url":{
				Type:schema.TypeString,
				Computed:true,
			},
			"recording_configuration_arn":{
				Type:schema.TypeString,
				Optional:true,
				Computed:true,
				ValidateFunc:verify.ValidARN,
			},
			names.AttrTags:tftags.TagsSchema(),
			names.AttrTagsAll:tftags.TagsSchemaComputed(),
			"type":{
				Type:schema.TypeString,
				Optional:true,
				Computed:true,
				ValidateFunc:validation.StringInSlice(ivs.ChannelType_Values(),false),
			},
		},		CustomizeDiff:verify.SetTagsDiff,
	}
}const(
	ResNameChannel="Channel"
)
funcresourceChannelCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).IVSConn(ctx)	in:=&ivs.CreateChannelInput{
		Tags:getTagsIn(ctx),
	}	ifv,ok:=d.GetOk("authorized");ok{
		in.Authorized=aws.Bool(v.(bool))
	}	ifv,ok:=d.GetOk("latency_mode");ok{
		in.LatencyMode=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("name");ok{
		in.Name=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("recording_configuration_arn");ok{
		in.RecordingConfigurationArn=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("type");ok{
		in.Type=aws.String(v.(string))
	}	out,err:=conn.CreateChannelWithContext(ctx,in)
	iferr!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionCreating,ResNameChannel,d.Get("name").(string),err)
	}	ifout==nil||out.Channel==nil{
		returncreate.DiagError(names.IVS,create.ErrActionCreating,ResNameChannel,d.Get("name").(string),errors.New("emptyoutput"))
	}	d.SetId(aws.StringValue(out.Channel.Arn))	if_,err:=waitChannelCreated(ctx,conn,d.Id(),d.Timeout(schema.TimeoutCreate));err!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionWaitingForCreation,ResNameChannel,d.Id(),err)
	}	returnresourceChannelRead(ctx,d,meta)
}
funcresourceChannelRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).IVSConn(ctx)	out,err:=FindChannelByID(ctx,conn,d.Id())	if!d.IsNewResource()&&tfresource.NotFound(err){
		log.Printf("[WARN]IVSChannel(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returnnil
	}	iferr!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionReading,ResNameChannel,d.Id(),err)
	}	d.Set("arn",out.Arn)
	d.Set("authorized",out.Authorized)
	d.Set("ingest_endpoint",out.IngestEndpoint)
	d.Set("latency_mode",out.LatencyMode)
	d.Set("name",out.Name)
	d.Set("playback_url",out.PlaybackUrl)
	d.Set("recording_configuration_arn",out.RecordingConfigurationArn)
	d.Set("type",out.Type)	returnnil
}
funcresourceChannelUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).IVSConn(ctx)	update:=false	arn:=d.Id()
	in:=&ivs.UpdateChannelInput{
		Arn:aws.String(arn),
	}	ifd.HasChanges("authorized"){
		in.Authorized=aws.Bool(d.Get("authorized").(bool))
		update=true
	}	ifd.HasChanges("latency_mode"){
		in.LatencyMode=aws.String(d.Get("latency_mode").(string))
		update=true
	}	ifd.HasChanges("name"){
		in.Name=aws.String(d.Get("name").(string))
		update=true
	}	ifd.HasChanges("recording_configuration_arn"){
		in.RecordingConfigurationArn=aws.String(d.Get("recording_configuration_arn").(string))
		update=true
	}	ifd.HasChanges("type"){
		in.Type=aws.String(d.Get("type").(string))
		update=true
	}	if!update{
		returnnil
	}	log.Printf("[DEBUG]UpdatingIVSChannel(%s):%#v",d.Id(),in)	out,err:=conn.UpdateChannelWithContext(ctx,in)
	iferr!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
	}	if_,err:=waitChannelUpdated(ctx,conn,*out.Channel.Arn,d.Timeout(schema.TimeoutUpdate),in);err!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionWaitingForUpdate,ResNameChannel,d.Id(),err)
	}	returnresourceChannelRead(ctx,d,meta)
}
funcresourceChannelDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).IVSConn(ctx)	log.Printf("[INFO]DeletingIVSChannel%s",d.Id())	_,err:=conn.DeleteChannelWithContext(ctx,&ivs.DeleteChannelInput{
		Arn:aws.String(d.Id()),
	})	iferr!=nil{
		iftfawserr.ErrCodeEquals(err,ivs.ErrCodeResourceNotFoundException){
			returnnil
		}		returncreate.DiagError(names.IVS,create.ErrActionDeleting,ResNameChannel,d.Id(),err)
	}	if_,err:=waitChannelDeleted(ctx,conn,d.Id(),d.Timeout(schema.TimeoutDelete));err!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionWaitingForDeletion,ResNameChannel,d.Id(),err)
	}	returnnil
}
