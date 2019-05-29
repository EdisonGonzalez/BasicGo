package main

import "fmt"

func main() {

	/*
		APUNTES PARA ENTENDER PUNTEROS

		1. Es una dirección de memoria
		2. En lugr del valor de una variable, tenemos la dirección en la que esta almacenada
		3. Supongamos que X,Y => as123d apuntan a la misma dirección de memoria. Y el valor dentro de esa memoria es
		as123d => 5
		Si X => as123d => 6, modifica el valor de la dirección as123d por el valor de 6, entonces
		Y => asd123d => 6, tendrá ahora por valor 6

		Observemos que aunque no modificamos nunca el valor de la variable Y, por estar apuntando a la misma dirección de X
		pòr tricotomia fue alterado

		Esta es una de las utilidades de los punteros, la sintaxis es:

		*T donde T es el tipo de dato, algunos ejemplos son: *int, *string, *Struct
		El valor zero o el valor inicial es nil (null)
	*/

	var puntero1, puntero2 *int
	entero := 5

	//puntero1 = 5 //Esta mal hecho porque estamos indicando que la dirección de memoria es 5
	puntero1 = &entero //& operador para "indicar la dirección de memoria de una variable" no el valor de la variable
	puntero2 = &entero

	fmt.Println(puntero1)
	fmt.Println(puntero2)

	/*
		LO ANTERIOR AL CORRERLO EJECUTA:

		AMAC02W40B5HV2D-EdisonGonzalez:Punteros edison.gonzalez$ go run punteros.go
		0xc00007e008
		0xc00007e008
	*/

	//para acceder al valor de un puntero entonces anteponemos el operador *, a la variable puntero
	fmt.Println(*puntero1)
	fmt.Println(*puntero2)

	//Ahora si modificamos el valor de un soo puntero tenemos
	*puntero1 = 7
	fmt.Println(puntero1)
	fmt.Println(puntero2)
	fmt.Println(*puntero1)
	fmt.Println(*puntero2)

	/*
		SIGUE SIENDO LA MISMA POSICIÓN DE MEMORIA INICIAL AUNQUE HA CAMBIADO EN LAS DIFEREMNTES EJECUCIONES DE ESTE CÓDIGO

		AMAC02W40B5HV2D-EdisonGonzalez:Punteros edison.gonzalez$ go run punteros.go
		0xc000082000
		0xc000082000
		5
		5
		0xc000082000
		0xc000082000
		7
		7

		Pero observemos que el valor cambia
	*/

}
