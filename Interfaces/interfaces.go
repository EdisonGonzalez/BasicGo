package main

import "fmt"

type User interface {
	//No definimos nada del método, solo el nombre y que retorna un entero
	Permisos() int // Puede ser un nivel de 1 a 5
	Nombre() string
}

type Admin struct {
	name string
}

//Métodos de la estructura Admin que implementan la interface User (NOTA: No se requieren KeyWords ni nada por el estilo)
func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.name
}

/*
	Aquí viene la utilidad de las interfaces, no se envia directamente que es de tipo Admin, no es el único tipo de
	usuarios que voy a manejar, pero entonces indico que todos los usuarios que manejo deben de contener lo que contiene
	la interfaz
*/
func auth(user User) string {
	var content string
	switch user.Permisos() {
	case 3:
		content = user.Nombre() + " tiene permisos de editor"
	case 5:
		content = user.Nombre() + " tiene permisos de administrador"
	default:
		content = user.Nombre() + " tiene permisos básicos"
	}
	return content
}

//Ahora creemos otro tipo de usuario:
type Editor struct {
	name string
}

//Métodos de la estructura Editor que implementan la interface User (NOTA: No se requieren KeyWords ni nada por el estilo)
func (this Editor) Permisos() int {
	return 3
}

func (this Editor) Nombre() string {
	return this.name
}

/*

	Una interfaz en Go, la podemos definir como una estructura que define métodos vacíos
	Una interfaz también puede definirse como un tipo de dato que puede pasarse entre funciones o a través del cuál se
	puede crear un arreglo o slice

*/

func main() {

	/*

		admin := Admin{"Edd"}
		fmt.Println(auth(admin)) //Observemos que aunque el metodo auth solicita un tipo User nosotros pasamos un tipo Admin
		//Esto es porque Admin implementa la inferfaz User y esto lo hace completamente válido
		editor := Editor{"Dan"}
		fmt.Println(auth(editor))

	*/

	//TODO EL ANTERIOR CÓDIGO SE PUDO RESUMIR EN:
	usuarios := []User{Admin{"Edd"}, Editor{"Dan"}} //Como implemementar un arreglo de estructuras que corresponden a una interfaz

	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}

	/*
		$ go run interfaces.go
		Edd tiene permisos de administrador
		Dan no tiene permisos de administrador
	*/

}
