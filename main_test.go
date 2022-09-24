package main

import "testing"

func TestSumar(t *testing.T) {
	resultado := sumar(10, 20)
	resultadoEsperado := 30
	if resultado != resultadoEsperado {
		t.Errorf("Error en %s. Se esperaba %d pero el resultado fue %d", t.Name(), resultadoEsperado, resultado)
	}
}
