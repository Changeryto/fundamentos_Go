package main

import "fmt"

func main() {
	// Suma
	x := 10
	y := 98

	result := x + y

	fmt.Println("Suma:", result)

	// Resta
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplicaci
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division (resultado)
	result = x / y
	fmt.Println("Division:", result)

	// Modulo o residuo (los flotantes no permiten obtener modulo)
	result = x % y
	fmt.Println("Modulo:", result)

	// Incremental
	i := 0
	i++
	fmt.Println("Incremental:", i)

	// Decremental
	i--
	fmt.Println("Decremental", i)

	// Calcular el area de un...

	// Rectangulo
	ladoA := 10
	ladoB := 14
	rectangulo := ladoA * ladoB
	fmt.Println("Rectangulo", rectangulo)

	// Trapecio
	baseMayor := 21
	baseMenor := 10
	altura := 7

	trapecio := baseMenor*altura + (((baseMayor-baseMenor)/2)*altura)/2

	fmt.Println("Trapecio:", trapecio)

	//Circulo

	const pi float64 = 3.1416
	radio := 10.0
	circulo := pi * radio * radio

	println("Circulo:", circulo)
}
