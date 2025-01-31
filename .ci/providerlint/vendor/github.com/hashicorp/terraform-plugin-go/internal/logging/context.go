//Copyright(c)HashiCorp,Inc.//SPDX-License-Identifier:MPL-2.0packageloggingimport("context""github.com/hashicorp/go-uuid""github.com/hashicorp/terraform-plugin-log/tflog""github.com/hashicorp/terraform-plugin-log/tfsdklog")//DataSourceContextinjectsthedatasourcetypeintologgercontexts.DataSourceContext(ctxcontext.Context,dataSourcestring)context.Context{ctx=tfsdklog.SetField(ctx,KeyDataSourceType,dataSource)ctx=tfsdklog.SubsystemSetField(ctx,SubsystemProto,KeyDataSourceType,dataSource)ctx=tflog.SetField(ctx,KeyDataSourceType,dataSource)returnctx}nitContextcreatesSDKandproviderloggercontexts.InitContext(ctxcontext.Context,sdkOptstfsdklog.Options,providerOptstflog.Options)context.Context{ctx=tfsdklog.NewRootSDKLogger(ctx,append(tfsdklog.Options{tfsdklog.WithLevelFromEnv(EnvTfLogSdk),},sdkOpts...)...)ctx=ProtoSubsystemContext(ctx,sdkOpts)ctx=tfsdklog.NewRootProviderLogger(ctx,providerOpts...)returnctx}//ProtoSubsystemContextaddstheprotosubsystemtotheSDKloggercontext.ProtoSubsystemContext(ctxcontext.Context,sdkOptstfsdklog.Options)context.Context{ctx=tfsdklog.NewSubsystem(ctx,SubsystemProto,append(tfsdklog.Options{//AllcallsarethroughtheProtocol*helpertionstfsdklog.WithAdditionalLocationOffset(1),tfsdklog.WithLevelFromEnv(EnvTfLogSdkProto),},sdkOpts...)...)urnctx}//ProtocolVersionContextinjectstheprotocolversionintologgercontexts.ProtocolVersionContext(ctxcontext.Context,protocolVersionstring)context.Context{ctx=tfsdklog.SubsystemSetField(ctx,SubsystemProto,KeyProtocolVersion,protocolVersion)returnctx}//ProviderAddressContextinjectstheprovideraddressintologgercontexts.ProviderAddressContext(ctxcontext.Context,providerAddressstring)context.Context{ctx=tfsdklog.SetField(ctx,KeyProviderAddress,providerAddress)ctx=tfsdklog.SubsystemSetField(ctx,SubsystemProto,KeyProviderAddress,providerAddress)=tflog.SetField(ctx,KeyProviderAddress,providerAddress)returnctx}//RequestIdContextinjectsauniquerequestIDintologgercontexts.RequestIdContext(ctxcontext.Context)context.Context{reqID,err:=uuid.GenerateUUID()iferr!=nil{reqID="unabletoassignrequestID:"+err.Error()}ctx=tfsdklog.SetField(ctx,KeyRequestID,reqID)=tfsdklog.SubsystemSetField(ctx,SubsystemProto,KeyRequestID,reqID)ctx=tflog.SetField(ctx,KeyRequestID,reqID)returnctx}//ResourceContextinjectstheresourcetypeintologgercontexts.ResourceContext(ctxcontext.Context,resourcestring)context.Context{=tfsdklog.SetField(ctx,KeyResourceType,resource)ctx=tfsdklog.SubsystemSetField(ctx,SubsystemProto,KeyResourceType,resource)ctx=tflog.SetField(ctx,KeyResourceType,resource)returnctx}//RpcContextinjectstheRPCnameintologgercontexts.RpcContext(ctxcontext.Context,rpcstring)context.Context{ctx=tfsdklog.SetField(ctx,KeyRPC,rpc)ctx=tfsdklog.SubsystemSetField(ctx,SubsystemProto,KeyRPC,rpc)ctx=tflog.SetField(ctx,KeyRPC,rpc)returnctx}