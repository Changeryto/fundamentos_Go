package main

import "fmt"

const PI float64 = 3.1416

func salutationsFunction(subject string) {
	fmt.Printf("Hola %s \n", subject)
}

func tripleArgument(a, b int32, c float64, d bool) {
	//b y c son del mismo tipo de dato
	fmt.Printf("Hay %b ramas en %b arboles.\nCon %E UFC? %t\n", a, b, c, d)
}

func cilinderVolume(radio, altura float64) {
	volume := PI * radio * radio * altura
	fmt.Printf("El cilindro tiene %e unidades cuadradas\n", volume)
}

func autoReturn(value float64) float64 { //Se especifica lo que se regresa
	return value
}

func simpleYDoble(a float64) (c, d float64) {
	return a, a * 2
}

func main() {
	salutationsFunction("Personas. Uwu")
	tripleArgument(12, 21, 3000.4532, true)
	cilinderVolume(21.1, 50.0)
	fmt.Println(autoReturn(42.01))

	// Funcion anonima
	var num int64 = 64
	doble := func() int64 {
		return num * 2
	}
	fmt.Println(doble())

	// Escapar del segundo return para no declararlo
	v1, v2 := simpleYDoble(90.876)
	fmt.Println("v1 y v1: ", v1, v2)

	v3, _ := simpleYDoble(0.00987)
	fmt.Println(v3)
}
