package main

import (
	"fmt"
	"github.com/EraldCaka/http-rpc-server/types"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Server struct{}

func (s *Server) Sum(args types.ArgsSum, reply *int) error {
	*reply = args.Item1 + args.Item2
	return nil
}

func main() {
	server := rpc.NewServer()
	err := server.Register(&Server{})
	if err != nil {
		return
	}

	listener, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server has started in port 4444...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
