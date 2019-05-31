package main

import "fmt"

/*
	LOS CANALES

	Su relación es inmediata con Go Routines (El uso de la KeyWord go sobre funciones o parte de código para que se ejecute
	de manera concurrente) puesto que un canal permite la comunicación entre rutinas o ejecución de código concurrentes,
	vía un canal

*/

func main() {
	//Para la creación de un canal debemos de usar la función make
	channel := make(chan string) //chan indica que creamos un canal y después sigue el tipo de dato que enviaremos entre canales

	//La go routine siempre solicitando info
	go func(channel chan string) { //la declaración de nuestra función debe esperar el channel, que debe ser igual al anterior
		//Ciclo infinito
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name //Como enviar info a un canal, de lado derecho del operador flecha (<-) la imfo, de lado izquierdo el canal
		}
	}(channel) //Enviando el channel

	//El script se detiene vía la espera de info del canal
	msg := <-channel //Recibiendo la info el canal ya aparece de lado derecho y la info (Que esta en una var) del lado izquierdo, esto con respecto al operador flecha

	fmt.Println("Este es el mensaje que se obtiene desde el canal: " + msg) //acá imprimimos eventualmente lo que sale del canal

	//Si lo ejecutamos dos veces, el lo permite. Puesto que espera hasta que el canal envié la info
	msg = <-channel //Recibiendo la info el canal ya aparece de lado derecho y la info (Que esta en una var) del lado izquierdo, esto con respecto al operador flecha

	fmt.Println("Este es el segundo mensjae que se obtiene desde el canal: " + msg) //acá imprimimos eventualmente lo que sale del canal

	/*
		OJO: variable = <-channel ESPERA HASTA QUE SE OBTENGA INFO DEL CANAL
		ASÍ PUES SI CAMBIAMOS LAS LINEAS DE ESPERA DEL CANAL Y DE IMPRESIÓN DEL CANAL POR OTRO FOR INFINITO, ASÍ:

		for {
			msg := <-channel
			fmt.Println("Este es el mensaje que se obtiene desde el canal: " + msg)
		}

		LA EJECUCIÓN FUESE INFINITA
	*/

}
