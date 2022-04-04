package main //section7
import (
	"fmt"
)

func main() {
	fmt.Println("関数を返す関数")
	f := ReturnFunc()
	f()
	fmt.Println("関数を引数に取る関数")
	CallFunction(func() {
		fmt.Println("無名関数だニャン")
	})
}

// "func()"と書いたところが返り値の型
func ReturnFunc() func() {
	return func() {
		fmt.Println("関数なのだずら～")
	}
}

func CallFunction(f func()) {
	f()
}
