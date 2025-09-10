package calculadora

func Soma(a, b int) int {
    return a + b
}

func Subtracao(a, b int) int {
    return a - b
}

func Multiplicacao(a, b int) int {
    return a * b
}

func Divisao(a, b int) (int, error) {
    if b == 0 {
        return 0, &divisaoPorZeroError{}
    }
    return a / b, nil
}

type divisaoPorZeroError struct{}

func (e *divisaoPorZeroError) Error() string {
    return "erro: divisão por zero"
}
