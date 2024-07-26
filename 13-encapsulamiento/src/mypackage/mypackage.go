package mypackage

import "fmt"

// Comentarios obligatorios

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  uint16
}

type carPrivate struct {
	brand string
	year  uint16
}

// Imprime un Hola subject
func Salutation(subject string) {
	fmt.Println("Hola", appendUwU(subject))
}

// La siguiente funcion es privada
func appendUwU(text string) string {
	return (text + " UwU")
}

// SI LA FUNCION NO REQUIERE UN ACCESO PUBLICO, ES MEJOR MANTENERLA PRIVADA.
