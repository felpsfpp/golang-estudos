package main

import "fmt"

func main() { //thread 1
	canal := make(chan string) //Vazio
	//canal é como um tunel de comunicaçao entre go rotinas, entre funçoes que o go executa em background

	go func() { //thread 2
		canal <- "Olá, canal!" //Preenchido
	}()

	fmt.Println(<-canal) //thread 1 //Vazio
}
