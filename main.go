package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type currentTime struct {
	Time string `json:"time"`
}

func main() {
	http.HandleFunc("/time", api)

	server := &http.Server{
		Addr: ":8795",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)

	rfc3339 := currentTime{
		Time: t,
	}

	w.Header().Set("Content-Type", "application/json")

	result, err := json.Marshal(rfc3339)
	if err != nil {
		panic(err)
	}
	w.Write(result)
}
