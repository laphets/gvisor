// automatically generated by stateify.

package host

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (d *descriptor) StateTypeName() string {
	return "pkg/sentry/fs/host.descriptor"
}

func (d *descriptor) StateFields() []string {
	return []string{
		"origFD",
		"wouldBlock",
	}
}

// +checklocksignore
func (d *descriptor) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.origFD)
	stateSinkObject.Save(1, &d.wouldBlock)
}

// +checklocksignore
func (d *descriptor) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.origFD)
	stateSourceObject.Load(1, &d.wouldBlock)
	stateSourceObject.AfterLoad(d.afterLoad)
}

func (f *fileOperations) StateTypeName() string {
	return "pkg/sentry/fs/host.fileOperations"
}

func (f *fileOperations) StateFields() []string {
	return []string{
		"iops",
		"dirCursor",
	}
}

func (f *fileOperations) beforeSave() {}

// +checklocksignore
func (f *fileOperations) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.iops)
	stateSinkObject.Save(1, &f.dirCursor)
}

func (f *fileOperations) afterLoad() {}

// +checklocksignore
func (f *fileOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadWait(0, &f.iops)
	stateSourceObject.Load(1, &f.dirCursor)
}

func (f *filesystem) StateTypeName() string {
	return "pkg/sentry/fs/host.filesystem"
}

func (f *filesystem) StateFields() []string {
	return []string{}
}

func (f *filesystem) beforeSave() {}

// +checklocksignore
func (f *filesystem) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *filesystem) afterLoad() {}

// +checklocksignore
func (f *filesystem) StateLoad(stateSourceObject state.Source) {
}

func (i *inodeOperations) StateTypeName() string {
	return "pkg/sentry/fs/host.inodeOperations"
}

func (i *inodeOperations) StateFields() []string {
	return []string{
		"fileState",
		"cachingInodeOps",
	}
}

func (i *inodeOperations) beforeSave() {}

// +checklocksignore
func (i *inodeOperations) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.fileState)
	stateSinkObject.Save(1, &i.cachingInodeOps)
}

func (i *inodeOperations) afterLoad() {}

// +checklocksignore
func (i *inodeOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadWait(0, &i.fileState)
	stateSourceObject.Load(1, &i.cachingInodeOps)
}

func (i *inodeFileState) StateTypeName() string {
	return "pkg/sentry/fs/host.inodeFileState"
}

func (i *inodeFileState) StateFields() []string {
	return []string{
		"descriptor",
		"sattr",
		"savedUAttr",
	}
}

func (i *inodeFileState) beforeSave() {}

// +checklocksignore
func (i *inodeFileState) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	if !state.IsZeroValue(&i.queue) {
		state.Failf("queue is %#v, expected zero", &i.queue)
	}
	stateSinkObject.Save(0, &i.descriptor)
	stateSinkObject.Save(1, &i.sattr)
	stateSinkObject.Save(2, &i.savedUAttr)
}

// +checklocksignore
func (i *inodeFileState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadWait(0, &i.descriptor)
	stateSourceObject.LoadWait(1, &i.sattr)
	stateSourceObject.Load(2, &i.savedUAttr)
	stateSourceObject.AfterLoad(i.afterLoad)
}

func (c *ConnectedEndpoint) StateTypeName() string {
	return "pkg/sentry/fs/host.ConnectedEndpoint"
}

func (c *ConnectedEndpoint) StateFields() []string {
	return []string{
		"ref",
		"queue",
		"path",
		"srfd",
		"stype",
	}
}

// +checklocksignore
func (c *ConnectedEndpoint) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.ref)
	stateSinkObject.Save(1, &c.queue)
	stateSinkObject.Save(2, &c.path)
	stateSinkObject.Save(3, &c.srfd)
	stateSinkObject.Save(4, &c.stype)
}

// +checklocksignore
func (c *ConnectedEndpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.ref)
	stateSourceObject.Load(1, &c.queue)
	stateSourceObject.Load(2, &c.path)
	stateSourceObject.LoadWait(3, &c.srfd)
	stateSourceObject.Load(4, &c.stype)
	stateSourceObject.AfterLoad(c.afterLoad)
}

func (t *TTYFileOperations) StateTypeName() string {
	return "pkg/sentry/fs/host.TTYFileOperations"
}

func (t *TTYFileOperations) StateFields() []string {
	return []string{
		"fileOperations",
		"session",
		"fgProcessGroup",
		"termios",
	}
}

func (t *TTYFileOperations) beforeSave() {}

// +checklocksignore
func (t *TTYFileOperations) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.fileOperations)
	stateSinkObject.Save(1, &t.session)
	stateSinkObject.Save(2, &t.fgProcessGroup)
	stateSinkObject.Save(3, &t.termios)
}

func (t *TTYFileOperations) afterLoad() {}

// +checklocksignore
func (t *TTYFileOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.fileOperations)
	stateSourceObject.Load(1, &t.session)
	stateSourceObject.Load(2, &t.fgProcessGroup)
	stateSourceObject.Load(3, &t.termios)
}

func init() {
	state.Register((*descriptor)(nil))
	state.Register((*fileOperations)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*inodeOperations)(nil))
	state.Register((*inodeFileState)(nil))
	state.Register((*ConnectedEndpoint)(nil))
	state.Register((*TTYFileOperations)(nil))
}
