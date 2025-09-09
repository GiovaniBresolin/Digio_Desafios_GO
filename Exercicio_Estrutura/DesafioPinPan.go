package main

import "fmt"

func main() {
	fmt.Println("Iniciando a brincadeira Pin Pan de 1 a 100:")

	// Itera sobre os números de 1 a 100.
	for i := 1; i <= 100; i++ {
		// É crucial verificar primeiro a condição de ser múltiplo de ambos (3 e 5).
		// Se não, o programa entraria na condição de 'múltiplo de 3' ou 'múltiplo de 5'
		// e nunca imprimiria "PinPan".
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 { // Se for múltiplo apenas de 3
			fmt.Println("Pin")
		} else if i%5 == 0 { // Se for múltiplo apenas de 5
			fmt.Println("Pan")
		} else { // Se não for múltiplo de 3 nem de 5
			fmt.Println(i)
		}
	}
}
