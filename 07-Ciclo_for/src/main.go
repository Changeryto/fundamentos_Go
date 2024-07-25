package main

import "fmt"

func main() {
	// Solo existe la palabra reservada for para los bucles.
	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println("Valores de i: ", i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println("Valores de counter: ", counter)
		counter++
	}

	// Ejercicio
	var lesser int8 = 127
	for lesser > 0 {
		fmt.Println("Disminuyendo desde ", lesser)
		lesser--
	}

	// For forever
	var forever int8 = 0
	for {
		fmt.Println("Por siempre ", forever)
		forever++
	}
}
