package main

import (
	"github.com/EraldCaka/http-rpc-server/types"
	"net/rpc/jsonrpc"
	"testing"
)

func TestIntegration(t *testing.T) {
	client, err := jsonrpc.Dial("tcp", "localhost:4444")
	if err != nil {
		t.Error(err)
	}

	var reply int
	args := types.ArgsSum{Item1: 2, Item2: 3}
	err = client.Call("Server.Sum", args, &reply)
	if err != nil {
		t.Error(err)
	}

	if reply != 5 {
		t.Errorf("Expected sum to be 5, got %d", reply)
	}
}
