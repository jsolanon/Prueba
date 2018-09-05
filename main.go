package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type song struct {
	id int 'json:"id"'
	artist string 'json:"artist"'
	song string 'json:"song"'
	genre int 'json:"genre"'
	lenght int 'json:"lenght"'
}

type genre struct {
	id int 'json:"id"'
	name string 'json:"name"'
}

type songs []song

var mainDB *sql.DB

func main(){
	db, errOpenDB := sql.Open("sqlite3", "jrdd.db")
	checkErr(errOpenDB)
	mainDB = db

	r := pat.New()
	//r.Del("/songs/:id", http.HandlerFunc(deleteByID))
	r.Get("/songs/:artist", http.HandlerFunc(getByArtist))
	//r.Put("/songs/:id", http.HandlerFunc(updateByID))
	r.Get("/songs", http.HandlerFunc(getBySong))
	//r.Post("/songs", http.HandlerFunc(insert))

}