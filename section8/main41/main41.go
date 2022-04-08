package main //section8
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("制御構文")
	fmt.Println("defer")

	TestDefer() // STARTの後にdefer文が出力。

	// deferで複数処理は無名関数で書く。最後に()がいる このmain関数の最後に実行。
	defer func() {
		fmt.Println("defer処理で出力1")
		fmt.Println("defer処理で出力2")
	}()

	RunDefer() //defer3, defer2, defer1 と逆順に実行

	CreateTestFile() // main41.go と同じ位置にtest.txt作成!

}

func TestDefer() {
	defer fmt.Println("defer処理でENDを出力")

	fmt.Println("関数START")
}

func RunDefer() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
}

func CreateTestFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //deferをよく使うパターン
	file.Write([]byte("Hello, Go world!"))
}
