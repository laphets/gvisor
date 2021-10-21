package dpdk

import (
	"fmt"

	"gvisor.dev/gvisor/pkg/context"

	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/errors/linuxerr"
	"gvisor.dev/gvisor/pkg/hostarch"
	"gvisor.dev/gvisor/pkg/marshal"
	"gvisor.dev/gvisor/pkg/sentry/arch"
	"gvisor.dev/gvisor/pkg/sentry/fs"
	"gvisor.dev/gvisor/pkg/sentry/fs/fsutil"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
	ktime "gvisor.dev/gvisor/pkg/sentry/kernel/time"
	"gvisor.dev/gvisor/pkg/sentry/socket"
	"gvisor.dev/gvisor/pkg/syserr"
	"gvisor.dev/gvisor/pkg/usermem"
	"gvisor.dev/gvisor/pkg/waiter"
)

type socketOpsCommon struct {
	socket.SendReceiveTimeout

	family   int            // Read-only.
	stype    linux.SockType // Read-only.
	protocol int            // Read-only.
	queue    waiter.Queue

	client *UdsClient
	id     int16
}

// Connect implements the connect(2) linux unix.
func (s *socketOpsCommon) Connect(t *kernel.Task, sockaddr []byte, blocking bool) *syserr.Error {
	req := &SocketReq{
		Id:   s.id,
		Size: int16(len(sockaddr)),
		Op:   OpConnect,
	}
	copy(req.Data[:], sockaddr)
	fmt.Printf("%+v\n", req)
	fmt.Println("Request Connect")
	_, err := s.client.Request(req)
	if err != nil {
		return syserr.ErrAborted
	}
	fmt.Println("Request Connect Done")
	return nil
}

// Accept implements the accept4(2) linux unix.
// Returns fd, real peer address length and error. Real peer address
// length is only set if len(peer) > 0.
func (s *socketOpsCommon) Accept(t *kernel.Task, peerRequested bool, flags int, blocking bool) (int32, linux.SockAddr, uint32, *syserr.Error) {
	fmt.Println("Accept not implemented")
	return 0, nil, 0, syserr.ErrNotSupported
}

// Bind implements the bind(2) linux unix.
func (s *socketOpsCommon) Bind(t *kernel.Task, sockaddr []byte) *syserr.Error {
	fmt.Println("Bind not implemented")
	return syserr.ErrNotSupported
}

// Listen implements the listen(2) linux unix.
func (s *socketOpsCommon) Listen(t *kernel.Task, backlog int) *syserr.Error {
	fmt.Println("Listen not implemented")
	return syserr.ErrNotSupported
}

// Shutdown implements the shutdown(2) linux unix.
func (s *socketOpsCommon) Shutdown(t *kernel.Task, how int) *syserr.Error {
	fmt.Println("Shutdown not implemented")
	return syserr.ErrNotSupported
}

// GetSockOpt implements the getsockopt(2) linux unix.
func (s *socketOpsCommon) GetSockOpt(t *kernel.Task, level int, name int, outPtr hostarch.Addr, outLen int) (marshal.Marshallable, *syserr.Error) {
	fmt.Println("GetSockOpt not implemented")
	return nil, syserr.ErrNotSupported
}

// SetSockOpt implements the setsockopt(2) linux unix.
func (s *socketOpsCommon) SetSockOpt(t *kernel.Task, level int, name int, opt []byte) *syserr.Error {
	fmt.Println("SetSockOpt not implemented")
	return syserr.ErrNotSupported
}

// GetSockName implements the getsockname(2) linux unix.
//
// addrLen is the address length to be returned to the application, not
// necessarily the actual length of the address.
func (s *socketOpsCommon) GetSockName(t *kernel.Task) (addr linux.SockAddr, addrLen uint32, err *syserr.Error) {
	fmt.Println("GetSockName not implemented")
	return nil, 0, syserr.ErrNotSupported
}

// GetPeerName implements the getpeername(2) linux unix.
//
// addrLen is the address length to be returned to the application, not
// necessarily the actual length of the address.
func (s *socketOpsCommon) GetPeerName(t *kernel.Task) (addr linux.SockAddr, addrLen uint32, err *syserr.Error) {
	fmt.Println("GetPeerName not implemented")
	return nil, 0, syserr.ErrNotSupported
}

