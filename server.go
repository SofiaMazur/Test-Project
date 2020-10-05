package main

import (
	"encoding/json"
	"net/http"
	"time"
	"log"
)

func main() {

	type Time struct {
		Time string `json:"time"`
	}

	http.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
		time := Time{time.Now().Format(time.RFC3339)}
		res, _ := json.Marshal(time)

		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(200)
		rw.Write(res)
	})

	log.Println("Staring HTTP server...")
	log.Fatal(http.ListenAndServe(":8795", nil))
}