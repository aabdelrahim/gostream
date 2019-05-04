package song

import (
	"golang.org/x/net/context"
)

// ServiceInterface is a song service-layer interface
type ServiceInterface interface {
	Add(ctx context.Context, req *AddSongRequest) error
	Get(ctx context.Context, req *GetSongRequest) (*GetSongResponse, error)
	Update(ctx context.Context, req *UpdateSongRequest) error
	Delete(ctx context.Context, req *DeleteSongRequest) error
}

// RepoInterface is a song repo-layer interface
type RepoInterface interface {
	Add(ctx context.Context, song *Song) error
	Get(ctx context.Context, name string) ([]*Song, error)
}

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
	Name        string   `json:"name"`
	Artists     []string `json:"artists"`
	Audio       []byte   `json:"audio"`
	SongID      string   `json:"songID"`
	AudioFormat string   `json:"audioFormat"`
	FilePath    string   `json:"filePath"`
}
