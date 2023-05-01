// Package processorfeatures provides details on the available processor features
package processorfeatures

import "os"

type ProcessorFeature struct {
	i   uint32 // Feature number
	v   string // Version when defined by standard.
	s   string // Flag name
	doc string // Documentation
}

// IsProcessorFeaturesPresent returns true when the feature identified by the number is present.
func IsProcessorFeaturesPresent(i uint32) (bool, error) {
	return isProcessorFeaturesPresent(i)
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
		err = os.Setenv("GOARCH", "arm64")
		if err != nil {
			return
		}
		return
	}
	err = os.Setenv("GOARCH", "arm")
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
	// TODO Check that 6 is set for all or return 5
	v = "6" // Default
	err = os.Setenv("GOARM", v)
	if err != nil {
		return
	}
	return
}
