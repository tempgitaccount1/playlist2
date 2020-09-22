package server

import (
	"context"
	"fmt"
	"net"

	pb "playlist/proto"

	"google.golang.org/grpc"
)

type Playlist struct {
}

func Test() {
	fmt.Printf("huh")
}

func New() *Playlist {
	return &Playlist{}
}

func (p *Playlist) Svc() *pb.PlaylistManagerService {
	return &pb.PlaylistManagerService{}
}

func (p *Playlist) Start() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterPlaylistManagerService(server, p.Svc())
	server.Serve(lis)
	return nil
}

func (p *Playlist) NewPlaylist(ctx context.Context, r *pb.NewPlaylistRequest) (*pb.NewPlaylistReply, error) {
	return &pb.NewPlaylistReply{
		Message: fmt.Sprintf("new playlist %v", r.Name),
	}, nil
}
