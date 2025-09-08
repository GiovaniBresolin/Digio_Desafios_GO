package main

import "fmt"

func main() {
	// Constante para o ponto de ebulição da água em Kelvin.
	const pontoEbulicaoAguaK = 373.15

	// Fator de conversão de Kelvin para Celsius.
	// A fórmula é C = K - 273.15
	const fatorConversao = 273.15

	// Calcula o valor em graus Celsius.
	var pontoEbulicaoAguaC = pontoEbulicaoAguaK - fatorConversao

	// Imprime o resultado na tela.
	fmt.Printf("O ponto de ebulição da água em Kelvin é %gK.\n", pontoEbulicaoAguaK)
	fmt.Printf("Convertendo para Celsius, o valor é %g°C.\n", pontoEbulicaoAguaC)
}
