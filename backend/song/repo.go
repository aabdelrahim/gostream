package song

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
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
	path := "./musicLibrary/" + song.Name + "." + song.AudioFormat
	err := ioutil.WriteFile(path, song.Audio, 0644)

	f, err := os.Create("./musicLibrary/songmetadata")
	check(err)
	defer f.Close()

	f.Write([]byte(stringifySong(song, path)))

	check(err)
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringifySong(song *Song, path string) string {
	resp := song.SongID
	resp += "- " + song.Name + "\n"
	resp += " path: " + path + "\n"
	for _, v := range song.Artists {
		resp += " Artist: " + v + "\n"
	}

	return resp
}
