//go:build linux

package processorfeatures

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var ProcessorFeatures = []ProcessorFeature{
	//cpu_feature_t of https://manpages.ubuntu.com/manpages/kinetic/man3/libcpuid.3.html
	{0, "v1", "fpu", "CPU_FEATURE_FPU"},
	{1, "", "vme", "CPU_FEATURE_VME"},
	{2, "", "de", "CPU_FEATURE_DE"},
	{3, "", "pse", "CPU_FEATURE_PSE"},
	{4, "", "tsc", "CPU_FEATURE_TSC"},
	{5, "", "msr", "CPU_FEATURE_MSR"},
	{6, "", "pae", "CPU_FEATURE_PAE"},
	{7, "", "mce", "CPU_FEATURE_MCE"},
	{8, "v1", "cx8", "CPU_FEATURE_CX8"},
	{9, "", "apic", "CPU_FEATURE_APIC"},
	{10, "", "mtrr", "CPU_FEATURE_MTRR"},
	{11, "", "sep", "CPU_FEATURE_SEP"},
	{12, "", "pge", "CPU_FEATURE_PGE"},
	{13, "", "mca", "CPU_FEATURE_MCA"},
	{14, "v1", "cmov", "CPU_FEATURE_CMOV"},
	{15, "", "pat", "CPU_FEATURE_PAT"},
	{16, "", "pse36", "CPU_FEATURE_PSE36"},
	{17, "", "pn", "CPU_FEATURE_PN"},
	{18, "", "clflush", "CPU_FEATURE_CLFLUSH"},
	{19, "", "dts", "CPU_FEATURE_DTS"},
	{20, "", "acpi", "CPU_FEATURE_ACPI"},
	{21, "v1", "mmx", "CPU_FEATURE_MMX"},
	{22, "v1", "fxsr", "CPU_FEATURE_FXSR"},
	{23, "v1", "sse", "CPU_FEATURE_SSE"},
	{24, "v1", "sse2", "CPU_FEATURE_SSE2"},
	{25, "", "ss", "CPU_FEATURE_SS"},
	{26, "", "ht", "CPU_FEATURE_HT"},
	{27, "", "tm", "CPU_FEATURE_TM"},
	{28, "", "ia64", "CPU_FEATURE_IA64"},
	{29, "", "pbe", "CPU_FEATURE_PBE"},
	{30, "", "pni", "CPU_FEATURE_PNI"},
	{31, "", "pclmul", "CPU_FEATURE_PCLMUL"},
	{32, "", "dts64", "CPU_FEATURE_DTS64"},
	{33, "", "monitor", "CPU_FEATURE_MONITOR"},
	{34, "", "ds_cpl", "CPU_FEATURE_DS_CPL"},
	{35, "", "vmx", "CPU_FEATURE_VMX"},
	{36, "", "smx", "CPU_FEATURE_SMX"},
	{37, "", "est", "CPU_FEATURE_EST"},
	{38, "", "tm2", "CPU_FEATURE_TM2"},
	{39, "v2", "ssse3", "CPU_FEATURE_SSSE3"},
	{40, "", "cid", "CPU_FEATURE_CID"},
	{41, "", "cx16", "CPU_FEATURE_CX16"},
	{42, "", "xtpr", "CPU_FEATURE_XTPR"},
	{43, "", "pdcm", "CPU_FEATURE_PDCM"},
	{44, "", "dca", "CPU_FEATURE_DCA"},
	{45, "v2", "sse4_1", "CPU_FEATURE_SSE4_1"},
	{46, "v2", "sse4_2", "CPU_FEATURE_SSE4_2"},
	{47, "v1", "syscall", "CPU_FEATURE_SYSCALL aka SCE"},
	{48, "", "xd", "CPU_FEATURE_XD"},
	{49, "v3", "movbe", "CPU_FEATURE_MOVBE"},
	{50, "v2", "popcnt", "CPU_FEATURE_POPCNT"},
	{51, "", "aes", "CPU_FEATURE_AES"},
	{52, "", "xsave", "CPU_FEATURE_XSAVE"},
	// Ubuntu seems to only report depending instruction like xsave (see below)
	{53, "xv3", "osxsave", "non-privileged copy of OSXSAVE supported"},
	{54, "v3", "avx", "CPU_FEATURE_AVX"},
	{55, "", "mmxext", "CPU_FEATURE_MMXEXT"},
	{56, "", "3dnow", "CPU_FEATURE_3DNOW"},
	{57, "", "3dnowext", "CPU_FEATURE_3DNOWEXT"},
	{58, "", "nx", "CPU_FEATURE_NX"},
	{59, "", "fxsr_opt", "CPU_FEATURE_FXSR_OPT"},
	{60, "", "rdtscp", "CPU_FEATURE_RDTSCP"},
	{61, "", "lm", "CPU_FEATURE_LM"},
	{62, "v2", "lahf_lm", "CPU_FEATURE_LAHF_LM"},
	{63, "", "cmp_legacy", "CPU_FEATURE_CMP_LEGACY"},
	{64, "", "svm", "CPU_FEATURE_SVM"},
	{65, "", "abm", "CPU_FEATURE_ABM"},
	{66, "", "misalignsse", "CPU_FEATURE_MISALIGNSSE"},
	{67, "", "sse4a", "CPU_FEATURE_SSE4A"},
	{68, "", "3dnowprefetch", "CPU_FEATURE_3DNOWPREFETCH"},
	{69, "", "osvw", "CPU_FEATURE_OSVW"},
	{70, "", "ibs", "CPU_FEATURE_IBS"},
	{71, "", "sse5", "CPU_FEATURE_SSE5"},
	{72, "", "skinit", "CPU_FEATURE_SKINIT"},
	{73, "", "wdt", "CPU_FEATURE_WDT"},
	{74, "", "ts", "CPU_FEATURE_TS"},
	{75, "", "fid", "CPU_FEATURE_FID"},
	{76, "", "vid", "CPU_FEATURE_VID"},
	{77, "", "ttp", "CPU_FEATURE_TTP"},
	{78, "", "tm_amd", "CPU_FEATURE_TM_AMD"},
	{79, "", "stc", "CPU_FEATURE_STC"},
	{80, "", "100mhzsteps", "CPU_FEATURE_100MHZSTEPS"},
	{81, "", "hwpstate", "CPU_FEATURE_HWPSTATE"},
	{82, "", "constant_tsc", "CPU_FEATURE_CONSTANT_TSC"},
	{83, "", "xop", "CPU_FEATURE_XOP"},
	{84, "", "fma3", "CPU_FEATURE_FMA3"},
	{85, "", "fma4", "CPU_FEATURE_FMA4"},
	{86, "", "tbm", "CPU_FEATURE_TBM"},
	{87, "v3", "f16c", "CPU_FEATURE_F16C"},
	{88, "", "rdrand", "CPU_FEATURE_RDRAND"},
	{89, "", "x2apic", "CPU_FEATURE_X2APIC"},
	{90, "", "cpb", "CPU_FEATURE_CPB"},
	{91, "", "aperfmperf", "CPU_FEATURE_APERFMPERF"},
	{92, "", "pfi", "CPU_FEATURE_PFI"},
	{93, "", "pa", "CPU_FEATURE_PA"},
	{94, "v3", "avx2", "CPU_FEATURE_AVX2"},
	{95, "v3", "bmi1", "CPU_FEATURE_BMI1"},
	{96, "v3", "bmi2", "CPU_FEATURE_BMI2"},
	{97, "", "hle", "CPU_FEATURE_HLE"},
	{98, "", "rtm", "CPU_FEATURE_RTM"},
	{99, "v4", "avx512f", "CPU_FEATURE_AVX512F"},
	{100, "v4", "avx512dq", "CPU_FEATURE_AVX512DQ"},
	{101, "", "avx512pf", "CPU_FEATURE_AVX512PF"},
	{102, "", "avx512er", "CPU_FEATURE_AVX512ER"},
	{103, "v4", "avx512cd", "CPU_FEATURE_AVX512CD"},
	{104, "", "sha_ni", "CPU_FEATURE_SHA_NI"},
	{105, "v4", "avx512bw", "CPU_FEATURE_AVX512BW"},
	{106, "v4", "avx512vl", "CPU_FEATURE_AVX512VL"},
	{107, "", "sgx", "CPU_FEATURE_SGX"},
	{108, "", "rdseed", "CPU_FEATURE_RDSEED"},
	{109, "", "adx", "CPU_FEATURE_ADX"},
	{110, "", "avx512vnni", "CPU_FEATURE_AVX512VNNI"},
	{111, "", "avx512vbmi", "CPU_FEATURE_AVX512VBMI"},
	{112, "", "avx512vbmi2", "CPU_FEATURE_AVX512VBMI2"},
	// https://en.wikipedia.org/wiki/X86-64
	// x version prefix means that feature as documented does not seem reported
	{113, "xv1", "fxsave", "CPU_FEATURE_OSFXSR"},
	{114, "xv2", "cmpxchg16b", "CPU_FEATURE_CMPXCHG16B"},
	{115, "xv2", "sse3", "CPU_FEATURE_SSE3"},
	{116, "xv3", "lzcnt", "CPU_FEATURE_LZCNT"},
	// Ubuntu on Github CI also reports
	{117, "", "pdpe1gb", ""},
	{118, "", "rep_good", ""},
	{119, "", "nopl", ""},
	{120, "", "xtopology", ""},
	{121, "", "cpuid", ""},
	{122, "", "pclmulqdq", ""},
	{123, "v3", "fma", ""},
	{124, "", "pcid", ""},
	{125, "", "hypervisor", ""},
	{126, "", "invpcid_single", ""},
	{127, "", "pti", ""},
	{128, "", "fsgsbase", ""},
	{129, "", "smep", ""},
	{130, "", "erms", ""},
	{131, "", "invpcid", ""},
	{132, "", "smap", ""},
	{133, "", "clflushopt", ""},
	{134, "", "xsaveopt", ""},
	{135, "", "xsavec", ""},
	{136, "", "xsaves", ""},
	{137, "", "md_clear", ""},
}

