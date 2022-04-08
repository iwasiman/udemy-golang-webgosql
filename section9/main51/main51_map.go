package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("マップ")

	var m = map[string]int{"a": 100, "b": 200}
	fmt.Println(m)                           // map[a:100 b:200]
	m2 := map[string]int{"a": 100, "b": 200} // 暗黙宣言
	m3 := map[int]string{
		1: "aaa",
		2: "bbb", // このカンマ必要
	}
	fmt.Println(m2, m3)

	m4 := make(map[int]string)
	m4[1] = "ja"
	m4[2] = "en"
	fmt.Println(m4, m4[1], m4[2], m3[3]) // map[1:ja 2:en] ja en (空文字)
	s, ok := m4[1]
	if !ok {
		fmt.Println("err!")
	}
	fmt.Println(s, ok) // ja true
	s, ok = m4[9999]
	if !ok {
		fmt.Println("err!") //出る
	}
	fmt.Println(s, ok) // (空文字) false

	delete(m4, 1)
	fmt.Println(m4, len(m4)) // map[2:en] 1

}
