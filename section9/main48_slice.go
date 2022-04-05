package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("スライス copy")

	sl1 := []int{100, 200}
	sl2 := sl1
	sl1[0] = 99
	fmt.Println(sl1, sl2) // []99 100] [99 100]で参照型は両方変わる。

	copy1 := []int{1, 2, 3, 4, 5}
	copy2 := make([]int, 5, 10)
	fmt.Println(copy1, copy2)    // [1 2 3 4 5] [0 0 0 0 0]
	n := copy(copy2, copy1)      // nはコピーに成功した数
	fmt.Println(n, copy1, copy2) // 5 [1 2 3 4 5] [1 2 3 4 5]

}
