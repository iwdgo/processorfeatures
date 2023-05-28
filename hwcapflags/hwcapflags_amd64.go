//go:build amd64

// Code generated using 'go generate'; DO NOT EDIT.
// package buildref contains the the HWCAP flags for each architecture
//
//go:generate go run mkhwcapflags.go
package hwcapflags

import "github.com/iwdgo/processorfeatures"

var ProcessorFeatures = []processorfeatures.ProcessorFeature{
	{0, "", "64_none", "R_X86_64_NONE  0"},
	{1, "", "64_64", "R_X86_64_64  1"},
	{2, "", "64_pc32", "R_X86_64_PC32  2"},
	{3, "", "64_got32", "R_X86_64_GOT32  3"},
	{4, "", "64_plt32", "R_X86_64_PLT32  4"},
	{5, "", "64_copy", "R_X86_64_COPY  5"},
	{6, "", "64_glob_dat", "R_X86_64_GLOB_DAT 6"},
	{7, "", "64_jump_slot", "R_X86_64_JUMP_SLOT 7"},
	{8, "", "64_relative", "R_X86_64_RELATIVE 8"},
	{9, "", "64_gotpcrel", "R_X86_64_GOTPCREL 9"},
}
