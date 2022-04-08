package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("チャネル select")

	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)
	// ch2 <- "AAA"
	// // ch1の中はまだ空なので受信できず、ここでパニック。プログラム全体が止まってしまう問題あり。
	// v1 := <-ch1 // fatal error: all goroutines are asleep - deadlock!
	// v2 := <-ch2
	// fmt.Println(v1)
	// fmt.Println(v2)

	ch2 <- "AAA" // これだけだと、select文で１回分岐して終了
	ch1 <- 123
	ch2 <- "BBB"
	ch1 <- 456 // これで実行すると、実行のたびにランダムにAAAか123。BBBと456は永遠に来ない。
	select {
	case v1 := <-ch1:
		fmt.Println("ch1から受信", v1)
	case v2 := <-ch2:
		fmt.Println("ch2から受信", v2)
	default:
		// どちらにも当てはまらない ch1,ch2に送信していればここには来ない。
		fmt.Println("どちらでもないぞい")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)
	// 無名関数によるゴルーチン
	go func() {
		for {
			// ch3に入ったらch4に渡す
			i := <-ch3
			fmt.Println("ゴル ch3からch4に渡すよ")
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			// ch4に入ったらch5に渡す
			i2 := <-ch4
			fmt.Println("ゴル ch4からch5に渡すよ")
			ch5 <- i2 - 1
		}
	}()

	n := 0
	// nが増えていく無限ループ。
Label1:
	for {
		select {
		case ch3 <- n:
			fmt.Println("ch3 received n:", n)
			n++
		case i3 := <-ch5:
			fmt.Println("ch5 received i3:", i3)
		default:
			//fmt.Println("どれでもないよん") //1回しかこない?と思ったらランダム
			// ここで終了判定する時は、for文の外に置いたラベルで行う。
			if n > 100 {
				break Label1
			}
		}
		// こちらが普通。
		// if n > 100 {
		// 	break
		// }
	}

}
