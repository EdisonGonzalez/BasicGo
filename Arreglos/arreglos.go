package main

import "fmt"

//import "fmt"

func main() {
	/*

		DATOS:
			- int, valor por defecto el cero
			- string, valor por defecto es cadena vacia
			- bool, valor por defecto es falso
	*/

	var arreglo [3]int //Arreglo estático, no se puede ejecutar en tiempo de ejecución, se establece desde la compilación
	//Los arreglos de go no tienen variabilidad de datos en la estructura

	fmt.Println(arreglo) //Así al imprimir el arreglo salen por defecto todos en cero

	//Ahora si fuesen cadenas:
	var arreglo_cadenas [5]string
	fmt.Println(arreglo_cadenas)

	arr_ini := [3]int{1, 2} //No necesariamente se deben de inicializar todos los campos
	arr_ini[2] = 20
	fmt.Println(arr_ini)
	//Otra forma de impresión
	for i := 0; i < len(arr_ini); i++ {
		fmt.Println(arr_ini[i])
	}

	//-----------------------------------------------------------------------------------------------------------------
	//Arreglos multidimensionales
	var matriz [3][2]int
	matriz[0][1] = 1
	fmt.Println(matriz)
}
