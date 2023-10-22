package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "net/http"
)

//	type contact struct{
//		Name string `json:"name"`
//		Number int `json:"number"`
//	}
var db = make(map[string]int)

func main() {
	fmt.Println("server is starting on port 4000")
	db["ashir"] = 8157987955
	r := mux.NewRouter()
	r.HandleFunc("/contact", getContactHandler).Methods("Get")
	r.HandleFunc("/contact", postContactHandler).Methods("Post")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal(err)
	}
}

// show all contacts in map
func getContactHandler(w http.ResponseWriter, r *http.Request) {
	var response = map[string]interface{}{
		"message": "success",
		"data":    db,
	}
	dbjson, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(dbjson)
}

// add contacts into map
func postContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloooo"))
}
