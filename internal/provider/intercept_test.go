//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageprovider

import(
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)

funcTestInterceptorsWhy(t*testing.T){
	t.Parallel()

	varinterceptorsinterceptorItems

	interceptors=append(interceptors,interceptorItem{
		when:Before,
		why:Create,
		interceptor:interceptorFunc(func(ctxcontext.Context,dschemaResourceData,metaany,whenwhen,whywhy,diagsdiag.Diagnostics)(context.Context,diag.Diagnostics){
returnctx,diags
		}),
	})
	interceptors=append(interceptors,interceptorItem{
		when:After,
		why:Delete,
		interceptor:interceptorFunc(func(ctxcontext.Context,dschemaResourceData,metaany,whenwhen,whywhy,diagsdiag.Diagnostics)(context.Context,diag.Diagnostics){
returnctx,diags
		}),
	})
	interceptors=append(interceptors,interceptorItem{
		when:Before,
		why:Create,
		interceptor:interceptorFunc(func(ctxcontext.Context,dschemaResourceData,metaany,whenwhen,whywhy,diagsdiag.Diagnostics)(context.Context,diag.Diagnostics){
returnctx,diags
		}),
	})

	ifgot,want:=len(interceptors.why(Create)),2;got!=want{
		t.Errorf("lengthofinterceptors.Why(Create)=%v,want%v",got,want)
	}
	ifgot,want:=len(interceptors.why(Read)),0;got!=want{
		t.Errorf("lengthofinterceptors.Why(Read)=%v,want%v",got,want)
	}
	ifgot,want:=len(interceptors.why(Update)),0;got!=want{
		t.Errorf("lengthofinterceptors.Why(Update)=%v,want%v",got,want)
	}
	ifgot,want:=len(interceptors.why(Delete)),1;got!=want{
		t.Errorf("lengthofinterceptors.Why(Delete)=%v,want%v",got,want)
	}
}

funcTestInterceptedHandler(t*testing.T){
	t.Parallel()

	varinterceptorsinterceptorItems

	interceptors=append(interceptors,interceptorItem{
		when:Before,
		why:Create,
		interceptor:interceptorFunc(func(ctxcontext.Context,dschemaResourceData,metaany,whenwhen,whywhy,diagsdiag.Diagnostics)(context.Context,diag.Diagnostics){
returnctx,diags
		}),
	})
	interceptors=append(interceptors,interceptorItem{
		when:After,
		why:Delete,
		interceptor:interceptorFunc(func(ctxcontext.Context,dschemaResourceData,metaany,whenwhen,whywhy,diagsdiag.Diagnostics)(context.Context,diag.Diagnostics){
returnctx,diags
		}),
	})
	interceptors=append(interceptors,interceptorItem{
		when:Before,
		why:Create,
		interceptor:interceptorFunc(func(ctxcontext.Context,dschemaResourceData,metaany,whenwhen,whywhy,diagsdiag.Diagnostics)(context.Context,diag.Diagnostics){
returnctx,diags
		}),
	})

	varreadschema.ReadContextFunc=func(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
		vardiagsdiag.Diagnostics
		returnsdkdiag.AppendErrorf(diags,"readerror")
	}
	bootstrapContext:=func(ctxcontext.Context,metaany)context.Context{
		returnctx
	}

	diags:=interceptedHandler(bootstrapContext,interceptors,read,Read)(context.Background(),nil,42)
	ifgot,want:=len(diags),1;got!=want{
		t.Errorf("lengthofdiags=%v,want%v",got,want)
	}
}
