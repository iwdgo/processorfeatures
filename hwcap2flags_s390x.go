//go:build s390x

// Code generated using 'go generate'; DO NOT EDIT.
// package buildref contains the the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
	{0, "", "90_none", "R_390_NONE  0"},
	{1, "", "90_8", "R_390_8   1"},
	{2, "", "90_12", "R_390_12  2"},
	{3, "", "90_16", "R_390_16  3"},
	{4, "", "90_32", "R_390_32  4"},
	{5, "", "90_pc32", "R_390_PC32  5"},
	{6, "", "90_got12", "R_390_GOT12  6"},
	{7, "", "90_got32", "R_390_GOT32  7"},
	{8, "", "90_plt32", "R_390_PLT32  8"},
	{9, "", "90_copy", "R_390_COPY  9"},
	{10, "", "90_glob_dat", "R_390_GLOB_DAT  10"},
	{11, "", "90_jmp_slot", "R_390_JMP_SLOT  11"},
	{12, "", "90_relative", "R_390_RELATIVE  12"},
	{13, "", "90_gotoff32", "R_390_GOTOFF32  13"},
	{14, "", "90_gotpc", "R_390_GOTPC  14"},
	{15, "", "90_got16", "R_390_GOT16  15"},
	{16, "", "90_pc16", "R_390_PC16  16"},
	{17, "", "90_pc16dbl", "R_390_PC16DBL  17"},
	{18, "", "90_plt16dbl", "R_390_PLT16DBL  18"},
	{19, "", "90_pc32dbl", "R_390_PC32DBL  19"},
	{20, "", "90_plt32dbl", "R_390_PLT32DBL  20"},
	{21, "", "90_gotpcdbl", "R_390_GOTPCDBL  21"},
	{22, "", "90_64", "R_390_64  22"},
	{23, "", "90_pc64", "R_390_PC64  23"},
	{24, "", "90_got64", "R_390_GOT64  24"},
	{25, "", "90_plt64", "R_390_PLT64  25"},
	{26, "", "90_gotent", "R_390_GOTENT  26"},
	{27, "", "90_gotoff16", "R_390_GOTOFF16  27"},
	{28, "", "90_gotoff64", "R_390_GOTOFF64  28"},
	{29, "", "90_gotplt12", "R_390_GOTPLT12  29"},
	{30, "", "90_gotplt16", "R_390_GOTPLT16  30"},
	{31, "", "90_gotplt32", "R_390_GOTPLT32  31"},
	{32, "", "90_gotplt64", "R_390_GOTPLT64  32"},
	{33, "", "90_gotpltent", "R_390_GOTPLTENT  33"},
	{34, "", "90_pltoff16", "R_390_PLTOFF16  34"},
	{35, "", "90_pltoff32", "R_390_PLTOFF32  35"},
	{36, "", "90_pltoff64", "R_390_PLTOFF64  36"},
	{37, "", "90_tls_load", "R_390_TLS_LOAD  37"},
	{38, "", "90_tls_gdcall", "R_390_TLS_GDCALL 38"},
}
