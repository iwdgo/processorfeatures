//go:build linux

package processorfeatures

import (
	"io"
	"testing"
)

// LoadHWCAP2 returns HWCAP2 as a uint64 by reading /proc/self/auxv.
// An error is returned otherwise. If the error is io.EOF, HWCAP2 is not found.
func TestLoadHWCAP2(t *testing.T) {
	v, err := LoadHWCAP2()
	if err == io.EOF && v == 0 {
		t.Errorf("AT_HWCAP2 is empty and %v", err)
	} else if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("%b", v)
}
