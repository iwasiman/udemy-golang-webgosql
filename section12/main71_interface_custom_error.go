package main //section12
import (
	"fmt"
)

// Goの定義ではこうなっている。
/*
type error interface {
	Error() string
}
*/

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return "エラー内容だずら:" + e.Message
}

func RaiseError() error {
	// MyError構造体のアドレスを返す。
	return &MyError{Message: "カスタムエラー発生", ErrCode: 101}
}

// これはコンパイルエラーで定義できない。
// func RaiseError2() MyError {
// 	// MyError構造体のアドレスを返す。
// 	return &MyError{Message: "カスタムエラー発生", ErrCode: 101}
// }

func main() {
	fmt.Println("セクション12: interface")
	fmt.Println("interface カスタムエラー")

	err := RaiseError()
	fmt.Println(err)         // エラー内容だずら:カスタムエラー発生 が出る。なぜ? *errはコンパイルエラー。
	fmt.Println(err.Error()) // エラー内容だずら:カスタムエラー発生

	//fmt.Println(err.Message) // ここではerrはError型なので、フィールドにはアクセスできない。
	e, ok := err.(*MyError)
	if ok {
		// ちゃんと e. でVSCodeの補完が出る。すごい。
		fmt.Println("MyError型のフィールドにアクセス", e.Message, e.ErrCode)
	}

}
