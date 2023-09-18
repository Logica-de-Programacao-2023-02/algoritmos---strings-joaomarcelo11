package main

import (
	"fmt"
	"strings"
)

func main() {

	var x, padrao string
	fmt.Print("Digite uma string:")
	fmt.Scan(&x)
	fmt.Print("Digite um padrão:")
	fmt.Scan(&padrao)

	var indexes []int

	for i := len(x) - 1; i >= 0; i -= len(padrao) {
		idx := strings.Index(x[i:], padrao)
		if idx != -1 {
			indexes = append(indexes, i+idx)
		}

		fmt.Println(indexes)
	}

}
