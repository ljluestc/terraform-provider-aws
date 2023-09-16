// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package rdsimport (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)// compareActualEngineVersion sets engine version related attributes
//
// `engine_version_actual` is always set to newVersion
//
// `engine_version` is set to newVersion unless:
//- old and pending versions are equal (ie. the update is awaiting a
//maintenance window)
//- old and new versions are not exactly equal, but match after accounting
//for an omitted patch value in the configuration (ie. old="1.3",
//new="1.3.27" will not trigger a set)
funcet("engine_version_actual", newVersion)	if oldVersion != "" && oldVersion == pendingVersion {
return
	}	newVersionSubstr := newVersion
	if len(newVersion) > len(oldVersion) {
newVersionSubstr = string([]byte(newVersion)[0 : len(oldVersion)+1])
	}
	if oldVersion != newVersion && oldVersion+"." == newVersionSubstr {
return
	}	d.Set("engine_version", newVersion)
}
