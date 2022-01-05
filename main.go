package main

import (
    "errors"
    "fmt"
    "log"
    "net/http"
)

type User struct {
    email string
    password  string
}

func personCreate(w http.ResponseWriter, r *http.Request) {
    var u User

    err := decodeJSONBody(w, r, &u)
    if err != nil {
        var mr *malformedRequest
        if errors.As(err, &mr) {
            http.Error(w, mr.msg, mr.status)
        } else {
            log.Println(err.Error())
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
        return
    }

    fmt.Fprintf(w, "User: %+v", u)

    
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/person/create", personCreate)

    log.Println("Starting server on :4000...")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}