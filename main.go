package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type User struct {
    Email string
    Password  string
}

func personCreate(w http.ResponseWriter, r *http.Request) {
    // Declare a new Person struct.
    var u User

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
    	fmt.Println("fail" )
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


    // Do something with the Person struct...
    fmt.Println("succ" )
    w.Header().Add("Content-Type", "application/json")
    if u.Email == "donny" {
    	w.WriteHeader(http.StatusCreated)
    }
    else{
    	http.Error(w, err.Error(), http.StatusBadRequest)
    }

    data, err := json.Marshal(u)
    w.Write(data)

}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/person/create", personCreate)

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}