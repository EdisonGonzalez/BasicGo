package main

import (
	"encoding/json"
	"net/http"
)

type Video struct {
	Fuente      string `json:"fuente"` //Convención de Go, los atributos que inician en letra mayúscula son públicos
	Titulo      string `json:"titulo"`
	Propietario string `json:"propietario"`
	Duracion    string `json:"duracion"`
	Topico      string `json:"tema_video"`
	minuscula   string //`json:"llave_especial"` -> struct field minuscula has json tag but is not exported -> //Un carácter que inicia en minúscula, no es público
}

type Videos []Video

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		video := Video{"YouTube", "Que es Blockchain", "Canal Vlog", "5:32", "Tecnología", "Ejemplo no se muestra"}
		json.NewEncoder(w).Encode(video)

		videos := Videos{
			Video{"YouTube", "Que es Machine Learning", "Canal DotCSV", "7:32", "Tecnología", "Ejemplo no se muestra"},
			Video{"YouTube", "Que es 4.0 Revolución Industrial", "Canal Edd", "3:46", "Tecnología", "Ejemplo no se muestra"},
			Video{"YouTube", "Que es IoT", "Canal Bit", "2:13", "Tecnología", "Ejemplo no se muestra"},
		}
		json.NewEncoder(w).Encode(videos)
	})
	http.ListenAndServe(":8001", nil)

	//ANTES DE PONER `json:...` PARA QUE SALGA EN EL PIPE DEL INDEX
	/*
		La salida al consultar http://localhost:8001/ es:

		{"Fuente":"YouTube","Titulo":"Que es Blockchain","Propietario":"Canal Vlog","Duracion":"5:32","Topico":"Tecnología"}

		Observemos que no sale el camppo: "minuscula":"Ejemplo no se muestra"
	*/

	//ANTES DE PONER `json:...` PARA QUE SE ELIJA COMO SALIR EN FORMATO JSON, DESPUÉS
	/*
		La salida al consultar http://localhost:8001/ es:

		{"fuente":"YouTube","titulo":"Que es Blockchain","propietario":"Canal Vlog","duracion":"5:32","tema_video":"Tecnología"}
	*/
}
