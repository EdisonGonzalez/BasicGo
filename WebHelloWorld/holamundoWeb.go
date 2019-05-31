package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	/*
		EXPLICACIÓN DE FUNCIÓN ListenAndServe()

		ENG:
		Listen on the TCP network address ADDR and then calls Serve with HANDLER to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
		The handler is typically nil, in which case the DefaultServeMux is used.
		ListenAndServe always returns a non-nil error.

		ESP:
		Escuche en la dirección de red TCP ADDR y luego llame a Servir con HANDLER para manejar las solicitudes de las conexiones entrantes. Las conexiones aceptadas están configuradas para habilitar TCP keep-alives.
		El controlador suele ser nulo, en cuyo caso se utiliza el DefaultServeMux.
		ListenAndServe siempre devuelve un error no nulo.
	*/

	//El 404 no encuentra urls en el servidor

	//Respondamos a la url / (Home) (Manejador global, la coincidencia de todas las rutas)
	http.HandleFunc("/", handler) //Handle(pattern string, handler Handler)

	http.HandleFunc("/otraruta", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Petición nueva ruta")
		io.WriteString(w, "Eyyyyy Heeeelllow!!") //Esto es lo que aparecerá en la página

		/*
			EN CONSOLA CADA VEZ QUE SE RECARGA LA PÁGINA APARECE. RUTA DIFERENTE A HOME (/otraruta)
			$ go run holamundoWeb.go
			Hay una nueva petición
			Petición nueva ruta

		*/
	})

	http.ListenAndServe(":8000", nil) //Esto ejecuta un servidor y aparece un 404 y debe ser resuelto antes de

}

func handler(w http.ResponseWriter, r *http.Request) { //Esta función recibe dos argumentos y en ese orden para poder responder a una petición web ("Tener esa firma")
	//Los argumentos son así, el mismo compilador lo indica al no ponerlo
	/*

		cannot use handler (type func(http.ResponseWriter, *http.Request)) as type http.Handler in argument to http.Handle:
		func(http.ResponseWriter, *http.Request) does not implement http.Handler (missing ServeHTTP method)go

	*/
	//http.ResponseWriter, es una estructura que nos permite definir como responder a la petición
	//*http.Request, es un puntero a la información de la petición (Que nos envio el navegador)
	fmt.Println("Hay una nueva petición")
	io.WriteString(w, "Hola") //Esto es lo que aparecerá en la página

	/*
		EN CONSOLA CADA VEZ QUE SE RECARGA LA PÁGINA APARECE. RUTA HOME (/) Y CUALQUIER RUTA POR DEFECTO
		$go run holamundoWeb.go
		Hay una nueva petición
		Hay una nueva petición
	*/
}
