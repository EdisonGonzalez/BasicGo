package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22
	//fmt.Println(edad)  //Esta funciona correctamente
	//fmt.Println("Mi edad es: " + edad)  //Aquí presenta un error con el manejo de conversión, no se puede convertir "mi edad es: " a tipo entero
	//Tampoco puede concatenar un número con una cadena
	edad_str := strconv.Itoa(edad)
	fmt.Println("Mi edad es: " + edad_str)

	fmt.Println("\n\nAhora entonces tenemos otra forma\n")

	edad_string := "22"
	edad_int, _ := strconv.Atoi(edad_string) //Una función de go, puede retornar múltiples datos, cuidado
	//Lo anterior retorna el entero convertido del String y un error
	//El operador _ (Guión bajo) es usado para indicar que usted si recibe el resultado pero no va a hacer nada con ello
	fmt.Println(edad_int + 10) //Acá se suman enteros
}
