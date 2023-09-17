//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packageresourcegroupsimport(
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/slices"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)//@SDKResource("aws_resourcegroups_resource",name="Resource")
funcResourceResource()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceResourceCreate,
		ReadWithoutTimeout:resourceResourceRead,
		DeleteWithoutTimeout:resourceResourceDelete,		Timeouts:&schema.ResourceTimeout{
			Create:schema.DefaultTimeout(5*time.Minute),
			Delete:schema.DefaultTimeout(5*time.Minute),
		},		Schema:map[string]*schema.Schema{
			"group_arn":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
			},
			"resource_arn":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
			},
			"resource_type":{
				Type:schema.TypeString,
				Computed:true,
			},
		},
	}
}funcresourceResourceCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).ResourceGroupsConn(ctx)	groupARN:=d.Get("group_arn").(string)
	resourceARN:=d.Get("resource_arn").(string)
	id:=strings.Join([]string{strings.Split(strings.ToLower(groupARN),"/")[1],strings.Split(resourceARN,"/")[1]},"_")
	input:=&resourcegroups.GroupResourcesInput{
		Group:aws.String(groupARN),
		ResourceArns:aws.StringSlice([]string{resourceARN}),
	}	output,err:=conn.GroupResourcesWithContext(ctx,input)	iferr==nil{
		err=FailedResourcesError(output.Failed)
	}	iferr!=nil{
		returndiag.Errorf("creatingResourceGroupsResource(%s):%s",id,err)
	}	d.SetId(id)	if_,err:=waitResourceCreated(ctx,conn,groupARN,resourceARN,d.Timeout(schema.TimeoutDelete));err!=nil{
		returndiag.Errorf("waitingforResourceGroupsResource(%s)create:%s",d.Id(),err)
	}	returnresourceResourceRead(ctx,d,meta)
}funcresourceResourceRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).ResourceGroupsConn(ctx)	output,err:=FindResourceByTwoPartKey(ctx,conn,d.Get("group_arn").(string),d.Get("resource_arn").(string))	if!d.IsNewResource()&&tfresource.NotFound(err){
		log.Printf("[WARN]ResourceGroupsResource(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returnnil
	}	iferr!=nil{
		returndiag.Errorf("readingResourceGroupsResource(%s):%s",d.Id(),err)
	}	d.Set("resource_arn",output.Identifier.ResourceArn)
	d.Set("resource_type",output.Identifier.ResourceType)	returnnil
}funcresourceResourceDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).ResourceGroupsConn(ctx)	groupARN:=d.Get("group_arn").(string)
	resourceARN:=d.Get("resource_arn").(string)
	log.Printf("[INFO]DeletingResourceGroupsResource:%s",d.Id())
	output,err:=conn.UngroupResourcesWithContext(ctx,&resourcegroups.UngroupResourcesInput{
		Group:aws.String(groupARN),
		ResourceArns:aws.StringSlice([]string{resourceARN}),
	})	iferr==nil{
		err=FailedResourcesError(output.Failed)
	}	iferr!=nil{
		returndiag.Errorf("deletingResourceGroupsResource(%s):%s",d.Id(),err)
	}	if_,err:=waitResourceDeleted(ctx,conn,groupARN,resourceARN,d.Timeout(schema.TimeoutDelete));err!=nil{
		returndiag.Errorf("waitingforResourceGroupsResource(%s)delete:%s",d.Id(),err)
	}	returnnil
}funcFindResourceByTwoPartKey(ctxcontext.Context,conn*resourcegroups.ResourceGroups,groupARN,resourceARNstring)(*resourcegroups.ListGroupResourcesItem,error){
	input:=&resourcegroups.ListGroupResourcesInput{
		Group:aws.String(groupARN),
	}
	varoutput[]*resourcegroups.ListGroupResourcesItem	err:=conn.ListGroupResourcesPagesWithContext(ctx,input,func(page*resourcegroups.ListGroupResourcesOutput,lastPagebool)bool{
		ifpage==nil{
			return!lastPage
		}		output=append(output,page.Resources...)		return!lastPage
	})	iftfawserr.ErrCodeEquals(err,resourcegroups.ErrCodeNotFoundException){
		returnnil,&retry.NotFoundError{
			LastError:err,
			LastRequest:input,
		}
	}	iferr!=nil{
		returnnil,err
	}	output=slices.Filter(output,func(v*resourcegroups.ListGroupResourcesItem)bool{
		returnv.Identifier!=nil&&aws.StringValue(v.Identifier.ResourceArn)==resourceARN
	})	returntfresource.AssertSinglePtrResult(output)
}funcstatusResource(ctxcontext.Context,conn*resourcegroups.ResourceGroups,groupARN,resourceARNstring)retry.StateRefreshFunc{
	returnfunc()(interface{},string,error){
		output,err:=FindResourceByTwoPartKey(ctx,conn,groupARN,resourceARN)		iftfresource.NotFound(err){
			returnnil,"",nil
		}		iferr!=nil{
			returnnil,"",err
		}		ifoutput.Status==nil{
			returnoutput,"",nil
		}		returnoutput,aws.StringValue(output.Status.Name),nil
	}
}funcwaitResourceCreated(ctxcontext.Context,conn*resourcegroups.ResourceGroups,groupARN,resourceARNstring,timeouttime.Duration)(*resourcegroups.ListGroupResourcesItem,error){
	stateConf:=&retry.StateChangeConf{
		Pending:[]string{resourcegroups.ResourceStatusValuePending},
		Target:[]string{""},
		Refresh:statusResource(ctx,conn,groupARN,resourceARN),
		Timeout:timeout,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)	ifoutput,ok:=outputRaw.(*resourcegroups.ListGroupResourcesItem);ok{
		returnoutput,err
	}	returnnil,err
}funcwaitResourceDeleted(ctxcontext.Context,conn*resourcegroups.ResourceGroups,groupARN,resourceARNstring,timeouttime.Duration)(*resourcegroups.ListGroupResourcesItem,error){
	stateConf:=&retry.StateChangeConf{
		Pending:[]string{resourcegroups.ResourceStatusValuePending},
		Target:[]string{},
		Refresh:statusResource(ctx,conn,groupARN,resourceARN),
		Timeout:timeout,
	}	outputRaw,err:=stateConf.WaitForStateContext(ctx)	ifoutput,ok:=outputRaw.(*resourcegroups.ListGroupResourcesItem);ok{
		returnoutput,err
	}	returnnil,err
}funcFailedResourceError(apiObject*resourcegroups.FailedResource)error{
	ifapiObject==nil{
		returnnil
	}	returnawserr.New(aws.StringValue(apiObject.ErrorCode),aws.StringValue(apiObject.ErrorMessage),nil)
}funcFailedResourcesError(apiObjects[]*resourcegroups.FailedResource)error{
	varerrs[]error	for_,apiObject:=rangeapiObjects{
		ifapiObject==nil{
			continue
		}		err:=FailedResourceError(apiObject)		iferr!=nil{
			errs=append(errs,fmt.Errorf("%s:%w",aws.StringValue(apiObject.ResourceArn),err))
		}
	}	returnerrors.Join(errs...)
}
