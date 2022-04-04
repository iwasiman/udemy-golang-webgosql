package main //section3
import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("型変換")
	var i int = 1
	fl64 := float64(i)
	fmt.Println("intからfloatに変換", fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Printf("float型からint型に変換したi2 = %T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Println("floatからintに変換", i3) // 端数切捨てで10になる野に注意。
	fmt.Printf("i3 = %T\n", i3)

	var s100 string = "100" // string型
	i100, _ := strconv.Atoi(s100)
	// 文字列から数値への変換は標準パッケージを使う。
	fmt.Println("文字列からintに変換", i100) //
	fmt.Printf("i100 = %T\n", i100)

	i101, err := strconv.Atoi("文字列だよ")
	if err != nil {
		fmt.Println(err) // strconv.Atoi: parsing "文字列だよ": invalid syntax
	}
	fmt.Println("文字列からintに変換", i101) //初期値の0のまま

	var i200 int = 200
	s2 := strconv.Itoa(i200)
	fmt.Println("intから文字列に変換", s2) //
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println("文字列hのバイト配列", b) // [72 101 108 108 111 32 87 111 114 108 100]
	h2 := string(b)
	fmt.Println("バイト配列をstring()で戻す", h2) // Hello World




}