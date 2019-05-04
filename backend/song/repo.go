package song

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/lib/pq" //adding postgres parsing for database connections
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

	artists := strings.Join(song.Artists, " | ")

	path := "./musicLibrary/" + song.Name + "-" + artists + "." + song.AudioFormat

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
	defer tx.Rollback()

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

// GetMulti returns all songs that match the song name
func (r Repo) GetMulti(ctx context.Context, name string) ([]*Song, error) {
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

// Get returns the song data for the single song with matching songID
func (r Repo) Get(ctx context.Context, songID string) (*Song, error) {
	fmt.Printf(">>> Add Repo Method called <<<\n\n")

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Creating DB transaction failed\n")
		return nil, err
	}
	rows, err := tx.Query(`SELECT * FROM gostream.song WHERE songID =  $1`, songID)
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
	var foundSong *Song
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

		foundSong = &Song{
			Name:        songName,
			Artists:     strings.Split(artists, " | "),
			Audio:       songaudio,
			AudioFormat: audioFormat,
			SongID:      id,
			FilePath:    filePath,
		}
	}
	return foundSong, nil
}

// Update updates the song with matching songID with newSongData in database and file
func (r Repo) Update(ctx context.Context, newSongData *Song, filePath string) error {
	fmt.Printf(">>> Update Repo Method called <<<\n\n")

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Creating DB transaction failed\n")
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE gostream.song
		SET name = $1, artists = $2, audioFormat = $3
		WHERE songID = $4;`,
		newSongData.Name, newSongData.Artists, newSongData.AudioFormat, newSongData.SongID)
	if err != nil {
		fmt.Printf("Updating song data failed: %v\n", err)
		return err
	}

	err = ioutil.WriteFile(filePath, newSongData.Audio, 0644)
	if err != nil {
		fmt.Printf("Failed to write updated song bytes to file: %v\n", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Committing transaction failed: %v\n", err)
		return err
	}
	return nil
}

// Delete removes the song with matching songID from database and removes song file
func (r Repo) Delete(ctx context.Context, songID string, filePath string) error {
	fmt.Printf(">>> Delete Repo Method called <<<\n\n")

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Creating DB transaction failed: %v\n", err)
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM gostream.song
		WHERE songID = $1;`,
		songID)
	if err != nil {
		fmt.Printf("Failed to delete song from database: %v\n", err)
		return err
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Printf("Failed to remove file at: %v - %v\n", filePath, err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Committing transaction failed: %v\n", err)
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
