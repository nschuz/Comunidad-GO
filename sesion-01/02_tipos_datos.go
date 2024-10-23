package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Tipos de datos básicos en Go

	// 1. Enteros (int, int8, int16, int32, int64)
	var entero int = 42
	var enteroChico int8 = -128 // Rango de -128 a 127
	fmt.Println("Entero:", entero, "Entero de 8 bits:", enteroChico)

	// 2. Números sin signo (uint, uint8, uint16, uint32, uint64)
	var enteroPositivo uint = 42
	var byteAlias uint8 = 255 // byte es un alias de uint8
	fmt.Println("Entero positivo:", enteroPositivo, "Byte (uint8):", byteAlias)

	// 3. Números flotantes (float32, float64)
	var decimalPequeno float32 = 3.1416872
	var decimalGrande float64 = 3.141592653589793
	fmt.Println("Decimal pequeño (float32):", decimalPequeno, "Decimal grande (float64):", decimalGrande)

	// 4. Booleanos (bool)
	var esVerdad bool = true
	var esFalso bool = false
	fmt.Println("Es verdad:", esVerdad, "Es falso:", esFalso)

	// 5. Strings (cadenas de texto)
	var saludo string = "Hola, mundo"
	fmt.Println("Saludo:", saludo)

	// 6. Caracteres (rune) - representan un carácter Unicode
	var caracter rune = 'A' // rune es un alias para int32
	fmt.Println("Carácter:", caracter, "Valor Unicode:", int(caracter))

	//

	fmt.Println("El tipo de dato es: ", fmt.Sprintf("%T", enteroChico))
	fmt.Printf("El tipo de dato es: %T\n", enteroChico)
	fmt.Println("Tipo de dato: ", reflect.TypeOf(enteroChico))

}
	