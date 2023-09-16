// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !darwin
// +build !windows,!darwin

package robustio

import (
	"io/ioutil"
	"os"
)


 rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}


 readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)



 removeAll(path string) error {
urn os.RemoveAll(path)
}


 isEphemeralError(err error) bool {
	return false
}
