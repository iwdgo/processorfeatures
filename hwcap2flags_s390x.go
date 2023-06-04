//go:build s390x

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
	{0, "", "none", "NONE  0"},
	{1, "", "8", "8   1"},
	{2, "", "12", "12  2"},
	{3, "", "16", "16  3"},
	{4, "", "32", "32  4"},
	{5, "", "pc32", "PC32  5"},
	{6, "", "got12", "GOT12  6"},
	{7, "", "got32", "GOT32  7"},
	{8, "", "plt32", "PLT32  8"},
	{9, "", "copy", "COPY  9"},
	{10, "", "glob_dat", "GLOB_DAT  10"},
	{11, "", "jmp_slot", "JMP_SLOT  11"},
	{12, "", "relative", "RELATIVE  12"},
	{13, "", "gotoff32", "GOTOFF32  13"},
	{14, "", "gotpc", "GOTPC  14"},
	{15, "", "got16", "GOT16  15"},
	{16, "", "pc16", "PC16  16"},
	{17, "", "pc16dbl", "PC16DBL  17"},
	{18, "", "plt16dbl", "PLT16DBL  18"},
	{19, "", "pc32dbl", "PC32DBL  19"},
	{20, "", "plt32dbl", "PLT32DBL  20"},
	{21, "", "gotpcdbl", "GOTPCDBL  21"},
	{22, "", "64", "64  22"},
	{23, "", "pc64", "PC64  23"},
	{24, "", "got64", "GOT64  24"},
	{25, "", "plt64", "PLT64  25"},
	{26, "", "gotent", "GOTENT  26"},
	{27, "", "gotoff16", "GOTOFF16  27"},
	{28, "", "gotoff64", "GOTOFF64  28"},
	{29, "", "gotplt12", "GOTPLT12  29"},
	{30, "", "gotplt16", "GOTPLT16  30"},
	{31, "", "gotplt32", "GOTPLT32  31"},
	{32, "", "gotplt64", "GOTPLT64  32"},
	{33, "", "gotpltent", "GOTPLTENT  33"},
	{34, "", "pltoff16", "PLTOFF16  34"},
	{35, "", "pltoff32", "PLTOFF32  35"},
	{36, "", "pltoff64", "PLTOFF64  36"},
	{37, "", "tls_load", "TLS_LOAD  37"},
	{38, "", "tls_gdcall", "TLS_GDCALL 38"},
}
