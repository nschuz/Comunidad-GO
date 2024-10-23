package main

import "fmt"

func main() {

	// const PI float32 = 3.141592653589793
	// const EULER float32 = 2.718281828459045

	const (
		PI     float32 = 3.141592653589793
		EULER  float32 = 2.718281828459045
		GRAVED float32 = 9.81
	)

	fmt.Println("PI:", PI, "EULER:", EULER, "GRAVED:", GRAVED)

	const (
		sunday = iota * 7
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday)
}
