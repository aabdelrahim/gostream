package song

import (
	"fmt"

	pb "github.com/aabdelrahim/gostream/api"
	"golang.org/x/net/context"
)

// AddSongRequest is a domain representation of the proto messages
type AddSongRequest struct {
	Song *Song `json:"song"`
}

// UpdateSongRequest is a domain representation of the proto messages
type UpdateSongRequest struct {
	SongID      string `json:"songID"`
	UpdatedSong *Song  `json:"updatedSong"`
}

// GetSongRequest is a domain representation of the proto messages
type GetSongRequest struct {
	Name    string   `json:"name"`
	Artists []string `json:"artists"`
	SongID  string   `json:"songID"`
}

// GetSongResponse is a domain representation of the proto messages
type GetSongResponse struct {
	FoundSongs []*Song `json:"foundSongs"`
}

// DeleteSongRequest is a domain representation of the proto messages
type DeleteSongRequest struct {
	SongID string `json:"songID"`
}

// Song is a domain representation of the proto messages
type Song struct {
	Name    string   `json:"name"`
	Artists []string `json:"artists"`
	Audio   []byte   `json:"audio"`
	SongID  string   `json:"songID"`
}

// Service is the service interface used to instantiate a server
type Service struct{}

// Add is the API validation method for incoming rpc requests
func (s Service) Add(ctx context.Context, req *pb.AddSongRequest) (*pb.Empty, error) {
	if req == nil {
		fmt.Printf("Request was empty\n")
		return &pb.Empty{}, nil
	}
	fmt.Printf(">>>>> Add Song Request Alert <<<<<\n")

	newSong := req.GetNewSong()

	fmt.Printf("New Song Request: %v\n", &newSong)

	request := &AddSongRequest{
		Song: &Song{
			Name:    req.NewSong.Name,
			Artists: req.NewSong.Artists,
		},
	}

	fmt.Printf("Adding song with request: %v\n", request)

	return &pb.Empty{}, nil
}

// Get is the API validation method for incoming rpc requests
func (s Service) Get(ctx context.Context, req *pb.GetSongRequest) (*pb.GetSongResponse, error) {
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

	return &pb.GetSongResponse{}, nil
}

// Update is the API validation method for incoming rpc requests
func (s Service) Update(ctx context.Context, req *pb.UpdateSongRequest) (*pb.Empty, error) {
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
			Name:    updatedSong.GetName(),
			Artists: updatedSong.GetArtists(),
			Audio:   updatedSong.GetAudio(),
		},
	}

	fmt.Printf("Updating song with request: %v\n", request)

	return &pb.Empty{}, nil
}

// Delete is the API validation method for incoming rpc requests
func (s Service) Delete(ctx context.Context, req *pb.DeleteSongRequest) (*pb.Empty, error) {
	if req == nil {
		fmt.Printf("Request was empty\n")
		return &pb.Empty{}, nil
	}
	fmt.Printf(">>>>> Delete Song Request Alert <<<<<\n")

	request := &DeleteSongRequest{
		SongID: req.GetSongID(),
	}

	fmt.Printf("Deleting song with request: %v\n", request)

	return &pb.Empty{}, nil
}

// type Service interface {
// 	Add(ctx context.Context, req *pb.AddSongRequest) (*pb.Empty, error)
// 	Get(ctx context.Context, req *pb.GetSongRequest) (*pb.GetSongResponse, error)
// 	Update(ctx context.Context, req *pb.UpdateSongRequest) (*pb.Empty, error)
// 	Delete(ctx context.Context, req *pb.DeleteSongRequest) (*pb.Empty, error)
// }
