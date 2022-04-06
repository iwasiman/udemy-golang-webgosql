package main //section11
import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type MyInt int

func (mi MyInt) Print() {
	fmt.Print("このMyInt型の値は", mi)
}

func main() {
	fmt.Println("セクション11: 構造体")
	fmt.Println("struct 独自型")

	var mi MyInt
	fmt.Println("myInt", mi) // 0
	fmt.Printf("%T\n", mi)   // main.MyInt型

	// 他のデータ型とは演算できない。
	//var i = 100
	//mi = mi + i // invalid operation: mi + i (mismatched types MyInt and int)

	// // 特有のメソッドを持った独自型が定義できる。
	mi.Print() // このMyInt型の値は0
}
