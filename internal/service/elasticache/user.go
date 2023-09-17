//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packageelasticacheimport(
	"context"
	"log"
	"strings"
	"time"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)//@SDKResource("aws_elasticache_user",name="User")
//@Tags(identifierAttribute="arn")
funcurn&schema.Resource{
		CreateWithoutTimeout:resourceUserCreate,
		ReadWithoutTimeout:resourceUserRead,
		UpdateWithoutTimeout:resourceUserUpdate,
		DeleteWithoutTimeout:resourceUserDelete,		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},		CustomizeDiff:verify.SetTagsDiff,		Timeouts:&schema.ResourceTimeout{
			Create:schema.DefaultTimeout(5*time.Minute),
			Update:schema.DefaultTimeout(5*time.Minute),
			Delete:schema.DefaultTimeout(5*time.Minute),
		},		Schema:map[string]*schema.Schema{
			"access_string":{
				Type:schema.TypeString,
				Required:true,
			},
			"arn":{
				Type:schema.TypeString,
				Computed:true,
			},
			"authentication_mode":{
				Type:schema.TypeList,
				Optional:true,
				Computed:true,
				MaxItems:1,
				Elem:&schema.Resource{
					Schema:map[string]*schema.Schema{
						"passwords":{
							Type:schema.TypeSet,
							Optional:true,
							MinItems:1,
							Sensitive:true,
							Elem:&schema.Schema{
								Type:schema.TypeString,
							},
						},
						"password_count":{
							Type:schema.TypeInt,
							Computed:true,
						},
						"type":{
							Type:schema.TypeString,
							Required:true,
							ValidateFunc:validation.StringInSlice(elasticache.InputAuthenticationType_Values(),false),
						},
					},
				},
			},
			"engine":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
				ValidateFunc:validation.StringInSlice([]string{"REDIS"},false),
				DiffSuppressFunc:
func(k,old,newstring,d*schema.ResourceData)bool{
					returnstrings.Eq
func},
			},
			"no_password_required":{
				Type:schema.TypeBool,
				Optional:true,
				Default:false,
			},
			"passwords":{
				Type:schema.TypeSet,
				Optional:true,
				MaxItems:2,
				Elem:&schema.Schema{
					Type:schema.TypeString,
					ValidateFunc:validation.StringLenBetween(16,128),
				},
				Sensitive:true,
			},
			names.AttrTags:tftags.TagsSchema(),
			names.AttrTagsAll:tftags.TagsSchemaComputed(),
			"user_id":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
			},
			"user_name":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
			},
		},
	}
}
funcresourceUserCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
func
	userID:=d.Get("user_id").(string)
	input:=&elasticache.CreateUserInput{
		AccessString:aws.String(d.Get("access_string").(string)),
		Engine:aws.String(d.Get("engine").(string)),
		NoPasswordRequired:aws.Bool(d.Get("no_password_required").(bool)),
		Tags:getTagsIn(ctx),
		UserId:aws.String(userID),
		UserName:aws.String(d.Get("user_name").(string)),
	}	ifv,ok:=d.GetOk("authentication_mode");ok&&len(v.([]interface{}))>0&&v.([]interface{})[0]!=nil{
		input.AuthenticationMode=expandAuthenticationMode(v.([]interface{})[0].(map[string]interface{}))
	}	ifv,ok:=d.GetOk("passwords");ok&&v.(*schema.Set).Len()>0{
		input.Passwords=flex.ExpandStringSet(v.(*schema.Set))
	}	output,err:=conn.CreateUserWithContext(ctx,input)	//Somepartitions(e.g.ISO)maynotsupporttag-on-create.
	ifinput.Tags!=nil&&errs.IsUnsupportedOperationInPartitionError(conn.PartitionID,err){
		input.Tags=nil		output,err=conn.CreateUserWithContext(ctx,input)
	}	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"creatingElastiCacheUser(%s):%s",userID,err)
	}	d.SetId(aws.StringValue(output.UserId))	if_,err:=waitUserCreated(ctx,conn,d.Id(),d.Timeout(schema.TimeoutCreate));err!=nil{
		returnsdkdiag.AppendErrorf(diags,"waitingforElastiCacheUser(%s)create:%s",d.Id(),err)
	}	//Forpartitionsnotsupportingtag-on-create,attempttagaftercreate.
	iftags:=getTagsIn(ctx);input.Tags==nil&&len(tags)>0{
		err:=createTags(ctx,conn,aws.StringValue(output.ARN),tags)		//Ifdefaulttagsonly,continue.Otherwise,error.
		ifv,ok:=d.GetOk(names.AttrTags);(!ok||len(v.(map[string]interface{}))==0)&&errs.IsUnsupportedOperationInPartitionError(conn.PartitionID,err){
			returnappend(diags,resourceUserRead(ctx,d,meta)...)
		}		iferr!=nil{
			returnsdkdiag.AppendErrorf(diags,"settingElastiCacheUser(%s)tags:%s",d.Id(),err)
		}
	}	returnappend(diags,resourceUserRead(ctx,d,meta)...)
}
funcresourceUserRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).ElastiCacheConn(ctx)
funcr,err:=FindUserByID(ctx,conn,d.Id())	if!d.IsNewResource()&&tfresource.NotFound(err){
		log.Printf("[WARN]ElastiCacheUser(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returndiags
	}	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"readingElastiCacheUser(%s):%s",d.Id(),err)
	}	d.Set("access_string",user.AccessString)
	d.Set("arn",user.ARN)
	ifv:=user.Authentication;v!=nil{
		authenticationMode:=map[string]interface{}{
			"passwords":d.Get("authentication_mode.0.passwords"),
			"password_count":aws.Int64Value(v.PasswordCount),
			"type":aws.StringValue(v.Type),
		}		iferr:=d.Set("authentication_mode",[]interface{}{authenticationMode});err!=nil{
			returnsdkdiag.AppendErrorf(diags,"settingauthentication_mode:%s",err)
		}
	}else{
		d.Set("authentication_mode",nil)
	}
	d.Set("engine",user.Engine)
	d.Set("user_id",user.UserId)
	d.Set("user_name",user.UserName)	returndiags
}
funcresourceUserUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).ElastiCacheConn(ctx)
funcput:=&elasticache.ModifyUserInput{
			UserId:aws.String(d.Id()),
		}		ifd.HasChange("access_string"){
			input.AccessString=aws.String(d.Get("access_string").(string))
		}		ifd.HasChange("authentication_mode"){
			ifv,ok:=d.GetOk("authentication_mode");ok&&len(v.([]interface{}))>0&&v.([]interface{})[0]!=nil{
				input.AuthenticationMode=expandAuthenticationMode(v.([]interface{})[0].(map[string]interface{}))
			}
		}		ifd.HasChange("no_password_required"){
			input.NoPasswordRequired=aws.Bool(d.Get("no_password_required").(bool))
		}		ifd.HasChange("passwords"){
			input.Passwords=flex.ExpandStringSet(d.Get("passwords").(*schema.Set))
		}		_,err:=conn.ModifyUserWithContext(ctx,input)		iferr!=nil{
			returnsdkdiag.AppendErrorf(diags,"updatingElastiCacheUser(%s):%s",d.Id(),err)
		}		if_,err:=waitUserUpdated(ctx,conn,d.Id(),d.Timeout(schema.TimeoutUpdate));err!=nil{
			returnsdkdiag.AppendErrorf(diags,"waitingforElastiCacheUser(%s)update:%s",d.Id(),err)
		}
	}	returnappend(diags,resourceUserRead(ctx,d,meta)...)
}
funcresourceUserDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).ElastiCacheConn(ctx)	log.Printf("[INFO]DeletingElastiCacheUser:%s",d.Id())
funcerId:aws.String(d.Id()),
	})	iftfawserr.ErrCodeEquals(err,elasticache.ErrCodeUserNotFoundFault){
		returndiags
	}	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"deletingElastiCacheUser(%s):%s",d.Id(),err)
	}	if_,err:=waitUserDeleted(ctx,conn,d.Id(),d.Timeout(schema.TimeoutDelete));err!=nil{
		returnsdkdiag.AppendErrorf(diags,"waitingforElastiCacheUser(%s)delete:%s",d.Id(),err)
	}	returndiags
}
funcFindUserByID(ctxcontext.Context,conn*elasticache.ElastiCache,idstring)(*elasticache.User,error){
	input:=&elasticache.DescribeUsersInput{
		UserId:aws.String(id),
	}	output,err:=conn.DescribeUsersWithContext(ctx,input)
functfawserr.ErrCodeEquals(err,elasticache.ErrCodeUserNotFoundFault){
		returnnil,&retry.NotFoundError{
			LastError:err,
			LastRequest:input,
		}
	}	iferr!=nil{
		returnnil,err
	}	ifoutput==nil||len(output.Users)==0||output.Users[0]==nil{
		returnnil,tfresource.NewEmptyResultError(input)
	}	ifcount:=len(output.Users);count>1{
		returnnil,tfresource.NewTooManyResultsError(count,input)
	}	returnoutput.Users[0],nil
}
funcstatusUser(ctxcontext.Context,conn*elasticache.ElastiCache,idstring)retry.StateRefreshFunc{
	return
func()(interface{},string,error){
		output,err:=FindUserByID(ctx,conn,id)		iftfresource.NotFound(err){
			returnnil,"",nil
		}
funcerr!=nil{
			retur
func		returnoutput,aws.StringValue(output.Status),nil
	}
}const(
	userStatusActive="active"
	userStatusCreating="creating"
	userStatusDeleting="deleting"
	userStatusModifying="modifying"
)
funcwaitUserCreated(ctxcontext.Context,conn*elasticache.ElastiCache,idstring,timeouttime.Duration)(*elasticache.User,error){
	stateConf:=&retry.StateChangeConf{
		Pending:[]string{userStatusCreating},
		Target:[]string{userStatusActive},
		Refresh:statusUser(ctx,conn,id),
		Timeout:timeout,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)
funcoutput,ok:=outputRaw.(*elasticache.User);ok{
		returnoutput,err
	}	returnnil,err
}
funcwaitUserUpdated(ctxcontext.Context,conn*elasticache.ElastiCache,idstring,timeouttime.Duration)(*elasticache.User,error){
	stateConf:=&retry.StateChangeConf{
		Pending:[]string{userStatusModifying},
		Target:[]string{userStatusActive},
		Refresh:statusUser(ctx,conn,id),
		Timeout:timeout,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)
functurnoutput,err
	}	returnnil,err
}
funcwaitUserDeleted(ctxcontext.Context,conn*elasticache.ElastiCache,idstring,timeouttime.Duration)(*elasticache.User,error){
	stateConf:=&retry.StateChangeConf{
		Pending:[]string{userStatusDeleting},
		Target:[]string{},
		Refresh:statusUser(ctx,conn,id),
		Timeout:timeout,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)	ifoutput,ok:=outputRaw.(*elasticache.User);ok{
func	returnnil,err
}
funcexpandAuthenticationMode(tfMapmap[string]interface{})*elasticache.AuthenticationMode{
	iftfMap==nil{
		returnnil
	}	apiObject:=&elasticache.AuthenticationMode{}	ifv,ok:=tfMap["passwords"].(*schema.Set);ok&&v.Len()>0{
		apiObject.Passwords=flex.ExpandStringSet(v)
	}	ifv,ok:=tfMap["type"].(string);ok&&v!=""{
func	returnapiObject
}
