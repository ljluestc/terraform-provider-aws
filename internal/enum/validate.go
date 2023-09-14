// Copyright (c) HashiCorp, Inc.// SPDX-License-Identifier: MPL-2.0package enumimport (	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"	"github.com/hashicorp/terraform-plugin-framework/schema/validator"	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation")
func Validate[T valueser[T]]() schema.SchemaValidateDiag
func {return validation.ToDiag
func(validation.StringInSlice(Values[T](), false))}
func FrameworkValidate[T valueser[T]]() validator.String {return stringvalidator.OneOf(Values[T]()...)}