package main

import (
	"fmt"
)

func main() {

	// PRIMER FORMA DE DECLARAR VARIABLES

	//Declarar variable sin valor
	var x int
	var y string
	var z bool

	//Inicializar variables
	x = 10
	y = "Hello"
	z = true

	fmt.Println(x, y, z)

	//Declarar multiples variables

	var a, b, c int
	a = 10
	b = 20
	c = 30
	fmt.Println(a, b, c)

	//###########################################
	// SEGUNDA FORMA DE DECLARAR VARIABLES

	// Declarar vARAIBLES E INICIALIZARLAS

	var verdura string = "tomate"
	var fruta string = "naranja"

	fmt.Println(verdura, fruta)

	// Declarar variables sin asignar tipo

	var comida = "pizza"
	fmt.Println(comida)

	//############################################
	// TERCERA FORMA DE DECLARAR VARIABLES

	//Operador corto u operador gopher :=

	mascota := "gopher"
	fmt.Println(mascota)
	index := 10
	fmt.Println(index)
}
