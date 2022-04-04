package main //section8
import "fmt"

func main() {
	fmt.Println("制御構文")
	fmt.Println("ラベル付きFor")

Loop:
	for {
		for {
			for {
				fmt.Println("スタート")
				break Loop
			}
			//fmt.Println("処理しない") //unreachable
		}
		//fmt.Println("処理しない") //unreachable
	}
	fmt.Println("END")

Loop2:
	// 0 1
	// 1 1
	// 1 1 で終わる。
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop2
			}
			fmt.Println(i, j)
		}
		fmt.Println("処理しないよ")
	}

}
