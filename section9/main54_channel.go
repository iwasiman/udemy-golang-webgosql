package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("チャネル")
	// 複数のゴルーチン間でのデータの受け渡しをする為に設計されたデータ構造
	// 「キュー」的なもの。FIFOで先入先出

	var ch1 chan int // 両方

	// var ch2 <-chan int //受信専用チャネル
	// var ch3 chan<- int //送信専用

	ch1 = make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int, 5)

	fmt.Println("容量", cap(ch1), cap(ch2), cap(ch3)) // 0, 0, 5

	ch3 <- 1000
	fmt.Println("len実行で要素数：", len(ch3)) // 1
	ch3 <- 2000
	ch3 <- 3000
	fmt.Println("len実行で要素数：", len(ch3))                  // 3になる
	i1 := <-ch3                                          // 矢印でデータを送信
	fmt.Println("チャネルからデータを渡す：", i1, "ch3要素数", len(ch3)) //1000 2
	i2 := <-ch3
	fmt.Println("チャネルからデータを渡す：", i2, "ch3要素数", len(ch3))    //2000 1に減る
	fmt.Println("チャネルからデータを渡す：", <-ch3, "ch3要素数", len(ch3)) //3000 0に減る

	//ch3は容量5
	ch3 <- 1
	fmt.Println("チャネルへデータを渡す ch3要素数", len(ch3)) // 1
	ch3 <- 2
	fmt.Println("チャネルへデータを渡す ch3要素数", len(ch3)) // 2
	ch3 <- 3
	fmt.Println("チャネルへデータを渡す ch3要素数", len(ch3)) // 3
	ch3 <- 4
	fmt.Println("チャネルへデータを渡す ch3要素数", len(ch3)) // 4
	ch3 <- 5
	fmt.Println("チャネルへデータを渡す ch3要素数", len(ch3)) // 5
	//ch3 <- 6 //atal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch3)                          //一つ送信するので容量4に減る
	fmt.Println("チャネルから１つ送信 ch3要素数", len(ch3))  // 4
	ch3 <- 6                                    // これは貯めて大丈夫。
	fmt.Println("チャネルへデータを渡す ch3要素数", len(ch3)) //5

}
