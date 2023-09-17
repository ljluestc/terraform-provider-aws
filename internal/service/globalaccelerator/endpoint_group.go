//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageglobalaccelerator

import(
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

//@SDKResource("aws_globalaccelerator_endpoint_group")
funcResourceEndpointGroup()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceEndpointGroupCreate,
		ReadWithoutTimeout:resourceEndpointGroupRead,
		UpdateWithoutTimeout:resourceEndpointGroupUpdate,
		DeleteWithoutTimeout:resourceEndpointGroupDelete,

		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},

		Timeouts:&schema.ResourceTimeout{
			Create:schema.DefaultTimeout(30*time.Minute),
			Update:schema.DefaultTimeout(30*time.Minute),
			Delete:schema.DefaultTimeout(30*time.Minute),
		},

		Schema:map[string]*schema.Schema{
			"arn":{
				Type:schema.TypeString,
				Computed:true,
			},
			"endpoint_configuration":{
				Type:schema.TypeSet,
				Optional:true,
				Elem:&schema.Resource{
					Schema:map[string]*schema.Schema{
						"client_ip_preservation_enabled":{
							Type:schema.TypeBool,
							Optional:true,
							Computed:true,
						},
						"endpoint_id":{
							Type:schema.TypeString,
							Optional:true,
							ValidateFunc:validation.StringLenBetween(1,255),
						},
						"weight":{
							Type:schema.TypeInt,
							Optional:true,
							ValidateFunc:validation.IntBetween(0,255),
						},
					},
				},
			},
			"endpoint_group_region":{
				Type:schema.TypeString,
				Optional:true,
				Computed:true,
				ForceNew:true,
				ValidateFunc:verify.ValidRegionName,
			},
			"health_check_interval_seconds":{
				Type:schema.TypeInt,
				Optional:true,
				Default:30,
				ValidateFunc:validation.IntBetween(10,30),
			},
			"health_check_path":{
				Type:schema.TypeString,
				Optional:true,
				Computed:true,
				ValidateFunc:validation.StringLenBetween(1,255),
			},
			"health_check_port":{
				Type:schema.TypeInt,
				Optional:true,
				Computed:true,
				ValidateFunc:validation.IsPortNumber,
			},
			"health_check_protocol":{
				Type:schema.TypeString,
				Optional:true,
				Default:globalaccelerator.HealthCheckProtocolTcp,
				ValidateFunc:validation.StringInSlice(globalaccelerator.HealthCheckProtocol_Values(),false),
			},
			"listener_arn":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
				ValidateFunc:verify.ValidARN,
			},
			"port_override":{
				Type:schema.TypeSet,
				Optional:true,
				MaxItems:10,
				Elem:&schema.Resource{
					Schema:map[string]*schema.Schema{
						"endpoint_port":{
							Type:schema.TypeInt,
							Required:true,
							ValidateFunc:validation.IsPortNumber,
						},
						"listener_port":{
							Type:schema.TypeInt,
							Required:true,
							ValidateFunc:validation.IsPortNumber,
						},
					},
				},
			},
			"threshold_count":{
				Type:schema.TypeInt,
				Optional:true,
				Default:3,
				ValidateFunc:validation.IntBetween(1,10),
			},
			"traffic_dial_percentage":{
				Type:schema.TypeFloat,
				Optional:true,
				Default:100.0,
				ValidateFunc:validation.FloatBetween(0.0,100.0),
			},
		},
	}
}
funcresourceEndpointGroupCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).GlobalAcceleratorConn(ctx)

	input:=&globalaccelerator.CreateEndpointGroupInput{
		EndpointGroupRegion:aws.String(meta.(*conns.AWSClient).Region),
		IdempotencyToken:aws.String(id.UniqueId()),
		ListenerArn:aws.String(d.Get("listener_arn").(string)),
	}

	ifv,ok:=d.GetOk("endpoint_configuration");ok&&v.(*schema.Set).Len()>0{
		input.EndpointConfigurations=expandEndpointConfigurations(v.(*schema.Set).List())
	}

	ifv,ok:=d.GetOk("endpoint_group_region");ok{
		input.EndpointGroupRegion=aws.String(v.(string))
	}

	ifv,ok:=d.GetOk("health_check_interval_seconds");ok{
		input.HealthCheckIntervalSeconds=aws.Int64(int64(v.(int)))
	}

	ifv,ok:=d.GetOk("health_check_path");ok{
		input.HealthCheckPath=aws.String(v.(string))
	}

	ifv,ok:=d.GetOk("health_check_port");ok{
		input.HealthCheckPort=aws.Int64(int64(v.(int)))
	}

	ifv,ok:=d.GetOk("health_check_protocol");ok{
		input.HealthCheckProtocol=aws.String(v.(string))
	}

	ifv,ok:=d.GetOk("port_override");ok&&v.(*schema.Set).Len()>0{
		input.PortOverrides=expandPortOverrides(v.(*schema.Set).List())
	}

	ifv,ok:=d.GetOk("threshold_count");ok{
		input.ThresholdCount=aws.Int64(int64(v.(int)))
	}

	ifv,ok:=d.Get("traffic_dial_percentage").(float64);ok{
		input.TrafficDialPercentage=aws.Float64(v)
	}

	resp,err:=conn.CreateEndpointGroupWithContext(ctx,input)

	iferr!=nil{
		returndiag.Errorf("creatingGlobalAcceleratorEndpointGroup:%s",err)
	}

	d.SetId(aws.StringValue(resp.EndpointGroup.EndpointGroupArn))

	acceleratorARN,err:=ListenerOrEndpointGroupARNToAcceleratorARN(d.Id())

	iferr!=nil{
		returndiag.FromErr(err)
	}

	if_,err:=waitAcceleratorDeployed(ctx,conn,acceleratorARN,d.Timeout(schema.TimeoutCreate));err!=nil{
		returndiag.Errorf("waitingforGlobalAcceleratorAccelerator(%s)deployment:%s",acceleratorARN,err)
	}

	returnresourceEndpointGroupRead(ctx,d,meta)
}
funcresourceEndpointGroupRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).GlobalAcceleratorConn(ctx)

	endpointGroup,err:=FindEndpointGroupByARN(ctx,conn,d.Id())

	if!d.IsNewResource()&&tfresource.NotFound(err){
		log.Printf("[WARN]GlobalAcceleratorendpointgroup(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returnnil
	}

	iferr!=nil{
		returndiag.Errorf("readingGlobalAcceleratorEndpointGroup(%s):%s",d.Id(),err)
	}

	listenerARN,err:=EndpointGroupARNToListenerARN(d.Id())

	iferr!=nil{
		returndiag.FromErr(err)
	}

	d.Set("arn",endpointGroup.EndpointGroupArn)
	iferr:=d.Set("endpoint_configuration",flattenEndpointDescriptions(endpointGroup.EndpointDescriptions));err!=nil{
		returndiag.Errorf("settingendpoint_configuration:%s",err)
	}
	d.Set("endpoint_group_region",endpointGroup.EndpointGroupRegion)
	d.Set("health_check_interval_seconds",endpointGroup.HealthCheckIntervalSeconds)
	d.Set("health_check_path",endpointGroup.HealthCheckPath)
	d.Set("health_check_port",endpointGroup.HealthCheckPort)
	d.Set("health_check_protocol",endpointGroup.HealthCheckProtocol)
	d.Set("listener_arn",listenerARN)
	iferr:=d.Set("port_override",flattenPortOverrides(endpointGroup.PortOverrides));err!=nil{
		returndiag.Errorf("settingport_override:%s",err)
	}
	d.Set("threshold_count",endpointGroup.ThresholdCount)
	d.Set("traffic_dial_percentage",endpointGroup.TrafficDialPercentage)

	returnnil
}
funcresourceEndpointGroupUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).GlobalAcceleratorConn(ctx)

	input:=&globalaccelerator.UpdateEndpointGroupInput{
		EndpointGroupArn:aws.String(d.Id()),
	}

	ifv,ok:=d.GetOk("endpoint_configuration");ok&&v.(*schema.Set).Len()>0{
		input.EndpointConfigurations=expandEndpointConfigurations(v.(*schema.Set).List())
	}else{
		input.EndpointConfigurations=[]*globalaccelerator.EndpointConfiguration{}
	}

	ifv,ok:=d.GetOk("health_check_interval_seconds");ok{
		input.HealthCheckIntervalSeconds=aws.Int64(int64(v.(int)))
	}

	ifv,ok:=d.GetOk("health_check_path");ok{
		input.HealthCheckPath=aws.String(v.(string))
	}

	ifv,ok:=d.GetOk("health_check_port");ok{
		input.HealthCheckPort=aws.Int64(int64(v.(int)))
	}

	ifv,ok:=d.GetOk("health_check_protocol");ok{
		input.HealthCheckProtocol=aws.String(v.(string))
	}

	ifv,ok:=d.GetOk("port_override");ok&&v.(*schema.Set).Len()>0{
		input.PortOverrides=expandPortOverrides(v.(*schema.Set).List())
	}else{
		input.PortOverrides=[]*globalaccelerator.PortOverride{}
	}

	ifv,ok:=d.GetOk("threshold_count");ok{
		input.ThresholdCount=aws.Int64(int64(v.(int)))
	}

	ifv,ok:=d.Get("traffic_dial_percentage").(float64);ok{
		input.TrafficDialPercentage=aws.Float64(v)
	}

	_,err:=conn.UpdateEndpointGroupWithContext(ctx,input)

	iferr!=nil{
		returndiag.Errorf("updatingGlobalAcceleratorEndpointGroup(%s):%s",d.Id(),err)
	}

	acceleratorARN,err:=ListenerOrEndpointGroupARNToAcceleratorARN(d.Id())

	iferr!=nil{
		returndiag.FromErr(err)
	}

	if_,err:=waitAcceleratorDeployed(ctx,conn,acceleratorARN,d.Timeout(schema.TimeoutUpdate));err!=nil{
		returndiag.Errorf("waitingforGlobalAcceleratorAccelerator(%s)deployment:%s",acceleratorARN,err)
	}

	returnresourceEndpointGroupRead(ctx,d,meta)
}
funcresourceEndpointGroupDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).GlobalAcceleratorConn(ctx)

	log.Printf("[DEBUG]DeletingGlobalAcceleratorEndpointGroup:%s",d.Id())
	_,err:=conn.DeleteEndpointGroupWithContext(ctx,&globalaccelerator.DeleteEndpointGroupInput{
		EndpointGroupArn:aws.String(d.Id()),
	})

	iftfawserr.ErrCodeEquals(err,globalaccelerator.ErrCodeEndpointGroupNotFoundException){
		returnnil
	}

	iferr!=nil{
		returndiag.Errorf("deletingGlobalAcceleratorEndpointGroup(%s):%s",d.Id(),err)
	}

	acceleratorARN,err:=ListenerOrEndpointGroupARNToAcceleratorARN(d.Id())

	iferr!=nil{
		returndiag.FromErr(err)
	}

	if_,err:=waitAcceleratorDeployed(ctx,conn,acceleratorARN,d.Timeout(schema.TimeoutDelete));err!=nil{
		returndiag.Errorf("waitingforGlobalAcceleratorAccelerator(%s)deployment:%s",acceleratorARN,err)
	}

	returnnil
}
funcFindEndpointGroupByARN(ctxcontext.Context,conn*globalaccelerator.GlobalAccelerator,arnstring)(*globalaccelerator.EndpointGroup,error){
	input:=&globalaccelerator.DescribeEndpointGroupInput{
		EndpointGroupArn:aws.String(arn),
	}

	returnfindEndpointGroup(ctx,conn,input)
}
funcfindEndpointGroup(ctxcontext.Context,conn*globalaccelerator.GlobalAccelerator,input*globalaccelerator.DescribeEndpointGroupInput)(*globalaccelerator.EndpointGroup,error){
	output,err:=conn.DescribeEndpointGroupWithContext(ctx,input)

	iftfawserr.ErrCodeEquals(err,globalaccelerator.ErrCodeEndpointGroupNotFoundException){
		returnnil,&retry.NotFoundError{
			LastError:err,
			LastRequest:input,
		}
	}

	iferr!=nil{
		returnnil,err
	}

	ifoutput==nil||output.EndpointGroup==nil{
		returnnil,tfresource.NewEmptyResultError(input)
	}

	returnoutput.EndpointGroup,nil
}
funcexpandEndpointConfiguration(tfMapmap[string]interface{})*globalaccelerator.EndpointConfiguration{
	iftfMap==nil{
		returnnil
	}

	apiObject:=&globalaccelerator.EndpointConfiguration{}

	ifv,ok:=tfMap["client_ip_preservation_enabled"].(bool);ok{
		apiObject.ClientIPPreservationEnabled=aws.Bool(v)
	}

	ifv,ok:=tfMap["endpoint_id"].(string);ok&&v!=""{
		apiObject.EndpointId=aws.String(v)
	}

	ifv,ok:=tfMap["weight"].(int);ok{
		apiObject.Weight=aws.Int64(int64(v))
	}

	returnapiObject
}
funcexpandEndpointConfigurations(tfList[]interface{})[]*globalaccelerator.EndpointConfiguration{
	iflen(tfList)==0{
		returnnil
	}

	varapiObjects[]*globalaccelerator.EndpointConfiguration

	for_,tfMapRaw:=rangetfList{
		tfMap,ok:=tfMapRaw.(map[string]interface{})

		if!ok{
			continue
		}

		apiObject:=expandEndpointConfiguration(tfMap)

		ifapiObject==nil{
			continue
		}

		apiObjects=append(apiObjects,apiObject)
	}

	returnapiObjects
}
funcexpandPortOverride(tfMapmap[string]interface{})*globalaccelerator.PortOverride{
	iftfMap==nil{
		returnnil
	}

	apiObject:=&globalaccelerator.PortOverride{}

	ifv,ok:=tfMap["endpoint_port"].(int);ok&&v!=0{
		apiObject.EndpointPort=aws.Int64(int64(v))
	}

	ifv,ok:=tfMap["listener_port"].(int);ok&&v!=0{
		apiObject.ListenerPort=aws.Int64(int64(v))
	}

	returnapiObject
}
funcexpandPortOverrides(tfList[]interface{})[]*globalaccelerator.PortOverride{
	iflen(tfList)==0{
		returnnil
	}

	varapiObjects[]*globalaccelerator.PortOverride

	for_,tfMapRaw:=rangetfList{
		tfMap,ok:=tfMapRaw.(map[string]interface{})

		if!ok{
			continue
		}

		apiObject:=expandPortOverride(tfMap)

		ifapiObject==nil{
			continue
		}

		apiObjects=append(apiObjects,apiObject)
	}

	returnapiObjects
}
funcflattenEndpointDescription(apiObject*globalaccelerator.EndpointDescription)map[string]interface{}{
	ifapiObject==nil{
		returnnil
	}

	tfMap:=map[string]interface{}{}

	ifv:=apiObject.ClientIPPreservationEnabled;v!=nil{
		tfMap["client_ip_preservation_enabled"]=aws.BoolValue(v)
	}

	ifv:=apiObject.EndpointId;v!=nil{
		tfMap["endpoint_id"]=aws.StringValue(v)
	}

	ifv:=apiObject.Weight;v!=nil{
		tfMap["weight"]=aws.Int64Value(v)
	}

	returntfMap
}
funcflattenEndpointDescriptions(apiObjects[]*globalaccelerator.EndpointDescription)[]interface{}{
	iflen(apiObjects)==0{
		returnnil
	}

	vartfList[]interface{}

	for_,apiObject:=rangeapiObjects{
		ifapiObject==nil{
			continue
		}

		tfList=append(tfList,flattenEndpointDescription(apiObject))
	}

	returntfList
}
funcflattenPortOverride(apiObject*globalaccelerator.PortOverride)map[string]interface{}{
	ifapiObject==nil{
		returnnil
	}

	tfMap:=map[string]interface{}{}

	ifv:=apiObject.EndpointPort;v!=nil{
		tfMap["endpoint_port"]=aws.Int64Value(v)
	}

	ifv:=apiObject.ListenerPort;v!=nil{
		tfMap["listener_port"]=aws.Int64Value(v)
	}

	returntfMap
}
funcflattenPortOverrides(apiObjects[]*globalaccelerator.PortOverride)[]interface{}{
	iflen(apiObjects)==0{
		returnnil
	}

	vartfList[]interface{}

	for_,apiObject:=rangeapiObjects{
		ifapiObject==nil{
			continue
		}

		tfList=append(tfList,flattenPortOverride(apiObject))
	}

	returntfList
}
