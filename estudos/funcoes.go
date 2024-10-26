//funçoes

//definiçao de funçao
func saudacao (nome string) string {
	return "Olá " + nome
}

//chamada de funçao
func main() {
	mensagem := saudacao("Alice")
	fmt.Println(mensagem)
}

//funçoes com multiplos retornos
func dividir (a, b float64) (float64, error) {
	if b == 0 {
        return 0, fmt.error("Não é possivel dividir por zero")
    }
    return a / b, nil
}

//funçoes Anonimas
anonima := func(x int) int {
	return x * x
}
resultado := anonima(5)
fmt.Println(resultado) // Output: 25