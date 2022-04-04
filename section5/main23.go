package main //section5
import (
	"fmt"
)

// 定数は関数の外に書くことが多い。
const Pi = 3.14 // 1文字目大文字でパッケージの外から参照できる

const (
	URL = "http://somewhere.co.jp" // カンマで区切らないのに注意
	SiteName = "Sample Iketeru Site"
)

const (
	A = 1
	B // 1になる
	C // 1になる
	D = "ddd"
	E // ddd になる
	F // ddd になる
)

//var Big int = 92233333333333333333333333333333 //cannot use 9223... (untyped int constant) as int value in variable declaration (overflows)
const Big = 92233333333333333333333333333333 //定数だと型の最大値が関係ない。

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println("定数")
	fmt.Println("定数Pi", Pi)
	// 定数は後から値を変えられない。
	//Pi = 3 // cannot assign to Pi (untyped float constant 3.14)
	fmt.Println("複数の定数を一度に表示", URL, SiteName)
	fmt.Println("初期値を省略した定数を一度に表示", A, B, C, D, E, F)
	fmt.Println("iota識別子で宣言した定数", c0, c1, c2) // 0, 1, 2

}