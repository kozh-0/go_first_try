package main

// * go build && ./main - компилирует и запускает проект
// * go run main.go - тоже но в одну команду
// * go fmt . - форматирует код в текущей папке
// * go install - собирает проект и кладет в GOPATH/bin
// * go get - скачивает зависимости

import (
	"errors"
	"fmt"
)

// int и float разные int8 int16 int32 int64 и т.п.
// Обычный int 32 или 64 в зависимости от архитектуры системы

// Мораль - думать о типизации, чтобы не дебажить
// var intNum int = 32767
// intNum = intNum + 1
// fmt.Println(intNum) >>>  -3ё2768

// float разных типов может давать разные номера
func first() {
var intNum int8 = 5
	var floatNum float64 = 54321.32 
	fmt.Println("LMFAO", intNum, floatNum)

// Нельзя вычислять разные типы данных др с др. можно их привести к одному формату 
	var floatNum32 float32 = 10.2
	var intNum32 int32 = 2
	var res float32 = floatNum32 + float32(intNum32)
	// 3/2 >>> 1 | округляет вниз
	fmt.Println(res, 3/2, 3%2)


	var str string = "alo \nkek" 
	var str1 string = `alo
	kek` 

	// len считает байты, для длины строки надо импортить utf8.RuneCountInString 
	fmt.Println(len(str + str1))

	var myRune rune = 'a' // >>> 87 (короче это странная юзлес залупа типа Symbol в JS)
	fmt.Println(myRune)
	
	var myBoolean bool = false
	fmt.Println(myBoolean)

	// var kek = "test" === kek := "test"
	// var var1, var2 := 1, 2

	// как в js короче, чтобы не изменять
	const pi float32 = 3.14
	fmt.Println(pi)

	// func, if ==============================================

	kek("HAHA")

	var result, remainder, err = intDevision(7,2)
	// Также можно написать через Switch, break в конце каждого case: не надо писать, в отличии от других языков
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The res of division is %v", 0)
	} else {
		fmt.Println("Smth went wrong...")
	}
	fmt.Printf("Результат %v; Остаток - %v \n\n", result, remainder)

}

func kek(str string) {
	fmt.Println(str + " ale ya ne vizival")
}
func intDevision(intNum1 int, intNum2 int) (int, int, error) {
	var err error // технически это nil
	if intNum2 == 0 {
		err = errors.New("can't devide by 0")
		return 0, 0, err
	}

	var res int = intNum1 / intNum2
	var remainder int = intNum1 % intNum2
	return res, remainder, err
}