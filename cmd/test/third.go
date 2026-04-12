package main

import "fmt"

type Person struct {
	name string
	age int8
	car Car
}
type Car struct {
	model string
	year int
}

func third(){
	var Dima = Person{name: "Dima", age: 26, car: Car{
		model: "bmw 34", year: 2000,
		},
	} // или {"Dima", 26}
	Dima.car = Car{ model: "KIA RIO", year: 2017}
	fmt.Printf("3d file: %v", Dima)	
}

func milesToKm(miles float32) {
	
}