[![Go Reference](https://pkg.go.dev/badge/github.com/iwdgo/processorfeatures.svg)](https://pkg.go.dev/github.com/iwdgo/processorfeatures)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/processorfeatures)](https://goreportcard.com/report/github.com/iwdgo/processorfeatures)
[![codecov](https://codecov.io/gh/iwdgo/processorfeatures/branch/master/graph/badge.svg)](https://codecov.io/gh/iwdgo/processorfeatures)

[![Go](https://github.com/iwdgo/processorfeatures/actions/workflows/go.yml/badge.svg)](https://github.com/iwdgo/processorfeatures/actions/workflows/go.yml)
[![Build Status](https://app.travis-ci.com/iwdgo/processorfeatures.svg?branch=master)](https://app.travis-ci.com/iwdgo/processorfeatures)

# Processor characteristics 

[Processor variables](https://go.dev/doc/install/source#environment) direct the Go compiler to use some processor instructions.
The variable has an effect only when cross-compiling or re-compiling go locally.
Purpose of the module is to detect processor characteristics and set the relevant processor variable accordingly.

## Windows

Win32 API call `IsProcessorFeaturePresent` is used to collect the features of the processor.

### AMD64

Several checks are available but not every instruction of each version.
Impact should be limited as the standard expects instructions to all be present for each version.
Level is set when all verifiable features are present.
Details are available in the structure `ProcessorFeatures`.

### ARM

The availability of VFP defines level 6 or higher for GOARM.
If v8 instruction set is available, GOARM is unset and GOARCH is set to ARM64. 

# References

### AMD64
[Processor features on Win32 API](https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-isprocessorfeaturepresent)
[Microarchitecture_levels](https://en.wikipedia.org/wiki/X86-64#Microarchitecture_levels)

A more detailed [discussion](https://github.com/golang/go/issues/50589) confirms that default is GOAMD64 is v1.

### ARM

https://stackoverflow.com/questions/19844575/how-to-do-division-in-arm

### Version history

`v0.1.0` Published. Windows only.
