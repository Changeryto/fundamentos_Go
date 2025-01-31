package main

import (
	"fmt"
	"strings"
)

func esPalindromo(prueba string) bool {
	var reverso string

	for i := len(prueba) - 1; i >= 0; i-- {
		reverso += string(prueba[i])
	}
	var resultado bool = strings.ToLower(prueba) == strings.ToLower(reverso)
	return resultado
}

func main() {
	// Array
	// Se inicializa por si solo en zeros
	// De tamanio inmutable, mas eficiente
	var array [4]int // Array de 4 enteros
	fmt.Println(array, "Elementos", len(array), "Capacidad", cap(array))
	array[1] = 123
	array[3] = 9128
	fmt.Println(array, "Elementos", len(array), "Capacidad", cap(array))

	// Slices
	// Declaracion similar, no se le indica la cantidad de valores que va a tener.
	slice := []int{0, 1, 2, 3, 4, 3, 2, 1, 0}
	fmt.Println(slice, "Elementos", len(slice), "Capacidad", cap(slice))

	// Metodos en el slice
	fmt.Println(slice[0])   // Elemento 0
	fmt.Println(slice[:3])  // Elemento 0 a 2, ultimo elemento exclusivo
	fmt.Println(slice[2:4]) // Elemento 2 al 3, ultimo elemento exclusivo
	fmt.Println(slice[4:])  // Elemento 4 al ultimo

	// Append. Sintaxis parecida a R!
	slice = append(slice, 7)

	// Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // Debe mapearse para que cada elemento se agregue de forma independiente.
	fmt.Println(slice)

	// Append invertido
	slice = append(newSlice, slice...)
	fmt.Println(slice)

	// Recorriendo los arreglos
	stringsSlice := []string{"Jola", "k", "aze", "uwu"}

	for i, valor := range stringsSlice {
		fmt.Printf("Indice %d valor %s\n", i, valor)
	}

	// Si solo se busca uno de los dos valores
	for _, valor := range stringsSlice {
		fmt.Println("Valor", valor)
	}

	// Es palindromo?
	const prueba string = "Amor a roma"
	fmt.Println("Es palindromo", prueba, esPalindromo(prueba))

	// El paquete strings tiene mas funciones para manejo de strings.
}
