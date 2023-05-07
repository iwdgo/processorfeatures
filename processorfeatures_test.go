package processorfeatures

import (
	"os"
	"runtime"
	"testing"
)

func TestIsProcessorFeaturesPresent_all(t *testing.T) {
	c := 0
	for _, f := range ProcessorFeatures {
		b, err := IsProcessorFeaturesPresent(f.i)
		if err != nil {
			t.Fatal(err)
		}
		if b {
			t.Logf("%s: %s", f.s, f.doc)
			c++
		}
	}
	t.Logf("%v features of %v are available", c, len(ProcessorFeatures))
}

func TestSetGOAMD64v(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skip("Not an AMD64 processor. Skipping.")
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
		t.Skip("Not an ARM processor. Skipping.")
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
	// TODO Should display all complete or missing items
	v := os.Getenv("GOAMD64")
	t.Logf("%s", v)
	// windows v1: CX8, MMX, SSE, SSE2 - not checked: CMOV, FPU, FXSR, OSFXSR, SCE
	// windows v2: CMPXCHG16B, SSE3, SSE4_1, SSE4_2, SSSE3 - not checked: LAHF-SAHF, POPCNT,
	// windows v3: AVX, AVX2 - not checked: BMI1, BMI2, F16C, FMA, LZCNT, MOVBE, OSXSAVE
	// windows v4: AVX512F - not checked: AVX512BW, AVX512CD, AVX512DQ, AVX512VL
	for _, check := range []string{"v1", "v2", "v3", "v4"} {
		if m := IsVersionComplete(check); len(m) == 0 {
			v = check
		} else {
			// Maximum level is set, all features must be missing
			t.Logf("last complete level is %s", v)
			i := 0
			for _, f := range ProcessorFeatures {
				if f.v == check {
					i++
				}
			}
			if i != len(m) {
				t.Fatalf("GOAMD64 == %s but %s is incomplete", v, check)
			}
		}
	}
}
