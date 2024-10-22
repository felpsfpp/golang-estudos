//Estruturas de dados

//Arrays e Slices
//Array (tamanho fixo)
var numeros [5]int = [5]int{1, 2, 3, 4, 5}

//Slice (tamanho dinamico)
numeros := []int{1, 2, 3, 4, 5}

//adicionar elementos a um slice
numeros = append(numeros, 6)

//maps (dicionarios)
idades := map[string]int{
	"Alice": 30,
	"Bob":   25,
}
fmt.Println(idades["Alice"]) // Output: 30
//adicionar um novo par chave-valor
idades["Charlie"] = 28

//structs
type Pessoa struct {
    Nome      string
    Idade     int
}

func main() {
	p := Pessoa{
		Nome:      "Alice",
        Idade:    30,
    }
	fmt.Println(p.Nome) // Output: Alice
}