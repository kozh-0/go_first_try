package main

import "fmt"

type Person struct {
	name string
	age int8
}

func third(){
	var Dima Person = Person{name: "Dima", age: 26} // или {"Dima", 26}
	fmt.Printf("3d file: %v", Dima)	
}