package main

import (
	"fmt"
	"strconv"
)

func main() {
	var palavra string
	fmt.Println("Digite uma palavra:")
	fmt.Scan(&palavra)

	if isDecrescente(palavra) {
		fmt.Println("A palavra é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A palavra não é uma sequência numérica decrescente.")
	}
}

func isDecrescente(palavra string) bool {
	for i := 0; i < len(palavra)-1; i++ {
		num1, _ := strconv.Atoi(string(palavra[i]))
		num2, _ := strconv.Atoi(string(palavra[i+1]))
		if num1 <= num2 {
			return false
		}
	}
	return true
}