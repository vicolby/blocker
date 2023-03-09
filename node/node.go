package node

import (
	"context"
	"fmt"

	"github.com/vicolby/blocker/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version string
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		version: "1.0.0",
	}
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	peer, _ := peer.FromContext(ctx)
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	fmt.Printf("Received handshake from %s, version %s, height %d\n", peer.Addr, v.Version, v.Height)
	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("Received transaction from", peer)
	return &proto.Ack{}, nil
}
