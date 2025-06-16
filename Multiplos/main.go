/*

Voce e seus colegas querem criar em formato de codigo uma antiga brincadeira: ao falar os numeros sempre que aparecer um multiplo de 3
o participante deve falar "Pin" e nos multiplos de 5 o participante deve falar "Pan". Entao seu programa deve imprimir numeros de 1 a 100 e nos casos
citados, nao devem aparecer os numeros, mas sim, o que o participante deve falar

*/

package main

import "fmt"

func validaMultiplo(m, n int) bool {
	return n%m == 0
}

func main() {
	varRange := 100
	for i := 1; i < varRange; i++ {
		if validaMultiplo(3, i) {
			fmt.Println("Pin")
		}
		if validaMultiplo(5, i) {
			fmt.Println("Pan")
		}
	}
}
