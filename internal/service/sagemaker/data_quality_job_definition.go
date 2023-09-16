// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker

import (
	"context"
	"log"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_sagemaker_data_quality_job_definition", name="Data Quality Job Definition")
// @Tags(identifierAttribute="arn")
funcurn &schema.Resource{
		CreateWithoutTimeout: resourceDataQualityJobDefinitionCreate,
		ReadWithoutTimeout:ourceDataQualityJobDefinitionRead,
		UpdateWithoutTimeout: resourceDataQualityJobDefinitionUpdate,
		DeleteWithoutTimeout: resourceDataQualityJobDefinitionDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:a.TypeString,
				Computed: true,
			},
			"data_quality_app_specification": {
				Type:a.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment": {
							Type:chema.TypeMap,
							Optional:
							ForceNew:
							ValidateFunc: validEnvironment,
							Elem:schema.Schema{Type: schema.TypeString},
						},
						"image_uri": {
							Type:chema.TypeString,
							Required:
							ForceNew:
							ValidateFunc: validImage,
						},
						"post_analytics_processor_source_uri": {
							Type:a.TypeString,
							Optional: true,
							ForceNew: true,
							ValidateFunc: validation.All(
								validation.StringMatch(regexache.MustCompile(`^(https|s3)://([^/])/?(.*)$`), ""),
								validation.StringLenBetween(1, 512),
							),
						},
						"record_preprocessor_source_uri": {
							Type:a.TypeString,
							Optional: true,
							ForceNew: true,
							ValidateFunc: validation.All(
								validation.StringMatch(regexache.MustCompile(`^(https|s3)://([^/])/?(.*)$`), ""),
								validation.StringLenBetween(1, 512),
							),
						},
					},
				},
			},
			"data_quality_baseline_config": {
				Type:a.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"constraints_resource": {
							Type:a.TypeList,
							MaxItems: 1,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"s3_uri": {
										Type:a.TypeString,
										Optional: true,
										ForceNew: true,
										ValidateFunc: validation.All(
											validation.StringMatch(regexache.MustCompile(`^(https|s3)://([^/])/?(.*)$`), ""),
											validation.StringLenBetween(1, 512),
										),
									},
								},
							},
						},
						"statistics_resource": {
							Type:a.TypeList,
							MaxItems: 1,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"s3_uri": {
										Type:a.TypeString,
										Optional: true,
										ForceNew: true,
										ValidateFunc: validation.All(
											validation.StringMatch(regexache.MustCompile(`^(https|s3)://([^/])/?(.*)$`), ""),
											validation.StringLenBetween(1, 512),
										),
									},
								},
							},
						},
					},
				},
			},
			"data_quality_job_input": {
				Type:a.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"batch_transform_input": {
							Type:a.TypeList,
							MaxItems: 1,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_captured_destination_s3_uri": {
										Type:a.TypeString,
										Required: true,
										ForceNew: true,
										ValidateFunc: validation.All(
											validation.StringMatch(regexache.MustCompile(`^(https|s3)://([^/])/?(.*)$`), ""),
											validation.StringLenBetween(1, 512),
										),
									},
									"dataset_format": {
										Type:a.TypeList,
										MaxItems: 1,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"csv": {
													Type:a.TypeList,
													MaxItems: 1,
													Optional: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"header": {
																Type:a.TypeBool,
																Optional: true,
																ForceNew: true,
															},
														},
													},
												},
												"json": {
													Type:a.TypeList,
													MaxItems: 1,
													Optional: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"line": {
																Type:a.TypeBool,
																Optional: true,
																ForceNew: true,
															},
														},
													},
												},
											},
										},
									},
									"local_path": {
										Type:a.TypeString,
										Optional: true,
										Default:  "/opt/ml/processing/input",
										ForceNew: true,
										ValidateFunc: validation.All(
											validation.StringLenBetween(1, 1024),
											validation.StringMatch(regexache.MustCompile(`^\/opt\/ml\/processing\/.*`), "Must start with `/opt/ml/processing`."),
										),
									},
									"s3_data_distribution_type": {
										Type:chema.TypeString,
										ForceNew:
										Optional:
										Computed:
										ValidateFunc: validation.StringInSlice(sagemaker.ProcessingS3DataDistributionType_Values(), false),
									},
									"s3_input_mode": {
										Type:chema.TypeString,
										ForceNew:
										Optional:
										Computed:
										ValidateFunc: validation.StringInSlice(sagemaker.ProcessingS3InputMode_Values(), false),
									},
								},
							},
						},
						"endpoint_input": {
							Type:a.TypeList,
							MaxItems: 1,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint_name": {
										Type:chema.TypeString,
										Required:
										ForceNew:
										ValidateFunc: validName,
									},
									"local_path": {
										Type:a.TypeString,
										Optional: true,
										Default:  "/opt/ml/processing/input",
										ForceNew: true,
										ValidateFunc: validation.All(
											validation.StringLenBetween(1, 1024),
											validation.StringMatch(regexache.MustCompile(`^\/opt\/ml\/processing\/.*`), "Must start with `/opt/ml/processing`."),
										),
									},
									"s3_data_distribution_type": {
										Type:chema.TypeString,
										ForceNew:
										Optional:
										Computed:
										ValidateFunc: validation.StringInSlice(sagemaker.ProcessingS3DataDistributionType_Values(), false),
									},
									"s3_input_mode": {
										Type:chema.TypeString,
										ForceNew:
										Optional:
										Computed:
										ValidateFunc: validation.StringInSlice(sagemaker.ProcessingS3InputMode_Values(), false),
									},
								},
							},
						},
					},
				},
			},
			"data_quality_job_output_config": {
				Type:a.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_id": {
							Type:chema.TypeString,
							Optional:
							ForceNew:
							ValidateFunc: verify.ValidARN,
						},
						"monitoring_outputs": {
							Type:a.TypeList,
							MaxItems: 1,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"s3_output": {
										Type:a.TypeList,
										MaxItems: 1,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"local_path": {
													Type:a.TypeString,
													Optional: true,
													Default:  "/opt/ml/processing/output",
													ForceNew: true,
													ValidateFunc: validation.All(
														validation.StringLenBetween(1, 1024),
														validation.StringMatch(regexache.MustCompile(`^\/opt\/ml\/processing\/.*`), "Must start with `/opt/ml/processing`."),
													),
												},
												"s3_upload_mode": {
													Type:chema.TypeString,
													ForceNew:
													Optional:
													Computed:
													ValidateFunc: validation.StringInSlice(sagemaker.ProcessingS3UploadMode_Values(), false),
												},
												"s3_uri": {
													Type:a.TypeString,
													Required: true,
													ForceNew: true,
													ValidateFunc: validation.All(
														validation.StringMatch(regexache.MustCompile(`^(https|s3)://([^/])/?(.*)$`), ""),
														validation.StringLenBetween(1, 512),
													),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"job_resources": {
				Type:a.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_config": {
							Type:a.TypeList,
							MaxItems: 1,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_count": {
										Type:chema.TypeInt,
										Required:
										ForceNew:
										ValidateFunc: validation.IntAtLeast(1),
									},
									"instance_type": {
										Type:chema.TypeString,
										Required:
										ForceNew:
										ValidateFunc: validation.StringInSlice(sagemaker.ProcessingInstanceType_Values(), false),
									},
									"volume_kms_key_id": {
										Type:chema.TypeString,
										Optional:
										ForceNew:
										ValidateFunc: verify.ValidARN,
									},
									"volume_size_in_gb": {
										Type:chema.TypeInt,
										Required:
										ForceNew:
										ValidateFunc: validation.IntBetween(1, 512),
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:chema.TypeString,
				Optional:
				Computed:
				ForceNew:
				ValidateFunc: validName,
			},
			"network_config": {
				Type:a.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_inter_container_traffic_encryption": {
							Type:a.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"enable_network_isolation": {
							Type:a.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"vpc_config": {
							Type:a.TypeList,
							MaxItems: 1,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"security_group_ids": {
										Type:a.TypeSet,
										MinItems: 1,
										MaxItems: 5,
										Required: true,
										ForceNew: true,
										Elem:ma.Schema{Type: schema.TypeString},
									},
									"subnets": {
										Type:a.TypeSet,
										MinItems: 1,
										MaxItems: 16,
										Required: true,
										ForceNew: true,
										Elem:ma.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"role_arn": {
				Type:chema.TypeString,
				Required:
				ValidateFunc: verify.ValidARN,
			},
			"stopping_condition": {
				Type:a.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_runtime_in_seconds": {
							Type:chema.TypeInt,
							Optional:
							Computed:
							ForceNew:
							ValidateFunc: validation.IntBetween(1, 3600),
						},
					},
				},
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},

		CustomizeDiff: verify.SetTagsDiff,
	}
}

func resourceDataQualityJobDefinitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).SageMakerConn(ctx)

	var name string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	} else {
		name = id.UniqueId()
	}

	var roleArn string
	if v, ok := d.GetOk("role_arn"); ok {
		roleArn = v.(string)
	}

	createOpts := &sagemaker.CreateDataQualityJobDefinitionInput{
		DataQualityAppSpecification: expandDataQualityAppSpecification(d.Get("data_quality_app_specification").([]interface{})),
		DataQualityJobInput:xpandDataQualityJobInput(d.Get("data_quality_job_input").([]interface{})),
		DataQualityJobOutputConfig:  expandMonitoringOutputConfig(d.Get("data_quality_job_output_config").([]interface{})),
		JobDefinitionName:g(name),
		JobResources:ngResources(d.Get("job_resources").([]interface{})),
		RoleArn:,
		Tags:
	}

	if v, ok := d.GetOk("data_quality_baseline_config"); ok && len(v.([]interface{})) > 0 {
		createOpts.DataQualityBaselineConfig = expandDataQualityBaselineConfig(v.([]interface{}))
	}

	if v, ok := d.GetOk("network_config"); ok && len(v.([]interface{})) > 0 {
		createOpts.NetworkConfig = expandMonitoringNetworkConfig(v.([]interface{}))
	}

	if v, ok := d.GetOk("stopping_condition"); ok && len(v.([]interface{})) > 0 {
		createOpts.StoppingCondition = expandMonitoringStoppingCondition(v.([]interface{}))
	}

	_, err := conn.CreateDataQualityJobDefinitionWithContext(ctx, createOpts)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating SageMaker Data Quality Job Definition (%s): %s", name, err)
	}

	d.SetId(name)

	return append(diags, resourceDataQualityJobDefinitionRead(ctx, d, meta)...)
}

func resourceDataQualityJobDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	jobDefinition, err := FindDataQualityJobDefinitionByName(ctx, conn, d.Id())

	if !d.IsNewResource() && tfawserr.ErrCodeEquals(err, sagemaker.ErrCodeResourceNotFound) {
		log.Printf("[WARN] SageMaker Data Quality Job Definition (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	d.Set("arn", jobDefinition.JobDefinitionArn)
	d.Set("name", jobDefinition.JobDefinitionName)
	d.Set("role_arn", jobDefinition.RoleArn)

	if err := d.Set("data_quality_app_specification", flattenDataQualityAppSpecification(jobDefinition.DataQualityAppSpecification)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting data_quality_app_specification for SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	if err := d.Set("data_quality_baseline_config", flattenDataQualityBaselineConfig(jobDefinition.DataQualityBaselineConfig)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting data_quality_baseline_config for SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	if err := d.Set("data_quality_job_input", flattenDataQualityJobInput(jobDefinition.DataQualityJobInput)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting data_quality_job_input for SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	if err := d.Set("data_quality_job_output_config", flattenMonitoringOutputConfig(jobDefinition.DataQualityJobOutputConfig)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting data_quality_job_output_config for SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	if err := d.Set("job_resources", flattenMonitoringResources(jobDefinition.JobResources)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting job_resources for SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	if err := d.Set("network_config", flattenMonitoringNetworkConfig(jobDefinition.NetworkConfig)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting network_config for SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	if err := d.Set("stopping_condition", flattenMonitoringStoppingCondition(jobDefinition.StoppingCondition)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting stopping_condition for SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	return diags
}

func resourceDataQualityJobDefinitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

func
	return append(diags, resourceDataQualityJobDefinitionRead(ctx, d, meta)...)
}

func resourceDataQualityJobDefinitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).SageMakerConn(ctx)

funcerr := conn.DeleteDataQualityJobDefinitionWithContext(ctx, &sagemaker.DeleteDataQualityJobDefinitionInput{
		JobDefinitionName: aws.String(d.Id()),
	})

	if tfawserr.ErrCodeEquals(err, sagemaker.ErrCodeResourceNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting SageMaker Data Quality Job Definition (%s): %s", d.Id(), err)
	}

	return diags
}

func FindDataQualityJobDefinitionByName(ctx context.Context, conn *sagemaker.SageMaker, name string) (*sagemaker.DescribeDataQualityJobDefinitionOutput, error) {
	input := &sagemaker.DescribeDataQualityJobDefinitionInput{
		JobDefinitionName: aws.String(name),
	}

func
	if tfawserr.ErrCodeEquals(err, sagemaker.ErrCodeResourceNotFound) {
		return nil, &retry.NotFoundError{
			LastError:,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output, nil
}

func flattenDataQualityAppSpecification(config *sagemaker.DataQualityAppSpecification) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}
funcconfig.ImageUri != nil {
		m["image_uri"] = aws.StringValue(config.ImageUri)
	}

	if config.Environment != nil {
		m["environment"] = aws.StringValueMap(config.Environment)
	}

	if config.PostAnalyticsProcessorSourceUri != nil {
		m["post_analytics_processor_source_uri"] = aws.StringValue(config.PostAnalyticsProcessorSourceUri)
	}

	if config.RecordPreprocessorSourceUri != nil {
		m["record_preprocessor_source_uri"] = aws.StringValue(config.RecordPreprocessorSourceUri)
	}

	return []map[string]interface{}{m}
}

func flattenDataQualityBaselineConfig(config *sagemaker.DataQualityBaselineConfig) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

func"constraints_resource"] = flattenConstraintsResource(config.ConstraintsResource)
	}

	if config.StatisticsResource != nil {
		m["statistics_resource"] = flattenMonitoringStatisticsResource(config.StatisticsResource)
	}

	return []map[string]interface{}{m}
}

func flattenConstraintsResource(config *sagemaker.MonitoringConstraintsResource) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.S3Uri != nil {
func

	return []map[string]interface{}{m}
}

func flattenMonitoringStatisticsResource(config *sagemaker.MonitoringStatisticsResource) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.S3Uri != nil {
		m["s3_uri"] = aws.StringValue(config.S3Uri)
func
	return []map[string]interface{}{m}
}

func flattenDataQualityJobInput(config *sagemaker.DataQualityJobInput) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.EndpointInput != nil {
		m["endpoint_input"] = flattenEndpointInput(config.EndpointInput)
	}
funcconfig.BatchTransformInput != nil {
		m["batch_transform_input"] = flattenBatchTransformInput(config.BatchTransformInput)
	}

	return []map[string]interface{}{m}
}

func flattenBatchTransformInput(config *sagemaker.BatchTransformInput_) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.LocalPath != nil {
		m["local_path"] = aws.StringValue(config.LocalPath)
	}

func"data_captured_destination_s3_uri"] = aws.StringValue(config.DataCapturedDestinationS3Uri)
	}

	if config.DatasetFormat != nil {
		m["dataset_format"] = flattenMonitoringDatasetFormat(config.DatasetFormat)
	}

	if config.S3DataDistributionType != nil {
		m["s3_data_distribution_type"] = aws.StringValue(config.S3DataDistributionType)
	}

	if config.S3InputMode != nil {
		m["s3_input_mode"] = aws.StringValue(config.S3InputMode)
	}

	return []map[string]interface{}{m}
}

