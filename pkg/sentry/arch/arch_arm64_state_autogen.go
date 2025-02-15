// automatically generated by stateify.

//go:build arm64 && arm64 && arm64
// +build arm64,arm64,arm64

package arch

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (c *context64) StateTypeName() string {
	return "pkg/sentry/arch.context64"
}

func (c *context64) StateFields() []string {
	return []string{
		"State",
		"sigFPState",
	}
}

func (c *context64) beforeSave() {}

// +checklocksignore
func (c *context64) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.State)
	stateSinkObject.Save(1, &c.sigFPState)
}

func (c *context64) afterLoad() {}

// +checklocksignore
func (c *context64) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.State)
	stateSourceObject.Load(1, &c.sigFPState)
}

func init() {
	state.Register((*context64)(nil))
}
