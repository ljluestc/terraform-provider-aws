//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagemedialive

import(
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

//@SDKResource("aws_medialive_channel",name="Channel")
//@Tags(identifierAttribute="arn")
funcResourceChannel()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceChannelCreate,
		ReadWithoutTimeout:resourceChannelRead,
		UpdateWithoutTimeout:resourceChannelUpdate,
		DeleteWithoutTimeout:resourceChannelDelete,

		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},

		Timeouts:&schema.ResourceTimeout{
			Create:schema.DefaultTimeout(15*time.Minute),
			Update:schema.DefaultTimeout(15*time.Minute),
			Delete:schema.DefaultTimeout(15*time.Minute),
		},

		SchemaFunc:func()map[string]*schema.Schema{
			returnmap[string]*schema.Schema{
				"arn":{
					Type:schema.TypeString,
					Computed:true,
				},
				"cdi_input_specification":{
					Type:schema.TypeList,
					Optional:true,
					MaxItems:1,
					Elem:&schema.Resource{
						Schema:map[string]*schema.Schema{
							"resolution":{
								Type:schema.TypeString,
								Required:true,
								ValidateDiagFunc:enum.Validate[types.CdiInputResolution](),
							},
						},
					},
				},
				"channel_class":{
					Type:schema.TypeString,
					Required:true,
					ForceNew:true,
					ValidateDiagFunc:enum.Validate[types.ChannelClass](),
				},
				"channel_id":{
					Type:schema.TypeString,
					Computed:true,
				},
				"destinations":{
					Type:schema.TypeSet,
					Required:true,
					MinItems:1,
					Elem:&schema.Resource{
						Schema:map[string]*schema.Schema{
							"id":{
								Type:schema.TypeString,
								Required:true,
							},
							"media_package_settings":{
								Type:schema.TypeSet,
								Optional:true,
								Elem:&schema.Resource{
									Schema:map[string]*schema.Schema{
										"channel_id":{
											Type:schema.TypeString,
											Required:true,
										},
									},
								},
							},
							"multiplex_settings":{
								Type:schema.TypeList,
								Optional:true,
								MaxItems:1,
								Elem:&schema.Resource{
									Schema:map[string]*schema.Schema{
										"multiplex_id":{
											Type:schema.TypeString,
											Required:true,
										},
										"program_name":{
											Type:schema.TypeString,
											Required:true,
										},
									},
								},
							},
							"settings":{
								Type:schema.TypeSet,
								Optional:true,
								Elem:&schema.Resource{
									Schema:map[string]*schema.Schema{
										"password_param":{
											Type:schema.TypeString,
											Optional:true,
										},
										"stream_name":{
											Type:schema.TypeString,
											Optional:true,
										},
										"url":{
											Type:schema.TypeString,
											Optional:true,
										},
										"username":{
											Type:schema.TypeString,
											Optional:true,
										},
									},
								},
							},
						},
					},
				},
				"encoder_settings":func()*schema.Schema{
					returnchannelEncoderSettingsSchema()
				}(),
				"input_attachments":{
					Type:schema.TypeSet,
					Required:true,
					Elem:&schema.Resource{
						Schema:map[string]*schema.Schema{
							"automatic_input_failover_settings":{
								Type:schema.TypeList,
								Optional:true,
								MaxItems:1,
								Elem:&schema.Resource{
									Schema:map[string]*schema.Schema{
										"secondary_input_id":{
											Type:schema.TypeString,
											Required:true,
										},
										"error_clear_time_msec":{
											Type:schema.TypeInt,
											Optional:true,
										},
										"failover_condition":{
											Type:schema.TypeSet,
											Optional:true,
											Elem:&schema.Resource{
												Schema:map[string]*schema.Schema{
													"failover_condition_settings":{
														Type:schema.TypeList,
														Optional:true,
														MaxItems:1,
														Elem:&schema.Resource{
															Schema:map[string]*schema.Schema{
																"audio_silence_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"audio_selector_name":{
																				Type:schema.TypeString,
																				Required:true,
																			},
																			"audio_silence_threshold_msec":{
																				Type:schema.TypeInt,
																				Optional:true,
																			},
																		},
																	},
																},
																"input_loss_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"input_loss_threshold_msec":{
																				Type:schema.TypeInt,
																				Optional:true,
																			},
																		},
																	},
																},
																"video_black_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"black_detect_threshold":{
																				Type:schema.TypeFloat,
																				Optional:true,
																			},
																			"video_black_threshold_msec":{
																				Type:schema.TypeInt,
																				Optional:true,
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
										"input_preference":{
											Type:schema.TypeString,
											Optional:true,
											ValidateDiagFunc:enum.Validate[types.InputPreference](),
										},
									},
								},
							},
							"input_attachment_name":{
								Type:schema.TypeString,
								Required:true,
							},
							"input_id":{
								Type:schema.TypeString,
								Required:true,
							},
							"input_settings":{
								Type:schema.TypeList,
								Optional:true,
								Computed:true,
								MaxItems:1,
								Elem:&schema.Resource{
									Schema:map[string]*schema.Schema{
										"audio_selector":{
											Type:schema.TypeList,
											Optional:true,
											Elem:&schema.Resource{
												Schema:map[string]*schema.Schema{
													"name":{
														Type:schema.TypeString,
														Required:true,
													},
													"selector_settings":{
														Type:schema.TypeList,
														Optional:true,
														MaxItems:1,
														Elem:&schema.Resource{
															Schema:map[string]*schema.Schema{
																"audio_hls_rendition_selection":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"group_id":{
																				Type:schema.TypeString,
																				Required:true,
																			},
																			"name":{
																				Type:schema.TypeString,
																				Required:true,
																			},
																		},
																	},
																},
																"audio_language_selection":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"language_code":{
																				Type:schema.TypeString,
																				Required:true,
																			},
																			"language_selection_policy":{
																				Type:schema.TypeString,
																				Optional:true,
																				ValidateDiagFunc:enum.Validate[types.AudioLanguageSelectionPolicy](),
																			},
																		},
																	},
																},
																"audio_pid_selection":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"pid":{
																				Type:schema.TypeInt,
																				Required:true,
																			},
																		},
																	},
																},
																"audio_track_selection":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"dolby_e_decode":{
																				Type:schema.TypeList,
																				Optional:true,
																				MaxItems:1,
																				Elem:&schema.Resource{
																					Schema:map[string]*schema.Schema{
																						"program_selection":{
																							Type:schema.TypeString,
																							Required:true,
																							ValidateDiagFunc:enum.Validate[types.DolbyEProgramSelection](),
																						},
																					},
																				},
																			},
																			"tracks":{
																				Type:schema.TypeSet,
																				Required:true,
																				Elem:&schema.Resource{
																					Schema:map[string]*schema.Schema{
																						"track":{
																							Type:schema.TypeInt,
																							Required:true,
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
												},
											},
										},
										"caption_selector":{
											Type:schema.TypeList,
											Optional:true,
											Elem:&schema.Resource{
												Schema:map[string]*schema.Schema{
													"name":{
														Type:schema.TypeString,
														Required:true,
													},
													"language_code":{
														Type:schema.TypeString,
														Optional:true,
													},
													"selector_settings":{
														Type:schema.TypeList,
														Optional:true,
														MaxItems:1,
														Elem:&schema.Resource{
															Schema:map[string]*schema.Schema{
																"ancillary_source_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"source_ancillary_channel_number":{
																				Type:schema.TypeInt,
																				Optional:true,
																			},
																		},
																	},
																},
																"arib_source_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{},//noexportedelementsinthislist
																	},
																},
																"dvb_sub_source_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"ocr_language":{
																				Type:schema.TypeString,
																				Optional:true,
																				ValidateDiagFunc:enum.Validate[types.DvbSubOcrLanguage](),
																			},
																			"pid":{
																				Type:schema.TypeInt,
																				Optional:true,
																				ValidateFunc:validation.IntAtLeast(1),
																			},
																		},
																	},
																},
																"embedded_source_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"convert_608_to_708":{
																				Type:schema.TypeString,
																				Optional:true,
																				ValidateDiagFunc:enum.Validate[types.EmbeddedConvert608To708](),
																			},
																			"scte20_detection":{
																				Type:schema.TypeString,
																				Optional:true,
																				ValidateDiagFunc:enum.Validate[types.EmbeddedScte20Detection](),
																			},
																			"source_608_channel_number":{
																				Type:schema.TypeInt,
																				Optional:true,
																			},
																		},
																	},
																},
																"scte20_source_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"convert_608_to_708":{
																				Type:schema.TypeString,
																				Optional:true,
																				ValidateDiagFunc:enum.Validate[types.Scte20Convert608To708](),
																			},
																			"source_608_channel_number":{
																				Type:schema.TypeInt,
																				Optional:true,
																			},
																		},
																	},
																},
																"scte27_source_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"ocr_language":{
																				Type:schema.TypeString,
																				Optional:true,
																				ValidateDiagFunc:enum.Validate[types.Scte27OcrLanguage](),
																			},
																			"pid":{
																				Type:schema.TypeInt,
																				Optional:true,
																			},
																		},
																	},
																},
																"teletext_source_settings":{
																	Type:schema.TypeList,
																	Optional:true,
																	MaxItems:1,
																	Elem:&schema.Resource{
																		Schema:map[string]*schema.Schema{
																			"output_rectangle":{
																				Type:schema.TypeList,
																				Optional:true,
																				MaxItems:1,
																				Elem:&schema.Resource{
																					Schema:map[string]*schema.Schema{
																						"height":{
																							Type:schema.TypeFloat,
																							Required:true,
																						},
																						"left_offset":{
																							Type:schema.TypeFloat,
																							Required:true,
																						},
																						"top_offset":{
																							Type:schema.TypeFloat,
																							Required:true,
																						},
																						"width":{
																							Type:schema.TypeFloat,
																							Required:true,
																						},
																					},
																				},
																			},
																			"page_number":{
																				Type:schema.TypeString,
																				Optional:true,
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
										"deblock_filter":{
											Type:schema.TypeString,
											Optional:true,
											ValidateDiagFunc:enum.Validate[types.InputDeblockFilter](),
										},
										"denoise_filter":{
											Type:schema.TypeString,
											Optional:true,
											ValidateDiagFunc:enum.Validate[types.InputDenoiseFilter](),
										},
										"filter_strength":{
											Type:schema.TypeInt,
											Optional:true,
											ValidateDiagFunc:validation.ToDiagFunc(validation.IntBetween(1,5)),
										},
										"input_filter":{
											Type:schema.TypeString,
											Optional:true,
											Computed:true,
											ValidateDiagFunc:enum.Validate[types.InputFilter](),
										},
										"network_input_settings":{
											Type:schema.TypeList,
											Optional:true,
											MaxItems:1,
											Elem:&schema.Resource{
												Schema:map[string]*schema.Schema{
													"hls_input_settings":{
														Type:schema.TypeList,
														Optional:true,
														MaxItems:1,
														Elem:&schema.Resource{
															Schema:map[string]*schema.Schema{
																"bandwidth":{
																	Type:schema.TypeInt,
																	Optional:true,
																},
																"buffer_segments":{
																	Type:schema.TypeInt,
																	Optional:true,
																},
																"retries":{
																	Type:schema.TypeInt,
																	Optional:true,
																},
																"retry_interval":{
																	Type:schema.TypeInt,
																	Optional:true,
																},
																"scte35_source":{
																	Type:schema.TypeString,
																	Optional:true,
																	ValidateDiagFunc:enum.Validate[types.HlsScte35SourceType](),
																},
															},
														},
													},
													"server_validation":{
														Type:schema.TypeString,
														Optional:true,
														ValidateDiagFunc:enum.Validate[types.NetworkInputServerValidation](),
													},
												},
											},
										},
										"scte35_pid":{
											Type:schema.TypeInt,
											Optional:true,
										},
										"smpte2038_data_preference":{
											Type:schema.TypeString,
											Optional:true,
											ValidateDiagFunc:enum.Validate[types.Smpte2038DataPreference](),
										},
										"source_end_behavior":{
											Type:schema.TypeString,
											Optional:true,
											ValidateDiagFunc:enum.Validate[types.InputSourceEndBehavior](),
										},
										"video_selector":{
											Type:schema.TypeList,
											Optional:true,
											MaxItems:1,
											Elem:&schema.Resource{
												Schema:map[string]*schema.Schema{
													"color_space":{
														Type:schema.TypeString,
														Optional:true,
														ValidateDiagFunc:enum.Validate[types.VideoSelectorColorSpace](),
													},
													//TODOimplementcolor_space_settings
													"color_space_usage":{
														Type:schema.TypeString,
														Optional:true,
														ValidateDiagFunc:enum.Validate[types.VideoSelectorColorSpaceUsage](),
													},
													//TODOimplementselector_settings
												},
											},
										},
									},
								},
							},
						},
					},
				},
				"input_specification":{
					Type:schema.TypeList,
					Required:true,
					MaxItems:1,
					Elem:&schema.Resource{
						Schema:map[string]*schema.Schema{
							"codec":{
								Type:schema.TypeString,
								Required:true,
								ValidateDiagFunc:enum.Validate[types.InputCodec](),
							},
							"maximum_bitrate":{
								Type:schema.TypeString,
								Required:true,
								ValidateDiagFunc:enum.Validate[types.InputMaximumBitrate](),
							},
							"input_resolution":{
								Type:schema.TypeString,
								Required:true,
								ValidateDiagFunc:enum.Validate[types.InputResolution](),
							},
						},
					},
				},
				"log_level":{
					Type:schema.TypeString,
					Optional:true,
					Computed:true,
					ValidateDiagFunc:enum.Validate[types.LogLevel](),
				},
				"maintenance":{
					Type:schema.TypeList,
					Optional:true,
					Computed:true,
					MaxItems:1,
					Elem:&schema.Resource{
						Schema:map[string]*schema.Schema{
							"maintenance_day":{
								Type:schema.TypeString,
								Required:true,
								ValidateDiagFunc:enum.Validate[types.MaintenanceDay](),
							},
							"maintenance_start_time":{
								Type:schema.TypeString,
								Required:true,
							},
						},
					},
				},
				"name":{
					Type:schema.TypeString,
					Required:true,
				},
				"role_arn":{
					Type:schema.TypeString,
					Optional:true,
					ValidateDiagFunc:validation.ToDiagFunc(verify.ValidARN),
				},
				"start_channel":{
					Type:schema.TypeBool,
					Optional:true,
					Default:false,
				},
				"vpc":{
					Type:schema.TypeList,
					Optional:true,
					MaxItems:1,
					ForceNew:true,
					Elem:&schema.Resource{
						Schema:map[string]*schema.Schema{
							"availability_zones":{
								Type:schema.TypeList,
								Computed:true,
								Elem:&schema.Schema{Type:schema.TypeString},
							},
							"public_address_allocation_ids":{
								Type:schema.TypeList,
								Required:true,
								Elem:&schema.Schema{Type:schema.TypeString},
							},
							"security_group_ids":{
								Type:schema.TypeList,
								Optional:true,
								Computed:true,
								MaxItems:5,
								Elem:&schema.Schema{Type:schema.TypeString},
							},
							"subnet_ids":{
								Type:schema.TypeList,
								Required:true,
								Elem:&schema.Schema{Type:schema.TypeString},
							},
						},
					},
				},
				names.AttrTags:tftags.TagsSchema(),
				names.AttrTagsAll:tftags.TagsSchemaComputed(),
			}
		},

		CustomizeDiff:verify.SetTagsDiff,
	}
}

