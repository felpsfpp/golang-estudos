package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d recebeu %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { //thread 1
	canal := make(chan int)
	qtdWorkers := 3 //cria 3 workers

	for i := range qtdWorkers {
		go worker(i, canal) // Executa worker em uma goroutine separada
	}

	for i := range 10 {
		canal <- i //fica travado ate o canal ser esvaziado. quando esvaziar, o loop continua // Envia valores para o canal
	}
}
