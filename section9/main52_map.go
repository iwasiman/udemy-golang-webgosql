package main //section9
import (
	"fmt"
)

func main() {
	fmt.Println("セクション9 参照型")
	fmt.Println("マップ for")

	m := map[string]int{
		"Apple":  99,
		"Banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
