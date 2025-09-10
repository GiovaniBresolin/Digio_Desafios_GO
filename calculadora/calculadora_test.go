package calculadora

import (
    "testing"
)

func TestSoma(t *testing.T) {
    resultado := Soma(5, 5)
    esperado := 10
    if resultado != esperado {
        t.Errorf("Soma(5, 5) = %d; esperado %d", resultado, esperado)
    }
}

func TestSubtracao(t *testing.T) {
    resultado := Subtracao(10, 5)
    esperado := 5
    if resultado != esperado {
        t.Errorf("Subtracao(10, 5) = %d; esperado %d", resultado, esperado)
    }
}

func TestMultiplicacao(t *testing.T) {
    resultado := Multiplicacao(5, 5)
    esperado := 25
    if resultado != esperado {
        t.Errorf("Multiplicacao(5, 5) = %d; esperado %d", resultado, esperado)
    }
}

func TestDivisao(t *testing.T) {
    resultado, err := Divisao(10, 2)
    if err != nil {
        t.Errorf("Divisao(10, 2) retornou um erro inesperado: %v", err)
    }
    esperado := 5
    if resultado != esperado {
        t.Errorf("Divisao(10, 2) = %d; esperado %d", resultado, esperado)
    }
}

func TestDivisaoPorZero(t *testing.T) {
    _, err := Divisao(10, 0)
    if err == nil {
        t.Errorf("Divisao(10, 0) deveria retornar um erro, mas não retornou")
    }
}
