package main

import (
	"database/sql"
	"fmt"
)

type Song struct {
	Artist    string `json:"artist"`
	Name      string `json:"title"`
	Genretype string `json:"name"`
	Duration  int    `json:"duration"`
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u *Song) getSongByName(db *sql.DB) error {
	statement := fmt.Sprintf("select title,  artist, name, duration from songs inner join genres on (songs.genre=genres.id) where title='%s'", u.Name)
	return db.QueryRow(statement).Scan(&u.Name, &u.Artist, &u.Genretype, &u.Duration)
}

func getSongsByArtist(db *sql.DB, artist string) ([]Song, error) {
	statement := "SELECT title,  artist, name, duration from songs inner join genres on (songs.genre=genres.id) where artist='" + artist + "'"
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	songs := []Song{}
	for rows.Next() {
		var u Song
		if err := rows.Scan(&u.Name, &u.Artist, &u.Genretype, &u.Duration); err != nil {
			return nil, err
		}
		songs = append(songs, u)
	}

	return songs, nil
}

func getSongsByGenre(db *sql.DB, genre string) ([]Song, error) {
	fmt.Println("---genre: ", genre)
	statement := "SELECT title,  artist, name, duration from songs inner join genres on (songs.genre=genres.id) where name='" + genre + "'"
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	songs := []Song{}
	for rows.Next() {
		var u Song
		if err := rows.Scan(&u.Name, &u.Artist, &u.Genretype, &u.Duration); err != nil {
			return nil, err
		}
		songs = append(songs, u)
	}

	return songs, nil
}

func getSongsByDurationRange(db *sql.DB, start, top int) ([]Song, error) {
	statement := fmt.Sprintf("SELECT title,  artist, name, duration FROM songs inner join genres on (songs.genre=genres.id) where duration between %d and %d", start, top)
	rows, err := db.Query(statement)
	//fmt.Println("+++++++++++++++rows: ", rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	songs := []Song{}
	for rows.Next() {
		var u Song
		if err := rows.Scan(&u.Name, &u.Artist, &u.Genretype, &u.Duration); err != nil {
			return nil, err
		}
		songs = append(songs, u)
	}
	return songs, nil
}
