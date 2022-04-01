package main //section3
import "fmt"

// 変数

var externali5 = 10
func main() {
	// 明示的な定義
	// var 変数名 型=値
	var i int = 100
	fmt.Println(i);
	var s string = "Hello?"
	fmt.Println(s);
	var t, f bool = true, false // まとめて代入できる
	fmt.Println(t, f)
	var (
		i2 int = 200;
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)
	var i3 int
	var s3 string
	fmt.Println(i3, s3) // 代入しないと初期値0、文字列は空文字
	i3 = 300
	s3 = "代入するよ"
	fmt.Println(i3, s3) // 変わる

	// 暗黙的な定義 変数名 := 値
	// 関数内でだけ定義できる。varで明示的な定義はできる。
	i4 := 400
	fmt.Println(i4)
	i4 = 500
	fmt.Println(i4)
	//i4 := 999 //no new variables on left side of :=
	//i4 = "文字列にしよう" //cannot use "文字列にしよう" (untyped string constant) as int value in assignment

	fmt.Println(externali5) // できる
	outer()
	//fmt.Println(s4) // undeclared name: s4  スコープ外だから。

	// 変数は宣言したら使うのが必須。面白い。
	//var sNotUserd string = "つかわれないやつ" //sNotUserd declared but not used

	// 整数型
	// int8、int16、int32, int64
	var seisuI int = 100
	var seisuI2 int64 = 200
	fmt.Println(seisuI + 50)
	//fmt.Println(seisuI + seisuI2) //invalid operation: seisuI + seisuI2 (mismatched types int and int64)

	fmt.Printf("%T\n", seisuI2) // %Tが型を表示。int64と出る。
	fmt.Printf("%T\n", int32(seisuI2)) // 型変換してint32になる。

	//浮動小数点型
	var fl64 float64 = 2.4
	fmt.Println(fl64)
	fl:=3.2 // 自動でfloat64型になる
	fmt.Println(fl64 + fl) //5.6
	var fl32 float32 = 1.2 // float32は現在はあまり使わない。
	fmt.Printf("%T, %T,%T\n", fl64, fl, fl32) // float64, 64, 32
	fmt.Printf("%T, %T,%T\n", fl64, fl, float64(fl32)) // float64, 64, 64

	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := zero / zero
	fmt.Println(pinf, ninf, nan) // +Inf(正の無限大), -Inf(負の無限大), NaN(非数)

}


func outer() {
	var s4 string = "outerだよ"
	fmt.Println(s4)
}