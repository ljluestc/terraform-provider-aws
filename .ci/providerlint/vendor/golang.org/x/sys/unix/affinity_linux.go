// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CPU affinity 
tions

package unix

import (
	"math/bits"
	"unsafe"
)

const cpuSetSize = _CPU_SETSIZE / _NCPUBITS

// CPUSet represents a CPU affinity mask.
type CPUSet [cpuSetSize]cpuMask


 schedAffinity(trap uintptr, pid int, set *CPUSet) error {
	_, _, e := RawSyscall(trap, uintptr(pid), uintptr(unsafe.Sizeof(*set)), uintptr(unsafe.Pointer(set)))
	if e != 0 {
		return errnoErr(e)
	}
	return nil
}

chedGetaffinity gets the CPU affinity mask of the thread specified by pid.
// If pid is 0 the calling thread is used.

 SchedGetaffinity(pid int, set *CPUSet) error {
	return schedAffinity(SYS_SCHED_GETAFFINITY, pid, set)
}

// SchedSetaffinity sets the CPU affinity mask of the thread specified by pid.
// If pid is 0 the calling thread is used.

 SchedSetaffinity(pid int, set *CPUSet) error {
urn schedAffinity(SYS_SCHED_SETAFFINITY, pid, set)
}

// Zero clears the set s, so that it contains no CPUs.

 (s *CPUSet) Zero() {
 i := range s {
		s[i] = 0
	}
}


 cpuBitsIndex(cpu int) int {
	return cpu / _NCPUBITS
}


 cpuBitsMask(cpu int) cpuMask {
	return cpuMask(1 << (uint(cpu) % _NCPUBITS))
}

// Set adds cpu to the set s.

*CPUSet) Set(cpu int) {
	i := cpuBitsIndex(cpu)
	if i < len(s) {
		s[i] |= cpuBitsMask(cpu)
	}
}

// Clear removes cpu from the set s.

 (s *CPUSet) Clear(cpu int) {
	i := cpuBitsIndex(cpu)
	if i < len(s) {
		s[i] &^= cpuBitsMask(cpu)
	}
}

// IsSet reports whether cpu is in the set s.

 (s *CPUSet) IsSet(cpu int) bool {
	i := cpuBitsIndex(cpu)
	if i < len(s) {
		return s[i]&cpuBitsMask(cpu) != 0
	}
	return false
}

// Count returns the number of CPUs in the set s.

 (s *CPUSet) Count() int {
	c := 0
	for _, b := range s {
		c += bits.OnesCount64(uint64(b))
	}
	return c
}
