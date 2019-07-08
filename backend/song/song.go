package song

import (
	"fmt"

	pb "github.com/aabdelrahim/gostream/api"
	"github.com/grpc/grpc-go/status"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

// CreateServer creates a new instance of the song server ready to be used
func CreateServer(s *Service) *Server {
	return &Server{Service: s}
}

// Server is the Server object used to instantiate a server
type Server struct {
	Service ServiceInterface
}

// Add is the API validation method for incoming rpc requests
func (s Server) Add(ctx context.Context, req *pb.AddSongRequest) (*pb.Empty, error) {
	if req == nil {
		fmt.Printf("Request was empty\n")
		return &pb.Empty{}, nil
	}
	fmt.Printf(">>> Add Song Request Alert <<<\n")

	newSong := req.GetNewSong()

	request := &AddSongRequest{
		Song: &Song{
			Name:        newSong.Name,
			Artists:     newSong.Artists,
			Audio:       newSong.Audio,
			AudioFormat: newSong.Audioformat,
		},
	}

	err := s.Service.Add(ctx, request)
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

// Get is the API validation method for incoming rpc requests
func (s Server) Get(ctx context.Context, req *pb.GetSongRequest) (*pb.GetSongResponse, error) {
	if req == nil {
		fmt.Printf("Request was empty\n")
		return nil, nil
	}

	fmt.Printf(">>>>> Get Song Request Alert <<<<<\n")

	request := &GetSongRequest{
		Name:    req.GetName(),
		Artists: req.GetArtists(),
		SongID:  req.GetSongID(),
	}

	resp, err := s.Service.Get(ctx, request)
	if err != nil {
		return nil, err
	}
	getSongResults := make([]*pb.GetSongResponse_GetSongResult, len(resp.FoundSongs))

	for k, v := range resp.FoundSongs {
		getSongResults[k] = &pb.GetSongResponse_GetSongResult{
			Song:   songToProto(v),
			SongID: v.SongID,
		}
	}

	return &pb.GetSongResponse{Result: getSongResults}, nil
}

// Update is the API validation method for incoming rpc requests
func (s Server) Update(ctx context.Context, req *pb.UpdateSongRequest) (*pb.Empty, error) {
	if req == nil {
		fmt.Printf("Request was empty\n")
		err := status.Error(codes.FailedPrecondition, "Request was empty")
		return &pb.Empty{}, err
	}

	if req.SongID == "" {
		fmt.Printf("SongID is required to update a song\n")
		err := status.Error(codes.NotFound, "SongID is required to update a song")
		return &pb.Empty{}, err
	}

	fmt.Printf(">>>>> Update Song Request Alert <<<<<\n")

	updatedSong := req.GetUpdatedSong()

	fmt.Printf("New Song Request: %v\n", &updatedSong)

	request := &UpdateSongRequest{
		UpdatedSong: &Song{
			Name:        updatedSong.GetName(),
			Artists:     updatedSong.GetArtists(),
			Audio:       updatedSong.GetAudio(),
			AudioFormat: updatedSong.GetAudioformat(),
			SongID:      req.GetSongID(),
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

	if req.SongID == "" {
		fmt.Printf("SongID is required to update a song\n")
		err := status.Error(codes.NotFound, "SongID is required to update a song")
		return &pb.Empty{}, err
	}

	fmt.Printf(">>>>> Delete Song Request Alert <<<<<\n")

	request := &DeleteSongRequest{
		SongID: req.GetSongID(),
	}

	fmt.Printf("Deleting song with request: %v\n", request)

	s.Service.Delete(ctx, request)

	return &pb.Empty{}, nil
}

func songToProto(in *Song) *pb.Song {
	return &pb.Song{
		Name:        in.Name,
		Audio:       in.Audio,
		Artists:     in.Artists,
		Audioformat: in.AudioFormat,
	}
}
