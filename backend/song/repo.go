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

// Get get audio files for a song
func (r Repo) Get(ctx context.Context, name string) ([]*Song, error) {
	fmt.Printf(">>> Add Repo Method called <<<\n\n")

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Creating DB transaction failed\n")
		return nil, err
	}
	rows, err := tx.Query(`SELECT * FROM gostream.song WHERE name like  '%' || $1 || '%'`, name)
	if err != nil {
		fmt.Printf("Get songs query failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	var id string
	var songName string
	var artists string
	var audioFormat string
	var filePath string
	var foundSongs []*Song
	for rows.Next() {
		err = rows.Scan(&id, &songName, &artists, &audioFormat, &filePath)
		if err != nil {
			fmt.Printf("Error during row scan: %v", err)
			return nil, err
		}
		fmt.Println(id, songName, artists, audioFormat, filePath)
		songaudio, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file at %s: %v", filePath, err)
		}

		foundSongs = append(foundSongs, &Song{
			Name:        songName,
			Artists:     strings.Split(artists, " | "),
			Audio:       songaudio,
			AudioFormat: audioFormat,
			SongID:      id,
		})
	}
	return foundSongs, nil
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
