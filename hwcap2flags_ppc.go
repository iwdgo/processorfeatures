//go:build linux && ppc

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var MachineFeatures = []ProcessorFeature{
	{0, "", "none", "NONE  0"},
	{1, "", "addr32", "ADDR32  1"},
	{2, "", "addr24", "ADDR24  2"},
	{3, "", "addr16", "ADDR16  3"},
	{4, "", "addr16_lo", "ADDR16_LO  4"},
	{5, "", "addr16_hi", "ADDR16_HI  5"},
	{6, "", "addr16_ha", "ADDR16_HA  6"},
	{7, "", "addr14", "ADDR14  7"},
	{8, "", "addr14_brtaken", "ADDR14_BRTAKEN 8"},
	{9, "", "addr14_brntaken", "ADDR14_BRNTAKEN 9"},
	{10, "", "rel24", "REL24  10"},
	{11, "", "rel14", "REL14  11"},
	{12, "", "rel14_brtaken", "REL14_BRTAKEN 12"},
	{13, "", "rel14_brntaken", "REL14_BRNTAKEN 13"},
	{14, "", "got16", "GOT16  14"},
	{15, "", "got16_lo", "GOT16_LO  15"},
	{16, "", "got16_hi", "GOT16_HI  16"},
	{17, "", "got16_ha", "GOT16_HA  17"},
	{18, "", "pltrel24", "PLTREL24  18"},
	{19, "", "copy", "COPY  19"},
	{20, "", "glob_dat", "GLOB_DAT  20"},
	{21, "", "jmp_slot", "JMP_SLOT  21"},
	{22, "", "relative", "RELATIVE  22"},
	{23, "", "local24pc", "LOCAL24PC  23"},
	{24, "", "uaddr32", "UADDR32  24"},
	{25, "", "uaddr16", "UADDR16  25"},
	{26, "", "rel32", "REL32  26"},
	{27, "", "plt32", "PLT32  27"},
	{28, "", "pltrel32", "PLTREL32  28"},
	{29, "", "plt16_lo", "PLT16_LO  29"},
	{30, "", "plt16_hi", "PLT16_HI  30"},
	{31, "", "plt16_ha", "PLT16_HA  31"},
	{32, "", "sdarel16", "SDAREL16  32"},
	{33, "", "sectoff", "SECTOFF  33"},
	{34, "", "sectoff_lo", "SECTOFF_LO 34"},
	{35, "", "sectoff_hi", "SECTOFF_HI 35"},
	{36, "", "sectoff_ha", "SECTOFF_HA 36"},
}
