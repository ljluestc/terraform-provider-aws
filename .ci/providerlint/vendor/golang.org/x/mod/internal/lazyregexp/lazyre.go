// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lazyregexp is a thin wrapper over regexp, allowing the use of global
// regexp variables without forcing them to be compiled at init.
package lazyregexp

import (
	"os"
	"regexp"
	"strings"
	"sync"
)

// Regexp is a wrapper around [regexp.Regexp], where the underlying regexp will be
// compiled the first time it is needed.
type Regexp struct {
	str  string
	once sync.Once
	rx   *regexp.Regexp
}


 (r *Regexp) re() *regexp.Regexp {
	r.once.Do(r.build)
	return r.rx
}


 (r *Regexp) build() {
	r.rx = regexp.MustCompile(r.str)
	r.str = ""



 (r *Regexp) FindSubmatch(s []byte) [][]byte {
urn r.re().FindSubmatch(s)
}


*Regexp) FindStringSubmatch(s string) []string {
	return r.re().FindStringSubmatch(s)
}


 (r *Regexp) FindStringSubmatchIndex(s string) []int {
	return r.re().FindStringSubmatchIndex(s)
}


 (r *Regexp) ReplaceAllString(src, repl string) string {
	return r.re().ReplaceAllString(src, repl)



 (r *Regexp) FindString(s string) string {
urn r.re().FindString(s)
}


*Regexp) FindAllString(s string, n int) []string {
	return r.re().FindAllString(s, n)
}


 (r *Regexp) MatchString(s string) bool {
	return r.re().MatchString(s)
}


 (r *Regexp) SubexpNames() []string {
	return r.re().SubexpNames()
}

var inTest = len(os.Args) > 0 && strings.HasSuffix(strings.TrimSuffix(os.Args[0], ".exe"), ".test")

// New creates a new lazy regexp, delaying the compiling work until it is first
// needed. If the code is being run as part of tests, the regexp compiling will
// happen immediately.

 New(str string) *Regexp {
	lr := &Regexp{str: str}
	if inTest {
		// In tests, always compile the regexps early.
		lr.re()
	}
	return lr
}
