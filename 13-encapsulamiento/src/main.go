package main

// Se importa el path desde el src inicial
import (
	//"FundamentosGo/13-encapsulamiento/src/mypackage"
	alias "FundamentosGo/13-encapsulamiento/src/mypackage"
	"fmt"
)

func main() {
	// Solo son accesibles las variables de clase publicas, las cuales empiezan con mayuscula.
	var myCar alias.CarPublic
	var newCar alias.CarPublic
	myCar.Brand = "Toyota"
	myCar.Year = 2001
	newCar.Brand = "Ford"
	newCar.Year = 1981

	fmt.Println(myCar)
	fmt.Println(newCar)

	alias.Salutation("Peter")

}
