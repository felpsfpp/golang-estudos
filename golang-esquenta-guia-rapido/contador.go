package main

import (
	"fmt"
	"time"
)

func contador(qtd int) {
	for i := range qtd {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() { //1
	go contador(10) //thread 2
	go contador(10) //thread 3
	contador(10)
	//o "go" serve para colocar cada açao em uma thread
}
