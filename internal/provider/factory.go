//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageprovider

import(
"context"

"github.com/hashicorp/terraform-plugin-framework/providerserver"
"github.com/hashicorp/terraform-plugin-go/tfprotov5"
"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/provider/fwprovider"
)

//ProtoV5ProviderServerFactoryreturnsamuxedterraform-plugin-goprotocolv5providerfactoryfunction.
//Thisfactoryfunctionissuitableforusewiththeterraform-plugin-goServefunction.
//Theprimary(PluginSDK)providerserverisalsoreturned(usefulfortesting).
funcProtoV5ProviderServerFactory(ctxcontext.Context)(func()tfprotov5.ProviderServer,*schema.Provider,error){
primary,err:=New(ctx)

iferr!=nil{
returnnil,nil,err
}

servers:=[]func()tfprotov5.ProviderServer{
primary.GRPCProvider,
providerserver.NewProtocol5(fwprovider.New(primary)),
}

muxServer,err:=tf5muxserver.NewMuxServer(ctx,servers...)

iferr!=nil{
returnnil,nil,err
}

returnmuxServer.ProviderServer,primary,nil
}
