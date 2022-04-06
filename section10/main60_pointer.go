package main //section10
import (
	"fmt"
)

func Double(i int) {
	i = i * 2
}
func DoubleForPointer(i *int) {
	*i = *i * 2
}
func DoubleForSlice(s []int) {
	for i, v := range s {
		s[i] = v * 2

	}
}

func main() {
	fmt.Println("セクション10: ポインタ")
	fmt.Println("ポインタ")

	var n int = 100
	fmt.Println("nとアドレス番地", n, &n) // 100 0xc0000a8058 &をつけるとアドレス番地を指す。アドレス演算子

	// 基本型は値渡し、コピーされたnが2倍されて元のnは変わらない。
	Double(n)
	fmt.Println("関数を呼んだ後 nとアドレス番地", n, &n) // かわらない。

	// 型名の前に* をつける。nのアドレスを指す
	var p *int = &n
	// *をつけるのがデリファレンス
	fmt.Println("ポインタのP", p, *p) // p:0xc0000a8058 *p: 100
	*p = 300
	fmt.Println("pの指す先(*p)のnの値を300に変更", n, &n) //n=300 番地は同じ!!!
	n = 500
	fmt.Println("nを500に変更、nを指しているpを出力", p, *p) // 番地は同じ!!! *p=300

	// nのアドレスを渡すとn自身が更新される
	DoubleForPointer(&n)
	fmt.Println("nを出力", n, &n) //n=1000 番地は同じ!!!
	fmt.Println("pを出力", p, *p) // 番地は同じ!!! *p=1000
	// ポインタを渡すときは変数名だけ
	DoubleForPointer(p)
	fmt.Println("ポインタを使って更新 nを出力", n, &n) //n=2000 番地は同じ!!!
	fmt.Println("ポインタを使って更新 pを出力", p, *p) // 番地は同じ!!! *p=2000

	var sl []int = []int{1, 2, 3}
	fmt.Println("スライスを出力", sl)       // 1 2 3
	DoubleForSlice(sl)               // スライスは元から参照渡しなのでポインタ不要
	fmt.Println("関数通過後 スライスを出力", sl) // 2 4 6

}
