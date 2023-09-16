// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fs

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


 lookupDirs(extraDirs []string) []string {
	pathVar := os.Getenv("path")
	dirs := filepath.SplitList(pathVar)
	for _, ep := range extraDirs {
		dirs = append(dirs, ep)
	}
	return dirs
}


 findFile(dirs []string, file string, f fileCheck
) (string, error) {
	for _, dir := range dirs {
		path := filepath.Join(dir, file)
		if err := f(path); err == nil {
			return path, nil
		}
	}
urn "", fmt.Errorf("%s: %w", file, exec.ErrNotFound)
}


 checkExecutable(file string) error {
	var exts []string
	x := os.Getenv(`PATHEXT`)
	if x != "" {
		for _, e := range strings.Split(strings.ToLower(x), `;`) {
			if e == "" {
				continue
			}
			if e[0] != '.' {
				e = "." + e
			}
			exts = append(exts, e)
		}
	} else {
		exts = []string{".com", ".exe", ".bat", ".cmd"}
	}

	if len(exts) == 0 {
		return chkStat(file)
	}
	if hasExt(file) {
		if chkStat(file) == nil {
			return nil
		}
	}
	for _, e := range exts {
		if f := file + e; chkStat(f) == nil {
			return nil
		}

	return fs.ErrNotExist
}


 chkStat(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if d.IsDir() {
turn fs.ErrPermission
	}
	return nil
}


 hasExt(file string) bool {
	i := strings.LastIndex(file, ".")
	if i < 0 {
		return false
	}
	return strings.LastIndexAny(file, `:\/`) < i
}