func flattenMonitoringDatasetFormat(config *sagemaker.MonitoringDatasetFormat) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.Csv != nil {
		m["csv"] = flattenMonitoringCSVDatasetFormat(config.Csv)
	}

	if config.Json != nil {
func

	return []map[string]interface{}{m}
}

func flattenMonitoringCSVDatasetFormat(config *sagemaker.MonitoringCsvDatasetFormat) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.Header != nil {
		m["header"] = aws.BoolValue(config.Header)
	}

	return []map[string]interface{}{m}
}
func flattenMonitoringJSONDatasetFormat(config *sagemaker.MonitoringJsonDatasetFormat) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.Line != nil {
		m["line"] = aws.BoolValue(config.Line)
	}

	return []map[string]interface{}{m}
}

funcconfig == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.EndpointName != nil {
		m["endpoint_name"] = aws.StringValue(config.EndpointName)
	}

	if config.LocalPath != nil {
		m["local_path"] = aws.StringValue(config.LocalPath)
	}

func"s3_data_distribution_type"] = aws.StringValue(config.S3DataDistributionType)
	}

	if config.S3InputMode != nil {
		m["s3_input_mode"] = aws.StringValue(config.S3InputMode)
	}

	return []map[string]interface{}{m}
}

func flattenMonitoringOutputConfig(config *sagemaker.MonitoringOutputConfig) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.KmsKeyId != nil {
		m["kms_key_id"] = aws.StringValue(config.KmsKeyId)
	}

	if config.MonitoringOutputs != nil {
		m["monitoring_outputs"] = flattenMonitoringOutputs(config.MonitoringOutputs)
	}

	return []map[string]interface{}{m}
