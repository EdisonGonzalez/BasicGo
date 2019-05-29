package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// -----------------------------------------------------------------------------------------------------------------
	// IMPRIMIR

	edad := 5

	fmt.Println("Mi edad es ", edad)
	fmt.Print("La edad es: \n", edad)
	fmt.Printf("\n\nLa edad es: %d\n\n", edad)
	/*
		%v, maneja el valor de la variable por defecto (Comodin)
		%t, impresión de booleanos
		%f, para flotantes. También tiene %e (Formato cientifico, 1.43e+01). A su vez %b (Formato especial, 8050p-49)
		%s, para el manejo de cadenas
	*/

	// -----------------------------------------------------------------------------------------------------------------
	// LEER DESDE TERMINAL

	/*
		%v, sirve solo para leer Strings (OJO)
	*/

	var edad_sol int
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad_sol)
	fmt.Printf("La edad ingresada es: %d\n", edad_sol)

	var nombre string
	fmt.Println("\nColoca tu nombre: ")
	fmt.Scanf("%v\n", &nombre)      //Esto puede ser también con %s
	fmt.Printf("Hola %v\n", nombre) //Esto puede ser también con %s

	//Bufio ahorra la necesidad de manejo de direcciones de memoria pero requiere de la generación de un reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	text, error := reader.ReadString('\n') //El parametro delimita cuando el lector deja de obtener la cadena, en este caso con un enter o salto de linea
	//Recuerde '' declaran carácter, "" decalaran string
	if error != nil { //Nil significa error
		fmt.Println(error)
	} else {
		fmt.Println("Hola " + text)
	}
}
