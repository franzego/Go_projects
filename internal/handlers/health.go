package handlers

import (
    "encoding/json"
    "net/http"
)

type HealthResponse struct {
    Status string `json: "status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    resp := HealthResponse{Status: "ok"}

    w.Header().Set("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    err := json.NewEncoder(w).Encode(resp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
