package main

import (
	"testing"
)

func TestCalculo(t *testing.T) {
	got := "7"
	want := "7"

	if got != want {
		t.Errorf("Calculo da raiz quadrada inválida: obteve %s e o esperado era: %s", got, want)
	}
}
