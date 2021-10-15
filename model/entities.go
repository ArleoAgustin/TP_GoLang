package model

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(t string, l int, v string) Result {
	return Result{t, l, v}
}

func SetCadena(c string) (Result, error) {

	t := c[0:2]
	l := c[2:4]
	v := c[4:]
	s, _ := strconv.Atoi(l) //lo convierto a entero
	nObj := NewResult(t, s, v)

	if len(nObj.Value) == nObj.Length { //se verifica que el value coincida con la longuitud
		if esCorrecto(nObj) {
			return nObj, nil
		}
	}

	return Result{}, errors.New("Error, cadena invalida")
}

func esCorrecto(r Result) bool {

	var cont = 0
	if r.Type == "NN" {

		for i := 0; i < len(r.Value); i++ {

			digit := (r.Value[i : i+1])

			if digit >= "0" && digit <= "9" {
				cont++
			}
		}
		if cont == r.Length {
			return true
		}
		return false

	} else if r.Type == "TX" {
		for _, v := range r.Value {

			if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
				cont++
			}
		}
		if cont == r.Length {
			return true
		}
	}
	return false
}
