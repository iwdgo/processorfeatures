//go:build linux && sparc

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var MachineFeatures = []ProcessorFeature{
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
	{24, "", "plt32", "PLT32  24"},
	{25, "", "hiplt22", "HIPLT22  25"},
	{26, "", "loplt10", "LOPLT10  26"},
	{27, "", "pcplt32", "PCPLT32  27"},
	{28, "", "pcplt22", "PCPLT22  28"},
	{29, "", "pcplt10", "PCPLT10  29"},
	{30, "", "10", "10  30"},
	{31, "", "11", "11  31"},
	{32, "", "64", "64  32"},
	{33, "", "olo10", "OLO10  33"},
	{34, "", "hh22", "HH22  34"},
	{35, "", "hm10", "HM10  35"},
	{36, "", "lm22", "LM22  36"},
	{37, "", "pc_hh22", "PC_HH22  37"},
	{38, "", "pc_hm10", "PC_HM10  38"},
	{39, "", "pc_lm22", "PC_LM22  39"},
	{40, "", "wdisp16", "WDISP16  40"},
	{41, "", "wdisp19", "WDISP19  41"},
	{42, "", "glob_jmp", "GLOB_JMP 42"},
	{43, "", "7", "7  43"},
	{44, "", "5", "5  44"},
	{45, "", "6", "6  45"},
	{46, "", "disp64", "DISP64  46"},
	{47, "", "plt64", "PLT64  47"},
	{48, "", "hix22", "HIX22  48"},
	{49, "", "lox10", "LOX10  49"},
	{50, "", "h44", "H44  50"},
	{51, "", "m44", "M44  51"},
	{52, "", "l44", "L44  52"},
	{53, "", "register", "REGISTER 53"},
	{54, "", "ua64", "UA64  54"},
	{55, "", "ua16", "UA16  55"},
	{56, "", "tls_gd_hi22", "TLS_GD_HI22 56"},
	{57, "", "tls_gd_lo10", "TLS_GD_LO10 57"},
	{58, "", "tls_gd_add", "TLS_GD_ADD 58"},
	{59, "", "tls_gd_call", "TLS_GD_CALL 59"},
	{60, "", "tls_ldm_hi22", "TLS_LDM_HI22 60"},
	{61, "", "tls_ldm_lo10", "TLS_LDM_LO10 61"},
	{62, "", "tls_ldm_add", "TLS_LDM_ADD 62"},
	{63, "", "tls_ldm_call", "TLS_LDM_CALL 63"},
	{64, "", "tls_ldo_hix22", "TLS_LDO_HIX22 64"},
	{65, "", "tls_ldo_lox10", "TLS_LDO_LOX10 65"},
	{66, "", "tls_ldo_add", "TLS_LDO_ADD 66"},
	{67, "", "tls_ie_hi22", "TLS_IE_HI22 67"},
	{68, "", "tls_ie_lo10", "TLS_IE_LO10 68"},
	{69, "", "tls_ie_ld", "TLS_IE_LD 69"},
	{70, "", "tls_ie_ldx", "TLS_IE_LDX 70"},
	{71, "", "tls_ie_add", "TLS_IE_ADD 71"},
	{72, "", "tls_le_hix22", "TLS_LE_HIX22 72"},
	{73, "", "tls_le_lox10", "TLS_LE_LOX10 73"},
	{74, "", "tls_dtpmod32", "TLS_DTPMOD32 74"},
	{75, "", "tls_dtpmod64", "TLS_DTPMOD64 75"},
	{76, "", "tls_dtpoff32", "TLS_DTPOFF32 76"},
	{77, "", "tls_dtpoff64", "TLS_DTPOFF64 77"},
	{78, "", "tls_tpoff32", "TLS_TPOFF32 78"},
	{79, "", "tls_tpoff64", "TLS_TPOFF64 79"},
	{80, "", "gotdata_hix22", "GOTDATA_HIX22 80"},
	{81, "", "gotdata_lox10", "GOTDATA_LOX10 81"},
	{82, "", "gotdata_op_hix22", "GOTDATA_OP_HIX22 82"},
	{83, "", "gotdata_op_lox10", "GOTDATA_OP_LOX10 83"},
	{84, "", "gotdata_op", "GOTDATA_OP 84"},
	{85, "", "h34", "H34  85"},
	{86, "", "size32", "SIZE32  86"},
	{87, "", "size64", "SIZE64  87"},
	{250, "", "gnu_vtinherit", "GNU_VTINHERIT 250"},
	{251, "", "gnu_vtentry", "GNU_VTENTRY 251"},
	{252, "", "rev32", "REV32  252"},
	{253, "", "num", "NUM  253"},
}
