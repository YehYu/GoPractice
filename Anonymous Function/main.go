package main

import "fmt"

//匿名函式 閉包(Closure)特性
type Getter = func() int
type Setter = func(int)

func x_getter_setter(x int) (Getter, Setter) {
	fmt.Printf("the parameter :\tx (%p) = %d\n", &x, x)

	getter := func() int {
		fmt.Printf("getter invoked:\tx (%p) = %d\n", &x, x)
		return x
	}
	setter := func(n int) {
		x = n
		fmt.Printf("setter invoked:\tx (%p) = %d\n", &x, x)
	}
	return getter, setter
}

func main() {
	//匿名函式
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2)) //3

	//匿名函式 省略變數
	sum := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(sum) //3

	//從另一個函式取得匿名函式
	getFunc := func() func(a, b int) int {
		return add
	}
	sum2 := getFunc()(1, 2)
	fmt.Println(sum2) //3

	//匿名函式 閉包(Closure)特性
	getX, setX := x_getter_setter(10) //0xc0000120d8 10
	fmt.Println(getX())               //0xc0000120d8 10
	setX(20)                          //0xc0000120d8 20
	fmt.Println(getX())               //0xc0000120d8 20

}
