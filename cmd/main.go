package main

import (
	"log"
	"strings"

	goarray "github.com/lcnssantos/gorray"
)

func main() {
	data := goarray.New(&[]string{"A"})

	newArray := data.Concat(goarray.New(&[]string{"A"}), goarray.New(&[]string{"C"}))

	log.Println(data)
	result := newArray.Map(func(element string) interface{} {
		return strings.ToLower(element)
	})

	log.Println(result)
}
