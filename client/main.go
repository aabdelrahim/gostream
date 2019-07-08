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
	// getSong(context.Background(), client, "newSong", []string{}, "")
	updateSong(context.Background(), client, "", "testSong", "testfile", []string{}, "mp3")
}

func addSong(ctx context.Context, client pb.SongServiceClient, songName string, filename string, audioFormat string, artists []string) {
	dat, err := ioutil.ReadFile(filename + "." + audioFormat)
	check(err)
	newSong := &pb.Song{Name: songName, Artists: artists, Audio: dat, Audioformat: audioFormat}
	request := &pb.AddSongRequest{NewSong: newSong}
	fmt.Printf("Adding song with request\nsongName: %s\nartists: %s\n", songName, artists)
	resp, err := client.Add(ctx, request)
	check(err)
	fmt.Printf("Got response: %v\n", resp)
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

func updateSong(ctx context.Context, client pb.SongServiceClient, SongID string, songName string, filename string, artists []string, audioFormat string) {
	dat, err := ioutil.ReadFile(filename + "." + audioFormat)
	check(err)
	updatedSong := &pb.Song{Name: songName, Artists: artists, Audio: dat, Audioformat: audioFormat}
	request := &pb.UpdateSongRequest{SongID: SongID, UpdatedSong: updatedSong}
	fmt.Printf("Updating song with ID: %s to\nname: %s\nartists: %s\n", SongID, songName, artists)
	resp, err := client.Update(ctx, request)
	check(err)
	fmt.Printf("Got response: %v\n", resp)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
