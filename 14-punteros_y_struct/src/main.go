package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (thisPC pc) ping() {
	fmt.Println(thisPC.brand, "Pong")
}

// Observa que es thisPC de clase puntero a pc
func (thisPC *pc) halfRAM() {
	thisPC.ram = (thisPC.ram / 2)
}

func main() {
	a := 50
	// Apuntar a la direccion de memoria
	b := &a
	// Deslocalizacion (acceder al valor al que se apunta)
	c := *b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Si se modifica el valor al que se apunta con una direccion de memoria, tambien se modifica el valor con su variable original.

	*b = 100
	fmt.Println(a)

	// La desreferencia deja de apuntar a la memoria.
	c = 200
	fmt.Println(a)
	fmt.Println(c)

	// LOS PUNTEROS SON MUY USADOS PARA CREAR FUNCIONES,
	// PERO SOBRE TODO PARA LLEVAR FUNCIONES DE UNA LIBRERIA
	// A OTRO PAQUETE MUCHO MAS EFICIENTE.

	myPC := pc{ram: 16, disk: 200, brand: "Lenovo"}
	fmt.Println(myPC)

	myPC.ping()

	// Ejemplo con una funcion (observa linea 15)
	fmt.Println("Original:", myPC)
	myPC.halfRAM()
	myPC.halfRAM()
	fmt.Println("Cuarto:", myPC)
}
