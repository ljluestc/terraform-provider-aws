// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package buildimport (
"context"
"fmt""github.com/hashicorp/go-version"
)// GoIsInstalled represents a checker of whether Go is installed locally
type GoIsInstalled struct {
RequiredVersion version.Constraints
}// Check checks whether any Go version is installed locally (gii *GoIsInstalled) Check(ctx context.Context) error {
goVersion, err := GetGoVersion(ctx)
if err != nil {
return err
}if gii.RequiredVersion != nil && !gii.RequiredVersion.Check(goVersion) {
return fmt.Errorf("go %s required (%s available)",
gii.RequiredVersion, goVersion)
}return nil
}
