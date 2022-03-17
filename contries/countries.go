package countries

import (
	"errors"
	"fmt"
)

var myCountry string

var collecion []string
var errNotFound error = errors.New("Country not found")

//si queremos que sea publica y que sea utiliozada por paquetes externos hay que declararla en mayuscula
//Add agrega el valor de un pais a la coleccion de paises
func Add(country string) {
	collecion = append(collecion, country)
}

// SetMyCountry selecciona un pais como pais favorito
func SetMyCountry(country string) error {
	/*for _, c := range collecion {
		if c == country {
			myCountry = country
		}
	}*/
	if !isInCollecion(country) {
		return errNotFound
	}
	myCountry = country
	return nil
}

//List lista todos los paises
func List() {
	for i, c := range collecion {
		meCountryMsg := ""
		if c == myCountry {
			meCountryMsg = "[me country]"
		}
		fmt.Printf("%d. %s %s \n", i+1, c, meCountryMsg)
	}
}

//isInCollecion verifica si el pais esta en la colecion
func isInCollecion(country string) bool {

	for _, c := range collecion {
		if c == country {
			return true
		}
	}
	return false
}
