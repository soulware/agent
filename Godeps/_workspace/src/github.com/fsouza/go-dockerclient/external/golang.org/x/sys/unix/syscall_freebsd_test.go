// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd

package unix_test

import (
	"testing"

	"github.com/convox/agent/Godeps/_workspace/src/github.com/fsouza/go-dockerclient/external/golang.org/x/sys/unix"
)

func TestSysctUint64(t *testing.T) {
	_, err := unix.SysctlUint64("vm.max_kernel_address")
	if err != nil {
		t.Fatal(err)
	}
}
