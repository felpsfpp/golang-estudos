//Concorrencia

//Go suporta concorrencia nativa usando goroutines e channels.

//goroutines
//Uma goroutine é uma função executada em paralelo com outras goroutines:

func digaOla() {
	fmt.Println("Olá")
}

func main() {
	go digaOla() // Executa digaOla em uma goroutine separada
	fmt.Println("Mundo")

	time.Sleep(1 * time.Second) // Espera a goroutine digaOla terminar
}

//canais (Channels)
//comunicaçao segura entre goroutines

func digaOla(canal chan string) {
	canal <- "Olá do canal!"
}

func main() {
	canal := make(chan string)
	go digaOla(canal)     // Executa digaOla em uma goroutine separada
	mensagem := <-canal   // Espera a goroutine digaOla enviar uma mensagem no canal
	fmt.Println(mensagem) // Output: Olá do canal!
}

