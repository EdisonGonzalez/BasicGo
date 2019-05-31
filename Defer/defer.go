package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Defer, agrega la función o la ejecución de este código al Stack
	El Stack es una serie de funciones que se deben de ejecutar cuando la función retorne

	Defer suele usar cuando requerimos que se cierre un archivo, un enlace a una base de datos, etc.
*/

func main() {
	ejecucion := readFile()
	fmt.Printf("Resultado de la ejecución: %t . (Dentro de main)\n", ejecucion)
}

func readFile() bool {
	archivo, errror := os.Open("./../hola2.txt")

	//defer archivo.Close() //Así se puede copiar pero así no vemos un print de que se ejecuto defer
	defer func() {
		archivo.Close()
		fmt.Println("Cerrando el archivo, dentro de func defer")
	}()

	if errror != nil {
		fmt.Println("Hubo un error")
	}

	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() { //Cada vez que se llama el método Scan, este coge una linea de texto
		i++
		linea := scanner.Text()
		fmt.Printf("%d. ", i)
		fmt.Println(linea)
	}

	if true {
		fmt.Println("Dummy")
		return true
	}

	//Parte del código que nunca se ejecuta por lo anterior
	//Pero si se ejecuta la función Defer
	fmt.Println("Nunca me ejecutó")
	return true
}
