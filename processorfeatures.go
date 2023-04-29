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
