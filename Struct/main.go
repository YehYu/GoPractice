package main

import "fmt"

type Point struct {
	X, Y int
}

//鏈狀參考
type Node struct {
	point *Point
	next  *Node
}

//如果你建立了一個結構的實例，並將之指定給另一個結構變數，那麼會進行值域的複製
func changeX(point Point) {
	point.X = 20
}

func changeXP(point *Point) {
	point.X = 20
}

func main() {
	// $T{} 的寫法與 new(T) 是等效的 使用$T{}可以同時指定結構的值域。
	point1 := &Point{10, 20}
	point2 := new(Point)
	point2.X = 10
	point2.Y = 30
	fmt.Println(point1) // &{10 20}
	fmt.Println(point2) // &{10 30}
	changeX(*point1)
	changeXP(point2)
	fmt.Println(point1) // &{10 20}
	fmt.Println(point2) // &{20 30}

	//鏈狀參考
	node := new(Node)

	node.point = &Point{10, 20}
	node.next = &Node{point: new(Point)}
	node.next.point.X = 10
	node.next.point.Y = 30

	fmt.Println(node.point)      // &{10 20}
	fmt.Println(node.next.point) // &{10 30}
}
