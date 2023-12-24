package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}
type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	fmt.Println("Server is on port 8080")
	http.ListenAndServe(":8080", nil)
}
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var reqData JsonRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}
	if reqData.Message == "" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}
	fmt.Println("Received message:", reqData.Message)
	resp := JsonResponse{
		Status:  "success",
		Message: "Data successfully received",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
