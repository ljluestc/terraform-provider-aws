//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagequicksight

import(
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

//statusfetchestheDataSourceanditsStatus

funcstatus(ctxcontext.Context,conn*quicksight.QuickSight,accountId,datasourceIdstring)retry.StateRefresh
func{
	return
func()(interface{},string,error){
		input:=&quicksight.DescribeDataSourceInput{
			AwsAccountId:aws.String(accountId),
			DataSourceId:aws.String(datasourceId),
		}

		output,err:=conn.DescribeDataSourceWithContext(ctx,input)

		iferr!=nil{
			returnnil,"",err
		}

		ifoutput==nil||output.DataSource==nil{
			returnnil,"",nil
		}

		returnoutput.DataSource,aws.StringValue(output.DataSource.Status),nil
	}
}

//FetchTemplatestatus

funcstatusTemplate(ctxcontext.Context,conn*quicksight.QuickSight,idstring)retry.StateRefresh
func{
	return
func()(interface{},string,error){
		out,err:=FindTemplateByID(ctx,conn,id)
		iftfresource.NotFound(err){
			returnnil,"",nil
		}

		iferr!=nil{
			returnnil,"",err
		}

		returnout,*out.Version.Status,nil
	}
}

//FetchDashboardstatus

funcstatusDashboard(ctxcontext.Context,conn*quicksight.QuickSight,idstring)retry.StateRefresh
func{
	return
func()(interface{},string,error){
		out,err:=FindDashboardByID(ctx,conn,id)
		iftfresource.NotFound(err){
			returnnil,"",nil
		}

		iferr!=nil{
			returnnil,"",err
		}

		returnout,*out.Version.Status,nil
	}
}

//FetchAnalysisstatus

funcstatusAnalysis(ctxcontext.Context,conn*quicksight.QuickSight,idstring)retry.StateRefresh
func{
	return
func()(interface{},string,error){
		out,err:=FindAnalysisByID(ctx,conn,id)
		iftfresource.NotFound(err){
			returnnil,"",nil
		}

		iferr!=nil{
			returnnil,"",err
		}

		returnout,*out.Status,nil
	}
}

//FetchThemestatus

funcstatusTheme(ctxcontext.Context,conn*quicksight.QuickSight,idstring)retry.StateRefresh
func{
	return
func()(interface{},string,error){
		out,err:=FindThemeByID(ctx,conn,id)
		iftfresource.NotFound(err){
			returnnil,"",nil
		}

		iferr!=nil{
			returnnil,"",err
		}

		returnout,*out.Version.Status,nil
	}
}
