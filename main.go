package main

import (
	"fmt"
	countries "packages/contries"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	fmt.Println(id)

	countries.Add("USA")
	countries.Add("MExico")
	countries.Add("Japan")
	countries.Add("Spain")

	countries.List()

	err := countries.SetMyCountry("Japan")

	if err != nil {
		panic(err)
	}
}

//comando para instalar paquetes de terceros
// go get github.com/google/uuid
