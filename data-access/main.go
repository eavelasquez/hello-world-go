package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// capture the connection properties
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		AllowNativePasswords: true,
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_NAME"),
	}

	// get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	// ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %s", err)
	}

	fmt.Println("Successfully connected to database!")

	// query for albums by artist - hard-code name "John Coltrane" here.
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatalf("Error querying for albums: %s", err)
	}

	// print the albums
	fmt.Printf("Albums found: %v\n", albums)

	// query for album by ID - hard-code ID 2 here.
	alb, err := albumByID(2)
	if err != nil {
		log.Fatalf("Error querying for album: %s", err)
	}

	// print the album
	fmt.Printf("Album found: %v\n", alb)

	// add an album
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatalf("Error adding album: %s", err)
	}

	// print the ID of the newly inserted album
	fmt.Printf("Album added with ID: %v\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// an albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM albums WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("error querying for albums by %q: %v", name, err)
	}
	defer rows.Close()

	// loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album

		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			return nil, fmt.Errorf("error scanning album row into Album struct: %v", err)
		}
		albums = append(albums, alb)
	}

	// check for errors from iterating over rows.
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return albums, nil
}

// albumByID queries for a the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// an album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM albums WHERE id = ?", id)

	// use Scan to assign column data to struct fields.
	err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("album with ID %d not found", id)
		}
		return alb, fmt.Errorf("error scanning album row into Album struct: %v", err)
	}

	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album's ID.
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("error inserting album into database: %v", err)
	}

	// get the ID of the newly inserted album.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting ID of newly inserted album: %v", err)
	}

	return id, nil
}
