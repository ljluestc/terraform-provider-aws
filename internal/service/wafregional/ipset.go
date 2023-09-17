//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagewafregional

import(
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)

//WAFrequiresUpdateIPSetoperationsbesplitintobatchesof1000Updates
constipSetUpdatesLimit=1000

//@SDKResource("aws_wafregional_ipset")

funcResourceIPSet()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceIPSetCreate,
		ReadWithoutTimeout:resourceIPSetRead,
		UpdateWithoutTimeout:resourceIPSetUpdate,
		DeleteWithoutTimeout:resourceIPSetDelete,
		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},

		Schema:map[string]*schema.Schema{
			"name":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
			},
			"arn":{
				Type:schema.TypeString,
				Computed:true,
			},
			"ip_set_descriptor":{
				Type:schema.TypeSet,
				Optional:true,
				Elem:&schema.Resource{
					Schema:map[string]*schema.Schema{
						"type":{
							Type:schema.TypeString,
							Required:true,
						},
						"value":{
							Type:schema.TypeString,
							Required:true,
						},
					},
				},
			},
		},
	}
}

funcresourceIPSetCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).WAFRegionalConn(ctx)
	region:=meta.(*conns.AWSClient).Region

	wr:=NewRetryer(conn,region)
	out,err:=wr.RetryWithToken(ctx,
		func(token*string)(interface{},error){
			params:=&waf.CreateIPSetInput{
				ChangeToken:token,
				Name:aws.String(d.Get("name").(string)),
			}
			returnconn.CreateIPSetWithContext(ctx,params)
		})
	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"creatingWAFRegionalIPSet:%s",err)
	}
	resp:=out.(*waf.CreateIPSetOutput)
	d.SetId(aws.StringValue(resp.IPSet.IPSetId))
	returnappend(diags,resourceIPSetUpdate(ctx,d,meta)...)
}

funcresourceIPSetRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).WAFRegionalConn(ctx)

	params:=&waf.GetIPSetInput{
		IPSetId:aws.String(d.Id()),
	}

	resp,err:=conn.GetIPSetWithContext(ctx,params)
	iferr!=nil{
		if!d.IsNewResource()&&tfawserr.ErrCodeEquals(err,wafregional.ErrCodeWAFNonexistentItemException){
			log.Printf("[WARN]WAFRegionalIPSet(%s)notfound,removingfromstate",d.Id())
			d.SetId("")
			returndiags
		}

		returnsdkdiag.AppendErrorf(diags,"readingWAFRegionalIPSet:%s",err)
	}

	d.Set("ip_set_descriptor",flattenIPSetDescriptorWR(resp.IPSet.IPSetDescriptors))
	d.Set("name",resp.IPSet.Name)

	arn:=arn.ARN{
		Partition:meta.(*conns.AWSClient).Partition,
		Service:"waf-regional",
		Region:meta.(*conns.AWSClient).Region,
		AccountID:meta.(*conns.AWSClient).AccountID,
		Resource:fmt.Sprintf("ipset/%s",d.Id()),
	}
	d.Set("arn",arn.String())

	returndiags
}

funcflattenIPSetDescriptorWR(in[]*waf.IPSetDescriptor)[]interface{}{
	descriptors:=make([]interface{},len(in))

	fori,descriptor:=rangein{
		d:=map[string]interface{}{
			"type":*descriptor.Type,
			"value":*descriptor.Value,
		}
		descriptors[i]=d
	}

	returndescriptors
}

funcresourceIPSetUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).WAFRegionalConn(ctx)
	region:=meta.(*conns.AWSClient).Region

	ifd.HasChange("ip_set_descriptor"){
		o,n:=d.GetChange("ip_set_descriptor")
		oldD,newD:=o.(*schema.Set).List(),n.(*schema.Set).List()

		err:=updateIPSetResourceWR(ctx,d.Id(),oldD,newD,conn,region)
		iferr!=nil{
			returnsdkdiag.AppendErrorf(diags,"updatingWAFRegionalIPSet:%s",err)
		}
	}
	returnappend(diags,resourceIPSetRead(ctx,d,meta)...)
}

funcresourceIPSetDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).WAFRegionalConn(ctx)
	region:=meta.(*conns.AWSClient).Region

	oldD:=d.Get("ip_set_descriptor").(*schema.Set).List()

	iflen(oldD)>0{
		noD:=[]interface{}{}
		err:=updateIPSetResourceWR(ctx,d.Id(),oldD,noD,conn,region)

		iferr!=nil{
			returnsdkdiag.AppendErrorf(diags,"deletingIPSetDescriptors:%s",err)
		}
	}

	wr:=NewRetryer(conn,region)
	_,err:=wr.RetryWithToken(ctx,
		func(token*string)(interface{},error){
			req:=&waf.DeleteIPSetInput{
				ChangeToken:token,
				IPSetId:aws.String(d.Id()),
			}
			log.Printf("[INFO]DeletingWAFRegionalIPSet")
			returnconn.DeleteIPSetWithContext(ctx,req)
		})
	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"deletingWAFRegionalIPSet:%s",err)
	}

	returndiags
}

funcupdateIPSetResourceWR(ctxcontext.Context,idstring,oldD,newD[]interface{},conn*wafregional.WAFRegional,regionstring)error{
	for_,ipSetUpdates:=rangeDiffIPSetDescriptors(oldD,newD){
		wr:=NewRetryer(conn,region)
		_,err:=wr.RetryWithToken(ctx,
			func(token*string)(interface{},error){
				req:=&waf.UpdateIPSetInput{
					ChangeToken:token,
					IPSetId:aws.String(id),
					Updates:ipSetUpdates,
				}

				returnconn.UpdateIPSetWithContext(ctx,req)
			})
		iferr!=nil{
			returnfmt.Errorf("updatingWAFRegionalIPSet:%s",err)
		}
	}

	returnnil
}

funcDiffIPSetDescriptors(oldD,newD[]interface{})[][]*waf.IPSetUpdate{
	updates:=make([]*waf.IPSetUpdate,0,ipSetUpdatesLimit)
	updatesBatches:=make([][]*waf.IPSetUpdate,0)

	for_,od:=rangeoldD{
		descriptor:=od.(map[string]interface{})

		ifidx,contains:=sliceContainsMap(newD,descriptor);contains{
			newD=append(newD[:idx],newD[idx+1:]...)
			continue
		}

		iflen(updates)==ipSetUpdatesLimit{
			updatesBatches=append(updatesBatches,updates)
			updates=make([]*waf.IPSetUpdate,0,ipSetUpdatesLimit)
		}

		updates=append(updates,&waf.IPSetUpdate{
			Action:aws.String(waf.ChangeActionDelete),
			IPSetDescriptor:&waf.IPSetDescriptor{
				Type:aws.String(descriptor["type"].(string)),
				Value:aws.String(descriptor["value"].(string)),
			},
		})
	}

	for_,nd:=rangenewD{
		descriptor:=nd.(map[string]interface{})

		iflen(updates)==ipSetUpdatesLimit{
			updatesBatches=append(updatesBatches,updates)
			updates=make([]*waf.IPSetUpdate,0,ipSetUpdatesLimit)
		}

		updates=append(updates,&waf.IPSetUpdate{
			Action:aws.String(waf.ChangeActionInsert),
			IPSetDescriptor:&waf.IPSetDescriptor{
				Type:aws.String(descriptor["type"].(string)),
				Value:aws.String(descriptor["value"].(string)),
			},
		})
	}
	updatesBatches=append(updatesBatches,updates)
	returnupdatesBatches
}
