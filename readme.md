# Consumiendo API Rest con GO

_Como consumir una api utilizando GO como cliente_

## Comenzando 🚀

### Pre-requisitos 📋

_Tener GO instalado_

_Tener Docker instalado (Opcional), sólo si se va a utilizar el servidor de ejemplo_

_En la página oficial están los pasos de instalación_

### Instalación 🔧

_Abre una terminal en donde se encuentra la carpeta con la aplicación_

_Entrar en la carpeta "cliente"_

_Ejecutar el siguiente comando_

```
go run cliente.go
```

_También se puede de la siguiente forma para obtener el ejecutable_

```
go build cliente.go

./cliente
```

## Ejemplo de uso

_Cuando ejecutamos la aplicación nos pedirá unos datos que debemos llenar_

...

Ingresa la url a dónde deseeas enviar los datos: http://104.154.87.38:3000/

Ingresa la cantidad de hilos que desea utilizar: 3

Ingresa la cantidad de datos que desea enviar: 20

Ingresa la ruta del archivo: input.json

...

_En la url colocamos la url que deseamos, o podemos utilizar la que está en la carpeta web_

_Ingresamos la cantidad de hilos que queremos utilizar para estar mandando la información_

_La cantidad de datos totales que contiene el archivo de entrada_

_La ruta del archivo que queremos leer_

_El programá realiza las acciones necesaria _

## Servidor Web con Docker

_nos movemos a la carpeta llamada "web"_

_Allí encontraremos un archivo Dockerfile_

_corremos el siguiente comando _

...
docker build -t algun_nombre .
...

_Esto creará una imágen con el servidor listo para ser utilizado en un container_

_Ahora podemos crear el container_

...
docker run -p miPuerto:3000 --name miNombre nombreImagen
...

_Y ya podemos entrar al servidor_

## Servidor Web local

_nos movemos a la carpeta llamada "web"_

_corremos el siguiente comando_

...
go get -u github.com/gorilla/mux
...

_Esto instalará un paquete para poder crear el servidor_

_Ejecutamos la applicación_

...
go run app.go
...

_O también construir el ejecutable_

...
go build app.go

./app
...

_Y ya podemos entrar al servidor con localhost_

## Construido con 🛠️

* [visualstudio code](https://code.visualstudio.com/) - El editor de código utilizado
* [Go](https://golang.org/) - El lenguaje utilizado
* [Mux](https://github.com/gorilla/mux) - Para el servidor 
* [Docker](https://www.docker.com/) - Usado para generar contenedores

