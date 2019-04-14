package song

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

// CreateRepo creates a new instance of Repo ready to be used
func CreateRepo() *Repo {
	return &Repo{}
}

// Repo is the repo layer used for managing files
type Repo struct{}

// ConnectToDatabase creates the initial connection to the databse when the repo is created
func ConnectToDatabase() {
	db, err := sql.Open("postgres", "user=postgres password=docker port=5432 sslmode=disable")
	check(err)
	tx, err := db.Begin()
	check(err)
	defer tx.Commit()
	_, err = tx.Query("CREATE SCHEMA IF NOT EXISTS gostream;")
	transactionCheck(err, tx)
	_, err = tx.Query(`
		CREATE TABLE IF NOT EXISTS gostream.song (
			songID SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			artists TEXT DEFAULT 'unknown',
			audioFormat TEXT NOT NULL,
			filePath TEXT NOT NULL,
			createdOn TIMESTAMP DEFAULT timezone('utc', now())
		);
	`)
	transactionCheck(err, tx)
	fmt.Printf("\n > DATABASE INITIALIZED <\n")
}

// Add creates a new audio file for a given song
func (r Repo) Add(ctx context.Context, song *Song) error {
	fmt.Printf(">>> Add Repo Method called <<<\n\n")
	path := "./musicLibrary/" + song.Name + "." + song.AudioFormat
	err := ioutil.WriteFile(path, song.Audio, 0644)

	// f, err := os.Create("./musicLibrary/songmetadata")
	// check(err)
	// defer f.Close()

	// f.Write([]byte(stringifySong(song, path)))

	check(err)
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func transactionCheck(e error, tx *sql.Tx) {
	if e != nil {
		tx.Rollback()
		panic(e)
	}
}

// func stringifySong(song *Song, path string) string {
// 	resp := song.SongID
// 	resp += "- " + song.Name + "\n"
// 	resp += " path: " + path + "\n"
// 	for _, v := range song.Artists {
// 		resp += " Artist: " + v + "\n"
// 	}

// 	return resp
// }
