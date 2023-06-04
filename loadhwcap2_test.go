//go:build linux

package processorfeatures

import (
	"testing"
)

// LoadHWCAP2 returns HWCAP2 as a uint64 by reading /proc/self/auxv.
// An error is returned otherwise. If the error is io.EOF, HWCAP2 is not found.
func TestLoadHWCAP2(t *testing.T) {
	v, err := LoadHWCAP2()
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	if err == io.EOF && (v == nil || v == 0) {
		t.Errorf("%v returned without HWCAP2 %v", err, v)
	}
	t.Logf("%b", v)
}
