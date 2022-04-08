package main //section8
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("制御構文")
	fmt.Println("go 並行処理")

	// 以下の実装だとSub Loopしか走らない、Main Loopに永遠に届かない
	//sub()
	// 以下の実装だとゴルーチンでSubとMainが同時に永遠に走る。すごいぞGo!
	go sub()
	go sub()
	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Microsecond)
	}

}

func sub() {
	// 無限ループさせる
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Microsecond)
	}
}