func
func flattenMonitoringOutputs(list []*sagemaker.MonitoringOutput) []map[string]interface{} {
	outputs := make([]map[string]interface{}, 0, len(list))

	for _, lRaw := range list {
		m := make(map[string]interface{})
		m["s3_output"] = flattenMonitoringS3Output(lRaw.S3Output)
		outputs = append(outputs, m)
	}

	return outputs
}

func flattenMonitoringS3Output(config *sagemaker.MonitoringS3Output) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

func
	if config.LocalPath != nil {
		m["local_path"] = aws.StringValue(config.LocalPath)
	}

	if config.S3UploadMode != nil {
		m["s3_upload_mode"] = aws.StringValue(config.S3UploadMode)
	}

	if config.S3Uri != nil {
		m["s3_uri"] = aws.StringValue(config.S3Uri)
	}
funcurn []map[string]interface{}{m}
}

func flattenMonitoringResources(config *sagemaker.MonitoringResources) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.ClusterConfig != nil {
		m["cluster_config"] = flattenMonitoringClusterConfig(config.ClusterConfig)
	}

	return []map[string]interface{}{m}
}

func flattenMonitoringClusterConfig(config *sagemaker.MonitoringClusterConfig) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

func
	if config.InstanceCount != nil {
		m["instance_count"] = aws.Int64Value(config.InstanceCount)
	}

	if config.InstanceType != nil {
		m["instance_type"] = aws.StringValue(config.InstanceType)
	}

	if config.VolumeKmsKeyId != nil {
		m["volume_kms_key_id"] = aws.StringValue(config.VolumeKmsKeyId)
	}

	if config.VolumeSizeInGB != nil {
func

	return []map[string]interface{}{m}
}

