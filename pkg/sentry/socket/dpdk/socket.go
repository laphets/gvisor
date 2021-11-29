package dpdk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	"gvisor.dev/gvisor/pkg/context"

	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/errors/linuxerr"
	"gvisor.dev/gvisor/pkg/fdnotifier"
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

const (
	sizeofInt32 = 4

	// sizeofSockaddr is the size in bytes of the largest sockaddr type
	// supported by this package.
	sizeofSockaddr = unix.SizeofSockaddrInet6 // sizeof(sockaddr_in6) > sizeof(sockaddr_in)

	// maxControlLen is the maximum size of a control message buffer used in a
	// recvmsg or sendmsg unix.
	maxControlLen = 1024
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
	rsp, err := s.client.Request(req)
	if err != nil {
		return syserr.ErrAborted
	}
	if rsp.Result < 0 {
		fmt.Println("Request Connect Error")
		return syserr.ErrAborted
	}
	fmt.Println("Request Connect Done")
	return nil
}

// Accept implements the accept4(2) linux unix.
// The accept() system call is used with connection-based socket, in our case, type is SOCK_STREAM
// Returns fd, Real peer address, real peer address length, and error.
// length is only set if len(peer) > 0.
func (s *socketOpsCommon) Accept(t *kernel.Task, peerRequested bool, flags int, blocking bool) (int32, linux.SockAddr, uint32, *syserr.Error) {

	// update stack idAlloc
	stack := t.NetworkContext()
	tStack := stack.(*Stack)
	tStack.idAlloc += 1
	id := tStack.idAlloc

	req := &SocketReq{
		Id:   s.id, // sockfd
		Size: 0,
		Op:   OpAccept,
	}

	fmt.Printf("current s.id=%d Accept newfd %d\n", s.id, id)
	buf := bytes.NewBuffer(req.Data[:0])
	err := binary.Write(buf, binary.LittleEndian, int32(id))
	if err != nil {
		log.Fatal("encode error:", err)
	}
	fmt.Println("Log: send accept request...\n")
	rsp, err := s.client.Request(req) // receive from SocketRsp
	if err != nil {
		return 0, nil, 0, syserr.ErrAborted
	}

	peerAddrlen := rsp.Size
	peerAddr := rsp.Data[:peerAddrlen]
	fmt.Printf("peerAddrLen %d\n", peerAddrlen)
	fmt.Println("PeerAddr", peerAddr)

	// create a new TCP socket
	stype := linux.SOCK_STREAM & linux.SOCK_TYPE_MASK
	f, e := newSocketFile(t, unix.AF_INET, stype, unix.IPPROTO_TCP, id, s.client, unix.SOCK_STREAM&unix.SOCK_NONBLOCK != 0)
	if e != nil {
		return 0, nil, 0, syserr.ErrAborted
	}

	fd, _ := t.NewFDFrom(0, f, kernel.FDFlags{
		CloseOnExec: flags&unix.SOCK_CLOEXEC != 0,
	})
	t.Kernel().RecordSocket(f)

	fmt.Println("Request Accept Done")
	return fd, socket.UnmarshalSockAddr(s.family, peerAddr), uint32(peerAddrlen), nil
}

// Bind implements the bind(2) linux unix.
func (s *socketOpsCommon) Bind(t *kernel.Task, sockaddr []byte) *syserr.Error {
	req := &SocketReq{
		Id:   s.id, // sockfd
		Size: int16(len(sockaddr)),
		Op:   OpBind,
	}
	// copy sockaddr to req.Data
	copy(req.Data[:], sockaddr)
	rsp, err := s.client.Request(req) // receive from SocketRsp
	if err != nil {
		return syserr.ErrAborted
	}
	if rsp.Result < 0 {
		fmt.Println("Request Bind Error")
		return syserr.ErrAborted
	}
	fmt.Println("Request Bind Done")
	return nil
}

