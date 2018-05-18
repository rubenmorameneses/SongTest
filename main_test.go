package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("root", "22245714watata", "test")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}
func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func clearTable() {
	a.DB.Exec("DELETE FROM genres")
	a.DB.Exec("DELETE FROM songs")
	a.DB.Exec("ALTER TABLE genres AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS genres (
  id int,
  name varchar(64)
)
;
INSERT INTO genres (id, name) VALUES
  (1, 'Rock'),
  (2, 'Country'),
  (3, 'Rap'),
  (4, 'Classical'),
  (5, 'Indie Rock'),
  (6, 'Noise Rock'),
  (7, 'Latin Pop Rock'),
  (8, 'Classic Rock'),
  (9, 'Pop')
;

CREATE TABLE IF NOT EXISTS songs (
  artist varchar(1024),
  title varchar(1024),
  genre int,
  duration int
);

INSERT INTO songs (artist, title, genre, duration) VALUES
  ('424', 'Gala', 5, 189),
  ('Colornoise', 'Amalie', 6, 246),
  ('Los Waldners', 'Horacio', 7, 165),
  ('Beatles', 'Strawberry Fields Forever', 8, 245),
  ('Chubby Checker', 'The Twist', 9, 235),
  ('Santana', 'Smooth', 9, 167),
  ('Bobby Darin', 'Mack the Knife', 1, 245),
  ('LeAnn Rhimes', 'How Do I Live', 2, 237),
  ('LMFAO', 'Party Rock Anthem', 3, 189),
  ('The Black Eyed Peas', 'I Gotta Feeling', 3, 219),
  ('Los Del Rio', 'Macarena', 9, 159),
  ('Olivia Newton-John', 'Physical', 9, 195),
  ('Debby Boone', 'You Light Up My Life', 9, 245),
  ('Beatles', 'Hey Jude', 8, 162)
;
`

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/songs", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentSong(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/song/Sinking", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Song not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Song not found'. Got '%s'", m["error"])
	}
}

func TestGetSong(t *testing.T) {
	clearTable()
	addSongs(1)
	req, _ := http.NewRequest("GET", "/song/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func addSongs(count int) {
	if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO Songs VALUES('%s', '%s', %d,  %d)", ("Song " + strconv.Itoa(i+1)), ("Artist " + strconv.Itoa(i+1)), 1, (i + 1))
		a.DB.Exec(statement)
	}
}
