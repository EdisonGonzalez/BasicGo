package main

import "fmt"

type User struct { //El Keyword type indica que va a definirse un nuevo tipo, el nombre de ese tipo es el que sigue, y struct porque va a ser una estructura
	edad             int
	nombre, apellido string //nombre string, apellido string
}

func main() {
	/*
		Recordemos que en Go tenemos un montón de variables primitivas como:
		1. Enteros
		2. Strings
		3. Flotantes
		4. Booleanos

		Abstraer en otro tipo de datos como las estructuras, basicamente es la forma en como defino un tipo de dato
	*/

	var edison User
	fmt.Println(edison)
	/*
		LA ESTRUCTURA INICIALIZO LAS VARIABLES INTERNAS EN SUS VALORES POR DEFECTO
		AMAC02W40B5HV2D-EdisonGonzalez:Structs edison.gonzalez$ go run estructuras.go
		{0  }
	*/
	fmt.Println(User{nombre: "David"})
	/*
		Para inicializar la variable puede hacerse:

		gerente := User{nombre:"Tom",apellido:"Blandón"}
		var secre User{edad:28,nombre:"Karla"} //Perop esta no es la ideal
		usuario := User{20,"",""} //Otra forma es declarar sin clave valor pero declarar todos en orden
	*/
	edison = User{24, "David", "González"}
	fmt.Println(edison)

	//Otro aspecto interesante es la existencia del Keyword New, el cual retorna un puntero
	user := new(User)    //user es un puntero no una estructura en este caso
	fmt.Println(user)    //La sálida es: &{0  }
	user.nombre = "Juan" //Se puede acceder a modificar las variables de la estructura a la cual apunta el puntero sin necesidad de ser *
	user.apellido = "Valverde"
	user.edad = 37
	fmt.Println((*user).nombre) //también puede ser user.nombre
	fmt.Println(user)           //Imprime con el & indicando que es un puntero
}
