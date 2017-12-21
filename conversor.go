package main

import (
	"fmt"     // Para lidar com formatacao de strings
	"os"      // Para operacoes que lidam com o sistema operacional
	"strconv" // Para conversao de strings
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1] // Considerando o segundo elemento e removendo o ultimo

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", unidadeDestino)
		os.Exit(1)
	}

	for i, v := range valoresOrigem {
		valoresOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf(
				"O valor %s na posição %d não é um número válido!\n",
				v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valoresOrigem*1.8 + 32
		} else {
			valorDestino = valoresOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			valoresOrigem, unidadeOrigem,
			valorDestino, unidadeDestino)
	}
}
