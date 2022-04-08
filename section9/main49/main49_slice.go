package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("スライス for")
	sl := []string{"A", "b", "C"}
	for i, v := range sl {
		fmt.Println(i, v)
		// 0 A
		// 1 b
		// 2 C
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("古典的forループ", sl[i])
		// A
		// b
		// C
	}

}
