package main

import (
	"fmt"
)

type Account struct {
	id      string
	name    string
	balance float64
}

type Point struct {
	x, y int
}

//移至方法名稱之前成為函式呼叫的接收者（Receiver） 類似物件導向的method(Account.Deposit) 可以傳值或是傳參考
func (ac *Account) Deposit(amount float64) {
	if amount <= 0 {
		panic("必須存入正數")
	}
	ac.balance += amount
}

//Go 語言中不允許方法重載（Overload），因此，對於以下的程式，是會發生 String 重複宣告的編譯錯誤：
/*
func String(account *Account) string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        account.id, account.name, account.balance)
}


func String(point *Point) string {    // String redeclared in this block 的編譯錯誤
    return fmt.Sprintf("Point{%d %d}", point.x, point.y)
}
*/

//若是將函式定義為方法，就不會有這個問題，Go 可以從方法的接收者辨別，該使用哪個 String 方法
func (ac *Account) String() string {
	return fmt.Sprintf("Account{%s %s %.2f}",
		ac.id, ac.name, ac.balance)
}

func (p *Point) String() string {
	return fmt.Sprintf("Point{%d %d}", p.x, p.y)
}

func main() {
	account := &Account{"1234-5678", "Justin Lin", 1000}
	point := &Point{20, 30}
	account.Deposit(500)
	fmt.Println(account.balance)  // 1500
	fmt.Println(account.String()) // Account{"1234-5678", "Justin Lin", 1500}
	fmt.Println(point.String())   //Point{20 30}

	//方法作為值
	deposit := (*Account).Deposit
	deposit(account, 500)
	fmt.Println(account.balance) // 2000

	//另一個取得方法的方式是方法值（Method value），這會保有取得方法當時的接收者
	accountDeposit := account.Deposit
	accountDeposit(500)
	fmt.Println(account.balance) // 2500

}
