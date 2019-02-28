package song

import (
	"fmt"

	"golang.org/x/net/context"
)

// Service is the service layer structure for handling domain logic
type Service struct{}

// Add is the service method for handling domain logic
func (s Service) Add(ctx context.Context, req *AddSongRequest) error {
	fmt.Printf(">>> Add Service Method called <<<")
	return nil
}

// Get is the service method for handling domain logic
func (s Service) Get(ctx context.Context, req *GetSongRequest) (*GetSongResponse, error) {
	fmt.Printf(">>> Get Service Method called <<<")
	return nil, nil
}

// Update is the service method for handling domain logic
func (s Service) Update(ctx context.Context, req *UpdateSongRequest) error {
	fmt.Printf(">>> Update Service Method called <<<")
	return nil
}

// Delete is the service method for handling domain logic
func (s Service) Delete(ctx context.Context, req *DeleteSongRequest) error {
	fmt.Printf(">>> Delete Service Method called <<<")
	return nil
}
