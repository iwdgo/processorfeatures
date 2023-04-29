// Package processorfeatures provides details on the available processor features
package processorfeatures

import "os"

type ProcessorFeature struct {
	i   uint32 // Feature number
	v   string // Version when defined by standard.
	s   string // Flag name
	doc string // Documentation
}

// SetGOAMD64v sets GOAMD64 for the current go session. Default v1 is used when no higher version is complete.
// Version set is returned or an error.
func SetGOAMD64v() (v string, err error) {
	for _, v = range []string{"v4", "v3", "v2"} {
		if m := IsVersionComplete(v); len(m) == 0 {
			err = os.Setenv("GOAMD64", v)
			return
		}
	}
	v = "v1"
	err = os.Setenv("GOAMD64", v)
	return
}

// IsVersionComplete verifies all features listed for version v
func IsVersionComplete(v string) (missing []uint32) {
	for _, f := range ProcessorFeatures {
		if f.v == v {
			if b, err := IsProcessorFeaturesPresent(f.i); err == nil {
				if !b {
					missing = append(missing, f.i)
				}
			} else {
				// TODO error is only about loading DLL
				missing = append(missing, f.i)
			}
		}
	}
	return missing
}

// SetGOARMv sets GOARM and GOARM64 and returns ARM64, 7 or 6 depending on identified processor features.
func SetGOARMv() (v string, err error) {
	v = "arm64"
	if m := IsVersionComplete(v); len(m) == 0 {
		err = os.Unsetenv("GOARM")
		if err != nil {
			return
		}
		err = os.Setenv("GOARCH", "ARM64")
		if err != nil {
			return
		}
		return
	}
	err = os.Setenv("GOARCH", "ARM")
	if err != nil {
		return
	}
	v = "7"
	if m := IsVersionComplete(v); len(m) == 0 {
		err = os.Setenv("GOARM", v)
		if err != nil {
			return
		}
		return
	}
	// TODO Check that 6 is set for all
	v = "6" // Default
	err = os.Setenv("GOARM", v)
	if err != nil {
		return
	}
	return
}
