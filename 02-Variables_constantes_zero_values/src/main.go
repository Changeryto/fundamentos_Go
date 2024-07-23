package main

import "fmt"

func main() {
	//Declaración de constantes
	const pi float64 = 3.1416
	fmt.Println("El valor de Pi es", pi)

	//Forma alternativa de declarar constantes, sin declarar tipo.
	const pi2 = 3.141624
	fmt.Println("El valor de pi con mas precision es ", pi2)

	// Declaracion de variables.
	// En este caso puede declararse una variable con la notacion de dos puntos.
	// El tipo de variable solo se declara explicitamente cuando se usa var.
	base := 12
	var altura int = 14
	var area int = base * altura

	fmt.Println("Un cuadrado con base, altura y areas ", base, altura, area)

	//Zero values (valor por defecto o default)

	var a int     //0
	var b float64 //+0.000000e+000
	var c string  // ""
	var d bool    //false

	println("Zero values", a, b, c, d)

	// Ejercicio del area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("La base de un cuadrado y su area son", baseCuadrado, areaCuadrado)

}
