//go:build riscv

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
	{0, "", "none", "NONE            0"},
	{1, "", "32", "32              1"},
	{2, "", "64", "64              2"},
	{3, "", "relative", "RELATIVE        3"},
	{4, "", "copy", "COPY            4"},
	{5, "", "jump_slot", "JUMP_SLOT       5"},
	{6, "", "tls_dtpmod32", "TLS_DTPMOD32    6"},
	{7, "", "tls_dtpmod64", "TLS_DTPMOD64    7"},
	{8, "", "tls_dtprel32", "TLS_DTPREL32    8"},
	{9, "", "tls_dtprel64", "TLS_DTPREL64    9"},
	{10, "", "tls_tprel32", "TLS_TPREL32     10"},
	{11, "", "tls_tprel64", "TLS_TPREL64     11"},
}
