package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/vicolby/blocker/node"
	"github.com/vicolby/blocker/proto"
	"google.golang.org/grpc"
)

func main() {
	node := node.NewNode()

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)

	}
	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("Node is running on port 8080")

	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()

	grpcServer.Serve(ln)
}

func makeTransaction() {
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	c := proto.NewNodeClient(client)

	tx := &proto.Version{
		Version: "1.0.0",
	}

	_, err = c.Handshake(context.TODO(), tx)

	if err != nil {
		log.Fatal(err)
	}
}
