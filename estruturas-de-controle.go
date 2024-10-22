//estruturas de controle

//condicionais(if, else, else if)
if idade >= 18 {
	fmt.Println("Você é maior de idade")
} else {
	fmt.Println("Você é menor de idade")
}

//switch
dia := "segunda"
switch dia {
case "segunda":
	fmt.Println("Inicio da semana")
case "sexta":
	fmt.Println("Fim da semana")
default:
	fmt.Println("Meio da semana")
}

//laços (for)
//basico
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

//como um while
i := 0
for i < 10 {
	fmt.Println(i)
	i++
}

//iteraçao sobre slices ou arrays
nums := []int{2, 4, 6, 8}
for index, value := range nums {
	fmt.Println(index, value)
}