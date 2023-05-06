package main

import (
	"fmt"
	"net"

	"smtp_server/pkg/api"
	smtp "smtp_server/pkg/smtp"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &smtp.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	fmt.Println("[SERVER STARTED]")
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(l); err != nil {
		fmt.Println(err)
	}
}
