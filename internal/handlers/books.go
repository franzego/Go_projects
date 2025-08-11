package handlers

import (
    "net/http"
    "encoding/json"
)


type Book struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
}

func BookHandler(w http.ResponseWriter, r *http.Request) {

    book := []Book {
        {ID: 1, Title: "No Longer Human", Author: "Osamu Dazai"},
        {ID: 2, Title: "Crime and Punishment", Author: "Fyodor Dostoevsky"},
    }

    w.Header().Set("content-type", "application/json")
 //   w,Header(http.StatusOK)
     err := json.NewEncoder(w).Encode(book)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
