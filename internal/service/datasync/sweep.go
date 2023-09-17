//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

//go:buildsweep
//+buildsweep

packagedatasync

import(
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/sweep"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

funcinit(){
	resource.AddTestSweepers("aws_datasync_agent",&resource.Sweeper{
		Name:"aws_datasync_agent",
		F:sweepAgents,
		Dependencies:[]string{
			"aws_datasync_location",
		},
	})

	//Pseudo-resourceforanyDataSynclocationresourcetype.
	resource.AddTestSweepers("aws_datasync_location",&resource.Sweeper{
		Name:"aws_datasync_location",
		F:sweepLocations,
		Dependencies:[]string{
			"aws_datasync_task",
		},
	})

	resource.AddTestSweepers("aws_datasync_task",&resource.Sweeper{
		Name:"aws_datasync_task",
		F:sweepTasks,
	})
}

funcsweepAgents(regionstring)error{
	ctx:=sweep.Context(region)
	client,err:=sweep.SharedRegionalSweepClient(ctx,region)
	iferr!=nil{
		returnfmt.Errorf("errorgettingclient:%s",err)
	}
	conn:=client.DataSyncConn(ctx)
	input:=&datasync.ListAgentsInput{}
	sweepResources:=make([]sweep.Sweepable,0)

	err=conn.ListAgentsPagesWithContext(ctx,input,func(page*datasync.ListAgentsOutput,lastPagebool)bool{
		ifpage==nil{
			return!lastPage
		}

		for_,v:=rangepage.Agents{
			r:=ResourceAgent()
			d:=r.Data(nil)
			d.SetId(aws.StringValue(v.AgentArn))

			sweepResources=append(sweepResources,sweep.NewSweepResource(r,d,client))
		}

		return!lastPage
	})

	ifsweep.SkipSweepError(err){
		log.Printf("[WARN]SkippingDataSyncAgentsweepfor%s:%s",region,err)
		returnnil
	}

	iferr!=nil{
		returnfmt.Errorf("errorlistingDataSyncAgents(%s):%w",region,err)
	}

	err=sweep.SweepOrchestrator(ctx,sweepResources)

	iferr!=nil{
		returnfmt.Errorf("errorsweepingDataSyncAgents(%s):%w",region,err)
	}

	returnnil
}

funcsweepLocations(regionstring)error{
	ctx:=sweep.Context(region)
	client,err:=sweep.SharedRegionalSweepClient(ctx,region)
	iferr!=nil{
		returnfmt.Errorf("errorgettingclient:%w",err)
	}
	conn:=client.DataSyncConn(ctx)
	input:=&datasync.ListLocationsInput{}
	sweepResources:=make([]sweep.Sweepable,0)

	err=conn.ListLocationsPagesWithContext(ctx,input,func(page*datasync.ListLocationsOutput,lastPagebool)bool{
		ifpage==nil{
			return!lastPage
		}

		for_,v:=rangepage.Locations{
			sweepable:=&sweepableLocation{
				arn:aws.StringValue(v.LocationArn),
				conn:conn,
			}

			sweepResources=append(sweepResources,sweepable)
		}

		return!lastPage
	})

	ifsweep.SkipSweepError(err){
		log.Printf("[WARN]SkippingDataSyncLocationsweepfor%s:%s",region,err)
		returnnil
	}

	iferr!=nil{
		returnfmt.Errorf("errorlistingDataSyncLocations(%s):%w",region,err)
	}

	err=sweep.SweepOrchestrator(ctx,sweepResources)

	iferr!=nil{
		returnfmt.Errorf("errorsweepingDataSyncLocations(%s):%w",region,err)
	}

	returnnil
}

typesweepableLocationstruct{
	arnstring
	conn*datasync.DataSync
}

func(sweepable*sweepableLocation)Delete(ctxcontext.Context,timeouttime.Duration,optFns...tfresource.OptionsFunc)error{
	log.Printf("[DEBUG]DeletingDataSyncLocation:%s",sweepable.arn)
	_,err:=sweepable.conn.DeleteLocationWithContext(ctx,&datasync.DeleteLocationInput{
		LocationArn:aws.String(sweepable.arn),
	})

	iftfawserr.ErrMessageContains(err,datasync.ErrCodeInvalidRequestException,"notfound"){
		returnnil
	}

	iferr!=nil{
		returnfmt.Errorf("deletingDataSyncLocation(%s):%w",sweepable.arn,err)
	}

	returnnil
}

funcsweepTasks(regionstring)error{
	ctx:=sweep.Context(region)
	client,err:=sweep.SharedRegionalSweepClient(ctx,region)
	iferr!=nil{
		returnfmt.Errorf("errorgettingclient:%w",err)
	}
	conn:=client.DataSyncConn(ctx)
	input:=&datasync.ListTasksInput{}
	sweepResources:=make([]sweep.Sweepable,0)

	err=conn.ListTasksPagesWithContext(ctx,input,func(page*datasync.ListTasksOutput,lastPagebool)bool{
		ifpage==nil{
			return!lastPage
		}

		for_,v:=rangepage.Tasks{
			r:=ResourceTask()
			d:=r.Data(nil)
			d.SetId(aws.StringValue(v.TaskArn))

			sweepResources=append(sweepResources,sweep.NewSweepResource(r,d,client))
		}

		return!lastPage
	})

	ifsweep.SkipSweepError(err){
		log.Printf("[WARN]SkippingDataSyncTasksweepfor%s:%s",region,err)
		returnnil
	}

	iferr!=nil{
		returnfmt.Errorf("errorlistingDataSyncTasks(%s):%w",region,err)
	}

	err=sweep.SweepOrchestrator(ctx,sweepResources)

	iferr!=nil{
		returnfmt.Errorf("errorsweepingDataSyncTasks(%s):%w",region,err)
	}

	returnnil
}
