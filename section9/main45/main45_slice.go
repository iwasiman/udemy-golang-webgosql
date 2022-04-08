package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("スライス")

	var sl []int    //[]に要素数を書かないのがスライス。
	fmt.Println(sl) // []
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2) // [100 200]

	sl3 := []string{"A", "B"}
	fmt.Println(sl3) // [A B]
	sl4 := make([]int, 5)
	fmt.Println(sl4) // [0 0 0 0 0]

	sl2[0] = 999
	fmt.Println(sl2) // [999 200]
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[2:4])            // [3 4] インデックス0始まり2番目から4の手前まで。
	fmt.Println(sl5[:2])             // [1, 2] インデックス0始まり2番目の手前まで。
	fmt.Println(sl5[2:])             // [3 4 5] インデックス0始まり2番目~最後まで。
	fmt.Println(sl5[:])              // 全部
	fmt.Println(sl5[len(sl5)-1])     // 最後の 5
	fmt.Println(sl5[0 : len(sl5)-1]) // 最初から最後の5を除く [1 2 3 4]

}
