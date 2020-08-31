package main

import "fmt"

type People struct {
	name string
	age  int
}

func (people People) Hello() string {
	return `Hi! I am ` + people.name
}

type Animal struct {
	//name   string //成員重覆且本身沒複寫，出現 ambiguous selector 編譯錯誤
	weight string
}

func (animal Animal) Hello() string {
	return `weight : ` + animal.weight + `kg`
}

type America struct {
	People  //America 繼承 People
	Animal  //America 繼承 Animal, 可以多重繼承 //但如果方法或成員重覆且本身沒複寫的話，就會出現 ambiguous selector 編譯錯誤
	country string
}

//成員與方法是會被繼承下來的，也可以被覆寫：
func (america America) Hello() string {
	return america.People.Hello() + ` and ` + america.Animal.Hello()
}

type Taiwanese struct {
	People  //Taiwanese 繼承 People
	country string
	name    string //成員也可以複寫
}

//成員與方法是會被繼承下來的，也可以被覆寫：
func (taiwanese Taiwanese) Hello() string {
	return `你好！我是` + taiwanese.name
}

func main() {
	miles := Taiwanese{People{`Miles`, 18}, `Taiwan`, `麥爾斯`}
	fmt.Println(miles)         // {{Miles 18} Taiwan 麥爾斯}
	fmt.Println(miles.country) // Taiwan
	fmt.Println(miles.name)    // 麥爾斯
	fmt.Println(miles.age)     // 18
	fmt.Println(miles.Hello()) // 你好！我是麥爾斯
	// 指定結構
	fmt.Println(miles.People.name)    // Miles
	fmt.Println(miles.People.age)     // 18
	fmt.Println(miles.People.Hello()) // Hi! I am Miles

	lion := America{People{`lion`, 3}, Animal{`63.5`}, `USA`}
	fmt.Println(lion)         // {{lion 3} {63.5} USA}
	fmt.Println(lion.country) // USA
	fmt.Println(lion.name)    // lion
	fmt.Println(lion.age)     // 3
	fmt.Println(lion.weight)  // 63.5
	fmt.Println(lion.Hello()) // Hi! I am lion and weight : 63.5kg
	// 指定結構
	fmt.Println(lion.People.name)    // lion
	fmt.Println(lion.People.age)     // 3
	fmt.Println(lion.People.Hello()) // Hi! I am lion
	fmt.Println(lion.Animal.Hello()) // weight : 63.5kg
}
