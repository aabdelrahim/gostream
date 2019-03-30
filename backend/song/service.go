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
	repo *Repo
}

// Add is the service method for handling domain logic
func (s Service) Add(ctx context.Context, req *AddSongRequest) error {
	fmt.Printf(">>> Add Service Method called <<<\n\n")
	err := s.repo.Add(ctx, req.Song)
	return err
}

// Get is the service method for handling domain logic
func (s Service) Get(ctx context.Context, req *GetSongRequest) (*GetSongResponse, error) {
	fmt.Printf(">>> Get Service Method called <<<\n\n")
	return nil, nil
}

// Update is the service method for handling domain logic
func (s Service) Update(ctx context.Context, req *UpdateSongRequest) error {
	fmt.Printf(">>> Update Service Method called <<<\n\n")
	return nil
}

// Delete is the service method for handling domain logic
func (s Service) Delete(ctx context.Context, req *DeleteSongRequest) error {
	fmt.Printf(">>> Delete Service Method called <<<\n\n")
	return nil
}
