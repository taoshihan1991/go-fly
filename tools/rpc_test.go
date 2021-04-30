package tools

import (
	"go-fly-muti/frpc"
	"testing"
)

func TestClientRpc(t *testing.T) {
	frpc.ClientRpc()
}
func TestServerRpc(t *testing.T) {
	frpc.NewRpcServer("127.0.0.1:8082")
}
