//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageprovider

import(
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
)

typemockServicestruct{}

func(t*mockService)FrameworkDataSources(ctxcontext.Context)[]*types.ServicePackageFrameworkDataSource{
	return[]*types.ServicePackageFrameworkDataSource{}
}

func(t*mockService)FrameworkResources(ctxcontext.Context)[]*types.ServicePackageFrameworkResource{
	return[]*types.ServicePackageFrameworkResource{}
}

func(t*mockService)SDKDataSources(ctxcontext.Context)[]*types.ServicePackageSDKDataSource{
	return[]*types.ServicePackageSDKDataSource{}
}

func(t*mockService)SDKResources(ctxcontext.Context)[]*types.ServicePackageSDKResource{
	return[]*types.ServicePackageSDKResource{}
}

func(t*mockService)ServicePackageName()string{
	return"TestService"
}

func(t*mockService)ListTags(ctxcontext.Context,metaany,identifierstring)error{
	tags:=tftags.New(ctx,map[string]string{
		"tag1":"value1",
	})
	ifinContext,ok:=tftags.FromContext(ctx);ok{
		inContext.TagsOut=types.Some(tags)
	}

	returnerrors.New("testerror")
}

func(t*mockService)UpdateTags(context.Context,any,string,string,any)error{
	returnnil
}

funcTestTagsResourceInterceptor(t*testing.T){
	t.Parallel()

	varinterceptorsinterceptorItems

	sp:=&types.ServicePackageResourceTags{
		IdentifierAttribute:"id",
	}

	tags:=tagsResourceInterceptor{
		tags:sp,
		updateFunc:tagsUpdateFunc,
		readFunc:tagsReadFunc,
	}

	interceptors=append(interceptors,interceptorItem{
		when:Finally,
		why:Update,
		interceptor:tags,
	})

	conn:=&conns.AWSClient{
		ServicePackages:map[string]conns.ServicePackage{
"Test":&mockService{},
		},
		DefaultTagsConfig:expandDefaultTags(context.Background(),map[string]interface{}{
"tag":"",
		}),
		IgnoreTagsConfig:expandIgnoreTags(context.Background(),map[string]interface{}{
"tag2":"tag",
		}),
	}

	bootstrapContext:=func(ctxcontext.Context,metaany)context.Context{
		ctx=conns.NewResourceContext(ctx,"Test","aws_test")
		ifv,ok:=meta.(*conns.AWSClient);ok{
ctx=tftags.NewContext(ctx,v.DefaultTagsConfig,v.IgnoreTagsConfig)
		}

		returnctx
	}

	ctx:=bootstrapContext(context.Background(),conn)
	d:=&resourceData{}

	for_,v:=rangeinterceptors{
		vardiagsdiag.Diagnostics
		_,diags=v.interceptor.run(ctx,d,conn,v.when,v.why,diags)
		ifgot,want:=len(diags),1;got!=want{
t.Errorf("lengthofdiags=%v,want%v",got,want)
		}
	}
}

typeresourceDatastruct{}

func(d*resourceData)GetRawConfig()cty.Value{
	returncty.ObjectVal(map[string]cty.Value{
		"tags":cty.MapVal(map[string]cty.Value{
"tag1":cty.StringVal("value1"),
		}),
	})
}

func(d*resourceData)GetRawPlan()cty.Value{
	returncty.ObjectVal(map[string]cty.Value{
		"tags_all":cty.MapVal(map[string]cty.Value{
"tag1":cty.UnknownVal(cty.String),
		}),
	})
}

func(d*resourceData)GetRawState()cty.Value{//nosemgrep:ci.aws-in-func-name
	returncty.Value{}
}

func(d*resourceData)Get(keystring)any{
	returnnil
}

func(d*resourceData)Id()string{
	return"id"
}

func(d*resourceData)Set(string,any)error{
	returnnil
}

func(d*resourceData)GetChange(keystring)(interface{},interface{}){
	returnnil,nil
}

func(d*resourceData)HasChange(keystring)bool{
	returnfalse
}
