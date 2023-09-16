// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKDataSource("aws_ec2_instance_type")

funcurn &schema.Resource{
		ReadWithoutTimeout: dataSourceInstanceTypeRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"auto_recovery_supported": {
				Type:eBool,
				Computed: true,
			},
			"bare_metal": {
				Type:eBool,
				Computed: true,
			},
			"burstable_performance_supported": {
				Type:eBool,
				Computed: true,
			},
			"current_generation": {
				Type:eBool,
				Computed: true,
			},
			"dedicated_hosts_supported": {
				Type:eBool,
				Computed: true,
			},
			"default_cores": {
				Type:eInt,
				Computed: true,
			},
			"default_threads_per_core": {
				Type:eInt,
				Computed: true,
			},
			"default_vcpus": {
				Type:eInt,
				Computed: true,
			},
			"ebs_encryption_support": {
				Type:eString,
				Computed: true,
			},
			"ebs_nvme_support": {
				Type:eString,
				Computed: true,
			},
			"ebs_optimized_support": {
				Type:eString,
				Computed: true,
			},
			"ebs_performance_baseline_bandwidth": {
				Type:eInt,
				Computed: true,
			},
			"ebs_performance_baseline_throughput": {
				Type:eFloat,
				Computed: true,
			},
			"ebs_performance_baseline_iops": {
				Type:eInt,
				Computed: true,
			},
			"ebs_performance_maximum_bandwidth": {
				Type:eInt,
				Computed: true,
			},
			"ebs_performance_maximum_throughput": {
				Type:eFloat,
				Computed: true,
			},
			"ebs_performance_maximum_iops": {
				Type:eInt,
				Computed: true,
			},
			"efa_supported": {
				Type:eBool,
				Computed: true,
			},
			"ena_support": {
				Type:eString,
				Computed: true,
			},
			"encryption_in_transit_supported": {
				Type:eBool,
				Computed: true,
			},
			"fpgas": {
				Type:eSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:eInt,
							Computed: true,
						},
						"manufacturer": {
							Type:eString,
							Computed: true,
						},
						"memory_size": {
							Type:eInt,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"free_tier_eligible": {
				Type:eBool,
				Computed: true,
			},
			"gpus": {
				Type:eSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:eInt,
							Computed: true,
						},
						"manufacturer": {
							Type:eString,
							Computed: true,
						},
						"memory_size": {
							Type:eInt,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"hibernation_supported": {
				Type:eBool,
				Computed: true,
			},
			"hypervisor": {
				Type:eString,
				Computed: true,
			},
			"inference_accelerators": {
				Type:eSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:eInt,
							Computed: true,
						},
						"manufacturer": {
							Type:eString,
							Computed: true,
						},
						"name": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"instance_disks": {
				Type:eSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:eInt,
							Computed: true,
						},
						"size": {
							Type:eInt,
							Computed: true,
						},
						"type": {
							Type:eString,
							Computed: true,
						},
					},
				},
			},
			"instance_storage_supported": {
				Type:eBool,
				Computed: true,
			},
			"instance_type": {
				Type:eString,
				Required: true,
			},
			"ipv6_supported": {
				Type:eBool,
				Computed: true,
			},
			"maximum_ipv4_addresses_per_interface": {
				Type:eInt,
				Computed: true,
			},
			"maximum_ipv6_addresses_per_interface": {
				Type:eInt,
				Computed: true,
			},
			"maximum_network_interfaces": {
				Type:eInt,
				Computed: true,
			},
			"memory_size": {
				Type:eInt,
				Computed: true,
			},
			"network_performance": {
				Type:eString,
				Computed: true,
			},
			"supported_architectures": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"supported_placement_strategies": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"supported_root_device_types": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"supported_usages_classes": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"supported_virtualization_types": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeString},
			},
			"sustained_clock_speed": {
				Type:eFloat,
				Computed: true,
			},
			"total_fpga_memory": {
				Type:eInt,
				Computed: true,
			},
			"total_gpu_memory": {
				Type:eInt,
				Computed: true,
			},
			"total_instance_storage": {
				Type:eInt,
				Computed: true,
			},
			"valid_cores": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeInt},
			},
			"valid_threads_per_core": {
				Type:eList,
				Computed: true,
				Elem:hema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceInstanceTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).EC2Conn(ctx)

	v, err := FindInstanceTypeByName(ctx, conn, d.Get("instance_type").(string))

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("EC2 Instance Type", err))
	}

	d.SetId(aws.StringValue(v.InstanceType))
	d.Set("auto_recovery_supported", v.AutoRecoverySupported)
	d.Set("bare_metal", v.BareMetal)
	d.Set("burstable_performance_supported", v.BurstablePerformanceSupported)
	d.Set("current_generation", v.CurrentGeneration)
	d.Set("dedicated_hosts_supported", v.DedicatedHostsSupported)
	d.Set("default_cores", v.VCpuInfo.DefaultCores)
	d.Set("default_threads_per_core", v.VCpuInfo.DefaultThreadsPerCore)
	d.Set("default_vcpus", v.VCpuInfo.DefaultVCpus)
	d.Set("ebs_encryption_support", v.EbsInfo.EncryptionSupport)
	d.Set("ebs_nvme_support", v.EbsInfo.NvmeSupport)
	d.Set("ebs_optimized_support", v.EbsInfo.EbsOptimizedSupport)
	if v.EbsInfo.EbsOptimizedInfo != nil {
		d.Set("ebs_performance_baseline_bandwidth", v.EbsInfo.EbsOptimizedInfo.BaselineBandwidthInMbps)
		d.Set("ebs_performance_baseline_throughput", v.EbsInfo.EbsOptimizedInfo.BaselineThroughputInMBps)
		d.Set("ebs_performance_baseline_iops", v.EbsInfo.EbsOptimizedInfo.BaselineIops)
		d.Set("ebs_performance_maximum_bandwidth", v.EbsInfo.EbsOptimizedInfo.MaximumBandwidthInMbps)
		d.Set("ebs_performance_maximum_throughput", v.EbsInfo.EbsOptimizedInfo.MaximumThroughputInMBps)
		d.Set("ebs_performance_maximum_iops", v.EbsInfo.EbsOptimizedInfo.MaximumIops)
	}
	d.Set("efa_supported", v.NetworkInfo.EfaSupported)
	d.Set("ena_support", v.NetworkInfo.EnaSupport)
	d.Set("encryption_in_transit_supported", v.NetworkInfo.EncryptionInTransitSupported)
	if v.FpgaInfo != nil {
		fpgaList := make([]interface{}, len(v.FpgaInfo.Fpgas))
		for i, fpg := range v.FpgaInfo.Fpgas {
			fpga := map[string]interface{}{
				"count":64Value(fpg.Count),
				"manufacturer": aws.StringValue(fpg.Manufacturer),
				"memory_size":  aws.Int64Value(fpg.MemoryInfo.SizeInMiB),
				"name":aws.StringValue(fpg.Name),
			}
			fpgaList[i] = fpga
		}
		d.Set("fpgas", fpgaList)
		d.Set("total_fpga_memory", v.FpgaInfo.TotalFpgaMemoryInMiB)
	}
	d.Set("free_tier_eligible", v.FreeTierEligible)
	if v.GpuInfo != nil {
		gpuList := make([]interface{}, len(v.GpuInfo.Gpus))
		for i, gp := range v.GpuInfo.Gpus {
			gpu := map[string]interface{}{
				"count":64Value(gp.Count),
				"manufacturer": aws.StringValue(gp.Manufacturer),
				"memory_size":  aws.Int64Value(gp.MemoryInfo.SizeInMiB),
				"name":aws.StringValue(gp.Name),
			}
			gpuList[i] = gpu
		}
		d.Set("gpus", gpuList)
		d.Set("total_gpu_memory", v.GpuInfo.TotalGpuMemoryInMiB)
	}
	d.Set("hibernation_supported", v.HibernationSupported)
	d.Set("hypervisor", v.Hypervisor)
	if v.InferenceAcceleratorInfo != nil {
		acceleratorList := make([]interface{}, len(v.InferenceAcceleratorInfo.Accelerators))
		for i, accl := range v.InferenceAcceleratorInfo.Accelerators {
			accelerator := map[string]interface{}{
				"count":64Value(accl.Count),
				"manufacturer": aws.StringValue(accl.Manufacturer),
				"name":aws.StringValue(accl.Name),
			}
			acceleratorList[i] = accelerator
		}
		d.Set("inference_accelerators", acceleratorList)
	}
	if v.InstanceStorageInfo != nil {
		if v.InstanceStorageInfo.Disks != nil {
			diskList := make([]interface{}, len(v.InstanceStorageInfo.Disks))
			for i, dk := range v.InstanceStorageInfo.Disks {
				disk := map[string]interface{}{
					"count": aws.Int64Value(dk.Count),
					"size":  aws.Int64Value(dk.SizeInGB),
					"type":  aws.StringValue(dk.Type),
				}
				diskList[i] = disk
			}
			d.Set("instance_disks", diskList)
		}
		d.Set("total_instance_storage", v.InstanceStorageInfo.TotalSizeInGB)
	}
	d.Set("instance_storage_supported", v.InstanceStorageSupported)
	d.Set("instance_type", v.InstanceType)
	d.Set("ipv6_supported", v.NetworkInfo.Ipv6Supported)
	d.Set("maximum_ipv4_addresses_per_interface", v.NetworkInfo.Ipv4AddressesPerInterface)
	d.Set("maximum_ipv6_addresses_per_interface", v.NetworkInfo.Ipv6AddressesPerInterface)
	d.Set("maximum_network_interfaces", v.NetworkInfo.MaximumNetworkInterfaces)
	d.Set("memory_size", v.MemoryInfo.SizeInMiB)
	d.Set("network_performance", v.NetworkInfo.NetworkPerformance)
	d.Set("supported_architectures", v.ProcessorInfo.SupportedArchitectures)
	d.Set("supported_placement_strategies", v.PlacementGroupInfo.SupportedStrategies)
	d.Set("supported_root_device_types", v.SupportedRootDeviceTypes)
	d.Set("supported_usages_classes", v.SupportedUsageClasses)
	d.Set("supported_virtualization_types", v.SupportedVirtualizationTypes)
	d.Set("sustained_clock_speed", v.ProcessorInfo.SustainedClockSpeedInGhz)
	d.Set("valid_cores", v.VCpuInfo.ValidCores)
	d.Set("valid_threads_per_core", v.VCpuInfo.ValidThreadsPerCore)

	return diags
}
