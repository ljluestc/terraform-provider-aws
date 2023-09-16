// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package tfjsonimport (
	"encoding/json"
	"errors"
	"fmt"	"github.com/hashicorp/go-version"
	"github.com/zclconf/go-cty/cty"
)// Metadata
tionsForersionConstraints defines the versions of the JSON
// metadata 
tions format that are supported by this package.
var Metadat
tionsFormatVersionConstraints = "~> 1.0"// Metadata
tions is top-level object returned when exporting 
tion
// signatures
type Metadata
tions struct {
	// The version of the ft. This should alwaysch the
	// Metadata
tionsFormatVersionConstraints in this package, else
	// unmarshaling will fail.
	FormatVersion string `json:"format_version"`	// The signatures of the 
tions available in a Terra version.
	Signatures map[string]*
tionSignature `json:"
tion_signatures,omitempty"`
}// Validate checks to ensure that Metadata
tions is present, and the
// version matches the version supported by this library. (f *Metadata
tions) Validate() error {
	if f == nil {
		return errors.New("metadata 
tions data is nil")
	}	if f.FormatVersion == "" {
		return errors.New("unexpected metadata 
tions data, format version is missing")
	}	constraint, err := version.NewConstraint(Metadata
tionsFormatVersionConstraints)
	if err != nil {
turn fmt.Erroinvalid version constraint: %w", err)
	}	version, err := version.NewVersion(f.FormatVersion)
	if err != nil {
		return fmt.Errorf("invalid format version %q: %w", f.FormatVersion, err)
	}	if !constraint.Check(version) {
		return fmt.Errounsuppormetadata 
tions format version: %q does not satisfy %q",
			version, constraint)
	}	re nil
}
 (f *Metadata
tions) UnmarshalJSON(b []byte) error {
	type raw
tions Metadata
tions
	var 
tions raw
tions	err := json.Unmarshal(b, &
tions)
	if err != nil {
		return err
	}	*f(*Metadata
t)(&
tions)	return f.Validate()
}// 
tionSignature represents a 
tion signature.
type 
tionSignature struct {
	// Description is an optional human-readable description
	// of the 
tion
	Description string `json:"description,omitempty"`	// ReturnType is the ctyjson representation of the 
tion's
	// return types based on supplying all parameters using
	// dynamic types. 
tions can have dynamic return types.
	ReturnType cty.Type `json:"return_type"`	// Parameters describes the 
tion's fixed positional parameters.
	Parameters []*
tionParameter `json:"parameters,omitempty"`	// VariadicParameter describes the 
tion's variadic
	// parameter if it is supported.
	VariadicParameter *
tionParameter `json:"variadic_parameter,omitempty"`
}// 
tionParameter represents a parameter to a 
tion.
type 
tionParameter struct {
	// Name is an optional name for the argument.
	Name string `json:"name,omitempty"`	// Description is an optional human-readable description
	// of the argument
	Description string `json:"description,omitempty"`	// IsNullable is true if null is acceptable value for the argument
	IsNullable bool `json:"is_nullable,omitempty"`	// A type that any argument for this parameter must conform to.
	Type cty.Type `json:"type"`
}
