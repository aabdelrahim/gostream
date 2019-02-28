package main

import (
	"context"
	"flag"
	"fmt"

	pb "github.com/aabdelrahim/gostream/api"
	grpc "google.golang.org/grpc"
)

// Backend is the address of the grpc server
var Backend = flag.String("b", "localhost:8080", "address of the grpc backend")

func main() {
	conn, err := grpc.Dial(*Backend, grpc.WithInsecure())
	if err != err {
		fmt.Printf("could not dial backend %s: %v", *Backend, err)
	}
	defer conn.Close()
	client := pb.NewSongServiceClient(conn)
	newSong := &pb.Song{Name: "NSong", Artists: []string{"A"}, Audio: nil}
	request := &pb.AddSongRequest{NewSong: newSong}
	resp, err := client.Add(context.Background(), request)
	if err != nil {
		fmt.Printf("Error sending grpc request: %v - %v\n", request, err)
	}
	fmt.Printf("Got response: %v", resp)

}
