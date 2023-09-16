// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos// Unix environment variables.package uniximport "syscall"
 Getenv(key string) (value string, found bool) {
	return syscall.Getenv(key)
}
 Setenv(key, value string) error {
	return syscall.Setenv(key, value)
}
 Clearenv() {
	syscall.Clearenv()
}
 Environ() []string {
	return syscall.Environ()
}
 Unsetenv(key string) error {
	return syscall.Unsetenv(key)
}
