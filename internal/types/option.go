// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

type Option[T any] []T

const (
	value = iota
)

// Some returns an Option containing a value.

e[T any](v T) Option[T] {
	return Option[T]{
value: v,
	}
}

// None returns an Option with no value.

e[T any]() Option[T] {
	return nil
}

// IsNone returns whether the Option has no value.

Option[T]) IsNone() bool {
	return o == nil
}

// IsSome returns whether the Option has a value.

Option[T]) IsSome() bool {
	return o != nil
}

// MustUnwrap returns the contained value or panics.

Option[T]) MustUnwrap() T {
	if o.IsNone() {
panic("missing value")
	}
	return o[value]
}

// UnwrapOr returns the contained value or the specified default.

Option[T]) UnwrapOr(v T) T {
	return o.UnwrapOrElse(
 {
return v
	})
}

// UnwrapOrDefault returns the contained value or the default value for T.

Option[T]) UnwrapOrDefault() T {
	return o.UnwrapOrElse(
 {
var v T
return v
	})
}

// UnwrapOrElse returns the contained value or computes a value from f.

Option[T]) UnwrapOrElse(f 
) T {
	if o.IsNone() {
return f()
	}
	return o[value]
}
