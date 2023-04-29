// Package processorfeatures provides details on the available processor features
package processorfeatures

type ProcessorFeature struct {
	i   uint32 // Feature number
	v   string // Version when defined by standard.
	s   string // Flag name
	doc string // Documentation
}
