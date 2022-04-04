package main //section8
import (
	"fmt"
)

func main() {
	fmt.Println("制御構文")
	fmt.Println("for")
	// 条件を書かないとwhile(true)的な無限ループ。
	i := 0
	for {
		i++
		if i == 5 {
			break
		}
		fmt.Println("むげんるーぷ～")
	}

	// 条件付きfor文 Python的
	point := 0
	for point < 10 {
		fmt.Println("条件付きfor文", point)
		point++
	}

	// 「古典的」for文
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("古典的for文", i)
	}

	// 配列でfor文
	arr := [3]int{1, 2, 3}
	// 1,2,3と3回回る
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 範囲式for range識別子を使う。
	for i, v := range arr {
		//変数名を使わないときは_ で
		fmt.Println("範囲式 index:", i, "値", v)
	}

	// スライスをちょとだけ 配列と同じ動き。
	sl := []string{"Python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println("範囲式でスライス index:", i, "値", v)
	}

	// マップをちょっとだけ
	m := map[string]int{"Python": 100, "PHP": 50, "Go": 200}
	for k, v := range m {
		fmt.Println("範囲式でマップ キー:", k, "値", v)
	}

}
