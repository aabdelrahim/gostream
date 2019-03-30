package song

import (
	"fmt"

	pb "github.com/aabdelrahim/gostream/api"
	"golang.org/x/net/context"
)

// CreateServer creates a new instance of the song server ready to be used
func CreateServer(s *Service) *Server {
	return &Server{Service: s}
}

// Server is the Server object used to instantiate a server
type Server struct {
	Service *Service
}

// Add is the API validation method for incoming rpc requests
func (s Server) Add(ctx context.Context, req *pb.AddSongRequest) (*pb.Empty, error) {
	if req == nil {
		fmt.Printf("Request was empty\n")
		return &pb.Empty{}, nil
	}
	fmt.Printf(">>>>> Add Song Request Alert <<<<<\n")

	newSong := req.GetNewSong()

	fmt.Printf("New Song Request: %v\n", &newSong)

	request := &AddSongRequest{
		Song: &Song{
			Name:        req.NewSong.Name,
			Artists:     req.NewSong.Artists,
			AudioFromat: req.NewSong.Audioformat,
		},
	}

	fmt.Printf("Adding song with request: %v\n", request)

	s.Service.Add(ctx, request)

	return &pb.Empty{}, nil
}

// Get is the API validation method for incoming rpc requests
func (s Server) Get(ctx context.Context, req *pb.GetSongRequest) (*pb.GetSongResponse, error) {
	fmt.Printf("Getting song with request: %v\n", &req)
	if req == nil {
		fmt.Printf("Request was empty\n")
		return nil, nil
	}

	fmt.Printf(">>>>> Add Song Request Alert <<<<<\n")

	request := &GetSongRequest{
		Name:    req.GetName(),
		Artists: req.GetArtists(),
		SongID:  req.GetSongID(),
	}

	fmt.Printf("Getting song with request: %v\n", request)

	s.Service.Get(ctx, request)

	return &pb.GetSongResponse{}, nil
}

// Update is the API validation method for incoming rpc requests
func (s Server) Update(ctx context.Context, req *pb.UpdateSongRequest) (*pb.Empty, error) {
	if req == nil {
		fmt.Printf("Request was empty\n")
		return &pb.Empty{}, nil
	}
	fmt.Printf(">>>>> Update Song Request Alert <<<<<\n")

	updatedSong := req.GetUpdatedSong()

	fmt.Printf("New Song Request: %v\n", &updatedSong)

	request := &UpdateSongRequest{
		SongID: req.GetSongID(),
		UpdatedSong: &Song{
			Name:        updatedSong.GetName(),
			Artists:     updatedSong.GetArtists(),
			Audio:       updatedSong.GetAudio(),
			AudioFromat: updatedSong.GetAudioformat(),
		},
	}

	fmt.Printf("Updating song with request: %v\n", request)

	s.Service.Update(ctx, request)

	return &pb.Empty{}, nil
}

// Delete is the API validation method for incoming rpc requests
func (s Server) Delete(ctx context.Context, req *pb.DeleteSongRequest) (*pb.Empty, error) {
	if req == nil {
		fmt.Printf("Request was empty\n")
		return &pb.Empty{}, nil
	}
	fmt.Printf(">>>>> Delete Song Request Alert <<<<<\n")

	request := &DeleteSongRequest{
		SongID: req.GetSongID(),
	}

	fmt.Printf("Deleting song with request: %v\n", request)

	s.Service.Delete(ctx, request)

	return &pb.Empty{}, nil
}
