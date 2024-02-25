package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type currentTime struct {
	Time string `json:"time"`
}

func api(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	rfc3339 := currentTime{
		Time: t.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(rfc3339)

	if err != nil {
		panic(err)
	}

	w.Write(result)
}

func main() {
	http.HandleFunc("/time", api)
	panic(http.ListenAndServe(":8795", nil))
}
