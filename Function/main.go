package main

import (
	"fmt"
	"reflect"
)

//定義新的型態
type addFunc func(int, int) int

func add(a, b int) (result int) {
	result = a + b
	return
}

//function 當變數
func cmd(a, b int, add addFunc) int {
	return add(a, b)
}

func addmore(numbers ...int) (sum int) { //傳入值是不定的，可以用 ... 來表示
	sum = 0
	for _, num := range numbers { // numbers 型態會是 Slice []int ，所以可以用 for range 走訪
		sum += num
	}

	return
}

func main() {
	var add2 addFunc
	//匿名函式
	add2 = func(a, b int) int {
		return a + b
	}

	fmt.Println(add2(1, 2))     // 3
	fmt.Println(cmd(1, 2, add)) // 3
	//inline func
	fmt.Println(cmd(1, 2, func(a, b int) int {
		return a + b
	})) // 3
	fmt.Println(reflect.TypeOf(add))  // func(int, int) int
	fmt.Println(reflect.TypeOf(add2)) // func(int, int) int

	sum := addmore(1, 2, 3, 4, 5)

	fmt.Println(sum) // 15
}
