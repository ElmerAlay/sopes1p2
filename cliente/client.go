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

// Sincronizar los hilos
var wg sync.WaitGroup

// url : la ruta a la que deseamos enviar
var url string

func main() {
	// path : ruta del archivo a leer
	var path string

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

	aux := len(persons)
	j := 0

	for i := len(persons); i < amount; i++ {
		persons = append(persons, persons[j])
		j++

		if j == aux {
			j = 0
		}
	}

	cant := amount / threads
	threadsNumber(persons, cant, threads, amount)
}

func threadsNumber(persons []Person, cant int, threads int, amount int) {
	//fmt.Println(cant)
	if threads == 1 {
		mostrar(persons, 0, amount)
	} else if threads == 2 {
		wg.Add(2)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, amount)
		wg.Wait()
	} else if threads == 3 {
		wg.Add(3)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, amount)
		wg.Wait()
	} else if threads == 4 {
		wg.Add(4)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, amount)
		wg.Wait()
	} else if threads == 5 {
		wg.Add(5)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, amount)
		wg.Wait()
	} else if threads == 6 {
		wg.Add(6)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, amount)
		wg.Wait()
	} else if threads == 7 {
		wg.Add(7)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, amount)
		wg.Wait()
	} else if threads == 8 {
		wg.Add(8)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, amount)
		wg.Wait()
	} else if threads == 9 {
		wg.Add(9)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, amount)
		wg.Wait()
	} else if threads == 10 {
		wg.Add(10)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, amount)
		wg.Wait()
	} else if threads == 11 {
		wg.Add(11)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, amount)
		wg.Wait()
	} else if threads == 12 {
		wg.Add(12)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, amount)
		wg.Wait()
	} else if threads == 13 {
		wg.Add(13)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, cant*12)
		go mostrar(persons, cant*12, amount)
		wg.Wait()
	} else if threads == 14 {
		wg.Add(14)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, cant*12)
		go mostrar(persons, cant*12, cant*13)
		go mostrar(persons, cant*13, amount)
		wg.Wait()
	} else if threads == 15 {
		wg.Add(15)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, cant*12)
		go mostrar(persons, cant*12, cant*13)
		go mostrar(persons, cant*13, cant*14)
		go mostrar(persons, cant*14, amount)
		wg.Wait()
	} else if threads == 16 {
		wg.Add(16)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, cant*12)
		go mostrar(persons, cant*12, cant*13)
		go mostrar(persons, cant*13, cant*14)
		go mostrar(persons, cant*14, cant*15)
		go mostrar(persons, cant*15, amount)
		wg.Wait()
	} else if threads == 17 {
		wg.Add(17)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, cant*12)
		go mostrar(persons, cant*12, cant*13)
		go mostrar(persons, cant*13, cant*14)
		go mostrar(persons, cant*14, cant*15)
		go mostrar(persons, cant*15, cant*16)
		go mostrar(persons, cant*16, amount)
		wg.Wait()
	} else if threads == 18 {
		wg.Add(18)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, cant*12)
		go mostrar(persons, cant*12, cant*13)
		go mostrar(persons, cant*13, cant*14)
		go mostrar(persons, cant*14, cant*15)
		go mostrar(persons, cant*15, cant*16)
		go mostrar(persons, cant*16, cant*17)
		go mostrar(persons, cant*17, amount)
		wg.Wait()
	} else if threads == 19 {
		wg.Add(19)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, cant*12)
		go mostrar(persons, cant*12, cant*13)
		go mostrar(persons, cant*13, cant*14)
		go mostrar(persons, cant*14, cant*15)
		go mostrar(persons, cant*15, cant*16)
		go mostrar(persons, cant*16, cant*17)
		go mostrar(persons, cant*17, cant*18)
		go mostrar(persons, cant*18, amount)
		wg.Wait()
	} else if threads == 20 {
		wg.Add(20)
		go mostrar(persons, 0, cant)
		go mostrar(persons, cant, cant*2)
		go mostrar(persons, cant*2, cant*3)
		go mostrar(persons, cant*3, cant*4)
		go mostrar(persons, cant*4, cant*5)
		go mostrar(persons, cant*5, cant*6)
		go mostrar(persons, cant*6, cant*7)
		go mostrar(persons, cant*7, cant*8)
		go mostrar(persons, cant*8, cant*9)
		go mostrar(persons, cant*9, cant*10)
		go mostrar(persons, cant*10, cant*11)
		go mostrar(persons, cant*11, cant*12)
		go mostrar(persons, cant*12, cant*13)
		go mostrar(persons, cant*13, cant*14)
		go mostrar(persons, cant*14, cant*15)
		go mostrar(persons, cant*15, cant*16)
		go mostrar(persons, cant*16, cant*17)
		go mostrar(persons, cant*17, cant*18)
		go mostrar(persons, cant*18, cant*19)
		go mostrar(persons, cant*19, amount)
		wg.Wait()
	}
}

func mostrar(persons []Person, init int, cant int) {
	for i := init; i < cant; i++ {
		// Convertimos a json el objeto en turno
		jsonData, _ := json.Marshal(persons[i])

		// Hacemos la insersión en la ruta que nos dio el usuario y le mandamos el objeto json
		_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("No se ha podido enviar la información: %s\n", err)
		}
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}
