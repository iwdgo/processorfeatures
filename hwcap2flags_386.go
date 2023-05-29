//go:build 386

// Code generated using 'go generate'; DO NOT EDIT.
// package buildref contains the the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
	{0, "", "86_none", "R_386_NONE    0"},
	{1, "", "86_32", "R_386_32    1"},
	{2, "", "86_pc32", "R_386_PC32    2"},
	{3, "", "86_got32", "R_386_GOT32    3"},
	{4, "", "86_plt32", "R_386_PLT32    4"},
	{5, "", "86_copy", "R_386_COPY    5"},
	{6, "", "86_glob_dat", "R_386_GLOB_DAT    6"},
	{7, "", "86_jmp_slot", "R_386_JMP_SLOT    7"},
	{8, "", "86_relative", "R_386_RELATIVE    8"},
	{9, "", "86_gotoff", "R_386_GOTOFF    9"},
	{10, "", "86_gotpc", "R_386_GOTPC    10"},
	{11, "", "86_32plt", "R_386_32PLT    11"},
	{14, "", "86_tls_tpoff", "R_386_TLS_TPOFF    14"},
	{15, "", "86_tls_ie", "R_386_TLS_IE    15"},
	{16, "", "86_tls_gotie", "R_386_TLS_GOTIE    16"},
	{17, "", "86_tls_le", "R_386_TLS_LE    17"},
	{18, "", "86_tls_gd", "R_386_TLS_GD    18"},
	{19, "", "86_tls_ldm", "R_386_TLS_LDM    19"},
	{20, "", "86_16", "R_386_16    20"},
	{21, "", "86_pc16", "R_386_PC16    21"},
	{22, "", "86_8", "R_386_8     22"},
	{23, "", "86_pc8", "R_386_PC8    23"},
	{24, "", "86_tls_gd_32", "R_386_TLS_GD_32    24"},
	{25, "", "86_tls_gd_push", "R_386_TLS_GD_PUSH  25"},
	{26, "", "86_tls_gd_call", "R_386_TLS_GD_CALL  26"},
	{27, "", "86_tls_gd_pop", "R_386_TLS_GD_POP   27"},
	{28, "", "86_tls_ldm_32", "R_386_TLS_LDM_32   28"},
	{29, "", "86_tls_ldm_push", "R_386_TLS_LDM_PUSH 29"},
	{30, "", "86_tls_ldm_call", "R_386_TLS_LDM_CALL 30"},
	{31, "", "86_tls_ldm_pop", "R_386_TLS_LDM_POP  31"},
	{32, "", "86_tls_ldo_32", "R_386_TLS_LDO_32   32"},
	{33, "", "86_tls_ie_32", "R_386_TLS_IE_32    33"},
	{34, "", "86_tls_le_32", "R_386_TLS_LE_32    34"},
	{35, "", "86_tls_dtpmod32", "R_386_TLS_DTPMOD32 35"},
	{36, "", "86_tls_dtpoff32", "R_386_TLS_DTPOFF32 36"},
	{37, "", "86_tls_tpoff32", "R_386_TLS_TPOFF32  37"},
	{38, "", "86_size32", "R_386_SIZE32       38"},
	{39, "", "86_tls_gotdesc", "R_386_TLS_GOTDESC  39"},
	{40, "", "86_tls_desc_call", "R_386_TLS_DESC_CALL 40"},
	{41, "", "86_tls_desc", "R_386_TLS_DESC     41"},
	{42, "", "86_irelative", "R_386_IRELATIVE    42"},
	{43, "", "86_got32x", "R_386_GOT32X    43"},
	{44, "", "86_num", "R_386_NUM    44"},
}
