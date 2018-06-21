package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/resty.v1"
)

func main() {
	r := makeRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func makeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/citibike", getCitiBikes).Methods("GET")
	return r

}

func getCitiBikes(w http.ResponseWriter, r *http.Request) {
	resp, _ := resty.R().Get("https://gbfs.citibikenyc.com/gbfs/en/station_status.json")

	var stationStatus stationStatus
	if err := json.Unmarshal(resp.Body(), &stationStatus); err != nil {
		log.Fatal("Client unmarshal failed: " + err.Error())
	}

	json.NewEncoder(w).Encode(&stationStatus)
}
