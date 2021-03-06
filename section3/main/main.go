package main //section3
import "fmt"

// 変数

var externali5 = 10

func main() {
	// 明示的な定義
	// var 変数名 型=値
	var i int = 100
	fmt.Println(i)
	var s string = "Hello?"
	fmt.Println(s)
	var t, f bool = true, false // まとめて代入できる
	fmt.Println(t, f)
	var (
		i2 int    = 200
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

	fmt.Printf("%T\n", seisuI2)        // %Tが型を表示。int64と出る。
	fmt.Printf("%T\n", int32(seisuI2)) // 型変換してint32になる。

	//浮動小数点型
	var fl64 float64 = 2.4
	fmt.Println(fl64)
	fl := 3.2                                          // 自動でfloat64型になる
	fmt.Println(fl64 + fl)                             //5.6
	var fl32 float32 = 1.2                             // float32は現在はあまり使わない。
	fmt.Printf("%T, %T,%T\n", fl64, fl, fl32)          // float64, 64, 32
	fmt.Printf("%T, %T,%T\n", fl64, fl, float64(fl32)) // float64, 64, 64

	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := zero / zero
	fmt.Println(pinf, ninf, nan) // +Inf(正の無限大), -Inf(負の無限大), NaN(非数)

	// 論理値形
	var boolT, boolF bool = true, false
	fmt.Println("論理値型", boolT, boolF)
	// 文字列型 ダブルクォート必須。
	var sString string = "はろはろ～"
	fmt.Println("文字列型", sString)
	fmt.Printf("%T\n", sString)
	var s300 string = "200"
	fmt.Println("文字列型", s300)
	fmt.Printf("%T\n", s300)
	// 複数行はバッククォート
	fmt.Println(`test
	test 
	test`)
	// ダブルクォートを出す
	fmt.Println("\"")
	fmt.Println(`"`)
	// 文字列を[0]で出すとその中身ではなくバイト配列。また全角文字でやると復元されない。
	var sString2 string = "Hello"
	fmt.Println("文字列を取り出すよ ", sString2[0])                     // 72
	fmt.Println("文字列を取り出すよ string()で囲う ", string(sString2[0])) // H

}

func outer() {
	var s4 string = "outerだよ"
	fmt.Println(s4)
}

/*
階層について
udemy-golang-webgosql/
  section3/
    main.go
		main16.go
		main20.go
		main22.go

この構造でも、最初は
> cd udemy-golang-webgosql/section3
> go run main16.go
これでmain16.goの中のmain関数が問題なく動く。

しかし、go modules を使って >go mod init でgo.modファイルを作って管理しだすと、
途端にプロジェクトの階層構造を認識。
- main パッケージが複数ある、main関数が複数ある
- Userなどの同じ構造体が複数回定義されている
などなどでエラーが多数出る。

解決法方法は以下。

udemy-golang-webgosql/
  section3/
    main/
		  main.go
		main16/
		  main16.go この時、package main のまま
		main20/
		  main20.go
		main22/
			main22.go

Goの言語仕様ではディレクトリ名とパッケージ名が等しくなくても動く。PHPと同じ。
> cd udemy-golang-webgosql/section3/main16
> go run main16.go

これで、mainパッケージやmain関数がそれぞれのファイルごとにあっても問題なく動作。


*/
