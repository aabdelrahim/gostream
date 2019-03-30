package song

import (
	"context"
	"io/ioutil"
)

// CreateRepo creates a new instance of Repo ready to be used
func CreateRepo() *Repo {
	return &Repo{}
}

// Repo is the repo layer used for managing files
type Repo struct{}

// Add creates a new audio file for a given song
func (r Repo) Add(ctx context.Context, song *Song) error {
	err := ioutil.WriteFile(song.Name, song.Audio, 0)
	check(err)
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
