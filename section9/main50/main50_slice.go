package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("スライス 可変長引数")

	fmt.Println(Sum(1, 2, 3, 4)) // 10
	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...)) // スライスを渡すときは変数名...　。結果は6

}

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
