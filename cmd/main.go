package main

import (
	"log"

	goarray "github.com/lcnssantos/gorray"
)

func main() {
	data := goarray.New(&[]string{"Luciano", "Souza", "Santos"})

	element := data.Some(func(element string) bool {
		return true
	})

	log.Println(element)
}
