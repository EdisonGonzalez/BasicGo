package main

import "fmt"

func main() {
	// var x, y, z int //Go no permite que se declaren variables y no sean usadas
	/* La variable se puede inicializar en una sola linea, es decir, lo siguiente se resume a:
	var x int
	x = 23
	Así no se debe de declarar el tipo de variable
	*/
	//x := 23
	//Ahora para una cadena seria como:
	nombre := "Coco"
	nombre = "Edd" //Siempre que la variable ya este declarada no se puede hacer uso de := sino de = solo
	fmt.Println(nombre)
}
