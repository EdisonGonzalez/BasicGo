package main

import (
	"fmt"
)

func main() {
	/*

		SLICE

		Estructura construida en base de un arreglo. Los slices son más comunes que los arreglos puesto que estos nos permiten
		redimensionar, es decir, es dinámico

	*/
	// var slice []int //Declaración correcta
	arr := []int{1, 2, 3, 4} //En vez de declarar, inicializar un slice
	fmt.Println(arr)

	//Una diferencia entre un Slice y un Arreglo es que el arreglo si no se inicializa, al ser estático y saber la cantidad
	//de posiciones, los inicializa en cero, vacio o false. En cambio un Slice sin inicializar es nulo
	var array [3]int
	var slice []int
	fmt.Println(array)
	fmt.Println(slice)
	//Tal que se puede obtener:
	if slice == nil {
		fmt.Println("Slice esta vacío")
	}

	if arr == nil {
		fmt.Println("Slice arr vacío")
	} else {
		fmt.Printf("Tamaño de la array: %d\n", len(arr))
	}

	/*
		El Slice maneja:

			- Puntero al arreglo, el slice internamente maneja un arreglo
			- Tiene conocimiento de la longitud del arreglo al que apunta
			- Capacidad

	*/
	arreglo := [3]int{10, 20, 33}
	// A continuación slicing o cortado
	slice2 := arreglo[:2] //Particion desde posición (inicial:final) hasta la final, sin ser la final la inclusiva
	//Esto es entonces entre 0 y 2 sin incluir a 2, tal que se obtiene la posición 0 y 1
	fmt.Println(slice2)
}
