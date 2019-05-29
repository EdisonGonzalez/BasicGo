package main

import "fmt"

/*
	Anonymous fields o campos anonimos

	Son campos para las estructuras que nos permiten replicar el concepto de la herencia de la POO (Programación
	Orientada a objetos)

	Recordemos que Go no es un lenguaje de POO
*/

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "Muchachos, el día de hoy en campos anónimos"
}

func (this Human) hablar2() string {
	return "Muchachos, el día de hoy en campos anónimos (Versión 2)"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (this Tutor) hablar2() string {
	return this.Human.hablar2() + "  ->  Agregamos la sobre-escritura"
}

func main() {
	tutor := Tutor{Human{"Edd"}, Dummy{"Ot"}}
	//fmt.Println(tutor.name) //Ambiguous selector //Si solo existe un campo anonimo con ese nombre,
	//en este caso si name existe en Human y en Dummy, se debe especificar cual campo
	fmt.Println(tutor.Human.name)
	/*
		Y esto imprime el nombre de tutor, sin Tutor como estructura tener este campo

		$ go run camposanonimos.go
		Edd
	*/

	//Pero puedo usar funciones o atributos asociados con una estructura directamente desde Tutor si no se repite
	fmt.Println(tutor.hablar()) //Ya no hay ambiguedad ni necesidad de hacer tutor.Human.hablar()
	fmt.Println(tutor.Human.hablar2())
	//En la herencia se pueden reescribir los metodos, un ejemplo es:
	fmt.Println(tutor.hablar2()) //Hace lo que hace el método sobreescrito

	/*

		A ESTE PUNTO LA IMPRESIÓN ES:

		$go run camposanonimos.go
		Edd
		Muchachos, el día de hoy en campos anónimos
		Muchachos, el día de hoy en campos anónimos (Versión 2)
		Muchachos, el día de hoy en campos anónimos (Versión 2)  ->  Agregamos la sobre-escritura
	*/

}
