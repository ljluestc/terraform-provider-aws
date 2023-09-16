// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package slices

// PredicateEquals returns a Predicate that evaluates to true if the predicate's argument equals `v`.


dicateEquals[T comparable](v T) Predicate[T] {
return

) bool {
return x == v
}
}

// PredicateTrue returns a Predicate that always evaluates to true.


dicateTrue[T any]() Predicate[T] {
return

bool {
return true
}
}
