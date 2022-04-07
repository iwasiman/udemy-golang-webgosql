package main //section12
import (
	"fmt"
)

// 特定メソッドがあるインターフェースを定義
type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

// 構造体のフィールドにアクセスする際、Goは変数でもポインタ変数でも同じようにp.Nameのようにしていける。
func (p *Person) ToString() string {
	return fmt.Sprintf("Name ::: %v, Age ::: %v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number ::: %v, Model ::: %v", c.Number, c.Model)
}

func main() {
	fmt.Println("セクション12: interface")
	fmt.Println("interface 異なる型に共通の性質を付与する")
	// スライスでは異なる構造体は一緒に使えないが、インターフェスを使うことで一緒にできる。
	vs := []Stringfy{
		&Person{Name: "Troo", Age: 21},
		&Car{Number: "No-123", Model: "Ultra-Turbo typeF"},
	}
	// この時アドレス演算子&が必須の模様。
	// &を外すと cannot use (Car literal) (value of type Car) as Stringfy value in array or slice literal: Car does not implement Stringfy (method ToString has pointer receiver)

	// for文内で共通のメソッドを呼べる。
	for _, v := range vs {
		fmt.Println(v.ToString())
	}
	// Name ::: Troo, Age ::: 21
	// Number ::: No-123, Model ::: Ultra-Turbo typeF
}
