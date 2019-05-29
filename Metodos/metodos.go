package main

import "fmt"

type Car struct { //El Keyword type indica que va a definirse un nuevo tipo, el nombre de ese tipo es el que sigue, y struct porque va a ser una estructura
	valor, modelo                      int
	marca, linea, referencia, fundador string
}

/*
	¿Como agregar una función a una estructura?, Sintaxis:

	func (identificaEstructuraDentroDeFunción nombreEstructura) nombreFunción variableRetornar () {
		return string
	}
*/

/*
	func (carro Car) obtener_carro() string {
		return carro.marca + " " + carro.linea + " " + carro.referencia
	}


	LO ANTERIOR ES UNA FORMA PARA ASOCIARLO A POO SE USA THIS

*/
func (this Car) obtener_carro() string {
	return this.marca + " " + this.linea + " " + this.referencia
}

//Pero la estructura también se puede recibir como parametro un puntero en vez de una copia
//Cada vez que pasamos un argumento en Go, este argumento se pasa como una copia
func (this Car) configurar_fundador(f string) {
	this.fundador = f
}

func (this *Car) configurar_fundador_puntero(f string) {
	this.fundador = f
}

func main() {

	/*
		LOS MÉTODOS SON FUNCIONES QUE SE EJECUTAN

		A pesar de Go no ser un lenguaje orientado a objetos puesto que no tiene clases, si permite que nosotros agreguemos
		funciones a estructuras, lo que es muy similar y conveniente

	*/

	car1 := new(Car)
	car1.fundador = "Henry Ford"
	car1.marca = "Ford"
	car1.linea = "Fiesta"
	car1.referencia = "v1 Simple"
	car1.valor = 280000000
	car1.modelo = 2017

	fmt.Println(car1.obtener_carro())

	car1.configurar_fundador("Mr. Volta")
	fmt.Println(car1)

	/*
		$ go run metodos.go
		Ford Fiesta v1 Simple
		&{280000000 2017 Ford Fiesta v1 Simple Henry Ford}

		Observemos que el fundador no cambio de Henry Ford a Mr. Volta, puesto que la función configurar_fundador
		hace una copia de car1 en this y el cambio se hace sobre dicho objeto por eso es importante enviar punteros
		como parametros

		Asi pues si enviamos el puntero debe de cambiar el fundador (func (this *Car) configurar_fundador_puntero)
	*/
	car1.configurar_fundador_puntero("Mr. Volta")
	fmt.Println(car1)

	/*
		Y LA SALIDA ES:

		$ go run metodos.go
		Ford Fiesta v1 Simple
		&{280000000 2017 Ford Fiesta v1 Simple Henry Ford}
		&{280000000 2017 Ford Fiesta v1 Simple Mr. Volta}

		Y si cambia
	*/

	/*
		NOTA: LAS FUNCIONALIDADES QUE SON AGREGADAS A UNA ESTRUCTURA DEBEN ESTAR EN EL MISMO PAQUETE
	*/
}
