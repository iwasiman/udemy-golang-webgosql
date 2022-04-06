package main //section9
import (
	"fmt"
	"time"
)

// チャネルにデータが入ったらその整数を出力する関数
func receiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("receiver関数 チャネル名", name, "中のi", i)
	}
	fmt.Println("receiver関数 END!")

}

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("チャネルとクローズ")

	// makeで作ったチャネルはオープン状態。そこをクローズできる。
	ch1 := make(chan int, 2) // バッファ2
	close(ch1)
	//ch1 <- 1 // panic: send on closed channel クローズ済みのチャネルに送信はできない。
	fmt.Println(<-ch1) // 0 クローズ済みのチャネルからの送信はできる。

	i, ok := <-ch1
	fmt.Println(i, ok) // 0 false チャネル内のバッファが空 && クローズされているとfalseになる。
	// ここでiが0なのはチャネルから渡ってきたものでなく、int型の初期値だ。

	ch2 := make(chan int, 2) // バッファ2
	ch2 <- 999
	i21, ok21 := <-ch2
	fmt.Println("ch2だよ", i21, ok21) // 999 true
	//i21, ok21 = <-ch2 // もう中身がないからfatal error: all goroutines are asleep - deadlock!
	//fmt.Println("ch2だよ", i21, ok21)

	ch3 := make(chan int, 2)
	go receiver("ゴルチ1", ch3)
	go receiver("ゴルチ2", ch3)
	go receiver("ゴルチ3", ch3)
	i3 := 0
	for i3 < 100 {
		// こう書くと、i3:0-99の間ほとんどゴルチ1が処理、2と3はひとつだけ。講義とちょっと違う。
		ch3 <- i3
		i3++
	}
	close(ch3)
	fmt.Println("ch3クローズ")
	time.Sleep(3 * time.Second)

}
