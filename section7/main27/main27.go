package main //section7
import (
	"fmt"
)

func main() {
	fmt.Println("関数")
	i := Plus(1, 2)
	fmt.Println("関数Plusを実行", i)
	i2, _ := Div(9, 4) // 戻り値を捨てる場合
	fmt.Println("関数Divを実行", i2)
	i3, i4 := Div(9, 4)
	fmt.Println("関数Divを実行", i3, i4)
	fmt.Println("関数Doubleを実行", Double(200))
	fmt.Println("関数NoReturnを実行") // ここに, NoReturn()と書くとエラー。
	NoReturn()

	f := func(x, y int) int {
		return x + y
	}
	fmt.Println("無名関数が入った変数Fを実行", f(1, 2)) // 3
	iNoName := func(x, y int) int {
		return x + y
	}(1, 2) // こういう風に()で無名関数コール後も入れられる
	fmt.Println("無名関数を実行した結果が入ったiNoName", iNoName) // 3

}

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値を変数に取る場合
func Double(price int) (result int) {
	result = price * 2
	return // 戻り値に変数名を宣言しているので、return result を省略できる。
}

// Goにはvoidがない。
func NoReturn() {
	fmt.Println("何も返さないよ")
}
