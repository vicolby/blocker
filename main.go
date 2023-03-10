package main

import (
	"context"
	"log"

	"github.com/vicolby/blocker/node"
	"github.com/vicolby/blocker/proto"
	"google.golang.org/grpc"
)

func main() {
	makeNode("localhost:8080", []string{})
	makeNode("localhost:8081", []string{"localhost:8080"})

	// go func() {
	// 	for {
	// 		time.Sleep(2 * time.Second)
	// 		makeTransaction()
	// 	}
	// }()

	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr)
	if len(bootstrapNodes) > 0 {
		if err := n.BootstrapNetwork(bootstrapNodes); err != nil {
			log.Fatal(err)
		}
	}
	return n
}

func makeTransaction() {
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)
	tx := &proto.Version{
		Version:    "1.0.0",
		Height:     100,
		ListenAddr: "localhost:4000",
	}

	_, err = c.Handshake(context.TODO(), tx)

	if err != nil {
		log.Fatal(err)
	}
}
