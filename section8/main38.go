package main //section8
import "fmt"

func main() {
	fmt.Println("制御構文")
	fmt.Println("switch式スイッチ")

	// Goではcase式の中にbreakがいらない。
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1か4だよ") //ここ
	case 3, 4:
		fmt.Println("3か4だよ")
	default:
		fmt.Println("わからん")
	}

	// 変数の局所性を高める
	switch n := 3; n {
	case 1, 2:
		fmt.Println("1か4だよ")
	case 3, 4:
		fmt.Println("3か4だよ") // ここ
	default:
		fmt.Println("わからん")
	}

	n = 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("満たすよ")
	case n > 3 && n < 7:
		fmt.Println("3-7満たすよ") // ここ
	// case 6: 値と混在するとコンパイルエラー
	default:
		fmt.Println("わからん")

	}

}
