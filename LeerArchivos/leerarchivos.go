package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	FORMA SENCILLA DE LEER UN ARCHIVO
	Cargar todo el contenido del archivo dentro de la memoria (Cuidado con los archivos grandes)

	SEGUNDA FORMA DE LEER EL ARCHIVO ES CON UNAS LIBRERIAS DEL SISTEMA Y BUFIO
*/

func main() {
	/********************************************* PRIMERA PARTE *********************************************/
	//ReadFile, retorna dos variables. Error (Un ejemplo es que el archivo no se encontró) y El dato leido
	file_data, err := ioutil.ReadFile("./../hola.txt") //Como es una posición relativa, debe ejecutarse go run desde el directorio donde veamos este script de go

	if err != nil {
		fmt.Println("Hubo un error")
		fmt.Println(err)
	} else {
		fmt.Println(string(file_data)) //La infor de file_data esta en bytes. Por eso se hace el cast con string()
	}

	/*
		LA SALIDA DE EJECUCIÓN FUE:

		$ go run leerarchivos.go
			Hola mundo
		------- Aquí acaba -------

		SI CAMBIAMOS LA POSICIÓN RELATIVA DEL ARCHIVO DE ESTA MANERA:
		file_data, err := ioutil.ReadFile("./../../hola.txt")

		LA SALIDA DE EJECUCIÓN SERÍA:
		$ go run leerarchivos.go
		Hubo un error
		open ./../../hola.txt: no such file or directory
	*/

	/********************************************* SEGUNDA PARTE *********************************************/
	archivo, errror := os.Open("./../hola2.txt")

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

	archivo.Close() //Cerrar el archivo
	//Es importante, por ejemplo si el archivo estuviese en una USB para luego que la misma pueda ser extraida
}
