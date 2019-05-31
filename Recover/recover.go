package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Recover.

	Es la forma en como detenemos un Panic.
*/

func main() {
	execReadFile()
	fmt.Println("Nunca me voy a imprimir")
}

func execReadFile() {
	ejecucion := readFile()
	fmt.Printf("Resultado de la ejecución: %t . (Dentro de execReadFile)\n", ejecucion)
}

func readFile() bool {
	archivo, errror := os.Open("./../../../hola22.txt")

	//defer archivo.Close() //Así se puede copiar pero así no vemos un print de que se ejecuto defer
	defer func() {
		archivo.Close()
		fmt.Println("Cerrando el archivo, dentro de func defer")

		r := recover() //Esto lo que hace es detener a Panic()
		fmt.Println(r)
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

		$go run recover.go
		Cerrando el archivo, dentro de func defer
		panic: open ./../../../hola22.txt: no such file or directory

		goroutine 1 [running]:
		main.readFile(0xc00007cf00)
		/Users/edison.gonzalez/Documents/BasicGo/Recover/recover.go:38 +0x35a
		main.execReadFile()
				/Users/edison.gonzalez/Documents/BasicGo/Recover/recover.go:21 +0x26
		main.main()
				/Users/edison.gonzalez/Documents/BasicGo/Recover/recover.go:16 +0x22
		exit status 2

		SE OBSERVA COMO CON EL PANIC SE TRAZA LA RUTA DE TODO EL ERROR.
	*/

	/*
		EN UNA SEGUNDA EJECUCIÓN DESCOMENTANDO LA LÍNEA 33 Y 34, EXPERIMENTAMOS FUNCIONALIDAD RECOVER, EL RESULTADO ES:

		$ go run recover.go
		Cerrando el archivo, dentro de func defer
		open ./../../../hola22.txt: no such file or directory
		Resultado de la ejecución: false . (Dentro de execReadFile)
		Nunca me voy a imprimir

		AQUI SE OBSERVA COMO SE ELIMINA EL PANIC Y SE EJECUTA TANTO LA FUNC execReadFile, main QUE SON DESENCADENADOS
		POR readFile() Y ALLÍ ES DÓNDE SALE EL ERROR
	*/
}
