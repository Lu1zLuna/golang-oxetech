package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	/* // Exercício 1 - Soma de Inteiros e Float
	var var1 int
	var var2 float64

	fmt.Scanln(&var1)
	fmt.Scanln(&var2)
	somaIntEfloat(var1, var2) */

	/* // Exercício 2 - Concatenar Inteiro e String
	var var3 int
	var var4 string


	fmt.Printf("Digite um inteiro: ")
	fmt.Scanln(&var3)
	fmt.Printf("Digite uma string: ")
	fmt.Scanln(&var4)
	concatIntString(var3, var4)] */
	
	/* // Exercício 3 - Divisão de Float e Int

	var var5 float64
	var var6 int
	
	fmt.Printf("Digite um número com casas decimais: ")
	fmt.Scanln(&var5)
	fmt.Printf("Digite um inteiro: ")
	fmt.Scanln(&var6)

	divisaoFloatInt(var5, var6) */

	/* // Exercício 4 - Comparação de Strings
	var var7 string
	var var8 string

	fmt.Printf("Digite a primeira string: ")
	fmt.Scanln(&var7)
	fmt.Printf("Digite a segunda string: ")
	fmt.Scanln(&var8)

	CompararStrings(var7, var8) */

	/* // Exercício 5 - Operações com UInt e Int
	var var9 uint
	var var10 int

	fmt.Printf("Digite o primeiro valor: ")
	fmt.Scanln(&var9)
	fmt.Printf("Digite o segundo valor: ")
	fmt.Scanln(&var10)

	SomarUintEInt(var9, var10) */

	/* // Exercício 6 - Conversão de Booleano para Inteiro
	var var11 bool = true
	var var12 bool = false

	ConvertBoolToInt(var11)
	ConvertBoolToInt(var12) */

	/* // Exercício 7 - Representação Binária de Inteiro
	var var13 int64
	fmt.Printf("Digite um número inteiro: ")
	fmt.Scanln(&var13)

	MostrarBinarioDeInt(var13) */

	/* // Exercício 8 - Cálculo de Área de um Círculo
	var var14 float64
	fmt.Printf("Digite a área do círculo: ")
	fmt.Scanln(&var14)

	CalcularAreaCirculo(var14) */

	/* // Exercício 9 - Concatenação de Strings e Float
	var var15 string
	fmt.Scanln(&var15)

	ConcatStringFloat(var15) */

	// Exercício 10 - Conversão de Inteiro para Caractere
	var var16 byte
	fmt.Printf("Digite um número inteiro positivo entre 0 e 255: ")
	fmt.Scanln(&var16)
	ConvertIntToChar(var16)
}

func somaIntEfloat(x int, y float64) {
	var a int = x
	var b float64 = y
	
	var c float64 = float64(a) + b
	
	fmt.Println("1. Conversao int-float: ", c)
}

func concatIntString(x int, y string) {
	var n int = x
	var text string = y
	
	convertString := strconv.Itoa(n)
	result := convertString + text
	
	fmt.Println("2. String convertida: ", result)
}

func divisaoFloatInt(x float64, y int) {
	z := x / float64(y)
	
	fmt.Println("3. Divisão de Float e Int: ", z)
}

func CompararStrings(str1 string, str2 string) {
	var isEqual bool = str1 == str2
	
	fmt.Println("4. Resultado da comparacao: ", isEqual)
}

func SomarUintEInt(m uint, n int) {
	var sum uint = m + uint(n)
	
	fmt.Println("5. Resultado da soma: ", sum)
}

func ConvertBoolToInt(flag bool) {
	var intFlag int
	if flag {
		intFlag = 1
		} else {
			intFlag = 0
		}
		
		fmt.Println("4. Resultado da conversao: ", intFlag)
	}
	
	func MostrarBinarioDeInt(num int64) {
		var binaryStr string = strconv.FormatInt(num, 2)
		
		fmt.Println("5. Representacao em binario: ", binaryStr)
	}
	
	func CalcularAreaCirculo(radius float64) {
		var area float64 = math.Pi * math.Pow(radius, 2)
		
		fmt.Printf("5. Área do círculo: %.2f", area)
	}
	
	func ConcatStringFloat(message string) {
	var pi float64 = 3.14159
	var fullMessage string = strconv.FormatFloat(pi, 'f', -1, 64) + message
	
	fmt.Println("6. Mensagem concatenada:", fullMessage)
}

func ConvertIntToChar(ascii byte) {
	var char string = string(ascii)

	fmt.Println("10. Inteiro convertido para ASCII: ", char)
}