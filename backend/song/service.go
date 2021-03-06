package song

import (
	"fmt"

	"golang.org/x/net/context"
)

// CreateService creates a new instance of Service ready to be used
func CreateService(r *Repo) *Service {
	return &Service{repo: r}
}

// Service is the service layer structure for handling domain logic
type Service struct {
	repo RepoInterface
}

// Add is the service method for handling domain logic
func (s Service) Add(ctx context.Context, req *AddSongRequest) error {
	fmt.Printf(">>> Add Service Method called <<<\n")
	err := s.repo.Add(ctx, req.Song)
	return err
}

// Get is the service method for handling domain logic
func (s Service) Get(ctx context.Context, req *GetSongRequest) (*GetSongResponse, error) {
	fmt.Printf(">>> Get Service Method called <<<\n")
	foundSongs, err := s.repo.GetMulti(ctx, req.Name)
	return &GetSongResponse{
		FoundSongs: foundSongs,
	}, err
}

// Update is the service method for handling domain logic
func (s Service) Update(ctx context.Context, req *UpdateSongRequest) error {
	fmt.Printf(">>> Update Service Method called <<<\n")
	song, err := s.repo.Get(ctx, req.UpdatedSong.SongID)
	if err != nil {
		return err
	}
	err = s.repo.Update(ctx, req.UpdatedSong, song.FilePath)
	if err != nil {
		return err
	}
	return nil
}

// Delete is the service method for handling domain logic
func (s Service) Delete(ctx context.Context, req *DeleteSongRequest) error {
	fmt.Printf(">>> Delete Service Method called <<<\n")
	song, err := s.repo.Get(ctx, req.SongID)
	if err != nil {
		return err
	}
	if song == nil {
		return nil
	}
	err = s.repo.Delete(ctx, req.SongID, song.FilePath)
	if err != nil {
		return err
	}
	return nil
}
