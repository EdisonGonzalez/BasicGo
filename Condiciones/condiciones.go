package main

import (
	"fmt"
)

func main() {

	/*

		LOS OPERADORES MÁS COMUNES SON:

			== igual a
			!= diferente de
			< menor que
			> mayor que
			>= mayor o igual que
			<= menor o igual que
			&& AND (Concatenar condiciones), OJO: Devuelve verdadero si TODAS las condiciones dan verdadero
			|| OR (Concatena expresiones), OJO: Devuelve verdadero si AL MENOS UNA de las condiciones da verdadero
	*/

	if true {
		fmt.Println("La condicion siempre se cumple, Hola Mundo")

		//Los paréntesis son opcionales en la expresión condicional de un if
		//Para el manejo de errores de sintaxis las llaves deben de ir en la misma linea de if ó else
		//Asi se tenga una linea en el if, se necesitan los corchetes {}, de hecho el compilador pone los ; así no sean necesarios en el código

		//EJEMPLO
		x := 20
		y := 20

		if x > y {
			fmt.Printf("%d es mayor que %d\n", x, y)
		} else if y > x {
			fmt.Printf("%d es mayor que %d\n", y, x)
		} else {
			fmt.Printf("%d es igual a %d\n", x, y)
		}
	}
}
