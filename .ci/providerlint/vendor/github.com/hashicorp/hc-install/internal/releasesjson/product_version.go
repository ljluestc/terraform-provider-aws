// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package releasesjsonimport "github.com/hashicorp/go-version"// ProductVersion is a wrapper around a particular product version like
// "consul 0.5.1". A ProductVersion may have one or more builds.
type ProductVersion struct {
Name        string  `json:"name"`
RawVersion  string  `json:"version"`
Version     *version.Version `json:"-"`
SHASUMS     string  `json:"shasums,omitempty"`
SHASUMSSig  string  `json:"shasums_signature,omitempty"`
SHASUMSSigs []string`json:"shasums_signatures,omitempty"`
Builds      ProductBuilds    `json:"builds"`
}type ProductVersionsMap map[string]*ProductVersiontype ProductVersions []*ProductVersion
 (pv ProductVersions) Len() int {
return len(pv)
}
 (pv ProductVersions) Less(i, j int) bool {
return pv[i].Version.LessThan(pv[j].Version) (pv ProductVersions) Swap(i, j int) {
i], pv[j] = pv[j], pv[i]
}
 (pvm ProductVersionsMap) AsSlice() ProductVersions {
versions := make(ProductVersions, 0)for _, pVersion := range pvm {
versions = append(versions, pVersion)
}return versions
}
