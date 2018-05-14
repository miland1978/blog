// Copyright DI Miliy Andrew 2018 All Rights Reserved.

package main

import (
	"encoding/json"
	"log"
	"time"

	pb "github.com/miland1978/blog/grpc/edu"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCatalogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetFullCatalog(ctx, &pb.GetCatalogRequest{})
	if err != nil {
		log.Fatalf("could not get full catalog: %v", err)
	}
	buf, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		log.Fatalf("cannot convert to JSON: %v", err)
	}
	log.Printf("Catalog:\n%s", buf)
}
