package main

import (
	"fmt"

	"./primerpaquete"
)

/*
	Cuando creamos un package la función que se ejcuta inicialmente es init().
	Así como para un programa la ejecución del programa es con main.

	La ventaja de esto es poder pre-configurar todas la variables que requiere nuestro
	paquete para funcionar de forma correcta.
	Init se ejecuta cuando importamos el paquete de manera que cuando usemos las funciones del paquete
	ya todo marche de manera correcta.
*/

func main() {
	fmt.Println(primerpaquete.HolaMundo())
	//fmt.Println(primerpaquete.holaMundo2()) //Error es: cannot refer to unexported name primerpaquete.holaMundo2
	fmt.Println(primerpaquete.HolaMundo3()) //Haciendo referncia a una función pre-analizada
	fmt.Println(primerpaquete.H2)           //Imprimiendo variables o Macros del paquete
}
