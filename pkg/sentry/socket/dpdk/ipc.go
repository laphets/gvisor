package dpdk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type SocketOp int8

const (
	OpCreate SocketOp = iota
	OpBind
	OpListen
	OpAccept
	OpConnect
	OpShutdown
	OpSetSockopt
	OpGetSockopt
	OpRelease
	OpSend
	OpRecv
)

// data plane -> shm
// control plane -> uds

const SocketReqDataLen = 123

type SocketReq struct {
	Id   int16
	Size int16
	Op   SocketOp
	Data [SocketReqDataLen]byte
}

const SocketRspDataLen = 121
type SocketRsp struct {
	Id     int16
	Size   int16
	Op     SocketOp
	Result int16
	Data   [SocketRspDataLen]byte
}

const SockAddr = "/tmp/gvisor.sock"

type SocketChan struct {
}

type UdsClient struct {
	Conn    net.Conn
	chanMap map[int16](chan *SocketRsp)
}

func (c *UdsClient) Request(req *SocketReq) (*SocketRsp, error) {
	if _, ok := c.chanMap[req.Id]; !ok {
		c.chanMap[req.Id] = make(chan *SocketRsp, 10)
	}

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, *req); err != nil {
		return nil, err
	}
	if _, err := c.Conn.Write(buf.Bytes()); err != nil {
		return nil, err
	}
	rsp := <-c.chanMap[req.Id]
	return rsp, nil
}

func (c *UdsClient) dispatch() {
	for {
		rsp := &SocketRsp{}
		if err := binary.Read(c.Conn, binary.LittleEndian, rsp); err != nil {
			fmt.Println("error", err)
			break
		}
		if ch, ok := c.chanMap[rsp.Id]; ok {
			ch <- rsp
		}
	}
}

func newUdsClient() (*UdsClient, error) {
	conn, err := net.Dial("unix", SockAddr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connect to socket server")

	client := &UdsClient{}
	client.Conn = conn
	client.chanMap = make(map[int16]chan *SocketRsp)

	go client.dispatch()

	return client, nil
}
