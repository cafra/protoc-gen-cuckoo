package cuckoo

import (
	"fmt"
	"log"
	"net"

	pb "github.com/carolove/cuckoo/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	lis net.Listener
	s   *grpc.Server
}

func newRPCServer(port string) (*server, error) {
	srv := &server{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return nil, err
	}
	srv.lis = lis
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, srv)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	srv.s = s
	return srv, nil
}

func (srv *server) start() error {
	if err := srv.s.Serve(srv.lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}

func (srv *server) stop() error {
	srv.s.Stop()
	return srv.lis.Close()
}

func (s *server) AppendEntries(ctx context.Context, in *pb.AppendEntriesRequest) (*pb.AppendEntriesAck, error) {
	return nil, nil
}

func (s *server) RequestVote(ctx context.Context, in *pb.RequestVoteRequest) (*pb.RequestVoteAck, error) {
	resp := &pb.RequestVoteAck{}
	resp.Success = true
	resp.Msg = "hello world!@server"
	fmt.Println(in.Vote)
	return resp, nil
}
