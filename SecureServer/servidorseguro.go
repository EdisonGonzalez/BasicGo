package main

import (
	"net/http"
)

/*
	CONOCIDO COMO SERVIDOR SEGURO, ES LA MISMA APLICACIÓN DE SERVIDOR PERO EN EL CUÁL SE DESEA EVITAR QUE PUEDAN VER
	TODO EL CÓDIGO FUENTE DE NUESTRO SERVIDOR, PUES NOS PUEDEN HACKEAR O TIENE PASSWORDS O EN GENERAL INFORMACIÓN
	CONFIDENCIAL

	La solución planteada es utilizar otra función de http para servir los archivos, en ese caso servimos una carpeta
	pública y en especifico
*/

func main() {
	fileServer := http.FileServer(http.Dir("public")) //http.FileServer requiere un FileSystem root y retorna un Handler.
	//http.Dir requiere un string, donde se especifica una carpeta root en vez de una URL

	http.Handle("/", http.StripPrefix("/", fileServer))
	http.ListenAndServe(":8000", nil)
}
