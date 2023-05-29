//go:build ppc

// Code generated using 'go generate'; DO NOT EDIT.
// package buildref contains the the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
	{0, "", "pc_none", "R_PPC_NONE  0"},
	{1, "", "pc_addr32", "R_PPC_ADDR32  1"},
	{2, "", "pc_addr24", "R_PPC_ADDR24  2"},
	{3, "", "pc_addr16", "R_PPC_ADDR16  3"},
	{4, "", "pc_addr16_lo", "R_PPC_ADDR16_LO  4"},
	{5, "", "pc_addr16_hi", "R_PPC_ADDR16_HI  5"},
	{6, "", "pc_addr16_ha", "R_PPC_ADDR16_HA  6"},
	{7, "", "pc_addr14", "R_PPC_ADDR14  7"},
	{8, "", "pc_addr14_brtaken", "R_PPC_ADDR14_BRTAKEN 8"},
	{9, "", "pc_addr14_brntaken", "R_PPC_ADDR14_BRNTAKEN 9"},
	{10, "", "pc_rel24", "R_PPC_REL24  10"},
	{11, "", "pc_rel14", "R_PPC_REL14  11"},
	{12, "", "pc_rel14_brtaken", "R_PPC_REL14_BRTAKEN 12"},
	{13, "", "pc_rel14_brntaken", "R_PPC_REL14_BRNTAKEN 13"},
	{14, "", "pc_got16", "R_PPC_GOT16  14"},
	{15, "", "pc_got16_lo", "R_PPC_GOT16_LO  15"},
	{16, "", "pc_got16_hi", "R_PPC_GOT16_HI  16"},
	{17, "", "pc_got16_ha", "R_PPC_GOT16_HA  17"},
	{18, "", "pc_pltrel24", "R_PPC_PLTREL24  18"},
	{19, "", "pc_copy", "R_PPC_COPY  19"},
	{20, "", "pc_glob_dat", "R_PPC_GLOB_DAT  20"},
	{21, "", "pc_jmp_slot", "R_PPC_JMP_SLOT  21"},
	{22, "", "pc_relative", "R_PPC_RELATIVE  22"},
	{23, "", "pc_local24pc", "R_PPC_LOCAL24PC  23"},
	{24, "", "pc_uaddr32", "R_PPC_UADDR32  24"},
	{25, "", "pc_uaddr16", "R_PPC_UADDR16  25"},
	{26, "", "pc_rel32", "R_PPC_REL32  26"},
	{27, "", "pc_plt32", "R_PPC_PLT32  27"},
	{28, "", "pc_pltrel32", "R_PPC_PLTREL32  28"},
	{29, "", "pc_plt16_lo", "R_PPC_PLT16_LO  29"},
	{30, "", "pc_plt16_hi", "R_PPC_PLT16_HI  30"},
	{31, "", "pc_plt16_ha", "R_PPC_PLT16_HA  31"},
	{32, "", "pc_sdarel16", "R_PPC_SDAREL16  32"},
	{33, "", "pc_sectoff", "R_PPC_SECTOFF  33"},
	{34, "", "pc_sectoff_lo", "R_PPC_SECTOFF_LO 34"},
	{35, "", "pc_sectoff_hi", "R_PPC_SECTOFF_HI 35"},
	{36, "", "pc_sectoff_ha", "R_PPC_SECTOFF_HA 36"},
}
