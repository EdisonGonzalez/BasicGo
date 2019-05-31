package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "./index.html") //"./index.html -> Path donde esta el html a servir (Ruta relativa)
		fmt.Println(r.URL.Path[1:])
		http.ServeFile(w, r, r.URL.Path[1:]) //Puede servir todo lo que esta dentro de la carpeta de donde se lanza
		//r.URL.Path[1:] busca despues del carácter del home / en adelante, los recursos en carpeta y los sirve
		//Incluso en una subcarpeta: http://localhost:8000/nuevaruta/folder.html

	})
	http.ListenAndServe(":8000", nil)
}
