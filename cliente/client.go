package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Person estructura que almacenará los datos de la persona que contiene el archivo
type Person struct {
	Nombre        string `json:"Nombre"`
	Departamento  string `json:"Departamento"`
	Edad          int    `json:"Edad"`
	FormaContagio string `json:"Forma de contagio"`
	Estado        string `json:"Estado"`
}

var wg sync.WaitGroup
var num int

func main() {
	// path : ruta del archivo a leer
	// url : la ruta a la que deseamos enviar
	var path, url string

	// threads : cantidad de hilos a utilizar para enviar
	// amount : cantidad de datos del archivo que queremos enviar
	var threads, amount int

	fmt.Printf("Ingresa la url a dónde deseeas enviar los datos: ")
	fmt.Scanf("%s\n", &url)
	fmt.Printf("Ingresa la cantidad de hilos que desea utilizar: ")
	fmt.Scanf("%d\n", &threads)
	fmt.Printf("Ingresa la cantidad de datos que desea enviar: ")
	fmt.Scanf("%d\n", &amount)
	fmt.Printf("Ingresa la ruta del archivo: ")
	fmt.Scanf("%s\n", &path)

	readFile(path, threads, amount, url)
}

func readFile(path string, threads int, amount int, url string) {
	// content : aquí guardaremos el contenido del archivo
	content, err := ioutil.ReadFile("input.json")

	// si existe algún error, entraremos en modo pánico, con defer podemos recuperarnos e imprimir dónde ocurrió el error
	defer func() {
		fmt.Println(recover())
	}()

	if err != nil {
		fmt.Println("Error al leer el archivo")
		panic(err)
	}

	// persons : arreglo de ls estructura persona, aquí se almacenarán todos los datos leídos del archivo
	var persons []Person

	// Convertimos cada objeto del archivo json a la estructura Person y la almacenamos en el arreglo
	err2 := json.Unmarshal(content, &persons)
	if err2 != nil {
		fmt.Println("Error al decodificar el json")
		panic(err2)
	}

	cant := amount / threads
	threadsNumber(persons, cant, threads, url)
}

func threadsNumber(persons []Person, cant int, threads int, url string) {
	if threads == 1 {
		mostrar(persons, 0, len(persons), url)
	} else if threads == 2 {
		wg.Add(2)
		go mostrar(persons, 0, cant, url)
		go mostrar(persons, cant, len(persons), url)
		wg.Wait()
	} else if threads == 3 {
		wg.Add(3)
		go mostrar(persons, 0, cant, url)
		go mostrar(persons, cant, cant*2, url)
		go mostrar(persons, cant*2, len(persons), url)
		wg.Wait()
	} else if threads == 4 {
		wg.Add(4)
		go mostrar(persons, 0, cant, url)
		go mostrar(persons, cant, cant*2, url)
		go mostrar(persons, cant*2, cant*3, url)
		go mostrar(persons, cant*3, len(persons), url)
		wg.Wait()
	} else if threads == 5 {
		wg.Add(5)
		go mostrar(persons, 0, cant, url)
		go mostrar(persons, cant, cant*2, url)
		go mostrar(persons, cant*2, cant*3, url)
		go mostrar(persons, cant*3, cant*4, url)
		go mostrar(persons, cant*4, len(persons), url)
		wg.Wait()
	}
}

func mostrar(persons []Person, init int, cant int, url string) {
	for i := init; i < cant; i++ {
		// Convertimos a json el objeto en turno
		jsonData, _ := json.Marshal(persons[i])

		// Hacemos la insersión en la ruta que nos dio el usuario y le mandamos el objeto json
		response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("No se ha podido enviar la información: %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(data[0:0])
		}
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}
