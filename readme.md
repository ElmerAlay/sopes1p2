# Consumiendo API Rest con GO

_Como consumir una api utilizando GO como cliente_

## Comenzando 🚀

### Pre-requisitos 📋

_ Tener GO instalado _
_Tener Docker instalado (Opcional), sólo si se va a utilizar el servidor de ejemplo_
_En la página oficial están los pasos de instalación_

### Instalación 🔧
_ Abre una terminal en donde se encuentra la carpeta con la aplicación _
_ Entrar en la carpeta "cliente" _
_ Ejecutar el siguiente comando _

```
go run cliente.go
```
_También se puede de la siguiente forma para obtener el ejecutable_

```
go build cliente.go
```
...
./cliente
...

## Ejemplo de uso
_ Cuando ejecutamos la aplicación nos pedirá unos datos que debemos llenar _

...
Ingresa la url a dónde deseeas enviar los datos: http://104.154.87.38:3000/
Ingresa la cantidad de hilos que desea utilizar: 3
Ingresa la cantidad de datos que desea enviar: 20
Ingresa la ruta del archivo: input.json
...

_ En la url colocamos la url que deseamos, o podemos utilizar la que está en la carpeta web_
_ Ingresamos la cantidad de hilos que queremos utilizar para estar mandando la información _
_ La cantidad de datos totales que contiene el archivo de entrada _
_ La ruta del archivo que queremos leer _
_ El programá realiza las acciones necesarias _

## Servidor Web con Docker
_ nos movemos a la carpeta llamada "web" _
_ Allí encontraremos un archivo Dockerfile _
_ corremos el siguiente comando _

...
docker build -t algun_nombre .
...

_ Esto creará una imágen con el servidor listo para ser utilizado en un container _
_ Ahora podemos crear el container _

...
docker run -p miPuerto:3000 --name miNombre nombreImagen
...

_ Y ya podemos entrar al servidor _

## Servidor Web local
_ nos movemos a la carpeta llamada "web" _
_ corremos el siguiente comando _

...
go get -u github.com/gorilla/mux
...

_ Esto instalará un paquete para poder crear el servidor _
_ Ejecutamos la applicación _

...
go run app.go
...

_ O también construir el ejecutable _

...
go build app.go
./app
...

_ Y ya podemos entrar al servidor con localhost _

## Construido con 🛠️

* [visualstudio code](https://code.visualstudio.com/) - El editor de código utilizado
* [Go](https://golang.org/) - El lenguaje utilizado
* [Mux](https://github.com/gorilla/mux) - Para el servidor 
* [Docker](https://www.docker.com/) - Usado para generar contenedores

