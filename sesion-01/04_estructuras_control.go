package main

import (
	"fmt"
)

func main() {
	// 1. if - else: Estructura condicional básica
	num := 7
	if num%2 == 0 {
		fmt.Println("El número es par.")
	} else {
		fmt.Println("El número es impar.")
	}

	// 2. if - else if - else: Varias condiciones
	if num < 0 {
		fmt.Println("El número es negativo.")
	} else if num == 0 {
		fmt.Println("El número es cero.")
	} else {
		fmt.Println("El número es positivo.")
	}

	// 3. switch: Selección de casos basada en el valor de una variable
	dia := 4
	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	default:
		fmt.Println("Fin de semana")
	}

	// 4. switch sin expresión: Útil para evaluar condiciones complejas
	nombre := "Juan"
	switch {
	case len(nombre) > 5:
		fmt.Println("El nombre es largo.")
	case len(nombre) == 5:
		fmt.Println("El nombre tiene 5 letras.")
	default:
		fmt.Println("El nombre es corto.")
	}

	// 5. for: Bucle con condición
	for i := 0; i < 5; i++ {
		fmt.Printf("Valor de i: %d\n", i)
	}

	// // 6. for tipo while: Bucle que se ejecuta mientras se cumple la condición
	suma := 0
	for suma < 10 {
		suma++
		fmt.Printf("Suma actual: %d\n", suma)
	}

	// // 7. for infinito con break: Se usa para crear bucles infinitos que salen con break
	for {
		fmt.Println("Bucle infinito.")
		break // Salimos del bucle
	}

	// // 8. continue: Salta a la siguiente iteración del bucle
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Salto la iteración 2.")
			continue
		}
		fmt.Printf("Valor de i: %d\n", i)
	}

	

}
