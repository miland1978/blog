// Copyright DI Miliy Andrew 2018 All Rights Reserved.

//go:generate protoc -I ../edu/ --go_out=plugins=grpc:../edu ../edu/catalog.proto

package main

import (
	"log"
	"net"

	pb "github.com/miland1978/blog/grpc/edu"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetFullCatalog(context.Context, *pb.GetCatalogRequest) (*pb.GetCatalogReply, error) {
	return &pb.GetCatalogReply{}, nil
}

func (s *server) ListCatalog(context.Context, *pb.GetCatalogRequest) (*pb.GetCatalogReply, error) {
	return &pb.GetCatalogReply{}, nil
}

func (s *server) GetCourse(context.Context, *pb.GetCourseRequest) (*pb.GetCourseReply, error) {
	return &pb.GetCourseReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCatalogServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