var (
	filled         = false
	featuresstatus = make([]bool, len(ProcessorFeatures))
	machinestatus  = make([]bool, len(MachineFeatures))
)

// loadflags loads processor flags as returned by a command like `cat proc/cpuinfo`
func loadflags() error {
	// TODO Check flags for each processor
	flags, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		if err != io.EOF || flags == nil {
			return err
		}
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
	var found bool
	for {
		if i == len(features) {
			break
		}
		if j == len(features) {
			j = i
		}
		f = string(features[i])
		// flag will often be the next one
		if ProcessorFeatures[j].S == f {
			featuresstatus[j] = true
			i++
			j++
			continue
		}
		// search entry
		found = false
		for k, pf := range ProcessorFeatures {
			if f == pf.S {
				// flag is found but not at the same place, i.e. order has changed
				featuresstatus[k] = true
				j = k + 1
				found = true
				break
			}
		}
		if !found {
			log.Printf("feature %s is unknown", f)
		}
		i++
	}
	v, err := LoadHWCAP2()
	if err != nil {
		return err
	}
	s32 := fmt.Sprintf("%032b", v)
	log.Printf("%s", s32)
	i = 0
	for i < len(machinestatus) {
		machinestatus[i] = string(s32[i]) == "0"
		i++
	}
	return nil
}

func isProcessorFeaturePresent(i uint32) (bool, error) {
	if !filled {
		if err := loadflags(); err != nil {
			return false, err
		}
		filled = true
	}
	return featuresstatus[i], nil
}

func isMachineFeaturePresent(i uint32) (bool, error) {
	if !filled {
		if err := loadflags(); err != nil {
			return false, err
		}
		filled = true
	}
	return machinestatus[i], nil
}
