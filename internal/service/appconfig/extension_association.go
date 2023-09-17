//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageappconfig

import(
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appconfig"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

const(
	ResExtensionAssociation="ExtensionAssociation"
)

//@SDKResource("aws_appconfig_extension_association")
funcResourceExtensionAssociation()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceExtensionAssociationCreate,
		ReadWithoutTimeout:resourceExtensionAssociationRead,
		UpdateWithoutTimeout:resourceExtensionAssociationUpdate,
		DeleteWithoutTimeout:resourceExtensionAssociationDelete,

		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},

		Schema:map[string]*schema.Schema{
			"arn":{
				Type:schema.TypeString,
				Computed:true,
			},
			"extension_arn":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
			},
			"parameters":{
				Type:schema.TypeMap,
				Optional:true,
				Elem:&schema.Schema{Type:schema.TypeString},
			},
			"resource_arn":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
			},
			"extension_version":{
				Type:schema.TypeInt,
				Computed:true,
			},
		},
	}
}

funcresourceExtensionAssociationCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).AppConfigConn(ctx)

	in:=appconfig.CreateExtensionAssociationInput{
		ExtensionIdentifier:aws.String(d.Get("extension_arn").(string)),
		ResourceIdentifier:aws.String(d.Get("resource_arn").(string)),
	}

	ifv,ok:=d.GetOk("parameters");ok{
		in.Parameters=flex.ExpandStringMap(v.(map[string]interface{}))
	}

	out,err:=conn.CreateExtensionAssociationWithContext(ctx,&in)

	iferr!=nil{
		returncreate.DiagError(names.AppConfig,create.ErrActionCreating,ResExtensionAssociation,d.Get("extension_arn").(string),err)
	}

	ifout==nil{
		returncreate.DiagError(names.AppConfig,create.ErrActionCreating,ResExtensionAssociation,d.Get("extension_arn").(string),errors.New("NoExtensionAssociationreturnedwithcreaterequest."))
	}

	d.SetId(aws.StringValue(out.Id))

	returnresourceExtensionAssociationRead(ctx,d,meta)
}

funcresourceExtensionAssociationRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).AppConfigConn(ctx)

	out,err:=FindExtensionAssociationById(ctx,conn,d.Id())

	if!d.IsNewResource()&&tfresource.NotFound(err){
		create.LogNotFoundRemoveState(names.AppConfig,create.ErrActionReading,ResExtensionAssociation,d.Id())
		d.SetId("")
		returnnil
	}

	iferr!=nil{
		returncreate.DiagError(names.AppConfig,create.ErrActionReading,ResExtensionAssociation,d.Id(),err)
	}

	d.Set("arn",out.Arn)
	d.Set("extension_arn",out.ExtensionArn)
	d.Set("parameters",out.Parameters)
	d.Set("resource_arn",out.ResourceArn)
	d.Set("extension_version",out.ExtensionVersionNumber)

	returnnil
}

funcresourceExtensionAssociationUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).AppConfigConn(ctx)
	requestUpdate:=false

	in:=&appconfig.UpdateExtensionAssociationInput{
		ExtensionAssociationId:aws.String(d.Id()),
	}

	ifd.HasChange("parameters"){
		in.Parameters=flex.ExpandStringMap(d.Get("parameters").(map[string]interface{}))
		requestUpdate=true
	}

	ifrequestUpdate{
		out,err:=conn.UpdateExtensionAssociationWithContext(ctx,in)

		iferr!=nil{
			returncreate.DiagError(names.AppConfig,create.ErrActionWaitingForUpdate,ResExtensionAssociation,d.Id(),err)
		}

		ifout==nil{
			returncreate.DiagError(names.AppConfig,create.ErrActionWaitingForUpdate,ResExtensionAssociation,d.Id(),errors.New("NoExtensionAssociationreturnedwithupdaterequest."))
		}
	}

	returnresourceExtensionAssociationRead(ctx,d,meta)
}

funcresourceExtensionAssociationDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).AppConfigConn(ctx)

	_,err:=conn.DeleteExtensionAssociationWithContext(ctx,&appconfig.DeleteExtensionAssociationInput{
		ExtensionAssociationId:aws.String(d.Id()),
	})

	iferr!=nil{
		returncreate.DiagError(names.AppConfig,create.ErrActionDeleting,ResExtensionAssociation,d.Id(),err)
	}

	returnnil
}
