package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("チャネル for")
	// チャネルに対してもfor文が使える。

	ch1 := make(chan int, 3)
	// ch1 <- 1
	// ch1 <- 2
	// ch1 <- 3
	// // 1 2 3 を出して fatal error: all goroutines are asleep - deadlock!になる
	// for i := range ch1 {
	// 	fmt.Println(i)
	// }

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1) // クローズしてから出力すると 1 2 3 で正常。
	for i := range ch1 {
		fmt.Println(i)
	}
}
