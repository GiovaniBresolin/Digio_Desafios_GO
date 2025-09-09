package main

import "fmt"

func main() {
	// Imprime um título para a saída.
	fmt.Println("Números entre 1 e 100 divisíveis por 3:")

	// Itera sobre os números de 1 a 100.
	for i := 1; i <= 100; i++ {
		// Verifica se o número atual (i) é divisível por 3.
		// O operador '%' (módulo) retorna o resto de uma divisão.
		// Se o resto for 0, o número é perfeitamente divisível.
		if i%3 == 0 {
			// Imprime o número que atende à condição.
			fmt.Println(i)
		}DDD
	}
}
