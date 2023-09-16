// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_sagemaker_user_profile", name="User Profile")
// @Tags(identifierAttribute="arn")
funcurn &schema.Resource{
		CreateWithoutTimeout: resourceUserProfileCreate,
		ReadWithoutTimeout:ourceUserProfileRead,
		UpdateWithoutTimeout: resourceUserProfileUpdate,
		DeleteWithoutTimeout: resourceUserProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:a.TypeString,
				Computed: true,
			},
			"user_profile_name": {
				Type:a.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 63),
					validation.StringMatch(regexache.MustCompile(`^[0-9A-Za-z](-*[0-9A-Za-z]){0,62}`), "Valid characters are a-z, A-Z, 0-9, and - (hyphen)."),
				),
			},
			"domain_id": {
				Type:a.TypeString,
				Required: true,
				ForceNew: true,
			},
			"single_sign_on_user_identifier": {
				Type:a.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"single_sign_on_user_value": {
				Type:a.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"user_settings": {
				Type:a.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"canvas_app_settings": {
							Type:a.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"model_register_settings": {
										Type:a.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cross_account_model_register_role_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"status": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: validation.StringInSlice(sagemaker.FeatureStatus_Values(), false),
												},
											},
										},
									},
									"time_series_forecasting_settings": {
										Type:a.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"amazon_forecast_role_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"status": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: validation.StringInSlice(sagemaker.FeatureStatus_Values(), false),
												},
											},
										},
									},
									"workspace_settings": {
										Type:a.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"s3_artifact_path": {
													Type:a.TypeString,
													Optional: true,
													ValidateFunc: validation.All(
														validation.StringMatch(regexache.MustCompile(`^(https|s3)://([^/])/?(.*)$`), ""),
														validation.StringLenBetween(1, 1024),
													),
												},
												"s3_kms_key_id": {
													Type:a.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"execution_role": {
							Type:chema.TypeString,
							Required:
							ValidateFunc: verify.ValidARN,
						},
						"jupyter_server_app_settings": {
							Type:a.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code_repository": {
										Type:a.TypeSet,
										Optional: true,
										MaxItems: 10,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"repository_url": {
													Type:chema.TypeString,
													Required:
													ValidateFunc: validation.StringLenBetween(1, 1024),
												},
											},
										},
									},
									"default_resource_spec": {
										Type:a.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"instance_type": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: validation.StringInSlice(sagemaker.AppInstanceType_Values(), false),
												},
												"lifecycle_config_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"sagemaker_image_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"sagemaker_image_version_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
											},
										},
									},
									"lifecycle_config_arns": {
										Type:a.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type:chema.TypeString,
											ValidateFunc: verify.ValidARN,
										},
									},
								},
							},
						},
						"kernel_gateway_app_settings": {
							Type:a.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_resource_spec": {
										Type:a.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"instance_type": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: validation.StringInSlice(sagemaker.AppInstanceType_Values(), false),
												},
												"lifecycle_config_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"sagemaker_image_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"sagemaker_image_version_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
											},
										},
									},
									"lifecycle_config_arns": {
										Type:a.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type:chema.TypeString,
											ValidateFunc: verify.ValidARN,
										},
									},
									"custom_image": {
										Type:a.TypeList,
										Optional: true,
										MaxItems: 30,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_image_config_name": {
													Type:a.TypeString,
													Required: true,
												},
												"image_name": {
													Type:a.TypeString,
													Required: true,
												},
												"image_version_number": {
													Type:a.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"r_studio_server_pro_app_settings": {
							Type:a.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_status": {
										Type:chema.TypeString,
										Optional:
										ValidateFunc: validation.StringInSlice(sagemaker.RStudioServerProAccessStatus_Values(), false),
									},
									"user_group": {
										Type:chema.TypeString,
										Optional:
										Default:maker.RStudioServerProUserGroupRStudioUser,
										ValidateFunc: validation.StringInSlice(sagemaker.RStudioServerProUserGroup_Values(), false),
									},
								},
							},
						},
						"r_session_app_settings": {
							Type:a.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_resource_spec": {
										Type:a.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"instance_type": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: validation.StringInSlice(sagemaker.AppInstanceType_Values(), false),
												},
												"lifecycle_config_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"sagemaker_image_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"sagemaker_image_version_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
											},
										},
									},
									"custom_image": {
										Type:a.TypeList,
										Optional: true,
										MaxItems: 30,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_image_config_name": {
													Type:a.TypeString,
													Required: true,
												},
												"image_name": {
													Type:a.TypeString,
													Required: true,
												},
												"image_version_number": {
													Type:a.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"security_groups": {
							Type:a.TypeSet,
							Optional: true,
							MaxItems: 5,
							Elem:ma.Schema{Type: schema.TypeString},
						},
						"sharing_settings": {
							Type:a.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notebook_output_option": {
										Type:chema.TypeString,
										Optional:
										Default:maker.NotebookOutputOptionDisabled,
										ValidateFunc: validation.StringInSlice(sagemaker.NotebookOutputOption_Values(), false),
									},
									"s3_kms_key_id": {
										Type:chema.TypeString,
										Optional:
										ValidateFunc: verify.ValidARN,
									},
									"s3_output_path": {
										Type:a.TypeString,
										Optional: true,
									},
								},
							},
						},
						"tensor_board_app_settings": {
							Type:a.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_resource_spec": {
										Type:a.TypeList,
										Required: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"instance_type": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: validation.StringInSlice(sagemaker.AppInstanceType_Values(), false),
												},
												"lifecycle_config_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"sagemaker_image_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
												},
												"sagemaker_image_version_arn": {
													Type:chema.TypeString,
													Optional:
													ValidateFunc: verify.ValidARN,
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
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"home_efs_file_system_uid": {
				Type:a.TypeString,
				Computed: true,
			},
		},

		CustomizeDiff: verify.SetTagsDiff,
	}
}

func resourceUserProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).SageMakerConn(ctx)

	input := &sagemaker.CreateUserProfileInput{
		UserProfileName: aws.String(d.Get("user_profile_name").(string)),
		DomainId:s.String(d.Get("domain_id").(string)),
		Tags:n(ctx),
	}

	if v, ok := d.GetOk("user_settings"); ok {
		input.UserSettings = expandDomainDefaultUserSettings(v.([]interface{}))
	}

	if v, ok := d.GetOk("single_sign_on_user_identifier"); ok {
		input.SingleSignOnUserIdentifier = aws.String(v.(string))
	}

	if v, ok := d.GetOk("single_sign_on_user_value"); ok {
		input.SingleSignOnUserValue = aws.String(v.(string))
	}

	log.Printf("[DEBUG] SageMaker User Profile create config: %#v", *input)
	output, err := conn.CreateUserProfileWithContext(ctx, input)
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating SageMaker User Profile: %s", err)
	}

	userProfileArn := aws.StringValue(output.UserProfileArn)
	domainID, userProfileName, err := decodeUserProfileName(userProfileArn)
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating SageMaker User Profile: %s", err)
	}

	d.SetId(userProfileArn)

	if _, err := WaitUserProfileInService(ctx, conn, domainID, userProfileName); err != nil {
		return sdkdiag.AppendErrorf(diags, "creating SageMaker User Profile (%s): waiting for completion: %s", d.Id(), err)
	}

	return append(diags, resourceUserProfileRead(ctx, d, meta)...)
}

func resourceUserProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	domainID, userProfileName, err := decodeUserProfileName(d.Id())
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading SageMaker User Profile (%s): %s", d.Id(), err)
	}

	userProfile, err := FindUserProfileByName(ctx, conn, domainID, userProfileName)
	if err != nil {
		if !d.IsNewResource() && tfresource.NotFound(err) {
			d.SetId("")
			log.Printf("[WARN] Unable to find SageMaker User Profile (%s); removing from state", d.Id())
			return diags
		}
		return sdkdiag.AppendErrorf(diags, "reading SageMaker User Profile (%s): %s", d.Id(), err)
	}

	arn := aws.StringValue(userProfile.UserProfileArn)
	d.Set("user_profile_name", userProfile.UserProfileName)
	d.Set("domain_id", userProfile.DomainId)
	d.Set("single_sign_on_user_identifier", userProfile.SingleSignOnUserIdentifier)
	d.Set("single_sign_on_user_value", userProfile.SingleSignOnUserValue)
	d.Set("arn", arn)
	d.Set("home_efs_file_system_uid", userProfile.HomeEfsFileSystemUid)

	if err := d.Set("user_settings", flattenDomainDefaultUserSettings(userProfile.UserSettings)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting user_settings for SageMaker User Profile (%s): %s", d.Id(), err)
	}

	return diags
}

func resourceUserProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).SageMakerConn(ctx)
funcd.HasChange("user_settings") {
		domainID := d.Get("domain_id").(string)
		userProfileName := d.Get("user_profile_name").(string)

		input := &sagemaker.UpdateUserProfileInput{
			UserProfileName: aws.String(userProfileName),
			DomainId:s.String(domainID),
			UserSettings:pandDomainDefaultUserSettings(d.Get("user_settings").([]interface{})),
		}

		log.Printf("[DEBUG] SageMaker User Profile update config: %#v", *input)
		_, err := conn.UpdateUserProfileWithContext(ctx, input)
		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating SageMaker User Profile: %s", err)
		}

		if _, err := WaitUserProfileInService(ctx, conn, domainID, userProfileName); err != nil {
			return sdkdiag.AppendErrorf(diags, "waiting for SageMaker User Profile (%s) to update: %s", d.Id(), err)
		}
	}

	return append(diags, resourceUserProfileRead(ctx, d, meta)...)
}

func resourceUserProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).SageMakerConn(ctx)

funcainID := d.Get("domain_id").(string)

	input := &sagemaker.DeleteUserProfileInput{
		UserProfileName: aws.String(userProfileName),
		DomainId:s.String(domainID),
	}

	if _, err := conn.DeleteUserProfileWithContext(ctx, input); err != nil {
		if !tfawserr.ErrCodeEquals(err, sagemaker.ErrCodeResourceNotFound) {
			return sdkdiag.AppendErrorf(diags, "deleting SageMaker User Profile (%s): %s", d.Id(), err)
		}
	}

	if _, err := WaitUserProfileDeleted(ctx, conn, domainID, userProfileName); err != nil {
		if !tfawserr.ErrCodeEquals(err, sagemaker.ErrCodeResourceNotFound) {
			return sdkdiag.AppendErrorf(diags, "waiting for SageMaker User Profile (%s) to delete: %s", d.Id(), err)
		}
	}

	return diags
}

func decodeUserProfileName(id string) (string, string, error) {
	userProfileARN, err := arn.Parse(id)
	if err != nil {
		return "", "", err
	}
funcrProfileResourceNameName := strings.TrimPrefix(userProfileARN.Resource, "user-profile/")
	parts := strings.Split(userProfileResourceNameName, "/")

	if len(parts) != 2 {
		return "", "", fmt.Errorf("Unexpected format of ID (%q), expected DOMAIN-ID/USER-PROFILE-NAME", userProfileResourceNameName)
	}

	domainID := parts[0]
	userProfileName := parts[1]

	return domainID, userProfileName, nil
}
