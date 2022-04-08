package main //section8
import (
	"fmt"
)

func main() {
	fmt.Println("制御構文")
	fmt.Println("panicとrecover")
	// これがしてあるとpanicで停止せず続行する。
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("deferしておいたrecoverが実行されました。戻り値：", x)
		}
	}()
	panic("runtime error的なパニック!") // 例外をスローする的なもの。強制的にプログラム停止
	//fmt.Println("main() スタート") //unreachable

}
