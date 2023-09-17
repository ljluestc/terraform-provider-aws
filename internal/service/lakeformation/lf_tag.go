//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagelakeformation

import(
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
)

//ThisvalueisdefinedbyAWSAPI
constlfTagsValuesMaxBatchSize=50

//@SDKResource("aws_lakeformation_lf_tag")
funcResourceLFTag()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceLFTagCreate,
		ReadWithoutTimeout:resourceLFTagRead,
		UpdateWithoutTimeout:resourceLFTagUpdate,
		DeleteWithoutTimeout:resourceLFTagDelete,
		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},

		Schema:map[string]*schema.Schema{
			"catalog_id":{
				Type:schema.TypeString,
				ForceNew:true,
				Optional:true,
				Computed:true,
			},
			"key":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
				ValidateFunc:validation.StringLenBetween(1,128),
			},
			"values":{
				Type:schema.TypeSet,
				Required:true,
				MinItems:1,
				//SoftlimitstatedinAWSDoc
				//https://docs.aws.amazon.com/lake-formation/latest/dg/TBAC-notes.html
				MaxItems:1000,
				Elem:&schema.Schema{
					Type:schema.TypeString,
					ValidateFunc:validateLFTagValues(),
				},
				Set:schema.HashString,
			},
		},
	}
}

funcresourceLFTagCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).LakeFormationConn(ctx)

	tagKey:=d.Get("key").(string)
	tagValues:=d.Get("values").(*schema.Set)

	varcatalogIDstring
	ifv,ok:=d.GetOk("catalog_id");ok{
		catalogID=v.(string)
	}else{
		catalogID=meta.(*conns.AWSClient).AccountID
	}

	tagValueChunks:=splitLFTagValues(tagValues.List(),lfTagsValuesMaxBatchSize)

	input:=&lakeformation.CreateLFTagInput{
		CatalogId:aws.String(catalogID),
		TagKey:aws.String(tagKey),
		TagValues:flex.ExpandStringList(tagValueChunks[0]),
	}

	_,err:=conn.CreateLFTagWithContext(ctx,input)
	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"creatingLakeFormationLF-Tag:%s",err)
	}

	iflen(tagValueChunks)>1{
		tagValueChunks=tagValueChunks[1:]

		for_,v:=rangetagValueChunks{
			in:=&lakeformation.UpdateLFTagInput{
				CatalogId:aws.String(catalogID),
				TagKey:aws.String(tagKey),
				TagValuesToAdd:flex.ExpandStringList(v),
			}

			_,err:=conn.UpdateLFTagWithContext(ctx,in)
			iferr!=nil{
				returnsdkdiag.AppendErrorf(diags,"creatingLakeFormationLF-Tag:%s",err)
			}
		}
	}

	d.SetId(fmt.Sprintf("%s:%s",catalogID,tagKey))

	returnappend(diags,resourceLFTagRead(ctx,d,meta)...)
}

funcresourceLFTagRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).LakeFormationConn(ctx)

	catalogID,tagKey,err:=ReadLFTagID(d.Id())
	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"readingLakeFormationLF-Tag(%s):%s",d.Id(),err)
	}

	input:=&lakeformation.GetLFTagInput{
		CatalogId:aws.String(catalogID),
		TagKey:aws.String(tagKey),
	}

	output,err:=conn.GetLFTagWithContext(ctx,input)
	if!d.IsNewResource(){
		iftfawserr.ErrCodeEquals(err,lakeformation.ErrCodeEntityNotFoundException){
			log.Printf("[WARN]LakeFormationLF-Tag(%s)notfound,removingfromstate",d.Id())
			d.SetId("")
			returndiags
		}
	}

	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"readingLakeFormationLF-Tag(%s):%s",d.Id(),err)
	}

	d.Set("key",output.TagKey)
	d.Set("values",flex.FlattenStringSet(output.TagValues))
	d.Set("catalog_id",output.CatalogId)

	returndiags
}

funcresourceLFTagUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).LakeFormationConn(ctx)

	catalogID,tagKey,err:=ReadLFTagID(d.Id())
	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"updatingLakeFormationLF-Tag(%s):%s",d.Id(),err)
	}

	o,n:=d.GetChange("values")
	os:=o.(*schema.Set)
	ns:=n.(*schema.Set)
	toAdd:=ns.Difference(os)
	toDelete:=os.Difference(ns)

	vartoAddChunks,toDeleteChunks[][]interface{}
	iflen(toAdd.List())>0{
		toAddChunks=splitLFTagValues(toAdd.List(),lfTagsValuesMaxBatchSize)
	}

	iflen(toDelete.List())>0{
		toDeleteChunks=splitLFTagValues(toDelete.List(),lfTagsValuesMaxBatchSize)
	}

	for{
		iflen(toAddChunks)==0&&len(toDeleteChunks)==0{
			break
		}

		input:=&lakeformation.UpdateLFTagInput{
			CatalogId:aws.String(catalogID),
			TagKey:aws.String(tagKey),
		}

		toAddEnd,toDeleteEnd:=len(toAddChunks),len(toDeleteChunks)
		varindexAdd,indexDeleteint
		ifindexAdd<toAddEnd{
			input.TagValuesToAdd=flex.ExpandStringList(toAddChunks[indexAdd])
			indexAdd++
		}
		ifindexDelete<toDeleteEnd{
			input.TagValuesToDelete=flex.ExpandStringList(toDeleteChunks[indexDelete])
			indexDelete++
		}

		toAddChunks=toAddChunks[indexAdd:]
		toDeleteChunks=toDeleteChunks[indexDelete:]

		_,err=conn.UpdateLFTagWithContext(ctx,input)
		iferr!=nil{
			returnsdkdiag.AppendErrorf(diags,"updatingLakeFormationLF-Tag(%s):%s",d.Id(),err)
		}
	}

	returnappend(diags,resourceLFTagRead(ctx,d,meta)...)
}

funcresourceLFTagDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	vardiagsdiag.Diagnostics
	conn:=meta.(*conns.AWSClient).LakeFormationConn(ctx)

	catalogID,tagKey,err:=ReadLFTagID(d.Id())
	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"deletingLakeFormationLF-Tag(%s):%s",d.Id(),err)
	}

	input:=&lakeformation.DeleteLFTagInput{
		CatalogId:aws.String(catalogID),
		TagKey:aws.String(tagKey),
	}

	_,err=conn.DeleteLFTagWithContext(ctx,input)
	iferr!=nil{
		returnsdkdiag.AppendErrorf(diags,"deletingLakeFormationLF-Tag(%s):%s",d.Id(),err)
	}

	returndiags
}

funcReadLFTagID(idstring)(string,string,error){
	catalogID,tagKey,found:=strings.Cut(id,":")

	if!found{
		return"","",fmt.Errorf("unexpectedformatofID(%q),expectedCATALOG-ID:TAG-KEY",id)
	}

	returncatalogID,tagKey,nil
}

funcvalidateLFTagValues()schema.SchemaValidateFunc{
	returnvalidation.All(
		validation.StringLenBetween(1,255),
		validation.StringMatch(regexache.MustCompile(`^([\p{L}\p{Z}\p{N}_.:\*\/=+\-@%]*)$`),""),
	)
}

funcsplitLFTagValues(in[]interface{},sizeint)[][]interface{}{
	varout[][]interface{}

	for{
		iflen(in)==0{
			break
		}

		iflen(in)<size{
			size=len(in)
		}

		out=append(out,in[0:size])
		in=in[size:]
	}

	returnout
}
