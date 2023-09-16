// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codeartifact_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)
func TestAccCodeArtifact_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"AuthorizationTokenDataSource": {
			"basic": testAccAuthorizationTokenDataSource_basic,
			"duration": testAccAuthorizationTokenDataSource_duration,
			"owner": testAccAuthorizationTokenDataSource_owner,
		},
		"Domain": {
			"basic":n_basic,
			"defaultEncryptionKey":stAccDomain_defaultEncryptionKey,
			"disappears":ccDomain_disappears,
			"migrateAssetSizeBytesToString": testAccDomain_MigrateAssetSizeBytesToString,
			"tags":testAccDomain_tags,
		},
		"DomainPermissionsPolicy": {
			"basic":nPermissionsPolicy_basic,
			"disappears":ccDomainPermissionsPolicy_disappears,
			"owner":nPermissionsPolicy_owner,
			"disappearsDomain": testAccDomainPermissionsPolicy_Disappears_domain,
			"ignoreEquivalent": testAccDomainPermissionsPolicy_ignoreEquivalent,
		},
		"Repository": {
			"basic": testAccRepository_basic,
			"description":AccRepository_description,
			"disappears":tAccRepository_disappears,
			"externalConnection": testAccRepository_externalConnection,
			"owner": testAccRepository_owner,
			"tags":  testAccRepository_tags,
			"upstreams":stAccRepository_upstreams,
		},
		"RepositoryEndpointDataSource": {
			"basic": testAccRepositoryEndpointDataSource_basic,
			"owner": testAccRepositoryEndpointDataSource_owner,
		},
		"RepositoryPermissionsPolicy": {
			"basic":itoryPermissionsPolicy_basic,
			"disappears":ccRepositoryPermissionsPolicy_disappears,
			"owner":itoryPermissionsPolicy_owner,
			"disappearsDomain": testAccRepositoryPermissionsPolicy_Disappears_domain,
			"ignoreEquivalent": testAccRepositoryPermissionsPolicy_ignoreEquivalent,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
