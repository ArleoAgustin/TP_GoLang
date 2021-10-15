package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {

		r, _ := SetCadena(testData.Input)

		assert.Equal(t, r.Type, testData.Type, testData.Success)
		assert.Equal(t, r.Value, testData.Value, testData.Success)
		assert.Equal(t, r.Length, testData.Length, testData.Success)

		y, _ := SetCadena("NN040087")

		assert.Equal(t, y.Value, "0087", "funciona")
		assert.Equal(t, y.Type, "NN", "funciona")
		assert.Equal(t, y.Length, 4, "funciona")

		x, _ := SetCadena("TX02AV")

		assert.Equal(t, x.Value, "AV", "funciona")
		assert.Equal(t, x.Length, 2, "funciona")

		_, err := SetCadena("TX03A2C")

		assert.NotEqual(t, err, nil, "esta cadena da error")
	}
}
