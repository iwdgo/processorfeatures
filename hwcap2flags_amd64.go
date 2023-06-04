//go:build amd64

// Code generated using 'go generate'; DO NOT EDIT.
// hwcap2flags files contain the HWCAP flags for each architecture
//
//go:generate go run mkhwcap2flags.go
package processorfeatures

var AuxvFeatures = []ProcessorFeature{
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
}
