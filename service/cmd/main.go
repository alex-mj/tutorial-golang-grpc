package main

import (
	"fmt"
	"log"
	"net"

	api "github.com/alex-mj/simple-grpc-srvc/service/api"
	pb "github.com/alex-mj/simple-grpc-srvc/service/api/proto/math"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, gRPC-world! -=>")

	s := grpc.NewServer()
	pb.RegisterMathServiceServer(s, &api.MathServer{})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" ... gRPC-math-service started:8080")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
