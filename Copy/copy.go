package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, 2, 5)
	copy(copia, slice) //copy(dst,src)
	fmt.Println(slice)
	fmt.Println(copia)
	/*
		Lo anterior imprimiría esto:

		[1 2 3 4]
		[1 2]

		Porque sucede esto, debido a que la longitud de copia es de 3
		Es decir, la función de copy coge el minimo de elementos

		Asi para hacer una copia correcta se recomienda que se envie la longitud del arreglo a copiar
	*/

	slice = append(slice, 2, 3, 4)
	copia2 := make([]int, len(slice))
	copy(copia2, slice) //copy(dst,src)
	fmt.Println(slice)
	fmt.Println(copia2)

	//Otro detalle entonces es la modificación de los capacity
	slice = append(slice, 2, 3, 4)
	copia3 := make([]int, len(slice), cap(slice)*2)
	copy(copia3, slice) //copy(dst,src)
	fmt.Println(slice)
	fmt.Println(copia3)
}
