package primerpaquete

/*
	Las librerias no pueden ser ejecutadas de manera independiente, deben ser ejecutadas por otro archivo
*/

/*
	Cuidado que las funciones que están dentro del paquete escritas con minúscula no se pueden acceder
	desde un paquete externo, porque esa es la convención de Go, las funciones que empiezan con minúscula son privadas
	la que empiecen con mayúscula son públicas.
*/

var h string
var H2 string //Recordemos que debe de ser en mayúscula para que sea pública

func init() {
	h = "Holaaaa!!"
	H2 = ":3 ... Holiii!! :D"
}

func HolaMundo() string {
	return "Hola Mundo"
}

func holaMundo2() string {
	return "Hola Mundo"
}

func HolaMundo3() string {
	return h
}
