package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	// https://sourceforge.net/projects/mingw-w64/
)
//Como definir un objeto, defino la estructura
type Song struct {
	id int `json:"id"`
	artist string `json:"artist"`
	song string `json:"song"`
	//genre int 'json:"genre"'
	lenght int `json:"lenght"`
	genre *Genre `json:"genre"`
}

type Genre struct {
	id int `json:"id"`
	name string `json:"name"`
}

type songs []Song

//type genres []Genre

var SongsDB *sql.DB

func main(){
	db, errOpenDB := sql.Open("sqlite3", "jrdd.db")
	checkErr(errOpenDB)
	SongsDB = db

	r := pat.New()
	//r.Del("/songs/:id", http.HandlerFunc(deleteByID))
	r.Get("/songs/:artist", http.HandlerFunc(getByArtist))
	//r.Put("/songs/:id", http.HandlerFunc(updateByID))
	r.Get("/songs/:song", http.HandlerFunc(getBySong))
	//r.Post("/songs", http.HandlerFunc(insert))

	http.Handle("/", r)

	//log.Print(" Running on 12345")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func getByArtist(w http.ResponseWriter, r *http.Request) {
	artist := r.URL.Query().Get(":artist")
	stmt, err := SongsDB.Prepare(" SELECT S.artist, S.song, G.name, S.length FROM Songs S JOIN Genres G	ON S.genre = G.id WHERE S.artist = ?")
	checkErr(err)
	rows, errQuery := stmt.Query(artist)
	checkErr(errQuery)
	var song Song
	for rows.Next() {
		err = rows.Scan(&song.artist, &song.song, &song.genre, &song.lenght)
		checkErr(err)
	}
	jsonB, errMarshal := json.Marshal(song)
	checkErr(errMarshal)
	fmt.Fprintf(w, "%s", string(jsonB))
}

func getBySong(w http.ResponseWriter, r *http.Request) {
	song := r.URL.Query().Get(":song")
	stmt, err := SongsDB.Prepare(" SELECT S.artist, S.song, G.name, S.length FROM Songs S JOIN Genres G	ON S.genre = G.id WHERE S.song = ?")
	checkErr(err)
	rows, errQuery := stmt.Query(song)
	checkErr(errQuery)
	var song Song
	for rows.Next() {
		err = rows.Scan(&song.artist, &song.song, &song.genre, &song.lenght)
		checkErr(err)
	}
	jsonB, errMarshal := json.Marshal(song)
	checkErr(errMarshal)
	fmt.Fprintf(w, "%s", string(jsonB))
}