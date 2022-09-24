package main

import "fmt"

func main() {
	resultado := sumar(1, 2)
	fmt.Printf("El resultado de la suma es: %d", resultado)
}

func sumar(a int, b int) int {
	return a + b
}
