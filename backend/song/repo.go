package song

import (
	"context"
	"fmt"
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
	fmt.Printf(">>> Add Repo Method called <<<\n\n")
	err := ioutil.WriteFile("./"+song.Name+song.AudioFormat, song.Audio, 0644)
	check(err)
	return nil
}
q
func check(e error) {
	if e != nil {
		panic(e)
	}
}
