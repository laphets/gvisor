// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package raw

import (
	"io"

	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
	"gvisor.dev/gvisor/pkg/waiter"
)

// createOnlyEndpoint can be created, but all interactions have no effect or
// return errors.
//
// +stateify savable
type createOnlyEndpoint struct {
	tcpip.DefaultSocketOptionsHandler
	ops tcpip.SocketOptions
}

func newCreateOnlyEndpoint(stk *stack.Stack) tcpip.Endpoint {
	// ep.ops must be in a valid state for callers of ep.SocketOptions.
	var ep createOnlyEndpoint
	ep.ops.InitHandler(&ep, stk, tcpip.GetStackSendBufferLimits, tcpip.GetStackReceiveBufferLimits)
	ep.ops.SetSendBufferSize(32*1024, false /* notify */)
	ep.ops.SetReceiveBufferSize(32*1024, false /* notify */)
	return &ep
}

// Abort implements stack.TransportEndpoint.Abort.
func (*createOnlyEndpoint) Abort() {
	// No-op.
}

// Close implements tcpip.Endpoint.Close.
func (*createOnlyEndpoint) Close() {
	// No-op.
}

// ModerateRecvBuf implements tcpip.Endpoint.ModerateRecvBuf.
func (*createOnlyEndpoint) ModerateRecvBuf(int) {
	// No-op.
}

func (*createOnlyEndpoint) SetOwner(owner tcpip.PacketOwner) {
	// No-op.
}

// Read implements tcpip.Endpoint.Read.
func (*createOnlyEndpoint) Read(dst io.Writer, opts tcpip.ReadOptions) (tcpip.ReadResult, tcpip.Error) {
	return tcpip.ReadResult{}, &tcpip.ErrNotPermitted{}
}

// Write implements tcpip.Endpoint.Write.
func (*createOnlyEndpoint) Write(p tcpip.Payloader, opts tcpip.WriteOptions) (int64, tcpip.Error) {
	return 0, &tcpip.ErrNotPermitted{}
}

// Disconnect implements tcpip.Endpoint.Disconnect.
func (*createOnlyEndpoint) Disconnect() tcpip.Error {
	return &tcpip.ErrNotSupported{}
}

// Connect implements tcpip.Endpoint.Connect.
func (*createOnlyEndpoint) Connect(addr tcpip.FullAddress) tcpip.Error {
	return &tcpip.ErrNotPermitted{}
}

// Shutdown implements tcpip.Endpoint.Shutdown. It's a noop for raw sockets.
func (*createOnlyEndpoint) Shutdown(tcpip.ShutdownFlags) tcpip.Error {
	return &tcpip.ErrNotPermitted{}
}

// Listen implements tcpip.Endpoint.Listen.
func (*createOnlyEndpoint) Listen(int) tcpip.Error {
	return &tcpip.ErrNotSupported{}
}

// Accept implements tcpip.Endpoint.Accept.
func (*createOnlyEndpoint) Accept(*tcpip.FullAddress) (tcpip.Endpoint, *waiter.Queue, tcpip.Error) {
	return nil, nil, &tcpip.ErrNotSupported{}
}

// Bind implements tcpip.Endpoint.Bind.
func (*createOnlyEndpoint) Bind(addr tcpip.FullAddress) tcpip.Error {
	return &tcpip.ErrNotPermitted{}
}

// GetLocalAddress implements tcpip.Endpoint.GetLocalAddress.
func (*createOnlyEndpoint) GetLocalAddress() (tcpip.FullAddress, tcpip.Error) {
	return tcpip.FullAddress{}, &tcpip.ErrNotSupported{}
}

// GetRemoteAddress implements tcpip.Endpoint.GetRemoteAddress.
func (*createOnlyEndpoint) GetRemoteAddress() (tcpip.FullAddress, tcpip.Error) {
	return tcpip.FullAddress{}, &tcpip.ErrNotConnected{}
}

// Readiness implements tcpip.Endpoint.Readiness.
func (*createOnlyEndpoint) Readiness(mask waiter.EventMask) waiter.EventMask {
	return waiter.EventMask(0)
}

// SetSockOpt implements tcpip.Endpoint.SetSockOpt.
func (*createOnlyEndpoint) SetSockOpt(opt tcpip.SettableSocketOption) tcpip.Error {
	return &tcpip.ErrUnknownProtocolOption{}
}

func (*createOnlyEndpoint) SetSockOptInt(tcpip.SockOptInt, int) tcpip.Error {
	return &tcpip.ErrUnknownProtocolOption{}
}

// GetSockOpt implements tcpip.Endpoint.GetSockOpt.
func (*createOnlyEndpoint) GetSockOpt(tcpip.GettableSocketOption) tcpip.Error {
	return &tcpip.ErrUnknownProtocolOption{}
}

// GetSockOptInt implements tcpip.Endpoint.GetSockOptInt.
func (*createOnlyEndpoint) GetSockOptInt(opt tcpip.SockOptInt) (int, tcpip.Error) {
	return -1, &tcpip.ErrUnknownProtocolOption{}
}

// HandlePacket implements stack.RawTransportEndpoint.HandlePacket.
func (*createOnlyEndpoint) HandlePacket(pkt *stack.PacketBuffer) {
	panic("unreachable: createOnlyEndpoint should never be registered")
}

// State implements socket.Socket.State.
func (*createOnlyEndpoint) State() uint32 {
	return 0
}

// Wait implements stack.TransportEndpoint.Wait.
func (*createOnlyEndpoint) Wait() {
	// No-op.
}

// LastError implements tcpip.Endpoint.LastError.
func (*createOnlyEndpoint) LastError() tcpip.Error {
	return nil
}

// SocketOptions implements tcpip.Endpoint.SocketOptions.
func (ep *createOnlyEndpoint) SocketOptions() *tcpip.SocketOptions {
	return &ep.ops
}

// Info implements tcpip.Endpoint.Info.
func (*createOnlyEndpoint) Info() tcpip.EndpointInfo {
	return &stack.TransportEndpointInfo{}
}

// Stats returns a pointer to the endpoint stats.
func (*createOnlyEndpoint) Stats() tcpip.EndpointStats {
	return &tcpip.TransportEndpointStats{}
}
