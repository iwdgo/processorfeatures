//go:build windows

package processorfeatures

import (
	"os"
	"syscall"
)

var ProcessorFeatures = []ProcessorFeature{
	{25, "7", "PF_ARM_64BIT_LOADSTORE_ATOMIC", "The 64-bit load/store atomic instructions are available."},
	{24, "7", "PF_ARM_DIVIDE_INSTRUCTION_AVAILABLE", "The divide instructions are available."},
	{26, "7", "PF_ARM_EXTERNAL_CACHE_AVAILABLE", "The external cache is available."},
	{27, "7", "PF_ARM_FMAC_INSTRUCTIONS_AVAILABLE", "The floating-point multiply-accumulate instruction is available."},
	{18, "7", "PF_ARM_VFP_32_REGISTERS_AVAILABLE", "The VFP/Neon: 32 x 64bit register bank is present. This flag has the same meaning as PF_ARM_VFP_EXTENDED_REGISTERS."},
	// AMD https://en.wikipedia.org/wiki/3DNow! Deprecated
	{7, "", "PF_3DNOW_INSTRUCTIONS_AVAILABLE", "The 3D-Now instruction set is available."},
	{16, "", "PF_CHANNELS_ENABLED", "The processor channels are enabled."},
	{2, "v1", "PF_COMPARE_EXCHANGE_DOUBLE", "The atomic compare and exchange operation (cmpxchg) is available."},
	{14, "v2", "PF_COMPARE_EXCHANGE128", "The atomic compare and exchange 128-bit operation (cmpxchg16b) is available."},
	// Itanium
	{15, "", "PF_COMPARE64_EXCHANGE128", "The atomic compare 64 and exchange 128-bit operation (cmp8xchg16) is available."},
	// AMD & ARM https://learn.microsoft.com/en-us/cpp/intrinsics/fastfail?view=msvc-170
	{23, "", "PF_FASTFAIL_AVAILABLE", "_fastfail() is available."},
	// Floating point is emulated
	{1, "", "PF_FLOATING_POINT_EMULATED", "Floating-point operations are emulated using a software emulator.\nThis function returns a nonzero value if floating-point operations are emulated; otherwise, it returns zero."},
	// https://en.wikipedia.org/wiki/Pentium_FDIV_bug
	{0, "", "PF_FLOATING_POINT_PRECISION_ERRATA", "On a Pentium, a floating-point precision error can occur in rare circumstances."},
	{3, "v1", "PF_MMX_INSTRUCTIONS_AVAILABLE", "The MMX instruction set is available."},
	// https://learn.microsoft.com/en-us/windows/win32/win7appqual/dep-nx-protection
	{12, "", "PF_NX_ENABLED", "Data execution prevention is enabled."},
	// Windows 32-bit: https://learn.microsoft.com/en-us/windows/win32/memory/physical-address-extension
	{9, "", "PF_PAE_ENABLED", "The processor is PAE-enabled."},
	// https://en.wikipedia.org/wiki/Time_Stamp_Counter
	{8, "CR4", "PF_RDTSC_INSTRUCTION_AVAILABLE", "The RDTSC instruction is available."},
	// CR4 https://en.wikipedia.org/wiki/Control_register
	{22, "CR4", "PF_RDWRFSGSBASE_AVAILABLE", "RDFSBASE, RDGSBASE, WRFSBASE, and WRGSBASE instructions are available."},
	{20, "", "PF_SECOND_LEVEL_ADDRESS_TRANSLATION", "Second Level Address Translation is supported by the hardware."},
	{13, "v2", "PF_SSE3_INSTRUCTIONS_AVAILABLE", "The SSE3 instruction set is available."},
	{36, "v2", "PF_SSSE3_INSTRUCTIONS_AVAILABLE", "The SSSE3 instruction set is available."},
	{37, "v2", "PF_SSE4_1_INSTRUCTIONS_AVAILABLE", "The SSE4_1 instruction set is available."},
	{38, "v2", "PF_SSE4_2_INSTRUCTIONS_AVAILABLE", "The SSE4_2 instruction set is available."},
	{39, "v3", "PF_AVX_INSTRUCTIONS_AVAILABLE", "The AVX instruction set is available."},
	{40, "v3", "PF_AVX2_INSTRUCTIONS_AVAILABLE", "The AVX2 instruction set is available."},
	{41, "v4", "PF_AVX512F_INSTRUCTIONS_AVAILABLE", "The AVX512F instruction set is available."},
	// https://support.microsoft.com/en-us/windows/enable-virtualization-on-windows-11-pcs-c5578302-6e43-4b4b-a449-8ced115f58e1
	{21, "", "PF_VIRT_FIRMWARE_ENABLED", "Virtualization is enabled in the firmware and made available by the operating system."},
	{6, "v1", "PF_XMMI_INSTRUCTIONS_AVAILABLE", "The SSE instruction set is available."},
	{10, "v1", "PF_XMMI64_INSTRUCTIONS_AVAILABLE", "The SSE2 instruction set is available."},
	{17, "v2", "PF_XSAVE_ENABLED", "The processor implements the XSAVE and XRSTOR instructions."},
	{29, "ARM64", "PF_ARM_V8_INSTRUCTIONS_AVAILABLE", "This Arm processor implements the Arm v8 instructions set."},
	{30, "ARM64", "PF_ARM_V8_CRYPTO_INSTRUCTIONS_AVAILABLE", "This Arm processor implements the Arm v8 extra cryptographic instructions (for example, AES, SHA1 and SHA2)."},
	{31, "ARM64", "PF_ARM_V8_CRC32_INSTRUCTIONS_AVAILABLE", "This Arm processor implements the Arm v8 extra CRC32 instructions."},
	{34, "ARM64", "PF_ARM_V81_ATOMIC_INSTRUCTIONS_AVAILABLE", "This Arm processor implements the Arm v8.1 atomic instructions (for example, CAS, SWP)."},
	{43, "ARM64", "PF_ARM_V82_DP_INSTRUCTIONS_AVAILABLE", "This Arm processor implements the Arm v8.2 DP instructions (for example, SDOT, UDOT). This feature is optional in Arm v8.2 implementations and mandatory in Arm v8.4 implementations."},
	{44, "ARM64", "PF_ARM_V83_JSCVT_INSTRUCTIONS_AVAILABLE", "This Arm processor implements the Arm v8.3 JSCVT instructions (for example, FJCVTZS)."},
	{45, "ARM64", "PF_ARM_V83_LRCPC_INSTRUCTIONS_AVAILABLE", "This Arm processor implements the Arm v8.3 LRCPC instructions (for example, LDAPR). Note that certain Arm v8.2 CPUs may optionally support the LRCPC instructions."},
}

// IsProcessorFeaturesPresent returns true when the feature identified by the number is present.
func IsProcessorFeaturesPresent(i uint32) (bool, error) {
	d, e := syscall.LoadDLL("kernel32.dll")
	if e != nil {
		return false, e
	}
	p, e := d.FindProc("IsProcessorFeaturePresent")
	if e != nil {
		return false, e
	}
	r1, _, _ := syscall.SyscallN(p.Addr(), uintptr(i))
	return r1 != 0, nil
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
				// TODO error is only about loading DLL
				missing = append(missing, f.i)
			}
		}
	}
	return missing
}
