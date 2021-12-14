package main

import "testing"

func TestSoma(t *testing.T) {
	// Given
	valorEsperado1 := 15
	valorEsperado2 := 15
	resultadoEsperado := 30

	// When
	resultadoObtido := Soma(valorEsperado1, valorEsperado2)

	// Then
	if resultadoObtido != resultadoEsperado {
		t.Errorf("Resultado da soma é inválido! \nObtido: %d \nEsperado: %d", resultadoObtido, resultadoEsperado)
	}
}
