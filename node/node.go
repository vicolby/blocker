package node

import (
	"context"
	"net"
	"sync"

	"github.com/vicolby/blocker/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version    string
	listenAddr string
	logger     *zap.SugaredLogger

	peersLock sync.RWMutex
	peers     map[proto.NodeClient]*proto.Version

	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	logger, _ := zap.NewDevelopment()
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "1.0.0",
		logger:  logger.Sugar(),
	}
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peersLock.Lock()
	defer n.peersLock.Unlock()
	n.logger.Debugw("new peer connected", "addr", v.ListenAddr, "height", v.Height)
	n.peers[c] = v
}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peersLock.Lock()
	defer n.peersLock.Unlock()
	delete(n.peers, c)
}

func (n *Node) BootstrapNetwork(addrs []string) error {
	for _, addr := range addrs {
		c, err := makeNodeClient(addr)
		if err != nil {
			return err
		}

		v, err := c.Handshake(context.Background(), n.getVersion())
		if err != nil {
			n.logger.Errorf("failed to handshake with peer [%s]: %s", addr, err)
			continue
		}

		n.addPeer(c, v)
	}

	return nil
}

func (n *Node) Start(listenAddr string) error {
	n.listenAddr = listenAddr
	var (
		opts       = []grpc.ServerOption{}
		grpcServer = grpc.NewServer(opts...)
	)
	ln, err := net.Listen("tcp", listenAddr)
	n.logger.Debug("node started on port: ", listenAddr)

	if err != nil {
		return err
	}

	proto.RegisterNodeServer(grpcServer, n)
	return grpcServer.Serve(ln)
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	c, err := makeNodeClient(v.ListenAddr)
	if err != nil {
		return nil, err
	}

	n.addPeer(c, v)
	return n.getVersion(), nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	n.logger.Debug("Received transaction from", peer)
	return &proto.Ack{}, nil
}

func (n *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    n.version,
		Height:     100,
		ListenAddr: n.listenAddr,
	}
}

func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	conn, err := grpc.Dial(listenAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return proto.NewNodeClient(conn), nil
}
