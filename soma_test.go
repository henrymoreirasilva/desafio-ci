package main

import (
	"testing"
)

func TestSoma(t *testing.T)  {
	s := Soma(5, 5)

	if s != 10 {
		t.Errorf("Falha de teste unitário");
	}
}