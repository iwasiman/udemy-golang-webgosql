package main //section8
import (
	"fmt"
)

// 特殊な関数で、そのパッケージの中で最初に呼ばれる。
func init() {
	fmt.Println("initだよ")
}

// 複数回定義すると定義順に実行。しかし分けるメリットはない。
func init() {
	fmt.Println("initだよ2")
}

func main() {
	fmt.Println("制御構文")
	fmt.Println("init")
}
