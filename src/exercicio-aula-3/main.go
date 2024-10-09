package main

import (
	"fmt"
	"strings"
)

func main() {
	/* // 1. Soma de Elementos em um Array:
	array := []int{1, 2, 3, 4, 5}
	result := SomarElementosArray(array)

	fmt.Println(result) */

	/* 	// 2. Inverter um Array:
	   	array := []int{3, 5, 7, 9}
	   	fmt.Println(InvertArray(array)) */

	/* // 3. Contagem de Caracteres em uma String (Map):
	var word string
	fmt.Printf("Digite uma palavra: ")
	fmt.Scanln(&word)

	CountCharOnString(word) */

	// 4. Agrupar Palavras por Comprimento (Map):
	// Exemplo de uso
	palavras := []string{"Alagoas", "Maceió", "oxetech", "go", "golang", 
		"python", "javascript", "php"}
	resultado := AgruparPorComprimento(palavras)

	// Exibe o map resultante
	fmt.Println(resultado)

/* 	// 5. Números Pares em um Intervalo:
	paresEmIntervalo() */

	/* // 6. Soma dos Números Ímpares:
	somarImparesEntreIntervalo(1, 10)
	somarImparesEntreIntervalo(97, 10) */

	/* // 7. Verificar Número Primo:
	fmt.Println(VerificarNumeroPrimo(7)) */

	/* // 8. Calcular o Fatorial:
	fmt.Println(Fatorial(10)) */

	/* // 9. Máximo e Mínimo de um Array:
	array := []int{3, 5, 7, 2, 1, 6, 8, 3, 4, 15, 9, 28}
	maximum, minimum := MaxMinArray(array)
	fmt.Println("Minimo:", minimum)
	fmt.Println("Maximo:", maximum) */

	/* // 10. Contagem de Vogais:
	text := "Olá, testando!"
	fmt.Println(ContarVogais(text)) */
}

// 1.
func SomarElementosArray(array []int) int {
	var somaElementos int
	for v := range array {
		somaElementos += v
	}
	return somaElementos
}

// 2.
func InvertArray(array []int) []int {
	arrayInvertido := make([]int, len(array))
	
	for i := range array {
		arrayInvertido[i] = array[len(array)-1-i]
	}
	
	return arrayInvertido
}

// 3.
func CountCharOnString(word string) map[string]int {
	var count int
	for i := range word {
		count = i
	}
	
	wordMap := map[string]int{word: count + 1}
	
	for c, v := range wordMap {
		fmt.Println("Palavra:", c, " | Num de caracteres: ", v)
	}
	return wordMap
}

// 5.
func paresEmIntervalo() {
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}

// 6.
func somarImparesEntreIntervalo(min int, max int) {
	var soma int
	
	// Garantir que o intervalo seja sempre crescente
	if min > max {
		min, max = max, min
	}
	
	// Loop para percorrer o intervalo de min a max
	for i := min; i <= max; i++ {
		// Verifica se o número é ímpar
		if i % 2 != 0 {
			soma += i
		}
	}
	
	fmt.Println(soma)
}

// 7.
func VerificarNumeroPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num % 2 == 0 {
			return true
		}
	}
	return true
}

// 8.
func Fatorial(num int) int {
	resultado := 1
	
	if num < 0 {
		fmt.Println("Não existe fatorial de número negativo.")
        return 0
		} else if num == 0 {
			return 1
		}
		
		for i := 1; i <= num; i++ {
			resultado *= i
		}
		
		return resultado
	}
	
	// 9.
	func MaxMinArray(array []int) (int, int) {
		max := array[0]
		min := array[0]
		
		for i := 0; i < len(array); i++ {
			if array[i] > max {
				max = array[i]
			}
			if array[i] < min {
				min = array[i]
			}
		}
		
		return min, max
	}
	
	func ContarVogais(text string) int {
		vogais := "aeiouAEIOU"
		
		contador := 0
		for _, char := range text {
			if strings.ContainsRune(vogais, char) {
				contador++
			}
		}
		
		return contador
	}
	
	func AgruparPorComprimento(palavras []string) map[int][]string {
		grupos := make(map[int][]string)


		for _, palavra := range palavras {
			comprimento := len(palavra)
			// Adiciona a palavra ao array do map correspondente ao comprimento
			grupos[comprimento] = append(grupos[comprimento], palavra)
		}

		return grupos
	}