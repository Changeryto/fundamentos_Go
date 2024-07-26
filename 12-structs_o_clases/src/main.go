package main

import "fmt"

// Creando un struct "car"
type car struct {
	brand string
	year  int
}

func main() {
	// Instanciando
	myCar := car{
		brand: "Volvo",
		year:  1990,
	}

	fmt.Println("Imprimiento struct", myCar)

	// Otra manera de instanciar
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2019
	fmt.Println("Imprimiendo otro struct", otherCar)
}
