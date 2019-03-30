package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"

	pb "github.com/aabdelrahim/gostream/api"
	grpc "google.golang.org/grpc"
)

// Backend is the address of the grpc server
var Backend = flag.String("b", "localhost:8080", "address of the grpc backend")

const sampleRate = 44100
const seconds = 1

func main() {
	conn, err := grpc.Dial(*Backend, grpc.WithInsecure())
	if err != err {
		fmt.Printf("could not dial backend %s: %v", *Backend, err)
	}
	defer conn.Close()
	client := pb.NewSongServiceClient(conn)

	dat, err := ioutil.ReadFile("testfile.mp3")
	check(err)

	newSong := &pb.Song{Name: "NSong", Artists: []string{"A"}, Audio: dat, Audioformat: "mp3"}
	request := &pb.AddSongRequest{NewSong: newSong}
	resp, err := client.Add(context.Background(), request)
	if err != nil {
		fmt.Printf("Error sending grpc request: %v - %v\n", request, err)
	}
	fmt.Printf("Got response: %v", resp)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
