//go:build sparc

// Code generated using 'go generate'; DO NOT EDIT.
// package buildref contains the the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
	{0, "", "rc_none", "R_SPARC_NONE  0"},
	{1, "", "rc_8", "R_SPARC_8  1"},
	{2, "", "rc_16", "R_SPARC_16  2"},
	{3, "", "rc_32", "R_SPARC_32  3"},
	{4, "", "rc_disp8", "R_SPARC_DISP8  4"},
	{5, "", "rc_disp16", "R_SPARC_DISP16  5"},
	{6, "", "rc_disp32", "R_SPARC_DISP32  6"},
	{7, "", "rc_wdisp30", "R_SPARC_WDISP30  7"},
	{8, "", "rc_wdisp22", "R_SPARC_WDISP22  8"},
	{9, "", "rc_hi22", "R_SPARC_HI22  9"},
	{10, "", "rc_22", "R_SPARC_22  10"},
	{11, "", "rc_13", "R_SPARC_13  11"},
	{12, "", "rc_lo10", "R_SPARC_LO10  12"},
	{13, "", "rc_got10", "R_SPARC_GOT10  13"},
	{14, "", "rc_got13", "R_SPARC_GOT13  14"},
	{15, "", "rc_got22", "R_SPARC_GOT22  15"},
	{16, "", "rc_pc10", "R_SPARC_PC10  16"},
	{17, "", "rc_pc22", "R_SPARC_PC22  17"},
	{18, "", "rc_wplt30", "R_SPARC_WPLT30  18"},
	{19, "", "rc_copy", "R_SPARC_COPY  19"},
	{20, "", "rc_glob_dat", "R_SPARC_GLOB_DAT 20"},
	{21, "", "rc_jmp_slot", "R_SPARC_JMP_SLOT 21"},
	{22, "", "rc_relative", "R_SPARC_RELATIVE 22"},
	{23, "", "rc_ua32", "R_SPARC_UA32  23"},
}
