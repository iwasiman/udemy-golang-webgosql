package main //section9
import (
	"fmt"
	"time"
)

// チャネルにデータが入ったらその整数を出力する関数
func receiver(c chan int, chName string) {
	for {
		i := <-c
		fmt.Println("receiver関数 与えたチャネル", chName, "中のi", i)
	}

}

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("チャネルとゴルーチン")

	ch1 := make(chan int)
	ch2 := make(chan int)
	//fmt.Println(<-ch1) 中身がないのでFatal error: all goroutines are asleep - deadlock!

	// ゴルーチンを並行して走らせる
	go receiver(ch1, "ch1")
	go receiver(ch2, "ch2")

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond) // 50ms待つ
		i++
	}
	// これで実行するとiが0-99まで実行。iひとつにつきch1,ch2の順に表示されるとは限らない。
}
