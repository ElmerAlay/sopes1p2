package main

import (
	"fmt"
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

	router.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
