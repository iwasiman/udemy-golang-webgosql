package main //section7
import (
	"fmt"
)

func main() {
	fmt.Println("クロージャ")
	f := Later()
	fmt.Println("Later()が入ったfを実行", f("やあ"))  // 出力なし
	fmt.Println("Later()が入ったfを実行", f("おや"))  // 1つの前の やあ
	fmt.Println("Later()が入ったfを実行", f("どうよ")) // 1つ前の おや

	fmt.Println("ジェネレーター")
	// Goにはジェネレーターはないが、クロージャを使って実現可能。
	ints := integers()
	fmt.Println("ジェネレータぽい!", ints()) // 1,2,3,5 ...
	fmt.Println("ジェネレータぽい!", ints())
	fmt.Println("ジェネレータぽい!", ints())
	fmt.Println("ジェネレータぽい!", ints())
	fmt.Println("ジェネレータぽい!", ints())
	otherInts := integers()
	fmt.Println("別ジェネレータぽい!", otherInts()) // 1,2 ...
	fmt.Println("別ジェネレータぽい!", otherInts())

}

// "func(string) string" が戻り値の型。
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store // 前回の値が入る
		store = next
		return s // 前回の値を返す
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
