package song

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/lib/pq"
)

// CreateRepo creates a new instance of Repo ready to be used
func CreateRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// Repo is the repo layer used for managing files
type Repo struct {
	db *sql.DB
}

// ConnectToDatabase creates the initial connection to the databse when the repo is created
func ConnectToDatabase() *sql.DB {
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
	return db
}

// Add creates a new audio file for a given song
func (r Repo) Add(ctx context.Context, song *Song) error {
	fmt.Printf(">>> Add Repo Method called <<<\n\n")

	path := "./musicLibrary/" + song.Name + "." + song.AudioFormat

	err := ioutil.WriteFile(path, song.Audio, 0644)
	if err != nil {
		fmt.Printf("Writing Audio file failed\n")
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Creating DB transaction failed\n")
		return err
	}

	artists := strings.Join(song.Artists, " | ")

	query := `INSERT INTO gostream.song (
		name,
		artists,
		audioFormat,
		filePath
	) VALUES (
		$1,
		$2,
		$3,
		$4
	);`
	_, err = tx.Exec(query, song.Name, artists, song.AudioFormat, path)
	if err != nil {
		fmt.Printf("Inserting row failed: %v\n", err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
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
