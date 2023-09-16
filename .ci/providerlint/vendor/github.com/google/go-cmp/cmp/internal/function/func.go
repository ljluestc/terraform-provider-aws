// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package 
 provides 
ality for identifying 
 types.
package 


import (
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

type 
 int

const (
	_ 
 = iota

	tb
 
bool
	ttb

T) bool
	trb

R) bool
	tib

I) bool
	tr
 
R

	Equal             = ttb

T) bool
	EqualAssignable   = tib

I) bool; encapsulates 
T) bool
	Transformer       = tr
 
R
	ValueFilter       = ttb

T) bool
	Less              = ttb

T) bool
	ValuePredicate    = tb
 
bool
	KeyValuePredicate = trb

R) bool
)

var boolType = reflect.TypeOf(true)

// IsType reports whether the reflect.Type is of the specified 
 type.

ype(t reflect.Type, ft 
) bool {
	if t == nil || t.Kind() != reflect.
t.IsVariadic() {
		return false
	}
	ni, no := t.NumIn(), t.NumOut()
	switch ft {
	case tb
 
bool
		if ni == 1 && no == 1 && t.Out(0) == boolType {
			return true
		}
	case ttb
 
T) bool
		if ni == 2 && no == 1 && t.In(0) == t.In(1) && t.Out(0) == boolType {
			return true
		}
	case trb
 
R) bool
		if ni == 2 && no == 1 && t.Out(0) == boolType {
			return true
		}
	case tib
 
I) bool
		if ni == 2 && no == 1 && t.In(0).AssignableTo(t.In(1)) && t.Out(0) == boolType {
			return true
		}
	case tr
 
R
		if ni == 1 && no == 1 {
			return true
		}
	}
	return false
}

var lastIdentRx = regexp.MustCompile(`[_\p{L}][_\p{L}\p{N}]*$`)

// NameOf returns the name of the 
 value.

eOf(v reflect.Value) string {
	fnc := runtime.
C(v.Pointer())
	if fnc == nil {
		return "<unknown>"
	}
	fullName := fnc.Name() // e.g., "long/path/name/mypkg.(*MyType).(long/path/name/mypkg.myMethod)-fm"

	// Method closures have a "-fm" suffix.
	fullName = strings.TrimSuffix(fullName, "-fm")

	var name string
	for len(fullName) > 0 {
		inParen := strings.HasSuffix(fullName, ")")
		fullName = strings.TrimSuffix(fullName, ")")

		s := lastIdentRx.FindString(fullName)
		if s == "" {
			break
		}
		name = s + "." + name
		fullName = strings.TrimSuffix(fullName, s)

		if i := strings.LastIndexByte(fullName, '('); inParen && i >= 0 {
			fullName = fullName[:i]
		}
		fullName = strings.TrimSuffix(fullName, ".")
	}
	return strings.TrimSuffix(name, ".")
}
