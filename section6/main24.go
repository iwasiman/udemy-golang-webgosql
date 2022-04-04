package main //section6
import (
	"fmt"
)

func main() {
	fmt.Println("算術演算子")
	fmt.Println(1 + 2)
	fmt.Println(1 - 2) // -1
	fmt.Println(1 * 2) // 2
	fmt.Println(1 / 2) // 0
	fmt.Println(9 % 4) // 余り1

	n := 0
	n += 4 // n = n + 4
	fmt.Println("nは", n)
	n++
	fmt.Println("n++すると", n)
	n--
	fmt.Println("n--すると", n)
	s := "ABC"
	s += "def" // s = s + になる
	fmt.Println("文字列s", s)

	fmt.Println("比較演算子")
	fmt.Println("==", 1 == 1)
	fmt.Println("==", 1 == 2)
	fmt.Println("<=", 4 <= 8)
	fmt.Println("真偽値", true == true) // true
	fmt.Println("真偽値", true != false) // true

	fmt.Println("論理演算子")
	fmt.Println("右辺が偽", true && false == true) // false
	fmt.Println("右辺が真", true && false == false) // true
	fmt.Println("どちらかが真なら真", true || false == true) // true

	fmt.Println("NOT演算子", !true) // false
}