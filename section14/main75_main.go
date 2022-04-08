package main //section14
import (
	"fmt"
	alib "udemy-golang-webgosql/section14/alib" // VSCode+拡張だと自動でパスが挿入、手動で別名が必要。
)

// ここでファイル名に_testをつけるとテストファイルと認識されてしまうので注意。

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("セクション14: テスト")
	fmt.Println("テスト")

	fmt.Println("1で実行", IsOne(1))
	fmt.Println("0で実行", IsOne(0))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("alib.Averageを実行", alib.Average(s))
}
