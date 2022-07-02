package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Musik struct {
	ID       string    `json:"id"`
	Genre    string    `json:"genre"`
	Durasi   string    `json:"durasi"`
	Title    string    `json:"title"`
	Penyanyi *Penyanyi `json:"penyanyi"`
}

type Penyanyi struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var playlist []Musik

func getAllPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlist)
}

func deletePlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range playlist {
		if item.ID == params["id"] {
			playlist = append(playlist[:index], playlist[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(playlist)
}

func getPlaylistById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range playlist {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func addPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var musik Musik
	_ = json.NewDecoder(r.Body).Decode(&musik)
	musik.ID = strconv.Itoa(rand.Intn(1000000000))
	playlist = append(playlist, musik)
	json.NewEncoder(w).Encode(musik)
}

func updatePlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range playlist {
		if item.ID == params["id"] {
			playlist = append(playlist[:index], playlist[index+1:]...)
			var musik Musik
			_ = json.NewDecoder(r.Body).Decode(&musik)
			musik.ID = params["id"]
			playlist = append(playlist, musik)
			json.NewEncoder(w).Encode(musik)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	playlist = append(playlist, Musik{ID: "1",
		Genre:    "Pop",
		Title:    "Peri Cintaku",
		Durasi:   "4 Menit",
		Penyanyi: &Penyanyi{Firstname: "Ziva", Lastname: "Magnolya"}})

	playlist = append(playlist, Musik{ID: "2",
		Genre:    "Pop",
		Title:    "Melawan Restu",
		Durasi:   "5 Menit",
		Penyanyi: &Penyanyi{Firstname: "Mahalini", Lastname: "Raharja"}})

	r.HandleFunc("/playlist", getAllPlaylist).Methods("GET")
	r.HandleFunc("/playlist/{id}", getPlaylistById).Methods("GET")
	r.HandleFunc("/playlist", addPlaylist).Methods("POST")
	r.HandleFunc("/playlist/{id}", updatePlaylist).Methods("PUT")
	r.HandleFunc("/playlist/{id}", deletePlaylist).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
