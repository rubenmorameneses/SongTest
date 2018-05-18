package main

import (
	"database/sql"
	"fmt"
)

type Song struct {
	Artist    int    `json:"artist"`
	Name      string `json:"song"`
	Genretype Genre  `json:"genre"`
	Duration  int    `json:"duration"`
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u *Song) getSongByName(db *sql.DB) error {
	statement := fmt.Sprintf("select title,  artist, name, duration from songs inner join genres on (songs.genre=genres.id) where title=%s", u.Name)
	return db.QueryRow(statement).Scan()
}

func getSongsByArtist(db *sql.DB, artist string) ([]Song, error) {
	statement := fmt.Sprintf("SELECT title,  artist, name, duration from songs inner join genres on (songs.genre=genres.id) where title=%s", artist)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	songs := []Song{}
	for rows.Next() {
		var u Song
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		songs = append(songs, u)
	}
	return songs, nil
}
