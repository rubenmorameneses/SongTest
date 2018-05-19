package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// answers whatever output in jsons format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getSongsArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	art := vars["artist"]
	fmt.Println("artist: " + art)
	songs, err := getSongsByArtist(a.DB, art)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Song not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, songs)
}

func (a *App) getSongsByGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gen := vars["genre"]
	fmt.Println("gen: " + gen)
	songs, err := getSongsByGenre(a.DB, gen)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Song not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, songs)
}

func (a *App) getSongByTittle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tittle := vars["tittle"]
	fmt.Println("tittle: " + tittle)
	u := Song{Name: tittle}
	err := u.getSongByName(a.DB)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Song not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// means that exist!
	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/songsByArtist/{artist:[a-zA-Z0-9]+}", a.getSongsArtist).Methods("GET")
	a.Router.HandleFunc("/song/{tittle:[a-zA-Z0-9]+}", a.getSongByTittle).Methods("GET")
	a.Router.HandleFunc("/songByGenre/{genre:[a-zA-Z0-9]+}", a.getSongsByGenre).Methods("GET")
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
