//go:build 386

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
	{0, "", "none", "NONE    0"},
	{1, "", "32", "32    1"},
	{2, "", "pc32", "PC32    2"},
	{3, "", "got32", "GOT32    3"},
	{4, "", "plt32", "PLT32    4"},
	{5, "", "copy", "COPY    5"},
	{6, "", "glob_dat", "GLOB_DAT    6"},
	{7, "", "jmp_slot", "JMP_SLOT    7"},
	{8, "", "relative", "RELATIVE    8"},
	{9, "", "gotoff", "GOTOFF    9"},
	{10, "", "gotpc", "GOTPC    10"},
	{11, "", "32plt", "32PLT    11"},
	{14, "", "tls_tpoff", "TLS_TPOFF    14"},
	{15, "", "tls_ie", "TLS_IE    15"},
	{16, "", "tls_gotie", "TLS_GOTIE    16"},
	{17, "", "tls_le", "TLS_LE    17"},
	{18, "", "tls_gd", "TLS_GD    18"},
	{19, "", "tls_ldm", "TLS_LDM    19"},
	{20, "", "16", "16    20"},
	{21, "", "pc16", "PC16    21"},
	{22, "", "8", "8     22"},
	{23, "", "pc8", "PC8    23"},
	{24, "", "tls_gd_32", "TLS_GD_32    24"},
	{25, "", "tls_gd_push", "TLS_GD_PUSH  25"},
	{26, "", "tls_gd_call", "TLS_GD_CALL  26"},
	{27, "", "tls_gd_pop", "TLS_GD_POP   27"},
	{28, "", "tls_ldm_32", "TLS_LDM_32   28"},
	{29, "", "tls_ldm_push", "TLS_LDM_PUSH 29"},
	{30, "", "tls_ldm_call", "TLS_LDM_CALL 30"},
	{31, "", "tls_ldm_pop", "TLS_LDM_POP  31"},
	{32, "", "tls_ldo_32", "TLS_LDO_32   32"},
	{33, "", "tls_ie_32", "TLS_IE_32    33"},
	{34, "", "tls_le_32", "TLS_LE_32    34"},
	{35, "", "tls_dtpmod32", "TLS_DTPMOD32 35"},
	{36, "", "tls_dtpoff32", "TLS_DTPOFF32 36"},
	{37, "", "tls_tpoff32", "TLS_TPOFF32  37"},
	{38, "", "size32", "SIZE32       38"},
	{39, "", "tls_gotdesc", "TLS_GOTDESC  39"},
	{40, "", "tls_desc_call", "TLS_DESC_CALL 40"},
	{41, "", "tls_desc", "TLS_DESC     41"},
	{42, "", "irelative", "IRELATIVE    42"},
	{43, "", "got32x", "GOT32X    43"},
	{44, "", "num", "NUM    44"},
}
