//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packageredshiftserverlessimport(
	"context"
	"log"
	"strings"
	"time"	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/redshiftserverless"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)//@SDKResource("aws_redshiftserverless_namespace",name="Namespace")
//@Tags(identifierAttribute="arn")
funcResourceNamespace()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceNamespaceCreate,
		ReadWithoutTimeout:resourceNamespaceRead,
		UpdateWithoutTimeout:resourceNamespaceUpdate,
		DeleteWithoutTimeout:resourceNamespaceDelete,		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},		Schema:map[string]*schema.Schema{
			"admin_user_password":{
				Type:schema.TypeString,
				Optional:true,
				Sensitive:true,
			},
			"admin_username":{
				Type:schema.TypeString,
				Optional:true,
				Sensitive:true,
				Computed:true,
			},
			"arn":{
				Type:schema.TypeString,
				Computed:true,
			},
			"db_name":{
				Type:schema.TypeString,
				Optional:true,
				ForceNew:true,
				Computed:true,
			},
			"default_iam_role_arn":{
				Type:schema.TypeString,
				Optional:true,
				ValidateFunc:verify.ValidARN,
			},
			"iam_roles":{
				Type:schema.TypeSet,
				Optional:true,
				Computed:true,
				Elem:&schema.Schema{
					Type:schema.TypeString,
					ValidateFunc:verify.ValidARN,
				},
			},
			"kms_key_id":{
				Type:schema.TypeString,
				Optional:true,
				Computed:true,
				ValidateFunc:verify.ValidARN,
			},
			"log_exports":{
				Type:schema.TypeSet,
				Optional:true,
				Elem:&schema.Schema{
					Type:schema.TypeString,
					ValidateFunc:validation.StringInSlice(redshiftserverless.LogExport_Values(),false),
				},
			},
			"namespace_id":{
				Type:schema.TypeString,
				Computed:true,
			},
			"namespace_name":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
			},
			names.AttrTags:tftags.TagsSchema(),
			names.AttrTagsAll:tftags.TagsSchemaComputed(),
		},		CustomizeDiff:verify.SetTagsDiff,
	}
}funcresourceNamespaceCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).RedshiftServerlessConn(ctx)	name:=d.Get("namespace_name").(string)
	input:=&redshiftserverless.CreateNamespaceInput{
		NamespaceName:aws.String(name),
		Tags:getTagsIn(ctx),
	}	ifv,ok:=d.GetOk("admin_user_password");ok{
		input.AdminUserPassword=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("admin_username");ok{
		input.AdminUsername=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("db_name");ok{
		input.DbName=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("default_iam_role_arn");ok{
		input.DefaultIamRoleArn=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("iam_roles");ok&&v.(*schema.Set).Len()>0{
		input.IamRoles=flex.ExpandStringSet(v.(*schema.Set))
	}	ifv,ok:=d.GetOk("kms_key_id");ok{
		input.KmsKeyId=aws.String(v.(string))
	}	ifv,ok:=d.GetOk("log_exports");ok&&v.(*schema.Set).Len()>0{
		input.LogExports=flex.ExpandStringSet(v.(*schema.Set))
	}	output,err:=conn.CreateNamespaceWithContext(ctx,input)	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"creatingRedshiftServerlessNamespace(%s):%s",name,err)
	}	d.SetId(aws.StringValue(output.Namespace.NamespaceName))	returnappend(diags,resourceNamespaceRead(ctx,d,meta)...)
}funcresourceNamespaceRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).RedshiftServerlessConn(ctx)	output,err:=FindNamespaceByName(ctx,conn,d.Id())	if!d.IsNewResource()&&tfresource.NotFound(err){
		log.Printf("[WARN]RedshiftServerlessNamespace(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returndiags
	}	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"readingRedshiftServerlessNamespace(%s):%s",d.Id(),err)
	}	arn:=aws.StringValue(output.NamespaceArn)
	d.Set("admin_username",output.AdminUsername)
	d.Set("arn",arn)
	d.Set("db_name",output.DbName)
	d.Set("default_iam_role_arn",output.DefaultIamRoleArn)
	d.Set("iam_roles",flattenNamespaceIAMRoles(output.IamRoles))
	d.Set("kms_key_id",output.KmsKeyId)
	d.Set("log_exports",aws.StringValueSlice(output.LogExports))
	d.Set("namespace_id",output.NamespaceId)
	d.Set("namespace_name",output.NamespaceName)	returndiags
}funcresourceNamespaceUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).RedshiftServerlessConn(ctx)	ifd.HasChangesExcept("tags","tags_all"){
		input:=&redshiftserverless.UpdateNamespaceInput{
			NamespaceName:aws.String(d.Id()),
		}		ifd.HasChanges("admin_username","admin_user_password"){
			input.AdminUsername=aws.String(d.Get("admin_username").(string))
			input.AdminUserPassword=aws.String(d.Get("admin_user_password").(string))
		}		ifd.HasChange("default_iam_role_arn"){
			input.DefaultIamRoleArn=aws.String(d.Get("default_iam_role_arn").(string))
		}		ifd.HasChange("iam_roles"){
			input.IamRoles=flex.ExpandStringSet(d.Get("iam_roles").(*schema.Set))
		}		ifd.HasChange("kms_key_id"){
			input.KmsKeyId=aws.String(d.Get("kms_key_id").(string))
		}		ifd.HasChange("log_exports"){
			input.LogExports=flex.ExpandStringSet(d.Get("log_exports").(*schema.Set))
		}		_,err:=conn.UpdateNamespaceWithContext(ctx,input)		iferr!=nil{
			returnsdkdiag.AppendErrorf(diags,"updatingRedshiftServerlessNamespace(%s):%s",d.Id(),err)
		}		if_,err:=waitNamespaceUpdated(ctx,conn,d.Id());err!=nil{
			returnsdkdiag.AppendErrorf(diags,"waitingforRedshiftServerlessNamespace(%s)update:%s",d.Id(),err)
		}
	}	returnappend(diags,resourceNamespaceRead(ctx,d,meta)...)
}funcresourceNamespaceDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).RedshiftServerlessConn(ctx)	log.Printf("[DEBUG]DeletingRedshiftServerlessNamespace:%s",d.Id())
	_,err:=tfresource.RetryWhenAWSErrMessageContains(ctx,10*time.Minute,
		func()(interface{},error){
			returnconn.DeleteNamespaceWithContext(ctx,&redshiftserverless.DeleteNamespaceInput{
				NamespaceName:aws.String(d.Id()),
			})
		},
		//"ConflictException:Thereisanoperationrunningonthenamespace.Trydeletingthenamespaceagainlater."
		redshiftserverless.ErrCodeConflictException,"operationrunning")	iftfawserr.ErrCodeEquals(err,redshiftserverless.ErrCodeResourceNotFoundException){
		returndiags
	}	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"deletingRedshiftServerlessNamespace(%s):%s",d.Id(),err)
	}	if_,err:=waitNamespaceDeleted(ctx,conn,d.Id());err!=nil{
		returnsdkdiag.AppendErrorf(diags,"waitingforRedshiftServerlessNamespace(%s)delete:%s",d.Id(),err)
	}	returndiags
}var(
	reIAMRole=regexache.MustCompile(`^\s*IamRole\((.*)\)\s*$`)
)funcflattenNamespaceIAMRoles(iamRoles[]*string)[]string{
	vartfList[]string	for_,iamRole:=rangeiamRoles{
		iamRole:=aws.StringValue(iamRole)		ifarn.IsARN(iamRole){
			tfList=append(tfList,iamRole)
			continue
		}		//e.g."IamRole(applyStatus=in-sync,iamRoleArn=arn:aws:iam::123456789012:role/service-role/test)"
		ifm:=reIAMRole.FindStringSubmatch(iamRole);len(m)>0{
			varkeystring
			s:=m[1]
			fors!=""{
				key,s,_=strings.Cut(s,",")
				key=strings.TrimSpace(key)
				ifkey==""{
					continue
				}
				key,value,_:=strings.Cut(key,"=")
				ifkey=="iamRoleArn"{
					tfList=append(tfList,value)
					break
				}
			}			continue
		}
	}	returntfList
}