const(
	ResNameChannel="Channel"
)

funcresourceChannelCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).MediaLiveClient(ctx)

	in:=&medialive.CreateChannelInput{
		Name:ring(d.Get("name").(string)),
		RequestId:aws.String(id.UniqueId()),
		Tags:sIn(ctx),
	}

	ifv,ok:=d.GetOk("cdi_input_specification");ok&&len(v.([]interface{}))>0{
		in.CdiInputSpecification=expandChannelCdiInputSpecification(v.([]interface{}))
	}
	ifv,ok:=d.GetOk("channel_class");ok{
		in.ChannelClass=types.ChannelClass(v.(string))
	}
	ifv,ok:=d.GetOk("destinations");ok&&v.(*schema.Set).Len()>0{
		in.Destinations=expandChannelDestinations(v.(*schema.Set).List())
	}
	ifv,ok:=d.GetOk("encoder_settings");ok&&len(v.([]interface{}))>0{
		in.EncoderSettings=expandChannelEncoderSettings(v.([]interface{}))
	}
	ifv,ok:=d.GetOk("input_attachments");ok&&v.(*schema.Set).Len()>0{
		in.InputAttachments=expandChannelInputAttachments(v.(*schema.Set).List())
	}
	ifv,ok:=d.GetOk("input_specification");ok&&len(v.([]interface{}))>0{
		in.InputSpecification=expandChannelInputSpecification(v.([]interface{}))
	}
	ifv,ok:=d.GetOk("maintenance");ok&&len(v.([]interface{}))>0{
		in.Maintenance=expandChannelMaintenanceCreate(v.([]interface{}))
	}
	ifv,ok:=d.GetOk("role_arn");ok{
		in.RoleArn=aws.String(v.(string))
	}
	ifv,ok:=d.GetOk("vpc");ok&&len(v.([]interface{}))>0{
		in.Vpc=expandChannelVPC(v.([]interface{}))
	}

	out,err:=conn.CreateChannel(ctx,in)
	iferr!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionCreating,ResNameChannel,d.Get("name").(string),err)
	}

	ifout==nil||out.Channel==nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionCreating,ResNameChannel,d.Get("name").(string),errors.New("emptyoutput"))
	}

	d.SetId(aws.ToString(out.Channel.Id))

	if_,err:=waitChannelCreated(ctx,conn,d.Id(),d.Timeout(schema.TimeoutCreate));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionWaitingForCreation,ResNameChannel,d.Id(),err)
	}

	ifd.Get("start_channel").(bool){
		iferr:=startChannel(ctx,conn,d.Timeout(schema.TimeoutCreate),d.Id());err!=nil{
			returncreate.DiagError(names.MediaLive,create.ErrActionCreating,ResNameChannel,d.Get("name").(string),err)
		}
	}

	returnresourceChannelRead(ctx,d,meta)
}

funcresourceChannelRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).MediaLiveClient(ctx)

	out,err:=FindChannelByID(ctx,conn,d.Id())

	if!d.IsNewResource()&&tfresource.NotFound(err){
		log.Printf("[WARN]MediaLiveChannel(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returnnil
	}

	iferr!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionReading,ResNameChannel,d.Id(),err)
	}

	d.Set("arn",out.Arn)
	d.Set("name",out.Name)
	d.Set("channel_class",out.ChannelClass)
	d.Set("channel_id",out.Id)
	d.Set("log_level",out.LogLevel)
	d.Set("role_arn",out.RoleArn)

	iferr:=d.Set("cdi_input_specification",flattenChannelCdiInputSpecification(out.CdiInputSpecification));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("input_attachments",flattenChannelInputAttachments(out.InputAttachments));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("destinations",flattenChannelDestinations(out.Destinations));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("encoder_settings",flattenChannelEncoderSettings(out.EncoderSettings));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("input_specification",flattenChannelInputSpecification(out.InputSpecification));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("maintenance",flattenChannelMaintenance(out.Maintenance));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}
	iferr:=d.Set("vpc",flattenChannelVPC(out.Vpc));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionSetting,ResNameChannel,d.Id(),err)
	}

	returnnil
}

funcresourceChannelUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).MediaLiveClient(ctx)

	ifd.HasChangesExcept("tags","tags_all","start_channel"){
		in:=&medialive.UpdateChannelInput{
			ChannelId:aws.String(d.Id()),
		}

		ifd.HasChange("name"){
			in.Name=aws.String(d.Get("name").(string))
		}

		ifd.HasChange("cdi_input_specification"){
			in.CdiInputSpecification=expandChannelCdiInputSpecification(d.Get("cdi_input_specification").([]interface{}))
		}

		ifd.HasChange("destinations"){
			in.Destinations=expandChannelDestinations(d.Get("destinations").(*schema.Set).List())
		}

		ifd.HasChange("encoder_settings"){
			in.EncoderSettings=expandChannelEncoderSettings(d.Get("encoder_settings").([]interface{}))
		}

		ifd.HasChange("input_attachments"){
			in.InputAttachments=expandChannelInputAttachments(d.Get("input_attachments").(*schema.Set).List())
		}

		ifd.HasChange("input_specification"){
			in.InputSpecification=expandChannelInputSpecification(d.Get("input_specification").([]interface{}))
		}

		ifd.HasChange("log_level"){
			in.LogLevel=types.LogLevel(d.Get("log_level").(string))
		}

		ifd.HasChange("maintenance"){
			in.Maintenance=expandChannelMaintenanceUpdate(d.Get("maintenance").([]interface{}))
		}

		ifd.HasChange("role_arn"){
			in.RoleArn=aws.String(d.Get("role_arn").(string))
		}

		channel,err:=FindChannelByID(ctx,conn,d.Id())

		iferr!=nil{
			returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
		}

		ifchannel.State==types.ChannelStateRunning{
			iferr:=stopChannel(ctx,conn,d.Timeout(schema.TimeoutUpdate),d.Id());err!=nil{
				returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
			}
		}

		out,err:=conn.UpdateChannel(ctx,in)
		iferr!=nil{
			returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
		}

		if_,err:=waitChannelUpdated(ctx,conn,aws.ToString(out.Channel.Id),d.Timeout(schema.TimeoutUpdate));err!=nil{
			returncreate.DiagError(names.MediaLive,create.ErrActionWaitingForUpdate,ResNameChannel,d.Id(),err)
		}
	}

	ifd.Get("start_channel").(bool){
		iferr:=startChannel(ctx,conn,d.Timeout(schema.TimeoutUpdate),d.Id());err!=nil{
			returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Get("name").(string),err)
		}
	}

	ifd.HasChange("start_channel"){
		channel,err:=FindChannelByID(ctx,conn,d.Id())

		iferr!=nil{
			returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
		}

		switchd.Get("start_channel").(bool){
		casetrue:
			ifchannel.State==types.ChannelStateIdle{
				iferr:=startChannel(ctx,conn,d.Timeout(schema.TimeoutUpdate),d.Id());err!=nil{
					returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
				}
			}
		default:
			ifchannel.State==types.ChannelStateRunning{
				iferr:=stopChannel(ctx,conn,d.Timeout(schema.TimeoutUpdate),d.Id());err!=nil{
					returncreate.DiagError(names.MediaLive,create.ErrActionUpdating,ResNameChannel,d.Id(),err)
				}
			}
		}
	}

	returnresourceChannelRead(ctx,d,meta)
}

funcresourceChannelDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).MediaLiveClient(ctx)

	log.Printf("[INFO]DeletingMediaLiveChannel%s",d.Id())

	channel,err:=FindChannelByID(ctx,conn,d.Id())

	iferr!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionDeleting,ResNameChannel,d.Id(),err)
	}

	ifchannel.State==types.ChannelStateRunning{
		iferr:=stopChannel(ctx,conn,d.Timeout(schema.TimeoutDelete),d.Id());err!=nil{
			returncreate.DiagError(names.MediaLive,create.ErrActionDeleting,ResNameChannel,d.Id(),err)
		}
	}

	_,err=conn.DeleteChannel(ctx,&medialive.DeleteChannelInput{
		ChannelId:aws.String(d.Id()),
	})

	iferr!=nil{
		varnfe*types.NotFoundException
		iferrors.As(err,&nfe){
			returnnil
		}

		returncreate.DiagError(names.MediaLive,create.ErrActionDeleting,ResNameChannel,d.Id(),err)
	}

	if_,err:=waitChannelDeleted(ctx,conn,d.Id(),d.Timeout(schema.TimeoutDelete));err!=nil{
		returncreate.DiagError(names.MediaLive,create.ErrActionWaitingForDeletion,ResNameChannel,d.Id(),err)
	}

	returnnil
}

funcstartChannel(ctxcontext.Context,conn*medialive.Client,timeouttime.Duration,idstring)error{
	_,err:=conn.StartChannel(ctx,&medialive.StartChannelInput{
		ChannelId:aws.String(id),
	})

	iferr!=nil{
		returnfmt.Errorf("startingMedialiveChannel(%s):%s",id,err)
	}

	_,err=waitChannelStarted(ctx,conn,id,timeout)

	iferr!=nil{
		returnfmt.Errorf("waitingforMedialiveChannel(%s)start:%s",id,err)
	}

	returnnil
}

funcstopChannel(ctxcontext.Context,conn*medialive.Client,timeouttime.Duration,idstring)error{
	_,err:=conn.StopChannel(ctx,&medialive.StopChannelInput{
		ChannelId:aws.String(id),
	})

	iferr!=nil{
		returnfmt.Errorf("stoppingMedialiveChannel(%s):%s",id,err)
	}

	_,err=waitChannelStopped(ctx,conn,id,timeout)

	iferr!=nil{
		returnfmt.Errorf("waitingforMedialiveChannel(%s)stop:%s",id,err)
	}

	returnnil
}

funcwaitChannelCreated(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
		Pending:enum.Slice(types.ChannelStateCreating),
		Target:enum.Slice(types.ChannelStateIdle),
		Refresh:statusChannel(ctx,conn,id),
		Timeout:timeout,
		NotFoundChecks:20,
		ContinuousTargetOccurence:2,
	}

	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
		returnout,err
	}

	returnnil,err
}

funcwaitChannelUpdated(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
		Pending:enum.Slice(types.ChannelStateUpdating),
		Target:enum.Slice(types.ChannelStateIdle),
		Refresh:statusChannel(ctx,conn,id),
		Timeout:timeout,
		NotFoundChecks:20,
		ContinuousTargetOccurence:2,
	}

	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
		returnout,err
	}

	returnnil,err
}

funcwaitChannelDeleted(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
		Pending:enum.Slice(types.ChannelStateDeleting),
		Target:[]string{},
		Refresh:statusChannel(ctx,conn,id),
		Timeout:timeout,
	}

	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
		returnout,err
	}

	returnnil,err
}

funcwaitChannelStarted(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
		Pending:enum.Slice(types.ChannelStateStarting),
		Target:enum.Slice(types.ChannelStateRunning),
		Refresh:statusChannel(ctx,conn,id),
		Timeout:timeout,
	}

	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
		returnout,err
	}

	returnnil,err
}

funcwaitChannelStopped(ctxcontext.Context,conn*medialive.Client,idstring,timeouttime.Duration)(*medialive.DescribeChannelOutput,error){
	stateConf:=&retry.StateChangeConf{
		Pending:enum.Slice(types.ChannelStateStopping),
		Target:enum.Slice(types.ChannelStateIdle),
		Refresh:statusChannel(ctx,conn,id),
		Timeout:timeout,
	}

	outputRaw,err:=stateConf.WaitForStateContext(ctx)
	ifout,ok:=outputRaw.(*medialive.DescribeChannelOutput);ok{
		returnout,err
	}

	returnnil,err
}

funcstatusChannel(ctxcontext.Context,conn*medialive.Client,idstring)retry.StateRefreshFunc{
	returnfunc()(interface{},string,error){
		out,err:=FindChannelByID(ctx,conn,id)
		iftfresource.NotFound(err){
			returnnil,"",nil
		}

		iferr!=nil{
			returnnil,"",err
		}

		returnout,string(out.State),nil
	}
}

funcFindChannelByID(ctxcontext.Context,conn*medialive.Client,idstring)(*medialive.DescribeChannelOutput,error){
	in:=&medialive.DescribeChannelInput{
		ChannelId:aws.String(id),
	}
	out,err:=conn.DescribeChannel(ctx,in)
	iferr!=nil{
		varnfe*types.NotFoundException
		iferrors.As(err,&nfe){
			returnnil,&retry.NotFoundError{
				LastError:err,
				LastRequest:in,
			}
		}

		returnnil,err
	}

	ifout==nil{
		returnnil,tfresource.NewEmptyResultError(in)
	}

	//ChannelcanstillbefoundwithastateofDELETED.
	//Setresultasnotfoundwhenthestateisdeleted.
	ifout.State==types.ChannelStateDeleted{
		returnnil,&retry.NotFoundError{
			LastResponse:string(types.ChannelStateDeleted),
			LastRequest:in,
		}
	}

	returnout,nil
}

funcexpandChannelInputAttachments(tfList[]interface{})[]types.InputAttachment{
	varattachments[]types.InputAttachment
	for_,v:=rangetfList{
		m,ok:=v.(map[string]interface{})
		if!ok{
			continue
		}

		varatypes.InputAttachment
		ifv,ok:=m["input_attachment_name"].(string);ok{
			a.InputAttachmentName=aws.String(v)
		}
		ifv,ok:=m["input_id"].(string);ok{
			a.InputId=aws.String(v)
		}
		ifv,ok:=m["input_settings"].([]interface{});ok&&len(v)>0{
			a.InputSettings=expandInputAttachmentInputSettings(v)
		}
		ifv,ok:=m["automatic_input_failover_settings"].([]interface{});ok&&len(v)>0{
			a.AutomaticInputFailoverSettings=expandInputAttachmentAutomaticInputFailoverSettings(v)
		}

		attachments=append(attachments,a)
	}

	returnattachments
}

