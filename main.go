package main

import (
	"fmt"

	"entregableGo.com/model"
)

func main() {

	c := "TX03ABC"
	r, err := model.SetCadena(c)

	if err == nil {
		fmt.Println(r)
	} else {
		fmt.Println(err)
	}
}
