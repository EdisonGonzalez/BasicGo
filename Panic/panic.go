package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Panic.

	En el ejemplo de hoy, vamos a colocar un error, en este caso en la ruta del archivo para que no lo pueda abrir.

	Puesto que panic es una forma especial en la que nosotros podemos imprimir un error e identificar en que linea
	esta pasando. Panic hace un trazado de la ruta del error
*/

func main() {
	ejecucion := readFile()
	fmt.Printf("Resultado de la ejecución: %t . (Dentro de main)\n", ejecucion)
}

func readFile() bool {
	archivo, errror := os.Open("./../../../hola22.txt")

	//defer archivo.Close() //Así se puede copiar pero así no vemos un print de que se ejecuto defer
	defer func() {
		archivo.Close()
		fmt.Println("Cerrando el archivo, dentro de func defer")
	}()

	if errror != nil {
		panic(errror)
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

	/*
		EL RESULTADO DE EJECUCIÓN ES:

		$go run panic.go
		Cerrando el archivo, dentro de func defer
		panic: open ./../../../hola22.txt: no such file or directory

		goroutine 1 [running]:
		main.readFile(0x10b1900)
				/Users/edison.gonzalez/Documents/BasicGo/Panic/panic.go:33 +0x35a
		main.main()
				/Users/edison.gonzalez/Documents/BasicGo/Panic/panic.go:19 +0x26
		exit status 2

		SE OBSERVA QUE EL PRINTF DEL MAIN NO SE EJECUTÓ, PERO SI SE EJECUTE EL DEFER, DEBIDO A QUE SIEMPRE SE EJECUTA
		Y ES LO ÚNICO.
	*/
}