// Listen implements the listen(2) linux unix.
func (s *socketOpsCommon) Listen(t *kernel.Task, backlog int) *syserr.Error {
	req := &SocketReq{
		Id:   s.id, // sockfd
		Size: 0,
		Op:   OpListen,
	}

	buf := bytes.NewBuffer(req.Data[:0])
	err := binary.Write(buf, binary.LittleEndian, int32(backlog))
	if err != nil {
		log.Fatal("encode error:", err)
	}
	fmt.Printf("In Listen: backlog=%d\n", backlog)
	rsp, err := s.client.Request(req) // receive from SocketRsp
	if err != nil {
		return syserr.ErrAborted
	}
	if rsp.Result < 0 {
		return syserr.ErrAborted
	}
	fmt.Println("Request Listen Done")
	return nil
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

	// copy level, name, opt into req.Data
	req := &SocketReq{
		Id:   s.id, // sockfd
		Size: 12,
		Op:   OpSetSockopt,
	}
	buf := bytes.NewBuffer(req.Data[:0])
	err := binary.Write(buf, binary.LittleEndian, int32(level))
	err = binary.Write(buf, binary.LittleEndian, int32(name))
	err = binary.Write(buf, binary.LittleEndian, int32(binary.Size(opt)))
	if err != nil {
		fmt.Print("Error", err)
	}
	copy(req.Data[12:], opt) //SocketReqDataLen-optlen
	fmt.Printf("level=%d, name=%d, optlen=%d\n", level, name, binary.Size(opt))
	// fmt.Print("data after copy, and opt", req.Data, opt)
	rsp, err := s.client.Request(req) // receive from SocketRsp
	if err != nil {
		return syserr.ErrAborted
	}
	if rsp.Result < 0 {
		fmt.Println("Request SetSockOpt Error")
		return syserr.ErrAborted
	}
	fmt.Println("Request SetSockOpt Done")
	return nil
}

// GetSockName implements the getsockname(2) linux unix.
//
// addrLen is the address length to be returned to the application, not
// necessarily the actual length of the address.
func (s *socketOpsCommon) GetSockName(t *kernel.Task) (linux.SockAddr, uint32, *syserr.Error) {
	fmt.Println("Request GetSockName, s.id", s.id)
	req := &SocketReq{
		Id:   s.id,
		Size: 0,
		Op:   OpGetSockName,
	}
	rsp, err := s.client.Request(req)
	if err != nil {
		return nil, 0, syserr.ErrAborted
	}
	addrlen := rsp.Size
	addr := rsp.Data[:addrlen]
	fmt.Printf("addrlen %d\n", addrlen)
	fmt.Println("addr", addr)
	fmt.Println("Request GetSockName Done")
	return socket.UnmarshalSockAddr(s.family, addr), uint32(addrlen), nil
}

