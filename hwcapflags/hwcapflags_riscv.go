//go:build riscv

// Code generated using 'go generate'; DO NOT EDIT.
// package buildref contains the the HWCAP flags for each architecture
//
//go:generate go run mkhwcapflags.go
package hwcapflags

import "github.com/iwdgo/processorfeatures"

var ProcessorFeatures = []processorfeatures.ProcessorFeature{
	{0, "", "cv_none", "R_RISCV_NONE            0"},
	{1, "", "cv_32", "R_RISCV_32              1"},
	{2, "", "cv_64", "R_RISCV_64              2"},
	{3, "", "cv_relative", "R_RISCV_RELATIVE        3"},
	{4, "", "cv_copy", "R_RISCV_COPY            4"},
	{5, "", "cv_jump_slot", "R_RISCV_JUMP_SLOT       5"},
	{6, "", "cv_tls_dtpmod32", "R_RISCV_TLS_DTPMOD32    6"},
	{7, "", "cv_tls_dtpmod64", "R_RISCV_TLS_DTPMOD64    7"},
	{8, "", "cv_tls_dtprel32", "R_RISCV_TLS_DTPREL32    8"},
	{9, "", "cv_tls_dtprel64", "R_RISCV_TLS_DTPREL64    9"},
	{10, "", "cv_tls_tprel32", "R_RISCV_TLS_TPREL32     10"},
	{11, "", "cv_tls_tprel64", "R_RISCV_TLS_TPREL64     11"},
}
