//go:build linux && amd64

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var MachineFeatures = []ProcessorFeature{
	{0, "", "none", "NONE  0"},
	{1, "", "64", "64  1"},
	{2, "", "pc32", "PC32  2"},
	{3, "", "got32", "GOT32  3"},
	{4, "", "plt32", "PLT32  4"},
	{5, "", "copy", "COPY  5"},
	{6, "", "glob_dat", "GLOB_DAT 6"},
	{7, "", "jump_slot", "JUMP_SLOT 7"},
	{8, "", "relative", "RELATIVE 8"},
	{9, "", "gotpcrel", "GOTPCREL 9"},
	{10, "", "32", "32  10"},
	{11, "", "32s", "32S  11"},
	{12, "", "16", "16  12"},
	{13, "", "pc16", "PC16  13"},
	{14, "", "8", "8  14"},
	{15, "", "pc8", "PC8  15"},
	{16, "", "dtpmod64", "DTPMOD64 16"},
	{17, "", "dtpoff64", "DTPOFF64 17"},
	{18, "", "tpoff64", "TPOFF64 18"},
	{19, "", "tlsgd", "TLSGD  19"},
	{20, "", "tlsld", "TLSLD  20"},
	{21, "", "dtpoff32", "DTPOFF32 21"},
	{22, "", "gottpoff", "GOTTPOFF 22"},
	{23, "", "tpoff32", "TPOFF32 23"},
	{24, "", "pc64", "PC64  24"},
	{25, "", "gotoff64", "GOTOFF64 25"},
	{26, "", "gotpc32", "GOTPC32 26"},
	{27, "", "got64", "GOT64  27"},
	{28, "", "gotpcrel64", "GOTPCREL64 28"},
	{29, "", "gotpc64", "GOTPC64 29"},
	{30, "", "gotplt64", "GOTPLT64 30"},
	{31, "", "pltoff64", "PLTOFF64 31"},
	{32, "", "size32", "SIZE32  32"},
	{33, "", "size64", "SIZE64  33"},
	{34, "", "gotpc32_tlsdesc", "GOTPC32_TLSDESC 34"},
	{35, "", "tlsdesc_call", "TLSDESC_CALL   35"},
	{36, "", "tlsdesc", "TLSDESC        36"},
	{37, "", "irelative", "IRELATIVE 37"},
	{38, "", "relative64", "RELATIVE64 38"},
	{41, "", "gotpcrelx", "GOTPCRELX 41"},
	{42, "", "rex_gotpcrelx", "REX_GOTPCRELX 42"},
	{43, "", "num", "NUM  43"},
}
