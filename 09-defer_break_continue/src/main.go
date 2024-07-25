package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Defer
	// Las lineas con defer se ejecutan hasta antes de la salida del codigo.
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
	// Usa defer cuando se requiera una accion antes del final del codigo
	// Por ejemplo para cerrar una coneccion a base de datos.
	// Por buena practica, debe usarse solo una sentencia defer dentro de una funcion.

	// Continue y break
	// En ciclos for
	for i := 0; i < 130; i++ {
		fmt.Println("Convirtiendo", strconv.Itoa(i), "a string")
		if i%10 == 0 {
			fmt.Println("Multiplo de 10")
			continue
			// Usar cuando interesa que continue el bucle mayor
			// Al ejecutarse, se reinicia toda la iteracion ignorando el resto del bloque
		}

		if i%72 == 0 {
			fmt.Println("Multiplo de 72, ruptura")
			break
		}
	}
	fmt.Println("Fin del programa")
}
