package main

import "fmt"

// arr, slices, maps, loops
func second() {
	var firstArr [3]int32 // >>> [0,0,0]
	firstArr[1] = 228

	fmt.Println(firstArr)
	// ссылка на значение в памяти
	fmt.Println(&firstArr[0]) 
	// другие эелементы массива будут под номером +4 

	// ... чтобы не считать сколько элементов в типе будет, можно и цифру вставить
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr)
	
	// slice и arr одно и то же, но смешивать их нельзя
	intSlice := []int32{4,5,6}
	
	fmt.Println("ДО", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // >> после этого вместимость массива 6, длина 4, из-за внутренних приколов ГО
	fmt.Println("ПОСЛЕ",len(intSlice), cap(intSlice))
	
	intSlice2 := []int32{8,9} 
	intSlice2 = append(intSlice2, intSlice...) // spread с другой стороны кек
	fmt.Println(intSlice2, "\n")
	

}

func secondMapsAndLoops(){
	// Map - это объект как в JS, [string]int32 - ключ строка, значение инт32
	var myMap map[string]int32 = make(map[string]int32)
	fmt.Println(myMap)

	var myMap2 = map[string]int8{
		"kek": 7,
		"lol": 42,
	}

	
	// delete(myMap2, "kek") // удалить элемент из мапы
	fmt.Println(myMap2, myMap2["kek"]) // если элемента нет, вернет дефолтное значение типа, для int 0, stirng = "", bool = false

	// Мап возвращает ok 
	var kek, ok = myMap2["kek"] // ok = true/false если такой ключ есть
	if ok {
		fmt.Println("Значение есть", kek)
	} else {
		fmt.Println("Такого ключа нет")
	}

	// * The range is a form of the for loop that iterates over a slice or map.
	// перебор map или array через range, ключи вывыодит в random порядке
	// возвращает сначала ключ, потом значение (для array - индекс потом значение)
	for name, value := range myMap2 {
		fmt.Println("Ключ:", name, "Значение:", value)
	} 


	var i int = 0
	// while loop
	for i<10 {
		fmt.Println("Цикл по условию:", i)
		i = i + 1
	}

	for {
		if i>5 {
			break
		}
		fmt.Println("Бесконечный цикл с брейком:", i)
		i = i + 1
	}

	// full for loop
	for i:=0; i<5; i++ {
		fmt.Println("Полный цикл фор:", i)
	}


	var myString = "resume"
	for i, char := range myString {
		fmt.Printf("idx: %v; char: %v; Порядковый номер кодировки %v \n", i, string(char), char)
	}
}