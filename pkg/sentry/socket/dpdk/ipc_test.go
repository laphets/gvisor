package dpdk

import (
	"fmt"
	"testing"
)

func execute(t *testing.T) error {
	client, err := newUdsClient()
	if err != nil {
		return err
	}
	req := &SocketReq{
		Id:   1,
		Size: 0,
		Op:   OpCreate,
	}
	rsp, err := client.Request(req)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", rsp)
	return nil
}

func TestClient(t *testing.T) {
	if err := execute(t); err != nil {
		t.Fatal(err)
	}
}
