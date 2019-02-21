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
	fmt.Printf("Adding song with request: %v\n", &req)
	return nil, nil
}

// Get is the API validation method for incoming rpc requests
func (s Service) Get(ctx context.Context, req *pb.GetSongRequest) (*pb.GetSongResponse, error) {
	fmt.Printf("Getting song with request: %v\n", &req)
	return nil, nil
}

// Update is the API validation method for incoming rpc requests
func (s Service) Update(ctx context.Context, req *pb.UpdateSongRequest) (*pb.Empty, error) {
	fmt.Printf("Updating song with request: %v\n", &req)
	return nil, nil
}

// Delete is the API validation method for incoming rpc requests
func (s Service) Delete(ctx context.Context, req *pb.DeleteSongRequest) (*pb.Empty, error) {
	fmt.Printf("Deleting song with request: %v\n", &req)
	return nil, nil
}
