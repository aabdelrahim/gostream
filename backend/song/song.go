package song

import (
	"fmt"

	pb "github.com/aabdelrahim/gostream/api"
	"golang.org/x/net/context"
)

type AddSongRequest struct {
	Song Song `json:"song"`
}

type UpdateSongRequest struct {
	SongID      string `json:"songID"`
	UpdatedSong Song   `json:"updatedSong"`
}

type GetSongRequest struct {
	Name    string   `json:"name"`
	Artists []string `json:"artists"`
	SongID  string   `json:"songID"`
}

type GetSongResponse struct {
	FoundSongs []Song `json:"foundSongs"`
}

type DeleteSongRequest struct {
	SongID string `json:"songID"`
}

type Song struct {
	Name    string `json:"name"`
	Artists string `json:"artists"`
	Audio   string `json:"audio"`
	SongID  string `json:"songID"`
}

type SongService struct{}

func (s SongService) Add(ctx context.Context, req *pb.AddSongRequest) (*pb.Empty, error) {
	fmt.Printf("Adding song with request: %v\n", &req)
	return nil, nil
}

func (s SongService) Get(ctx context.Context, req *pb.GetSongRequest) (*pb.GetSongResponse, error) {
	fmt.Printf("Getting song with request: %v\n", &req)
	return nil, nil
}

func (s SongService) Update(ctx context.Context, req *pb.UpdateSongRequest) (*pb.Empty, error) {
	fmt.Printf("Updating song with request: %v\n", &req)
	return nil, nil
}

func (s SongService) Delete(ctx context.Context, req *pb.DeleteSongRequest) (*pb.Empty, error) {
	fmt.Printf("Deleting song with request: %v\n", &req)
	return nil, nil
}