func flattenMonitoringNetworkConfig(config *sagemaker.MonitoringNetworkConfig) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.EnableInterContainerTrafficEncryption != nil {
		m["enable_inter_container_traffic_encryption"] = aws.BoolValue(config.EnableInterContainerTrafficEncryption)
	}

	if config.EnableNetworkIsolation != nil {
		m["enable_network_isolation"] = aws.BoolValue(config.EnableNetworkIsolation)
	}

	if config.VpcConfig != nil {
		m["vpc_config"] = flattenVPCConfig(config.VpcConfig)
	}

	return []map[string]interface{}{m}
}
func flattenVPCConfig(config *sagemaker.VpcConfig) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{}

	if config.SecurityGroupIds != nil {
		m["security_group_ids"] = flex.FlattenStringSet(config.SecurityGroupIds)
	}

	if config.Subnets != nil {
		m["subnets"] = flex.FlattenStringSet(config.Subnets)
	}

	return []map[string]interface{}{m}
}

func flattenMonitoringStoppingCondition(config *sagemaker.MonitoringStoppingCondition) []map[string]interface{} {
	if config == nil {
		return []map[string]interface{}{}
	}
func= map[string]interface{}{}

	if config.MaxRuntimeInSeconds != nil {
		m["max_runtime_in_seconds"] = aws.Int64Value(config.MaxRuntimeInSeconds)
	}

	return []map[string]interface{}{m}
}

func expandDataQualityAppSpecification(configured []interface{}) *sagemaker.DataQualityAppSpecification {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.DataQualityAppSpecification{}

funcImageUri = aws.String(v)
	}

	if v, ok := m["environment"].(map[string]interface{}); ok && len(v) > 0 {
		c.Environment = flex.ExpandStringMap(v)
	}

	if v, ok := m["post_analytics_processor_source_uri"].(string); ok && v != "" {
		c.PostAnalyticsProcessorSourceUri = aws.String(v)
	}

	if v, ok := m["record_preprocessor_source_uri"].(string); ok && v != "" {
		c.RecordPreprocessorSourceUri = aws.String(v)
	}