// RecvMsg implements the recvmsg(2) linux unix.
//
// senderAddrLen is the address length to be returned to the application,
// not necessarily the actual length of the address.
//
// flags control how RecvMsg should be completed. msgFlags indicate how
// the RecvMsg call was completed. Note that control message truncation
// may still be required even if the MSG_CTRUNC bit is not set in
// msgFlags. In that case, the caller should set MSG_CTRUNC appropriately.
//
// If err != nil, the recv was not successful.
func (s *socketOpsCommon) RecvMsg(t *kernel.Task, dst usermem.IOSequence, flags int, haveDeadline bool, deadline ktime.Time, senderRequested bool, controlDataLen uint64) (n int, msgFlags int, senderAddr linux.SockAddr, senderAddrLen uint32, controlMessages socket.ControlMessages, err *syserr.Error) {
	fmt.Println("RecvMsg not implemented")
	return 0, 0, nil, 0, socket.ControlMessages{}, syserr.ErrNotSupported
}

// SendMsg implements the sendmsg(2) linux unix. SendMsg does not take
// ownership of the ControlMessage on error.
//
// If n > 0, err will either be nil or an error from t.Block.
func (s *socketOpsCommon) SendMsg(t *kernel.Task, src usermem.IOSequence, to []byte, flags int, haveDeadline bool, deadline ktime.Time, controlMessages socket.ControlMessages) (n int, err *syserr.Error) {
	fmt.Println("SendMsg not implemented")
	return 0, syserr.ErrNotSupported
}

// SetRecvTimeout sets the timeout (in ns) for recv operations. Zero means
// no timeout, and negative means DONTWAIT.
func (s *socketOpsCommon) SetRecvTimeout(nanoseconds int64) {
	fmt.Println("SetRecvTimeout not implemented")
}

// RecvTimeout gets the current timeout (in ns) for recv operations. Zero
// means no timeout, and negative means DONTWAIT.
func (s *socketOpsCommon) RecvTimeout() int64 {
	fmt.Println("RecvTimeout not implemented")
	return 0
}

// SetSendTimeout sets the timeout (in ns) for send operations. Zero means
// no timeout, and negative means DONTWAIT.
func (s *socketOpsCommon) SetSendTimeout(nanoseconds int64) {
	fmt.Println("SetSendTimeout not implemented")
}

// SendTimeout gets the current timeout (in ns) for send operations. Zero
// means no timeout, and negative means DONTWAIT.
func (s *socketOpsCommon) SendTimeout() int64 {
	fmt.Println("SendTimeout not implemented")
	return 0
}

// State returns the current state of the socket, as represented by Linux in
// procfs. The returned state value is protocol-specific.
func (s *socketOpsCommon) State() uint32 {
	fmt.Println("State not implemented")
	return 0
}

// Type returns the family, socket type and protocol of the socket.
func (s *socketOpsCommon) Type() (family int, skType linux.SockType, protocol int) {
	fmt.Println("Type not implemented")
	return 0, 0, 0
}

// EventRegister implements waiter.Waitable.EventRegister.
func (s *socketOpsCommon) EventRegister(e *waiter.Entry, mask waiter.EventMask) {
	fmt.Println("EventRegister not implemented")
}

// EventUnregister implements waiter.Waitable.EventUnregister.
func (s *socketOpsCommon) EventUnregister(e *waiter.Entry) {
	fmt.Println("EventUnregister not implemented")
}

// Ioctl implements fs.FileOperations.Ioctl.
func (s *socketOpsCommon) Ioctl(context.Context, *fs.File, usermem.IO, arch.SyscallArguments) (uintptr, error) {
	fmt.Println("Ioctl not implemented")
	return 0, linuxerr.ENOTTY
}

// Read implements fs.FileOperations.Read.
func (s *socketOperations) Read(ctx context.Context, _ *fs.File, dst usermem.IOSequence, _ int64) (int64, error) {
	fmt.Println("Read not implemented")
	return 0, linuxerr.EPERM
}

