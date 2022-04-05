package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("スライス append make len cap")
	sl := []int{100, 200}
	fmt.Println(sl)
	sl = append(sl, 300)
	fmt.Println(sl) // [100 200 300]
	sl = append(sl, 400, 500, 600)
	fmt.Println(sl) // [100 200 300 400 500 600]
	sl2 := make([]int, 5)
	fmt.Println(sl2, "要素数", len(sl2), "容量", cap(sl2)) // [0 0 0 0 0] 5 5
	sl3 := make([]int, 5, 10)
	fmt.Println(sl3, "要素数", len(sl3), "容量", cap(sl3)) // [0 0 0 0 0] 5 10

}
