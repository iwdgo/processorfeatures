//go:build linux && ppc64

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var MachineFeatures = []ProcessorFeature{
	{0, "", "none", "NONE  R_PPC_NONE"},
	{1, "", "addr32", "ADDR32  R_PPC_ADDR32"},
	{2, "", "addr24", "ADDR24  R_PPC_ADDR24"},
	{3, "", "addr16", "ADDR16  R_PPC_ADDR16"},
	{4, "", "addr16_lo", "ADDR16_LO R_PPC_ADDR16_LO"},
	{5, "", "addr16_hi", "ADDR16_HI R_PPC_ADDR16_HI"},
	{6, "", "addr16_ha", "ADDR16_HA R_PPC_ADDR16_HA"},
	{7, "", "addr14", "ADDR14  R_PPC_ADDR14"},
	{8, "", "addr14_brtaken", "ADDR14_BRTAKEN R_PPC_ADDR14_BRTAKEN"},
	{9, "", "addr14_brntaken", "ADDR14_BRNTAKEN R_PPC_ADDR14_BRNTAKEN"},
	{10, "", "rel24", "REL24  R_PPC_REL24"},
	{11, "", "rel14", "REL14  R_PPC_REL14"},
	{12, "", "rel14_brtaken", "REL14_BRTAKEN R_PPC_REL14_BRTAKEN"},
	{13, "", "rel14_brntaken", "REL14_BRNTAKEN R_PPC_REL14_BRNTAKEN"},
	{14, "", "got16", "GOT16  R_PPC_GOT16"},
	{15, "", "got16_lo", "GOT16_LO R_PPC_GOT16_LO"},
	{16, "", "got16_hi", "GOT16_HI R_PPC_GOT16_HI"},
	{17, "", "got16_ha", "GOT16_HA R_PPC_GOT16_HA"},
}
