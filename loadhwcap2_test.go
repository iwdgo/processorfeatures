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
	if err == nil {
		t.Logf("%b", v)
		return
	}
	if err == io.EOF && v == 0 {
		t.Skipf("AT_HWCAP2 not found")
	}
	if err.Error() == "unexpected bit in the high word" {
		t.Logf("Ignoring error. %v", v)
		return
	}
	t.Errorf("%v", err)
}
