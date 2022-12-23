package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFunc(t *testing.T) {

	valores := []float64{2, 3, 3, 4, 10, 2, 4, 5}
	valor_esperado := float64(2)
	respuesta := MinFunc(valores...)
	assert.Equal(t, valor_esperado, respuesta, "Minimo incorrecto")

}

func TestMaxFunc(t *testing.T) {

	valores := []float64{2, 3, 3, 4, 10, 2, 4, 5}
	valor_esperado := float64(10)
	respuesta := MaxFunc(valores...)
	assert.Equal(t, valor_esperado, respuesta, "Maximo incorrecto")

}

func TestAverageFunc(t *testing.T) {

	valores := []float64{10, 5}
	valor_esperado := float64(7.5)
	respuesta := AverageFunc(valores...)
	assert.Equal(t, valor_esperado, respuesta, "Average incorrecto")

}
