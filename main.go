package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}

var contactsDB = []Contact{
	// {ID: 1, Name: "Ameer", Number: 99987878},
}

func main() {
	fmt.Println("server is starting on port 4000")

	r := mux.NewRouter()
	r.HandleFunc("/contact", getContactHandler).Methods("Get")
	r.HandleFunc("/contact", postContactHandler).Methods("Post")
	r.HandleFunc("/contact/{id}", patchContactHandler).Methods("PATCH")
	// r.HandleFunc("/conatct/:id", deleteContactHandler).Methods("Delete")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal(err)
	}
}

// show all contacts in map
func getContactHandler(w http.ResponseWriter, r *http.Request) {
	var response = map[string]interface{}{
		"message": "success",
		"data":    contactsDB,
	}
	dbjson, err := json.MarshalIndent(response, "", "\t")
	// use newencoder for later
	if err != nil {
		log.Fatal(err)
	}
	w.Write(dbjson)
}

// add contacts into map
func postContactHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	var reqjsondata Contact

	err := json.Unmarshal(body, &reqjsondata)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reqjsondata)

	length := len(contactsDB)
	newContact := Contact{ID: length + 1, Name: reqjsondata.Name, Number: reqjsondata.Number}
	contactsDB = append(contactsDB, newContact)

	var response = map[string]interface{}{
		"message": "success",
		"data":    contactsDB,
	}
	dbjson, err := json.Marshal(response)

	w.Write(dbjson)
}

// update a contact
func patchContactHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(contactsDB[1])
	body, _ := io.ReadAll(r.Body)
	var reqjsondata Contact

	err = json.Unmarshal(body, &reqjsondata)
	if err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	if num >= 0 && num < len(contactsDB) {
		contactsDB[num+1] = Contact{
			ID:     reqjsondata.ID,
			Name:   reqjsondata.Name,
			Number: reqjsondata.Number,
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Contact Not Found", http.StatusNotFound)
	}

}

// // delete a contact
// func deleteContactHandler(w http.ResponseWriter, r *http.Request) {

// }
