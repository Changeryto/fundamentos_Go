package main

import "fmt"

func main() {
	// Println imprime a consola y salta la linea
	const helloMessage string = "Hello"
	const worldMessage string = "World"

	fmt.Println(helloMessage, worldMessage)
	fmt.Print("Sin salto de linea ", helloMessage, worldMessage)

	// Printf aparte de imprimir, agrega una funcion extra al string como valor de entrada.
	nombre := "Platzi"
	const cursos uint16 = 500

	fmt.Printf("%s tiene mas de %d cursos.\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos.\n", nombre, cursos) // Es preferible evitarlo.

	// Sprintf. Permite guardar a una nueva variable.
	message := fmt.Sprintf("%s tiene mas de %d cursos.\n", nombre, cursos)
	fmt.Printf(message)

	// Conocer tipo de dato
	fmt.Printf("helloMessage: %T \n", helloMessage)
}
