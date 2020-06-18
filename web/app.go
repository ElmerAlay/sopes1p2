package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person : estructura que almacenará el dato recibido o enviado
type Person struct {
	Name      string `json:"Nombre"`
	Depto     string `json:"Departamento"`
	Age       int    `json:"Edad"`
	Contagion string `json:"Forma de contagio"`
	State     string `json:"Estado"`
}

var persons []Person

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/", addPerson).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "applycation/json")
	json.NewEncoder(w).Encode(persons)
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Something wrong has passed")
	}

	json.Unmarshal(reqBody, &person)
	persons = append(persons, person)

	w.Header().Set("Content-Type", "applycation/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}
