//go:build linux

package processorfeatures

import (
	"bytes"
	"errors"
	"os"
)

var ProcessorFeatures = []ProcessorFeature{
	{0, "xv1", "fxsr", ""},
	{1, "v1", "fpu", ""},
	{2, "", "vme", ""},
	{3, "", "de", ""},
	{4, "", "pse", ""},
	{5, "", "tsc", ""},
	{6, "", "msr", ""},
	{7, "", "pae", ""},
	{8, "", "mce", ""},
	{9, "v1", "cx8", ""},
	{10, "", "apic", ""},
	{11, "", "sep", ""},
	{12, "", "mtrr", ""},
	{13, "", "pge", ""},
	{14, "", "mca", ""},
	{15, "v1", "cmov", ""},
	{16, "", "pat", ""},
	{17, "", "pse36", ""},
	{18, "", "clflush", ""},
	{19, "v1", "mmx", ""},
	{20, "v1", "fxsr", ""},
	{21, "v1", "sse", ""},
	{22, "v1", "sse2", ""},
	{23, "", "ss", ""},
	{24, "", "ht", ""},
	{25, "v1", "syscall", ""},
	{26, "", "nx", ""},
	{27, "", "pdpe1gb", ""},
	{28, "", "rdtscp", ""},
	{29, "", "lm", ""},
	{30, "", "constant_tsc", ""},
	{31, "", "rep_good", ""},
	{32, "", "nopl", ""},
	{33, "", "xtopology", ""},
	{34, "", "cpuid", ""},
	{35, "", "pni", ""},
	{36, "", "pclmulqdq", ""},
	{37, "v2", "ssse3", ""},
	{38, "v3", "fma", ""},
	{39, "", "cx16", ""},
	{40, "", "pcid", ""},
	{41, "v2", "sse4_1", ""},
	{42, "v2", "sse4_2", ""},
	{43, "v3", "movbe", ""},
	{44, "v2", "popcnt", ""},
	{45, "", "aes", ""},
	{46, "", "xsave", ""},
	{47, "v3", "avx", ""},
	{48, "v3", "f16c", ""},
	{49, "", "rdrand", ""},
	{50, "", "hypervisor", ""},
	{51, "v2", "lahf_lm", ""},
	{52, "", "abm", ""},
	{53, "", "3dnowprefetch", ""},
	{54, "", "invpcid_single", ""},
	{55, "", "pti", ""},
	{56, "", "fsgsbase", ""},
	{57, "v3", "bmi1", ""},
	{58, "", "hle", ""},
	{59, "v3", "avx2", ""},
	{60, "", "smep", ""},
	{61, "v3", "bmi2", ""},
	{62, "", "erms", ""},
	{63, "", "invpcid", ""},
	{64, "", "rtm", ""},
	{65, "v4", "avx512f", ""},
	{66, "v4", "avx512dq", ""},
	{67, "", "rdseed", ""},
	{68, "", "adx", ""},
	{69, "", "smap", ""},
	{70, "", "clflushopt", ""},
	{71, "v4", "avx512cd", ""},
	{72, "v4", "avx512bw", ""},
	{73, "v4", "avx512vl", ""},
	{74, "", "xsaveopt", ""},
	{75, "", "xsavec", ""},
	{76, "", "xsaves", ""},
	{77, "", "md_clear", ""},
	{78, "xv1", "osfxsr", "v1 - not reported on ubuntu"},
	{79, "xv2", "cmpxchg16b", "v2 - not reported on ubuntu"},
	{80, "xv2", "sse3", "v2 - not reported on ubuntu"},
	{81, "xv3", "lzcnt", "v3 - not reported on ubuntu"},
	{82, "xv3", "osxsave", "v3 - not reported on ubuntu"},
}

var (
	filled         = false
	featuresstatus = make([]bool, len(ProcessorFeatures))
)

func loadflags() error {
	// TODO Check flags for each processor
	flags, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return err
	}
	b := false
	_, flags, b = bytes.Cut(flags, []byte("flags\t\t:"))
	if !b {
		return errors.New("no flags section found")
	}
	flags, _, b = bytes.Cut(flags, []byte("\n"))
	if !b {
		return errors.New("flags are not truncated")
	}
	flags = bytes.Trim(flags, " ")
	features := bytes.Split(flags, []byte(" "))
	i, j := 0, 0
	var f string
	for {
		if i == len(features) {
			break
		}
		if j == len(features) {
			j = i
		}
		f = string(features[i])
		// flag will often be the next one
		if ProcessorFeatures[j].s == f {
			featuresstatus[j] = true
			i++
			j++
			continue
		}
		// search entry
		for k, pf := range ProcessorFeatures {
			if f == pf.s {
				// flag is found but not at the same place, i.e. order has changed
				featuresstatus[k] = true
				j = k + 1
			}
		}
		i++
	}
	return nil
}

// IsProcessorFeaturesPresent returns true when the feature identified by the number is present.
func IsProcessorFeaturesPresent(i uint32) (bool, error) {
	if !filled {
		if err := loadflags(); err != nil {
			return false, err
		}
		filled = true
	}
	return featuresstatus[i], nil
}

// SetGOAMD64v sets GOAMD64 for the current go session. Default v1 is used when no higher version is complete.
// Version set is returned or an error.
func SetGOAMD64v() (v string, err error) {
	for _, v = range []string{"v4", "v3", "v2"} {
		if m := IsVersionComplete(v); len(m) == 0 {
			err = os.Setenv("GOAMD64", v)
			if err != nil {
				return
			}
			return
		}
	}
	v = "v1"
	err = os.Setenv("GOAMD64", v)
	if err != nil {
		return
	}
	return
}

// SetGOARMv sets GOARM and GOARM64 and returns ARM64, 7 or 6 depending on identified processor features.
func SetGOARMv() (v string, err error) {
	v = "ARM64"
	if m := IsVersionComplete(v); len(m) == 0 {
		err = os.Unsetenv("GOARM")
		if err != nil {
			return
		}
		err = os.Setenv("GOARCH", "ARM64")
		if err != nil {
			return
		}
		return
	}
	err = os.Setenv("GOARCH", "ARM")
	if err != nil {
		return
	}
	v = "7"
	if m := IsVersionComplete(v); len(m) == 0 {
		err = os.Setenv("GOARM", v)
		if err != nil {
			return
		}
		return
	}
	// TODO Check that 6 is set for all
	v = "6" // Default
	err = os.Setenv("GOARM", v)
	if err != nil {
		return
	}
	return
}

func IsVersionComplete(v string) (missing []uint32) {
	for _, f := range ProcessorFeatures {
		if f.v == v {
			if b, err := IsProcessorFeaturesPresent(f.i); err == nil {
				if !b {
					missing = append(missing, f.i)
				}
			} else {
				missing = append(missing, f.i)
			}
		}
	}
	return missing
}
