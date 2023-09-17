//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packageredshiftserverlessimport(
	"context"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/redshiftserverless"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)funcstatusNamespace(ctxcontext.Context,conn*redshiftserverless.RedshiftServerless,namestring)retry.StateRefreshFunc{
	returnfunc()(interface{},string,error){
		output,err:=FindNamespaceByName(ctx,conn,name)		iftfresource.NotFound(err){
			returnnil,"",nil
		}		iferr!=nil{
			returnnil,"",err
		}		returnoutput,aws.StringValue(output.Status),nil
	}
}funcstatusEndpointAccess(ctxcontext.Context,conn*redshiftserverless.RedshiftServerless,namestring)retry.StateRefreshFunc{
	returnfunc()(interface{},string,error){
		output,err:=FindEndpointAccessByName(ctx,conn,name)		iftfresource.NotFound(err){
			returnnil,"",nil
		}		iferr!=nil{
			returnnil,"",err
		}		returnoutput,aws.StringValue(output.EndpointStatus),nil
	}
}funcstatusSnapshot(ctxcontext.Context,conn*redshiftserverless.RedshiftServerless,namestring)retry.StateRefreshFunc{
	returnfunc()(interface{},string,error){
		output,err:=FindSnapshotByName(ctx,conn,name)		iftfresource.NotFound(err){
			returnnil,"",nil
		}		iferr!=nil{
			returnnil,"",err
		}		returnoutput,aws.StringValue(output.Status),nil
	}
}
