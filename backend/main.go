package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/aabdelrahim/gostream/api"
	song "github.com/aabdelrahim/gostream/backend/song"
	grpc "google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	s := grpc.NewServer()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("could not listen to port %d: %v", *port, err)
	}
	songRepo := song.CreateRepo()
	songService := song.CreateService(songRepo)
	songServer := song.CreateServer(songService)

	pb.RegisterSongServiceServer(s, songServer)

	fmt.Printf("Preparing to serve requests\n")
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("could not serve %v", err)
	}
}
