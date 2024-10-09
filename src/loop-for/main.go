package main

import "fmt"

func main() {
	/* arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println("Índice:", i, "| Valor: ", v)
	} */

	s := "golang"
	for i, c := range s {
		fmt.Println(i, string(c))
	}
}
