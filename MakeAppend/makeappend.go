package main

import "fmt"

func main() {
	slice := make([]int, 3)
	fmt.Println(slice)
	//En este caso el tamaño y la capacidad son la misma
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	//cambiando la capacidad
	slice = make([]int, 3, 5) //Argumentos de make. Primero tipo de slice, luego la longitud, por último la capacidad
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	/*
		Capacidad

		Esta versión slice := make([]int,2) es menos eficiente que slice := make([]int, 2, 4)

		Puesto que internamente cuando un slice desborda su capacidad, la función append se la extiende reconstruyendo dicho Slice
	*/
	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
