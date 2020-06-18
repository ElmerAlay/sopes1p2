package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	var path string

	// threads : cantidad de hilos a utilizar para enviar
	// amount : cantidad de datos del archivo que queremos enviar
	var threads, amount int

	fmt.Printf("Ingresa la cantidad de hilos que desea utilizar: ")
	fmt.Scanf("%d\n", &threads)
	fmt.Printf("Ingresa la cantidad de datos que desea enviar: ")
	fmt.Scanf("%d\n", &amount)
	fmt.Printf("Ingresa la ruta del archivo: ")
	fmt.Scanf("%s\n", &path)

	readFile(path, threads, amount)
}

func readFile(path string, threads int, amount int) {
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
	threadsNumber(persons, cant, threads)
}

func threadsNumber(persons []Person, cant int, threads int) {
	if threads == 1 {
		mostrar(persons, 0, len(persons))
	} else if threads == 2 {
		wg.Add(2)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, len(persons))
		wg.Wait()
	} else if threads == 3 {
		wg.Add(3)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, len(persons))
		wg.Wait()
	} else if threads == 4 {
		wg.Add(4)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, len(persons))
		wg.Wait()
	} else if threads == 5 {
		wg.Add(5)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, len(persons))
		wg.Wait()
	}
}

func mostrar(persons []Person, init int, cant int) {
	for i := init; i < cant; i++ {
		fmt.Println(persons[i].Nombre)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}
