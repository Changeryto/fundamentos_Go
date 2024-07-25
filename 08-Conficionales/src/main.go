package main

import (
	"fmt"
	"log"
	"strconv"
)

func isPair(number uint64) bool {
	var query bool = number%2 == 0
	return query
}

func main() {
	valor1 := 2
	valor2 := 1

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// and

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Cada valor vale su nombre")
	} else if valor1 == 2 && valor2 == 1 {
		fmt.Println("Cada valor no vale lo mismo que su nombre")
	}

	// or
	if valor1 != 0 || valor2 == 2 {
		fmt.Println("Alguno de los dos valores son verdad.")
	}

	// Convertir texto a numero
	// Un valor convertido, y un lugar donde guardar un error de existir.
	value, err := strconv.Atoi("53")
	// Declara lo que se debe hacer si ocurre un error
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("No hay error")
	}
	fmt.Println("Valor convertido", value)

	// Reto par
	var prueba uint64 = 120831298

	if isPair(prueba) {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}

	// Switch case
	// Cuando se requieren multiples condiciones anidadas.
	// Y se va a iterar una variable.
	switch modulo := prueba % 2; modulo {
	case 0:
		fmt.Println("Switch case con un par", prueba)
	default:
		fmt.Println("Switch case con un impar", prueba)
	}

	// Switch sin condicion
	// Si es similar a un if
	thisValue := 0
	switch {
	case thisValue > 100:
		fmt.Println("Es mayor a 0")
	case thisValue < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("El valor es 0")
	}
	// Siempre es buena practica dejar un default
}
