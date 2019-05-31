package main

import (
	"fmt"
	"strings"
	"time"
)

/*
¿GOROUTINES Y SU RELACIÓN CON LA CONCURRENCIA? ¿QUE LE DIFERENCIA DE LOS THREADS DEL SISTEMA?

La concurrencia trata de dividir un problema en diferentes hilos o ejecuciones en paralelo, ejecuciones al mismo tiempo.

Go en lugar de hacer uso de los threads del sistema, hace uso de algo mucho más ligero conocido como Go Routines. Las
ventaja es que en una computadora hay un limite no muy alto de cuántos threads se pueden manejar en cambio con las go
routines se pueden manejar miles y miles de threads o hilos. ¿Como lo hace?, Go solo abre un threads cuando una Go
Routine se esta bloqueando (El programa se bloquee). El compilador se encarga de abrir y cerrar las go routines y
consigo los hilos. Otra ventaja es que no se requiere una función, simplemente a una sección de código
*/

func main() {
	mnlento("Edison")
	fmt.Println("Pero que espera tan berraca! y esto que es secuencial y sincrono tuvo que esperar los 6 segundos anteriores")

	//Ahora de manera Concurrente:
	go mnlento("Edison") //Solo anteponer la palabra Go en significado de concurrencua //Esta no se ejecutó puesto que el sistema no espera
	fmt.Println("Ahora si no tuvo que esperar todo")

	//Condicionemos a que espere hasta que alguien no presione enter
	var wait string
	fmt.Scanln(&wait) //Scanln es similar a Scan, pero detiene el escaneo en una nueva línea y después del elemento final debe haber una nueva línea o EOF.
}

func mnlento(name string) {
	letras := strings.Split(name, "") //Un Split sobre una cadena vacía significa que nos va a separar por carácter

	for _, letra := range letras { //_ es el index ignorado y letra es el termino obtenido del rango de letras
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
