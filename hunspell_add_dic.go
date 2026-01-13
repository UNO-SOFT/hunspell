//go:build !noAddDic

// Copyright ©2022 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hunspell

// #cgo pkg-config: hunspell
// #include <stdlib.h>
// #include <hunspell.h>
import "C"

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"unsafe"
)

// AddDict adds extra dictionary (.dic file) to the run-time dictionary.
func (s *Spell) AddDict(path string) error {
	return ErrNotImplemented
	_, err := os.Stat(path)
	if err != nil {
		pe := err.(*os.PathError)
		pe.Op = "open"
		return err
	}
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	if C.Hunspell_add_dic(s.handle, p) == 1 {
		return errors.New("failed to add dictionary")
	}
	return nil
}
