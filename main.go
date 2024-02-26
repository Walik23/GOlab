package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type ServerStatus struct {
	Message string `json:"message"`
}

func serverHandle(w http.ResponseWriter, req *http.Request) {
	status := ServerStatus{Message: "Server is running!"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

type TimeResponse struct {
	CurrentTime string `json:"time"`
}

func getTime(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()

	response := TimeResponse{
		CurrentTime: currentTime.Format(time.RFC3339),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/time", getTime)
	http.HandleFunc("/", serverHandle)

	http.ListenAndServe(":8795", nil)
}
