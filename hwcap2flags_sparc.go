//go:build sparc

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
	{0, "", "none", "NONE  0"},
	{1, "", "8", "8  1"},
	{2, "", "16", "16  2"},
	{3, "", "32", "32  3"},
	{4, "", "disp8", "DISP8  4"},
	{5, "", "disp16", "DISP16  5"},
	{6, "", "disp32", "DISP32  6"},
	{7, "", "wdisp30", "WDISP30  7"},
	{8, "", "wdisp22", "WDISP22  8"},
	{9, "", "hi22", "HI22  9"},
	{10, "", "22", "22  10"},
	{11, "", "13", "13  11"},
	{12, "", "lo10", "LO10  12"},
	{13, "", "got10", "GOT10  13"},
	{14, "", "got13", "GOT13  14"},
	{15, "", "got22", "GOT22  15"},
	{16, "", "pc10", "PC10  16"},
	{17, "", "pc22", "PC22  17"},
	{18, "", "wplt30", "WPLT30  18"},
	{19, "", "copy", "COPY  19"},
	{20, "", "glob_dat", "GLOB_DAT 20"},
	{21, "", "jmp_slot", "JMP_SLOT 21"},
	{22, "", "relative", "RELATIVE 22"},
	{23, "", "ua32", "UA32  23"},
}
