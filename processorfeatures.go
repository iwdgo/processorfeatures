// Package processorfeatures provides details on the available processor features
package processorfeatures

import (
	"os"
)

type ProcessorFeature struct {
	I   uint32 // Feature number
	V   string // Version when defined by standard.
	S   string // Flag name
	Doc string // Documentation
}

// IsProcessorFeaturePresent returns true when the feature identified by the number is present.
func IsProcessorFeaturePresent(i uint32) (bool, error) {
	return isProcessorFeaturePresent(i)
}

// IsMachineFeaturePresent returns true when the feature identified by the number is present.
func IsMachineFeaturePresent(i uint32) (bool, error) {
	return isMachineFeaturePresent(i)
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
		if f.V == v {
			if b, err := IsProcessorFeaturePresent(f.I); err == nil {
				if !b {
					missing = append(missing, f.I)
				}
			} else {
				missing = append(missing, f.I)
			}
		}
	}
	return missing
}

// IsMachineComplete verifies all machine features listed for version v
func IsMachineComplete(v string) (missing []uint32) {
	for _, f := range MachineFeatures {
		if f.V == v {
			if b, err := IsMachineFeaturePresent(f.I); err == nil {
				if !b {
					missing = append(missing, f.I)
				}
			} else {
				missing = append(missing, f.I)
			}
		}
	}
	return missing
}

// SetGOARMv sets GOARM and GOARCH and returns arm64, 7 or 6 depending on identified processor features.
func SetGOARMv() (v string, err error) {
	v = "arm64"
	// TODO Add check on abs64 as all ARM under docker have the same implementation
	if m := IsVersionComplete(v); len(m) == 0 && !Bits32 {
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
	// TODO Differentiate using machines
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
	v = "6"
	if m := IsVersionComplete(v); len(m) == 0 {
		err = os.Setenv("GOARM", v)
		if err != nil {
			return
		}
		return
	}
	v = "5" // Default
	err = os.Setenv("GOARM", v)
	if err != nil {
		return
	}
	return
}
