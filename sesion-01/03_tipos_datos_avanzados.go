package main

import "fmt"

func main() {

	// Tipos más avanzados

	// 7. Arrays - tamaño fijo
	var numeros [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", numeros)

	// 8. Slices - tamaño dinámico (basados en arrays)
	slice := []int{10, 20, 30}
	fmt.Println("Slice:", slice)

	// 9. Maps - tipo de datos clave-valor
	mapa := map[string]int{
		"A": 1,
		"B": 2,
	}
	fmt.Println("Mapa:", mapa)

	// 10. Structs - tipo de datos personalizado
	type Persona struct {
		Nombre string
		Edad   int
	}

	juan := Persona{"Juan", 25}
	fmt.Println("Persona:", juan)
	fmt.Println("Edad:", juan.Edad)

	// 11. Punteros - apuntan a la dirección de memoria de una variable
	var numero int = 10
	var puntero *int = &numero
	fmt.Println("Valor del puntero:", *puntero, "Dirección de memoria:", puntero)
}
