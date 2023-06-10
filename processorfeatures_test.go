package processorfeatures

import (
	"os"
	"runtime"
	"testing"
)

func TestIsProcessorFeaturePresent_all(t *testing.T) {
	c := 0
	for _, f := range ProcessorFeatures {
		b, err := IsProcessorFeaturePresent(f.I)
		if err != nil {
			t.Fatal(err)
		}
		if b {
			t.Logf("%s : %s", f.S, f.Doc)
			c++
		}
	}
	t.Logf("%v processor features of %v are available", c, len(ProcessorFeatures))
	c = 0
	for i, f := range MachineFeatures {
		// TODO ARM, machine features do not have sequential numbering
		b, err := IsMachineFeaturePresent(uint32(i))
		if err != nil {
			t.Fatal(err)
		}
		if b {
			t.Logf("%s", f.S)
			c++
		}
		if i >= 32 {
			// TODO Arm, number of bits is larger than known features
			break
		}
	}
	t.Logf("%v machine features of %v are available", c, len(MachineFeatures))
}

func TestSetGOAMD64v(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skipf("not amd64 but %s. Skipping.", runtime.GOARCH)
	}
	v, err := SetGOAMD64v()
	if err != nil {
		t.Fatal(err)
	}
	if s := os.Getenv("GOAMD64"); s != v {
		t.Errorf("got %s, want %s", s, v)
	}
}

func TestSetGOARMv(t *testing.T) {
	if runtime.GOARCH != "arm" && runtime.GOARCH != "arm64" {
		t.Skipf("not arm but %s. Skipping.", runtime.GOARCH)

	}
	v, err := SetGOARMv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
	if runtime.GOARCH == "arm64" {
		if s := os.Getenv("GOARCH"); s != v {
			t.Errorf("got %s, want %s", s, v)
		}
		if s := os.Getenv("GOARM"); s != "" {
			t.Errorf("ARM64: Unexpected GOARM %s", s)
		}
		return
	}
	if s := os.Getenv("GOARM"); s != v {
		t.Errorf("got %s, want %s", s, v)
	}
	if g, w := os.Getenv("GOARCH"), "arm"; g != w {
		t.Errorf("GOARCH: got %s, want %s", g, w)
	}
}

func TestIsVersionComplete(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skipf("not amd64 but %s. Skipping.", runtime.GOARCH)
	}
	v := os.Getenv("GOAMD64")
	t.Logf("GOAMD64 is set to %s", v)
	// windows v1: CX8, MMX, SSE, SSE2 - not checked: CMOV, FPU, FXSR, OSFXSR, SCE
	// windows v2: CMPXCHG16B, SSE3, SSE4_1, SSE4_2, SSSE3 - not checked: LAHF-SAHF, POPCNT,
	// windows v3: AVX, AVX2 - not checked: BMI1, BMI2, F16C, FMA, LZCNT, MOVBE, OSXSAVE
	// windows v4: AVX512F - not checked: AVX512BW, AVX512CD, AVX512DQ, AVX512VL
	for _, check := range []string{"v1", "v2", "v3", "v4"} {
		m := IsVersionComplete(check)
		if len(m) == 0 {
			v = check
			continue
		}
		t.Logf("last complete level is %s", v)
		i := 0 // total number of features of failed level
		for _, f := range ProcessorFeatures {
			if f.V == check {
				i++
			}
		}
		if i == len(m) {
			break
		}
		t.Logf("%s is missing %v features", check, len(m))
		for _, i := range m {
			t.Logf("%s requires %s (%s) which is undetected", ProcessorFeatures[i].V, ProcessorFeatures[i].S, ProcessorFeatures[i].Doc)
		}
	}
}
