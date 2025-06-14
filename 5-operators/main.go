package main

import "fmt"

func main() {
	// ARITMÉTICOS
	sum := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	multi := 10 * 5
	resto := 10 % 2

	fmt.Println(sum, sub, div, multi, resto)

	var number1 int16 = 10
	var number2 int16 = 25
	sum2 := number1 + number2
	fmt.Println(sum2)

	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	var var1 string = "String"
	var2 := "String2"
	fmt.Println(var1, var2)
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	fmt.Println("--------------")
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
