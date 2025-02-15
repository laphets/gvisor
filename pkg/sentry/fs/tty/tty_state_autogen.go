// automatically generated by stateify.

package tty

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (d *dirInodeOperations) StateTypeName() string {
	return "pkg/sentry/fs/tty.dirInodeOperations"
}

func (d *dirInodeOperations) StateFields() []string {
	return []string{
		"InodeSimpleAttributes",
		"msrc",
		"master",
		"replicas",
		"dentryMap",
		"next",
	}
}

func (d *dirInodeOperations) beforeSave() {}

// +checklocksignore
func (d *dirInodeOperations) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.InodeSimpleAttributes)
	stateSinkObject.Save(1, &d.msrc)
	stateSinkObject.Save(2, &d.master)
	stateSinkObject.Save(3, &d.replicas)
	stateSinkObject.Save(4, &d.dentryMap)
	stateSinkObject.Save(5, &d.next)
}

func (d *dirInodeOperations) afterLoad() {}

// +checklocksignore
func (d *dirInodeOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.InodeSimpleAttributes)
	stateSourceObject.Load(1, &d.msrc)
	stateSourceObject.Load(2, &d.master)
	stateSourceObject.Load(3, &d.replicas)
	stateSourceObject.Load(4, &d.dentryMap)
	stateSourceObject.Load(5, &d.next)
}

func (df *dirFileOperations) StateTypeName() string {
	return "pkg/sentry/fs/tty.dirFileOperations"
}

func (df *dirFileOperations) StateFields() []string {
	return []string{
		"di",
		"dirCursor",
	}
}

func (df *dirFileOperations) beforeSave() {}

// +checklocksignore
func (df *dirFileOperations) StateSave(stateSinkObject state.Sink) {
	df.beforeSave()
	stateSinkObject.Save(0, &df.di)
	stateSinkObject.Save(1, &df.dirCursor)
}

func (df *dirFileOperations) afterLoad() {}

// +checklocksignore
func (df *dirFileOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &df.di)
	stateSourceObject.Load(1, &df.dirCursor)
}

func (f *filesystem) StateTypeName() string {
	return "pkg/sentry/fs/tty.filesystem"
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

func (s *superOperations) StateTypeName() string {
	return "pkg/sentry/fs/tty.superOperations"
}

func (s *superOperations) StateFields() []string {
	return []string{}
}

func (s *superOperations) beforeSave() {}

// +checklocksignore
func (s *superOperations) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
}

func (s *superOperations) afterLoad() {}

// +checklocksignore
func (s *superOperations) StateLoad(stateSourceObject state.Source) {
}

func (l *lineDiscipline) StateTypeName() string {
	return "pkg/sentry/fs/tty.lineDiscipline"
}

func (l *lineDiscipline) StateFields() []string {
	return []string{
		"size",
		"inQueue",
		"outQueue",
		"termios",
		"column",
	}
}

func (l *lineDiscipline) beforeSave() {}

// +checklocksignore
func (l *lineDiscipline) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	if !state.IsZeroValue(&l.masterWaiter) {
		state.Failf("masterWaiter is %#v, expected zero", &l.masterWaiter)
	}
	if !state.IsZeroValue(&l.replicaWaiter) {
		state.Failf("replicaWaiter is %#v, expected zero", &l.replicaWaiter)
	}
	stateSinkObject.Save(0, &l.size)
	stateSinkObject.Save(1, &l.inQueue)
	stateSinkObject.Save(2, &l.outQueue)
	stateSinkObject.Save(3, &l.termios)
	stateSinkObject.Save(4, &l.column)
}

func (l *lineDiscipline) afterLoad() {}

// +checklocksignore
func (l *lineDiscipline) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.size)
	stateSourceObject.Load(1, &l.inQueue)
	stateSourceObject.Load(2, &l.outQueue)
	stateSourceObject.Load(3, &l.termios)
	stateSourceObject.Load(4, &l.column)
}

func (o *outputQueueTransformer) StateTypeName() string {
	return "pkg/sentry/fs/tty.outputQueueTransformer"
}

func (o *outputQueueTransformer) StateFields() []string {
	return []string{}
}

func (o *outputQueueTransformer) beforeSave() {}

// +checklocksignore
func (o *outputQueueTransformer) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
}

func (o *outputQueueTransformer) afterLoad() {}

// +checklocksignore
func (o *outputQueueTransformer) StateLoad(stateSourceObject state.Source) {
}

func (i *inputQueueTransformer) StateTypeName() string {
	return "pkg/sentry/fs/tty.inputQueueTransformer"
}

func (i *inputQueueTransformer) StateFields() []string {
	return []string{}
}

func (i *inputQueueTransformer) beforeSave() {}

// +checklocksignore
func (i *inputQueueTransformer) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *inputQueueTransformer) afterLoad() {}

// +checklocksignore
func (i *inputQueueTransformer) StateLoad(stateSourceObject state.Source) {
}

func (mi *masterInodeOperations) StateTypeName() string {
	return "pkg/sentry/fs/tty.masterInodeOperations"
}

func (mi *masterInodeOperations) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"d",
	}
}

func (mi *masterInodeOperations) beforeSave() {}

