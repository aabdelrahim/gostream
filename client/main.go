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
	// addSong(context.Background(), client, "testfile", "mp3", []string{}, "newSong")
	getSong(context.Background(), client, "newSong", []string{}, "")
}

func addSong(ctx context.Context, client pb.SongServiceClient, filename string, audioFormat string, artists []string, songName string) {
	dat, err := ioutil.ReadFile(filename + "." + audioFormat)
	check(err)
	newSong := &pb.Song{Name: songName, Artists: artists, Audio: dat, Audioformat: audioFormat}
	request := &pb.AddSongRequest{NewSong: newSong}
	fmt.Printf("Adding song with request\rsongName: %s\rartists: %v\n", songName, artists)
	resp, err := client.Add(ctx, request)
	if err != nil {
		fmt.Printf("Error sending grpc request: %v - %v\n", request, err)
	}
	fmt.Printf("Got response: %v", resp)
}

func getSong(ctx context.Context, client pb.SongServiceClient, songName string, artists []string, songID string) {
	request := &pb.GetSongRequest{Name: songName, Artists: artists, SongID: songID}
	fmt.Printf("Getting song with request: %v \n", request)
	resp, err := client.Get(ctx, request)
	check(err)
	for _, v := range resp.Result {
		fmt.Println(v.Song.Name, " has id: ", v.SongID)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