// GetPeerName implements the getpeername(2) linux unix.
//
// addrLen is the address length to be returned to the application, not
// necessarily the actual length of the address.
func (s *socketOpsCommon) GetPeerName(t *kernel.Task) (linux.SockAddr, uint32, *syserr.Error) {
	fmt.Println("Request GetPeerName s.id", s.id)
	req := &SocketReq{
		Id:   s.id,
		Size: 0,
		Op:   OpGetPeerName,
	}
	rsp, err := s.client.Request(req)
	if err != nil {
		return nil, 0, syserr.ErrAborted
	}
	addrlen := rsp.Size
	addr := rsp.Data[:addrlen]
	fmt.Printf("addrlen %d\n", addrlen)
	fmt.Println("addr", addr)
	fmt.Println("Request GetPeerName Done")
	return socket.UnmarshalSockAddr(s.family, addr), uint32(addrlen), nil
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
func (s *socketOpsCommon) RecvMsg(t *kernel.Task, dst usermem.IOSequence, flags int, haveDeadline bool, deadline ktime.Time, senderRequested bool, controlDataLen uint64) (int, int, linux.SockAddr, uint32, socket.ControlMessages, *syserr.Error) {
	req := &SocketReq{
		Id:   s.id,
		Size: int16(dst.NumBytes()),
		Op:   OpRecv,
	}
	rsp, err := s.client.Request(req)
	if err != nil || rsp.Result < 0 {
		return 0, 0, nil, 0, socket.ControlMessages{}, syserr.ErrAborted
	}

	if n, err := dst.CopyOut(t, rsp.Data[:rsp.Size]); err != nil || n != int(rsp.Size) {
		return 0, 0, nil, 0, socket.ControlMessages{}, syserr.ErrAborted
	}

	fmt.Println("Request RecvMsg Done")
	return int(rsp.Size), 0, nil, 0, socket.ControlMessages{}, nil
}

// SendMsg implements the sendmsg(2) linux unix. SendMsg does not take
// ownership of the ControlMessage on error.
//
// If n > 0, err will either be nil or an error from t.Block.
func (s *socketOpsCommon) SendMsg(t *kernel.Task, src usermem.IOSequence, to []byte, flags int, haveDeadline bool, deadline ktime.Time, controlMessages socket.ControlMessages) (int, *syserr.Error) {
	req := &SocketReq{
		Id:   s.id,
		Size: int16(src.NumBytes()),
		Op:   OpSend,
	}
	if n, err := src.CopyIn(t, req.Data[:]); err != nil || n != int(src.NumBytes()) {
		return 0, syserr.ErrNotSupported
	}
	rsp, err := s.client.Request(req)
	if err != nil {
		return 0, syserr.ErrAborted
	}
	fmt.Println("Request SendMsg Done")
	return int(rsp.Result), nil
}

// SetRecvTimeout sets the timeout (in ns) for recv operations. Zero means
// no timeout, and negative means DONTWAIT.
func (s *socketOpsCommon) SetRecvTimeout(nanoseconds int64) {
	s.SendReceiveTimeout.SetRecvTimeout(nanoseconds)
	fmt.Println("SetRecvTimeout Done")
}

// RecvTimeout gets the current timeout (in ns) for recv operations. Zero
// means no timeout, and negative means DONTWAIT.
func (s *socketOpsCommon) RecvTimeout() int64 {
	s.SendReceiveTimeout.RecvTimeout()
	fmt.Println("RecvTimeout Done")
	return 0
}

// SetSendTimeout sets the timeout (in ns) for send operations. Zero means
// no timeout, and negative means DONTWAIT.
func (s *socketOpsCommon) SetSendTimeout(nanoseconds int64) {
	s.SendReceiveTimeout.SetSendTimeout(nanoseconds)
	fmt.Println("SetSendTimeout")
}

// SendTimeout gets the current timeout (in ns) for send operations. Zero
// means no timeout, and negative means DONTWAIT.
func (s *socketOpsCommon) SendTimeout() int64 {
	s.SendReceiveTimeout.SendTimeout()
	fmt.Println("SendTimeout")
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
	// s.queue.EventRegister(e, mask)
	// _ = fdnotifier.UpdateFD(int32(s.id))
	fmt.Println("EventRegister Done")
}

// EventUnregister implements waiter.Waitable.EventUnregister.
func (s *socketOpsCommon) EventUnregister(e *waiter.Entry) {
	// s.queue.EventUnregister(e)
	// _ = fdnotifier.UpdateFD(int32(s.id))
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
	fmt.Println("Readiness Done")
	return fdnotifier.NonBlockingPoll(int32(s.id), mask)
}

// Release implements fs.FileOperations.Release.
func (s *socketOpsCommon) Release(ctx context.Context) {
	req := &SocketReq{
		Id:   s.id,
		Size: 0,
		Op:   OpRelease,
	}
	s.client.Request(req)
	fmt.Println("Request Release Done")
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
	fmt.Println("gvisor Socket Create starts")
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
	fmt.Println("new socket from dpdk")
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