funcexpandInputAttachmentInputSettings(tfList[]interface{})*types.InputSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.InputSettings
	ifv,ok:=m["audio_selector"].([]interface{});ok&&len(v)>0{
		out.AudioSelectors=expandInputAttachmentInputSettingsAudioSelectors(v)
	}
	ifv,ok:=m["caption_selector"].([]interface{});ok&&len(v)>0{
		out.CaptionSelectors=expandInputAttachmentInputSettingsCaptionSelectors(v)
	}
	ifv,ok:=m["deblock_filter"].(string);ok&&v!=""{
		out.DeblockFilter=types.InputDeblockFilter(v)
	}
	ifv,ok:=m["denoise_filter"].(string);ok&&v!=""{
		out.DenoiseFilter=types.InputDenoiseFilter(v)
	}
	ifv,ok:=m["filter_strength"].(int);ok{
		out.FilterStrength=int32(v)
	}
	ifv,ok:=m["input_filter"].(string);ok&&v!=""{
		out.InputFilter=types.InputFilter(v)
	}
	ifv,ok:=m["network_input_settings"].([]interface{});ok&&len(v)>0{
		out.NetworkInputSettings=expandInputAttachmentInputSettingsNetworkInputSettings(v)
	}
	ifv,ok:=m["scte35_pid"].(int);ok{
		out.Scte35Pid=int32(v)
	}
	ifv,ok:=m["smpte2038_data_preference"].(string);ok&&v!=""{
		out.Smpte2038DataPreference=types.Smpte2038DataPreference(v)
	}
	ifv,ok:=m["source_end_behavior"].(string);ok&&v!=""{
		out.SourceEndBehavior=types.InputSourceEndBehavior(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsAudioSelectors(tfList[]interface{})[]types.AudioSelector{
	varas[]types.AudioSelector
	for_,v:=rangetfList{
		m,ok:=v.(map[string]interface{})
		if!ok{
			continue
		}

		varatypes.AudioSelector
		ifv,ok:=m["name"].(string);ok&&v!=""{
			a.Name=aws.String(v)
		}
		ifv,ok:=m["selector_settings"].([]interface{});ok&&len(v)>0{
			a.SelectorSettings=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettings(v)
		}

		as=append(as,a)
	}

	returnas
}

funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettings(tfList[]interface{})*types.AudioSelectorSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AudioSelectorSettings
	ifv,ok:=m["audio_hls_rendition_selection"].([]interface{});ok&&len(v)>0{
		out.AudioHlsRenditionSelection=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioHlsRenditionSelection(v)
	}
	ifv,ok:=m["audio_language_selection"].([]interface{});ok&&len(v)>0{
		out.AudioLanguageSelection=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioLanguageSelection(v)
	}
	ifv,ok:=m["audio_pid_selection"].([]interface{});ok&&len(v)>0{
		out.AudioPidSelection=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioPidSelection(v)
	}
	ifv,ok:=m["audio_track_selection"].([]interface{});ok&&len(v)>0{
		out.AudioTrackSelection=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelection(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioHlsRenditionSelection(tfList[]interface{})*types.AudioHlsRenditionSelection{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AudioHlsRenditionSelection
	ifv,ok:=m["group_id"].(string);ok&&len(v)>0{
		out.GroupId=aws.String(v)
	}
	ifv,ok:=m["name"].(string);ok&&len(v)>0{
		out.Name=aws.String(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioLanguageSelection(tfList[]interface{})*types.AudioLanguageSelection{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AudioLanguageSelection
	ifv,ok:=m["language_code"].(string);ok&&len(v)>0{
		out.LanguageCode=aws.String(v)
	}
	ifv,ok:=m["language_selection_policy"].(string);ok&&len(v)>0{
		out.LanguageSelectionPolicy=types.AudioLanguageSelectionPolicy(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioPidSelection(tfList[]interface{})*types.AudioPidSelection{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AudioPidSelection
	ifv,ok:=m["pid"].(int);ok{
		out.Pid=int32(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelection(tfList[]interface{})*types.AudioTrackSelection{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AudioTrackSelection
	ifv,ok:=m["tracks"].(*schema.Set);ok&&v.Len()>0{
		out.Tracks=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionTracks(v.List())
	}
	ifv,ok:=m["dolby_e_decode"].([]interface{});ok&&len(v)>0{
		out.DolbyEDecode=expandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionDolbyEDecode(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionTracks(tfList[]interface{})[]types.AudioTrack{
	iflen(tfList)==0{
		returnnil
	}

	varout[]types.AudioTrack
	for_,v:=rangetfList{
		m,ok:=v.(map[string]interface{})
		if!ok{
			continue
		}

		varotypes.AudioTrack
		ifv,ok:=m["track"].(int);ok{
			o.Track=int32(v)
		}

		out=append(out,o)
	}

	returnout
}

funcexpandInputAttachmentInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionDolbyEDecode(tfList[]interface{})*types.AudioDolbyEDecode{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AudioDolbyEDecode
	ifv,ok:=m["program_selection"].(string);ok&&v!=""{
		out.ProgramSelection=types.DolbyEProgramSelection(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsCaptionSelectors(tfList[]interface{})[]types.CaptionSelector{
	iflen(tfList)==0{
		returnnil
	}

	varout[]types.CaptionSelector
	for_,v:=rangetfList{
		m,ok:=v.(map[string]interface{})
		if!ok{
			continue
		}

		varotypes.CaptionSelector
		ifv,ok:=m["name"].(string);ok&&v!=""{
			o.Name=aws.String(v)
		}
		ifv,ok:=m["language_code"].(string);ok&&v!=""{
			o.LanguageCode=aws.String(v)
		}
		ifv,ok:=m["selector_settings"].([]interface{});ok&&len(v)>0{
			o.SelectorSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettings(v)
		}

		out=append(out,o)
	}

	returnout
}

funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettings(tfList[]interface{})*types.CaptionSelectorSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.CaptionSelectorSettings
	ifv,ok:=m["ancillary_source_settings"].([]interface{});ok&&len(v)>0{
		out.AncillarySourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsAncillarySourceSettings(v)
	}
	ifv,ok:=m["arib_source_settings"].([]interface{});ok&&len(v)>0{
		out.AribSourceSettings=&types.AribSourceSettings{}//noexportedfields
	}
	ifv,ok:=m["dvb_sub_source_settings"].([]interface{});ok&&len(v)>0{
		out.DvbSubSourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsDvbSubSourceSettings(v)
	}
	ifv,ok:=m["embedded_source_settings"].([]interface{});ok&&len(v)>0{
		out.EmbeddedSourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsEmbeddedSourceSettings(v)
	}
	ifv,ok:=m["scte20_source_settings"].([]interface{});ok&&len(v)>0{
		out.Scte20SourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsScte20SourceSettings(v)
	}
	ifv,ok:=m["scte27_source_settings"].([]interface{});ok&&len(v)>0{
		out.Scte27SourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsScte27SourceSettings(v)
	}
	ifv,ok:=m["teletext_source_settings"].([]interface{});ok&&len(v)>0{
		out.TeletextSourceSettings=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettings(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsAncillarySourceSettings(tfList[]interface{})*types.AncillarySourceSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AncillarySourceSettings
	ifv,ok:=m["source_ancillary_channel_number"].(int);ok{
		out.SourceAncillaryChannelNumber=int32(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsDvbSubSourceSettings(tfList[]interface{})*types.DvbSubSourceSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.DvbSubSourceSettings
	ifv,ok:=m["ocr_language"].(string);ok&&v!=""{
		out.OcrLanguage=types.DvbSubOcrLanguage(v)
	}
	ifv,ok:=m["pid"].(int);ok{
		out.Pid=int32(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsEmbeddedSourceSettings(tfList[]interface{})*types.EmbeddedSourceSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.EmbeddedSourceSettings
	ifv,ok:=m["convert_608_to_708"].(string);ok&&v!=""{
		out.Convert608To708=types.EmbeddedConvert608To708(v)
	}
	ifv,ok:=m["scte20_detection"].(string);ok&&v!=""{
		out.Scte20Detection=types.EmbeddedScte20Detection(v)
	}
	ifv,ok:=m["source_608_channel_number"].(int);ok{
		out.Source608ChannelNumber=int32(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsScte20SourceSettings(tfList[]interface{})*types.Scte20SourceSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.Scte20SourceSettings
	ifv,ok:=m["convert_608_to_708"].(string);ok&&v!=""{
		out.Convert608To708=types.Scte20Convert608To708(v)
	}
	ifv,ok:=m["source_608_channel_number"].(int);ok{
		out.Source608ChannelNumber=int32(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsScte27SourceSettings(tfList[]interface{})*types.Scte27SourceSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.Scte27SourceSettings
	ifv,ok:=m["ocr_language"].(string);ok&&v!=""{
		out.OcrLanguage=types.Scte27OcrLanguage(v)
	}
	ifv,ok:=m["pid"].(int);ok{
		out.Pid=int32(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettings(tfList[]interface{})*types.TeletextSourceSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.TeletextSourceSettings
	ifv,ok:=m["output_rectangle"].([]interface{});ok&&len(v)>0{
		out.OutputRectangle=expandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettingsOutputRectangle(v)
	}
	ifv,ok:=m["page_number"].(string);ok&&v!=""{
		out.PageNumber=aws.String(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettingsOutputRectangle(tfList[]interface{})*types.CaptionRectangle{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.CaptionRectangle
	ifv,ok:=m["height"].(float32);ok{
		out.Height=float64(v)
	}
	ifv,ok:=m["left_offset"].(float32);ok{
		out.LeftOffset=float64(v)
	}
	ifv,ok:=m["top_offset"].(float32);ok{
		out.TopOffset=float64(v)
	}
	ifv,ok:=m["width"].(float32);ok{
		out.Width=float64(v)
	}

	return&out
}

funcexpandInputAttachmentInputSettingsNetworkInputSettings(tfList[]interface{})*types.NetworkInputSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.NetworkInputSettings
	ifv,ok:=m["hls_input_settings"].([]interface{});ok&&len(v)>0{
		out.HlsInputSettings=expandNetworkInputSettingsHLSInputSettings(v)
	}
	ifv,ok:=m["server_validation"].(string);ok&&v!=""{
		out.ServerValidation=types.NetworkInputServerValidation(v)
	}

	return&out
}

funcexpandNetworkInputSettingsHLSInputSettings(tfList[]interface{})*types.HlsInputSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.HlsInputSettings
	ifv,ok:=m["bandwidth"].(int);ok{
		out.Bandwidth=int32(v)
	}
	ifv,ok:=m["buffer_segments"].(int);ok{
		out.BufferSegments=int32(v)
	}
	ifv,ok:=m["retries"].(int);ok{
		out.Retries=int32(v)
	}
	ifv,ok:=m["retry_interval"].(int);ok{
		out.RetryInterval=int32(v)
	}
	ifv,ok:=m["scte35_source"].(string);ok&&v!=""{
		out.Scte35Source=types.HlsScte35SourceType(v)
	}

	return&out
}

funcexpandInputAttachmentAutomaticInputFailoverSettings(tfList[]interface{})*types.AutomaticInputFailoverSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AutomaticInputFailoverSettings
	ifv,ok:=m["secondary_input_id"].(string);ok&&v!=""{
		out.SecondaryInputId=aws.String(v)
	}
	ifv,ok:=m["error_clear_time_msec"].(int);ok{
		out.ErrorClearTimeMsec=int32(v)
	}
	ifv,ok:=m["failover_conditions"].(*schema.Set);ok&&v.Len()>0{
		out.FailoverConditions=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditions(v.List())
	}
	ifv,ok:=m["input_preference"].(string);ok&&v!=""{
		out.InputPreference=types.InputPreference(v)
	}

	return&out
}

funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditions(tfList[]interface{})[]types.FailoverCondition{
	iflen(tfList)==0{
		returnnil
	}

	varout[]types.FailoverCondition
	for_,v:=rangetfList{
		m,ok:=v.(map[string]interface{})
		if!ok{
			continue
		}

		varotypes.FailoverCondition
		ifv,ok:=m["failover_condition_settings"].([]interface{});ok&&len(v)>0{
			o.FailoverConditionSettings=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettings(v)
		}

		out=append(out,o)
	}

	returnout
}

funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettings(tfList[]interface{})*types.FailoverConditionSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.FailoverConditionSettings
	ifv,ok:=m["audio_silence_settings"].([]interface{});ok&&len(v)>0{
		out.AudioSilenceSettings=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsAudioSilenceSettings(v)
	}
	ifv,ok:=m["input_loss_settings"].([]interface{});ok&&len(v)>0{
		out.InputLossSettings=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsInputLossSettings(v)
	}
	ifv,ok:=m["video_black_settings"].([]interface{});ok&&len(v)>0{
		out.VideoBlackSettings=expandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsVideoBlackSettings(v)
	}

	return&out
}

funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsAudioSilenceSettings(tfList[]interface{})*types.AudioSilenceFailoverSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.AudioSilenceFailoverSettings
	ifv,ok:=m["audio_selector_name"].(string);ok&&v!=""{
		out.AudioSelectorName=aws.String(v)
	}
	ifv,ok:=m["audio_silence_threshold_msec"].(int);ok{
		out.AudioSilenceThresholdMsec=int32(v)
	}

	return&out
}

funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsInputLossSettings(tfList[]interface{})*types.InputLossFailoverSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.InputLossFailoverSettings
	ifv,ok:=m["input_loss_threshold_msec"].(int);ok{
		out.InputLossThresholdMsec=int32(v)
	}

	return&out
}

funcexpandInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsVideoBlackSettings(tfList[]interface{})*types.VideoBlackFailoverSettings{
	iftfList==nil{
		returnnil
	}

	m:=tfList[0].(map[string]interface{})

	varouttypes.VideoBlackFailoverSettings
	ifv,ok:=m["black_detect_threshold"].(float32);ok{
		out.BlackDetectThreshold=float64(v)
	}
	ifv,ok:=m["video_black_threshold_msec"].(int);ok{
		out.VideoBlackThresholdMsec=int32(v)
	}

	return&out
}

funcflattenChannelInputAttachments(tfList[]types.InputAttachment)[]interface{}{
	iflen(tfList)==0{
		returnnil
	}

	varout[]interface{}

	for_,item:=rangetfList{
		m:=map[string]interface{}{
			"input_id":ToString(item.InputId),
			"input_attachment_name":aws.ToString(item.InputAttachmentName),
			"input_settings":flattenInputAttachmentsInputSettings(item.InputSettings),
			"automatic_input_failover_settings":flattenInputAttachmentAutomaticInputFailoverSettings(item.AutomaticInputFailoverSettings),
		}

		out=append(out,m)
	}

	returnout
}

funcflattenInputAttachmentsInputSettings(in*types.InputSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"audio_selector":flattenInputAttachmentsInputSettingsAudioSelectors(in.AudioSelectors),
		"caption_selector":flattenInputAttachmentsInputSettingsCaptionSelectors(in.CaptionSelectors),
		"deblock_filter":string(in.DeblockFilter),
		"denoise_filter":string(in.DenoiseFilter),
		"filter_strength":int(in.FilterStrength),
		"input_filter":string(in.InputFilter),
		"network_input_settings":flattenInputAttachmentsInputSettingsNetworkInputSettings(in.NetworkInputSettings),
		"scte35_pid":n.Scte35Pid),
		"smpte2038_data_preference":string(in.Smpte2038DataPreference),
		"source_end_behavior":g(in.SourceEndBehavior),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsAudioSelectors(tfList[]types.AudioSelector)[]interface{}{
	iflen(tfList)==0{
		returnnil
	}

	varout[]interface{}

	for_,v:=rangetfList{
		m:=map[string]interface{}{
			"name":aws.ToString(v.Name),
			"selector_settings":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettings(v.SelectorSettings),
		}

		out=append(out,m)
	}

	returnout
}

funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettings(in*types.AudioSelectorSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"audio_hls_rendition_selection":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioHlsRenditionSelection(in.AudioHlsRenditionSelection),
		"audio_language_selection":nInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioLanguageSelection(in.AudioLanguageSelection),
		"audio_pid_selection":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioPidSelection(in.AudioPidSelection),
		"audio_track_selection":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelection(in.AudioTrackSelection),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioHlsRenditionSelection(in*types.AudioHlsRenditionSelection)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"group_id":aws.ToString(in.GroupId),
		"name":aws.ToString(in.Name),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioLanguageSelection(in*types.AudioLanguageSelection)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"language_code":aws.ToString(in.LanguageCode),
		"language_selection_policy":string(in.LanguageSelectionPolicy),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioPidSelection(in*types.AudioPidSelection)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"pid":int(in.Pid),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelection(in*types.AudioTrackSelection)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"dolby_e_decode":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionDolbyEDecode(in.DolbyEDecode),
		"tracks":flattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionTracks(in.Tracks),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionDolbyEDecode(in*types.AudioDolbyEDecode)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"program_selection":string(in.ProgramSelection),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsAudioSelectorsSelectorSettingsAudioTrackSelectionTracks(tfList[]types.AudioTrack)[]interface{}{
	iflen(tfList)==0{
		returnnil
	}

	varout[]interface{}

	for_,v:=rangetfList{
		m:=map[string]interface{}{
			"track":int(v.Track),
		}

		out=append(out,m)
	}

	returnout
}

funcflattenInputAttachmentsInputSettingsCaptionSelectors(tfList[]types.CaptionSelector)[]interface{}{
	iflen(tfList)==0{
		returnnil
	}

	varout[]interface{}

	for_,v:=rangetfList{
		m:=map[string]interface{}{
			"name":aws.ToString(v.Name),
			"language_code":aws.ToString(v.LanguageCode),
			"selector_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettings(v.SelectorSettings),
		}

		out=append(out,m)
	}

	returnout
}

funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettings(in*types.CaptionSelectorSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"ancillary_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsAncillarySourceSettings(in.AncillarySourceSettings),
		"arib_source_settings":rface{}{},//attributehasnoexportedfields
		"dvb_sub_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsDvbSubSourceSettings(in.DvbSubSourceSettings),
		"embedded_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsEmbeddedSourceSettings(in.EmbeddedSourceSettings),
		"scte20_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsScte20SourceSettings(in.Scte20SourceSettings),
		"scte27_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsScte27SourceSettings(in.Scte27SourceSettings),
		"teletext_source_settings":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettings(in.TeletextSourceSettings),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsAncillarySourceSettings(in*types.AncillarySourceSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"source_ancillary_channel_number":int(in.SourceAncillaryChannelNumber),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsDvbSubSourceSettings(in*types.DvbSubSourceSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"ocr_language":string(in.OcrLanguage),
		"pid":int(in.Pid),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsEmbeddedSourceSettings(in*types.EmbeddedSourceSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"convert_608_to_708":ng(in.Convert608To708),
		"scte20_detection":string(in.Scte20Detection),
		"source_608_channel_number":int(in.Source608ChannelNumber),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsScte20SourceSettings(in*types.Scte20SourceSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"convert_608_to_708":ng(in.Convert608To708),
		"source_608_channel_number":int(in.Source608ChannelNumber),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsScte27SourceSettings(in*types.Scte27SourceSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"ocr_language":string(in.OcrLanguage),
		"pid":int(in.Pid),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettings(in*types.TeletextSourceSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"output_rectangle":flattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettingsOutputRectangle(in.OutputRectangle),
		"page_number":String(in.PageNumber),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsCaptionSelectorsSelectorSettingsTeletextSourceSettingsOutputRectangle(in*types.CaptionRectangle)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"height":2(in.Height),
		"left_offset":float32(in.LeftOffset),
		"top_offset":float32(in.TopOffset),
		"width":32(in.Width),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentsInputSettingsNetworkInputSettings(in*types.NetworkInputSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"hls_input_settings":flattenNetworkInputSettingsHLSInputSettings(in.HlsInputSettings),
		"server_validation":string(in.ServerValidation),
	}

	return[]interface{}{m}
}

funcflattenNetworkInputSettingsHLSInputSettings(in*types.HlsInputSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"bandwidth":n.Bandwidth),
		"buffer_segments":int(in.BufferSegments),
		"retries":int(in.Retries),
		"retry_interval":int(in.RetryInterval),
		"scte35_source":string(in.Scte35Source),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentAutomaticInputFailoverSettings(in*types.AutomaticInputFailoverSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"secondary_input_id":aws.ToString(in.SecondaryInputId),
		"error_clear_time_msec":int(in.ErrorClearTimeMsec),
		"failover_conditions":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditions(in.FailoverConditions),
		"input_preference":(in.InputPreference),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditions(tfList[]types.FailoverCondition)[]interface{}{
	iflen(tfList)==0{
		returnnil
	}

	varout[]interface{}

	for_,item:=rangetfList{
		m:=map[string]interface{}{
			"failover_condition_settings":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettings(item.FailoverConditionSettings),
		}

		out=append(out,m)
	}
	returnout
}

funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettings(in*types.FailoverConditionSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"audio_silence_settings":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsAudioSilenceSettings(in.AudioSilenceSettings),
		"input_loss_settings":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsInputLossSettings(in.InputLossSettings),
		"video_black_settings":flattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsVideoBlackSettings(in.VideoBlackSettings),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsAudioSilenceSettings(in*types.AudioSilenceFailoverSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"audio_selector_name":aws.ToString(in.AudioSelectorName),
		"audio_silence_threshold_msec":int(in.AudioSilenceThresholdMsec),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsInputLossSettings(in*types.InputLossFailoverSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"input_loss_threshold_msec":int(in.InputLossThresholdMsec),
	}

	return[]interface{}{m}
}

funcflattenInputAttachmentAutomaticInputFailoverSettingsFailoverConditionsFailoverConditionSettingsVideoBlackSettings(in*types.VideoBlackFailoverSettings)[]interface{}{
	ifin==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"black_detect_threshold":float32(in.BlackDetectThreshold),
		"video_black_threshold_msec":int(in.VideoBlackThresholdMsec),
	}

	return[]interface{}{m}
}

funcexpandChannelCdiInputSpecification(tfList[]interface{})*types.CdiInputSpecification{
	iftfList==nil{
		returnnil
	}
	m:=tfList[0].(map[string]interface{})

	spec:=&types.CdiInputSpecification{}
	ifv,ok:=m["resolution"].(string);ok&&v!=""{
		spec.Resolution=types.CdiInputResolution(v)
	}

	returnspec
}

funcflattenChannelCdiInputSpecification(apiObject*types.CdiInputSpecification)[]interface{}{
	ifapiObject==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"resolution":string(apiObject.Resolution),
	}

	return[]interface{}{m}
}

funcexpandChannelDestinations(tfList[]interface{})[]types.OutputDestination{
	iftfList==nil{
		returnnil
	}

	vardestinations[]types.OutputDestination
	for_,v:=rangetfList{
		m,ok:=v.(map[string]interface{})
		if!ok{
			continue
		}

		vardtypes.OutputDestination
		ifv,ok:=m["id"].(string);ok{
			d.Id=aws.String(v)
		}
		ifv,ok:=m["media_package_settings"].(*schema.Set);ok&&v.Len()>0{
			d.MediaPackageSettings=expandChannelDestinationsMediaPackageSettings(v.List())
		}
		ifv,ok:=m["multiplex_settings"].([]interface{});ok&&len(v)>0{
			d.MultiplexSettings=expandChannelDestinationsMultiplexSettings(v)
		}
		ifv,ok:=m["settings"].(*schema.Set);ok&&v.Len()>0{
			d.Settings=expandChannelDestinationsSettings(v.List())
		}

		destinations=append(destinations,d)
	}

	returndestinations
}

funcexpandChannelDestinationsMediaPackageSettings(tfList[]interface{})[]types.MediaPackageOutputDestinationSettings{
	iftfList==nil{
		returnnil
	}

	varsettings[]types.MediaPackageOutputDestinationSettings
	for_,v:=rangetfList{
		m,ok:=v.(map[string]interface{})
		if!ok{
			continue
		}

		varstypes.MediaPackageOutputDestinationSettings
		ifv,ok:=m["channel_id"].(string);ok{
			s.ChannelId=aws.String(v)
		}

		settings=append(settings,s)
	}

	returnsettings
}

funcexpandChannelDestinationsMultiplexSettings(tfList[]interface{})*types.MultiplexProgramChannelDestinationSettings{
	iftfList==nil{
		returnnil
	}
	m:=tfList[0].(map[string]interface{})

	settings:=&types.MultiplexProgramChannelDestinationSettings{}
	ifv,ok:=m["multiplex_id"].(string);ok&&v!=""{
		settings.MultiplexId=aws.String(v)
	}
	ifv,ok:=m["program_name"].(string);ok&&v!=""{
		settings.ProgramName=aws.String(v)
	}

	returnsettings
}

funcexpandChannelDestinationsSettings(tfList[]interface{})[]types.OutputDestinationSettings{
	iftfList==nil{
		returnnil
	}

	varsettings[]types.OutputDestinationSettings
	for_,v:=rangetfList{
		m,ok:=v.(map[string]interface{})
		if!ok{
			continue
		}

		varstypes.OutputDestinationSettings
		ifv,ok:=m["password_param"].(string);ok{
			s.PasswordParam=aws.String(v)
		}
		ifv,ok:=m["stream_name"].(string);ok{
			s.StreamName=aws.String(v)
		}
		ifv,ok:=m["url"].(string);ok{
			s.Url=aws.String(v)
		}
		ifv,ok:=m["username"].(string);ok{
			s.Username=aws.String(v)
		}

		settings=append(settings,s)
	}

	returnsettings
}

funcflattenChannelDestinations(apiObject[]types.OutputDestination)[]interface{}{
	ifapiObject==nil{
		returnnil
	}

	vartfList[]interface{}
	for_,v:=rangeapiObject{
		m:=map[string]interface{}{
			"id":aws.ToString(v.Id),
			"media_package_settings":flattenChannelDestinationsMediaPackageSettings(v.MediaPackageSettings),
			"multiplex_settings":flattenChannelDestinationsMultiplexSettings(v.MultiplexSettings),
			"settings":nChannelDestinationsSettings(v.Settings),
		}

		tfList=append(tfList,m)
	}

	returntfList
}

funcflattenChannelDestinationsMediaPackageSettings(apiObject[]types.MediaPackageOutputDestinationSettings)[]interface{}{
	ifapiObject==nil{
		returnnil
	}

	vartfList[]interface{}
	for_,v:=rangeapiObject{
		m:=map[string]interface{}{
			"channel_id":aws.ToString(v.ChannelId),
		}

		tfList=append(tfList,m)
	}

	returntfList
}

funcflattenChannelDestinationsMultiplexSettings(apiObject*types.MultiplexProgramChannelDestinationSettings)[]interface{}{
	ifapiObject==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"multiplex_id":aws.ToString(apiObject.MultiplexId),
		"program_name":aws.ToString(apiObject.ProgramName),
	}

	return[]interface{}{m}
}

funcflattenChannelDestinationsSettings(apiObject[]types.OutputDestinationSettings)[]interface{}{
	ifapiObject==nil{
		returnnil
	}

	vartfList[]interface{}
	for_,v:=rangeapiObject{
		m:=map[string]interface{}{
			"password_param":aws.ToString(v.PasswordParam),
			"stream_name":aws.ToString(v.StreamName),
			"url":aws.ToString(v.Url),
			"username":oString(v.Username),
		}

		tfList=append(tfList,m)
	}

	returntfList
}

funcexpandChannelInputSpecification(tfList[]interface{})*types.InputSpecification{
	iftfList==nil{
		returnnil
	}
	m:=tfList[0].(map[string]interface{})

	spec:=&types.InputSpecification{}
	ifv,ok:=m["codec"].(string);ok&&v!=""{
		spec.Codec=types.InputCodec(v)
	}
	ifv,ok:=m["maximum_bitrate"].(string);ok&&v!=""{
		spec.MaximumBitrate=types.InputMaximumBitrate(v)
	}
	ifv,ok:=m["input_resolution"].(string);ok&&v!=""{
		spec.Resolution=types.InputResolution(v)
	}

	returnspec
}

funcflattenChannelInputSpecification(apiObject*types.InputSpecification)[]interface{}{
	ifapiObject==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"codec":string(apiObject.Codec),
		"maximum_bitrate":string(apiObject.MaximumBitrate),
		"input_resolution":string(apiObject.Resolution),
	}

	return[]interface{}{m}
}

funcexpandChannelMaintenanceCreate(tfList[]interface{})*types.MaintenanceCreateSettings{
	iftfList==nil{
		returnnil
	}
	m:=tfList[0].(map[string]interface{})

	settings:=&types.MaintenanceCreateSettings{}
	ifv,ok:=m["maintenance_day"].(string);ok&&v!=""{
		settings.MaintenanceDay=types.MaintenanceDay(v)
	}
	ifv,ok:=m["maintenance_start_time"].(string);ok&&v!=""{
		settings.MaintenanceStartTime=aws.String(v)
	}

	returnsettings
}

funcexpandChannelMaintenanceUpdate(tfList[]interface{})*types.MaintenanceUpdateSettings{
	iftfList==nil{
		returnnil
	}
	m:=tfList[0].(map[string]interface{})

	settings:=&types.MaintenanceUpdateSettings{}
	ifv,ok:=m["maintenance_day"].(string);ok&&v!=""{
		settings.MaintenanceDay=types.MaintenanceDay(v)
	}
	ifv,ok:=m["maintenance_start_time"].(string);ok&&v!=""{
		settings.MaintenanceStartTime=aws.String(v)
	}
	//NOTE:Thisfieldisonlyavailableintheupdatestruct.Toallowuserstosetascheduled
	//dateonupdate,itmaybeworthaddingtothebaseschema.
	//ifv,ok:=m["maintenance_scheduled_date"].(string);ok&&v!=""{
	//	settings.MaintenanceScheduledDate=aws.String(v)
	//}

	returnsettings
}

funcflattenChannelMaintenance(apiObject*types.MaintenanceStatus)[]interface{}{
	ifapiObject==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"maintenance_day":ng(apiObject.MaintenanceDay),
		"maintenance_start_time":aws.ToString(apiObject.MaintenanceStartTime),
	}

	return[]interface{}{m}
}

funcexpandChannelVPC(tfList[]interface{})*types.VpcOutputSettings{
	iftfList==nil{
		returnnil
	}
	m:=tfList[0].(map[string]interface{})

	settings:=&types.VpcOutputSettings{}
	ifv,ok:=m["security_group_ids"].([]string);ok&&len(v)>0{
		settings.SecurityGroupIds=v
	}
	ifv,ok:=m["subnet_ids"].([]string);ok&&len(v)>0{
		settings.SubnetIds=v
	}
	ifv,ok:=m["public_address_allocation_ids"].([]string);ok&&len(v)>0{
		settings.PublicAddressAllocationIds=v
	}

	returnsettings
}

funcflattenChannelVPC(apiObject*types.VpcOutputSettingsDescription)[]interface{}{
	ifapiObject==nil{
		returnnil
	}

	m:=map[string]interface{}{
		"security_group_ids":flex.FlattenStringValueList(apiObject.SecurityGroupIds),
		"subnet_ids":flex.FlattenStringValueList(apiObject.SubnetIds),
		//public_address_allocation_idsisnotincludedintheoutputstruct
	}

	return[]interface{}{m}
}