// +checklocksignore
func (mi *masterInodeOperations) StateSave(stateSinkObject state.Sink) {
	mi.beforeSave()
	stateSinkObject.Save(0, &mi.SimpleFileInode)
	stateSinkObject.Save(1, &mi.d)
}

func (mi *masterInodeOperations) afterLoad() {}

// +checklocksignore
func (mi *masterInodeOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mi.SimpleFileInode)
	stateSourceObject.Load(1, &mi.d)
}

func (mf *masterFileOperations) StateTypeName() string {
	return "pkg/sentry/fs/tty.masterFileOperations"
}

func (mf *masterFileOperations) StateFields() []string {
	return []string{
		"d",
		"t",
	}
}

func (mf *masterFileOperations) beforeSave() {}

// +checklocksignore
func (mf *masterFileOperations) StateSave(stateSinkObject state.Sink) {
	mf.beforeSave()
	stateSinkObject.Save(0, &mf.d)
	stateSinkObject.Save(1, &mf.t)
}

func (mf *masterFileOperations) afterLoad() {}

// +checklocksignore
func (mf *masterFileOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mf.d)
	stateSourceObject.Load(1, &mf.t)
}

func (q *queue) StateTypeName() string {
	return "pkg/sentry/fs/tty.queue"
}

func (q *queue) StateFields() []string {
	return []string{
		"readBuf",
		"waitBuf",
		"waitBufLen",
		"readable",
		"transformer",
	}
}

func (q *queue) beforeSave() {}

// +checklocksignore
func (q *queue) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.readBuf)
	stateSinkObject.Save(1, &q.waitBuf)
	stateSinkObject.Save(2, &q.waitBufLen)
	stateSinkObject.Save(3, &q.readable)
	stateSinkObject.Save(4, &q.transformer)
}

func (q *queue) afterLoad() {}

// +checklocksignore
func (q *queue) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.readBuf)
	stateSourceObject.Load(1, &q.waitBuf)
	stateSourceObject.Load(2, &q.waitBufLen)
	stateSourceObject.Load(3, &q.readable)
	stateSourceObject.Load(4, &q.transformer)
}

func (si *replicaInodeOperations) StateTypeName() string {
	return "pkg/sentry/fs/tty.replicaInodeOperations"
}

func (si *replicaInodeOperations) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"d",
		"t",
	}
}

func (si *replicaInodeOperations) beforeSave() {}

// +checklocksignore
func (si *replicaInodeOperations) StateSave(stateSinkObject state.Sink) {
	si.beforeSave()
	stateSinkObject.Save(0, &si.SimpleFileInode)
	stateSinkObject.Save(1, &si.d)
	stateSinkObject.Save(2, &si.t)
}

func (si *replicaInodeOperations) afterLoad() {}

// +checklocksignore
func (si *replicaInodeOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &si.SimpleFileInode)
	stateSourceObject.Load(1, &si.d)
	stateSourceObject.Load(2, &si.t)
}

func (sf *replicaFileOperations) StateTypeName() string {
	return "pkg/sentry/fs/tty.replicaFileOperations"
}

func (sf *replicaFileOperations) StateFields() []string {
	return []string{
		"si",
	}
}

func (sf *replicaFileOperations) beforeSave() {}

// +checklocksignore
func (sf *replicaFileOperations) StateSave(stateSinkObject state.Sink) {
	sf.beforeSave()
	stateSinkObject.Save(0, &sf.si)
}

func (sf *replicaFileOperations) afterLoad() {}

// +checklocksignore
func (sf *replicaFileOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &sf.si)
}

func (tm *Terminal) StateTypeName() string {
	return "pkg/sentry/fs/tty.Terminal"
}

func (tm *Terminal) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"n",
		"d",
		"ld",
		"masterKTTY",
		"replicaKTTY",
	}
}

func (tm *Terminal) beforeSave() {}

// +checklocksignore
func (tm *Terminal) StateSave(stateSinkObject state.Sink) {
	tm.beforeSave()
	stateSinkObject.Save(0, &tm.AtomicRefCount)
	stateSinkObject.Save(1, &tm.n)
	stateSinkObject.Save(2, &tm.d)
	stateSinkObject.Save(3, &tm.ld)
	stateSinkObject.Save(4, &tm.masterKTTY)
	stateSinkObject.Save(5, &tm.replicaKTTY)
}

func (tm *Terminal) afterLoad() {}

// +checklocksignore
func (tm *Terminal) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &tm.AtomicRefCount)
	stateSourceObject.Load(1, &tm.n)
	stateSourceObject.Load(2, &tm.d)
	stateSourceObject.Load(3, &tm.ld)
	stateSourceObject.Load(4, &tm.masterKTTY)
	stateSourceObject.Load(5, &tm.replicaKTTY)
}

func init() {
	state.Register((*dirInodeOperations)(nil))
	state.Register((*dirFileOperations)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*superOperations)(nil))
	state.Register((*lineDiscipline)(nil))
	state.Register((*outputQueueTransformer)(nil))
	state.Register((*inputQueueTransformer)(nil))
	state.Register((*masterInodeOperations)(nil))
	state.Register((*masterFileOperations)(nil))
	state.Register((*queue)(nil))
	state.Register((*replicaInodeOperations)(nil))
	state.Register((*replicaFileOperations)(nil))
	state.Register((*Terminal)(nil))
}
