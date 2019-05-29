package main

import "fmt"

func main() {
	/*

		En Go, solo existe el ciclo for, no se tiene el ciclo while ni demás

	*/
	// Las partes del for: inicialización, condición para continuar el ciclo, instrucción que se ejecuta cada iteración del ciclo
	fmt.Println("Primer Ejemplo")
	for i := 0; i < 10; i = i + 2 {
		fmt.Println(i)
	}

	fmt.Println("Segundo Ejemplo, es el caso de un while")
	z := 0
	for z < 10 {
		fmt.Println(z)
		z++
	}

	fmt.Println("Tercer Ejemplo, es el caso de un ciclo infinito")
	j := 0
	for {
		/*
			if j==2 {
				continue //simplemente se usa para que no realice nada en ese ciclo
			}
		*/

		fmt.Println(j)
		j++

		if j > 10 {
			break //La culminación de un ciclo infinito dado un caso en especifico
		}
	}
}
