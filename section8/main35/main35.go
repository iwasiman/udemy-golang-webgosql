package main //section8
import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("制御構文")
	fmt.Println("if")
	a := 1
	// GoのIF分の条件は()で囲わない。
	if a == 2 {
		fmt.Println("正しいよ")
	} else if a == 1 {
		fmt.Println("1だよ")
	} else {
		fmt.Println("違うよ")
	}

	// 条件式の前に簡易文を書く
	if b := 100; b == 100 {
		fmt.Println("bは100だよ")
	}
	x := 0
	if x := 2; true {
		fmt.Println("条件が真だよ x:", x) //x:2 このXはIF文の中だけがスコープ。
	}
	fmt.Println("if文の外の x:", x) //x:0

	fmt.Println("エラーハンドリング")
	var s string = "100"
	i, err := strconv.Atoi(s)
	fmt.Println("変換したよ", i, err) // 100, <nil>
	s = "AAA"
	i, err = strconv.Atoi(s)
	fmt.Println("変換したよ", i, err) // 0, strconv.Atoi: parsing "AAA": invalid syntax
	if err != nil {
		fmt.Println(err) // よく使うパターン
	}

}
