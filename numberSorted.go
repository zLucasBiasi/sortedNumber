package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var response int

	numberSorted := rand.Intn(20)

	fmt.Println(numberSorted)

	attemps := 3

	for {
		fmt.Println("digite um numero e vamos ver se voce vai acertar")
		fmt.Println("você terá três tentativas")
		fmt.Scanln(&response)
		if numberSorted == response {
			fmt.Println("voce acertou!!")
			break
		} else {
			fmt.Println("Tente novamente")
			attemps -= 1
			fmt.Printf("numero de tentativas: %d\n", attemps)
		}

		if attemps == 0 {
			fmt.Println("Voce perdeu")
			break
		}
	}

}
