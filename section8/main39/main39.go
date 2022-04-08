package main //section8
import "fmt"

func main() {
	fmt.Println("制御構文")
	fmt.Println("switch型スイッチ")
	anything("aaa") // aaa
	anything(1)     // 1 ただし型は失われているので演算はできない。

	// 型アサーションは２つの返り値を使えば2つ目に変換できたかが入る。
	var x interface{} = 3
	i := x.(int)       // 「型アサーション」で型を復元する
	fmt.Println(i * 5) // 15
	//f := x.(float64)
	//fmt.Println(f) // panic: interface conversion: interface {} is int, not float64
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64) // 0, false こうするとパニックにならない。

	fmt.Println("hantei() 実行")
	hantei(nil)
	hantei(12)
	hantei("aaa")
	hantei(3.14444444)

	fmt.Println("hanteiByCase() 実行")
	hanteiByCase(nil)
	hanteiByCase(12)
	hanteiByCase("aaa")
	hanteiByCase(3.14444444)
}

func anything(a interface{}) {
	fmt.Println(a)
}

// 型の判定をif文で
func hantei(x interface{}) {
	if x == nil {
		fmt.Println("nilだよ")
	} else if i, isInt := x.(int); isInt {
		fmt.Println("引数はint型だよ", i)
	} else if s, isString := x.(string); isString {
		fmt.Println("引数はstring型だよ", s)
	} else {
		fmt.Println("正体不明!")
	}
}

// 型の判定をswitch型スイッチで
func hanteiByCase(x interface{}) {
	// xを標準出力もできる。こちらは引数の値を出力。
	// vを出力すると、判定された後の値。
	// typeは変数名でなく識別子のようなので、出力はできない。
	switch v := x.(type) {
	case int:
		fmt.Println("引数はint型だよ", v)
	case string:
		fmt.Println("引数はstring型だよ", v)
	case nil:
		fmt.Println("nilだよ", v)
	default:
		fmt.Println("正体不明!", v)
	}
}