// Write implements fs.FileOperations.Write.
func (s *socketOperations) Write(ctx context.Context, _ *fs.File, src usermem.IOSequence, _ int64) (int64, error) {
	fmt.Println("Write not implemented")
	return 0, linuxerr.EPERM
}

// Readiness implements waiter.Waitable.Readiness.
func (s *socketOpsCommon) Readiness(mask waiter.EventMask) waiter.EventMask {
	fmt.Println("Readiness not implemented")
	return 0
}

// Release implements fs.FileOperations.Release.
func (s *socketOpsCommon) Release(ctx context.Context) {
	fmt.Println("Release not implemented")
}

type socketProvider struct {
	family int
}

// socketOperations implements fs.FileOperations and socket.Socket for a socket
// implemented using a dpdk socket.
type socketOperations struct {
	fsutil.FilePipeSeek             `state:"nosave"`
	fsutil.FileNotDirReaddir        `state:"nosave"`
	fsutil.FileNoFsync              `state:"nosave"`
	fsutil.FileNoMMap               `state:"nosave"`
	fsutil.FileNoSplice             `state:"nosave"`
	fsutil.FileNoopFlush            `state:"nosave"`
	fsutil.FileUseInodeUnstableAttr `state:"nosave"`

	socketOpsCommon
}

var _ = socket.Socket(&socketOperations{})

func newSocketFile(ctx context.Context, family int, stype linux.SockType, protocol int, id int16, client *UdsClient, nonblock bool) (*fs.File, *syserr.Error) {
	s := &socketOperations{
		socketOpsCommon: socketOpsCommon{
			family:   family,
			stype:    stype,
			protocol: protocol,
			client:   client,
			id:       id,
		},
	}
	// if err := fdnotifier.AddFD(int32(fd), &s.queue); err != nil {
	// 	return nil, syserr.FromError(err)
	// }
	dirent := socket.NewDirent(ctx, socketDevice)
	defer dirent.DecRef(ctx)
	return fs.NewFile(ctx, dirent, fs.FileFlags{NonBlocking: nonblock, Read: true, Write: true, NonSeekable: true}, s), nil
}

// Socket implements socket.Provider.Socket.
func (p *socketProvider) Socket(t *kernel.Task, stypeflags linux.SockType, protocol int) (*fs.File, *syserr.Error) {
	// Check that we are using the dpdk network stack.
	stack := t.NetworkContext()
	if stack == nil {
		return nil, nil
	}
	if _, ok := stack.(*Stack); !ok {
		return nil, nil
	}
	tStack := stack.(*Stack)
	client := tStack.client
	tStack.idAlloc += 1
	id := tStack.idAlloc
	fmt.Println("new socket from dpdk")
	// Only accept TCP and UDP.
	stype := stypeflags & linux.SOCK_TYPE_MASK
	switch stype {
	case unix.SOCK_STREAM:
		switch protocol {
		case 0, unix.IPPROTO_TCP:
			// ok
		default:
			return nil, nil
		}
	case unix.SOCK_DGRAM:
		switch protocol {
		case 0, unix.IPPROTO_UDP:
			// ok
		default:
			return nil, nil
		}
	default:
		return nil, nil
	}

	req := &SocketReq{
		Id:   id,
		Size: 0,
		Op:   OpCreate,
	}
	_, err := client.Request(req)
	if err != nil {
		return nil, syserr.ErrAborted
	}

	f, e := newSocketFile(t, p.family, stype, protocol, id, client, stypeflags&unix.SOCK_NONBLOCK != 0)
	if e != nil {
		return nil, e
	}

	return f, nil
}

// Pair implements socket.Provider.Pair.
func (p *socketProvider) Pair(*kernel.Task, linux.SockType, int) (*fs.File, *fs.File, *syserr.Error) {
	// Not supported by AF_INET/AF_INET6.
	return nil, nil, nil
}

func init() {
	for _, family := range []int{unix.AF_INET, unix.AF_INET6} {
		socket.RegisterProvider(family, &socketProvider{family})
	}
}
