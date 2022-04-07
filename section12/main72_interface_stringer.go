package main //section12
import (
	"fmt"
)

/*
type Stringer interface {
	String() string
}
*/

type Point struct {
	A int
	B string
}

// 構造体とメソッドを関連付ける
func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	fmt.Println("セクション12: interface")
	fmt.Println("interface Stringer")

	p := &Point{100, "ABC"}
	p2 := Point{200, "DEF"}
	fmt.Println(p)  // &{100 ABC}
	fmt.Println(p2) // {200 DEF}
	//紐づいたメソッド呼び出しは普通の変数、ポインタ双方で気にせず同じように呼べる。
	fmt.Println(p.String())  // <<100, ABC>>
	fmt.Println(p2.String()) // <<200, DEF>>

}
