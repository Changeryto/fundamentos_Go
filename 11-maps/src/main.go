package main

import "fmt"

// Objetivo similar a los mapas en Python
func main() {
	m := make(map[string]int) //Se crea un mapa, cuyas llaves con strings y sus valores son ints.

	m["Jose"] = 14
	m["Pepe"] = 24

	// La separacion es un espacio, no una coma.
	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
		// ESTA CLASE DE ITERACION ES EN ORDEN PSEUDOALEATORIO
	}

	// Encontrar un valor
	valorEncontrado := m["Jose"]
	fmt.Println(valorEncontrado)

	// Acceder a una llave inexistente arroja el zero value segun el tipo
	fmt.Println(m["Jair"])

	// Cambiar esto requiere obtener el valor ok

	value, ok := m["Jamon"]
	fmt.Println(value, ok)

	// Tambien pueden declararse directamente
	temps := map[string]float64{
		"Celsius": 273.15,
		"Kelvin":  0.0,
	}

	temp, tempExist := temps["Celsius"]
	fmt.Println(temp, tempExist)
}
