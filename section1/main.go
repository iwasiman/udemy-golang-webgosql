package main //ファイルが属するパッケージは常に１つ。

// import文を複数書いてもよいが、()にまとめる方がGoのお作法。
import (
	"fmt"
	"time"
)

// Hello Worldを実行するプログラム。
/*
 複数行コメント
*/
// mainパッケージの中のmain関数が最初に実行されるエントリポイントという仕様。
func main() {
	fmt.Println("Hello, Go world!")
	fmt.Println(time.Now())


}