// Voces e seus colegas querem criar um codigo que exiba todos os numeros entre 1 a 100 que sao divisiveis por 3

package main

import "fmt"

func DivisivelPTres(number int) bool {
	return number%3 == 0
}

func main() {
	var i int
	varRange := 100
	for i = 1; i < varRange; i++ {
		if DivisivelPTres(i) {
			fmt.Printf("O numero %d, e divisivel por 3.\n", i)
		}
	}
}
