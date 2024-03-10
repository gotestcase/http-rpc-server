package main

import (
	"fmt"
	"github.com/EraldCaka/http-rpc-server/types"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:4444")
	if err != nil {
		panic(err)
	}

	var reply int
	args := types.ArgsSum{Item1: 2, Item2: 3}
	err = client.Call("Server.Sum", args, &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sum:", reply)
}
