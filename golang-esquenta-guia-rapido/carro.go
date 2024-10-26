package main

import (
	"database/sql"
	"fmt"
)

type Veiculo interface {
	Andar()
}

type Carro struct {
	Modelo string
	ano    int
}

func (c Carro) Andar() {
	fmt.Println("O carro", c.Modelo, " está andando")
}

func VemAndarComigo(v Veiculo) {
	v.Andar()
}

func main() {
	//sqlite3
	db, err := sql.Open("sqlite3", "teste.db")
	if err != nil {
		panic(err)
	}
	defer db.Close() //vai ser executado no final da aplicaçao por conta do "defer"

	carro1 := Carro{Modelo: "Fusca", ano: 1969}
	insertCarro(db, carro1)
}

func insertCarro(db *sql.DB, carro Carro) {
	_, err := db.Exec("INSERT INTO carro (modelo, ano) VALUES (?, ?)", carro.Modelo, carro.Ano)
	if err != nil {
		panic(err)
	}

}
