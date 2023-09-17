// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package eksimport (
"time"
)const (
IdentityProviderConfigTypeOIDC = "oidc"
)const (
ResourcesSecrets = "secrets"
)func Resources_Values() []string {
return []string{
ResourcesSecrets,
}
}const (
propagationTimeout = 2 * time.Minute
)