funcurn c
}

func expandDataQualityBaselineConfig(configured []interface{}) *sagemaker.DataQualityBaselineConfig {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.DataQualityBaselineConfig{}

	if v, ok := m["constraints_resource"].([]interface{}); ok && len(v) > 0 {
		c.ConstraintsResource = expandMonitoringConstraintsResource(v)
	}

	if v, ok := m["statistics_resource"].([]interface{}); ok && len(v) > 0 {
		c.StatisticsResource = expandMonitoringStatisticsResource(v)
	}

	return c
}

func expandMonitoringConstraintsResource(configured []interface{}) *sagemaker.MonitoringConstraintsResource {
	if len(configured) == 0 {
		return nil
	}

func
	c := &sagemaker.MonitoringConstraintsResource{}

	if v, ok := m["s3_uri"].(string); ok && v != "" {
		c.S3Uri = aws.String(v)
	}

	return c
}

func expandMonitoringStatisticsResource(configured []interface{}) *sagemaker.MonitoringStatisticsResource {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.MonitoringStatisticsResource{}

	if v, ok := m["s3_uri"].(string); ok && v != "" {
func

	return c
}

func expandDataQualityJobInput(configured []interface{}) *sagemaker.DataQualityJobInput {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.DataQualityJobInput{}

	if v, ok := m["endpoint_input"].([]interface{}); ok && len(v) > 0 {
		c.EndpointInput = expandEndpointInput(v)
func
	if v, ok := m["batch_transform_input"].([]interface{}); ok && len(v) > 0 {
		c.BatchTransformInput = expandBatchTransformInput(v)
	}

	return c
}

func expandEndpointInput(configured []interface{}) *sagemaker.EndpointInput {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.EndpointInput{}
funcv, ok := m["endpoint_name"].(string); ok && v != "" {
		c.EndpointName = aws.String(v)
	}

	if v, ok := m["local_path"].(string); ok && v != "" {
		c.LocalPath = aws.String(v)
	}

	if v, ok := m["s3_data_distribution_type"].(string); ok && v != "" {
		c.S3DataDistributionType = aws.String(v)
	}

	if v, ok := m["s3_input_mode"].(string); ok && v != "" {
		c.S3InputMode = aws.String(v)
	}

	return c
}

func expandBatchTransformInput(configured []interface{}) *sagemaker.BatchTransformInput_ {
functurn nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.BatchTransformInput_{}

	if v, ok := m["data_captured_destination_s3_uri"].(string); ok && v != "" {
		c.DataCapturedDestinationS3Uri = aws.String(v)
	}

	if v, ok := m["dataset_format"].([]interface{}); ok && len(v) > 0 {
		c.DatasetFormat = expandMonitoringDatasetFormat(v)
	}

	if v, ok := m["local_path"].(string); ok && v != "" {
		c.LocalPath = aws.String(v)
	}

	if v, ok := m["s3_data_distribution_type"].(string); ok && v != "" {
		c.S3DataDistributionType = aws.String(v)
	}

	if v, ok := m["s3_input_mode"].(string); ok && v != "" {
		c.S3InputMode = aws.String(v)
	}

	return c
func
func expandMonitoringDatasetFormat(configured []interface{}) *sagemaker.MonitoringDatasetFormat {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.MonitoringDatasetFormat{}

	if v, ok := m["csv"].([]interface{}); ok && len(v) > 0 {
		c.Csv = expandMonitoringCSVDatasetFormat(v)
	}

	if v, ok := m["json"].([]interface{}); ok && len(v) > 0 {
		c.Json = expandMonitoringJSONDatasetFormat(v)
	}

	return c
}

func expandMonitoringJSONDatasetFormat(configured []interface{}) *sagemaker.MonitoringJsonDatasetFormat {
	if len(configured) == 0 {
		return nil
	}

	c := &sagemaker.MonitoringJsonDatasetFormat{}

	if configured[0] == nil {
		return c
	}

funcv, ok := m["line"]; ok {
		c.Line = aws.Bool(v.(bool))
	}

	return c
}

func expandMonitoringCSVDatasetFormat(configured []interface{}) *sagemaker.MonitoringCsvDatasetFormat {
	if len(configured) == 0 {
		return nil
	}

	c := &sagemaker.MonitoringCsvDatasetFormat{}

	if configured[0] == nil {
		return c
	}

	m := configured[0].(map[string]interface{})
	if v, ok := m["header"]; ok {
func

	return c
}

func expandMonitoringOutputConfig(configured []interface{}) *sagemaker.MonitoringOutputConfig {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.MonitoringOutputConfig{}

	if v, ok := m["kms_key_id"].(string); ok && v != "" {
		c.KmsKeyId = aws.String(v)
	}

	if v, ok := m["monitoring_outputs"].([]interface{}); ok && len(v) > 0 {
func

	return c
}

func expandMonitoringOutputs(configured []interface{}) []*sagemaker.MonitoringOutput {
	containers := make([]*sagemaker.MonitoringOutput, 0, len(configured))

	for _, lRaw := range configured {
		data := lRaw.(map[string]interface{})

		l := &sagemaker.MonitoringOutput{
			S3Output: expandMonitoringS3Output(data["s3_output"].([]interface{})),
		}
		containers = append(containers, l)
	}

	return containers
}
func expandMonitoringS3Output(configured []interface{}) *sagemaker.MonitoringS3Output {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.MonitoringS3Output{}

	if v, ok := m["local_path"].(string); ok && v != "" {
		c.LocalPath = aws.String(v)
	}

	if v, ok := m["s3_upload_mode"].(string); ok && v != "" {
		c.S3UploadMode = aws.String(v)
	}

	if v, ok := m["s3_uri"].(string); ok && v != "" {
		c.S3Uri = aws.String(v)
	}
funcurn c
}

func expandMonitoringResources(configured []interface{}) *sagemaker.MonitoringResources {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.MonitoringResources{}

	if v, ok := m["cluster_config"].([]interface{}); ok && len(v) > 0 {
		c.ClusterConfig = expandMonitoringClusterConfig(v)
	}
funcurn c
}

func expandMonitoringClusterConfig(configured []interface{}) *sagemaker.MonitoringClusterConfig {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.MonitoringClusterConfig{}

	if v, ok := m["instance_count"].(int); ok && v > 0 {
		c.InstanceCount = aws.Int64(int64(v))
	}

	if v, ok := m["instance_type"].(string); ok && v != "" {
		c.InstanceType = aws.String(v)
	}

	if v, ok := m["volume_kms_key_id"].(string); ok && v != "" {
		c.VolumeKmsKeyId = aws.String(v)
	}

funcVolumeSizeInGB = aws.Int64(int64(v))
	}

	return c
}

func expandMonitoringNetworkConfig(configured []interface{}) *sagemaker.MonitoringNetworkConfig {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.MonitoringNetworkConfig{}

	if v, ok := m["enable_inter_container_traffic_encryption"]; ok {
func

	if v, ok := m["enable_network_isolation"]; ok {
		c.EnableNetworkIsolation = aws.Bool(v.(bool))
	}

	if v, ok := m["vpc_config"].([]interface{}); ok && len(v) > 0 {
		c.VpcConfig = expandVPCConfig(v)
	}

	return c
}

func expandVPCConfig(configured []interface{}) *sagemaker.VpcConfig {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.VpcConfig{}

	if v, ok := m["security_group_ids"].(*schema.Set); ok && v.Len() > 0 {
		c.SecurityGroupIds = flex.ExpandStringSet(v)
	}

	if v, ok := m["subnets"].(*schema.Set); ok && v.Len() > 0 {
		c.Subnets = flex.ExpandStringSet(v)
func
	return c
}

func expandMonitoringStoppingCondition(configured []interface{}) *sagemaker.MonitoringStoppingCondition {
	if len(configured) == 0 {
		return nil
	}

	m := configured[0].(map[string]interface{})

	c := &sagemaker.MonitoringStoppingCondition{}

	if v, ok := m["max_runtime_in_seconds"].(int); ok && v > 0 {
		c.MaxRuntimeInSeconds = aws.Int64(int64(v))
	}

	return c
}
funcfunc