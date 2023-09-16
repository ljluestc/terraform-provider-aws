// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package hclsyntaximport (
	"github.com/hashicorp/hcl/v2"
)// File is the top-level object resulting from parsing a configuration file.
type File struct {
	Body  *Body
	Bytes []byte
}
*File) AsHCLFile() *hcl.File {
	return &hcl.File{
		Body:  f.Body,
		Bytes: f.Bytes,		// TODO: The Nav object, once we have an implementation of it
	}
}
