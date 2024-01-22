package main

import (
	"fmt"
	"log"
	"net"

	server "github.com/tovinhtuan/iam_project/cmd/iam-admin-api/server"
	iamgrpc "github.com/tovinhtuan/my_proto/gen_go/iam_service"
	"google.golang.org/grpc"
)

func main() {
	
	lis, err := net.Listen("tcp", "0.0.0.0:9601")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}
	s := grpc.NewServer()
	iamgrpc.RegisterIamServiceServer(s,  &server.Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("error while serve %v", err)
	}
	fmt.Printf("Starting server port 9601...")
}